// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: utxo.proto

package v4utxopb

import (
	proto "github.com/golang/protobuf/proto"
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

type Utxo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount        []byte `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	PublicKeyHash []byte `protobuf:"bytes,2,opt,name=public_key_hash,json=publicKeyHash,proto3" json:"public_key_hash,omitempty"`
	Txid          []byte `protobuf:"bytes,3,opt,name=txid,proto3" json:"txid,omitempty"`
	TxIndex       uint32 `protobuf:"varint,4,opt,name=tx_index,json=txIndex,proto3" json:"tx_index,omitempty"`
	UtxoType      uint32 `protobuf:"varint,5,opt,name=utxoType,proto3" json:"utxoType,omitempty"`
	Contract      string `protobuf:"bytes,6,opt,name=contract,proto3" json:"contract,omitempty"`
	NextUtxoKey   []byte `protobuf:"bytes,7,opt,name=nextUtxoKey,proto3" json:"nextUtxoKey,omitempty"`
}

func (x *Utxo) Reset() {
	*x = Utxo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_utxo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Utxo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Utxo) ProtoMessage() {}

func (x *Utxo) ProtoReflect() protoreflect.Message {
	mi := &file_utxo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Utxo.ProtoReflect.Descriptor instead.
func (*Utxo) Descriptor() ([]byte, []int) {
	return file_utxo_proto_rawDescGZIP(), []int{0}
}

func (x *Utxo) GetAmount() []byte {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *Utxo) GetPublicKeyHash() []byte {
	if x != nil {
		return x.PublicKeyHash
	}
	return nil
}

func (x *Utxo) GetTxid() []byte {
	if x != nil {
		return x.Txid
	}
	return nil
}

func (x *Utxo) GetTxIndex() uint32 {
	if x != nil {
		return x.TxIndex
	}
	return 0
}

func (x *Utxo) GetUtxoType() uint32 {
	if x != nil {
		return x.UtxoType
	}
	return 0
}

func (x *Utxo) GetContract() string {
	if x != nil {
		return x.Contract
	}
	return ""
}

func (x *Utxo) GetNextUtxoKey() []byte {
	if x != nil {
		return x.NextUtxoKey
	}
	return nil
}

var File_utxo_proto protoreflect.FileDescriptor

var file_utxo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x74, 0x78, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x75, 0x74,
	0x78, 0x6f, 0x70, 0x62, 0x22, 0xcf, 0x01, 0x0a, 0x04, 0x55, 0x74, 0x78, 0x6f, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f,
	0x6b, 0x65, 0x79, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x48, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x78, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x74, 0x78, 0x69,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x78, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x74, 0x78, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x74, 0x78, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x75, 0x74, 0x78, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x55, 0x74, 0x78, 0x6f,
	0x4b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x55,
	0x74, 0x78, 0x6f, 0x4b, 0x65, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_utxo_proto_rawDescOnce sync.Once
	file_utxo_proto_rawDescData = file_utxo_proto_rawDesc
)

func file_utxo_proto_rawDescGZIP() []byte {
	file_utxo_proto_rawDescOnce.Do(func() {
		file_utxo_proto_rawDescData = protoimpl.X.CompressGZIP(file_utxo_proto_rawDescData)
	})
	return file_utxo_proto_rawDescData
}

var file_utxo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_utxo_proto_goTypes = []interface{}{
	(*Utxo)(nil), // 0: utxopb.Utxo
}
var file_utxo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_utxo_proto_init() }
func file_utxo_proto_init() {
	if File_utxo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_utxo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Utxo); i {
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
			RawDescriptor: file_utxo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_utxo_proto_goTypes,
		DependencyIndexes: file_utxo_proto_depIdxs,
		MessageInfos:      file_utxo_proto_msgTypes,
	}.Build()
	File_utxo_proto = out.File
	file_utxo_proto_rawDesc = nil
	file_utxo_proto_goTypes = nil
	file_utxo_proto_depIdxs = nil
}
