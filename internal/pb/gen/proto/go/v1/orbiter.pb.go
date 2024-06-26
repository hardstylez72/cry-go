// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: v1/orbiter.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetSwapOptionsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromNetwork Network `protobuf:"varint,1,opt,name=from_network,json=fromNetwork,proto3,enum=shared.Network" json:"from_network,omitempty"`
	ToNetwork   Network `protobuf:"varint,2,opt,name=to_network,json=toNetwork,proto3,enum=shared.Network" json:"to_network,omitempty"`
	FromToken   Token   `protobuf:"varint,3,opt,name=from_token,json=fromToken,proto3,enum=shared.Token" json:"from_token,omitempty"`
	ToToken     Token   `protobuf:"varint,4,opt,name=to_token,json=toToken,proto3,enum=shared.Token" json:"to_token,omitempty"`
}

func (x *GetSwapOptionsReq) Reset() {
	*x = GetSwapOptionsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_orbiter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSwapOptionsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSwapOptionsReq) ProtoMessage() {}

func (x *GetSwapOptionsReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_orbiter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSwapOptionsReq.ProtoReflect.Descriptor instead.
func (*GetSwapOptionsReq) Descriptor() ([]byte, []int) {
	return file_v1_orbiter_proto_rawDescGZIP(), []int{0}
}

func (x *GetSwapOptionsReq) GetFromNetwork() Network {
	if x != nil {
		return x.FromNetwork
	}
	return Network_ARBITRUM
}

func (x *GetSwapOptionsReq) GetToNetwork() Network {
	if x != nil {
		return x.ToNetwork
	}
	return Network_ARBITRUM
}

func (x *GetSwapOptionsReq) GetFromToken() Token {
	if x != nil {
		return x.FromToken
	}
	return Token_USDT
}

func (x *GetSwapOptionsReq) GetToToken() Token {
	if x != nil {
		return x.ToToken
	}
	return Token_USDT
}

type GetSwapOptionsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Min   string  `protobuf:"bytes,1,opt,name=min,proto3" json:"min,omitempty"`
	Max   string  `protobuf:"bytes,2,opt,name=max,proto3" json:"max,omitempty"`
	Fee   string  `protobuf:"bytes,3,opt,name=fee,proto3" json:"fee,omitempty"`
	Error *string `protobuf:"bytes,4,opt,name=error,proto3,oneof" json:"error,omitempty"`
}

func (x *GetSwapOptionsRes) Reset() {
	*x = GetSwapOptionsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_orbiter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSwapOptionsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSwapOptionsRes) ProtoMessage() {}

func (x *GetSwapOptionsRes) ProtoReflect() protoreflect.Message {
	mi := &file_v1_orbiter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSwapOptionsRes.ProtoReflect.Descriptor instead.
func (*GetSwapOptionsRes) Descriptor() ([]byte, []int) {
	return file_v1_orbiter_proto_rawDescGZIP(), []int{1}
}

func (x *GetSwapOptionsRes) GetMin() string {
	if x != nil {
		return x.Min
	}
	return ""
}

func (x *GetSwapOptionsRes) GetMax() string {
	if x != nil {
		return x.Max
	}
	return ""
}

func (x *GetSwapOptionsRes) GetFee() string {
	if x != nil {
		return x.Fee
	}
	return ""
}

func (x *GetSwapOptionsRes) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ""
}

type GetToTokensReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromNetwork Network `protobuf:"varint,1,opt,name=from_network,json=fromNetwork,proto3,enum=shared.Network" json:"from_network,omitempty"`
	ToNetwork   Network `protobuf:"varint,2,opt,name=to_network,json=toNetwork,proto3,enum=shared.Network" json:"to_network,omitempty"`
	FromToken   Token   `protobuf:"varint,3,opt,name=from_token,json=fromToken,proto3,enum=shared.Token" json:"from_token,omitempty"`
}

func (x *GetToTokensReq) Reset() {
	*x = GetToTokensReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_orbiter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetToTokensReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetToTokensReq) ProtoMessage() {}

func (x *GetToTokensReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_orbiter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetToTokensReq.ProtoReflect.Descriptor instead.
func (*GetToTokensReq) Descriptor() ([]byte, []int) {
	return file_v1_orbiter_proto_rawDescGZIP(), []int{2}
}

func (x *GetToTokensReq) GetFromNetwork() Network {
	if x != nil {
		return x.FromNetwork
	}
	return Network_ARBITRUM
}

func (x *GetToTokensReq) GetToNetwork() Network {
	if x != nil {
		return x.ToNetwork
	}
	return Network_ARBITRUM
}

func (x *GetToTokensReq) GetFromToken() Token {
	if x != nil {
		return x.FromToken
	}
	return Token_USDT
}

type GetToTokensRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tokens []Token `protobuf:"varint,1,rep,packed,name=tokens,proto3,enum=shared.Token" json:"tokens,omitempty"`
}

func (x *GetToTokensRes) Reset() {
	*x = GetToTokensRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_orbiter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetToTokensRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetToTokensRes) ProtoMessage() {}

func (x *GetToTokensRes) ProtoReflect() protoreflect.Message {
	mi := &file_v1_orbiter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetToTokensRes.ProtoReflect.Descriptor instead.
func (*GetToTokensRes) Descriptor() ([]byte, []int) {
	return file_v1_orbiter_proto_rawDescGZIP(), []int{3}
}

func (x *GetToTokensRes) GetTokens() []Token {
	if x != nil {
		return x.Tokens
	}
	return nil
}

type GetFromTokensReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromNetwork Network `protobuf:"varint,1,opt,name=from_network,json=fromNetwork,proto3,enum=shared.Network" json:"from_network,omitempty"`
	ToNetwork   Network `protobuf:"varint,2,opt,name=to_network,json=toNetwork,proto3,enum=shared.Network" json:"to_network,omitempty"`
}

func (x *GetFromTokensReq) Reset() {
	*x = GetFromTokensReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_orbiter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFromTokensReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFromTokensReq) ProtoMessage() {}

func (x *GetFromTokensReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_orbiter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFromTokensReq.ProtoReflect.Descriptor instead.
func (*GetFromTokensReq) Descriptor() ([]byte, []int) {
	return file_v1_orbiter_proto_rawDescGZIP(), []int{4}
}

func (x *GetFromTokensReq) GetFromNetwork() Network {
	if x != nil {
		return x.FromNetwork
	}
	return Network_ARBITRUM
}

func (x *GetFromTokensReq) GetToNetwork() Network {
	if x != nil {
		return x.ToNetwork
	}
	return Network_ARBITRUM
}

type GetFromTokensRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tokens []Token `protobuf:"varint,1,rep,packed,name=tokens,proto3,enum=shared.Token" json:"tokens,omitempty"`
}

func (x *GetFromTokensRes) Reset() {
	*x = GetFromTokensRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_orbiter_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFromTokensRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFromTokensRes) ProtoMessage() {}

func (x *GetFromTokensRes) ProtoReflect() protoreflect.Message {
	mi := &file_v1_orbiter_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFromTokensRes.ProtoReflect.Descriptor instead.
func (*GetFromTokensRes) Descriptor() ([]byte, []int) {
	return file_v1_orbiter_proto_rawDescGZIP(), []int{5}
}

func (x *GetFromTokensRes) GetTokens() []Token {
	if x != nil {
		return x.Tokens
	}
	return nil
}

type GetNetworksReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetNetworksReq) Reset() {
	*x = GetNetworksReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_orbiter_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNetworksReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNetworksReq) ProtoMessage() {}

func (x *GetNetworksReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_orbiter_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNetworksReq.ProtoReflect.Descriptor instead.
func (*GetNetworksReq) Descriptor() ([]byte, []int) {
	return file_v1_orbiter_proto_rawDescGZIP(), []int{6}
}

type GetNetworksRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Networks []Network `protobuf:"varint,1,rep,packed,name=networks,proto3,enum=shared.Network" json:"networks,omitempty"`
}

func (x *GetNetworksRes) Reset() {
	*x = GetNetworksRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_orbiter_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNetworksRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNetworksRes) ProtoMessage() {}

func (x *GetNetworksRes) ProtoReflect() protoreflect.Message {
	mi := &file_v1_orbiter_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNetworksRes.ProtoReflect.Descriptor instead.
func (*GetNetworksRes) Descriptor() ([]byte, []int) {
	return file_v1_orbiter_proto_rawDescGZIP(), []int{7}
}

func (x *GetNetworksRes) GetNetworks() []Network {
	if x != nil {
		return x.Networks
	}
	return nil
}

var File_v1_orbiter_proto protoreflect.FileDescriptor

var file_v1_orbiter_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x62, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x6f, 0x72, 0x62, 0x69, 0x74, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x76, 0x31, 0x2f, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x02, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x53, 0x77, 0x61, 0x70, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x12, 0x32, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x2e, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x09, 0x74, 0x6f, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x2c, 0x0a, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x28, 0x0a, 0x08, 0x74, 0x6f, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x07, 0x74, 0x6f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x3a, 0x39, 0x92,
	0x41, 0x36, 0x0a, 0x34, 0xd2, 0x01, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0xd2, 0x01, 0x0a, 0x74, 0x6f, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0xd2, 0x01, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0xd2, 0x01, 0x08,
	0x74, 0x6f, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x87, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x53, 0x77, 0x61, 0x70, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x69, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x61, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x66, 0x65, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x3a,
	0x17, 0x92, 0x41, 0x14, 0x0a, 0x12, 0xd2, 0x01, 0x03, 0x6d, 0x69, 0x6e, 0xd2, 0x01, 0x03, 0x6d,
	0x61, 0x78, 0xd2, 0x01, 0x03, 0x66, 0x65, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0xd2, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x12, 0x32, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x0b, 0x66, 0x72,
	0x6f, 0x6d, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x2e, 0x0a, 0x0a, 0x74, 0x6f, 0x5f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x09,
	0x74, 0x6f, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x2c, 0x0a, 0x0a, 0x66, 0x72, 0x6f,
	0x6d, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x09, 0x66, 0x72,
	0x6f, 0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x3a, 0x2e, 0x92, 0x41, 0x2b, 0x0a, 0x29, 0xd2, 0x01,
	0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0xd2, 0x01, 0x0a,
	0x74, 0x6f, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0xd2, 0x01, 0x0a, 0x66, 0x72, 0x6f,
	0x6d, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x47, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x6f,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x06, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x3a, 0x0e, 0x92, 0x41, 0x0b, 0x0a, 0x09, 0xd2, 0x01, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x22, 0x99, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x12, 0x32, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x0b, 0x66, 0x72,
	0x6f, 0x6d, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x2e, 0x0a, 0x0a, 0x74, 0x6f, 0x5f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x09,
	0x74, 0x6f, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x3a, 0x21, 0x92, 0x41, 0x1e, 0x0a, 0x1c,
	0xd2, 0x01, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0xd2,
	0x01, 0x0a, 0x74, 0x6f, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x22, 0x49, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x12, 0x25, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e,
	0x32, 0x0d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x3a, 0x0e, 0x92, 0x41, 0x0b, 0x0a, 0x09, 0xd2, 0x01,
	0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x22, 0x4f, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x08, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0f, 0x2e,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x08,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x3a, 0x10, 0x92, 0x41, 0x0d, 0x0a, 0x0b, 0xd2,
	0x01, 0x08, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x32, 0xcb, 0x03, 0x0a, 0x0e, 0x4f,
	0x72, 0x62, 0x69, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x70, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x19,
	0x2e, 0x6f, 0x72, 0x62, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x72, 0x6f, 0x6d,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x6f, 0x72, 0x62, 0x69,
	0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x3a, 0x01, 0x2a, 0x22,
	0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x62, 0x69,
	0x74, 0x65, 0x72, 0x2f, 0x66, 0x72, 0x6f, 0x6d, 0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12,
	0x68, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x17,
	0x2e, 0x6f, 0x72, 0x62, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x6f, 0x72, 0x62, 0x69, 0x74, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x3a, 0x01, 0x2a, 0x22, 0x1c, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x62, 0x69, 0x74, 0x65, 0x72, 0x2f,
	0x74, 0x6f, 0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x74, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x53, 0x77, 0x61, 0x70, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x2e, 0x6f, 0x72,
	0x62, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x77, 0x61, 0x70, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x72, 0x62, 0x69, 0x74, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x77, 0x61, 0x70, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x3a, 0x01, 0x2a, 0x22, 0x1f,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x62, 0x69, 0x74,
	0x65, 0x72, 0x2f, 0x73, 0x77, 0x61, 0x70, 0x2d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x67, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x12, 0x17,
	0x2e, 0x6f, 0x72, 0x62, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x6f, 0x72, 0x62, 0x69, 0x74, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x73,
	0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x62, 0x69, 0x74, 0x65, 0x72, 0x2f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x42, 0x09, 0x5a, 0x07, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_orbiter_proto_rawDescOnce sync.Once
	file_v1_orbiter_proto_rawDescData = file_v1_orbiter_proto_rawDesc
)

func file_v1_orbiter_proto_rawDescGZIP() []byte {
	file_v1_orbiter_proto_rawDescOnce.Do(func() {
		file_v1_orbiter_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_orbiter_proto_rawDescData)
	})
	return file_v1_orbiter_proto_rawDescData
}

var file_v1_orbiter_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_v1_orbiter_proto_goTypes = []interface{}{
	(*GetSwapOptionsReq)(nil), // 0: orbiter.GetSwapOptionsReq
	(*GetSwapOptionsRes)(nil), // 1: orbiter.GetSwapOptionsRes
	(*GetToTokensReq)(nil),    // 2: orbiter.GetToTokensReq
	(*GetToTokensRes)(nil),    // 3: orbiter.GetToTokensRes
	(*GetFromTokensReq)(nil),  // 4: orbiter.GetFromTokensReq
	(*GetFromTokensRes)(nil),  // 5: orbiter.GetFromTokensRes
	(*GetNetworksReq)(nil),    // 6: orbiter.GetNetworksReq
	(*GetNetworksRes)(nil),    // 7: orbiter.GetNetworksRes
	(Network)(0),              // 8: shared.Network
	(Token)(0),                // 9: shared.Token
}
var file_v1_orbiter_proto_depIdxs = []int32{
	8,  // 0: orbiter.GetSwapOptionsReq.from_network:type_name -> shared.Network
	8,  // 1: orbiter.GetSwapOptionsReq.to_network:type_name -> shared.Network
	9,  // 2: orbiter.GetSwapOptionsReq.from_token:type_name -> shared.Token
	9,  // 3: orbiter.GetSwapOptionsReq.to_token:type_name -> shared.Token
	8,  // 4: orbiter.GetToTokensReq.from_network:type_name -> shared.Network
	8,  // 5: orbiter.GetToTokensReq.to_network:type_name -> shared.Network
	9,  // 6: orbiter.GetToTokensReq.from_token:type_name -> shared.Token
	9,  // 7: orbiter.GetToTokensRes.tokens:type_name -> shared.Token
	8,  // 8: orbiter.GetFromTokensReq.from_network:type_name -> shared.Network
	8,  // 9: orbiter.GetFromTokensReq.to_network:type_name -> shared.Network
	9,  // 10: orbiter.GetFromTokensRes.tokens:type_name -> shared.Token
	8,  // 11: orbiter.GetNetworksRes.networks:type_name -> shared.Network
	4,  // 12: orbiter.OrbiterService.GetFromTokens:input_type -> orbiter.GetFromTokensReq
	2,  // 13: orbiter.OrbiterService.GetToTokens:input_type -> orbiter.GetToTokensReq
	0,  // 14: orbiter.OrbiterService.GetSwapOptions:input_type -> orbiter.GetSwapOptionsReq
	6,  // 15: orbiter.OrbiterService.GetNetworks:input_type -> orbiter.GetNetworksReq
	5,  // 16: orbiter.OrbiterService.GetFromTokens:output_type -> orbiter.GetFromTokensRes
	3,  // 17: orbiter.OrbiterService.GetToTokens:output_type -> orbiter.GetToTokensRes
	1,  // 18: orbiter.OrbiterService.GetSwapOptions:output_type -> orbiter.GetSwapOptionsRes
	7,  // 19: orbiter.OrbiterService.GetNetworks:output_type -> orbiter.GetNetworksRes
	16, // [16:20] is the sub-list for method output_type
	12, // [12:16] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_v1_orbiter_proto_init() }
func file_v1_orbiter_proto_init() {
	if File_v1_orbiter_proto != nil {
		return
	}
	file_v1_shared_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_orbiter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSwapOptionsReq); i {
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
		file_v1_orbiter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSwapOptionsRes); i {
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
		file_v1_orbiter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetToTokensReq); i {
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
		file_v1_orbiter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetToTokensRes); i {
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
		file_v1_orbiter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFromTokensReq); i {
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
		file_v1_orbiter_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFromTokensRes); i {
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
		file_v1_orbiter_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNetworksReq); i {
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
		file_v1_orbiter_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNetworksRes); i {
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
	file_v1_orbiter_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_orbiter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_orbiter_proto_goTypes,
		DependencyIndexes: file_v1_orbiter_proto_depIdxs,
		MessageInfos:      file_v1_orbiter_proto_msgTypes,
	}.Build()
	File_v1_orbiter_proto = out.File
	file_v1_orbiter_proto_rawDesc = nil
	file_v1_orbiter_proto_goTypes = nil
	file_v1_orbiter_proto_depIdxs = nil
}
