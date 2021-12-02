// Code generated by counterfeiter. DO NOT EDIT.
package servicesfakes

import (
	"context"
	"sync"

	"github.com/weaveworks/weave-gitops/pkg/git"
	"github.com/weaveworks/weave-gitops/pkg/gitproviders"
	"github.com/weaveworks/weave-gitops/pkg/kube"
	"github.com/weaveworks/weave-gitops/pkg/services"
	"github.com/weaveworks/weave-gitops/pkg/services/app"
)

type FakeFactory struct {
	GetAppServiceStub        func(context.Context) (app.AppService, error)
	getAppServiceMutex       sync.RWMutex
	getAppServiceArgsForCall []struct {
		arg1 context.Context
	}
	getAppServiceReturns struct {
		result1 app.AppService
		result2 error
	}
	getAppServiceReturnsOnCall map[int]struct {
		result1 app.AppService
		result2 error
	}
	GetGitClientsStub        func(context.Context, gitproviders.Client, services.GitConfigParams) (git.Git, gitproviders.GitProvider, error)
	getGitClientsMutex       sync.RWMutex
	getGitClientsArgsForCall []struct {
		arg1 context.Context
		arg2 gitproviders.Client
		arg3 services.GitConfigParams
	}
	getGitClientsReturns struct {
		result1 git.Git
		result2 gitproviders.GitProvider
		result3 error
	}
	getGitClientsReturnsOnCall map[int]struct {
		result1 git.Git
		result2 gitproviders.GitProvider
		result3 error
	}
	GetKubeServiceStub        func() (kube.Kube, error)
	getKubeServiceMutex       sync.RWMutex
	getKubeServiceArgsForCall []struct {
	}
	getKubeServiceReturns struct {
		result1 kube.Kube
		result2 error
	}
	getKubeServiceReturnsOnCall map[int]struct {
		result1 kube.Kube
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFactory) GetAppService(arg1 context.Context) (app.AppService, error) {
	fake.getAppServiceMutex.Lock()
	ret, specificReturn := fake.getAppServiceReturnsOnCall[len(fake.getAppServiceArgsForCall)]
	fake.getAppServiceArgsForCall = append(fake.getAppServiceArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.GetAppServiceStub
	fakeReturns := fake.getAppServiceReturns
	fake.recordInvocation("GetAppService", []interface{}{arg1})
	fake.getAppServiceMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFactory) GetAppServiceCallCount() int {
	fake.getAppServiceMutex.RLock()
	defer fake.getAppServiceMutex.RUnlock()
	return len(fake.getAppServiceArgsForCall)
}

func (fake *FakeFactory) GetAppServiceCalls(stub func(context.Context) (app.AppService, error)) {
	fake.getAppServiceMutex.Lock()
	defer fake.getAppServiceMutex.Unlock()
	fake.GetAppServiceStub = stub
}

func (fake *FakeFactory) GetAppServiceArgsForCall(i int) context.Context {
	fake.getAppServiceMutex.RLock()
	defer fake.getAppServiceMutex.RUnlock()
	argsForCall := fake.getAppServiceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeFactory) GetAppServiceReturns(result1 app.AppService, result2 error) {
	fake.getAppServiceMutex.Lock()
	defer fake.getAppServiceMutex.Unlock()
	fake.GetAppServiceStub = nil
	fake.getAppServiceReturns = struct {
		result1 app.AppService
		result2 error
	}{result1, result2}
}

func (fake *FakeFactory) GetAppServiceReturnsOnCall(i int, result1 app.AppService, result2 error) {
	fake.getAppServiceMutex.Lock()
	defer fake.getAppServiceMutex.Unlock()
	fake.GetAppServiceStub = nil
	if fake.getAppServiceReturnsOnCall == nil {
		fake.getAppServiceReturnsOnCall = make(map[int]struct {
			result1 app.AppService
			result2 error
		})
	}
	fake.getAppServiceReturnsOnCall[i] = struct {
		result1 app.AppService
		result2 error
	}{result1, result2}
}

func (fake *FakeFactory) GetGitClients(arg1 context.Context, arg2 gitproviders.Client, arg3 services.GitConfigParams) (git.Git, gitproviders.GitProvider, error) {
	fake.getGitClientsMutex.Lock()
	ret, specificReturn := fake.getGitClientsReturnsOnCall[len(fake.getGitClientsArgsForCall)]
	fake.getGitClientsArgsForCall = append(fake.getGitClientsArgsForCall, struct {
		arg1 context.Context
		arg2 gitproviders.Client
		arg3 services.GitConfigParams
	}{arg1, arg2, arg3})
	stub := fake.GetGitClientsStub
	fakeReturns := fake.getGitClientsReturns
	fake.recordInvocation("GetGitClients", []interface{}{arg1, arg2, arg3})
	fake.getGitClientsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeFactory) GetGitClientsCallCount() int {
	fake.getGitClientsMutex.RLock()
	defer fake.getGitClientsMutex.RUnlock()
	return len(fake.getGitClientsArgsForCall)
}

func (fake *FakeFactory) GetGitClientsCalls(stub func(context.Context, gitproviders.Client, services.GitConfigParams) (git.Git, gitproviders.GitProvider, error)) {
	fake.getGitClientsMutex.Lock()
	defer fake.getGitClientsMutex.Unlock()
	fake.GetGitClientsStub = stub
}

func (fake *FakeFactory) GetGitClientsArgsForCall(i int) (context.Context, gitproviders.Client, services.GitConfigParams) {
	fake.getGitClientsMutex.RLock()
	defer fake.getGitClientsMutex.RUnlock()
	argsForCall := fake.getGitClientsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeFactory) GetGitClientsReturns(result1 git.Git, result2 gitproviders.GitProvider, result3 error) {
	fake.getGitClientsMutex.Lock()
	defer fake.getGitClientsMutex.Unlock()
	fake.GetGitClientsStub = nil
	fake.getGitClientsReturns = struct {
		result1 git.Git
		result2 gitproviders.GitProvider
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeFactory) GetGitClientsReturnsOnCall(i int, result1 git.Git, result2 gitproviders.GitProvider, result3 error) {
	fake.getGitClientsMutex.Lock()
	defer fake.getGitClientsMutex.Unlock()
	fake.GetGitClientsStub = nil
	if fake.getGitClientsReturnsOnCall == nil {
		fake.getGitClientsReturnsOnCall = make(map[int]struct {
			result1 git.Git
			result2 gitproviders.GitProvider
			result3 error
		})
	}
	fake.getGitClientsReturnsOnCall[i] = struct {
		result1 git.Git
		result2 gitproviders.GitProvider
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeFactory) GetKubeService() (kube.Kube, error) {
	fake.getKubeServiceMutex.Lock()
	ret, specificReturn := fake.getKubeServiceReturnsOnCall[len(fake.getKubeServiceArgsForCall)]
	fake.getKubeServiceArgsForCall = append(fake.getKubeServiceArgsForCall, struct {
	}{})
	stub := fake.GetKubeServiceStub
	fakeReturns := fake.getKubeServiceReturns
	fake.recordInvocation("GetKubeService", []interface{}{})
	fake.getKubeServiceMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFactory) GetKubeServiceCallCount() int {
	fake.getKubeServiceMutex.RLock()
	defer fake.getKubeServiceMutex.RUnlock()
	return len(fake.getKubeServiceArgsForCall)
}

func (fake *FakeFactory) GetKubeServiceCalls(stub func() (kube.Kube, error)) {
	fake.getKubeServiceMutex.Lock()
	defer fake.getKubeServiceMutex.Unlock()
	fake.GetKubeServiceStub = stub
}

func (fake *FakeFactory) GetKubeServiceReturns(result1 kube.Kube, result2 error) {
	fake.getKubeServiceMutex.Lock()
	defer fake.getKubeServiceMutex.Unlock()
	fake.GetKubeServiceStub = nil
	fake.getKubeServiceReturns = struct {
		result1 kube.Kube
		result2 error
	}{result1, result2}
}

func (fake *FakeFactory) GetKubeServiceReturnsOnCall(i int, result1 kube.Kube, result2 error) {
	fake.getKubeServiceMutex.Lock()
	defer fake.getKubeServiceMutex.Unlock()
	fake.GetKubeServiceStub = nil
	if fake.getKubeServiceReturnsOnCall == nil {
		fake.getKubeServiceReturnsOnCall = make(map[int]struct {
			result1 kube.Kube
			result2 error
		})
	}
	fake.getKubeServiceReturnsOnCall[i] = struct {
		result1 kube.Kube
		result2 error
	}{result1, result2}
}

func (fake *FakeFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getAppServiceMutex.RLock()
	defer fake.getAppServiceMutex.RUnlock()
	fake.getGitClientsMutex.RLock()
	defer fake.getGitClientsMutex.RUnlock()
	fake.getKubeServiceMutex.RLock()
	defer fake.getKubeServiceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFactory) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ services.Factory = new(FakeFactory)