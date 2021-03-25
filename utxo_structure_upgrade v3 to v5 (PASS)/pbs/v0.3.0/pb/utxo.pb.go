// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/dappley/go-dappley/core/utxo/pb/utxo.proto

package v3utxopb

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Utxo struct {
	Amount               []byte   `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	PublicKeyHash        []byte   `protobuf:"bytes,2,opt,name=public_key_hash,json=publicKeyHash,proto3" json:"public_key_hash,omitempty"`
	Txid                 []byte   `protobuf:"bytes,3,opt,name=txid,proto3" json:"txid,omitempty"`
	TxIndex              uint32   `protobuf:"varint,4,opt,name=tx_index,json=txIndex,proto3" json:"tx_index,omitempty"`
	UtxoType             uint32   `protobuf:"varint,5,opt,name=utxoType,proto3" json:"utxoType,omitempty"`
	Contract             string   `protobuf:"bytes,6,opt,name=contract,proto3" json:"contract,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Utxo) Reset()         { *m = Utxo{} }
func (m *Utxo) String() string { return proto.CompactTextString(m) }
func (*Utxo) ProtoMessage()    {}
func (*Utxo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8df3b4ec976b2cee, []int{0}
}

func (m *Utxo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Utxo.Unmarshal(m, b)
}
func (m *Utxo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Utxo.Marshal(b, m, deterministic)
}
func (m *Utxo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Utxo.Merge(m, src)
}
func (m *Utxo) XXX_Size() int {
	return xxx_messageInfo_Utxo.Size(m)
}
func (m *Utxo) XXX_DiscardUnknown() {
	xxx_messageInfo_Utxo.DiscardUnknown(m)
}

var xxx_messageInfo_Utxo proto.InternalMessageInfo

func (m *Utxo) GetAmount() []byte {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *Utxo) GetPublicKeyHash() []byte {
	if m != nil {
		return m.PublicKeyHash
	}
	return nil
}

func (m *Utxo) GetTxid() []byte {
	if m != nil {
		return m.Txid
	}
	return nil
}

func (m *Utxo) GetTxIndex() uint32 {
	if m != nil {
		return m.TxIndex
	}
	return 0
}

func (m *Utxo) GetUtxoType() uint32 {
	if m != nil {
		return m.UtxoType
	}
	return 0
}

func (m *Utxo) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

type UtxoList struct {
	Utxos                []*Utxo  `protobuf:"bytes,1,rep,name=utxos,proto3" json:"utxos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UtxoList) Reset()         { *m = UtxoList{} }
func (m *UtxoList) String() string { return proto.CompactTextString(m) }
func (*UtxoList) ProtoMessage()    {}
func (*UtxoList) Descriptor() ([]byte, []int) {
	return fileDescriptor_8df3b4ec976b2cee, []int{1}
}

func (m *UtxoList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UtxoList.Unmarshal(m, b)
}
func (m *UtxoList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UtxoList.Marshal(b, m, deterministic)
}
func (m *UtxoList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UtxoList.Merge(m, src)
}
func (m *UtxoList) XXX_Size() int {
	return xxx_messageInfo_UtxoList.Size(m)
}
func (m *UtxoList) XXX_DiscardUnknown() {
	xxx_messageInfo_UtxoList.DiscardUnknown(m)
}

var xxx_messageInfo_UtxoList proto.InternalMessageInfo

func (m *UtxoList) GetUtxos() []*Utxo {
	if m != nil {
		return m.Utxos
	}
	return nil
}

func init() {
	proto.RegisterType((*Utxo)(nil), "utxopb.Utxo")
	proto.RegisterType((*UtxoList)(nil), "utxopb.UtxoList")
}

func init() {
	proto.RegisterFile("github.com/dappley/go-dappley/core/utxo/pb/utxo.proto", fileDescriptor_8df3b4ec976b2cee)
}

var fileDescriptor_8df3b4ec976b2cee = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x89, 0xdb, 0xad, 0x75, 0xdc, 0x45, 0xc8, 0x41, 0xa2, 0xa7, 0xd2, 0x83, 0xf4, 0x62,
	0x0b, 0x8a, 0xef, 0xa0, 0xe8, 0xa9, 0xe8, 0xb9, 0x24, 0x69, 0xd8, 0x06, 0x77, 0x3b, 0xa1, 0x9d,
	0x42, 0xfa, 0x50, 0xbe, 0xa3, 0x24, 0xd5, 0x3d, 0xe5, 0xff, 0xf2, 0xcd, 0xc0, 0xfc, 0xf0, 0x72,
	0xb0, 0xd4, 0xcf, 0xaa, 0xd2, 0x78, 0xaa, 0x3b, 0xe9, 0xdc, 0xd1, 0x2c, 0xf5, 0x01, 0x1f, 0xff,
	0xa3, 0xc6, 0xd1, 0xd4, 0x33, 0x79, 0xac, 0x9d, 0x8a, 0x6f, 0xe5, 0x46, 0x24, 0xe4, 0x69, 0xc8,
	0x4e, 0x15, 0x3f, 0x0c, 0x92, 0x2f, 0xf2, 0xc8, 0x6f, 0x21, 0x95, 0x27, 0x9c, 0x07, 0x12, 0x2c,
	0x67, 0xe5, 0xae, 0xf9, 0x23, 0xfe, 0x00, 0x37, 0x6e, 0x56, 0x47, 0xab, 0xdb, 0x6f, 0xb3, 0xb4,
	0xbd, 0x9c, 0x7a, 0x71, 0x11, 0x07, 0xf6, 0xeb, 0xf7, 0xbb, 0x59, 0x5e, 0xe5, 0xd4, 0x73, 0x0e,
	0x09, 0x79, 0xdb, 0x89, 0x4d, 0x94, 0x31, 0xf3, 0x3b, 0xc8, 0xc8, 0xb7, 0x76, 0xe8, 0x8c, 0x17,
	0x49, 0xce, 0xca, 0x7d, 0x73, 0x49, 0xfe, 0x2d, 0x20, 0xbf, 0x87, 0x2c, 0x5c, 0xf0, 0xb9, 0x38,
	0x23, 0xb6, 0x51, 0x9d, 0x39, 0x38, 0x8d, 0x03, 0x8d, 0x52, 0x93, 0x48, 0x73, 0x56, 0x5e, 0x35,
	0x67, 0x2e, 0x2a, 0xc8, 0xc2, 0xb9, 0x1f, 0x76, 0x22, 0x5e, 0xc0, 0x36, 0xec, 0x4c, 0x82, 0xe5,
	0x9b, 0xf2, 0xfa, 0x69, 0x57, 0xad, 0x9d, 0xaa, 0x30, 0xd0, 0xac, 0x4a, 0xa5, 0xb1, 0xee, 0xf3,
	0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x35, 0x1b, 0xf1, 0x27, 0x01, 0x00, 0x00,
}
