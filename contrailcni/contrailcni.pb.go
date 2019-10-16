// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contrailcni.proto

package contrailcni

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

// The request message containing the user's name.
type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4886bde213a71f89, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_4886bde213a71f89, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// The request message containing the user's name.
type DelResult struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelResult) Reset()         { *m = DelResult{} }
func (m *DelResult) String() string { return proto.CompactTextString(m) }
func (*DelResult) ProtoMessage()    {}
func (*DelResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_4886bde213a71f89, []int{2}
}

func (m *DelResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelResult.Unmarshal(m, b)
}
func (m *DelResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelResult.Marshal(b, m, deterministic)
}
func (m *DelResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelResult.Merge(m, src)
}
func (m *DelResult) XXX_Size() int {
	return xxx_messageInfo_DelResult.Size(m)
}
func (m *DelResult) XXX_DiscardUnknown() {
	xxx_messageInfo_DelResult.DiscardUnknown(m)
}

var xxx_messageInfo_DelResult proto.InternalMessageInfo

func (m *DelResult) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

// The request message containing the user's name.
type CheckResult struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckResult) Reset()         { *m = CheckResult{} }
func (m *CheckResult) String() string { return proto.CompactTextString(m) }
func (*CheckResult) ProtoMessage()    {}
func (*CheckResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_4886bde213a71f89, []int{3}
}

func (m *CheckResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckResult.Unmarshal(m, b)
}
func (m *CheckResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckResult.Marshal(b, m, deterministic)
}
func (m *CheckResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckResult.Merge(m, src)
}
func (m *CheckResult) XXX_Size() int {
	return xxx_messageInfo_CheckResult.Size(m)
}
func (m *CheckResult) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckResult.DiscardUnknown(m)
}

var xxx_messageInfo_CheckResult proto.InternalMessageInfo

func (m *CheckResult) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type CNIArgs struct {
	ContainerID          string   `protobuf:"bytes,1,opt,name=ContainerID,proto3" json:"ContainerID,omitempty"`
	Netns                string   `protobuf:"bytes,2,opt,name=Netns,proto3" json:"Netns,omitempty"`
	IfName               string   `protobuf:"bytes,3,opt,name=IfName,proto3" json:"IfName,omitempty"`
	Args                 string   `protobuf:"bytes,4,opt,name=Args,proto3" json:"Args,omitempty"`
	Path                 string   `protobuf:"bytes,5,opt,name=Path,proto3" json:"Path,omitempty"`
	StdinData            string   `protobuf:"bytes,6,opt,name=StdinData,proto3" json:"StdinData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CNIArgs) Reset()         { *m = CNIArgs{} }
func (m *CNIArgs) String() string { return proto.CompactTextString(m) }
func (*CNIArgs) ProtoMessage()    {}
func (*CNIArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_4886bde213a71f89, []int{4}
}

func (m *CNIArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CNIArgs.Unmarshal(m, b)
}
func (m *CNIArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CNIArgs.Marshal(b, m, deterministic)
}
func (m *CNIArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CNIArgs.Merge(m, src)
}
func (m *CNIArgs) XXX_Size() int {
	return xxx_messageInfo_CNIArgs.Size(m)
}
func (m *CNIArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_CNIArgs.DiscardUnknown(m)
}

var xxx_messageInfo_CNIArgs proto.InternalMessageInfo

func (m *CNIArgs) GetContainerID() string {
	if m != nil {
		return m.ContainerID
	}
	return ""
}

func (m *CNIArgs) GetNetns() string {
	if m != nil {
		return m.Netns
	}
	return ""
}

func (m *CNIArgs) GetIfName() string {
	if m != nil {
		return m.IfName
	}
	return ""
}

func (m *CNIArgs) GetArgs() string {
	if m != nil {
		return m.Args
	}
	return ""
}

func (m *CNIArgs) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *CNIArgs) GetStdinData() string {
	if m != nil {
		return m.StdinData
	}
	return ""
}

type AddResult struct {
	CNIVersion           string      `protobuf:"bytes,1,opt,name=CNIVersion,proto3" json:"CNIVersion,omitempty"`
	Interfaces           []*Intf     `protobuf:"bytes,2,rep,name=Interfaces,proto3" json:"Interfaces,omitempty"`
	IPs                  []*IPConfig `protobuf:"bytes,3,rep,name=IPs,proto3" json:"IPs,omitempty"`
	Routes               []*Route    `protobuf:"bytes,4,rep,name=Routes,proto3" json:"Routes,omitempty"`
	DNSs                 *DNS        `protobuf:"bytes,5,opt,name=DNSs,proto3" json:"DNSs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AddResult) Reset()         { *m = AddResult{} }
func (m *AddResult) String() string { return proto.CompactTextString(m) }
func (*AddResult) ProtoMessage()    {}
func (*AddResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_4886bde213a71f89, []int{5}
}

func (m *AddResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddResult.Unmarshal(m, b)
}
func (m *AddResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddResult.Marshal(b, m, deterministic)
}
func (m *AddResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddResult.Merge(m, src)
}
func (m *AddResult) XXX_Size() int {
	return xxx_messageInfo_AddResult.Size(m)
}
func (m *AddResult) XXX_DiscardUnknown() {
	xxx_messageInfo_AddResult.DiscardUnknown(m)
}

var xxx_messageInfo_AddResult proto.InternalMessageInfo

func (m *AddResult) GetCNIVersion() string {
	if m != nil {
		return m.CNIVersion
	}
	return ""
}

func (m *AddResult) GetInterfaces() []*Intf {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

func (m *AddResult) GetIPs() []*IPConfig {
	if m != nil {
		return m.IPs
	}
	return nil
}

func (m *AddResult) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

func (m *AddResult) GetDNSs() *DNS {
	if m != nil {
		return m.DNSs
	}
	return nil
}

type Intf struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Mac                  string   `protobuf:"bytes,2,opt,name=Mac,proto3" json:"Mac,omitempty"`
	Sandbox              string   `protobuf:"bytes,3,opt,name=Sandbox,proto3" json:"Sandbox,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Intf) Reset()         { *m = Intf{} }
func (m *Intf) String() string { return proto.CompactTextString(m) }
func (*Intf) ProtoMessage()    {}
func (*Intf) Descriptor() ([]byte, []int) {
	return fileDescriptor_4886bde213a71f89, []int{6}
}

func (m *Intf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Intf.Unmarshal(m, b)
}
func (m *Intf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Intf.Marshal(b, m, deterministic)
}
func (m *Intf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Intf.Merge(m, src)
}
func (m *Intf) XXX_Size() int {
	return xxx_messageInfo_Intf.Size(m)
}
func (m *Intf) XXX_DiscardUnknown() {
	xxx_messageInfo_Intf.DiscardUnknown(m)
}

var xxx_messageInfo_Intf proto.InternalMessageInfo

func (m *Intf) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Intf) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

func (m *Intf) GetSandbox() string {
	if m != nil {
		return m.Sandbox
	}
	return ""
}

type IPConfig struct {
	Version              string   `protobuf:"bytes,1,opt,name=Version,proto3" json:"Version,omitempty"`
	Intf                 int32    `protobuf:"varint,2,opt,name=Intf,proto3" json:"Intf,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=Address,proto3" json:"Address,omitempty"`
	Gateway              string   `protobuf:"bytes,4,opt,name=Gateway,proto3" json:"Gateway,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IPConfig) Reset()         { *m = IPConfig{} }
func (m *IPConfig) String() string { return proto.CompactTextString(m) }
func (*IPConfig) ProtoMessage()    {}
func (*IPConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_4886bde213a71f89, []int{7}
}

func (m *IPConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPConfig.Unmarshal(m, b)
}
func (m *IPConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPConfig.Marshal(b, m, deterministic)
}
func (m *IPConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPConfig.Merge(m, src)
}
func (m *IPConfig) XXX_Size() int {
	return xxx_messageInfo_IPConfig.Size(m)
}
func (m *IPConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_IPConfig.DiscardUnknown(m)
}

var xxx_messageInfo_IPConfig proto.InternalMessageInfo

func (m *IPConfig) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *IPConfig) GetIntf() int32 {
	if m != nil {
		return m.Intf
	}
	return 0
}

func (m *IPConfig) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *IPConfig) GetGateway() string {
	if m != nil {
		return m.Gateway
	}
	return ""
}

type Route struct {
	Dst                  string   `protobuf:"bytes,1,opt,name=Dst,proto3" json:"Dst,omitempty"`
	GW                   string   `protobuf:"bytes,2,opt,name=GW,proto3" json:"GW,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_4886bde213a71f89, []int{8}
}

func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetDst() string {
	if m != nil {
		return m.Dst
	}
	return ""
}

func (m *Route) GetGW() string {
	if m != nil {
		return m.GW
	}
	return ""
}

type DNS struct {
	Nameservers          []string `protobuf:"bytes,1,rep,name=Nameservers,proto3" json:"Nameservers,omitempty"`
	Domain               string   `protobuf:"bytes,2,opt,name=Domain,proto3" json:"Domain,omitempty"`
	Search               []string `protobuf:"bytes,3,rep,name=Search,proto3" json:"Search,omitempty"`
	Options              []string `protobuf:"bytes,4,rep,name=Options,proto3" json:"Options,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DNS) Reset()         { *m = DNS{} }
func (m *DNS) String() string { return proto.CompactTextString(m) }
func (*DNS) ProtoMessage()    {}
func (*DNS) Descriptor() ([]byte, []int) {
	return fileDescriptor_4886bde213a71f89, []int{9}
}

func (m *DNS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DNS.Unmarshal(m, b)
}
func (m *DNS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DNS.Marshal(b, m, deterministic)
}
func (m *DNS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DNS.Merge(m, src)
}
func (m *DNS) XXX_Size() int {
	return xxx_messageInfo_DNS.Size(m)
}
func (m *DNS) XXX_DiscardUnknown() {
	xxx_messageInfo_DNS.DiscardUnknown(m)
}

var xxx_messageInfo_DNS proto.InternalMessageInfo

func (m *DNS) GetNameservers() []string {
	if m != nil {
		return m.Nameservers
	}
	return nil
}

func (m *DNS) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *DNS) GetSearch() []string {
	if m != nil {
		return m.Search
	}
	return nil
}

func (m *DNS) GetOptions() []string {
	if m != nil {
		return m.Options
	}
	return nil
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "contrailcni.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "contrailcni.HelloReply")
	proto.RegisterType((*DelResult)(nil), "contrailcni.DelResult")
	proto.RegisterType((*CheckResult)(nil), "contrailcni.CheckResult")
	proto.RegisterType((*CNIArgs)(nil), "contrailcni.CNIArgs")
	proto.RegisterType((*AddResult)(nil), "contrailcni.AddResult")
	proto.RegisterType((*Intf)(nil), "contrailcni.Intf")
	proto.RegisterType((*IPConfig)(nil), "contrailcni.IPConfig")
	proto.RegisterType((*Route)(nil), "contrailcni.Route")
	proto.RegisterType((*DNS)(nil), "contrailcni.DNS")
}

func init() { proto.RegisterFile("contrailcni.proto", fileDescriptor_4886bde213a71f89) }

var fileDescriptor_4886bde213a71f89 = []byte{
	// 571 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x6e, 0x96, 0xb6, 0x6b, 0x4e, 0x11, 0xda, 0xac, 0x31, 0xc2, 0x84, 0x50, 0x65, 0xfe, 0x06,
	0x17, 0x93, 0x28, 0xe2, 0x82, 0x3b, 0xaa, 0x44, 0x94, 0x5c, 0x10, 0x2a, 0x47, 0x62, 0xd7, 0x5e,
	0xe2, 0xb6, 0x11, 0xa9, 0xd3, 0xc5, 0x2e, 0xd0, 0xa7, 0xe1, 0xb9, 0x78, 0x13, 0x2e, 0xd1, 0x71,
	0x9c, 0x2d, 0x91, 0x40, 0x70, 0x77, 0xbe, 0xef, 0x7c, 0xf6, 0xf9, 0xb5, 0xe1, 0x38, 0x2d, 0xa5,
	0xae, 0x78, 0x5e, 0xa4, 0x32, 0xbf, 0xd8, 0x56, 0xa5, 0x2e, 0xc9, 0xb8, 0x45, 0x51, 0x0a, 0x77,
	0x3e, 0x88, 0xa2, 0x28, 0x99, 0xb8, 0xde, 0x09, 0xa5, 0x09, 0x81, 0xbe, 0xe4, 0x1b, 0xe1, 0x3b,
	0x13, 0xe7, 0xdc, 0x63, 0xc6, 0xa6, 0xcf, 0x00, 0xac, 0x66, 0x5b, 0xec, 0x89, 0x0f, 0x87, 0x1b,
	0xa1, 0x14, 0x5f, 0x35, 0xa2, 0x06, 0xd2, 0xc7, 0xe0, 0x85, 0xa2, 0x60, 0x42, 0xed, 0x0a, 0x4d,
	0x4e, 0x61, 0x58, 0x19, 0xcb, 0xa8, 0x46, 0xcc, 0x22, 0xfa, 0x14, 0xc6, 0xc1, 0x5a, 0xa4, 0x5f,
	0xfe, 0x21, 0xfb, 0xe1, 0xc0, 0x61, 0x10, 0x47, 0xb3, 0x6a, 0xa5, 0xc8, 0x04, 0xc6, 0x41, 0x29,
	0x35, 0xcf, 0xa5, 0xa8, 0xa2, 0xd0, 0x46, 0x6d, 0x53, 0xe4, 0x04, 0x06, 0xb1, 0xd0, 0x52, 0xf9,
	0x07, 0xc6, 0x57, 0x03, 0xbc, 0x3b, 0x5a, 0xc6, 0x58, 0x8d, 0x6b, 0x68, 0x8b, 0xb0, 0x46, 0xbc,
	0xd7, 0xef, 0xd7, 0x35, 0x9a, 0x18, 0x04, 0xfa, 0x0b, 0xae, 0xd7, 0xfe, 0xa0, 0xe6, 0xd0, 0x26,
	0x0f, 0xc1, 0x4b, 0x74, 0x96, 0xcb, 0x90, 0x6b, 0xee, 0x0f, 0x8d, 0xe3, 0x96, 0xa0, 0x3f, 0x1d,
	0xf0, 0x66, 0x59, 0x66, 0xeb, 0x78, 0x04, 0x10, 0xc4, 0xd1, 0x67, 0x51, 0xa9, 0xbc, 0x94, 0x36,
	0xc5, 0x16, 0x43, 0x5e, 0x01, 0x44, 0x52, 0x8b, 0x6a, 0xc9, 0x53, 0x81, 0x69, 0xba, 0xe7, 0xe3,
	0xe9, 0xf1, 0x45, 0x7b, 0x38, 0x91, 0xd4, 0x4b, 0xd6, 0x12, 0x91, 0xe7, 0xe0, 0x46, 0x0b, 0xe5,
	0xbb, 0x46, 0x7b, 0xaf, 0xab, 0x5d, 0x04, 0xa5, 0x5c, 0xe6, 0x2b, 0x86, 0x0a, 0xf2, 0x12, 0x86,
	0xac, 0xdc, 0x69, 0x81, 0x15, 0xa1, 0x96, 0x74, 0xb4, 0xc6, 0xc5, 0xac, 0x82, 0x3c, 0x81, 0x7e,
	0x18, 0x27, 0xca, 0xd4, 0x39, 0x9e, 0x1e, 0x75, 0x94, 0x61, 0x9c, 0x30, 0xe3, 0xa5, 0xef, 0xa1,
	0x8f, 0xe9, 0x60, 0x57, 0xe2, 0xd6, 0x36, 0x98, 0xee, 0x1d, 0x81, 0xfb, 0x91, 0xa7, 0xb6, 0xd3,
	0x68, 0xe2, 0x46, 0x24, 0x5c, 0x66, 0x57, 0xe5, 0x77, 0xdb, 0xe8, 0x06, 0xd2, 0x02, 0x46, 0x4d,
	0xaa, 0xa8, 0xea, 0xb6, 0xa7, 0x81, 0x18, 0x05, 0xa3, 0x99, 0x2b, 0x07, 0xac, 0x8e, 0xec, 0xc3,
	0xe1, 0x2c, 0xcb, 0x2a, 0xa1, 0x54, 0x73, 0xa7, 0x85, 0xe8, 0x99, 0x73, 0x2d, 0xbe, 0xf1, 0xbd,
	0x1d, 0x60, 0x03, 0xe9, 0x0b, 0x18, 0x98, 0x2a, 0x31, 0xc5, 0x50, 0x69, 0x1b, 0x06, 0x4d, 0x72,
	0x17, 0x0e, 0xe6, 0x97, 0x36, 0xe7, 0x83, 0xf9, 0x25, 0xbd, 0x06, 0x37, 0x8c, 0x13, 0xdc, 0x2c,
	0xac, 0x49, 0x89, 0xea, 0xab, 0xa8, 0x94, 0xef, 0x4c, 0x5c, 0xdc, 0xac, 0x16, 0x85, 0x3b, 0x14,
	0x96, 0x1b, 0x9e, 0x4b, 0x7b, 0xd8, 0x22, 0xe4, 0x13, 0xc1, 0xab, 0x74, 0x6d, 0xe6, 0xe3, 0x31,
	0x8b, 0x30, 0xbb, 0x4f, 0x5b, 0x9d, 0x97, 0xb2, 0x1e, 0x86, 0xc7, 0x1a, 0x38, 0xfd, 0xe5, 0xd4,
	0x6b, 0x8c, 0xdd, 0x0e, 0xe2, 0x88, 0xbc, 0x01, 0x77, 0x96, 0x65, 0xe4, 0xa4, 0x33, 0x02, 0xbb,
	0xf2, 0x67, 0xa7, 0x1d, 0xf6, 0x66, 0xcd, 0x68, 0x0f, 0x8f, 0x85, 0xa2, 0xf8, 0xaf, 0x63, 0x37,
	0x8f, 0x91, 0xf6, 0xc8, 0x5b, 0x18, 0x98, 0x67, 0xf7, 0x97, 0x83, 0x7e, 0x97, 0xbd, 0x7d, 0xa0,
	0xb4, 0x47, 0xde, 0xc1, 0x28, 0xe1, 0x7b, 0xf3, 0x03, 0x90, 0x07, 0x1d, 0x5d, 0xfb, 0xe7, 0x38,
	0xbb, 0xff, 0x27, 0xd7, 0xb6, 0xd8, 0xd3, 0xde, 0xd5, 0xd0, 0x7c, 0x3c, 0xaf, 0x7f, 0x07, 0x00,
	0x00, 0xff, 0xff, 0x8e, 0x57, 0x62, 0x98, 0x8d, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ContrailCNIClient is the client API for ContrailCNI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContrailCNIClient interface {
	Add(ctx context.Context, in *CNIArgs, opts ...grpc.CallOption) (*AddResult, error)
	Del(ctx context.Context, in *CNIArgs, opts ...grpc.CallOption) (*DelResult, error)
	Check(ctx context.Context, in *CNIArgs, opts ...grpc.CallOption) (*CheckResult, error)
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type contrailCNIClient struct {
	cc *grpc.ClientConn
}

func NewContrailCNIClient(cc *grpc.ClientConn) ContrailCNIClient {
	return &contrailCNIClient{cc}
}

func (c *contrailCNIClient) Add(ctx context.Context, in *CNIArgs, opts ...grpc.CallOption) (*AddResult, error) {
	out := new(AddResult)
	err := c.cc.Invoke(ctx, "/contrailcni.ContrailCNI/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contrailCNIClient) Del(ctx context.Context, in *CNIArgs, opts ...grpc.CallOption) (*DelResult, error) {
	out := new(DelResult)
	err := c.cc.Invoke(ctx, "/contrailcni.ContrailCNI/Del", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contrailCNIClient) Check(ctx context.Context, in *CNIArgs, opts ...grpc.CallOption) (*CheckResult, error) {
	out := new(CheckResult)
	err := c.cc.Invoke(ctx, "/contrailcni.ContrailCNI/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contrailCNIClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/contrailcni.ContrailCNI/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContrailCNIServer is the server API for ContrailCNI service.
type ContrailCNIServer interface {
	Add(context.Context, *CNIArgs) (*AddResult, error)
	Del(context.Context, *CNIArgs) (*DelResult, error)
	Check(context.Context, *CNIArgs) (*CheckResult, error)
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

// UnimplementedContrailCNIServer can be embedded to have forward compatible implementations.
type UnimplementedContrailCNIServer struct {
}

func (*UnimplementedContrailCNIServer) Add(ctx context.Context, req *CNIArgs) (*AddResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedContrailCNIServer) Del(ctx context.Context, req *CNIArgs) (*DelResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Del not implemented")
}
func (*UnimplementedContrailCNIServer) Check(ctx context.Context, req *CNIArgs) (*CheckResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (*UnimplementedContrailCNIServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterContrailCNIServer(s *grpc.Server, srv ContrailCNIServer) {
	s.RegisterService(&_ContrailCNI_serviceDesc, srv)
}

func _ContrailCNI_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CNIArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContrailCNIServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contrailcni.ContrailCNI/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContrailCNIServer).Add(ctx, req.(*CNIArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContrailCNI_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CNIArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContrailCNIServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contrailcni.ContrailCNI/Del",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContrailCNIServer).Del(ctx, req.(*CNIArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContrailCNI_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CNIArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContrailCNIServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contrailcni.ContrailCNI/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContrailCNIServer).Check(ctx, req.(*CNIArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContrailCNI_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContrailCNIServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contrailcni.ContrailCNI/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContrailCNIServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContrailCNI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "contrailcni.ContrailCNI",
	HandlerType: (*ContrailCNIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _ContrailCNI_Add_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _ContrailCNI_Del_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _ContrailCNI_Check_Handler,
		},
		{
			MethodName: "SayHello",
			Handler:    _ContrailCNI_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contrailcni.proto",
}
