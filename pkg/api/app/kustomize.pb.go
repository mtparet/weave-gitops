// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: api/app/kustomize.proto

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

type SourceRef_Kind int32

const (
	SourceRef_GitRepository  SourceRef_Kind = 0
	SourceRef_Bucket         SourceRef_Kind = 1
	SourceRef_HelmRepository SourceRef_Kind = 2
)

// Enum value maps for SourceRef_Kind.
var (
	SourceRef_Kind_name = map[int32]string{
		0: "GitRepository",
		1: "Bucket",
		2: "HelmRepository",
	}
	SourceRef_Kind_value = map[string]int32{
		"GitRepository":  0,
		"Bucket":         1,
		"HelmRepository": 2,
	}
)

func (x SourceRef_Kind) Enum() *SourceRef_Kind {
	p := new(SourceRef_Kind)
	*p = x
	return p
}

func (x SourceRef_Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SourceRef_Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_api_app_kustomize_proto_enumTypes[0].Descriptor()
}

func (SourceRef_Kind) Type() protoreflect.EnumType {
	return &file_api_app_kustomize_proto_enumTypes[0]
}

func (x SourceRef_Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SourceRef_Kind.Descriptor instead.
func (SourceRef_Kind) EnumDescriptor() ([]byte, []int) {
	return file_api_app_kustomize_proto_rawDescGZIP(), []int{1, 0}
}

type Interval struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hours   int64 `protobuf:"varint,1,opt,name=hours,proto3" json:"hours,omitempty"`
	Minutes int64 `protobuf:"varint,2,opt,name=minutes,proto3" json:"minutes,omitempty"`
	Seconds int64 `protobuf:"varint,3,opt,name=seconds,proto3" json:"seconds,omitempty"`
}

func (x *Interval) Reset() {
	*x = Interval{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_kustomize_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Interval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Interval) ProtoMessage() {}

func (x *Interval) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_kustomize_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Interval.ProtoReflect.Descriptor instead.
func (*Interval) Descriptor() ([]byte, []int) {
	return file_api_app_kustomize_proto_rawDescGZIP(), []int{0}
}

func (x *Interval) GetHours() int64 {
	if x != nil {
		return x.Hours
	}
	return 0
}

func (x *Interval) GetMinutes() int64 {
	if x != nil {
		return x.Minutes
	}
	return 0
}

func (x *Interval) GetSeconds() int64 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

type SourceRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind SourceRef_Kind `protobuf:"varint,1,opt,name=kind,proto3,enum=gitops_server.v1.SourceRef_Kind" json:"kind,omitempty"`
	Name string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SourceRef) Reset() {
	*x = SourceRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_kustomize_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceRef) ProtoMessage() {}

func (x *SourceRef) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_kustomize_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceRef.ProtoReflect.Descriptor instead.
func (*SourceRef) Descriptor() ([]byte, []int) {
	return file_api_app_kustomize_proto_rawDescGZIP(), []int{1}
}

func (x *SourceRef) GetKind() SourceRef_Kind {
	if x != nil {
		return x.Kind
	}
	return SourceRef_GitRepository
}

func (x *SourceRef) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Kustomization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string     `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Path      string     `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	SourceRef *SourceRef `protobuf:"bytes,4,opt,name=sourceRef,proto3" json:"sourceRef,omitempty"`
	Interval  *Interval  `protobuf:"bytes,5,opt,name=interval,proto3" json:"interval,omitempty"`
}

func (x *Kustomization) Reset() {
	*x = Kustomization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_kustomize_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Kustomization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Kustomization) ProtoMessage() {}

func (x *Kustomization) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_kustomize_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Kustomization.ProtoReflect.Descriptor instead.
func (*Kustomization) Descriptor() ([]byte, []int) {
	return file_api_app_kustomize_proto_rawDescGZIP(), []int{2}
}

func (x *Kustomization) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Kustomization) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Kustomization) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Kustomization) GetSourceRef() *SourceRef {
	if x != nil {
		return x.SourceRef
	}
	return nil
}

func (x *Kustomization) GetInterval() *Interval {
	if x != nil {
		return x.Interval
	}
	return nil
}

type AddKustomizationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepoName  string     `protobuf:"bytes,1,opt,name=repo_name,json=repoName,proto3" json:"repo_name,omitempty"`
	AppName   string     `protobuf:"bytes,2,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	Name      string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string     `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Path      string     `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	SourceRef *SourceRef `protobuf:"bytes,6,opt,name=sourceRef,proto3" json:"sourceRef,omitempty"`
	Interval  *Interval  `protobuf:"bytes,7,opt,name=interval,proto3" json:"interval,omitempty"`
}

func (x *AddKustomizationRequest) Reset() {
	*x = AddKustomizationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_kustomize_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddKustomizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddKustomizationRequest) ProtoMessage() {}

func (x *AddKustomizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_kustomize_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddKustomizationRequest.ProtoReflect.Descriptor instead.
func (*AddKustomizationRequest) Descriptor() ([]byte, []int) {
	return file_api_app_kustomize_proto_rawDescGZIP(), []int{3}
}

func (x *AddKustomizationRequest) GetRepoName() string {
	if x != nil {
		return x.RepoName
	}
	return ""
}

func (x *AddKustomizationRequest) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *AddKustomizationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddKustomizationRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *AddKustomizationRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *AddKustomizationRequest) GetSourceRef() *SourceRef {
	if x != nil {
		return x.SourceRef
	}
	return nil
}

func (x *AddKustomizationRequest) GetInterval() *Interval {
	if x != nil {
		return x.Interval
	}
	return nil
}

type AddKustomizationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success       bool           `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Kustomization *Kustomization `protobuf:"bytes,2,opt,name=kustomization,proto3" json:"kustomization,omitempty"`
}

func (x *AddKustomizationResponse) Reset() {
	*x = AddKustomizationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_kustomize_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddKustomizationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddKustomizationResponse) ProtoMessage() {}

func (x *AddKustomizationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_kustomize_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddKustomizationResponse.ProtoReflect.Descriptor instead.
func (*AddKustomizationResponse) Descriptor() ([]byte, []int) {
	return file_api_app_kustomize_proto_rawDescGZIP(), []int{4}
}

func (x *AddKustomizationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *AddKustomizationResponse) GetKustomization() *Kustomization {
	if x != nil {
		return x.Kustomization
	}
	return nil
}

type RemoveKustomizationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepoName          string `protobuf:"bytes,1,opt,name=repo_name,json=repoName,proto3" json:"repo_name,omitempty"`
	AppName           string `protobuf:"bytes,2,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	KustomizationName string `protobuf:"bytes,3,opt,name=kustomization_name,json=kustomizationName,proto3" json:"kustomization_name,omitempty"`
}

func (x *RemoveKustomizationRequest) Reset() {
	*x = RemoveKustomizationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_kustomize_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveKustomizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveKustomizationRequest) ProtoMessage() {}

func (x *RemoveKustomizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_kustomize_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveKustomizationRequest.ProtoReflect.Descriptor instead.
func (*RemoveKustomizationRequest) Descriptor() ([]byte, []int) {
	return file_api_app_kustomize_proto_rawDescGZIP(), []int{5}
}

func (x *RemoveKustomizationRequest) GetRepoName() string {
	if x != nil {
		return x.RepoName
	}
	return ""
}

func (x *RemoveKustomizationRequest) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *RemoveKustomizationRequest) GetKustomizationName() string {
	if x != nil {
		return x.KustomizationName
	}
	return ""
}

type RemoveKustomizationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RemoveKustomizationResponse) Reset() {
	*x = RemoveKustomizationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_kustomize_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveKustomizationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveKustomizationResponse) ProtoMessage() {}

func (x *RemoveKustomizationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_kustomize_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveKustomizationResponse.ProtoReflect.Descriptor instead.
func (*RemoveKustomizationResponse) Descriptor() ([]byte, []int) {
	return file_api_app_kustomize_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveKustomizationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_api_app_kustomize_proto protoreflect.FileDescriptor

var file_api_app_kustomize_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x6b, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x69, 0x7a, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x67, 0x69, 0x74, 0x6f, 0x70,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x62, 0x6f, 0x64, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x08, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x12, 0x14, 0x0a, 0x05, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x90, 0x01, 0x0a, 0x09, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x12, 0x34, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x65, 0x66, 0x2e, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x39, 0x0a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x11, 0x0a, 0x0d, 0x47, 0x69,
	0x74, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x48, 0x65, 0x6c,
	0x6d, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x10, 0x02, 0x22, 0xc8, 0x01,
	0x0a, 0x0d, 0x4b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x39, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x65, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x65, 0x66, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66,
	0x12, 0x36, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x08,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x22, 0x8a, 0x02, 0x0a, 0x17, 0x41, 0x64, 0x64,
	0x4b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x39, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x65, 0x66, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x12, 0x36, 0x0a,
	0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x08, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x22, 0x7b, 0x0a, 0x18, 0x41, 0x64, 0x64, 0x4b, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x45, 0x0a, 0x0d, 0x6b,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x6b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x83, 0x01, 0x0a, 0x1a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4b, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x6b, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x37, 0x0a, 0x1b, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x4b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x32, 0xea, 0x02, 0x0a, 0x10, 0x41, 0x70, 0x70, 0x4b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x9a, 0x01, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x29,
	0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x64, 0x64, 0x4b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x67, 0x69, 0x74, 0x6f,
	0x70, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64,
	0x4b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x36, 0x22, 0x31, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x7b, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x7d, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x7b, 0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x7d, 0x2f, 0x6b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x3a, 0x01, 0x2a, 0x12, 0xb8, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x2c,
	0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x67,
	0x69, 0x74, 0x6f, 0x70, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x51, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x4b, 0x2a, 0x46, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x7b, 0x72,
	0x65, 0x70, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x7b, 0x61,
	0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x6b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x6b, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x3a, 0x01, 0x2a, 0x42, 0xca,
	0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65,
	0x61, 0x76, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2f, 0x77, 0x65, 0x61, 0x76, 0x65, 0x2d, 0x67,
	0x69, 0x74, 0x6f, 0x70, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x70, 0x70, 0x73, 0x92, 0x41, 0x94, 0x01, 0x12, 0x7c, 0x0a, 0x22, 0x57, 0x65, 0x61, 0x76, 0x65,
	0x20, 0x47, 0x69, 0x74, 0x4f, 0x70, 0x73, 0x20, 0x41, 0x70, 0x70, 0x20, 0x4b, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x41, 0x50, 0x49, 0x12, 0x51, 0x54,
	0x68, 0x65, 0x20, 0x57, 0x65, 0x61, 0x76, 0x65, 0x20, 0x47, 0x69, 0x74, 0x4f, 0x70, 0x73, 0x20,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x41, 0x50, 0x49, 0x20,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x20, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x57, 0x65, 0x61, 0x76, 0x65, 0x20, 0x47, 0x69, 0x74,
	0x4f, 0x70, 0x73, 0x20, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x32, 0x03, 0x30, 0x2e, 0x31, 0x32, 0x09, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6a, 0x73, 0x6f, 0x6e,
	0x3a, 0x09, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_app_kustomize_proto_rawDescOnce sync.Once
	file_api_app_kustomize_proto_rawDescData = file_api_app_kustomize_proto_rawDesc
)

func file_api_app_kustomize_proto_rawDescGZIP() []byte {
	file_api_app_kustomize_proto_rawDescOnce.Do(func() {
		file_api_app_kustomize_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_app_kustomize_proto_rawDescData)
	})
	return file_api_app_kustomize_proto_rawDescData
}

var file_api_app_kustomize_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_app_kustomize_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_app_kustomize_proto_goTypes = []interface{}{
	(SourceRef_Kind)(0),                 // 0: gitops_server.v1.SourceRef.Kind
	(*Interval)(nil),                    // 1: gitops_server.v1.Interval
	(*SourceRef)(nil),                   // 2: gitops_server.v1.SourceRef
	(*Kustomization)(nil),               // 3: gitops_server.v1.Kustomization
	(*AddKustomizationRequest)(nil),     // 4: gitops_server.v1.AddKustomizationRequest
	(*AddKustomizationResponse)(nil),    // 5: gitops_server.v1.AddKustomizationResponse
	(*RemoveKustomizationRequest)(nil),  // 6: gitops_server.v1.RemoveKustomizationRequest
	(*RemoveKustomizationResponse)(nil), // 7: gitops_server.v1.RemoveKustomizationResponse
}
var file_api_app_kustomize_proto_depIdxs = []int32{
	0, // 0: gitops_server.v1.SourceRef.kind:type_name -> gitops_server.v1.SourceRef.Kind
	2, // 1: gitops_server.v1.Kustomization.sourceRef:type_name -> gitops_server.v1.SourceRef
	1, // 2: gitops_server.v1.Kustomization.interval:type_name -> gitops_server.v1.Interval
	2, // 3: gitops_server.v1.AddKustomizationRequest.sourceRef:type_name -> gitops_server.v1.SourceRef
	1, // 4: gitops_server.v1.AddKustomizationRequest.interval:type_name -> gitops_server.v1.Interval
	3, // 5: gitops_server.v1.AddKustomizationResponse.kustomization:type_name -> gitops_server.v1.Kustomization
	4, // 6: gitops_server.v1.AppKustomization.Add:input_type -> gitops_server.v1.AddKustomizationRequest
	6, // 7: gitops_server.v1.AppKustomization.Remove:input_type -> gitops_server.v1.RemoveKustomizationRequest
	5, // 8: gitops_server.v1.AppKustomization.Add:output_type -> gitops_server.v1.AddKustomizationResponse
	7, // 9: gitops_server.v1.AppKustomization.Remove:output_type -> gitops_server.v1.RemoveKustomizationResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_api_app_kustomize_proto_init() }
func file_api_app_kustomize_proto_init() {
	if File_api_app_kustomize_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_app_kustomize_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Interval); i {
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
		file_api_app_kustomize_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourceRef); i {
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
		file_api_app_kustomize_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Kustomization); i {
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
		file_api_app_kustomize_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddKustomizationRequest); i {
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
		file_api_app_kustomize_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddKustomizationResponse); i {
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
		file_api_app_kustomize_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveKustomizationRequest); i {
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
		file_api_app_kustomize_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveKustomizationResponse); i {
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
			RawDescriptor: file_api_app_kustomize_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_app_kustomize_proto_goTypes,
		DependencyIndexes: file_api_app_kustomize_proto_depIdxs,
		EnumInfos:         file_api_app_kustomize_proto_enumTypes,
		MessageInfos:      file_api_app_kustomize_proto_msgTypes,
	}.Build()
	File_api_app_kustomize_proto = out.File
	file_api_app_kustomize_proto_rawDesc = nil
	file_api_app_kustomize_proto_goTypes = nil
	file_api_app_kustomize_proto_depIdxs = nil
}