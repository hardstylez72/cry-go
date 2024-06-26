// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: v1/public.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SwapStatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SwapStatReq) Reset() {
	*x = SwapStatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_public_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwapStatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwapStatReq) ProtoMessage() {}

func (x *SwapStatReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_public_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwapStatReq.ProtoReflect.Descriptor instead.
func (*SwapStatReq) Descriptor() ([]byte, []int) {
	return file_v1_public_proto_rawDescGZIP(), []int{0}
}

type SwapStatRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ZkSyncETHUSDC   []*SwapStatItem        `protobuf:"bytes,1,rep,name=zkSyncETHUSDC,proto3" json:"zkSyncETHUSDC,omitempty"`
	ZkSyncUSDCETH   []*SwapStatItem        `protobuf:"bytes,2,rep,name=zkSyncUSDCETH,proto3" json:"zkSyncUSDCETH,omitempty"`
	StarknetETHUSDC []*SwapStatItem        `protobuf:"bytes,3,rep,name=starknetETHUSDC,proto3" json:"starknetETHUSDC,omitempty"`
	StarknetUSDCETH []*SwapStatItem        `protobuf:"bytes,4,rep,name=starknetUSDCETH,proto3" json:"starknetUSDCETH,omitempty"`
	Updated         *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *SwapStatRes) Reset() {
	*x = SwapStatRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_public_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwapStatRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwapStatRes) ProtoMessage() {}

func (x *SwapStatRes) ProtoReflect() protoreflect.Message {
	mi := &file_v1_public_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwapStatRes.ProtoReflect.Descriptor instead.
func (*SwapStatRes) Descriptor() ([]byte, []int) {
	return file_v1_public_proto_rawDescGZIP(), []int{1}
}

func (x *SwapStatRes) GetZkSyncETHUSDC() []*SwapStatItem {
	if x != nil {
		return x.ZkSyncETHUSDC
	}
	return nil
}

func (x *SwapStatRes) GetZkSyncUSDCETH() []*SwapStatItem {
	if x != nil {
		return x.ZkSyncUSDCETH
	}
	return nil
}

func (x *SwapStatRes) GetStarknetETHUSDC() []*SwapStatItem {
	if x != nil {
		return x.StarknetETHUSDC
	}
	return nil
}

func (x *SwapStatRes) GetStarknetUSDCETH() []*SwapStatItem {
	if x != nil {
		return x.StarknetUSDCETH
	}
	return nil
}

func (x *SwapStatRes) GetUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

type SwapStatItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      TaskType `protobuf:"varint,1,opt,name=type,proto3,enum=task.TaskType" json:"type,omitempty"`
	RateRatio float64  `protobuf:"fixed64,2,opt,name=rateRatio,proto3" json:"rateRatio,omitempty"`
	Slippage  float64  `protobuf:"fixed64,3,opt,name=slippage,proto3" json:"slippage,omitempty"`
	Sum       float64  `protobuf:"fixed64,4,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *SwapStatItem) Reset() {
	*x = SwapStatItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_public_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwapStatItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwapStatItem) ProtoMessage() {}

func (x *SwapStatItem) ProtoReflect() protoreflect.Message {
	mi := &file_v1_public_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwapStatItem.ProtoReflect.Descriptor instead.
func (*SwapStatItem) Descriptor() ([]byte, []int) {
	return file_v1_public_proto_rawDescGZIP(), []int{2}
}

func (x *SwapStatItem) GetType() TaskType {
	if x != nil {
		return x.Type
	}
	return TaskType_StargateBridge
}

func (x *SwapStatItem) GetRateRatio() float64 {
	if x != nil {
		return x.RateRatio
	}
	return 0
}

func (x *SwapStatItem) GetSlippage() float64 {
	if x != nil {
		return x.Slippage
	}
	return 0
}

func (x *SwapStatItem) GetSum() float64 {
	if x != nil {
		return x.Sum
	}
	return 0
}

var File_v1_public_proto protoreflect.FileDescriptor

var file_v1_public_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x76, 0x31, 0x2f, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0d, 0x0a, 0x0b, 0x53, 0x77, 0x61, 0x70,
	0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x71, 0x22, 0xe0, 0x02, 0x0a, 0x0b, 0x53, 0x77, 0x61, 0x70,
	0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x0d, 0x7a, 0x6b, 0x53, 0x79, 0x6e,
	0x63, 0x45, 0x54, 0x48, 0x55, 0x53, 0x44, 0x43, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x53, 0x77, 0x61, 0x70, 0x53, 0x74, 0x61, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x0d, 0x7a, 0x6b, 0x53, 0x79, 0x6e, 0x63, 0x45, 0x54, 0x48, 0x55,
	0x53, 0x44, 0x43, 0x12, 0x3a, 0x0a, 0x0d, 0x7a, 0x6b, 0x53, 0x79, 0x6e, 0x63, 0x55, 0x53, 0x44,
	0x43, 0x45, 0x54, 0x48, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2e, 0x53, 0x77, 0x61, 0x70, 0x53, 0x74, 0x61, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x0d, 0x7a, 0x6b, 0x53, 0x79, 0x6e, 0x63, 0x55, 0x53, 0x44, 0x43, 0x45, 0x54, 0x48, 0x12,
	0x3e, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x6e, 0x65, 0x74, 0x45, 0x54, 0x48, 0x55, 0x53,
	0x44, 0x43, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x2e, 0x53, 0x77, 0x61, 0x70, 0x53, 0x74, 0x61, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0f,
	0x73, 0x74, 0x61, 0x72, 0x6b, 0x6e, 0x65, 0x74, 0x45, 0x54, 0x48, 0x55, 0x53, 0x44, 0x43, 0x12,
	0x3e, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x6e, 0x65, 0x74, 0x55, 0x53, 0x44, 0x43, 0x45,
	0x54, 0x48, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x2e, 0x53, 0x77, 0x61, 0x70, 0x53, 0x74, 0x61, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0f,
	0x73, 0x74, 0x61, 0x72, 0x6b, 0x6e, 0x65, 0x74, 0x55, 0x53, 0x44, 0x43, 0x45, 0x54, 0x48, 0x12,
	0x34, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x3a, 0x23, 0x92, 0x41, 0x20, 0x0a, 0x1e, 0xd2, 0x01, 0x06, 0x7a,
	0x6b, 0x53, 0x79, 0x6e, 0x63, 0xd2, 0x01, 0x08, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x6e, 0x65, 0x74,
	0xd2, 0x01, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0xa9, 0x01, 0x0a, 0x0c, 0x53,
	0x77, 0x61, 0x70, 0x53, 0x74, 0x61, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x22, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x72, 0x61, 0x74, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x09, 0x72, 0x61, 0x74, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x6c, 0x69, 0x70, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x08, 0x73, 0x6c, 0x69, 0x70, 0x70, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x6d,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x3a, 0x29, 0x92, 0x41, 0x26,
	0x0a, 0x24, 0xd2, 0x01, 0x04, 0x74, 0x79, 0x70, 0x65, 0xd2, 0x01, 0x09, 0x72, 0x61, 0x74, 0x65,
	0x52, 0x61, 0x74, 0x69, 0x6f, 0xd2, 0x01, 0x08, 0x73, 0x6c, 0x69, 0x70, 0x70, 0x61, 0x67, 0x65,
	0xd2, 0x01, 0x03, 0x73, 0x75, 0x6d, 0x32, 0x6d, 0x0a, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x08, 0x53, 0x77, 0x61, 0x70, 0x53,
	0x74, 0x61, 0x74, 0x12, 0x13, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x53, 0x77, 0x61,
	0x70, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x2e, 0x53, 0x77, 0x61, 0x70, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x73, 0x22, 0x26, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67,
	0x77, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x73, 0x77, 0x61, 0x70,
	0x2d, 0x73, 0x74, 0x61, 0x74, 0x42, 0x09, 0x5a, 0x07, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_public_proto_rawDescOnce sync.Once
	file_v1_public_proto_rawDescData = file_v1_public_proto_rawDesc
)

func file_v1_public_proto_rawDescGZIP() []byte {
	file_v1_public_proto_rawDescOnce.Do(func() {
		file_v1_public_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_public_proto_rawDescData)
	})
	return file_v1_public_proto_rawDescData
}

var file_v1_public_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v1_public_proto_goTypes = []interface{}{
	(*SwapStatReq)(nil),           // 0: public.SwapStatReq
	(*SwapStatRes)(nil),           // 1: public.SwapStatRes
	(*SwapStatItem)(nil),          // 2: public.SwapStatItem
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(TaskType)(0),                 // 4: task.TaskType
}
var file_v1_public_proto_depIdxs = []int32{
	2, // 0: public.SwapStatRes.zkSyncETHUSDC:type_name -> public.SwapStatItem
	2, // 1: public.SwapStatRes.zkSyncUSDCETH:type_name -> public.SwapStatItem
	2, // 2: public.SwapStatRes.starknetETHUSDC:type_name -> public.SwapStatItem
	2, // 3: public.SwapStatRes.starknetUSDCETH:type_name -> public.SwapStatItem
	3, // 4: public.SwapStatRes.updated:type_name -> google.protobuf.Timestamp
	4, // 5: public.SwapStatItem.type:type_name -> task.TaskType
	0, // 6: public.publicService.SwapStat:input_type -> public.SwapStatReq
	1, // 7: public.publicService.SwapStat:output_type -> public.SwapStatRes
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_v1_public_proto_init() }
func file_v1_public_proto_init() {
	if File_v1_public_proto != nil {
		return
	}
	file_v1_shared_proto_init()
	file_v1_task_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_public_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwapStatReq); i {
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
		file_v1_public_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwapStatRes); i {
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
		file_v1_public_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwapStatItem); i {
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
			RawDescriptor: file_v1_public_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_public_proto_goTypes,
		DependencyIndexes: file_v1_public_proto_depIdxs,
		MessageInfos:      file_v1_public_proto_msgTypes,
	}.Build()
	File_v1_public_proto = out.File
	file_v1_public_proto_rawDesc = nil
	file_v1_public_proto_goTypes = nil
	file_v1_public_proto_depIdxs = nil
}
