// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: supernode.proto

package v1

import (
	context "context"
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

type Supernode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListenPort     int64 `protobuf:"varint,1,opt,name=ListenPort,proto3" json:"ListenPort,omitempty"`
	ManagementPort int64 `protobuf:"varint,2,opt,name=ManagementPort,proto3" json:"ManagementPort,omitempty"`
}

func (x *Supernode) Reset() {
	*x = Supernode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supernode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Supernode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Supernode) ProtoMessage() {}

func (x *Supernode) ProtoReflect() protoreflect.Message {
	mi := &file_supernode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Supernode.ProtoReflect.Descriptor instead.
func (*Supernode) Descriptor() ([]byte, []int) {
	return file_supernode_proto_rawDescGZIP(), []int{0}
}

func (x *Supernode) GetListenPort() int64 {
	if x != nil {
		return x.ListenPort
	}
	return 0
}

func (x *Supernode) GetManagementPort() int64 {
	if x != nil {
		return x.ManagementPort
	}
	return 0
}

type SupernodeManaged struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	ListenPort     int64  `protobuf:"varint,2,opt,name=ListenPort,proto3" json:"ListenPort,omitempty"`
	ManagementPort int64  `protobuf:"varint,3,opt,name=ManagementPort,proto3" json:"ManagementPort,omitempty"`
}

func (x *SupernodeManaged) Reset() {
	*x = SupernodeManaged{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supernode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupernodeManaged) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupernodeManaged) ProtoMessage() {}

func (x *SupernodeManaged) ProtoReflect() protoreflect.Message {
	mi := &file_supernode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupernodeManaged.ProtoReflect.Descriptor instead.
func (*SupernodeManaged) Descriptor() ([]byte, []int) {
	return file_supernode_proto_rawDescGZIP(), []int{1}
}

func (x *SupernodeManaged) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SupernodeManaged) GetListenPort() int64 {
	if x != nil {
		return x.ListenPort
	}
	return 0
}

func (x *SupernodeManaged) GetManagementPort() int64 {
	if x != nil {
		return x.ManagementPort
	}
	return 0
}

type SupernodeManagerListArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SupernodeManagerListArgs) Reset() {
	*x = SupernodeManagerListArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supernode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupernodeManagerListArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupernodeManagerListArgs) ProtoMessage() {}

func (x *SupernodeManagerListArgs) ProtoReflect() protoreflect.Message {
	mi := &file_supernode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupernodeManagerListArgs.ProtoReflect.Descriptor instead.
func (*SupernodeManagerListArgs) Descriptor() ([]byte, []int) {
	return file_supernode_proto_rawDescGZIP(), []int{2}
}

type SupernodeManagedId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *SupernodeManagedId) Reset() {
	*x = SupernodeManagedId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supernode_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupernodeManagedId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupernodeManagedId) ProtoMessage() {}

func (x *SupernodeManagedId) ProtoReflect() protoreflect.Message {
	mi := &file_supernode_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupernodeManagedId.ProtoReflect.Descriptor instead.
func (*SupernodeManagedId) Descriptor() ([]byte, []int) {
	return file_supernode_proto_rawDescGZIP(), []int{3}
}

func (x *SupernodeManagedId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SupernodeManagerListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SupernodesManaged []*SupernodeManaged `protobuf:"bytes,1,rep,name=SupernodesManaged,proto3" json:"SupernodesManaged,omitempty"`
}

func (x *SupernodeManagerListReply) Reset() {
	*x = SupernodeManagerListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supernode_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupernodeManagerListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupernodeManagerListReply) ProtoMessage() {}

func (x *SupernodeManagerListReply) ProtoReflect() protoreflect.Message {
	mi := &file_supernode_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupernodeManagerListReply.ProtoReflect.Descriptor instead.
func (*SupernodeManagerListReply) Descriptor() ([]byte, []int) {
	return file_supernode_proto_rawDescGZIP(), []int{4}
}

func (x *SupernodeManagerListReply) GetSupernodesManaged() []*SupernodeManaged {
	if x != nil {
		return x.SupernodesManaged
	}
	return nil
}

var File_supernode_proto protoreflect.FileDescriptor

var file_supernode_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72,
	0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x67, 0x6f, 0x6e, 0x32, 0x6e, 0x22, 0x53, 0x0a, 0x09,
	0x53, 0x75, 0x70, 0x65, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x4c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x4c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x72,
	0x74, 0x22, 0x6a, 0x0a, 0x10, 0x53, 0x75, 0x70, 0x65, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x50,
	0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x1a, 0x0a,
	0x18, 0x53, 0x75, 0x70, 0x65, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x67, 0x73, 0x22, 0x24, 0x0a, 0x12, 0x53, 0x75, 0x70,
	0x65, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x49, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x22,
	0x76, 0x0a, 0x19, 0x53, 0x75, 0x70, 0x65, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x59, 0x0a, 0x11,
	0x53, 0x75, 0x70, 0x65, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f,
	0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x67, 0x6f,
	0x6e, 0x32, 0x6e, 0x2e, 0x53, 0x75, 0x70, 0x65, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x64, 0x52, 0x11, 0x53, 0x75, 0x70, 0x65, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x73,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x32, 0xb7, 0x03, 0x0a, 0x10, 0x53, 0x75, 0x70, 0x65,
	0x72, 0x6e, 0x6f, 0x64, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x5f, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6a,
	0x74, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x67, 0x6f, 0x6e,
	0x32, 0x6e, 0x2e, 0x53, 0x75, 0x70, 0x65, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x1a, 0x2d, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c,
	0x69, 0x78, 0x2e, 0x67, 0x6f, 0x6e, 0x32, 0x6e, 0x2e, 0x53, 0x75, 0x70, 0x65, 0x72, 0x6e, 0x6f,
	0x64, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x49, 0x64, 0x22, 0x00, 0x12, 0x73, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6a, 0x74,
	0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x67, 0x6f, 0x6e, 0x32,
	0x6e, 0x2e, 0x53, 0x75, 0x70, 0x65, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x34, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78,
	0x2e, 0x67, 0x6f, 0x6e, 0x32, 0x6e, 0x2e, 0x53, 0x75, 0x70, 0x65, 0x72, 0x6e, 0x6f, 0x64, 0x65,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x63, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e,
	0x67, 0x6f, 0x6e, 0x32, 0x6e, 0x2e, 0x53, 0x75, 0x70, 0x65, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x49, 0x64, 0x1a, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70,
	0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x67,
	0x6f, 0x6e, 0x32, 0x6e, 0x2e, 0x53, 0x75, 0x70, 0x65, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x64, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65,
	0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x67, 0x6f, 0x6e, 0x32, 0x6e, 0x2e, 0x53, 0x75,
	0x70, 0x65, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x49, 0x64,
	0x1a, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72,
	0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x67, 0x6f, 0x6e, 0x32, 0x6e, 0x2e, 0x53, 0x75, 0x70,
	0x65, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x49, 0x64, 0x22,
	0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x6f, 0x6a, 0x6e, 0x74, 0x66, 0x78, 0x2f, 0x67, 0x6f, 0x6e, 0x32, 0x6e, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_supernode_proto_rawDescOnce sync.Once
	file_supernode_proto_rawDescData = file_supernode_proto_rawDesc
)

func file_supernode_proto_rawDescGZIP() []byte {
	file_supernode_proto_rawDescOnce.Do(func() {
		file_supernode_proto_rawDescData = protoimpl.X.CompressGZIP(file_supernode_proto_rawDescData)
	})
	return file_supernode_proto_rawDescData
}

var file_supernode_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_supernode_proto_goTypes = []interface{}{
	(*Supernode)(nil),                 // 0: com.pojtinger.felicitas.gon2n.Supernode
	(*SupernodeManaged)(nil),          // 1: com.pojtinger.felicitas.gon2n.SupernodeManaged
	(*SupernodeManagerListArgs)(nil),  // 2: com.pojtinger.felicitas.gon2n.SupernodeManagerListArgs
	(*SupernodeManagedId)(nil),        // 3: com.pojtinger.felicitas.gon2n.SupernodeManagedId
	(*SupernodeManagerListReply)(nil), // 4: com.pojtinger.felicitas.gon2n.SupernodeManagerListReply
}
var file_supernode_proto_depIdxs = []int32{
	1, // 0: com.pojtinger.felicitas.gon2n.SupernodeManagerListReply.SupernodesManaged:type_name -> com.pojtinger.felicitas.gon2n.SupernodeManaged
	0, // 1: com.pojtinger.felicitas.gon2n.SupernodeManager.Create:input_type -> com.pojtinger.felicitas.gon2n.Supernode
	2, // 2: com.pojtinger.felicitas.gon2n.SupernodeManager.List:input_type -> com.pojtinger.felicitas.gon2n.SupernodeManagerListArgs
	3, // 3: com.pojtinger.felicitas.gon2n.SupernodeManager.Get:input_type -> com.pojtinger.felicitas.gon2n.SupernodeManagedId
	3, // 4: com.pojtinger.felicitas.gon2n.SupernodeManager.Delete:input_type -> com.pojtinger.felicitas.gon2n.SupernodeManagedId
	3, // 5: com.pojtinger.felicitas.gon2n.SupernodeManager.Create:output_type -> com.pojtinger.felicitas.gon2n.SupernodeManagedId
	4, // 6: com.pojtinger.felicitas.gon2n.SupernodeManager.List:output_type -> com.pojtinger.felicitas.gon2n.SupernodeManagerListReply
	1, // 7: com.pojtinger.felicitas.gon2n.SupernodeManager.Get:output_type -> com.pojtinger.felicitas.gon2n.SupernodeManaged
	3, // 8: com.pojtinger.felicitas.gon2n.SupernodeManager.Delete:output_type -> com.pojtinger.felicitas.gon2n.SupernodeManagedId
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_supernode_proto_init() }
func file_supernode_proto_init() {
	if File_supernode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_supernode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Supernode); i {
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
		file_supernode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupernodeManaged); i {
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
		file_supernode_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupernodeManagerListArgs); i {
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
		file_supernode_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupernodeManagedId); i {
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
		file_supernode_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupernodeManagerListReply); i {
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
			RawDescriptor: file_supernode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_supernode_proto_goTypes,
		DependencyIndexes: file_supernode_proto_depIdxs,
		MessageInfos:      file_supernode_proto_msgTypes,
	}.Build()
	File_supernode_proto = out.File
	file_supernode_proto_rawDesc = nil
	file_supernode_proto_goTypes = nil
	file_supernode_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SupernodeManagerClient is the client API for SupernodeManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SupernodeManagerClient interface {
	Create(ctx context.Context, in *Supernode, opts ...grpc.CallOption) (*SupernodeManagedId, error)
	List(ctx context.Context, in *SupernodeManagerListArgs, opts ...grpc.CallOption) (*SupernodeManagerListReply, error)
	Get(ctx context.Context, in *SupernodeManagedId, opts ...grpc.CallOption) (*SupernodeManaged, error)
	Delete(ctx context.Context, in *SupernodeManagedId, opts ...grpc.CallOption) (*SupernodeManagedId, error)
}

type supernodeManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewSupernodeManagerClient(cc grpc.ClientConnInterface) SupernodeManagerClient {
	return &supernodeManagerClient{cc}
}

func (c *supernodeManagerClient) Create(ctx context.Context, in *Supernode, opts ...grpc.CallOption) (*SupernodeManagedId, error) {
	out := new(SupernodeManagedId)
	err := c.cc.Invoke(ctx, "/com.pojtinger.felicitas.gon2n.SupernodeManager/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supernodeManagerClient) List(ctx context.Context, in *SupernodeManagerListArgs, opts ...grpc.CallOption) (*SupernodeManagerListReply, error) {
	out := new(SupernodeManagerListReply)
	err := c.cc.Invoke(ctx, "/com.pojtinger.felicitas.gon2n.SupernodeManager/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supernodeManagerClient) Get(ctx context.Context, in *SupernodeManagedId, opts ...grpc.CallOption) (*SupernodeManaged, error) {
	out := new(SupernodeManaged)
	err := c.cc.Invoke(ctx, "/com.pojtinger.felicitas.gon2n.SupernodeManager/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supernodeManagerClient) Delete(ctx context.Context, in *SupernodeManagedId, opts ...grpc.CallOption) (*SupernodeManagedId, error) {
	out := new(SupernodeManagedId)
	err := c.cc.Invoke(ctx, "/com.pojtinger.felicitas.gon2n.SupernodeManager/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SupernodeManagerServer is the server API for SupernodeManager service.
type SupernodeManagerServer interface {
	Create(context.Context, *Supernode) (*SupernodeManagedId, error)
	List(context.Context, *SupernodeManagerListArgs) (*SupernodeManagerListReply, error)
	Get(context.Context, *SupernodeManagedId) (*SupernodeManaged, error)
	Delete(context.Context, *SupernodeManagedId) (*SupernodeManagedId, error)
}

// UnimplementedSupernodeManagerServer can be embedded to have forward compatible implementations.
type UnimplementedSupernodeManagerServer struct {
}

func (*UnimplementedSupernodeManagerServer) Create(context.Context, *Supernode) (*SupernodeManagedId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedSupernodeManagerServer) List(context.Context, *SupernodeManagerListArgs) (*SupernodeManagerListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedSupernodeManagerServer) Get(context.Context, *SupernodeManagedId) (*SupernodeManaged, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedSupernodeManagerServer) Delete(context.Context, *SupernodeManagedId) (*SupernodeManagedId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterSupernodeManagerServer(s *grpc.Server, srv SupernodeManagerServer) {
	s.RegisterService(&_SupernodeManager_serviceDesc, srv)
}

func _SupernodeManager_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Supernode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupernodeManagerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.pojtinger.felicitas.gon2n.SupernodeManager/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupernodeManagerServer).Create(ctx, req.(*Supernode))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupernodeManager_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupernodeManagerListArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupernodeManagerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.pojtinger.felicitas.gon2n.SupernodeManager/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupernodeManagerServer).List(ctx, req.(*SupernodeManagerListArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupernodeManager_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupernodeManagedId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupernodeManagerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.pojtinger.felicitas.gon2n.SupernodeManager/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupernodeManagerServer).Get(ctx, req.(*SupernodeManagedId))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupernodeManager_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupernodeManagedId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupernodeManagerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.pojtinger.felicitas.gon2n.SupernodeManager/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupernodeManagerServer).Delete(ctx, req.(*SupernodeManagedId))
	}
	return interceptor(ctx, in, info, handler)
}

var _SupernodeManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.pojtinger.felicitas.gon2n.SupernodeManager",
	HandlerType: (*SupernodeManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SupernodeManager_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _SupernodeManager_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _SupernodeManager_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SupernodeManager_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "supernode.proto",
}
