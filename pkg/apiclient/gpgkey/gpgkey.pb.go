// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.7.1
// source: server/gpgkey/gpgkey.proto

// GPG public key service
//
// GPG public key API performs CRUD actions against GnuPGPublicKey resources

package gpgkey

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

// Message to query the server for configured GPG public keys
type GnuPGPublicKeyQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The GPG key ID to query for
	KeyID string `protobuf:"bytes,1,opt,name=keyID,proto3" json:"keyID,omitempty"`
}

func (x *GnuPGPublicKeyQuery) Reset() {
	*x = GnuPGPublicKeyQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_gpgkey_gpgkey_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GnuPGPublicKeyQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GnuPGPublicKeyQuery) ProtoMessage() {}

func (x *GnuPGPublicKeyQuery) ProtoReflect() protoreflect.Message {
	mi := &file_server_gpgkey_gpgkey_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GnuPGPublicKeyQuery.ProtoReflect.Descriptor instead.
func (*GnuPGPublicKeyQuery) Descriptor() ([]byte, []int) {
	return file_server_gpgkey_gpgkey_proto_rawDescGZIP(), []int{0}
}

func (x *GnuPGPublicKeyQuery) GetKeyID() string {
	if x != nil {
		return x.KeyID
	}
	return ""
}

// Request to create one or more public keys on the server
type GnuPGPublicKeyCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Raw key data of the GPG key(s) to create
	Publickey *v1alpha1.GnuPGPublicKey `protobuf:"bytes,1,opt,name=publickey,proto3" json:"publickey,omitempty"`
	// Whether to upsert already existing public keys
	Upsert bool `protobuf:"varint,2,opt,name=upsert,proto3" json:"upsert,omitempty"`
}

func (x *GnuPGPublicKeyCreateRequest) Reset() {
	*x = GnuPGPublicKeyCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_gpgkey_gpgkey_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GnuPGPublicKeyCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GnuPGPublicKeyCreateRequest) ProtoMessage() {}

func (x *GnuPGPublicKeyCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_gpgkey_gpgkey_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GnuPGPublicKeyCreateRequest.ProtoReflect.Descriptor instead.
func (*GnuPGPublicKeyCreateRequest) Descriptor() ([]byte, []int) {
	return file_server_gpgkey_gpgkey_proto_rawDescGZIP(), []int{1}
}

func (x *GnuPGPublicKeyCreateRequest) GetPublickey() *v1alpha1.GnuPGPublicKey {
	if x != nil {
		return x.Publickey
	}
	return nil
}

func (x *GnuPGPublicKeyCreateRequest) GetUpsert() bool {
	if x != nil {
		return x.Upsert
	}
	return false
}

// Response to a public key creation request
type GnuPGPublicKeyCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of GPG public keys that have been created
	Created *v1alpha1.GnuPGPublicKeyList `protobuf:"bytes,1,opt,name=created,proto3" json:"created,omitempty"`
	// List of key IDs that haven been skipped because they already exist on the server
	Skipped []string `protobuf:"bytes,2,rep,name=skipped,proto3" json:"skipped,omitempty"`
}

func (x *GnuPGPublicKeyCreateResponse) Reset() {
	*x = GnuPGPublicKeyCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_gpgkey_gpgkey_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GnuPGPublicKeyCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GnuPGPublicKeyCreateResponse) ProtoMessage() {}

func (x *GnuPGPublicKeyCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_gpgkey_gpgkey_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GnuPGPublicKeyCreateResponse.ProtoReflect.Descriptor instead.
func (*GnuPGPublicKeyCreateResponse) Descriptor() ([]byte, []int) {
	return file_server_gpgkey_gpgkey_proto_rawDescGZIP(), []int{2}
}

func (x *GnuPGPublicKeyCreateResponse) GetCreated() *v1alpha1.GnuPGPublicKeyList {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *GnuPGPublicKeyCreateResponse) GetSkipped() []string {
	if x != nil {
		return x.Skipped
	}
	return nil
}

// Generic (empty) response for GPG public key CRUD requests
type GnuPGPublicKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GnuPGPublicKeyResponse) Reset() {
	*x = GnuPGPublicKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_gpgkey_gpgkey_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GnuPGPublicKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GnuPGPublicKeyResponse) ProtoMessage() {}

func (x *GnuPGPublicKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_gpgkey_gpgkey_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GnuPGPublicKeyResponse.ProtoReflect.Descriptor instead.
func (*GnuPGPublicKeyResponse) Descriptor() ([]byte, []int) {
	return file_server_gpgkey_gpgkey_proto_rawDescGZIP(), []int{3}
}

var File_server_gpgkey_gpgkey_proto protoreflect.FileDescriptor

var file_server_gpgkey_gpgkey_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x70, 0x67, 0x6b, 0x65, 0x79, 0x2f,
	0x67, 0x70, 0x67, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x67, 0x70,
	0x67, 0x6b, 0x65, 0x79, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x72, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x6a, 0x2f, 0x61, 0x72, 0x67, 0x6f, 0x2d, 0x63, 0x64, 0x2f,
	0x76, 0x32, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x2b, 0x0a, 0x13, 0x47, 0x6e, 0x75, 0x50, 0x47, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x44, 0x22, 0xa1, 0x01,
	0x0a, 0x1b, 0x47, 0x6e, 0x75, 0x50, 0x47, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x6a, 0x0a,
	0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x4c, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x72,
	0x67, 0x6f, 0x70, 0x72, 0x6f, 0x6a, 0x2e, 0x61, 0x72, 0x67, 0x6f, 0x5f, 0x63, 0x64, 0x2e, 0x76,
	0x32, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x47, 0x6e, 0x75, 0x50, 0x47, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x09,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70, 0x73,
	0x65, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x75, 0x70, 0x73, 0x65, 0x72,
	0x74, 0x22, 0xa4, 0x01, 0x0a, 0x1c, 0x47, 0x6e, 0x75, 0x50, 0x47, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x6a, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x50, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x61, 0x72, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x6a, 0x2e, 0x61, 0x72, 0x67, 0x6f, 0x5f, 0x63,
	0x64, 0x2e, 0x76, 0x32, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x47, 0x6e, 0x75, 0x50, 0x47, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x6b, 0x69, 0x70, 0x70, 0x65, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x6b, 0x69, 0x70, 0x70, 0x65, 0x64, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x6e, 0x75, 0x50,
	0x47, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0x8d, 0x04, 0x0a, 0x0d, 0x47, 0x50, 0x47, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x8e, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e,
	0x67, 0x70, 0x67, 0x6b, 0x65, 0x79, 0x2e, 0x47, 0x6e, 0x75, 0x50, 0x47, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x50, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x72, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x6a,
	0x2e, 0x61, 0x72, 0x67, 0x6f, 0x5f, 0x63, 0x64, 0x2e, 0x76, 0x32, 0x2e, 0x70, 0x6b, 0x67, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x6e, 0x75, 0x50, 0x47, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x17, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x70,
	0x67, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x91, 0x01, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1b, 0x2e,
	0x67, 0x70, 0x67, 0x6b, 0x65, 0x79, 0x2e, 0x47, 0x6e, 0x75, 0x50, 0x47, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x4c, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x72, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x6a,
	0x2e, 0x61, 0x72, 0x67, 0x6f, 0x5f, 0x63, 0x64, 0x2e, 0x76, 0x32, 0x2e, 0x70, 0x6b, 0x67, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x6e, 0x75, 0x50, 0x47, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19,
	0x12, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x70, 0x67, 0x6b, 0x65, 0x79,
	0x73, 0x2f, 0x7b, 0x6b, 0x65, 0x79, 0x49, 0x44, 0x7d, 0x12, 0x77, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x67, 0x70, 0x67, 0x6b, 0x65, 0x79, 0x2e, 0x47, 0x6e, 0x75,
	0x50, 0x47, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x67, 0x70, 0x67, 0x6b, 0x65,
	0x79, 0x2e, 0x47, 0x6e, 0x75, 0x50, 0x47, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x67, 0x70, 0x67, 0x6b, 0x65, 0x79, 0x73, 0x3a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x6b,
	0x65, 0x79, 0x12, 0x5e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x67,
	0x70, 0x67, 0x6b, 0x65, 0x79, 0x2e, 0x47, 0x6e, 0x75, 0x50, 0x47, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x1e, 0x2e, 0x67, 0x70, 0x67, 0x6b,
	0x65, 0x79, 0x2e, 0x47, 0x6e, 0x75, 0x50, 0x47, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x11, 0x2a, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x70, 0x67, 0x6b, 0x65,
	0x79, 0x73, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x72, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x6a, 0x2f, 0x61, 0x72, 0x67, 0x6f, 0x2d, 0x63,
	0x64, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2f, 0x67, 0x70, 0x67, 0x6b, 0x65, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_server_gpgkey_gpgkey_proto_rawDescOnce sync.Once
	file_server_gpgkey_gpgkey_proto_rawDescData = file_server_gpgkey_gpgkey_proto_rawDesc
)

func file_server_gpgkey_gpgkey_proto_rawDescGZIP() []byte {
	file_server_gpgkey_gpgkey_proto_rawDescOnce.Do(func() {
		file_server_gpgkey_gpgkey_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_gpgkey_gpgkey_proto_rawDescData)
	})
	return file_server_gpgkey_gpgkey_proto_rawDescData
}

var file_server_gpgkey_gpgkey_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_server_gpgkey_gpgkey_proto_goTypes = []interface{}{
	(*GnuPGPublicKeyQuery)(nil),          // 0: gpgkey.GnuPGPublicKeyQuery
	(*GnuPGPublicKeyCreateRequest)(nil),  // 1: gpgkey.GnuPGPublicKeyCreateRequest
	(*GnuPGPublicKeyCreateResponse)(nil), // 2: gpgkey.GnuPGPublicKeyCreateResponse
	(*GnuPGPublicKeyResponse)(nil),       // 3: gpgkey.GnuPGPublicKeyResponse
	(*v1alpha1.GnuPGPublicKey)(nil),      // 4: github.com.argoproj.argo_cd.v2.pkg.apis.application.v1alpha1.GnuPGPublicKey
	(*v1alpha1.GnuPGPublicKeyList)(nil),  // 5: github.com.argoproj.argo_cd.v2.pkg.apis.application.v1alpha1.GnuPGPublicKeyList
}
var file_server_gpgkey_gpgkey_proto_depIdxs = []int32{
	4, // 0: gpgkey.GnuPGPublicKeyCreateRequest.publickey:type_name -> github.com.argoproj.argo_cd.v2.pkg.apis.application.v1alpha1.GnuPGPublicKey
	5, // 1: gpgkey.GnuPGPublicKeyCreateResponse.created:type_name -> github.com.argoproj.argo_cd.v2.pkg.apis.application.v1alpha1.GnuPGPublicKeyList
	0, // 2: gpgkey.GPGKeyService.List:input_type -> gpgkey.GnuPGPublicKeyQuery
	0, // 3: gpgkey.GPGKeyService.Get:input_type -> gpgkey.GnuPGPublicKeyQuery
	1, // 4: gpgkey.GPGKeyService.Create:input_type -> gpgkey.GnuPGPublicKeyCreateRequest
	0, // 5: gpgkey.GPGKeyService.Delete:input_type -> gpgkey.GnuPGPublicKeyQuery
	5, // 6: gpgkey.GPGKeyService.List:output_type -> github.com.argoproj.argo_cd.v2.pkg.apis.application.v1alpha1.GnuPGPublicKeyList
	4, // 7: gpgkey.GPGKeyService.Get:output_type -> github.com.argoproj.argo_cd.v2.pkg.apis.application.v1alpha1.GnuPGPublicKey
	2, // 8: gpgkey.GPGKeyService.Create:output_type -> gpgkey.GnuPGPublicKeyCreateResponse
	3, // 9: gpgkey.GPGKeyService.Delete:output_type -> gpgkey.GnuPGPublicKeyResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_server_gpgkey_gpgkey_proto_init() }
func file_server_gpgkey_gpgkey_proto_init() {
	if File_server_gpgkey_gpgkey_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_gpgkey_gpgkey_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GnuPGPublicKeyQuery); i {
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
		file_server_gpgkey_gpgkey_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GnuPGPublicKeyCreateRequest); i {
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
		file_server_gpgkey_gpgkey_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GnuPGPublicKeyCreateResponse); i {
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
		file_server_gpgkey_gpgkey_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GnuPGPublicKeyResponse); i {
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
			RawDescriptor: file_server_gpgkey_gpgkey_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_gpgkey_gpgkey_proto_goTypes,
		DependencyIndexes: file_server_gpgkey_gpgkey_proto_depIdxs,
		MessageInfos:      file_server_gpgkey_gpgkey_proto_msgTypes,
	}.Build()
	File_server_gpgkey_gpgkey_proto = out.File
	file_server_gpgkey_gpgkey_proto_rawDesc = nil
	file_server_gpgkey_gpgkey_proto_goTypes = nil
	file_server_gpgkey_gpgkey_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GPGKeyServiceClient is the client API for GPGKeyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GPGKeyServiceClient interface {
	// List all available repository certificates
	List(ctx context.Context, in *GnuPGPublicKeyQuery, opts ...grpc.CallOption) (*v1alpha1.GnuPGPublicKeyList, error)
	// Get information about specified GPG public key from the server
	Get(ctx context.Context, in *GnuPGPublicKeyQuery, opts ...grpc.CallOption) (*v1alpha1.GnuPGPublicKey, error)
	// Create one or more GPG public keys in the server's configuration
	Create(ctx context.Context, in *GnuPGPublicKeyCreateRequest, opts ...grpc.CallOption) (*GnuPGPublicKeyCreateResponse, error)
	// Delete specified GPG public key from the server's configuration
	Delete(ctx context.Context, in *GnuPGPublicKeyQuery, opts ...grpc.CallOption) (*GnuPGPublicKeyResponse, error)
}

type gPGKeyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGPGKeyServiceClient(cc grpc.ClientConnInterface) GPGKeyServiceClient {
	return &gPGKeyServiceClient{cc}
}

func (c *gPGKeyServiceClient) List(ctx context.Context, in *GnuPGPublicKeyQuery, opts ...grpc.CallOption) (*v1alpha1.GnuPGPublicKeyList, error) {
	out := new(v1alpha1.GnuPGPublicKeyList)
	err := c.cc.Invoke(ctx, "/gpgkey.GPGKeyService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gPGKeyServiceClient) Get(ctx context.Context, in *GnuPGPublicKeyQuery, opts ...grpc.CallOption) (*v1alpha1.GnuPGPublicKey, error) {
	out := new(v1alpha1.GnuPGPublicKey)
	err := c.cc.Invoke(ctx, "/gpgkey.GPGKeyService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gPGKeyServiceClient) Create(ctx context.Context, in *GnuPGPublicKeyCreateRequest, opts ...grpc.CallOption) (*GnuPGPublicKeyCreateResponse, error) {
	out := new(GnuPGPublicKeyCreateResponse)
	err := c.cc.Invoke(ctx, "/gpgkey.GPGKeyService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gPGKeyServiceClient) Delete(ctx context.Context, in *GnuPGPublicKeyQuery, opts ...grpc.CallOption) (*GnuPGPublicKeyResponse, error) {
	out := new(GnuPGPublicKeyResponse)
	err := c.cc.Invoke(ctx, "/gpgkey.GPGKeyService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GPGKeyServiceServer is the server API for GPGKeyService service.
type GPGKeyServiceServer interface {
	// List all available repository certificates
	List(context.Context, *GnuPGPublicKeyQuery) (*v1alpha1.GnuPGPublicKeyList, error)
	// Get information about specified GPG public key from the server
	Get(context.Context, *GnuPGPublicKeyQuery) (*v1alpha1.GnuPGPublicKey, error)
	// Create one or more GPG public keys in the server's configuration
	Create(context.Context, *GnuPGPublicKeyCreateRequest) (*GnuPGPublicKeyCreateResponse, error)
	// Delete specified GPG public key from the server's configuration
	Delete(context.Context, *GnuPGPublicKeyQuery) (*GnuPGPublicKeyResponse, error)
}

// UnimplementedGPGKeyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGPGKeyServiceServer struct {
}

func (*UnimplementedGPGKeyServiceServer) List(context.Context, *GnuPGPublicKeyQuery) (*v1alpha1.GnuPGPublicKeyList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedGPGKeyServiceServer) Get(context.Context, *GnuPGPublicKeyQuery) (*v1alpha1.GnuPGPublicKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedGPGKeyServiceServer) Create(context.Context, *GnuPGPublicKeyCreateRequest) (*GnuPGPublicKeyCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedGPGKeyServiceServer) Delete(context.Context, *GnuPGPublicKeyQuery) (*GnuPGPublicKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterGPGKeyServiceServer(s *grpc.Server, srv GPGKeyServiceServer) {
	s.RegisterService(&_GPGKeyService_serviceDesc, srv)
}

func _GPGKeyService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GnuPGPublicKeyQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GPGKeyServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gpgkey.GPGKeyService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GPGKeyServiceServer).List(ctx, req.(*GnuPGPublicKeyQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _GPGKeyService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GnuPGPublicKeyQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GPGKeyServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gpgkey.GPGKeyService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GPGKeyServiceServer).Get(ctx, req.(*GnuPGPublicKeyQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _GPGKeyService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GnuPGPublicKeyCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GPGKeyServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gpgkey.GPGKeyService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GPGKeyServiceServer).Create(ctx, req.(*GnuPGPublicKeyCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GPGKeyService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GnuPGPublicKeyQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GPGKeyServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gpgkey.GPGKeyService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GPGKeyServiceServer).Delete(ctx, req.(*GnuPGPublicKeyQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _GPGKeyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gpgkey.GPGKeyService",
	HandlerType: (*GPGKeyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _GPGKeyService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _GPGKeyService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _GPGKeyService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _GPGKeyService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/gpgkey/gpgkey.proto",
}
