// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: api/app/apps.proto

package apps

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/api/httpbody"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type App struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Id          string `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *App) Reset() {
	*x = App{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_apps_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *App) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*App) ProtoMessage() {}

func (x *App) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_apps_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use App.ProtoReflect.Descriptor instead.
func (*App) Descriptor() ([]byte, []int) {
	return file_api_app_apps_proto_rawDescGZIP(), []int{0}
}

func (x *App) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *App) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *App) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *App) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AddAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepoName    string `protobuf:"bytes,1,opt,name=repo_name,json=repoName,proto3" json:"repo_name,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	DisplayName string `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
}

func (x *AddAppRequest) Reset() {
	*x = AddAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_apps_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAppRequest) ProtoMessage() {}

func (x *AddAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_apps_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAppRequest.ProtoReflect.Descriptor instead.
func (*AddAppRequest) Descriptor() ([]byte, []int) {
	return file_api_app_apps_proto_rawDescGZIP(), []int{1}
}

func (x *AddAppRequest) GetRepoName() string {
	if x != nil {
		return x.RepoName
	}
	return ""
}

func (x *AddAppRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddAppRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AddAppRequest) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

type AddAppResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	App     *App `protobuf:"bytes,2,opt,name=app,proto3" json:"app,omitempty"`
}

func (x *AddAppResponse) Reset() {
	*x = AddAppResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_apps_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAppResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAppResponse) ProtoMessage() {}

func (x *AddAppResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_apps_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAppResponse.ProtoReflect.Descriptor instead.
func (*AddAppResponse) Descriptor() ([]byte, []int) {
	return file_api_app_apps_proto_rawDescGZIP(), []int{2}
}

func (x *AddAppResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *AddAppResponse) GetApp() *App {
	if x != nil {
		return x.App
	}
	return nil
}

type GetAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepoName string `protobuf:"bytes,1,opt,name=repo_name,json=repoName,proto3" json:"repo_name,omitempty"`
	AppName  string `protobuf:"bytes,2,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
}

func (x *GetAppRequest) Reset() {
	*x = GetAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_apps_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppRequest) ProtoMessage() {}

func (x *GetAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_apps_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppRequest.ProtoReflect.Descriptor instead.
func (*GetAppRequest) Descriptor() ([]byte, []int) {
	return file_api_app_apps_proto_rawDescGZIP(), []int{3}
}

func (x *GetAppRequest) GetRepoName() string {
	if x != nil {
		return x.RepoName
	}
	return ""
}

func (x *GetAppRequest) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

type GetAppResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	App *App `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
}

func (x *GetAppResponse) Reset() {
	*x = GetAppResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_apps_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppResponse) ProtoMessage() {}

func (x *GetAppResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_apps_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppResponse.ProtoReflect.Descriptor instead.
func (*GetAppResponse) Descriptor() ([]byte, []int) {
	return file_api_app_apps_proto_rawDescGZIP(), []int{4}
}

func (x *GetAppResponse) GetApp() *App {
	if x != nil {
		return x.App
	}
	return nil
}

type ListAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepoName string `protobuf:"bytes,1,opt,name=repo_name,json=repoName,proto3" json:"repo_name,omitempty"`
}

func (x *ListAppRequest) Reset() {
	*x = ListAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_apps_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAppRequest) ProtoMessage() {}

func (x *ListAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_apps_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAppRequest.ProtoReflect.Descriptor instead.
func (*ListAppRequest) Descriptor() ([]byte, []int) {
	return file_api_app_apps_proto_rawDescGZIP(), []int{5}
}

func (x *ListAppRequest) GetRepoName() string {
	if x != nil {
		return x.RepoName
	}
	return ""
}

type ListAppResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Apps []*App `protobuf:"bytes,1,rep,name=apps,proto3" json:"apps,omitempty"` // A list of applications
}

func (x *ListAppResponse) Reset() {
	*x = ListAppResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_apps_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAppResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAppResponse) ProtoMessage() {}

func (x *ListAppResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_apps_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAppResponse.ProtoReflect.Descriptor instead.
func (*ListAppResponse) Descriptor() ([]byte, []int) {
	return file_api_app_apps_proto_rawDescGZIP(), []int{6}
}

func (x *ListAppResponse) GetApps() []*App {
	if x != nil {
		return x.Apps
	}
	return nil
}

type RemoveAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepoName  string `protobuf:"bytes,1,opt,name=repo_name,json=repoName,proto3" json:"repo_name,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	AutoMerge bool   `protobuf:"varint,4,opt,name=autoMerge,proto3" json:"autoMerge,omitempty"`
}

func (x *RemoveAppRequest) Reset() {
	*x = RemoveAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_apps_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveAppRequest) ProtoMessage() {}

func (x *RemoveAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_apps_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveAppRequest.ProtoReflect.Descriptor instead.
func (*RemoveAppRequest) Descriptor() ([]byte, []int) {
	return file_api_app_apps_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveAppRequest) GetRepoName() string {
	if x != nil {
		return x.RepoName
	}
	return ""
}

func (x *RemoveAppRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RemoveAppRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *RemoveAppRequest) GetAutoMerge() bool {
	if x != nil {
		return x.AutoMerge
	}
	return false
}

type RemoveAppResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RemoveAppResponse) Reset() {
	*x = RemoveAppResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_apps_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveAppResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveAppResponse) ProtoMessage() {}

func (x *RemoveAppResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_apps_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveAppResponse.ProtoReflect.Descriptor instead.
func (*RemoveAppResponse) Descriptor() ([]byte, []int) {
	return file_api_app_apps_proto_rawDescGZIP(), []int{8}
}

func (x *RemoveAppResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_api_app_apps_proto protoreflect.FileDescriptor

var file_api_app_apps_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x68, 0x74, 0x74, 0x70, 0x62, 0x6f, 0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x6e, 0x0a, 0x03, 0x41, 0x70, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x85, 0x01, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x53, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x41, 0x70,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x27, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x03, 0x61, 0x70, 0x70, 0x22, 0x47, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70,
	0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70,
	0x70, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x39, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x03, 0x61, 0x70, 0x70,
	0x22, 0x2d, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x3c, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x61, 0x70, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x04, 0x61, 0x70, 0x70, 0x73, 0x22, 0x7f, 0x0a,
	0x10, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x6f, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x75, 0x74, 0x6f, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x22, 0x2d,
	0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xe8, 0x03,
	0x0a, 0x04, 0x41, 0x70, 0x70, 0x73, 0x12, 0x70, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x41, 0x70, 0x70,
	0x12, 0x1f, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x7b, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x7d, 0x2f, 0x61, 0x70, 0x70, 0x3a, 0x01, 0x2a, 0x12, 0x78, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41,
	0x70, 0x70, 0x12, 0x1f, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x12, 0x23, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x7b, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x7d, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x7b, 0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x7d, 0x12, 0x71, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x73, 0x12, 0x20,
	0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x7b, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x7d, 0x2f, 0x61, 0x70, 0x70, 0x12, 0x80, 0x01, 0x0a, 0x09, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x41, 0x70, 0x70, 0x12, 0x22, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x70, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x24, 0x2a, 0x1f, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x7b,
	0x72, 0x65, 0x70, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x7b,
	0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x3a, 0x01, 0x2a, 0x42, 0xc4, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x61, 0x76, 0x65, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x2f, 0x77, 0x65, 0x61, 0x76, 0x65, 0x2d, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x92, 0x41, 0x8e,
	0x01, 0x12, 0x76, 0x0a, 0x1c, 0x57, 0x65, 0x61, 0x76, 0x65, 0x20, 0x47, 0x69, 0x74, 0x4f, 0x70,
	0x73, 0x20, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x41, 0x50,
	0x49, 0x12, 0x51, 0x54, 0x68, 0x65, 0x20, 0x57, 0x65, 0x61, 0x76, 0x65, 0x20, 0x47, 0x69, 0x74,
	0x4f, 0x70, 0x73, 0x20, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20,
	0x41, 0x50, 0x49, 0x20, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x20, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x57, 0x65, 0x61, 0x76, 0x65,
	0x20, 0x47, 0x69, 0x74, 0x4f, 0x70, 0x73, 0x20, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x32, 0x03, 0x30, 0x2e, 0x31, 0x32, 0x09, 0x61, 0x70, 0x70, 0x73, 0x2f,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x09, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_app_apps_proto_rawDescOnce sync.Once
	file_api_app_apps_proto_rawDescData = file_api_app_apps_proto_rawDesc
)

func file_api_app_apps_proto_rawDescGZIP() []byte {
	file_api_app_apps_proto_rawDescOnce.Do(func() {
		file_api_app_apps_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_app_apps_proto_rawDescData)
	})
	return file_api_app_apps_proto_rawDescData
}

var file_api_app_apps_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_app_apps_proto_goTypes = []interface{}{
	(*App)(nil),               // 0: gitops_server.v1.App
	(*AddAppRequest)(nil),     // 1: gitops_server.v1.AddAppRequest
	(*AddAppResponse)(nil),    // 2: gitops_server.v1.AddAppResponse
	(*GetAppRequest)(nil),     // 3: gitops_server.v1.GetAppRequest
	(*GetAppResponse)(nil),    // 4: gitops_server.v1.GetAppResponse
	(*ListAppRequest)(nil),    // 5: gitops_server.v1.ListAppRequest
	(*ListAppResponse)(nil),   // 6: gitops_server.v1.ListAppResponse
	(*RemoveAppRequest)(nil),  // 7: gitops_server.v1.RemoveAppRequest
	(*RemoveAppResponse)(nil), // 8: gitops_server.v1.RemoveAppResponse
}
var file_api_app_apps_proto_depIdxs = []int32{
	0, // 0: gitops_server.v1.AddAppResponse.app:type_name -> gitops_server.v1.App
	0, // 1: gitops_server.v1.GetAppResponse.app:type_name -> gitops_server.v1.App
	0, // 2: gitops_server.v1.ListAppResponse.apps:type_name -> gitops_server.v1.App
	1, // 3: gitops_server.v1.Apps.AddApp:input_type -> gitops_server.v1.AddAppRequest
	3, // 4: gitops_server.v1.Apps.GetApp:input_type -> gitops_server.v1.GetAppRequest
	5, // 5: gitops_server.v1.Apps.ListApps:input_type -> gitops_server.v1.ListAppRequest
	7, // 6: gitops_server.v1.Apps.RemoveApp:input_type -> gitops_server.v1.RemoveAppRequest
	2, // 7: gitops_server.v1.Apps.AddApp:output_type -> gitops_server.v1.AddAppResponse
	4, // 8: gitops_server.v1.Apps.GetApp:output_type -> gitops_server.v1.GetAppResponse
	6, // 9: gitops_server.v1.Apps.ListApps:output_type -> gitops_server.v1.ListAppResponse
	8, // 10: gitops_server.v1.Apps.RemoveApp:output_type -> gitops_server.v1.RemoveAppResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_app_apps_proto_init() }
func file_api_app_apps_proto_init() {
	if File_api_app_apps_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_app_apps_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*App); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_app_apps_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAppRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_app_apps_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAppResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_app_apps_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_app_apps_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_app_apps_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAppRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_app_apps_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAppResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_app_apps_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveAppRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_app_apps_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveAppResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_app_apps_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_app_apps_proto_goTypes,
		DependencyIndexes: file_api_app_apps_proto_depIdxs,
		MessageInfos:      file_api_app_apps_proto_msgTypes,
	}.Build()
	File_api_app_apps_proto = out.File
	file_api_app_apps_proto_rawDesc = nil
	file_api_app_apps_proto_goTypes = nil
	file_api_app_apps_proto_depIdxs = nil
}