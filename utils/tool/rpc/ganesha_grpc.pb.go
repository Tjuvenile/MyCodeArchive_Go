// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: protocol/ganesha.proto

package rpc

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

const (
	Ganesha_AddCeafsConf_FullMethodName           = "/pmg.Ganesha/AddCeafsConf"
	Ganesha_RemoveCeafsConf_FullMethodName        = "/pmg.Ganesha/RemoveCeafsConf"
	Ganesha_SetExportMgr_FullMethodName           = "/pmg.Ganesha/SetExportMgr"
	Ganesha_UpdateExportMgr_FullMethodName        = "/pmg.Ganesha/UpdateExportMgr"
	Ganesha_AddExportConf_FullMethodName          = "/pmg.Ganesha/AddExportConf"
	Ganesha_RemoveExportConf_FullMethodName       = "/pmg.Ganesha/RemoveExportConf"
	Ganesha_RemoveExportMgr_FullMethodName        = "/pmg.Ganesha/RemoveExportMgr"
	Ganesha_AddPoliciesConf_FullMethodName        = "/pmg.Ganesha/AddPoliciesConf"
	Ganesha_RemovePolicyConf_FullMethodName       = "/pmg.Ganesha/RemovePolicyConf"
	Ganesha_UpdatePolicyConf_FullMethodName       = "/pmg.Ganesha/UpdatePolicyConf"
	Ganesha_ReadLocalConfig_FullMethodName        = "/pmg.Ganesha/ReadLocalConfig"
	Ganesha_SaveConfigToLocal_FullMethodName      = "/pmg.Ganesha/SaveConfigToLocal"
	Ganesha_RemoveLocalConf_FullMethodName        = "/pmg.Ganesha/RemoveLocalConf"
	Ganesha_RemoveLocalConfRollbak_FullMethodName = "/pmg.Ganesha/RemoveLocalConfRollbak"
	Ganesha_GetNfsStats_FullMethodName            = "/pmg.Ganesha/GetNfsStats"
)

// GaneshaClient is the client API for Ganesha service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GaneshaClient interface {
	// init
	AddCeafsConf(ctx context.Context, in *AddCeafsConfRequest, opts ...grpc.CallOption) (*BaseReply, error)
	RemoveCeafsConf(ctx context.Context, in *RemoveCeafsConfRequest, opts ...grpc.CallOption) (*BaseReply, error)
	// export
	SetExportMgr(ctx context.Context, in *ExportIDRequest, opts ...grpc.CallOption) (*BaseReply, error)
	UpdateExportMgr(ctx context.Context, in *ExportIDRequest, opts ...grpc.CallOption) (*BaseReply, error)
	AddExportConf(ctx context.Context, in *AddExportConfRequest, opts ...grpc.CallOption) (*BaseReply, error)
	RemoveExportConf(ctx context.Context, in *RemoveExportConfRequest, opts ...grpc.CallOption) (*BaseReply, error)
	RemoveExportMgr(ctx context.Context, in *ExportIDRequest, opts ...grpc.CallOption) (*BaseReply, error)
	// export policy
	AddPoliciesConf(ctx context.Context, in *AddPoliciesConfRequest, opts ...grpc.CallOption) (*BaseReply, error)
	RemovePolicyConf(ctx context.Context, in *RemovePolicyConfRequest, opts ...grpc.CallOption) (*BaseReply, error)
	UpdatePolicyConf(ctx context.Context, in *UpdatePolicyConfRequest, opts ...grpc.CallOption) (*BaseReply, error)
	// config sync
	ReadLocalConfig(ctx context.Context, in *ReadLocalConfigRequest, opts ...grpc.CallOption) (Ganesha_ReadLocalConfigClient, error)
	SaveConfigToLocal(ctx context.Context, opts ...grpc.CallOption) (Ganesha_SaveConfigToLocalClient, error)
	RemoveLocalConf(ctx context.Context, in *RemoveLocalConfRequest, opts ...grpc.CallOption) (*BaseReply, error)
	RemoveLocalConfRollbak(ctx context.Context, in *RemoveLocalConfRollbakRequest, opts ...grpc.CallOption) (*BaseReply, error)
	// Get nfs stats
	GetNfsStats(ctx context.Context, in *GetNfsStatsRequest, opts ...grpc.CallOption) (*GetNfsStatsResponse, error)
}

type ganeshaClient struct {
	cc grpc.ClientConnInterface
}

func NewGaneshaClient(cc grpc.ClientConnInterface) GaneshaClient {
	return &ganeshaClient{cc}
}

func (c *ganeshaClient) AddCeafsConf(ctx context.Context, in *AddCeafsConfRequest, opts ...grpc.CallOption) (*BaseReply, error) {
	out := new(BaseReply)
	err := c.cc.Invoke(ctx, Ganesha_AddCeafsConf_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ganeshaClient) RemoveCeafsConf(ctx context.Context, in *RemoveCeafsConfRequest, opts ...grpc.CallOption) (*BaseReply, error) {
	out := new(BaseReply)
	err := c.cc.Invoke(ctx, Ganesha_RemoveCeafsConf_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ganeshaClient) SetExportMgr(ctx context.Context, in *ExportIDRequest, opts ...grpc.CallOption) (*BaseReply, error) {
	out := new(BaseReply)
	err := c.cc.Invoke(ctx, Ganesha_SetExportMgr_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ganeshaClient) UpdateExportMgr(ctx context.Context, in *ExportIDRequest, opts ...grpc.CallOption) (*BaseReply, error) {
	out := new(BaseReply)
	err := c.cc.Invoke(ctx, Ganesha_UpdateExportMgr_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ganeshaClient) AddExportConf(ctx context.Context, in *AddExportConfRequest, opts ...grpc.CallOption) (*BaseReply, error) {
	out := new(BaseReply)
	err := c.cc.Invoke(ctx, Ganesha_AddExportConf_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ganeshaClient) RemoveExportConf(ctx context.Context, in *RemoveExportConfRequest, opts ...grpc.CallOption) (*BaseReply, error) {
	out := new(BaseReply)
	err := c.cc.Invoke(ctx, Ganesha_RemoveExportConf_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ganeshaClient) RemoveExportMgr(ctx context.Context, in *ExportIDRequest, opts ...grpc.CallOption) (*BaseReply, error) {
	out := new(BaseReply)
	err := c.cc.Invoke(ctx, Ganesha_RemoveExportMgr_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ganeshaClient) AddPoliciesConf(ctx context.Context, in *AddPoliciesConfRequest, opts ...grpc.CallOption) (*BaseReply, error) {
	out := new(BaseReply)
	err := c.cc.Invoke(ctx, Ganesha_AddPoliciesConf_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ganeshaClient) RemovePolicyConf(ctx context.Context, in *RemovePolicyConfRequest, opts ...grpc.CallOption) (*BaseReply, error) {
	out := new(BaseReply)
	err := c.cc.Invoke(ctx, Ganesha_RemovePolicyConf_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ganeshaClient) UpdatePolicyConf(ctx context.Context, in *UpdatePolicyConfRequest, opts ...grpc.CallOption) (*BaseReply, error) {
	out := new(BaseReply)
	err := c.cc.Invoke(ctx, Ganesha_UpdatePolicyConf_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ganeshaClient) ReadLocalConfig(ctx context.Context, in *ReadLocalConfigRequest, opts ...grpc.CallOption) (Ganesha_ReadLocalConfigClient, error) {
	stream, err := c.cc.NewStream(ctx, &Ganesha_ServiceDesc.Streams[0], Ganesha_ReadLocalConfig_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &ganeshaReadLocalConfigClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Ganesha_ReadLocalConfigClient interface {
	Recv() (*ReadLocalConfigRes, error)
	grpc.ClientStream
}

type ganeshaReadLocalConfigClient struct {
	grpc.ClientStream
}

func (x *ganeshaReadLocalConfigClient) Recv() (*ReadLocalConfigRes, error) {
	m := new(ReadLocalConfigRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ganeshaClient) SaveConfigToLocal(ctx context.Context, opts ...grpc.CallOption) (Ganesha_SaveConfigToLocalClient, error) {
	stream, err := c.cc.NewStream(ctx, &Ganesha_ServiceDesc.Streams[1], Ganesha_SaveConfigToLocal_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &ganeshaSaveConfigToLocalClient{stream}
	return x, nil
}

type Ganesha_SaveConfigToLocalClient interface {
	Send(*SaveConfigToLocalRequest) error
	CloseAndRecv() (*BaseReply, error)
	grpc.ClientStream
}

type ganeshaSaveConfigToLocalClient struct {
	grpc.ClientStream
}

func (x *ganeshaSaveConfigToLocalClient) Send(m *SaveConfigToLocalRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *ganeshaSaveConfigToLocalClient) CloseAndRecv() (*BaseReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(BaseReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ganeshaClient) RemoveLocalConf(ctx context.Context, in *RemoveLocalConfRequest, opts ...grpc.CallOption) (*BaseReply, error) {
	out := new(BaseReply)
	err := c.cc.Invoke(ctx, Ganesha_RemoveLocalConf_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ganeshaClient) RemoveLocalConfRollbak(ctx context.Context, in *RemoveLocalConfRollbakRequest, opts ...grpc.CallOption) (*BaseReply, error) {
	out := new(BaseReply)
	err := c.cc.Invoke(ctx, Ganesha_RemoveLocalConfRollbak_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ganeshaClient) GetNfsStats(ctx context.Context, in *GetNfsStatsRequest, opts ...grpc.CallOption) (*GetNfsStatsResponse, error) {
	out := new(GetNfsStatsResponse)
	err := c.cc.Invoke(ctx, Ganesha_GetNfsStats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GaneshaServer is the server API for Ganesha service.
// All implementations must embed UnimplementedGaneshaServer
// for forward compatibility
type GaneshaServer interface {
	// init
	AddCeafsConf(context.Context, *AddCeafsConfRequest) (*BaseReply, error)
	RemoveCeafsConf(context.Context, *RemoveCeafsConfRequest) (*BaseReply, error)
	// export
	SetExportMgr(context.Context, *ExportIDRequest) (*BaseReply, error)
	UpdateExportMgr(context.Context, *ExportIDRequest) (*BaseReply, error)
	AddExportConf(context.Context, *AddExportConfRequest) (*BaseReply, error)
	RemoveExportConf(context.Context, *RemoveExportConfRequest) (*BaseReply, error)
	RemoveExportMgr(context.Context, *ExportIDRequest) (*BaseReply, error)
	// export policy
	AddPoliciesConf(context.Context, *AddPoliciesConfRequest) (*BaseReply, error)
	RemovePolicyConf(context.Context, *RemovePolicyConfRequest) (*BaseReply, error)
	UpdatePolicyConf(context.Context, *UpdatePolicyConfRequest) (*BaseReply, error)
	// config sync
	ReadLocalConfig(*ReadLocalConfigRequest, Ganesha_ReadLocalConfigServer) error
	SaveConfigToLocal(Ganesha_SaveConfigToLocalServer) error
	RemoveLocalConf(context.Context, *RemoveLocalConfRequest) (*BaseReply, error)
	RemoveLocalConfRollbak(context.Context, *RemoveLocalConfRollbakRequest) (*BaseReply, error)
	// Get nfs stats
	GetNfsStats(context.Context, *GetNfsStatsRequest) (*GetNfsStatsResponse, error)
	mustEmbedUnimplementedGaneshaServer()
}

// UnimplementedGaneshaServer must be embedded to have forward compatible implementations.
type UnimplementedGaneshaServer struct {
}

func (UnimplementedGaneshaServer) AddCeafsConf(context.Context, *AddCeafsConfRequest) (*BaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCeafsConf not implemented")
}
func (UnimplementedGaneshaServer) RemoveCeafsConf(context.Context, *RemoveCeafsConfRequest) (*BaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveCeafsConf not implemented")
}
func (UnimplementedGaneshaServer) SetExportMgr(context.Context, *ExportIDRequest) (*BaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetExportMgr not implemented")
}
func (UnimplementedGaneshaServer) UpdateExportMgr(context.Context, *ExportIDRequest) (*BaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateExportMgr not implemented")
}
func (UnimplementedGaneshaServer) AddExportConf(context.Context, *AddExportConfRequest) (*BaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddExportConf not implemented")
}
func (UnimplementedGaneshaServer) RemoveExportConf(context.Context, *RemoveExportConfRequest) (*BaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveExportConf not implemented")
}
func (UnimplementedGaneshaServer) RemoveExportMgr(context.Context, *ExportIDRequest) (*BaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveExportMgr not implemented")
}
func (UnimplementedGaneshaServer) AddPoliciesConf(context.Context, *AddPoliciesConfRequest) (*BaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPoliciesConf not implemented")
}
func (UnimplementedGaneshaServer) RemovePolicyConf(context.Context, *RemovePolicyConfRequest) (*BaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePolicyConf not implemented")
}
func (UnimplementedGaneshaServer) UpdatePolicyConf(context.Context, *UpdatePolicyConfRequest) (*BaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePolicyConf not implemented")
}
func (UnimplementedGaneshaServer) ReadLocalConfig(*ReadLocalConfigRequest, Ganesha_ReadLocalConfigServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadLocalConfig not implemented")
}
func (UnimplementedGaneshaServer) SaveConfigToLocal(Ganesha_SaveConfigToLocalServer) error {
	return status.Errorf(codes.Unimplemented, "method SaveConfigToLocal not implemented")
}
func (UnimplementedGaneshaServer) RemoveLocalConf(context.Context, *RemoveLocalConfRequest) (*BaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveLocalConf not implemented")
}
func (UnimplementedGaneshaServer) RemoveLocalConfRollbak(context.Context, *RemoveLocalConfRollbakRequest) (*BaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveLocalConfRollbak not implemented")
}
func (UnimplementedGaneshaServer) GetNfsStats(context.Context, *GetNfsStatsRequest) (*GetNfsStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNfsStats not implemented")
}
func (UnimplementedGaneshaServer) mustEmbedUnimplementedGaneshaServer() {}

// UnsafeGaneshaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GaneshaServer will
// result in compilation errors.
type UnsafeGaneshaServer interface {
	mustEmbedUnimplementedGaneshaServer()
}

func RegisterGaneshaServer(s grpc.ServiceRegistrar, srv GaneshaServer) {
	s.RegisterService(&Ganesha_ServiceDesc, srv)
}

func _Ganesha_AddCeafsConf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCeafsConfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GaneshaServer).AddCeafsConf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ganesha_AddCeafsConf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GaneshaServer).AddCeafsConf(ctx, req.(*AddCeafsConfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ganesha_RemoveCeafsConf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveCeafsConfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GaneshaServer).RemoveCeafsConf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ganesha_RemoveCeafsConf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GaneshaServer).RemoveCeafsConf(ctx, req.(*RemoveCeafsConfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ganesha_SetExportMgr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GaneshaServer).SetExportMgr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ganesha_SetExportMgr_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GaneshaServer).SetExportMgr(ctx, req.(*ExportIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ganesha_UpdateExportMgr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GaneshaServer).UpdateExportMgr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ganesha_UpdateExportMgr_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GaneshaServer).UpdateExportMgr(ctx, req.(*ExportIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ganesha_AddExportConf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddExportConfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GaneshaServer).AddExportConf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ganesha_AddExportConf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GaneshaServer).AddExportConf(ctx, req.(*AddExportConfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ganesha_RemoveExportConf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveExportConfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GaneshaServer).RemoveExportConf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ganesha_RemoveExportConf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GaneshaServer).RemoveExportConf(ctx, req.(*RemoveExportConfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ganesha_RemoveExportMgr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GaneshaServer).RemoveExportMgr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ganesha_RemoveExportMgr_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GaneshaServer).RemoveExportMgr(ctx, req.(*ExportIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ganesha_AddPoliciesConf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPoliciesConfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GaneshaServer).AddPoliciesConf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ganesha_AddPoliciesConf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GaneshaServer).AddPoliciesConf(ctx, req.(*AddPoliciesConfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ganesha_RemovePolicyConf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemovePolicyConfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GaneshaServer).RemovePolicyConf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ganesha_RemovePolicyConf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GaneshaServer).RemovePolicyConf(ctx, req.(*RemovePolicyConfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ganesha_UpdatePolicyConf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePolicyConfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GaneshaServer).UpdatePolicyConf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ganesha_UpdatePolicyConf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GaneshaServer).UpdatePolicyConf(ctx, req.(*UpdatePolicyConfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ganesha_ReadLocalConfig_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadLocalConfigRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GaneshaServer).ReadLocalConfig(m, &ganeshaReadLocalConfigServer{stream})
}

type Ganesha_ReadLocalConfigServer interface {
	Send(*ReadLocalConfigRes) error
	grpc.ServerStream
}

type ganeshaReadLocalConfigServer struct {
	grpc.ServerStream
}

func (x *ganeshaReadLocalConfigServer) Send(m *ReadLocalConfigRes) error {
	return x.ServerStream.SendMsg(m)
}

func _Ganesha_SaveConfigToLocal_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GaneshaServer).SaveConfigToLocal(&ganeshaSaveConfigToLocalServer{stream})
}

type Ganesha_SaveConfigToLocalServer interface {
	SendAndClose(*BaseReply) error
	Recv() (*SaveConfigToLocalRequest, error)
	grpc.ServerStream
}

type ganeshaSaveConfigToLocalServer struct {
	grpc.ServerStream
}

func (x *ganeshaSaveConfigToLocalServer) SendAndClose(m *BaseReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *ganeshaSaveConfigToLocalServer) Recv() (*SaveConfigToLocalRequest, error) {
	m := new(SaveConfigToLocalRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Ganesha_RemoveLocalConf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveLocalConfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GaneshaServer).RemoveLocalConf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ganesha_RemoveLocalConf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GaneshaServer).RemoveLocalConf(ctx, req.(*RemoveLocalConfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ganesha_RemoveLocalConfRollbak_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveLocalConfRollbakRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GaneshaServer).RemoveLocalConfRollbak(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ganesha_RemoveLocalConfRollbak_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GaneshaServer).RemoveLocalConfRollbak(ctx, req.(*RemoveLocalConfRollbakRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ganesha_GetNfsStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNfsStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GaneshaServer).GetNfsStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ganesha_GetNfsStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GaneshaServer).GetNfsStats(ctx, req.(*GetNfsStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Ganesha_ServiceDesc is the grpc.ServiceDesc for Ganesha service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ganesha_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pmg.Ganesha",
	HandlerType: (*GaneshaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCeafsConf",
			Handler:    _Ganesha_AddCeafsConf_Handler,
		},
		{
			MethodName: "RemoveCeafsConf",
			Handler:    _Ganesha_RemoveCeafsConf_Handler,
		},
		{
			MethodName: "SetExportMgr",
			Handler:    _Ganesha_SetExportMgr_Handler,
		},
		{
			MethodName: "UpdateExportMgr",
			Handler:    _Ganesha_UpdateExportMgr_Handler,
		},
		{
			MethodName: "AddExportConf",
			Handler:    _Ganesha_AddExportConf_Handler,
		},
		{
			MethodName: "RemoveExportConf",
			Handler:    _Ganesha_RemoveExportConf_Handler,
		},
		{
			MethodName: "RemoveExportMgr",
			Handler:    _Ganesha_RemoveExportMgr_Handler,
		},
		{
			MethodName: "AddPoliciesConf",
			Handler:    _Ganesha_AddPoliciesConf_Handler,
		},
		{
			MethodName: "RemovePolicyConf",
			Handler:    _Ganesha_RemovePolicyConf_Handler,
		},
		{
			MethodName: "UpdatePolicyConf",
			Handler:    _Ganesha_UpdatePolicyConf_Handler,
		},
		{
			MethodName: "RemoveLocalConf",
			Handler:    _Ganesha_RemoveLocalConf_Handler,
		},
		{
			MethodName: "RemoveLocalConfRollbak",
			Handler:    _Ganesha_RemoveLocalConfRollbak_Handler,
		},
		{
			MethodName: "GetNfsStats",
			Handler:    _Ganesha_GetNfsStats_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReadLocalConfig",
			Handler:       _Ganesha_ReadLocalConfig_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SaveConfigToLocal",
			Handler:       _Ganesha_SaveConfigToLocal_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "protocol/ganesha.proto",
}
