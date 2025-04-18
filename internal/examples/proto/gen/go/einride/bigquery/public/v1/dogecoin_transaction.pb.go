// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: einride/bigquery/public/v1/dogecoin_transaction.proto

package publicv1

import (
	date "google.golang.org/genproto/googleapis/type/date"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Protobuf schema for the BigQuery public table:
//
//	bigquery-public-data.crypto_dogecoin.transactions
type DogecoinTransaction struct {
	state               protoimpl.MessageState        `protogen:"open.v1"`
	Hash                string                        `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`                                                            // STRING REQUIRED
	Size                int64                         `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`                                                           // INTEGER NULLABLE
	VirtualSize         int64                         `protobuf:"varint,3,opt,name=virtual_size,json=virtualSize,proto3" json:"virtual_size,omitempty"`                          // INTEGER NULLABLE
	Version             int64                         `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`                                                     // INTEGER NULLABLE
	LockTime            int64                         `protobuf:"varint,5,opt,name=lock_time,json=lockTime,proto3" json:"lock_time,omitempty"`                                   // INTEGER NULLABLE
	BlockHash           string                        `protobuf:"bytes,6,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`                                 // STRING REQUIRED
	BlockNumber         int64                         `protobuf:"varint,7,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`                          // INTEGER REQUIRED
	BlockTimestamp      *timestamppb.Timestamp        `protobuf:"bytes,8,opt,name=block_timestamp,json=blockTimestamp,proto3" json:"block_timestamp,omitempty"`                  // TIMESTAMP REQUIRED
	BlockTimestampMonth *date.Date                    `protobuf:"bytes,9,opt,name=block_timestamp_month,json=blockTimestampMonth,proto3" json:"block_timestamp_month,omitempty"` // DATE REQUIRED
	InputCount          int64                         `protobuf:"varint,10,opt,name=input_count,json=inputCount,proto3" json:"input_count,omitempty"`                            // INTEGER NULLABLE
	OutputCount         int64                         `protobuf:"varint,11,opt,name=output_count,json=outputCount,proto3" json:"output_count,omitempty"`                         // INTEGER NULLABLE
	IsCoinbase          bool                          `protobuf:"varint,14,opt,name=is_coinbase,json=isCoinbase,proto3" json:"is_coinbase,omitempty"`                            // BOOLEAN NULLABLE
	Inputs              []*DogecoinTransaction_Input  `protobuf:"bytes,16,rep,name=inputs,proto3" json:"inputs,omitempty"`                                                       // RECORD REPEATED
	Outputs             []*DogecoinTransaction_Output `protobuf:"bytes,17,rep,name=outputs,proto3" json:"outputs,omitempty"`                                                     // RECORD REPEATED
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *DogecoinTransaction) Reset() {
	*x = DogecoinTransaction{}
	mi := &file_einride_bigquery_public_v1_dogecoin_transaction_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DogecoinTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DogecoinTransaction) ProtoMessage() {}

func (x *DogecoinTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_einride_bigquery_public_v1_dogecoin_transaction_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DogecoinTransaction.ProtoReflect.Descriptor instead.
func (*DogecoinTransaction) Descriptor() ([]byte, []int) {
	return file_einride_bigquery_public_v1_dogecoin_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *DogecoinTransaction) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *DogecoinTransaction) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *DogecoinTransaction) GetVirtualSize() int64 {
	if x != nil {
		return x.VirtualSize
	}
	return 0
}

func (x *DogecoinTransaction) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *DogecoinTransaction) GetLockTime() int64 {
	if x != nil {
		return x.LockTime
	}
	return 0
}

func (x *DogecoinTransaction) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *DogecoinTransaction) GetBlockNumber() int64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *DogecoinTransaction) GetBlockTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.BlockTimestamp
	}
	return nil
}

func (x *DogecoinTransaction) GetBlockTimestampMonth() *date.Date {
	if x != nil {
		return x.BlockTimestampMonth
	}
	return nil
}

func (x *DogecoinTransaction) GetInputCount() int64 {
	if x != nil {
		return x.InputCount
	}
	return 0
}

func (x *DogecoinTransaction) GetOutputCount() int64 {
	if x != nil {
		return x.OutputCount
	}
	return 0
}

func (x *DogecoinTransaction) GetIsCoinbase() bool {
	if x != nil {
		return x.IsCoinbase
	}
	return false
}

func (x *DogecoinTransaction) GetInputs() []*DogecoinTransaction_Input {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *DogecoinTransaction) GetOutputs() []*DogecoinTransaction_Output {
	if x != nil {
		return x.Outputs
	}
	return nil
}

type DogecoinTransaction_Input struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	Index                int64                  `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`                                                            // INTEGER REQUIRED
	SpentTransactionHash string                 `protobuf:"bytes,2,opt,name=spent_transaction_hash,json=spentTransactionHash,proto3" json:"spent_transaction_hash,omitempty"` // STRING NULLABLE
	SpentOutputIndex     int64                  `protobuf:"varint,3,opt,name=spent_output_index,json=spentOutputIndex,proto3" json:"spent_output_index,omitempty"`            // INTEGER NULLABLE
	ScriptAsm            string                 `protobuf:"bytes,4,opt,name=script_asm,json=scriptAsm,proto3" json:"script_asm,omitempty"`                                    // STRING NULLABLE
	ScriptHex            string                 `protobuf:"bytes,5,opt,name=script_hex,json=scriptHex,proto3" json:"script_hex,omitempty"`                                    // STRING NULLABLE
	Sequence             int64                  `protobuf:"varint,6,opt,name=sequence,proto3" json:"sequence,omitempty"`                                                      // INTEGER NULLABLE
	RequiredSignatures   int64                  `protobuf:"varint,7,opt,name=required_signatures,json=requiredSignatures,proto3" json:"required_signatures,omitempty"`        // INTEGER NULLABLE
	Type                 string                 `protobuf:"bytes,8,opt,name=type,proto3" json:"type,omitempty"`                                                               // STRING NULLABLE
	Addresses            []string               `protobuf:"bytes,9,rep,name=addresses,proto3" json:"addresses,omitempty"`                                                     // STRING REPEATED
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *DogecoinTransaction_Input) Reset() {
	*x = DogecoinTransaction_Input{}
	mi := &file_einride_bigquery_public_v1_dogecoin_transaction_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DogecoinTransaction_Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DogecoinTransaction_Input) ProtoMessage() {}

func (x *DogecoinTransaction_Input) ProtoReflect() protoreflect.Message {
	mi := &file_einride_bigquery_public_v1_dogecoin_transaction_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DogecoinTransaction_Input.ProtoReflect.Descriptor instead.
func (*DogecoinTransaction_Input) Descriptor() ([]byte, []int) {
	return file_einride_bigquery_public_v1_dogecoin_transaction_proto_rawDescGZIP(), []int{0, 0}
}

func (x *DogecoinTransaction_Input) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *DogecoinTransaction_Input) GetSpentTransactionHash() string {
	if x != nil {
		return x.SpentTransactionHash
	}
	return ""
}

func (x *DogecoinTransaction_Input) GetSpentOutputIndex() int64 {
	if x != nil {
		return x.SpentOutputIndex
	}
	return 0
}

func (x *DogecoinTransaction_Input) GetScriptAsm() string {
	if x != nil {
		return x.ScriptAsm
	}
	return ""
}

func (x *DogecoinTransaction_Input) GetScriptHex() string {
	if x != nil {
		return x.ScriptHex
	}
	return ""
}

func (x *DogecoinTransaction_Input) GetSequence() int64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *DogecoinTransaction_Input) GetRequiredSignatures() int64 {
	if x != nil {
		return x.RequiredSignatures
	}
	return 0
}

func (x *DogecoinTransaction_Input) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *DogecoinTransaction_Input) GetAddresses() []string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

type DogecoinTransaction_Output struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	Index              int64                  `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`                                                     // INTEGER REQUIRED
	ScriptAsm          string                 `protobuf:"bytes,2,opt,name=script_asm,json=scriptAsm,proto3" json:"script_asm,omitempty"`                             // STRING NULLABLE
	ScriptHex          string                 `protobuf:"bytes,3,opt,name=script_hex,json=scriptHex,proto3" json:"script_hex,omitempty"`                             // STRING NULLABLE
	RequiredSignatures int64                  `protobuf:"varint,4,opt,name=required_signatures,json=requiredSignatures,proto3" json:"required_signatures,omitempty"` // INTEGER NULLABLE
	Type               string                 `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`                                                        // STRING NULLABLE
	Addresses          []string               `protobuf:"bytes,6,rep,name=addresses,proto3" json:"addresses,omitempty"`                                              // STRING REPEATED
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *DogecoinTransaction_Output) Reset() {
	*x = DogecoinTransaction_Output{}
	mi := &file_einride_bigquery_public_v1_dogecoin_transaction_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DogecoinTransaction_Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DogecoinTransaction_Output) ProtoMessage() {}

func (x *DogecoinTransaction_Output) ProtoReflect() protoreflect.Message {
	mi := &file_einride_bigquery_public_v1_dogecoin_transaction_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DogecoinTransaction_Output.ProtoReflect.Descriptor instead.
func (*DogecoinTransaction_Output) Descriptor() ([]byte, []int) {
	return file_einride_bigquery_public_v1_dogecoin_transaction_proto_rawDescGZIP(), []int{0, 1}
}

func (x *DogecoinTransaction_Output) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *DogecoinTransaction_Output) GetScriptAsm() string {
	if x != nil {
		return x.ScriptAsm
	}
	return ""
}

func (x *DogecoinTransaction_Output) GetScriptHex() string {
	if x != nil {
		return x.ScriptHex
	}
	return ""
}

func (x *DogecoinTransaction_Output) GetRequiredSignatures() int64 {
	if x != nil {
		return x.RequiredSignatures
	}
	return 0
}

func (x *DogecoinTransaction_Output) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *DogecoinTransaction_Output) GetAddresses() []string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

var File_einride_bigquery_public_v1_dogecoin_transaction_proto protoreflect.FileDescriptor

var file_einride_bigquery_public_v1_dogecoin_transaction_proto_rawDesc = string([]byte{
	0x0a, 0x35, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x67,
	0x65, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65,
	0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x08, 0x0a,
	0x13, 0x44, 0x6f, 0x67, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x0f, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x45, 0x0a,
	0x15, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52,
	0x13, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x4d,
	0x6f, 0x6e, 0x74, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x63,
	0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69,
	0x73, 0x43, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x06, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x65, 0x69, 0x6e, 0x72,
	0x69, 0x64, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x67, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x52, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x12, 0x50, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x73, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x65, 0x69, 0x6e, 0x72,
	0x69, 0x64, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x67, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x1a, 0xbe, 0x02, 0x0a, 0x05, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x34, 0x0a, 0x16, 0x73, 0x70,
	0x65, 0x6e, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x73, 0x70, 0x65, 0x6e,
	0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x2c, 0x0a, 0x12, 0x73, 0x70, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x73, 0x70,
	0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x61, 0x73, 0x6d, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x41, 0x73, 0x6d, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x68, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x48, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x53,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x1a, 0xbf, 0x01, 0x0a, 0x06,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x61, 0x73, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x41, 0x73, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x68, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x48, 0x65, 0x78, 0x12, 0x2f, 0x0a, 0x13, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x42, 0xaa, 0x02,
	0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x62, 0x69,
	0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x76, 0x31,
	0x42, 0x18, 0x44, 0x6f, 0x67, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x61, 0x67, 0x6f,
	0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x69, 0x6e, 0x72,
	0x69, 0x64, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x45, 0x42, 0x50, 0xaa, 0x02, 0x1a, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e,
	0x42, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x1b, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x42, 0x69, 0x67,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x5c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x27, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x42, 0x69, 0x67, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x5c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1d, 0x45, 0x69, 0x6e,
	0x72, 0x69, 0x64, 0x65, 0x3a, 0x3a, 0x42, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x3a, 0x3a,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_einride_bigquery_public_v1_dogecoin_transaction_proto_rawDescOnce sync.Once
	file_einride_bigquery_public_v1_dogecoin_transaction_proto_rawDescData []byte
)

func file_einride_bigquery_public_v1_dogecoin_transaction_proto_rawDescGZIP() []byte {
	file_einride_bigquery_public_v1_dogecoin_transaction_proto_rawDescOnce.Do(func() {
		file_einride_bigquery_public_v1_dogecoin_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_einride_bigquery_public_v1_dogecoin_transaction_proto_rawDesc), len(file_einride_bigquery_public_v1_dogecoin_transaction_proto_rawDesc)))
	})
	return file_einride_bigquery_public_v1_dogecoin_transaction_proto_rawDescData
}

var file_einride_bigquery_public_v1_dogecoin_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_einride_bigquery_public_v1_dogecoin_transaction_proto_goTypes = []any{
	(*DogecoinTransaction)(nil),        // 0: einride.bigquery.public.v1.DogecoinTransaction
	(*DogecoinTransaction_Input)(nil),  // 1: einride.bigquery.public.v1.DogecoinTransaction.Input
	(*DogecoinTransaction_Output)(nil), // 2: einride.bigquery.public.v1.DogecoinTransaction.Output
	(*timestamppb.Timestamp)(nil),      // 3: google.protobuf.Timestamp
	(*date.Date)(nil),                  // 4: google.type.Date
}
var file_einride_bigquery_public_v1_dogecoin_transaction_proto_depIdxs = []int32{
	3, // 0: einride.bigquery.public.v1.DogecoinTransaction.block_timestamp:type_name -> google.protobuf.Timestamp
	4, // 1: einride.bigquery.public.v1.DogecoinTransaction.block_timestamp_month:type_name -> google.type.Date
	1, // 2: einride.bigquery.public.v1.DogecoinTransaction.inputs:type_name -> einride.bigquery.public.v1.DogecoinTransaction.Input
	2, // 3: einride.bigquery.public.v1.DogecoinTransaction.outputs:type_name -> einride.bigquery.public.v1.DogecoinTransaction.Output
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_einride_bigquery_public_v1_dogecoin_transaction_proto_init() }
func file_einride_bigquery_public_v1_dogecoin_transaction_proto_init() {
	if File_einride_bigquery_public_v1_dogecoin_transaction_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_einride_bigquery_public_v1_dogecoin_transaction_proto_rawDesc), len(file_einride_bigquery_public_v1_dogecoin_transaction_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_bigquery_public_v1_dogecoin_transaction_proto_goTypes,
		DependencyIndexes: file_einride_bigquery_public_v1_dogecoin_transaction_proto_depIdxs,
		MessageInfos:      file_einride_bigquery_public_v1_dogecoin_transaction_proto_msgTypes,
	}.Build()
	File_einride_bigquery_public_v1_dogecoin_transaction_proto = out.File
	file_einride_bigquery_public_v1_dogecoin_transaction_proto_goTypes = nil
	file_einride_bigquery_public_v1_dogecoin_transaction_proto_depIdxs = nil
}
