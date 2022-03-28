// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types.proto

package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PersonalInformation struct {
	// @gotags: gorm:"primaryKey;column:id"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// @gotags: gorm:"column:name"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// @gotags: gorm:"column:sex"
	Sex string `protobuf:"bytes,3,opt,name=sex,proto3" json:"sex,omitempty"`
	// @gotags: gorm:"column:tall"
	Tall float32 `protobuf:"fixed32,4,opt,name=tall,proto3" json:"tall,omitempty"`
	// @gotags: gorm:"column:weight"
	Weight float32 `protobuf:"fixed32,5,opt,name=weight,proto3" json:"weight,omitempty"`
	// @gotags: gorm:"column:age"
	Age                  int64    `protobuf:"varint,6,opt,name=age,proto3" json:"age,omitempty"`
	Nation               string   `protobuf:"bytes,7,opt,name=nation,proto3" json:"nation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PersonalInformation) Reset()         { *m = PersonalInformation{} }
func (m *PersonalInformation) String() string { return proto.CompactTextString(m) }
func (*PersonalInformation) ProtoMessage()    {}
func (*PersonalInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_0c5ad1892002dfbb, []int{0}
}
func (m *PersonalInformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonalInformation.Unmarshal(m, b)
}
func (m *PersonalInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonalInformation.Marshal(b, m, deterministic)
}
func (dst *PersonalInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonalInformation.Merge(dst, src)
}
func (m *PersonalInformation) XXX_Size() int {
	return xxx_messageInfo_PersonalInformation.Size(m)
}
func (m *PersonalInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonalInformation.DiscardUnknown(m)
}

var xxx_messageInfo_PersonalInformation proto.InternalMessageInfo

func (m *PersonalInformation) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PersonalInformation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PersonalInformation) GetSex() string {
	if m != nil {
		return m.Sex
	}
	return ""
}

func (m *PersonalInformation) GetTall() float32 {
	if m != nil {
		return m.Tall
	}
	return 0
}

func (m *PersonalInformation) GetWeight() float32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *PersonalInformation) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *PersonalInformation) GetNation() string {
	if m != nil {
		return m.Nation
	}
	return ""
}

func init() {
	proto.RegisterType((*PersonalInformation)(nil), "main.PersonalInformation")
}

func init() { proto.RegisterFile("types.proto", fileDescriptor_types_0c5ad1892002dfbb) }

var fileDescriptor_types_0c5ad1892002dfbb = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0x8e, 0xc1, 0xca, 0xc2, 0x40,
	0x0c, 0x84, 0xd9, 0x6d, 0xff, 0xfe, 0x18, 0x41, 0x24, 0x82, 0xe4, 0x58, 0x3c, 0xf5, 0xe4, 0xc5,
	0xa7, 0xf0, 0x26, 0x7d, 0x83, 0x48, 0x63, 0x5d, 0x68, 0xb3, 0xa5, 0x5d, 0x50, 0x1f, 0xc7, 0x37,
	0x95, 0x0d, 0xbd, 0x7d, 0x33, 0x4c, 0x32, 0x03, 0xdb, 0xf4, 0x99, 0x64, 0x39, 0x4f, 0x73, 0x4c,
	0x11, 0xcb, 0x91, 0x83, 0x9e, 0xbe, 0x0e, 0x0e, 0x37, 0x99, 0x97, 0xa8, 0x3c, 0x5c, 0xf5, 0x11,
	0xe7, 0x91, 0x53, 0x88, 0x8a, 0x3b, 0xf0, 0xa1, 0x23, 0x57, 0xbb, 0xa6, 0x68, 0x7d, 0xe8, 0x10,
	0xa1, 0x54, 0x1e, 0x85, 0x7c, 0xed, 0x9a, 0x4d, 0x6b, 0x8c, 0x7b, 0x28, 0x16, 0x79, 0x53, 0x61,
	0x56, 0xc6, 0x9c, 0x4a, 0x3c, 0x0c, 0x54, 0xd6, 0xae, 0xf1, 0xad, 0x31, 0x1e, 0xa1, 0x7a, 0x49,
	0xe8, 0x9f, 0x89, 0xfe, 0xcc, 0x5d, 0x55, 0xbe, 0xe6, 0x5e, 0xa8, 0xb2, 0x8a, 0x8c, 0x39, 0xa9,
	0xd6, 0x4e, 0xff, 0xf6, 0x72, 0x55, 0xf7, 0xca, 0x06, 0x5f, 0x7e, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xb9, 0xf3, 0x00, 0x3b, 0xbf, 0x00, 0x00, 0x00,
}
