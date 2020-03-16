// Code generated by protoc-gen-go. DO NOT EDIT.
// source: edge.proto

package gon2n

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Edge struct {
	AllowP2P             bool     `protobuf:"varint,1,opt,name=AllowP2P,proto3" json:"AllowP2P,omitempty"`
	AllowRouting         bool     `protobuf:"varint,2,opt,name=AllowRouting,proto3" json:"AllowRouting,omitempty"`
	CommunityName        string   `protobuf:"bytes,3,opt,name=CommunityName,proto3" json:"CommunityName,omitempty"`
	DisablePMTUDiscovery bool     `protobuf:"varint,4,opt,name=DisablePMTUDiscovery,proto3" json:"DisablePMTUDiscovery,omitempty"`
	DisableMulticast     bool     `protobuf:"varint,5,opt,name=DisableMulticast,proto3" json:"DisableMulticast,omitempty"`
	DynamicIPMode        bool     `protobuf:"varint,6,opt,name=DynamicIPMode,proto3" json:"DynamicIPMode,omitempty"`
	EncryptionKey        string   `protobuf:"bytes,7,opt,name=EncryptionKey,proto3" json:"EncryptionKey,omitempty"`
	LocalPort            int64    `protobuf:"varint,8,opt,name=LocalPort,proto3" json:"LocalPort,omitempty"`
	ManagementPort       int64    `protobuf:"varint,9,opt,name=ManagementPort,proto3" json:"ManagementPort,omitempty"`
	RegisterInterval     int64    `protobuf:"varint,10,opt,name=RegisterInterval,proto3" json:"RegisterInterval,omitempty"`
	RegisterTTL          int64    `protobuf:"varint,11,opt,name=RegisterTTL,proto3" json:"RegisterTTL,omitempty"`
	SupernodeHostPort    string   `protobuf:"bytes,12,opt,name=SupernodeHostPort,proto3" json:"SupernodeHostPort,omitempty"`
	TypeOfService        int64    `protobuf:"varint,13,opt,name=TypeOfService,proto3" json:"TypeOfService,omitempty"`
	EncryptionMethod     int64    `protobuf:"varint,14,opt,name=EncryptionMethod,proto3" json:"EncryptionMethod,omitempty"`
	DeviceName           string   `protobuf:"bytes,15,opt,name=DeviceName,proto3" json:"DeviceName,omitempty"`
	AddressMode          string   `protobuf:"bytes,16,opt,name=AddressMode,proto3" json:"AddressMode,omitempty"`
	DeviceIP             string   `protobuf:"bytes,17,opt,name=DeviceIP,proto3" json:"DeviceIP,omitempty"`
	DeviceNetmask        string   `protobuf:"bytes,18,opt,name=DeviceNetmask,proto3" json:"DeviceNetmask,omitempty"`
	DeviceMACAddress     string   `protobuf:"bytes,19,opt,name=DeviceMACAddress,proto3" json:"DeviceMACAddress,omitempty"`
	MTU                  int64    `protobuf:"varint,20,opt,name=MTU,proto3" json:"MTU,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Edge) Reset()         { *m = Edge{} }
func (m *Edge) String() string { return proto.CompactTextString(m) }
func (*Edge) ProtoMessage()    {}
func (*Edge) Descriptor() ([]byte, []int) {
	return fileDescriptor_cab1176173a95651, []int{0}
}

func (m *Edge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Edge.Unmarshal(m, b)
}
func (m *Edge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Edge.Marshal(b, m, deterministic)
}
func (m *Edge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Edge.Merge(m, src)
}
func (m *Edge) XXX_Size() int {
	return xxx_messageInfo_Edge.Size(m)
}
func (m *Edge) XXX_DiscardUnknown() {
	xxx_messageInfo_Edge.DiscardUnknown(m)
}

var xxx_messageInfo_Edge proto.InternalMessageInfo

func (m *Edge) GetAllowP2P() bool {
	if m != nil {
		return m.AllowP2P
	}
	return false
}

func (m *Edge) GetAllowRouting() bool {
	if m != nil {
		return m.AllowRouting
	}
	return false
}

func (m *Edge) GetCommunityName() string {
	if m != nil {
		return m.CommunityName
	}
	return ""
}

func (m *Edge) GetDisablePMTUDiscovery() bool {
	if m != nil {
		return m.DisablePMTUDiscovery
	}
	return false
}

func (m *Edge) GetDisableMulticast() bool {
	if m != nil {
		return m.DisableMulticast
	}
	return false
}

func (m *Edge) GetDynamicIPMode() bool {
	if m != nil {
		return m.DynamicIPMode
	}
	return false
}

func (m *Edge) GetEncryptionKey() string {
	if m != nil {
		return m.EncryptionKey
	}
	return ""
}

func (m *Edge) GetLocalPort() int64 {
	if m != nil {
		return m.LocalPort
	}
	return 0
}

func (m *Edge) GetManagementPort() int64 {
	if m != nil {
		return m.ManagementPort
	}
	return 0
}

func (m *Edge) GetRegisterInterval() int64 {
	if m != nil {
		return m.RegisterInterval
	}
	return 0
}

func (m *Edge) GetRegisterTTL() int64 {
	if m != nil {
		return m.RegisterTTL
	}
	return 0
}

func (m *Edge) GetSupernodeHostPort() string {
	if m != nil {
		return m.SupernodeHostPort
	}
	return ""
}

func (m *Edge) GetTypeOfService() int64 {
	if m != nil {
		return m.TypeOfService
	}
	return 0
}

func (m *Edge) GetEncryptionMethod() int64 {
	if m != nil {
		return m.EncryptionMethod
	}
	return 0
}

func (m *Edge) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

func (m *Edge) GetAddressMode() string {
	if m != nil {
		return m.AddressMode
	}
	return ""
}

func (m *Edge) GetDeviceIP() string {
	if m != nil {
		return m.DeviceIP
	}
	return ""
}

func (m *Edge) GetDeviceNetmask() string {
	if m != nil {
		return m.DeviceNetmask
	}
	return ""
}

func (m *Edge) GetDeviceMACAddress() string {
	if m != nil {
		return m.DeviceMACAddress
	}
	return ""
}

func (m *Edge) GetMTU() int64 {
	if m != nil {
		return m.MTU
	}
	return 0
}

type EdgeManaged struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	AllowP2P             bool     `protobuf:"varint,2,opt,name=AllowP2P,proto3" json:"AllowP2P,omitempty"`
	AllowRouting         bool     `protobuf:"varint,3,opt,name=AllowRouting,proto3" json:"AllowRouting,omitempty"`
	CommunityName        string   `protobuf:"bytes,4,opt,name=CommunityName,proto3" json:"CommunityName,omitempty"`
	DisablePMTUDiscovery bool     `protobuf:"varint,5,opt,name=DisablePMTUDiscovery,proto3" json:"DisablePMTUDiscovery,omitempty"`
	DisableMulticast     bool     `protobuf:"varint,6,opt,name=DisableMulticast,proto3" json:"DisableMulticast,omitempty"`
	DynamicIPMode        bool     `protobuf:"varint,7,opt,name=DynamicIPMode,proto3" json:"DynamicIPMode,omitempty"`
	LocalPort            int64    `protobuf:"varint,8,opt,name=LocalPort,proto3" json:"LocalPort,omitempty"`
	ManagementPort       int64    `protobuf:"varint,9,opt,name=ManagementPort,proto3" json:"ManagementPort,omitempty"`
	RegisterInterval     int64    `protobuf:"varint,10,opt,name=RegisterInterval,proto3" json:"RegisterInterval,omitempty"`
	RegisterTTL          int64    `protobuf:"varint,11,opt,name=RegisterTTL,proto3" json:"RegisterTTL,omitempty"`
	SupernodeHostPort    string   `protobuf:"bytes,12,opt,name=SupernodeHostPort,proto3" json:"SupernodeHostPort,omitempty"`
	TypeOfService        int64    `protobuf:"varint,13,opt,name=TypeOfService,proto3" json:"TypeOfService,omitempty"`
	EncryptionMethod     int64    `protobuf:"varint,14,opt,name=EncryptionMethod,proto3" json:"EncryptionMethod,omitempty"`
	DeviceName           string   `protobuf:"bytes,15,opt,name=DeviceName,proto3" json:"DeviceName,omitempty"`
	AddressMode          string   `protobuf:"bytes,16,opt,name=AddressMode,proto3" json:"AddressMode,omitempty"`
	DeviceIP             string   `protobuf:"bytes,17,opt,name=DeviceIP,proto3" json:"DeviceIP,omitempty"`
	DeviceNetmask        string   `protobuf:"bytes,18,opt,name=DeviceNetmask,proto3" json:"DeviceNetmask,omitempty"`
	DeviceMACAddress     string   `protobuf:"bytes,19,opt,name=DeviceMACAddress,proto3" json:"DeviceMACAddress,omitempty"`
	MTU                  int64    `protobuf:"varint,20,opt,name=MTU,proto3" json:"MTU,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EdgeManaged) Reset()         { *m = EdgeManaged{} }
func (m *EdgeManaged) String() string { return proto.CompactTextString(m) }
func (*EdgeManaged) ProtoMessage()    {}
func (*EdgeManaged) Descriptor() ([]byte, []int) {
	return fileDescriptor_cab1176173a95651, []int{1}
}

func (m *EdgeManaged) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EdgeManaged.Unmarshal(m, b)
}
func (m *EdgeManaged) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EdgeManaged.Marshal(b, m, deterministic)
}
func (m *EdgeManaged) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EdgeManaged.Merge(m, src)
}
func (m *EdgeManaged) XXX_Size() int {
	return xxx_messageInfo_EdgeManaged.Size(m)
}
func (m *EdgeManaged) XXX_DiscardUnknown() {
	xxx_messageInfo_EdgeManaged.DiscardUnknown(m)
}

var xxx_messageInfo_EdgeManaged proto.InternalMessageInfo

func (m *EdgeManaged) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *EdgeManaged) GetAllowP2P() bool {
	if m != nil {
		return m.AllowP2P
	}
	return false
}

func (m *EdgeManaged) GetAllowRouting() bool {
	if m != nil {
		return m.AllowRouting
	}
	return false
}

func (m *EdgeManaged) GetCommunityName() string {
	if m != nil {
		return m.CommunityName
	}
	return ""
}

func (m *EdgeManaged) GetDisablePMTUDiscovery() bool {
	if m != nil {
		return m.DisablePMTUDiscovery
	}
	return false
}

func (m *EdgeManaged) GetDisableMulticast() bool {
	if m != nil {
		return m.DisableMulticast
	}
	return false
}

func (m *EdgeManaged) GetDynamicIPMode() bool {
	if m != nil {
		return m.DynamicIPMode
	}
	return false
}

func (m *EdgeManaged) GetLocalPort() int64 {
	if m != nil {
		return m.LocalPort
	}
	return 0
}

func (m *EdgeManaged) GetManagementPort() int64 {
	if m != nil {
		return m.ManagementPort
	}
	return 0
}

func (m *EdgeManaged) GetRegisterInterval() int64 {
	if m != nil {
		return m.RegisterInterval
	}
	return 0
}

func (m *EdgeManaged) GetRegisterTTL() int64 {
	if m != nil {
		return m.RegisterTTL
	}
	return 0
}

func (m *EdgeManaged) GetSupernodeHostPort() string {
	if m != nil {
		return m.SupernodeHostPort
	}
	return ""
}

func (m *EdgeManaged) GetTypeOfService() int64 {
	if m != nil {
		return m.TypeOfService
	}
	return 0
}

func (m *EdgeManaged) GetEncryptionMethod() int64 {
	if m != nil {
		return m.EncryptionMethod
	}
	return 0
}

func (m *EdgeManaged) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

func (m *EdgeManaged) GetAddressMode() string {
	if m != nil {
		return m.AddressMode
	}
	return ""
}

func (m *EdgeManaged) GetDeviceIP() string {
	if m != nil {
		return m.DeviceIP
	}
	return ""
}

func (m *EdgeManaged) GetDeviceNetmask() string {
	if m != nil {
		return m.DeviceNetmask
	}
	return ""
}

func (m *EdgeManaged) GetDeviceMACAddress() string {
	if m != nil {
		return m.DeviceMACAddress
	}
	return ""
}

func (m *EdgeManaged) GetMTU() int64 {
	if m != nil {
		return m.MTU
	}
	return 0
}

type EdgeManagerListArgs struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EdgeManagerListArgs) Reset()         { *m = EdgeManagerListArgs{} }
func (m *EdgeManagerListArgs) String() string { return proto.CompactTextString(m) }
func (*EdgeManagerListArgs) ProtoMessage()    {}
func (*EdgeManagerListArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_cab1176173a95651, []int{2}
}

func (m *EdgeManagerListArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EdgeManagerListArgs.Unmarshal(m, b)
}
func (m *EdgeManagerListArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EdgeManagerListArgs.Marshal(b, m, deterministic)
}
func (m *EdgeManagerListArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EdgeManagerListArgs.Merge(m, src)
}
func (m *EdgeManagerListArgs) XXX_Size() int {
	return xxx_messageInfo_EdgeManagerListArgs.Size(m)
}
func (m *EdgeManagerListArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_EdgeManagerListArgs.DiscardUnknown(m)
}

var xxx_messageInfo_EdgeManagerListArgs proto.InternalMessageInfo

type EdgeManagedId struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EdgeManagedId) Reset()         { *m = EdgeManagedId{} }
func (m *EdgeManagedId) String() string { return proto.CompactTextString(m) }
func (*EdgeManagedId) ProtoMessage()    {}
func (*EdgeManagedId) Descriptor() ([]byte, []int) {
	return fileDescriptor_cab1176173a95651, []int{3}
}

func (m *EdgeManagedId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EdgeManagedId.Unmarshal(m, b)
}
func (m *EdgeManagedId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EdgeManagedId.Marshal(b, m, deterministic)
}
func (m *EdgeManagedId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EdgeManagedId.Merge(m, src)
}
func (m *EdgeManagedId) XXX_Size() int {
	return xxx_messageInfo_EdgeManagedId.Size(m)
}
func (m *EdgeManagedId) XXX_DiscardUnknown() {
	xxx_messageInfo_EdgeManagedId.DiscardUnknown(m)
}

var xxx_messageInfo_EdgeManagedId proto.InternalMessageInfo

func (m *EdgeManagedId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type EdgeManagerListReply struct {
	EdgesManaged         []*EdgeManaged `protobuf:"bytes,1,rep,name=EdgesManaged,proto3" json:"EdgesManaged,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *EdgeManagerListReply) Reset()         { *m = EdgeManagerListReply{} }
func (m *EdgeManagerListReply) String() string { return proto.CompactTextString(m) }
func (*EdgeManagerListReply) ProtoMessage()    {}
func (*EdgeManagerListReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_cab1176173a95651, []int{4}
}

func (m *EdgeManagerListReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EdgeManagerListReply.Unmarshal(m, b)
}
func (m *EdgeManagerListReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EdgeManagerListReply.Marshal(b, m, deterministic)
}
func (m *EdgeManagerListReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EdgeManagerListReply.Merge(m, src)
}
func (m *EdgeManagerListReply) XXX_Size() int {
	return xxx_messageInfo_EdgeManagerListReply.Size(m)
}
func (m *EdgeManagerListReply) XXX_DiscardUnknown() {
	xxx_messageInfo_EdgeManagerListReply.DiscardUnknown(m)
}

var xxx_messageInfo_EdgeManagerListReply proto.InternalMessageInfo

func (m *EdgeManagerListReply) GetEdgesManaged() []*EdgeManaged {
	if m != nil {
		return m.EdgesManaged
	}
	return nil
}

func init() {
	proto.RegisterType((*Edge)(nil), "gon2n.Edge")
	proto.RegisterType((*EdgeManaged)(nil), "gon2n.EdgeManaged")
	proto.RegisterType((*EdgeManagerListArgs)(nil), "gon2n.EdgeManagerListArgs")
	proto.RegisterType((*EdgeManagedId)(nil), "gon2n.EdgeManagedId")
	proto.RegisterType((*EdgeManagerListReply)(nil), "gon2n.EdgeManagerListReply")
}

func init() {
	proto.RegisterFile("edge.proto", fileDescriptor_cab1176173a95651)
}

var fileDescriptor_cab1176173a95651 = []byte{
	// 592 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x95, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0xc7, 0xeb, 0x38, 0x49, 0x9b, 0x49, 0xdb, 0xaf, 0xdd, 0xe6, 0x93, 0x56, 0x01, 0x41, 0x64,
	0x21, 0x14, 0x21, 0xa8, 0x44, 0x90, 0x7a, 0x8f, 0x9a, 0x0a, 0x2c, 0xea, 0x62, 0xb9, 0xe9, 0x03,
	0xb8, 0xf6, 0x60, 0x2c, 0x6c, 0x6f, 0xb4, 0xde, 0x14, 0xf9, 0x69, 0xb9, 0x73, 0xe0, 0x19, 0xd0,
	0x8e, 0x4b, 0x63, 0xc7, 0xa9, 0x68, 0xaf, 0x88, 0x9b, 0xfd, 0x9b, 0xff, 0xcc, 0xce, 0x78, 0xff,
	0x23, 0x03, 0x60, 0x18, 0xe1, 0xf1, 0x42, 0x0a, 0x25, 0x58, 0x27, 0x12, 0xd9, 0x24, 0xb3, 0x7e,
	0x76, 0xa0, 0x7d, 0x16, 0x46, 0xc8, 0x86, 0xb0, 0x33, 0x4d, 0x12, 0xf1, 0xcd, 0x9d, 0xb8, 0xdc,
	0x18, 0x19, 0xe3, 0x1d, 0xef, 0xee, 0x9d, 0x59, 0xb0, 0x4b, 0xcf, 0x9e, 0x58, 0xaa, 0x38, 0x8b,
	0x78, 0x8b, 0xe2, 0x35, 0xc6, 0x5e, 0xc0, 0xde, 0xa9, 0x48, 0xd3, 0x65, 0x16, 0xab, 0xe2, 0xc2,
	0x4f, 0x91, 0x9b, 0x23, 0x63, 0xdc, 0xf3, 0xea, 0x90, 0x4d, 0x60, 0x30, 0x8b, 0x73, 0xff, 0x3a,
	0x41, 0xd7, 0x99, 0x5f, 0xcd, 0xe2, 0x3c, 0x10, 0x37, 0x28, 0x0b, 0xde, 0xa6, 0x8a, 0x1b, 0x63,
	0xec, 0x15, 0x1c, 0xdc, 0x72, 0x67, 0x99, 0xa8, 0x38, 0xf0, 0x73, 0xc5, 0x3b, 0xa4, 0x6f, 0x70,
	0xdd, 0xc5, 0xac, 0xc8, 0xfc, 0x34, 0x0e, 0x6c, 0xd7, 0x11, 0x21, 0xf2, 0x2e, 0x09, 0xeb, 0x50,
	0xab, 0xce, 0xb2, 0x40, 0x16, 0x0b, 0x15, 0x8b, 0xec, 0x23, 0x16, 0x7c, 0xbb, 0xec, 0xb5, 0x06,
	0xd9, 0x53, 0xe8, 0x9d, 0x8b, 0xc0, 0x4f, 0x5c, 0x21, 0x15, 0xdf, 0x19, 0x19, 0x63, 0xd3, 0x5b,
	0x01, 0xf6, 0x12, 0xf6, 0x1d, 0x3f, 0xf3, 0x23, 0x4c, 0x31, 0x53, 0x24, 0xe9, 0x91, 0x64, 0x8d,
	0xea, 0xee, 0x3d, 0x8c, 0xe2, 0x5c, 0xa1, 0xb4, 0x33, 0x85, 0xf2, 0xc6, 0x4f, 0x38, 0x90, 0xb2,
	0xc1, 0xd9, 0x08, 0xfa, 0xbf, 0xd9, 0x7c, 0x7e, 0xce, 0xfb, 0x24, 0xab, 0x22, 0xf6, 0x1a, 0x0e,
	0x2f, 0x97, 0x0b, 0x94, 0x99, 0x08, 0xf1, 0x83, 0xc8, 0xcb, 0x83, 0x77, 0xa9, 0xfb, 0x66, 0x40,
	0xcf, 0x39, 0x2f, 0x16, 0xf8, 0xe9, 0xf3, 0x25, 0xca, 0x9b, 0x38, 0x40, 0xbe, 0x47, 0x15, 0xeb,
	0x50, 0x77, 0xb8, 0x1a, 0xdc, 0x41, 0xf5, 0x45, 0x84, 0x7c, 0xbf, 0xec, 0x70, 0x9d, 0xb3, 0x67,
	0x00, 0x33, 0xd4, 0x59, 0x74, 0xc5, 0xff, 0xd1, 0xc1, 0x15, 0xa2, 0x27, 0x98, 0x86, 0xa1, 0xc4,
	0x3c, 0xa7, 0xaf, 0x7f, 0x40, 0x82, 0x2a, 0xd2, 0x3e, 0x2b, 0xf5, 0xb6, 0xcb, 0x0f, 0x29, 0x7c,
	0xf7, 0x4e, 0xb7, 0x57, 0xd6, 0x42, 0x95, 0xfa, 0xf9, 0x57, 0xce, 0xca, 0x7b, 0xa9, 0x41, 0xf2,
	0x03, 0x01, 0x67, 0x7a, 0x7a, 0x5b, 0x99, 0x1f, 0x91, 0xb0, 0xc1, 0xd9, 0x01, 0x98, 0xce, 0xfc,
	0x8a, 0x0f, 0x68, 0x1c, 0xfd, 0x68, 0x7d, 0xef, 0x40, 0x5f, 0x1b, 0xbe, 0xbc, 0xa6, 0x90, 0xed,
	0x43, 0xcb, 0x0e, 0xc9, 0xf1, 0x3d, 0xaf, 0x65, 0x87, 0xb5, 0x3d, 0x68, 0xfd, 0x61, 0x0f, 0xcc,
	0x87, 0xec, 0x41, 0xfb, 0x31, 0x7b, 0xd0, 0x79, 0xe4, 0x1e, 0x74, 0x1f, 0xba, 0x07, 0xdb, 0x9b,
	0xf6, 0xe0, 0x9f, 0xc3, 0xff, 0x46, 0x87, 0xff, 0x0f, 0x47, 0x2b, 0x83, 0xcb, 0xf3, 0x38, 0x57,
	0x53, 0x19, 0xe5, 0xd6, 0x73, 0xd8, 0xab, 0xf8, 0xde, 0x6e, 0x38, 0xdf, 0xba, 0x80, 0xc1, 0x5a,
	0x9e, 0x87, 0x8b, 0xa4, 0x60, 0x27, 0xb0, 0xab, 0x79, 0x7e, 0x9b, 0xc9, 0x8d, 0x91, 0x39, 0xee,
	0x4f, 0xd8, 0x31, 0xfd, 0x40, 0x8e, 0x2b, 0x35, 0xbd, 0x9a, 0x6e, 0xf2, 0xc3, 0xa8, 0x6e, 0x9a,
	0x64, 0x6f, 0xa0, 0x7b, 0x2a, 0xd1, 0x57, 0xc8, 0xfa, 0x95, 0xdc, 0xe1, 0xa0, 0x59, 0xc8, 0x0e,
	0xad, 0x2d, 0x36, 0x85, 0xb6, 0xee, 0x81, 0x0d, 0x1b, 0xf1, 0xbb, 0x99, 0x86, 0x4f, 0x36, 0xc7,
	0xa8, 0x6f, 0x6b, 0x8b, 0xbd, 0x05, 0xf3, 0x3d, 0x2a, 0xb6, 0xf1, 0x84, 0xe1, 0x86, 0x01, 0xac,
	0x2d, 0x76, 0x02, 0xdd, 0x19, 0x26, 0xa8, 0xf0, 0x9e, 0xac, 0x7b, 0xba, 0xbd, 0xee, 0xd2, 0x5f,
	0xf5, 0xdd, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x36, 0xc0, 0xe9, 0x63, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EdgeManagerClient is the client API for EdgeManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EdgeManagerClient interface {
	Create(ctx context.Context, in *Edge, opts ...grpc.CallOption) (*EdgeManagedId, error)
	List(ctx context.Context, in *EdgeManagerListArgs, opts ...grpc.CallOption) (*EdgeManagerListReply, error)
	Get(ctx context.Context, in *EdgeManagedId, opts ...grpc.CallOption) (*EdgeManaged, error)
	Delete(ctx context.Context, in *EdgeManagedId, opts ...grpc.CallOption) (*EdgeManagedId, error)
}

type edgeManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewEdgeManagerClient(cc grpc.ClientConnInterface) EdgeManagerClient {
	return &edgeManagerClient{cc}
}

func (c *edgeManagerClient) Create(ctx context.Context, in *Edge, opts ...grpc.CallOption) (*EdgeManagedId, error) {
	out := new(EdgeManagedId)
	err := c.cc.Invoke(ctx, "/gon2n.EdgeManager/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *edgeManagerClient) List(ctx context.Context, in *EdgeManagerListArgs, opts ...grpc.CallOption) (*EdgeManagerListReply, error) {
	out := new(EdgeManagerListReply)
	err := c.cc.Invoke(ctx, "/gon2n.EdgeManager/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *edgeManagerClient) Get(ctx context.Context, in *EdgeManagedId, opts ...grpc.CallOption) (*EdgeManaged, error) {
	out := new(EdgeManaged)
	err := c.cc.Invoke(ctx, "/gon2n.EdgeManager/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *edgeManagerClient) Delete(ctx context.Context, in *EdgeManagedId, opts ...grpc.CallOption) (*EdgeManagedId, error) {
	out := new(EdgeManagedId)
	err := c.cc.Invoke(ctx, "/gon2n.EdgeManager/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EdgeManagerServer is the server API for EdgeManager service.
type EdgeManagerServer interface {
	Create(context.Context, *Edge) (*EdgeManagedId, error)
	List(context.Context, *EdgeManagerListArgs) (*EdgeManagerListReply, error)
	Get(context.Context, *EdgeManagedId) (*EdgeManaged, error)
	Delete(context.Context, *EdgeManagedId) (*EdgeManagedId, error)
}

// UnimplementedEdgeManagerServer can be embedded to have forward compatible implementations.
type UnimplementedEdgeManagerServer struct {
}

func (*UnimplementedEdgeManagerServer) Create(ctx context.Context, req *Edge) (*EdgeManagedId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedEdgeManagerServer) List(ctx context.Context, req *EdgeManagerListArgs) (*EdgeManagerListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedEdgeManagerServer) Get(ctx context.Context, req *EdgeManagedId) (*EdgeManaged, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedEdgeManagerServer) Delete(ctx context.Context, req *EdgeManagedId) (*EdgeManagedId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterEdgeManagerServer(s *grpc.Server, srv EdgeManagerServer) {
	s.RegisterService(&_EdgeManager_serviceDesc, srv)
}

func _EdgeManager_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Edge)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EdgeManagerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gon2n.EdgeManager/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EdgeManagerServer).Create(ctx, req.(*Edge))
	}
	return interceptor(ctx, in, info, handler)
}

func _EdgeManager_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EdgeManagerListArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EdgeManagerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gon2n.EdgeManager/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EdgeManagerServer).List(ctx, req.(*EdgeManagerListArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _EdgeManager_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EdgeManagedId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EdgeManagerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gon2n.EdgeManager/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EdgeManagerServer).Get(ctx, req.(*EdgeManagedId))
	}
	return interceptor(ctx, in, info, handler)
}

func _EdgeManager_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EdgeManagedId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EdgeManagerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gon2n.EdgeManager/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EdgeManagerServer).Delete(ctx, req.(*EdgeManagedId))
	}
	return interceptor(ctx, in, info, handler)
}

var _EdgeManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gon2n.EdgeManager",
	HandlerType: (*EdgeManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _EdgeManager_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _EdgeManager_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _EdgeManager_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _EdgeManager_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "edge.proto",
}
