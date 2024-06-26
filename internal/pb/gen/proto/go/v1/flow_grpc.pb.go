// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: v1/flow.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FlowServiceClient is the client API for FlowService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FlowServiceClient interface {
	UpdateFlow(ctx context.Context, in *UpdateFlowRequest, opts ...grpc.CallOption) (*UpdateFlowResponse, error)
	UpdateFlowV2(ctx context.Context, in *UpdateFlowV2Request, opts ...grpc.CallOption) (*UpdateFlowV2Response, error)
	CreateFlow(ctx context.Context, in *CreateFlowRequest, opts ...grpc.CallOption) (*CreateFlowResponse, error)
	GetFlow(ctx context.Context, in *GetFlowRequest, opts ...grpc.CallOption) (*GetFlowResponse, error)
	ListFlow(ctx context.Context, in *ListFlowRequest, opts ...grpc.CallOption) (*ListFlowResponse, error)
	DeleteFlow(ctx context.Context, in *DeleteFlowRequest, opts ...grpc.CallOption) (*DeleteFlowResponse, error)
	CopyFlow(ctx context.Context, in *CopyFlowReq, opts ...grpc.CallOption) (*CopyFlowRes, error)
	ShareFlow(ctx context.Context, in *ShareFlowReq, opts ...grpc.CallOption) (*ShareFlowRes, error)
	HideFlow(ctx context.Context, in *HideFlowReq, opts ...grpc.CallOption) (*HideFlowRes, error)
	SharedFlows(ctx context.Context, in *SharedFlowsReq, opts ...grpc.CallOption) (*SharedFlowsRes, error)
	SharedFlow(ctx context.Context, in *SharedFlowReq, opts ...grpc.CallOption) (*SharedFlowRes, error)
	UseSharedFlow(ctx context.Context, in *UseSharedFlowReq, opts ...grpc.CallOption) (*UseSharedFlowRes, error)
	CreateFlowV2(ctx context.Context, in *CreateFlowV2Req, opts ...grpc.CallOption) (*CreateFlowV2Res, error)
	FlowPreview(ctx context.Context, in *FlowPreviewReq, opts ...grpc.CallOption) (*FlowPreviewRes, error)
	OnlyRandomFlowPreview(ctx context.Context, in *OnlyRandomFlowPreviewReq, opts ...grpc.CallOption) (*OnlyRandomFlowPreviewRes, error)
	OnlyRandomFlowFromTokens(ctx context.Context, in *OnlyRandomFlowPreviewReq, opts ...grpc.CallOption) (*OnlyRandomFlowFromTokensRes, error)
	GetFlowV2(ctx context.Context, in *GetFlowV2Req, opts ...grpc.CallOption) (*GetFlowV2Res, error)
}

type flowServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFlowServiceClient(cc grpc.ClientConnInterface) FlowServiceClient {
	return &flowServiceClient{cc}
}

func (c *flowServiceClient) UpdateFlow(ctx context.Context, in *UpdateFlowRequest, opts ...grpc.CallOption) (*UpdateFlowResponse, error) {
	out := new(UpdateFlowResponse)
	err := c.cc.Invoke(ctx, "/flow.FlowService/UpdateFlow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowServiceClient) UpdateFlowV2(ctx context.Context, in *UpdateFlowV2Request, opts ...grpc.CallOption) (*UpdateFlowV2Response, error) {
	out := new(UpdateFlowV2Response)
	err := c.cc.Invoke(ctx, "/flow.FlowService/UpdateFlowV2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowServiceClient) CreateFlow(ctx context.Context, in *CreateFlowRequest, opts ...grpc.CallOption) (*CreateFlowResponse, error) {
	out := new(CreateFlowResponse)
	err := c.cc.Invoke(ctx, "/flow.FlowService/CreateFlow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowServiceClient) GetFlow(ctx context.Context, in *GetFlowRequest, opts ...grpc.CallOption) (*GetFlowResponse, error) {
	out := new(GetFlowResponse)
	err := c.cc.Invoke(ctx, "/flow.FlowService/GetFlow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowServiceClient) ListFlow(ctx context.Context, in *ListFlowRequest, opts ...grpc.CallOption) (*ListFlowResponse, error) {
	out := new(ListFlowResponse)
	err := c.cc.Invoke(ctx, "/flow.FlowService/ListFlow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowServiceClient) DeleteFlow(ctx context.Context, in *DeleteFlowRequest, opts ...grpc.CallOption) (*DeleteFlowResponse, error) {
	out := new(DeleteFlowResponse)
	err := c.cc.Invoke(ctx, "/flow.FlowService/DeleteFlow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowServiceClient) CopyFlow(ctx context.Context, in *CopyFlowReq, opts ...grpc.CallOption) (*CopyFlowRes, error) {
	out := new(CopyFlowRes)
	err := c.cc.Invoke(ctx, "/flow.FlowService/CopyFlow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowServiceClient) ShareFlow(ctx context.Context, in *ShareFlowReq, opts ...grpc.CallOption) (*ShareFlowRes, error) {
	out := new(ShareFlowRes)
	err := c.cc.Invoke(ctx, "/flow.FlowService/ShareFlow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowServiceClient) HideFlow(ctx context.Context, in *HideFlowReq, opts ...grpc.CallOption) (*HideFlowRes, error) {
	out := new(HideFlowRes)
	err := c.cc.Invoke(ctx, "/flow.FlowService/HideFlow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowServiceClient) SharedFlows(ctx context.Context, in *SharedFlowsReq, opts ...grpc.CallOption) (*SharedFlowsRes, error) {
	out := new(SharedFlowsRes)
	err := c.cc.Invoke(ctx, "/flow.FlowService/SharedFlows", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowServiceClient) SharedFlow(ctx context.Context, in *SharedFlowReq, opts ...grpc.CallOption) (*SharedFlowRes, error) {
	out := new(SharedFlowRes)
	err := c.cc.Invoke(ctx, "/flow.FlowService/SharedFlow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowServiceClient) UseSharedFlow(ctx context.Context, in *UseSharedFlowReq, opts ...grpc.CallOption) (*UseSharedFlowRes, error) {
	out := new(UseSharedFlowRes)
	err := c.cc.Invoke(ctx, "/flow.FlowService/UseSharedFlow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowServiceClient) CreateFlowV2(ctx context.Context, in *CreateFlowV2Req, opts ...grpc.CallOption) (*CreateFlowV2Res, error) {
	out := new(CreateFlowV2Res)
	err := c.cc.Invoke(ctx, "/flow.FlowService/CreateFlowV2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowServiceClient) FlowPreview(ctx context.Context, in *FlowPreviewReq, opts ...grpc.CallOption) (*FlowPreviewRes, error) {
	out := new(FlowPreviewRes)
	err := c.cc.Invoke(ctx, "/flow.FlowService/FlowPreview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowServiceClient) OnlyRandomFlowPreview(ctx context.Context, in *OnlyRandomFlowPreviewReq, opts ...grpc.CallOption) (*OnlyRandomFlowPreviewRes, error) {
	out := new(OnlyRandomFlowPreviewRes)
	err := c.cc.Invoke(ctx, "/flow.FlowService/OnlyRandomFlowPreview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowServiceClient) OnlyRandomFlowFromTokens(ctx context.Context, in *OnlyRandomFlowPreviewReq, opts ...grpc.CallOption) (*OnlyRandomFlowFromTokensRes, error) {
	out := new(OnlyRandomFlowFromTokensRes)
	err := c.cc.Invoke(ctx, "/flow.FlowService/OnlyRandomFlowFromTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowServiceClient) GetFlowV2(ctx context.Context, in *GetFlowV2Req, opts ...grpc.CallOption) (*GetFlowV2Res, error) {
	out := new(GetFlowV2Res)
	err := c.cc.Invoke(ctx, "/flow.FlowService/GetFlowV2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FlowServiceServer is the server API for FlowService service.
// All implementations must embed UnimplementedFlowServiceServer
// for forward compatibility
type FlowServiceServer interface {
	UpdateFlow(context.Context, *UpdateFlowRequest) (*UpdateFlowResponse, error)
	UpdateFlowV2(context.Context, *UpdateFlowV2Request) (*UpdateFlowV2Response, error)
	CreateFlow(context.Context, *CreateFlowRequest) (*CreateFlowResponse, error)
	GetFlow(context.Context, *GetFlowRequest) (*GetFlowResponse, error)
	ListFlow(context.Context, *ListFlowRequest) (*ListFlowResponse, error)
	DeleteFlow(context.Context, *DeleteFlowRequest) (*DeleteFlowResponse, error)
	CopyFlow(context.Context, *CopyFlowReq) (*CopyFlowRes, error)
	ShareFlow(context.Context, *ShareFlowReq) (*ShareFlowRes, error)
	HideFlow(context.Context, *HideFlowReq) (*HideFlowRes, error)
	SharedFlows(context.Context, *SharedFlowsReq) (*SharedFlowsRes, error)
	SharedFlow(context.Context, *SharedFlowReq) (*SharedFlowRes, error)
	UseSharedFlow(context.Context, *UseSharedFlowReq) (*UseSharedFlowRes, error)
	CreateFlowV2(context.Context, *CreateFlowV2Req) (*CreateFlowV2Res, error)
	FlowPreview(context.Context, *FlowPreviewReq) (*FlowPreviewRes, error)
	OnlyRandomFlowPreview(context.Context, *OnlyRandomFlowPreviewReq) (*OnlyRandomFlowPreviewRes, error)
	OnlyRandomFlowFromTokens(context.Context, *OnlyRandomFlowPreviewReq) (*OnlyRandomFlowFromTokensRes, error)
	GetFlowV2(context.Context, *GetFlowV2Req) (*GetFlowV2Res, error)
	mustEmbedUnimplementedFlowServiceServer()
}

// UnimplementedFlowServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFlowServiceServer struct {
}

func (UnimplementedFlowServiceServer) UpdateFlow(context.Context, *UpdateFlowRequest) (*UpdateFlowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFlow not implemented")
}
func (UnimplementedFlowServiceServer) UpdateFlowV2(context.Context, *UpdateFlowV2Request) (*UpdateFlowV2Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFlowV2 not implemented")
}
func (UnimplementedFlowServiceServer) CreateFlow(context.Context, *CreateFlowRequest) (*CreateFlowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFlow not implemented")
}
func (UnimplementedFlowServiceServer) GetFlow(context.Context, *GetFlowRequest) (*GetFlowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFlow not implemented")
}
func (UnimplementedFlowServiceServer) ListFlow(context.Context, *ListFlowRequest) (*ListFlowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFlow not implemented")
}
func (UnimplementedFlowServiceServer) DeleteFlow(context.Context, *DeleteFlowRequest) (*DeleteFlowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFlow not implemented")
}
func (UnimplementedFlowServiceServer) CopyFlow(context.Context, *CopyFlowReq) (*CopyFlowRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CopyFlow not implemented")
}
func (UnimplementedFlowServiceServer) ShareFlow(context.Context, *ShareFlowReq) (*ShareFlowRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShareFlow not implemented")
}
func (UnimplementedFlowServiceServer) HideFlow(context.Context, *HideFlowReq) (*HideFlowRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HideFlow not implemented")
}
func (UnimplementedFlowServiceServer) SharedFlows(context.Context, *SharedFlowsReq) (*SharedFlowsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SharedFlows not implemented")
}
func (UnimplementedFlowServiceServer) SharedFlow(context.Context, *SharedFlowReq) (*SharedFlowRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SharedFlow not implemented")
}
func (UnimplementedFlowServiceServer) UseSharedFlow(context.Context, *UseSharedFlowReq) (*UseSharedFlowRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UseSharedFlow not implemented")
}
func (UnimplementedFlowServiceServer) CreateFlowV2(context.Context, *CreateFlowV2Req) (*CreateFlowV2Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFlowV2 not implemented")
}
func (UnimplementedFlowServiceServer) FlowPreview(context.Context, *FlowPreviewReq) (*FlowPreviewRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FlowPreview not implemented")
}
func (UnimplementedFlowServiceServer) OnlyRandomFlowPreview(context.Context, *OnlyRandomFlowPreviewReq) (*OnlyRandomFlowPreviewRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnlyRandomFlowPreview not implemented")
}
func (UnimplementedFlowServiceServer) OnlyRandomFlowFromTokens(context.Context, *OnlyRandomFlowPreviewReq) (*OnlyRandomFlowFromTokensRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnlyRandomFlowFromTokens not implemented")
}
func (UnimplementedFlowServiceServer) GetFlowV2(context.Context, *GetFlowV2Req) (*GetFlowV2Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFlowV2 not implemented")
}
func (UnimplementedFlowServiceServer) mustEmbedUnimplementedFlowServiceServer() {}

// UnsafeFlowServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FlowServiceServer will
// result in compilation errors.
type UnsafeFlowServiceServer interface {
	mustEmbedUnimplementedFlowServiceServer()
}

func RegisterFlowServiceServer(s grpc.ServiceRegistrar, srv FlowServiceServer) {
	s.RegisterService(&FlowService_ServiceDesc, srv)
}

func _FlowService_UpdateFlow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFlowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowServiceServer).UpdateFlow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.FlowService/UpdateFlow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowServiceServer).UpdateFlow(ctx, req.(*UpdateFlowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowService_UpdateFlowV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFlowV2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowServiceServer).UpdateFlowV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.FlowService/UpdateFlowV2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowServiceServer).UpdateFlowV2(ctx, req.(*UpdateFlowV2Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowService_CreateFlow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFlowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowServiceServer).CreateFlow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.FlowService/CreateFlow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowServiceServer).CreateFlow(ctx, req.(*CreateFlowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowService_GetFlow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFlowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowServiceServer).GetFlow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.FlowService/GetFlow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowServiceServer).GetFlow(ctx, req.(*GetFlowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowService_ListFlow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFlowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowServiceServer).ListFlow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.FlowService/ListFlow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowServiceServer).ListFlow(ctx, req.(*ListFlowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowService_DeleteFlow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFlowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowServiceServer).DeleteFlow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.FlowService/DeleteFlow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowServiceServer).DeleteFlow(ctx, req.(*DeleteFlowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowService_CopyFlow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CopyFlowReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowServiceServer).CopyFlow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.FlowService/CopyFlow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowServiceServer).CopyFlow(ctx, req.(*CopyFlowReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowService_ShareFlow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShareFlowReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowServiceServer).ShareFlow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.FlowService/ShareFlow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowServiceServer).ShareFlow(ctx, req.(*ShareFlowReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowService_HideFlow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HideFlowReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowServiceServer).HideFlow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.FlowService/HideFlow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowServiceServer).HideFlow(ctx, req.(*HideFlowReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowService_SharedFlows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SharedFlowsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowServiceServer).SharedFlows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.FlowService/SharedFlows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowServiceServer).SharedFlows(ctx, req.(*SharedFlowsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowService_SharedFlow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SharedFlowReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowServiceServer).SharedFlow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.FlowService/SharedFlow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowServiceServer).SharedFlow(ctx, req.(*SharedFlowReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowService_UseSharedFlow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UseSharedFlowReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowServiceServer).UseSharedFlow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.FlowService/UseSharedFlow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowServiceServer).UseSharedFlow(ctx, req.(*UseSharedFlowReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowService_CreateFlowV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFlowV2Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowServiceServer).CreateFlowV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.FlowService/CreateFlowV2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowServiceServer).CreateFlowV2(ctx, req.(*CreateFlowV2Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowService_FlowPreview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlowPreviewReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowServiceServer).FlowPreview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.FlowService/FlowPreview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowServiceServer).FlowPreview(ctx, req.(*FlowPreviewReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowService_OnlyRandomFlowPreview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnlyRandomFlowPreviewReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowServiceServer).OnlyRandomFlowPreview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.FlowService/OnlyRandomFlowPreview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowServiceServer).OnlyRandomFlowPreview(ctx, req.(*OnlyRandomFlowPreviewReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowService_OnlyRandomFlowFromTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnlyRandomFlowPreviewReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowServiceServer).OnlyRandomFlowFromTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.FlowService/OnlyRandomFlowFromTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowServiceServer).OnlyRandomFlowFromTokens(ctx, req.(*OnlyRandomFlowPreviewReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowService_GetFlowV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFlowV2Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowServiceServer).GetFlowV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.FlowService/GetFlowV2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowServiceServer).GetFlowV2(ctx, req.(*GetFlowV2Req))
	}
	return interceptor(ctx, in, info, handler)
}

// FlowService_ServiceDesc is the grpc.ServiceDesc for FlowService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FlowService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flow.FlowService",
	HandlerType: (*FlowServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateFlow",
			Handler:    _FlowService_UpdateFlow_Handler,
		},
		{
			MethodName: "UpdateFlowV2",
			Handler:    _FlowService_UpdateFlowV2_Handler,
		},
		{
			MethodName: "CreateFlow",
			Handler:    _FlowService_CreateFlow_Handler,
		},
		{
			MethodName: "GetFlow",
			Handler:    _FlowService_GetFlow_Handler,
		},
		{
			MethodName: "ListFlow",
			Handler:    _FlowService_ListFlow_Handler,
		},
		{
			MethodName: "DeleteFlow",
			Handler:    _FlowService_DeleteFlow_Handler,
		},
		{
			MethodName: "CopyFlow",
			Handler:    _FlowService_CopyFlow_Handler,
		},
		{
			MethodName: "ShareFlow",
			Handler:    _FlowService_ShareFlow_Handler,
		},
		{
			MethodName: "HideFlow",
			Handler:    _FlowService_HideFlow_Handler,
		},
		{
			MethodName: "SharedFlows",
			Handler:    _FlowService_SharedFlows_Handler,
		},
		{
			MethodName: "SharedFlow",
			Handler:    _FlowService_SharedFlow_Handler,
		},
		{
			MethodName: "UseSharedFlow",
			Handler:    _FlowService_UseSharedFlow_Handler,
		},
		{
			MethodName: "CreateFlowV2",
			Handler:    _FlowService_CreateFlowV2_Handler,
		},
		{
			MethodName: "FlowPreview",
			Handler:    _FlowService_FlowPreview_Handler,
		},
		{
			MethodName: "OnlyRandomFlowPreview",
			Handler:    _FlowService_OnlyRandomFlowPreview_Handler,
		},
		{
			MethodName: "OnlyRandomFlowFromTokens",
			Handler:    _FlowService_OnlyRandomFlowFromTokens_Handler,
		},
		{
			MethodName: "GetFlowV2",
			Handler:    _FlowService_GetFlowV2_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/flow.proto",
}
