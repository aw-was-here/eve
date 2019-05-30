// Code generated by protoc-gen-go. DO NOT EDIT.
// source: devmodel.proto

package config

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// XXX duplicate of definition in zmet.proto with uniq names (ZCio vs Zio)
// Types of I/O adapters that can be assigned to applications
// Deprecate in future, as we would stop using the old style definiations
type ZCioType int32

const (
	ZCioType_ZCioNop   ZCioType = 0
	ZCioType_ZCioEth   ZCioType = 1
	ZCioType_ZCioUSB   ZCioType = 2
	ZCioType_ZCioCOM   ZCioType = 3
	ZCioType_ZCioHDMI  ZCioType = 4
	ZCioType_ZCioOther ZCioType = 255
)

var ZCioType_name = map[int32]string{
	0:   "ZCioNop",
	1:   "ZCioEth",
	2:   "ZCioUSB",
	3:   "ZCioCOM",
	4:   "ZCioHDMI",
	255: "ZCioOther",
}

var ZCioType_value = map[string]int32{
	"ZCioNop":   0,
	"ZCioEth":   1,
	"ZCioUSB":   2,
	"ZCioCOM":   3,
	"ZCioHDMI":  4,
	"ZCioOther": 255,
}

func (x ZCioType) String() string {
	return proto.EnumName(ZCioType_name, int32(x))
}

func (ZCioType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9fb58492383773ea, []int{0}
}

type SWAdapterType int32

const (
	SWAdapterType_IGNORE SWAdapterType = 0
	SWAdapterType_VLAN   SWAdapterType = 1
	SWAdapterType_BOND   SWAdapterType = 2
)

var SWAdapterType_name = map[int32]string{
	0: "IGNORE",
	1: "VLAN",
	2: "BOND",
}

var SWAdapterType_value = map[string]int32{
	"IGNORE": 0,
	"VLAN":   1,
	"BOND":   2,
}

func (x SWAdapterType) String() string {
	return proto.EnumName(SWAdapterType_name, int32(x))
}

func (SWAdapterType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9fb58492383773ea, []int{1}
}

type PhyIoType int32

const (
	PhyIoType_PhyIoNoop    PhyIoType = 0
	PhyIoType_PhyIoNetEth  PhyIoType = 1
	PhyIoType_PhyIoUSB     PhyIoType = 2
	PhyIoType_PhyIoCOM     PhyIoType = 3
	PhyIoType_PhyIoAudio   PhyIoType = 4
	PhyIoType_PhyIoNetWLAN PhyIoType = 5
	PhyIoType_PhyIoNetWWAN PhyIoType = 6
	PhyIoType_PhyIoHDMI    PhyIoType = 7
	PhyIoType_PhyIoOther   PhyIoType = 255
)

var PhyIoType_name = map[int32]string{
	0:   "PhyIoNoop",
	1:   "PhyIoNetEth",
	2:   "PhyIoUSB",
	3:   "PhyIoCOM",
	4:   "PhyIoAudio",
	5:   "PhyIoNetWLAN",
	6:   "PhyIoNetWWAN",
	7:   "PhyIoHDMI",
	255: "PhyIoOther",
}

var PhyIoType_value = map[string]int32{
	"PhyIoNoop":    0,
	"PhyIoNetEth":  1,
	"PhyIoUSB":     2,
	"PhyIoCOM":     3,
	"PhyIoAudio":   4,
	"PhyIoNetWLAN": 5,
	"PhyIoNetWWAN": 6,
	"PhyIoHDMI":    7,
	"PhyIoOther":   255,
}

func (x PhyIoType) String() string {
	return proto.EnumName(PhyIoType_name, int32(x))
}

func (PhyIoType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9fb58492383773ea, []int{2}
}

// How does EVE should use them, for what purpose it is for
type PhyIoMemberUsage int32

const (
	PhyIoMemberUsage_PhyIoUsageNone      PhyIoMemberUsage = 0
	PhyIoMemberUsage_PhyIoUsageMgmt      PhyIoMemberUsage = 1
	PhyIoMemberUsage_PhyIoUsageShared    PhyIoMemberUsage = 2
	PhyIoMemberUsage_PhyIoUsageDedicated PhyIoMemberUsage = 3
)

var PhyIoMemberUsage_name = map[int32]string{
	0: "PhyIoUsageNone",
	1: "PhyIoUsageMgmt",
	2: "PhyIoUsageShared",
	3: "PhyIoUsageDedicated",
}

var PhyIoMemberUsage_value = map[string]int32{
	"PhyIoUsageNone":      0,
	"PhyIoUsageMgmt":      1,
	"PhyIoUsageShared":    2,
	"PhyIoUsageDedicated": 3,
}

func (x PhyIoMemberUsage) String() string {
	return proto.EnumName(PhyIoMemberUsage_name, int32(x))
}

func (PhyIoMemberUsage) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9fb58492383773ea, []int{3}
}

type SWAdapterParams struct {
	AType SWAdapterType `protobuf:"varint,1,opt,name=aType,proto3,enum=SWAdapterType" json:"aType,omitempty"`
	// vlan
	UnderlayInterface string `protobuf:"bytes,8,opt,name=underlayInterface,proto3" json:"underlayInterface,omitempty"`
	VlanId            uint32 `protobuf:"varint,9,opt,name=vlanId,proto3" json:"vlanId,omitempty"`
	// OR : repeated physical interfaces for bond0
	Bondgroup            []string `protobuf:"bytes,10,rep,name=bondgroup,proto3" json:"bondgroup,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SWAdapterParams) Reset()         { *m = SWAdapterParams{} }
func (m *SWAdapterParams) String() string { return proto.CompactTextString(m) }
func (*SWAdapterParams) ProtoMessage()    {}
func (*SWAdapterParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fb58492383773ea, []int{0}
}

func (m *SWAdapterParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SWAdapterParams.Unmarshal(m, b)
}
func (m *SWAdapterParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SWAdapterParams.Marshal(b, m, deterministic)
}
func (m *SWAdapterParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SWAdapterParams.Merge(m, src)
}
func (m *SWAdapterParams) XXX_Size() int {
	return xxx_messageInfo_SWAdapterParams.Size(m)
}
func (m *SWAdapterParams) XXX_DiscardUnknown() {
	xxx_messageInfo_SWAdapterParams.DiscardUnknown(m)
}

var xxx_messageInfo_SWAdapterParams proto.InternalMessageInfo

func (m *SWAdapterParams) GetAType() SWAdapterType {
	if m != nil {
		return m.AType
	}
	return SWAdapterType_IGNORE
}

func (m *SWAdapterParams) GetUnderlayInterface() string {
	if m != nil {
		return m.UnderlayInterface
	}
	return ""
}

func (m *SWAdapterParams) GetVlanId() uint32 {
	if m != nil {
		return m.VlanId
	}
	return 0
}

func (m *SWAdapterParams) GetBondgroup() []string {
	if m != nil {
		return m.Bondgroup
	}
	return nil
}

// systemAdapters, are the higher l2 concept built on physicalIOs.
// systemAdapters, gives all the required bits to turn the physical IOs
// into useful IP endpoints
// These endpoints can be further used to connect to controller or
// can be shared between workload/services running on the node.
type SystemAdapter struct {
	// name of the adapter; hardware-specific e.g., eth0
	Name         string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AllocDetails *SWAdapterParams `protobuf:"bytes,20,opt,name=allocDetails,proto3" json:"allocDetails,omitempty"`
	// this is part of the freelink group
	// deprecate: look at PhysicalIO->Usage
	FreeUplink bool `protobuf:"varint,2,opt,name=freeUplink,proto3" json:"freeUplink,omitempty"`
	// this is part of the uplink group
	// deprecate: look at PhysicalIO->Usage
	Uplink bool `protobuf:"varint,3,opt,name=uplink,proto3" json:"uplink,omitempty"`
	// attach this network config for this adapter
	NetworkUUID string `protobuf:"bytes,4,opt,name=networkUUID,proto3" json:"networkUUID,omitempty"`
	// if its static network we need ip address
	Addr string `protobuf:"bytes,5,opt,name=addr,proto3" json:"addr,omitempty"`
	// alias/logical name which will be reported to zedcloud
	// and used for app instances
	LogicalName          string   `protobuf:"bytes,6,opt,name=logicalName,proto3" json:"logicalName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SystemAdapter) Reset()         { *m = SystemAdapter{} }
func (m *SystemAdapter) String() string { return proto.CompactTextString(m) }
func (*SystemAdapter) ProtoMessage()    {}
func (*SystemAdapter) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fb58492383773ea, []int{1}
}

func (m *SystemAdapter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemAdapter.Unmarshal(m, b)
}
func (m *SystemAdapter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemAdapter.Marshal(b, m, deterministic)
}
func (m *SystemAdapter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemAdapter.Merge(m, src)
}
func (m *SystemAdapter) XXX_Size() int {
	return xxx_messageInfo_SystemAdapter.Size(m)
}
func (m *SystemAdapter) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemAdapter.DiscardUnknown(m)
}

var xxx_messageInfo_SystemAdapter proto.InternalMessageInfo

func (m *SystemAdapter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SystemAdapter) GetAllocDetails() *SWAdapterParams {
	if m != nil {
		return m.AllocDetails
	}
	return nil
}

func (m *SystemAdapter) GetFreeUplink() bool {
	if m != nil {
		return m.FreeUplink
	}
	return false
}

func (m *SystemAdapter) GetUplink() bool {
	if m != nil {
		return m.Uplink
	}
	return false
}

func (m *SystemAdapter) GetNetworkUUID() string {
	if m != nil {
		return m.NetworkUUID
	}
	return ""
}

func (m *SystemAdapter) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *SystemAdapter) GetLogicalName() string {
	if m != nil {
		return m.LogicalName
	}
	return ""
}

// Given additional details for EVE softwar to how to treat this
// interface. Example policies could be limit use of LTE interface
// or only use Eth1 only if Eth0 is not available etc
type PhyIOUsagePolicy struct {
	// Used only when one other normal uplinks are down
	FreeUplink           bool     `protobuf:"varint,1,opt,name=freeUplink,proto3" json:"freeUplink,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PhyIOUsagePolicy) Reset()         { *m = PhyIOUsagePolicy{} }
func (m *PhyIOUsagePolicy) String() string { return proto.CompactTextString(m) }
func (*PhyIOUsagePolicy) ProtoMessage()    {}
func (*PhyIOUsagePolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fb58492383773ea, []int{2}
}

func (m *PhyIOUsagePolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PhyIOUsagePolicy.Unmarshal(m, b)
}
func (m *PhyIOUsagePolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PhyIOUsagePolicy.Marshal(b, m, deterministic)
}
func (m *PhyIOUsagePolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhyIOUsagePolicy.Merge(m, src)
}
func (m *PhyIOUsagePolicy) XXX_Size() int {
	return xxx_messageInfo_PhyIOUsagePolicy.Size(m)
}
func (m *PhyIOUsagePolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_PhyIOUsagePolicy.DiscardUnknown(m)
}

var xxx_messageInfo_PhyIOUsagePolicy proto.InternalMessageInfo

func (m *PhyIOUsagePolicy) GetFreeUplink() bool {
	if m != nil {
		return m.FreeUplink
	}
	return false
}

// PhysicalIO:
//    Absolute low level description of physical buses and ports that are
//    available on given platfrom.
//    Collection of these IOs, connstitue what we would call as hardware
//    model. Each physical IO is manageable and visible to EVE software, and
//    it can be further configured to either provide IP connectivity or
//    directly be given to workloads
type PhysicalIO struct {
	Ptype PhyIoType `protobuf:"varint,1,opt,name=ptype,proto3,enum=PhyIoType" json:"ptype,omitempty"`
	// physical label typically printed on box.
	// Example Eth0, Eth1, Wifi0, ComA, ComB
	Phylabel string `protobuf:"bytes,2,opt,name=phylabel,proto3" json:"phylabel,omitempty"`
	// the hardware bus address
	Phyaddrs map[string]string `protobuf:"bytes,3,rep,name=phyaddrs,proto3" json:"phyaddrs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// provides the ability to model designer to rename the physicalIO
	// port to more understandable
	// For example Eth0->Mgmt0
	//  or USBA->ConfigDiskA etc
	Logicallabel string `protobuf:"bytes,4,opt,name=logicallabel,proto3" json:"logicallabel,omitempty"`
	// Assignment Group, is unique label that is applied across PhysicalIOs
	// EntireGroup can be assigned to application or nothing at all
	Assigngrp   string            `protobuf:"bytes,5,opt,name=assigngrp,proto3" json:"assigngrp,omitempty"`
	Usage       PhyIoMemberUsage  `protobuf:"varint,6,opt,name=usage,proto3,enum=PhyIoMemberUsage" json:"usage,omitempty"`
	UsagePolicy *PhyIOUsagePolicy `protobuf:"bytes,7,opt,name=usagePolicy,proto3" json:"usagePolicy,omitempty"`
	// physical and logical attributes
	//    For example in WWAN to which firmware version to laod etc
	Cbattr               map[string]string `protobuf:"bytes,8,rep,name=cbattr,proto3" json:"cbattr,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PhysicalIO) Reset()         { *m = PhysicalIO{} }
func (m *PhysicalIO) String() string { return proto.CompactTextString(m) }
func (*PhysicalIO) ProtoMessage()    {}
func (*PhysicalIO) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fb58492383773ea, []int{3}
}

func (m *PhysicalIO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PhysicalIO.Unmarshal(m, b)
}
func (m *PhysicalIO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PhysicalIO.Marshal(b, m, deterministic)
}
func (m *PhysicalIO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhysicalIO.Merge(m, src)
}
func (m *PhysicalIO) XXX_Size() int {
	return xxx_messageInfo_PhysicalIO.Size(m)
}
func (m *PhysicalIO) XXX_DiscardUnknown() {
	xxx_messageInfo_PhysicalIO.DiscardUnknown(m)
}

var xxx_messageInfo_PhysicalIO proto.InternalMessageInfo

func (m *PhysicalIO) GetPtype() PhyIoType {
	if m != nil {
		return m.Ptype
	}
	return PhyIoType_PhyIoNoop
}

func (m *PhysicalIO) GetPhylabel() string {
	if m != nil {
		return m.Phylabel
	}
	return ""
}

func (m *PhysicalIO) GetPhyaddrs() map[string]string {
	if m != nil {
		return m.Phyaddrs
	}
	return nil
}

func (m *PhysicalIO) GetLogicallabel() string {
	if m != nil {
		return m.Logicallabel
	}
	return ""
}

func (m *PhysicalIO) GetAssigngrp() string {
	if m != nil {
		return m.Assigngrp
	}
	return ""
}

func (m *PhysicalIO) GetUsage() PhyIoMemberUsage {
	if m != nil {
		return m.Usage
	}
	return PhyIoMemberUsage_PhyIoUsageNone
}

func (m *PhysicalIO) GetUsagePolicy() *PhyIOUsagePolicy {
	if m != nil {
		return m.UsagePolicy
	}
	return nil
}

func (m *PhysicalIO) GetCbattr() map[string]string {
	if m != nil {
		return m.Cbattr
	}
	return nil
}

func init() {
	proto.RegisterEnum("ZCioType", ZCioType_name, ZCioType_value)
	proto.RegisterEnum("SWAdapterType", SWAdapterType_name, SWAdapterType_value)
	proto.RegisterEnum("PhyIoType", PhyIoType_name, PhyIoType_value)
	proto.RegisterEnum("PhyIoMemberUsage", PhyIoMemberUsage_name, PhyIoMemberUsage_value)
	proto.RegisterType((*SWAdapterParams)(nil), "sWAdapterParams")
	proto.RegisterType((*SystemAdapter)(nil), "SystemAdapter")
	proto.RegisterType((*PhyIOUsagePolicy)(nil), "PhyIOUsagePolicy")
	proto.RegisterType((*PhysicalIO)(nil), "PhysicalIO")
	proto.RegisterMapType((map[string]string)(nil), "PhysicalIO.CbattrEntry")
	proto.RegisterMapType((map[string]string)(nil), "PhysicalIO.PhyaddrsEntry")
}

func init() { proto.RegisterFile("devmodel.proto", fileDescriptor_9fb58492383773ea) }

var fileDescriptor_9fb58492383773ea = []byte{
	// 757 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdb, 0x6e, 0xeb, 0x44,
	0x14, 0xad, 0x73, 0x6b, 0xb2, 0x73, 0xe9, 0x9c, 0xa1, 0xe2, 0x98, 0x0a, 0x81, 0x15, 0x1d, 0x89,
	0xa8, 0x02, 0x47, 0xca, 0x01, 0x89, 0xcb, 0x53, 0xdb, 0x54, 0x87, 0x48, 0x34, 0xa9, 0x5c, 0x42,
	0xa5, 0x23, 0x5e, 0x26, 0xf6, 0xae, 0x63, 0xd5, 0xf1, 0x58, 0xe3, 0x71, 0x90, 0xf9, 0x95, 0xfe,
	0x11, 0x7f, 0xc2, 0x57, 0x80, 0x66, 0xc6, 0x49, 0x9c, 0xf2, 0xc4, 0xdb, 0x5e, 0x6b, 0x5f, 0xb2,
	0xd6, 0xde, 0xe3, 0xc0, 0x20, 0xc0, 0xed, 0x86, 0x07, 0x18, 0xbb, 0xa9, 0xe0, 0x92, 0x0f, 0x5f,
	0x2c, 0x38, 0xcb, 0x1e, 0xaf, 0x02, 0x96, 0x4a, 0x14, 0xf7, 0x4c, 0xb0, 0x4d, 0x46, 0xdf, 0x41,
	0x93, 0xfd, 0x5a, 0xa4, 0x68, 0x5b, 0x8e, 0x35, 0x1a, 0x4c, 0x06, 0xee, 0xbe, 0x40, 0xb1, 0x9e,
	0x49, 0xd2, 0xaf, 0xe1, 0x4d, 0x9e, 0x04, 0x28, 0x62, 0x56, 0xcc, 0x12, 0x89, 0xe2, 0x89, 0xf9,
	0x68, 0xb7, 0x1d, 0x6b, 0xd4, 0xf1, 0xfe, 0x9b, 0xa0, 0x9f, 0x42, 0x6b, 0x1b, 0xb3, 0x64, 0x16,
	0xd8, 0x1d, 0xc7, 0x1a, 0xf5, 0xbd, 0x12, 0xd1, 0xcf, 0xa1, 0xb3, 0xe2, 0x49, 0x10, 0x0a, 0x9e,
	0xa7, 0x36, 0x38, 0xf5, 0x51, 0xc7, 0x3b, 0x10, 0xc3, 0xbf, 0x2d, 0xe8, 0x3f, 0x14, 0x99, 0xc4,
	0x4d, 0x29, 0x80, 0x52, 0x68, 0x24, 0x6c, 0x63, 0xa4, 0x75, 0x3c, 0x1d, 0xd3, 0x6f, 0xa1, 0xc7,
	0xe2, 0x98, 0xfb, 0x53, 0x94, 0x2c, 0x8a, 0x33, 0xfb, 0xdc, 0xb1, 0x46, 0xdd, 0x09, 0x71, 0x5f,
	0xf9, 0xf2, 0x8e, 0xaa, 0xe8, 0x17, 0x00, 0x4f, 0x02, 0x71, 0x99, 0xc6, 0x51, 0xf2, 0x6c, 0xd7,
	0x1c, 0x6b, 0xd4, 0xf6, 0x2a, 0x8c, 0x52, 0x9c, 0x9b, 0x5c, 0x5d, 0xe7, 0x4a, 0x44, 0x1d, 0xe8,
	0x26, 0x28, 0xff, 0xe0, 0xe2, 0x79, 0xb9, 0x9c, 0x4d, 0xed, 0x86, 0x16, 0x52, 0xa5, 0x94, 0x46,
	0x16, 0x04, 0xc2, 0x6e, 0x1a, 0x8d, 0x2a, 0x56, 0x5d, 0x31, 0x0f, 0x23, 0x9f, 0xc5, 0x73, 0x25,
	0xbf, 0x65, 0xba, 0x2a, 0xd4, 0x70, 0x02, 0xe4, 0x7e, 0x5d, 0xcc, 0x16, 0xcb, 0x8c, 0x85, 0x78,
	0xcf, 0xe3, 0xc8, 0x2f, 0x5e, 0x69, 0xb4, 0x5e, 0x6b, 0x1c, 0xfe, 0x55, 0x07, 0xb8, 0x5f, 0x17,
	0x99, 0x1a, 0x32, 0x5b, 0x50, 0x07, 0x9a, 0xa9, 0x3c, 0x1c, 0x0e, 0x5c, 0x35, 0x90, 0x9b, 0xa3,
	0xe9, 0x04, 0xbd, 0x80, 0x76, 0xba, 0x2e, 0x62, 0xb6, 0xc2, 0x58, 0x5b, 0xee, 0x78, 0x7b, 0x4c,
	0xbf, 0xd3, 0x39, 0xa5, 0x36, 0xb3, 0xeb, 0x4e, 0x7d, 0xd4, 0x9d, 0x7c, 0xe6, 0x1e, 0x86, 0xab,
	0x50, 0xe7, 0x6e, 0x13, 0x29, 0x0a, 0x6f, 0x5f, 0x4a, 0x87, 0xd0, 0x2b, 0x6d, 0x98, 0xb1, 0x66,
	0x21, 0x47, 0x9c, 0xba, 0x32, 0xcb, 0xb2, 0x28, 0x4c, 0x42, 0x91, 0x96, 0x6b, 0x39, 0x10, 0xf4,
	0x2b, 0x68, 0xe6, 0xca, 0xb4, 0xde, 0xca, 0x60, 0xf2, 0xc6, 0xc8, 0xbe, 0xc3, 0xcd, 0x0a, 0x85,
	0xde, 0x86, 0x67, 0xf2, 0xf4, 0x3d, 0x74, 0xf3, 0xc3, 0x76, 0xec, 0x53, 0x7d, 0x67, 0x53, 0x5e,
	0x5d, 0x9b, 0x57, 0xad, 0xa2, 0x63, 0x68, 0xf9, 0x2b, 0x26, 0xa5, 0xb0, 0xdb, 0xda, 0xd4, 0xdb,
	0xaa, 0xa9, 0x1b, 0x9d, 0x31, 0x96, 0xca, 0xb2, 0x8b, 0x9f, 0xa0, 0x7f, 0xe4, 0x95, 0x12, 0xa8,
	0x3f, 0x63, 0x51, 0x3e, 0x39, 0x15, 0xd2, 0x73, 0x68, 0x6e, 0x59, 0x9c, 0x63, 0xb9, 0x43, 0x03,
	0x7e, 0xac, 0x7d, 0x6f, 0x5d, 0xfc, 0x00, 0xdd, 0xca, 0xcc, 0xff, 0xd3, 0x7a, 0xf9, 0x3b, 0xb4,
	0x3f, 0xde, 0x44, 0xfa, 0x5c, 0xb4, 0x0b, 0xa7, 0x2a, 0x9e, 0xf3, 0x94, 0x9c, 0xec, 0xc0, 0xad,
	0x5c, 0x13, 0x6b, 0x07, 0x96, 0x0f, 0xd7, 0xa4, 0xb6, 0x03, 0x37, 0x8b, 0x3b, 0x52, 0xa7, 0x3d,
	0xd3, 0xff, 0xf3, 0xf4, 0x6e, 0x46, 0x1a, 0x74, 0x00, 0x1d, 0x85, 0x16, 0x72, 0x8d, 0x82, 0xfc,
	0x63, 0x5d, 0x8e, 0xa1, 0x7f, 0xf4, 0x19, 0x53, 0x80, 0xd6, 0xec, 0xc3, 0x7c, 0xe1, 0xdd, 0x92,
	0x13, 0xda, 0x86, 0xc6, 0x6f, 0xbf, 0x5c, 0xcd, 0x89, 0xa5, 0xa2, 0xeb, 0xc5, 0x7c, 0x4a, 0x6a,
	0x97, 0x2f, 0x16, 0x74, 0xf6, 0xef, 0x87, 0xf6, 0x4b, 0x30, 0xe7, 0x5a, 0xd2, 0x19, 0x74, 0x0d,
	0x44, 0x69, 0x64, 0xf5, 0xa0, 0xad, 0x09, 0xa3, 0x6b, 0x87, 0x8c, 0xb0, 0x81, 0x7e, 0xa4, 0x33,
	0x7e, 0x95, 0x07, 0x11, 0x27, 0x0d, 0x4a, 0xa0, 0xb7, 0x6b, 0x7e, 0x54, 0xbf, 0xda, 0x3c, 0x62,
	0x1e, 0xaf, 0xe6, 0xa4, 0xb5, 0xff, 0x3d, 0xed, 0xe6, 0x94, 0x9e, 0x95, 0x23, 0xf6, 0x76, 0x22,
	0xf3, 0xb5, 0x54, 0x5f, 0x09, 0xa5, 0x30, 0x30, 0x1a, 0x14, 0x9a, 0xf3, 0x04, 0xc9, 0xc9, 0x31,
	0x77, 0x17, 0x6e, 0x24, 0xb1, 0xe8, 0x79, 0xd9, 0xab, 0xb9, 0x87, 0x35, 0x13, 0x18, 0x90, 0x1a,
	0x7d, 0x0b, 0x9f, 0x1c, 0xd8, 0x29, 0x06, 0x91, 0xcf, 0x24, 0x06, 0xa4, 0x7e, 0xfd, 0x01, 0xbe,
	0xf4, 0xf9, 0xc6, 0xfd, 0x13, 0x03, 0x0c, 0x98, 0xeb, 0xc7, 0x3c, 0x0f, 0xdc, 0x3c, 0x43, 0xb1,
	0x8d, 0x7c, 0x34, 0xff, 0xa2, 0x1f, 0xdf, 0x85, 0x91, 0x5c, 0xe7, 0x2b, 0xd7, 0xe7, 0x9b, 0x71,
	0xfc, 0xf4, 0x0d, 0x06, 0x21, 0x8e, 0x71, 0x8b, 0x63, 0x96, 0x46, 0xe3, 0x90, 0x8f, 0x7d, 0x9e,
	0x3c, 0x45, 0xe1, 0xaa, 0xa5, 0x8b, 0xdf, 0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x23, 0x0d,
	0x34, 0x84, 0x05, 0x00, 0x00,
}