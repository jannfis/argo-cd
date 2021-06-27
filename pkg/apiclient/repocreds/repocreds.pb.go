// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.7.1
// source: server/repocreds/repocreds.proto

// Repository Service
//
// Repository Service API performs CRUD actions against repository resources

package repocreds

import (
	context "context"
	v1alpha1 "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "k8s.io/api/core/v1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// RepoCredsQuery is a query for RepoCreds resources
type RepoCredsQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Repo URL for query
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *RepoCredsQuery) Reset() {
	*x = RepoCredsQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_repocreds_repocreds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepoCredsQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepoCredsQuery) ProtoMessage() {}

func (x *RepoCredsQuery) ProtoReflect() protoreflect.Message {
	mi := &file_server_repocreds_repocreds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepoCredsQuery.ProtoReflect.Descriptor instead.
func (*RepoCredsQuery) Descriptor() ([]byte, []int) {
	return file_server_repocreds_repocreds_proto_rawDescGZIP(), []int{0}
}

func (x *RepoCredsQuery) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type RepoCredsDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *RepoCredsDeleteRequest) Reset() {
	*x = RepoCredsDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_repocreds_repocreds_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepoCredsDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepoCredsDeleteRequest) ProtoMessage() {}

func (x *RepoCredsDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_repocreds_repocreds_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepoCredsDeleteRequest.ProtoReflect.Descriptor instead.
func (*RepoCredsDeleteRequest) Descriptor() ([]byte, []int) {
	return file_server_repocreds_repocreds_proto_rawDescGZIP(), []int{1}
}

func (x *RepoCredsDeleteRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// RepoCredsResponse is a response to most repository credentials requests
type RepoCredsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RepoCredsResponse) Reset() {
	*x = RepoCredsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_repocreds_repocreds_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepoCredsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepoCredsResponse) ProtoMessage() {}

func (x *RepoCredsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_repocreds_repocreds_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepoCredsResponse.ProtoReflect.Descriptor instead.
func (*RepoCredsResponse) Descriptor() ([]byte, []int) {
	return file_server_repocreds_repocreds_proto_rawDescGZIP(), []int{2}
}

// RepoCreateRequest is a request for creating repository credentials config
type RepoCredsCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Repository definition
	Creds *v1alpha1.RepoCreds `protobuf:"bytes,1,opt,name=creds,proto3" json:"creds,omitempty"`
	// Whether to create in upsert mode
	Upsert bool `protobuf:"varint,2,opt,name=upsert,proto3" json:"upsert,omitempty"`
}

func (x *RepoCredsCreateRequest) Reset() {
	*x = RepoCredsCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_repocreds_repocreds_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepoCredsCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepoCredsCreateRequest) ProtoMessage() {}

func (x *RepoCredsCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_repocreds_repocreds_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepoCredsCreateRequest.ProtoReflect.Descriptor instead.
func (*RepoCredsCreateRequest) Descriptor() ([]byte, []int) {
	return file_server_repocreds_repocreds_proto_rawDescGZIP(), []int{3}
}

func (x *RepoCredsCreateRequest) GetCreds() *v1alpha1.RepoCreds {
	if x != nil {
		return x.Creds
	}
	return nil
}

func (x *RepoCredsCreateRequest) GetUpsert() bool {
	if x != nil {
		return x.Upsert
	}
	return false
}

// RepoCredsUpdateRequest is a request for updating existing repository credentials config
type RepoCredsUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Creds *v1alpha1.RepoCreds `protobuf:"bytes,1,opt,name=creds,proto3" json:"creds,omitempty"`
}

func (x *RepoCredsUpdateRequest) Reset() {
	*x = RepoCredsUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_repocreds_repocreds_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepoCredsUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepoCredsUpdateRequest) ProtoMessage() {}

func (x *RepoCredsUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_repocreds_repocreds_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepoCredsUpdateRequest.ProtoReflect.Descriptor instead.
func (*RepoCredsUpdateRequest) Descriptor() ([]byte, []int) {
	return file_server_repocreds_repocreds_proto_rawDescGZIP(), []int{4}
}

func (x *RepoCredsUpdateRequest) GetCreds() *v1alpha1.RepoCreds {
	if x != nil {
		return x.Creds
	}
	return nil
}

var File_server_repocreds_repocreds_proto protoreflect.FileDescriptor

var file_server_repocreds_repocreds_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x63, 0x72, 0x65,
	0x64, 0x73, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x72, 0x65, 0x70, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x73, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x6b, 0x38, 0x73,
	0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x67, 0x6f,
	0x70, 0x72, 0x6f, 0x6a, 0x2f, 0x61, 0x72, 0x67, 0x6f, 0x2d, 0x63, 0x64, 0x2f, 0x76, 0x32, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a,
	0x0e, 0x52, 0x65, 0x70, 0x6f, 0x43, 0x72, 0x65, 0x64, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x22, 0x2a, 0x0a, 0x16, 0x52, 0x65, 0x70, 0x6f, 0x43, 0x72, 0x65, 0x64, 0x73, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x13, 0x0a,
	0x11, 0x52, 0x65, 0x70, 0x6f, 0x43, 0x72, 0x65, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x16, 0x52, 0x65, 0x70, 0x6f, 0x43, 0x72, 0x65, 0x64, 0x73,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x5d, 0x0a,
	0x05, 0x63, 0x72, 0x65, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x72, 0x67, 0x6f, 0x70, 0x72,
	0x6f, 0x6a, 0x2e, 0x61, 0x72, 0x67, 0x6f, 0x5f, 0x63, 0x64, 0x2e, 0x76, 0x32, 0x2e, 0x70, 0x6b,
	0x67, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f,
	0x43, 0x72, 0x65, 0x64, 0x73, 0x52, 0x05, 0x63, 0x72, 0x65, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x70, 0x73, 0x65, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x75, 0x70,
	0x73, 0x65, 0x72, 0x74, 0x22, 0x77, 0x0a, 0x16, 0x52, 0x65, 0x70, 0x6f, 0x43, 0x72, 0x65, 0x64,
	0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x5d,
	0x0a, 0x05, 0x63, 0x72, 0x65, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x47, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x72, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x6a, 0x2e, 0x61, 0x72, 0x67, 0x6f, 0x5f, 0x63, 0x64, 0x2e, 0x76, 0x32, 0x2e, 0x70,
	0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x70,
	0x6f, 0x43, 0x72, 0x65, 0x64, 0x73, 0x52, 0x05, 0x63, 0x72, 0x65, 0x64, 0x73, 0x32, 0x9c, 0x05,
	0x0a, 0x10, 0x52, 0x65, 0x70, 0x6f, 0x43, 0x72, 0x65, 0x64, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x9e, 0x01, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x12, 0x19, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x73, 0x2e, 0x52, 0x65, 0x70,
	0x6f, 0x43, 0x72, 0x65, 0x64, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x4b, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x72, 0x67, 0x6f, 0x70, 0x72, 0x6f,
	0x6a, 0x2e, 0x61, 0x72, 0x67, 0x6f, 0x5f, 0x63, 0x64, 0x2e, 0x76, 0x32, 0x2e, 0x70, 0x6b, 0x67,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x43,
	0x72, 0x65, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13,
	0x12, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x63, 0x72,
	0x65, 0x64, 0x73, 0x12, 0xab, 0x01, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x12, 0x21, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x73, 0x2e,
	0x52, 0x65, 0x70, 0x6f, 0x43, 0x72, 0x65, 0x64, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x47, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x72, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x6a, 0x2e, 0x61, 0x72, 0x67,
	0x6f, 0x5f, 0x63, 0x64, 0x2e, 0x76, 0x32, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x43, 0x72, 0x65, 0x64, 0x73, 0x22,
	0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x70, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x73, 0x3a, 0x05, 0x63, 0x72, 0x65, 0x64,
	0x73, 0x12, 0xb7, 0x01, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x12, 0x21, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x73, 0x2e, 0x52, 0x65,
	0x70, 0x6f, 0x43, 0x72, 0x65, 0x64, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x47, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x61, 0x72, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x6a, 0x2e, 0x61, 0x72, 0x67, 0x6f, 0x5f,
	0x63, 0x64, 0x2e, 0x76, 0x32, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x43, 0x72, 0x65, 0x64, 0x73, 0x22, 0x2c, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x26, 0x1a, 0x1d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x70, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x73, 0x2f, 0x7b, 0x63, 0x72, 0x65, 0x64, 0x73, 0x2e,
	0x75, 0x72, 0x6c, 0x7d, 0x3a, 0x05, 0x63, 0x72, 0x65, 0x64, 0x73, 0x12, 0x7f, 0x0a, 0x1b, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x21, 0x2e, 0x72, 0x65, 0x70,
	0x6f, 0x63, 0x72, 0x65, 0x64, 0x73, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x43, 0x72, 0x65, 0x64, 0x73,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x72, 0x65, 0x70, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x73, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x43, 0x72,
	0x65, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x19, 0x2a, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x70,
	0x6f, 0x63, 0x72, 0x65, 0x64, 0x73, 0x2f, 0x7b, 0x75, 0x72, 0x6c, 0x7d, 0x42, 0x38, 0x5a, 0x36,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x6a, 0x2f, 0x61, 0x72, 0x67, 0x6f, 0x2d, 0x63, 0x64, 0x2f, 0x76, 0x32, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x72, 0x65, 0x70,
	0x6f, 0x63, 0x72, 0x65, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_repocreds_repocreds_proto_rawDescOnce sync.Once
	file_server_repocreds_repocreds_proto_rawDescData = file_server_repocreds_repocreds_proto_rawDesc
)

func file_server_repocreds_repocreds_proto_rawDescGZIP() []byte {
	file_server_repocreds_repocreds_proto_rawDescOnce.Do(func() {
		file_server_repocreds_repocreds_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_repocreds_repocreds_proto_rawDescData)
	})
	return file_server_repocreds_repocreds_proto_rawDescData
}

var file_server_repocreds_repocreds_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_server_repocreds_repocreds_proto_goTypes = []interface{}{
	(*RepoCredsQuery)(nil),         // 0: repocreds.RepoCredsQuery
	(*RepoCredsDeleteRequest)(nil), // 1: repocreds.RepoCredsDeleteRequest
	(*RepoCredsResponse)(nil),      // 2: repocreds.RepoCredsResponse
	(*RepoCredsCreateRequest)(nil), // 3: repocreds.RepoCredsCreateRequest
	(*RepoCredsUpdateRequest)(nil), // 4: repocreds.RepoCredsUpdateRequest
	(*v1alpha1.RepoCreds)(nil),     // 5: github.com.argoproj.argo_cd.v2.pkg.apis.application.v1alpha1.RepoCreds
	(*v1alpha1.RepoCredsList)(nil), // 6: github.com.argoproj.argo_cd.v2.pkg.apis.application.v1alpha1.RepoCredsList
}
var file_server_repocreds_repocreds_proto_depIdxs = []int32{
	5, // 0: repocreds.RepoCredsCreateRequest.creds:type_name -> github.com.argoproj.argo_cd.v2.pkg.apis.application.v1alpha1.RepoCreds
	5, // 1: repocreds.RepoCredsUpdateRequest.creds:type_name -> github.com.argoproj.argo_cd.v2.pkg.apis.application.v1alpha1.RepoCreds
	0, // 2: repocreds.RepoCredsService.ListRepositoryCredentials:input_type -> repocreds.RepoCredsQuery
	3, // 3: repocreds.RepoCredsService.CreateRepositoryCredentials:input_type -> repocreds.RepoCredsCreateRequest
	4, // 4: repocreds.RepoCredsService.UpdateRepositoryCredentials:input_type -> repocreds.RepoCredsUpdateRequest
	1, // 5: repocreds.RepoCredsService.DeleteRepositoryCredentials:input_type -> repocreds.RepoCredsDeleteRequest
	6, // 6: repocreds.RepoCredsService.ListRepositoryCredentials:output_type -> github.com.argoproj.argo_cd.v2.pkg.apis.application.v1alpha1.RepoCredsList
	5, // 7: repocreds.RepoCredsService.CreateRepositoryCredentials:output_type -> github.com.argoproj.argo_cd.v2.pkg.apis.application.v1alpha1.RepoCreds
	5, // 8: repocreds.RepoCredsService.UpdateRepositoryCredentials:output_type -> github.com.argoproj.argo_cd.v2.pkg.apis.application.v1alpha1.RepoCreds
	2, // 9: repocreds.RepoCredsService.DeleteRepositoryCredentials:output_type -> repocreds.RepoCredsResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_server_repocreds_repocreds_proto_init() }
func file_server_repocreds_repocreds_proto_init() {
	if File_server_repocreds_repocreds_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_repocreds_repocreds_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepoCredsQuery); i {
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
		file_server_repocreds_repocreds_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepoCredsDeleteRequest); i {
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
		file_server_repocreds_repocreds_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepoCredsResponse); i {
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
		file_server_repocreds_repocreds_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepoCredsCreateRequest); i {
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
		file_server_repocreds_repocreds_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepoCredsUpdateRequest); i {
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
			RawDescriptor: file_server_repocreds_repocreds_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_repocreds_repocreds_proto_goTypes,
		DependencyIndexes: file_server_repocreds_repocreds_proto_depIdxs,
		MessageInfos:      file_server_repocreds_repocreds_proto_msgTypes,
	}.Build()
	File_server_repocreds_repocreds_proto = out.File
	file_server_repocreds_repocreds_proto_rawDesc = nil
	file_server_repocreds_repocreds_proto_goTypes = nil
	file_server_repocreds_repocreds_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RepoCredsServiceClient is the client API for RepoCredsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RepoCredsServiceClient interface {
	// ListRepositoryCredentials gets a list of all configured repository credential sets
	ListRepositoryCredentials(ctx context.Context, in *RepoCredsQuery, opts ...grpc.CallOption) (*v1alpha1.RepoCredsList, error)
	// CreateRepositoryCredentials creates a new repository credential set
	CreateRepositoryCredentials(ctx context.Context, in *RepoCredsCreateRequest, opts ...grpc.CallOption) (*v1alpha1.RepoCreds, error)
	// UpdateRepositoryCredentials updates a repository credential set
	UpdateRepositoryCredentials(ctx context.Context, in *RepoCredsUpdateRequest, opts ...grpc.CallOption) (*v1alpha1.RepoCreds, error)
	// DeleteRepositoryCredentials deletes a repository credential set from the configuration
	DeleteRepositoryCredentials(ctx context.Context, in *RepoCredsDeleteRequest, opts ...grpc.CallOption) (*RepoCredsResponse, error)
}

type repoCredsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRepoCredsServiceClient(cc grpc.ClientConnInterface) RepoCredsServiceClient {
	return &repoCredsServiceClient{cc}
}

func (c *repoCredsServiceClient) ListRepositoryCredentials(ctx context.Context, in *RepoCredsQuery, opts ...grpc.CallOption) (*v1alpha1.RepoCredsList, error) {
	out := new(v1alpha1.RepoCredsList)
	err := c.cc.Invoke(ctx, "/repocreds.RepoCredsService/ListRepositoryCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoCredsServiceClient) CreateRepositoryCredentials(ctx context.Context, in *RepoCredsCreateRequest, opts ...grpc.CallOption) (*v1alpha1.RepoCreds, error) {
	out := new(v1alpha1.RepoCreds)
	err := c.cc.Invoke(ctx, "/repocreds.RepoCredsService/CreateRepositoryCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoCredsServiceClient) UpdateRepositoryCredentials(ctx context.Context, in *RepoCredsUpdateRequest, opts ...grpc.CallOption) (*v1alpha1.RepoCreds, error) {
	out := new(v1alpha1.RepoCreds)
	err := c.cc.Invoke(ctx, "/repocreds.RepoCredsService/UpdateRepositoryCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoCredsServiceClient) DeleteRepositoryCredentials(ctx context.Context, in *RepoCredsDeleteRequest, opts ...grpc.CallOption) (*RepoCredsResponse, error) {
	out := new(RepoCredsResponse)
	err := c.cc.Invoke(ctx, "/repocreds.RepoCredsService/DeleteRepositoryCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RepoCredsServiceServer is the server API for RepoCredsService service.
type RepoCredsServiceServer interface {
	// ListRepositoryCredentials gets a list of all configured repository credential sets
	ListRepositoryCredentials(context.Context, *RepoCredsQuery) (*v1alpha1.RepoCredsList, error)
	// CreateRepositoryCredentials creates a new repository credential set
	CreateRepositoryCredentials(context.Context, *RepoCredsCreateRequest) (*v1alpha1.RepoCreds, error)
	// UpdateRepositoryCredentials updates a repository credential set
	UpdateRepositoryCredentials(context.Context, *RepoCredsUpdateRequest) (*v1alpha1.RepoCreds, error)
	// DeleteRepositoryCredentials deletes a repository credential set from the configuration
	DeleteRepositoryCredentials(context.Context, *RepoCredsDeleteRequest) (*RepoCredsResponse, error)
}

// UnimplementedRepoCredsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRepoCredsServiceServer struct {
}

func (*UnimplementedRepoCredsServiceServer) ListRepositoryCredentials(context.Context, *RepoCredsQuery) (*v1alpha1.RepoCredsList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRepositoryCredentials not implemented")
}
func (*UnimplementedRepoCredsServiceServer) CreateRepositoryCredentials(context.Context, *RepoCredsCreateRequest) (*v1alpha1.RepoCreds, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRepositoryCredentials not implemented")
}
func (*UnimplementedRepoCredsServiceServer) UpdateRepositoryCredentials(context.Context, *RepoCredsUpdateRequest) (*v1alpha1.RepoCreds, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRepositoryCredentials not implemented")
}
func (*UnimplementedRepoCredsServiceServer) DeleteRepositoryCredentials(context.Context, *RepoCredsDeleteRequest) (*RepoCredsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRepositoryCredentials not implemented")
}

func RegisterRepoCredsServiceServer(s *grpc.Server, srv RepoCredsServiceServer) {
	s.RegisterService(&_RepoCredsService_serviceDesc, srv)
}

func _RepoCredsService_ListRepositoryCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepoCredsQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoCredsServiceServer).ListRepositoryCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repocreds.RepoCredsService/ListRepositoryCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoCredsServiceServer).ListRepositoryCredentials(ctx, req.(*RepoCredsQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepoCredsService_CreateRepositoryCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepoCredsCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoCredsServiceServer).CreateRepositoryCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repocreds.RepoCredsService/CreateRepositoryCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoCredsServiceServer).CreateRepositoryCredentials(ctx, req.(*RepoCredsCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepoCredsService_UpdateRepositoryCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepoCredsUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoCredsServiceServer).UpdateRepositoryCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repocreds.RepoCredsService/UpdateRepositoryCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoCredsServiceServer).UpdateRepositoryCredentials(ctx, req.(*RepoCredsUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepoCredsService_DeleteRepositoryCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepoCredsDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoCredsServiceServer).DeleteRepositoryCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repocreds.RepoCredsService/DeleteRepositoryCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoCredsServiceServer).DeleteRepositoryCredentials(ctx, req.(*RepoCredsDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RepoCredsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "repocreds.RepoCredsService",
	HandlerType: (*RepoCredsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRepositoryCredentials",
			Handler:    _RepoCredsService_ListRepositoryCredentials_Handler,
		},
		{
			MethodName: "CreateRepositoryCredentials",
			Handler:    _RepoCredsService_CreateRepositoryCredentials_Handler,
		},
		{
			MethodName: "UpdateRepositoryCredentials",
			Handler:    _RepoCredsService_UpdateRepositoryCredentials_Handler,
		},
		{
			MethodName: "DeleteRepositoryCredentials",
			Handler:    _RepoCredsService_DeleteRepositoryCredentials_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/repocreds/repocreds.proto",
}
