// Code generated by protoc-gen-go. DO NOT EDIT.
// source: execution/record.proto

package execution

import (
	fmt "fmt"
	_ "github.com/cube2222/octosql"
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

type RecordData struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecordData) Reset()         { *m = RecordData{} }
func (m *RecordData) String() string { return proto.CompactTextString(m) }
func (*RecordData) ProtoMessage()    {}
func (*RecordData) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd7ef896dd4ca095, []int{0}
}

func (m *RecordData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordData.Unmarshal(m, b)
}
func (m *RecordData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordData.Marshal(b, m, deterministic)
}
func (m *RecordData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordData.Merge(m, src)
}
func (m *RecordData) XXX_Size() int {
	return xxx_messageInfo_RecordData.Size(m)
}
func (m *RecordData) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordData.DiscardUnknown(m)
}

var xxx_messageInfo_RecordData proto.InternalMessageInfo

func (m *RecordData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*RecordData)(nil), "execution.RecordData")
}

func init() { proto.RegisterFile("execution/record.proto", fileDescriptor_fd7ef896dd4ca095) }

var fileDescriptor_fd7ef896dd4ca095 = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0xad, 0x48, 0x4d,
	0x2e, 0x2d, 0xc9, 0xcc, 0xcf, 0xd3, 0x2f, 0x4a, 0x4d, 0xce, 0x2f, 0x4a, 0xd1, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x8b, 0x4b, 0xf1, 0x94, 0x25, 0xe6, 0x94, 0xa6, 0x16, 0x43, 0x24,
	0x94, 0x14, 0xb8, 0xb8, 0x82, 0xc0, 0x0a, 0x5d, 0x12, 0x4b, 0x12, 0x85, 0x84, 0xb8, 0x58, 0x52,
	0x12, 0x4b, 0x12, 0x25, 0x18, 0x15, 0x18, 0x35, 0x78, 0x82, 0xc0, 0x6c, 0x27, 0xf5, 0x28, 0xd5,
	0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xe4, 0xd2, 0xa4, 0x54, 0x23,
	0x23, 0x23, 0x23, 0xfd, 0xfc, 0xe4, 0x92, 0xfc, 0xe2, 0xc2, 0x1c, 0x7d, 0xb8, 0xc1, 0x49, 0x6c,
	0x60, 0x13, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x48, 0xc0, 0x56, 0x84, 0x00, 0x00,
	0x00,
}
