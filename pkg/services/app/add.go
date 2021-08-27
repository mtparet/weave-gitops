package app

import (
	"bytes"
	"context"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/fluxcd/go-git-providers/github"
	"github.com/fluxcd/go-git-providers/gitprovider"
	"github.com/weaveworks/weave-gitops/pkg/git"
	"github.com/weaveworks/weave-gitops/pkg/gitproviders"
	"github.com/weaveworks/weave-gitops/pkg/utils"

	wego "github.com/weaveworks/weave-gitops/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

type ConfigType string

type ConfigMode string

type ResourceKind string

type ResourceRef struct {
	kind           ResourceKind
	name           string
	repositoryPath string
}

type AppResourceInfo struct {
	wego.Application
	clusterName string
	targetName  string
}

const (
	ConfigTypeUserRepo ConfigType = ""
	ConfigTypeNone     ConfigType = "NONE"

	ConfigModeClusterOnly  ConfigMode = "clusterOnly"
	ConfigModeUserRepo     ConfigMode = "userRepo"
	ConfigModeExternalRepo ConfigMode = "externalRepo"

	ResourceKindApplication    ResourceKind = "Application"
	ResourceKindSecret         ResourceKind = "Secret"
	ResourceKindGitRepository  ResourceKind = "GitRepository"
	ResourceKindHelmRepository ResourceKind = "HelmRepository"
	ResourceKindKustomization  ResourceKind = "Kustomization"
	ResourceKindHelmRelease    ResourceKind = "HelmRelease"

	WeGOAppIdentifierLabelKey = "wego.weave.works/app-identifier"
)

type AddParams struct {
	Dir                        string
	Name                       string
	Url                        string
	Path                       string
	Branch                     string
	PrivateKey                 string
	DeploymentType             string
	Chart                      string
	SourceType                 wego.SourceType
	AppConfigUrl               string
	Namespace                  string
	DryRun                     bool
	AutoMerge                  bool
	GitProviderToken           string
	HelmReleaseTargetNamespace string
}

const (
	DefaultPath           = "./"
	DefaultBranch         = "main"
	DefaultDeploymentType = "kustomize"
)

// Three models:
// --app-config-url=none
//
// - Source created for user repo (GitRepository or HelmRepository)
// - app.yaml created for app
// - HelmRelease or Kustomize created for app dir within user repo
// - app.yaml, Source, Helm Release or Kustomize applied directly to cluster
//
// --app-config-url=<URL>
//
// - Separate GOAT repo
// - Source created for GOAT repo
// - Kustomize created for targets/<target name> directory in GOAT repo
// - Kustomize created for apps/<app name> directory within GOAT repo
// - Source, Kustomizes applied directly to cluster
// - app.yaml created for app
// - app.yaml placed in apps/<app name>/app.yaml in GOAT repo
// - Source created for user repo (GitRepository or HelmRepository)
// - User repo Source placed in targets/<target name>/<app-name>/<app name>-gitops-runtime.yaml in GOAT repo
// - HelmRelease or Kustomize referencing user repo source created for user app dir within user repo
// - User app dir HelmRelease or Kustomize placed in targets/<target name>/<app name>/<app name>-gitops-runtime.yaml in GOAT repo
// - PR created or commit directly pushed for GOAT repo
//
// --app-config-url="" (default)
//
// - Source created for user repo (GitRepository only)
// - Kustomize created for .wego/targets/<target name> directory in user repo
// - Kustomize created for .wego/apps/<app name> directory within user repo
// - Source, Kustomizes applied directly to cluster
// - app.yaml created for app
// - app.yaml placed in apps/<app name>/app.yaml in .wego directory within user repo
// - HelmRelease or Kustomize referencing user repo source created for app dir within user repo
// - User app dir HelmRelease or Kustomize placed in targets/<target name>/<app name>/<app name>-gitops-runtime.yaml in .wego
//   directory within user repo
// - PR created or commit directly pushed for user repo

func (a *App) Add(params AddParams) error {
	ctx := context.Background()

	gitProvider, err := a.GitProviderFactory(params.GitProviderToken)
	if err != nil {
		return err
	}

	params, err = a.updateParametersIfNecessary(gitProvider, params)
	if err != nil {
		return fmt.Errorf("could not update parameters: %w", err)
	}

	if params.SourceType != wego.SourceTypeHelm && params.PrivateKey != "" {
		err = a.Git.ValidateAccess(ctx, params.Url, params.Branch)
		if err != nil {
			return fmt.Errorf("error validating access for app %s. %w", params.Url, err)
		}
	}

	a.printAddSummary(params)

	if err := IsClusterReady(a.Logger, a.Kube); err != nil {
		return err
	}

	clusterName, err := a.Kube.GetClusterName(ctx)
	if err != nil {
		return err
	}

	info := getAppResourceInfo(makeWegoApplication(params), clusterName)

	appHash := info.getAppHash()

	apps, err := a.Kube.GetApplications(ctx, params.Namespace)
	if err != nil {
		return err
	}
	for _, app := range apps {
		existingHash := getAppResourceInfo(app, clusterName).getAppHash()

		if appHash == existingHash {
			return fmt.Errorf("unable to create resource, resource already exists in cluster")
		}
	}

	var secretRef string
	if wego.SourceType(params.SourceType) == wego.SourceTypeGit {
		secretRef, err = a.createAndUploadDeployKey(info, params.DryRun, info.Spec.URL, gitProvider)
		if err != nil {
			return fmt.Errorf("could not generate deploy key: %w", err)
		}
	}

	switch strings.ToUpper(info.Spec.ConfigURL) {
	case string(ConfigTypeNone):
		return a.addAppWithNoConfigRepo(info, params.DryRun, secretRef, appHash)
	case string(ConfigTypeUserRepo):
		return a.addAppWithConfigInAppRepo(info, params, gitProvider, secretRef, appHash)
	default:

		// TODO: @jpellizzari As of https://github.com/weaveworks/weave-gitops/pull/587,
		// we are not using deploy keys for external repos yet.
		// Followup created here: https://github.com/weaveworks/weave-gitops/issues/592
		gitClient, err := a.temporaryGitClientFactory(a.Osys, params.PrivateKey)
		if err != nil {
			return fmt.Errorf("error selecting auth method for external config repo: %w", err)
		}
		// Overriding the internal git service
		a.Git = gitClient
		return a.addAppWithConfigInExternalRepo(info, params, gitProvider, secretRef, appHash)
	}
}

func (a *App) printAddSummary(params AddParams) {
	a.Logger.Println("Adding application:\n")
	a.Logger.Println("Name: %s", params.Name)
	a.Logger.Println("URL: %s", params.Url)
	a.Logger.Println("Path: %s", params.Path)
	a.Logger.Println("Branch: %s", params.Branch)
	a.Logger.Println("Type: %s", params.DeploymentType)

	if params.Chart != "" {
		a.Logger.Println("Chart: %s", params.Chart)
	}

	a.Logger.Println("")
}

const maxKubernetesResourceNameLength = 63

func (a *App) updateParametersIfNecessary(gitProvider gitproviders.GitProvider, params AddParams) (AddParams, error) {
	params.SourceType = wego.SourceTypeGit

	// making sure the config url is in good format
	if strings.ToUpper(params.AppConfigUrl) != string(ConfigTypeNone) &&
		strings.ToUpper(params.AppConfigUrl) != string(ConfigTypeUserRepo) {
		params.AppConfigUrl = utils.SanitizeRepoUrl(params.AppConfigUrl)
	}

	switch {
	case params.Chart != "":
		params.SourceType = wego.SourceTypeHelm
		params.DeploymentType = string(wego.DeploymentTypeHelm)
		params.Path = params.Chart
		if params.Name == "" {
			if nameTooLong(params.Chart) {
				return params, fmt.Errorf("chart name %q is too long to use as application name; please specify name with '--name'", params.Chart)
			}

			params.Name = params.Chart
		}
		if params.Url == "" {
			return params, fmt.Errorf("--url must be specified for helm repositories")
		}
	case params.Url == "":
		// Git repository -- identifying repo url if not set by the user
		url, err := a.getGitRemoteUrl(params)

		if err != nil {
			return params, err
		}

		params.Url = url
	default:
		// making sure url is in the correct format
		_, err := url.Parse(params.Url)
		if err != nil {
			return params, fmt.Errorf("error validating url %w", err)
		}
		params.Url = utils.SanitizeRepoUrl(params.Url)

		// resetting Dir param since Url has priority over it
		params.Dir = ""
	}

	if params.Name == "" {
		repoName := utils.UrlToRepoName(params.Url)
		if nameTooLong(repoName) {
			return params, fmt.Errorf("url base name %q is too long to use as application name; please specify name with '--name'", repoName)
		}

		params.Name = generateResourceName(params.Url)
	}

	if params.Path == "" {
		params.Path = DefaultPath
	}

	if params.DeploymentType == "" {
		params.DeploymentType = DefaultDeploymentType
	}

	if params.Branch == "" {
		params.Branch = DefaultBranch

		if params.SourceType == wego.SourceTypeGit {
			branch, err := gitProvider.GetDefaultBranch(params.Url)
			if err != nil {
				return params, err
			} else {
				params.Branch = branch
			}
		}
	}

	if nameTooLong(params.Name) {
		return params, fmt.Errorf("application name too long: %s; must be <= 63 characters", params.Name)
	}

	// Validate namespace argument for helm
	if params.HelmReleaseTargetNamespace != "" {
		if nserr := utils.ValidateNamespace(params.HelmReleaseTargetNamespace); nserr != nil {
			return params, nserr
		}
	}

	return params, nil
}

func (a *App) getGitRemoteUrl(params AddParams) (string, error) {
	repo, err := a.Git.Open(params.Dir)
	if err != nil {
		return "", fmt.Errorf("failed to open repository: %s: %w", params.Dir, err)
	}

	remote, err := repo.Remote("origin")
	if err != nil {
		return "", fmt.Errorf("failed to find the origin remote in the repository: %w", err)
	}

	urls := remote.Config().URLs
	if len(urls) == 0 {
		return "", fmt.Errorf("remote config in %s does not have an url", params.Dir)
	}

	return utils.SanitizeRepoUrl(urls[0]), nil
}

func (a *App) addAppWithNoConfigRepo(info *AppResourceInfo, dryRun bool, secretRef string, appHash string) error {
	// Returns the source, app spec and kustomization
	source, appGoat, appSpec, err := a.generateAppManifests(info, secretRef, appHash)
	if err != nil {
		return fmt.Errorf("could not generate application GitOps Automation manifests: %w", err)
	}

	a.Logger.Actionf("Applying manifests to the cluster")
	return a.applyToCluster(info, dryRun, source, appGoat, appSpec)
}

func (a *App) addAppWithConfigInAppRepo(info *AppResourceInfo, params AddParams, gitProvider gitproviders.GitProvider, secretRef string, appHash string) error {
	// Returns the source, app spec and kustomization
	source, appGoat, appSpec, err := a.generateAppManifests(info, secretRef, appHash)
	if err != nil {
		return fmt.Errorf("could not generate application GitOps Automation manifests: %w", err)
	}

	appWegoGoat, err := a.generateAppWegoManifests(info)
	if err != nil {
		return fmt.Errorf("could not create GitOps automation for .wego directory: %w", err)
	}

	// a local directory has not been passed, so we clone the repo passed in the --url
	if params.Dir == "" {
		a.Logger.Actionf("Cloning %s", info.Spec.URL)
		remover, err := a.cloneRepo(info.Spec.URL, info.Spec.Branch, params.DryRun)
		if err != nil {
			return fmt.Errorf("failed to clone application repo: %w", err)
		}
		defer remover()
	}

	if !params.DryRun {
		if !params.AutoMerge {
			if err := a.createPullRequestToRepo(info, gitProvider, info.Spec.URL, appHash, appSpec, appGoat, source); err != nil {
				return err
			}
		} else {
			a.Logger.Actionf("Writing manifests to disk")

			if err := a.writeAppYaml(info, appSpec); err != nil {
				return fmt.Errorf("failed writing app.yaml to disk: %w", err)
			}

			if err := a.writeAppGoats(info, source, appGoat); err != nil {
				return fmt.Errorf("failed writing app.yaml to disk: %w", err)
			}
		}
	}

	a.Logger.Actionf("Applying manifests to the cluster")
	if err := a.applyToCluster(info, params.DryRun, source, appWegoGoat); err != nil {
		return fmt.Errorf("could not apply manifests to the cluster: %w", err)
	}

	return a.commitAndPush(func(fname string) bool {
		return strings.Contains(fname, ".wego")
	})
}

func (a *App) addAppWithConfigInExternalRepo(info *AppResourceInfo, params AddParams, gitProvider gitproviders.GitProvider, appSecretRef string, appHash string) error {
	appConfigSecretName, err := a.createAndUploadDeployKey(info, params.DryRun, info.Spec.ConfigURL, gitProvider)
	if err != nil {
		return fmt.Errorf("could not generate deploy key: %w", err)
	}

	// Returns the source, app spec and kustomization
	appSource, appGoat, appSpec, err := a.generateAppManifests(info, appSecretRef, appHash)
	if err != nil {
		return fmt.Errorf("could not generate application GitOps Automation manifests: %w", err)
	}

	configBranch, err := gitProvider.GetDefaultBranch(info.Spec.ConfigURL)
	if err != nil {
		return fmt.Errorf("could not determine default branch for config repository: %w", err)
	}

	targetSource, targetGoats, err := a.generateExternalRepoManifests(info, appConfigSecretName, configBranch)
	if err != nil {
		return fmt.Errorf("could not generate target GitOps Automation manifests: %w", err)
	}

	remover, err := a.cloneRepo(info.Spec.ConfigURL, configBranch, params.DryRun)
	if err != nil {
		return fmt.Errorf("failed to clone configuration repo: %w", err)
	}
	defer remover()

	if !params.DryRun {
		if !params.AutoMerge {
			if err := a.createPullRequestToRepo(info, gitProvider, info.Spec.ConfigURL, appHash, appSpec, appGoat, appSource); err != nil {
				return err
			}
		} else {
			a.Logger.Actionf("Writing manifests to disk")

			if err := a.writeAppYaml(info, appSpec); err != nil {
				return fmt.Errorf("failed writing app.yaml to disk: %w", err)
			}

			if err := a.writeAppGoats(info, appSource, appGoat); err != nil {
				return fmt.Errorf("failed writing application gitops manifests to disk: %w", err)
			}
		}
	}

	a.Logger.Actionf("Applying manifests to the cluster")
	if err := a.applyToCluster(info, params.DryRun, targetSource, targetGoats); err != nil {
		return fmt.Errorf("could not apply manifests to the cluster: %w", err)
	}

	return a.commitAndPush()
}

func (a *App) generateAppManifests(info *AppResourceInfo, secretRef string, appHash string) ([]byte, []byte, []byte, error) {
	var sourceManifest, appManifest, appGoatManifest []byte
	var err error
	a.Logger.Generatef("Generating Source manifest")
	sourceManifest, err = a.generateSource(info, secretRef)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("could not set up GitOps for user repository: %w", err)
	}

	a.Logger.Generatef("Generating GitOps automation manifests")
	appGoatManifest, err = a.generateApplicationGoat(info)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("could not create GitOps automation for '%s': %w", info.Name, err)
	}

	a.Logger.Generatef("Generating Application spec manifest")
	appManifest, err = generateAppYaml(info, appHash)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("could not create app.yaml for '%s': %w", info.Name, err)
	}

	return sourceManifest, appGoatManifest, appManifest, nil
}

func (a *App) generateAppWegoManifests(info *AppResourceInfo) ([]byte, error) {
	appsDirManifest, err := a.Flux.CreateKustomization(
		info.automationAppsDirKustomizationName(),
		info.Name,
		info.appYamlDir(),
		info.Namespace)
	if err != nil {
		return nil, fmt.Errorf("could not create app dir kustomization for '%s': %w", info.Name, err)
	}

	targetDirManifest, err := a.Flux.CreateKustomization(
		info.automationTargetDirKustomizationName(),
		info.Name,
		info.appAutomationDir(),
		info.Namespace)
	if err != nil {
		return nil, fmt.Errorf("could not create target dir kustomization for '%s': %w", info.Name, err)
	}

	manifests := bytes.Join([][]byte{appsDirManifest, targetDirManifest}, []byte(""))

	return bytes.ReplaceAll(manifests, []byte("path: ./wego"), []byte("path: .wego")), nil
}

func (a *App) generateExternalRepoManifests(info *AppResourceInfo, secretRef, branch string) ([]byte, []byte, error) {
	repoName := generateResourceName(info.Spec.ConfigURL)

	targetSource, err := a.Flux.CreateSourceGit(repoName, info.Spec.ConfigURL, branch, secretRef, info.Namespace)
	if err != nil {
		return nil, nil, fmt.Errorf("could not generate target source manifests: %w", err)
	}

	appGoat, err := a.Flux.CreateKustomization(
		info.automationAppsDirKustomizationName(),
		repoName,
		info.appYamlDir(),
		info.Namespace)
	if err != nil {
		return nil, nil, fmt.Errorf("could not generate app dir kustomization for '%s': %w", info.Name, err)
	}

	targetGoat, err := a.Flux.CreateKustomization(
		info.automationTargetDirKustomizationName(),
		repoName,
		info.appAutomationDir(),
		info.Namespace)
	if err != nil {
		return nil, nil, fmt.Errorf("could not generate target dir kustomization for '%s': %w", info.Name, err)
	}

	manifests := bytes.Join([][]byte{targetGoat, appGoat}, []byte(""))

	return targetSource, manifests, nil
}

func (a *App) commitAndPush(filters ...func(string) bool) error {
	a.Logger.Actionf("Committing and pushing wego updates for application")

	_, err := a.Git.Commit(git.Commit{
		Author:  git.Author{Name: "Weave Gitops", Email: "weave-gitops@weave.works"},
		Message: "Add App manifests",
	}, filters...)
	if err != nil && err != git.ErrNoStagedFiles {
		return fmt.Errorf("failed to update the repository: %w", err)
	}

	if err == nil {
		a.Logger.Actionf("Pushing app changes to repository")
		if err = a.Git.Push(context.Background()); err != nil {
			return fmt.Errorf("failed to push changes: %w", err)
		}
	} else {
		a.Logger.Successf("App is up to date")
	}

	return nil
}

func (a *App) createAndUploadDeployKey(info *AppResourceInfo, dryRun bool, repoUrl string, gitProvider gitproviders.GitProvider) (string, error) {
	if repoUrl == "" {
		return "", nil
	}

	secretRefName := info.appSecretName(repoUrl)
	if dryRun {
		return secretRefName.String(), nil
	}

	repoUrl = utils.SanitizeRepoUrl(repoUrl)

	owner, err := utils.GetOwnerFromUrl(repoUrl)
	if err != nil {
		return "", err
	}

	repoName := utils.UrlToRepoName(repoUrl)

	accountType, err := gitProvider.GetAccountType(owner)
	if err != nil {
		return "", err
	}

	repoInfo, err := gitProvider.GetRepoInfo(accountType, owner, repoName)
	if err != nil {
		return "", err
	}

	if repoInfo != nil && repoInfo.Visibility != nil && *repoInfo.Visibility == gitprovider.RepositoryVisibilityPublic {
		return "", nil
	}

	deployKeyExists, err := gitProvider.DeployKeyExists(owner, repoName)
	if err != nil {
		return "", fmt.Errorf("failed check for existing deploy key: %w", err)
	}

	secretPresent, err := a.Kube.SecretPresent(context.Background(), secretRefName.String(), info.Namespace)
	if err != nil {
		return "", fmt.Errorf("failed check for existing secret: %w", err)
	}

	if !deployKeyExists || !secretPresent {
		// TODO: @jpellizzari As of https://github.com/weaveworks/weave-gitops/pull/587/files,
		// deployKeyExists and secretPresent should always be true, meaning we will never hit this block.
		// A lot of our unit tests rely on the individual function calls being called in this block,
		// so they were left in for now and will be handled in a follow up PR.
		a.Logger.Generatef("Generating deploy key for repo %s", repoUrl)
		secret, err := a.Flux.CreateSecretGit(secretRefName.String(), repoUrl, info.Namespace)
		if err != nil {
			return "", fmt.Errorf("could not create git secret: %w", err)
		}
		var secretData corev1.Secret
		err = yaml.Unmarshal(secret, &secretData)
		if err != nil {
			return string(secret), fmt.Errorf("failed to unmarshal created secret: %w", err)
		}

		deployKey := []byte(secretData.StringData["identity.pub"])

		if err := gitProvider.UploadDeployKey(owner, repoName, deployKey); err != nil {
			return "", fmt.Errorf("error uploading deploy key: %w", err)
		}

		if out, err := a.Kube.Apply(secret, info.Namespace); err != nil {
			return "", fmt.Errorf("could not apply secret manifest: %s: %w", string(out), err)
		}
	}

	return secretRefName.String(), nil
}

func (a *App) generateSource(info *AppResourceInfo, secretRef string) ([]byte, error) {
	switch info.Spec.SourceType {
	case wego.SourceTypeGit:
		sourceManifest, err := a.Flux.CreateSourceGit(info.Name, info.Spec.URL, info.Spec.Branch, secretRef, info.Namespace)
		if err != nil {
			return nil, fmt.Errorf("could not create git source: %w", err)
		}

		return sourceManifest, nil
	case wego.SourceTypeHelm:
		return a.Flux.CreateSourceHelm(info.Name, info.Spec.URL, info.Namespace)
	default:
		return nil, fmt.Errorf("unknown source type: %v", info.Spec.SourceType)
	}
}

func (a *App) generateApplicationGoat(info *AppResourceInfo) ([]byte, error) {
	switch info.Spec.DeploymentType {
	case wego.DeploymentTypeKustomize:
		return a.Flux.CreateKustomization(info.Name, info.Name, info.Spec.Path, info.Namespace)
	case wego.DeploymentTypeHelm:
		switch info.Spec.SourceType {
		case wego.SourceTypeHelm:
			return a.Flux.CreateHelmReleaseHelmRepository(info.Name, info.Spec.Path, info.Namespace, info.Spec.HelmTargetNamespace)
		case wego.SourceTypeGit:
			return a.Flux.CreateHelmReleaseGitRepository(info.Name, info.Name, info.Spec.Path, info.Namespace, info.Spec.HelmTargetNamespace)
		default:
			return nil, fmt.Errorf("invalid source type: %v", info.Spec.SourceType)
		}
	default:
		return nil, fmt.Errorf("invalid deployment type: %v", info.Spec.DeploymentType)
	}
}

func (a *App) applyToCluster(info *AppResourceInfo, dryRun bool, manifests ...[]byte) error {
	if dryRun {
		for _, manifest := range manifests {
			fmt.Fprintf(a.Osys.Stdout(), "%s\n", manifest)
		}
		return nil
	}

	for _, manifest := range manifests {
		if out, err := a.Kube.Apply(manifest, info.Namespace); err != nil {
			return fmt.Errorf("could not apply manifest: %s: %w", string(out), err)
		}
	}

	return nil
}

func (a *App) cloneRepo(url string, branch string, dryRun bool) (func(), error) {
	if dryRun {
		return func() {}, nil
	}

	url = utils.SanitizeRepoUrl(url)

	repoDir, err := ioutil.TempDir("", "user-repo-")
	if err != nil {
		return nil, fmt.Errorf("failed creating temp. directory to clone repo: %w", err)
	}

	_, err = a.Git.Clone(context.Background(), repoDir, url, branch)
	if err != nil {
		return nil, fmt.Errorf("failed cloning user repo: %s: %w", url, err)
	}

	return func() {
		os.RemoveAll(repoDir)
	}, nil
}

func (a *App) writeAppYaml(info *AppResourceInfo, manifest []byte) error {
	return a.Git.Write(info.appYamlPath(), manifest)
}

func (a *App) writeAppGoats(info *AppResourceInfo, sourceManifest, deployManifest []byte) error {
	if err := a.Git.Write(info.appAutomationSourcePath(), sourceManifest); err != nil {
		return err
	}

	return a.Git.Write(info.appAutomationDeployPath(), deployManifest)
}

func makeWegoApplication(params AddParams) wego.Application {
	gvk := wego.GroupVersion.WithKind(wego.ApplicationKind)
	app := wego.Application{
		TypeMeta: metav1.TypeMeta{
			Kind:       gvk.Kind,
			APIVersion: gvk.GroupVersion().String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      params.Name,
			Namespace: params.Namespace,
		},
		Spec: wego.ApplicationSpec{
			ConfigURL:           params.AppConfigUrl,
			Branch:              params.Branch,
			URL:                 params.Url,
			Path:                params.Path,
			DeploymentType:      wego.DeploymentType(params.DeploymentType),
			SourceType:          wego.SourceType(params.SourceType),
			HelmTargetNamespace: params.HelmReleaseTargetNamespace,
		},
	}

	return app
}

func generateAppYaml(info *AppResourceInfo, appHash string) ([]byte, error) {
	app := info.Application

	app.ObjectMeta.Labels = map[string]string{
		WeGOAppIdentifierLabelKey: appHash,
	}

	b, err := yaml.Marshal(&app)
	if err != nil {
		return nil, fmt.Errorf("could not marshal yaml: %w", err)
	}

	return sanitizeK8sYaml(b), nil
}

func generateResourceName(url string) string {
	return hashNameIfTooLong(strings.ReplaceAll(utils.UrlToRepoName(url), "_", "-"))
}

func (a *App) createPullRequestToRepo(info *AppResourceInfo, gitProvider gitproviders.GitProvider, repo string, appHash string, appYaml []byte, goatSource, goatDeploy []byte) error {
	repoName := generateResourceName(repo)

	appPath := info.appYamlPath()
	goatSourcePath := info.appAutomationSourcePath()
	goatDeployPath := info.appAutomationDeployPath()

	appContent := string(appYaml)
	goatSourceContent := string(goatSource)
	goatDeployContent := string(goatDeploy)

	files := []gitprovider.CommitFile{
		{
			Path:    &appPath,
			Content: &appContent,
		},
		{
			Path:    &goatSourcePath,
			Content: &goatSourceContent,
		},
		{
			Path:    &goatDeployPath,
			Content: &goatDeployContent,
		},
	}

	owner, err := utils.GetOwnerFromUrl(repo)
	if err != nil {
		return fmt.Errorf("failed to retrieve owner: %w", err)
	}

	accountType, err := gitProvider.GetAccountType(owner)
	if err != nil {
		return fmt.Errorf("failed to retrieve account type: %w", err)
	}

	if accountType == gitproviders.AccountTypeOrg {
		orgRepoRef := gitproviders.NewOrgRepositoryRef(github.DefaultDomain, owner, repoName)
		prLink, err := gitProvider.CreatePullRequestToOrgRepo(orgRepoRef, info.Spec.Branch, appHash, files, utils.GetCommitMessage(), fmt.Sprintf("wego add %s", info.Name), fmt.Sprintf("Added yamls for %s", info.Name))
		if err != nil {
			return fmt.Errorf("unable to create pull request: %w", err)
		}
		a.Logger.Println("Pull Request created: %s\n", prLink.Get().WebURL)
		return nil
	}

	userRepoRef := gitproviders.NewUserRepositoryRef(github.DefaultDomain, owner, repoName)
	prLink, err := gitProvider.CreatePullRequestToUserRepo(userRepoRef, info.Spec.Branch, appHash, files, utils.GetCommitMessage(), fmt.Sprintf("wego add %s", info.Name), fmt.Sprintf("Added yamls for %s", info.Name))
	if err != nil {
		return fmt.Errorf("unable to create pull request: %w", err)
	}
	a.Logger.Println("Pull Request created: %s\n", prLink.Get().WebURL)
	return nil
}

func getAppResourceInfo(app wego.Application, clusterName string) *AppResourceInfo {
	return &AppResourceInfo{
		Application: app,
		clusterName: clusterName,
		targetName:  clusterName,
	}
}

func (a *AppResourceInfo) configMode() ConfigMode {
	if strings.ToUpper(a.Spec.ConfigURL) == string(ConfigTypeNone) {
		return ConfigModeClusterOnly
	}

	if a.Spec.ConfigURL == string(ConfigTypeUserRepo) || a.Spec.ConfigURL == a.Spec.URL {
		return ConfigModeUserRepo
	}

	return ConfigModeExternalRepo
}

func (a *AppResourceInfo) automationRoot() string {
	root := "."

	if a.configMode() == ConfigModeUserRepo {
		root = ".wego"
	}

	return root
}

func (a *AppResourceInfo) appYamlPath() string {
	return filepath.Join(a.appYamlDir(), "app.yaml")
}

func (a *AppResourceInfo) appYamlDir() string {
	return filepath.Join(a.automationRoot(), "apps", a.Name)
}

func (a *AppResourceInfo) appAutomationSourcePath() string {
	return filepath.Join(a.appAutomationDir(), fmt.Sprintf("%s-gitops-source.yaml", a.Name))
}

func (a *AppResourceInfo) appAutomationDeployPath() string {
	return filepath.Join(a.appAutomationDir(), fmt.Sprintf("%s-gitops-deploy.yaml", a.Name))
}

func (a *AppResourceInfo) appAutomationDir() string {
	return filepath.Join(a.automationRoot(), "targets", a.clusterName, a.Name)
}

func (a *AppResourceInfo) appSourceName() string {
	return a.Name
}

func (a *AppResourceInfo) appDeployName() string {
	return a.Name
}

func (a *AppResourceInfo) appResourceName() string {
	return a.Name
}

type GeneratedSecretName string

func (s GeneratedSecretName) String() string {
	return string(s)
}

func (a *AppResourceInfo) appSecretName(repoURL string) GeneratedSecretName {
	return CreateAppSecretName(a.clusterName, repoURL)
}

func nameTooLong(name string) bool {
	return len(name) > maxKubernetesResourceNameLength
}

func hashNameIfTooLong(name string) string {
	if !nameTooLong(name) {
		return name
	}

	return fmt.Sprintf("wego-%x", md5.Sum([]byte(name)))
}

func CreateAppSecretName(targetName string, repoURL string) GeneratedSecretName {
	return GeneratedSecretName(hashNameIfTooLong(fmt.Sprintf("wego-%s-%s", targetName, generateResourceName(repoURL))))
}

func (a *AppResourceInfo) automationAppsDirKustomizationName() string {
	return hashNameIfTooLong(fmt.Sprintf("%s-apps-dir", a.Name))
}

func (a *AppResourceInfo) automationTargetDirKustomizationName() string {
	return hashNameIfTooLong(fmt.Sprintf("%s-%s", a.targetName, a.Name))
}

func (a *AppResourceInfo) sourceKind() ResourceKind {
	result := ResourceKindGitRepository

	if a.Spec.SourceType == "helm" {
		result = ResourceKindHelmRepository
	}

	return result
}

func (a *AppResourceInfo) deployKind() ResourceKind {
	result := ResourceKindKustomization

	if a.Spec.DeploymentType == "helm" {
		result = ResourceKindHelmRelease
	}

	return result
}

func (a *AppResourceInfo) clusterResources() []ResourceRef {
	resources := []ResourceRef{}

	// Application GOAT, common to all three modes
	appPath := a.appYamlPath()
	automationSourcePath := a.appAutomationSourcePath()
	automationDeployPath := a.appAutomationDeployPath()

	if a.configMode() == ConfigModeClusterOnly {
		appPath = ""
		automationSourcePath = ""
		automationDeployPath = ""
	}

	resources = append(
		resources,
		ResourceRef{
			kind:           a.sourceKind(),
			name:           a.appSourceName(),
			repositoryPath: automationSourcePath},
		ResourceRef{
			kind:           a.deployKind(),
			name:           a.appDeployName(),
			repositoryPath: automationDeployPath},
		ResourceRef{
			kind:           ResourceKindApplication,
			name:           a.appResourceName(),
			repositoryPath: appPath})

	// Secret for deploy key associated with app repository;
	// common to all three modes when not using upstream Helm repository
	if a.sourceKind() == ResourceKindGitRepository {
		resources = append(
			resources,
			ResourceRef{
				kind: ResourceKindSecret,
				name: a.appSecretName(a.Spec.URL).String()})
	}

	if strings.ToUpper(a.Spec.ConfigURL) == string(ConfigTypeNone) {
		// Only app resources present in cluster; no resources to manage config
		return resources
	}

	// App dir and target dir resources are common to app and external repo modes
	resources = append(
		resources,
		// Kustomization for .wego/apps/<app-name> directory
		ResourceRef{
			kind: ResourceKindKustomization,
			name: a.automationAppsDirKustomizationName()},
		// Kustomization for .wego/targets/<cluster-name>/<app-name> directory
		ResourceRef{
			kind: ResourceKindKustomization,
			name: a.automationTargetDirKustomizationName()})

	// External repo adds a secret and source for the external repo
	if a.Spec.ConfigURL != string(ConfigTypeUserRepo) && a.Spec.ConfigURL != a.Spec.URL {
		// Config stored in external repo
		resources = append(
			resources,
			// Secret for deploy key associated with config repository
			ResourceRef{
				kind: ResourceKindSecret,
				name: a.appSecretName(a.Spec.ConfigURL).String()},
			// Source for config repository
			ResourceRef{
				kind: ResourceKindGitRepository,
				name: generateResourceName(a.Spec.ConfigURL)})
	}

	return resources
}

func (a *AppResourceInfo) clusterResourcePaths() []string {
	if a.configMode() == ConfigModeClusterOnly {
		return []string{}
	}

	return []string{a.appYamlPath(), a.appAutomationSourcePath(), a.appAutomationDeployPath()}
}

func (info *AppResourceInfo) getAppHash() string {
	var getHash = func(inputs ...string) string {
		final := []byte(strings.Join(inputs, ""))
		return fmt.Sprintf("%x", md5.Sum(final))
	}

	if info.Spec.DeploymentType == wego.DeploymentTypeHelm {
		return "wego-" + getHash(info.Spec.URL, info.Name, info.Spec.Branch)
	} else {
		return "wego-" + getHash(info.Spec.URL, info.Spec.Path, info.Spec.Branch)
	}
}

// NOTE: ready to save the targets automation in phase 2
// func (a *App) writeTargetGoats(basePath string, name string, manifests ...[]byte) error {
//  goatPath := filepath.Join(basePath, "targets", fmt.Sprintf("%s-gitops-runtime.yaml", name))

//  goat := bytes.Join(manifests, []byte(""))
//  return a.Git.Write(goatPath, goat)
// }

// Remove some problematic fields before saving the yaml files.
// K8s/reconcilers will populate these fields after creation.
// https://github.com/fluxcd/flux2/blob/0ae39d5a0a5220c177b29e71fc8824babd1e0d7c/cmd/flux/export.go#L111
func sanitizeK8sYaml(data []byte) []byte {
	out := []byte("---\n")
	data = bytes.Replace(data, []byte("  creationTimestamp: null\n"), []byte(""), 1)
	data = bytes.Replace(data, []byte("status: {}\n"), []byte(""), 1)
	return append(out, data...)
}