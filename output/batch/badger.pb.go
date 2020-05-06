// Code generated by protoc-gen-go. DO NOT EDIT.
// source: output/badger/badger.proto

package batch

import (
	fmt "fmt"
	execution "github.com/cube2222/octosql/execution"
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
	Ids                  []*execution.RecordID `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	IsUndo               bool                  `protobuf:"varint,2,opt,name=isUndo,proto3" json:"isUndo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RecordData) Reset()         { *m = RecordData{} }
func (m *RecordData) String() string { return proto.CompactTextString(m) }
func (*RecordData) ProtoMessage()    {}
func (*RecordData) Descriptor() ([]byte, []int) {
	return fileDescriptor_926b22d4ed6f055f, []int{0}
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

func (m *RecordData) GetIds() []*execution.RecordID {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *RecordData) GetIsUndo() bool {
	if m != nil {
		return m.IsUndo
	}
	return false
}

func init() {
	proto.RegisterType((*RecordData)(nil), "execution.RecordData")
}

func init() { proto.RegisterFile("output/badger/badger.proto", fileDescriptor_926b22d4ed6f055f) }

var fileDescriptor_926b22d4ed6f055f = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xca, 0x2f, 0x2d, 0x29,
	0x28, 0x2d, 0xd1, 0x4f, 0x4a, 0x4c, 0x49, 0x4f, 0x2d, 0x82, 0x52, 0x7a, 0x05, 0x45, 0xf9, 0x25,
	0xf9, 0x42, 0x9c, 0xa9, 0x15, 0xa9, 0xc9, 0xa5, 0x25, 0x99, 0xf9, 0x79, 0x52, 0x62, 0x70, 0xa6,
	0x7e, 0x51, 0x6a, 0x72, 0x7e, 0x51, 0x0a, 0x44, 0x89, 0x92, 0x37, 0x17, 0x57, 0x10, 0x98, 0xef,
	0x92, 0x58, 0x92, 0x28, 0xa4, 0xca, 0xc5, 0x9c, 0x99, 0x52, 0x2c, 0xc1, 0xa8, 0xc0, 0xac, 0xc1,
	0x6d, 0x24, 0xac, 0x07, 0xd7, 0xa3, 0x07, 0x51, 0xe3, 0xe9, 0x12, 0x04, 0x92, 0x17, 0x12, 0xe3,
	0x62, 0xcb, 0x2c, 0x0e, 0xcd, 0x4b, 0xc9, 0x97, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x08, 0x82, 0xf2,
	0x9c, 0xb4, 0xa3, 0x34, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x93,
	0x4b, 0x93, 0x52, 0x8d, 0x8c, 0x8c, 0x8c, 0xf4, 0xf3, 0x93, 0x4b, 0xf2, 0x8b, 0x0b, 0x73, 0xf4,
	0x51, 0x5c, 0x9a, 0xc4, 0x06, 0x76, 0x80, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x09, 0x68,
	0x97, 0xc1, 0x00, 0x00, 0x00,
}