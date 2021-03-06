// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ping/pb/ping.proto

package pb

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

type Ping struct {
	ReqID                []byte   `protobuf:"bytes,1,opt,name=ReqID,proto3" json:"ReqID,omitempty"`
	Req                  bool     `protobuf:"varint,2,opt,name=Req,proto3" json:"Req,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_ping_b16150388e45212b, []int{0}
}
func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (dst *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(dst, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetReqID() []byte {
	if m != nil {
		return m.ReqID
	}
	return nil
}

func (m *Ping) GetReq() bool {
	if m != nil {
		return m.Req
	}
	return false
}

func (m *Ping) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Ping)(nil), "pb.Ping")
}

func init() { proto.RegisterFile("ping/pb/ping.proto", fileDescriptor_ping_b16150388e45212b) }

var fileDescriptor_ping_b16150388e45212b = []byte{
	// 118 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0xc8, 0xcc, 0x4b,
	0xd7, 0x2f, 0x48, 0xd2, 0x07, 0xd1, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x4c, 0x05, 0x49,
	0x4a, 0x1e, 0x5c, 0x2c, 0x01, 0x99, 0x79, 0xe9, 0x42, 0x22, 0x5c, 0xac, 0x41, 0xa9, 0x85, 0x9e,
	0x2e, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x10, 0x8e, 0x90, 0x00, 0x17, 0x73, 0x50, 0x6a,
	0xa1, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x47, 0x10, 0x88, 0x29, 0x24, 0xc1, 0xc5, 0xee, 0x9b, 0x5a,
	0x5c, 0x9c, 0x98, 0x9e, 0x2a, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe3, 0x3a, 0xb1, 0x44,
	0x31, 0x15, 0x24, 0x25, 0xb1, 0x81, 0x8d, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x91, 0x7d,
	0xfd, 0x39, 0x70, 0x00, 0x00, 0x00,
}
