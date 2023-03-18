// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: order_service.proto

package pb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Pong, error)
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
	DeleteOrder(ctx context.Context, in *DeleteOrderRequest, opts ...grpc.CallOption) (*DeleteOrderResponse, error)
	UpdateOrder(ctx context.Context, in *UpdateOrderStatusRequest, opts ...grpc.CallOption) (*UpdateOrderStatusResponse, error)
	HandleOrder(ctx context.Context, in *HandleOrderRequest, opts ...grpc.CallOption) (*HandleOrderResponse, error)
	GetWaitingOrderBySupplier(ctx context.Context, in *GetWaitingOrderBySupplierRequest, opts ...grpc.CallOption) (*GetWaitingOrderBySupplierResponse, error)
	GetWaitingOrderByCustomer(ctx context.Context, in *GetWaitingOrderByCustomerRequest, opts ...grpc.CallOption) (*GetWaitingOrderByCustomerResponse, error)
	GetOrderByProductId(ctx context.Context, in *GetOrderByProductIdRequest, opts ...grpc.CallOption) (*GetOrderByProductIdResponse, error)
	CheckOrderIsHandled(ctx context.Context, in *CheckOrderIsHandledRequest, opts ...grpc.CallOption) (*CheckOrderIsHandledResponse, error)
	GetHandledOrderByCustomer(ctx context.Context, in *GetHandledOrderByCustomerRequest, opts ...grpc.CallOption) (*GetHandledOrderByCustomerResponse, error)
	GetHandledOrderBySupllier(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetHandledOrderBySupplierResponse, error)
	GetSoldProduct(ctx context.Context, in *GetSoldProductRequest, opts ...grpc.CallOption) (*GetSoldProductResponse, error)
	GetCancelOrderByCustomer(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetHandledOrderByCustomerResponse, error)
	GetCancelOrderBySupplier(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetHandledOrderBySupplierResponse, error)
	GetAddressOrder(ctx context.Context, in *GetAddressOrderRequest, opts ...grpc.CallOption) (*GetAddressOrderResponse, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := c.cc.Invoke(ctx, "/ecommerce.OrderService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	out := new(CreateOrderResponse)
	err := c.cc.Invoke(ctx, "/ecommerce.OrderService/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) DeleteOrder(ctx context.Context, in *DeleteOrderRequest, opts ...grpc.CallOption) (*DeleteOrderResponse, error) {
	out := new(DeleteOrderResponse)
	err := c.cc.Invoke(ctx, "/ecommerce.OrderService/DeleteOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) UpdateOrder(ctx context.Context, in *UpdateOrderStatusRequest, opts ...grpc.CallOption) (*UpdateOrderStatusResponse, error) {
	out := new(UpdateOrderStatusResponse)
	err := c.cc.Invoke(ctx, "/ecommerce.OrderService/UpdateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) HandleOrder(ctx context.Context, in *HandleOrderRequest, opts ...grpc.CallOption) (*HandleOrderResponse, error) {
	out := new(HandleOrderResponse)
	err := c.cc.Invoke(ctx, "/ecommerce.OrderService/HandleOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetWaitingOrderBySupplier(ctx context.Context, in *GetWaitingOrderBySupplierRequest, opts ...grpc.CallOption) (*GetWaitingOrderBySupplierResponse, error) {
	out := new(GetWaitingOrderBySupplierResponse)
	err := c.cc.Invoke(ctx, "/ecommerce.OrderService/GetWaitingOrderBySupplier", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetWaitingOrderByCustomer(ctx context.Context, in *GetWaitingOrderByCustomerRequest, opts ...grpc.CallOption) (*GetWaitingOrderByCustomerResponse, error) {
	out := new(GetWaitingOrderByCustomerResponse)
	err := c.cc.Invoke(ctx, "/ecommerce.OrderService/GetWaitingOrderByCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrderByProductId(ctx context.Context, in *GetOrderByProductIdRequest, opts ...grpc.CallOption) (*GetOrderByProductIdResponse, error) {
	out := new(GetOrderByProductIdResponse)
	err := c.cc.Invoke(ctx, "/ecommerce.OrderService/GetOrderByProductId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CheckOrderIsHandled(ctx context.Context, in *CheckOrderIsHandledRequest, opts ...grpc.CallOption) (*CheckOrderIsHandledResponse, error) {
	out := new(CheckOrderIsHandledResponse)
	err := c.cc.Invoke(ctx, "/ecommerce.OrderService/CheckOrderIsHandled", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetHandledOrderByCustomer(ctx context.Context, in *GetHandledOrderByCustomerRequest, opts ...grpc.CallOption) (*GetHandledOrderByCustomerResponse, error) {
	out := new(GetHandledOrderByCustomerResponse)
	err := c.cc.Invoke(ctx, "/ecommerce.OrderService/GetHandledOrderByCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetHandledOrderBySupllier(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetHandledOrderBySupplierResponse, error) {
	out := new(GetHandledOrderBySupplierResponse)
	err := c.cc.Invoke(ctx, "/ecommerce.OrderService/GetHandledOrderBySupllier", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetSoldProduct(ctx context.Context, in *GetSoldProductRequest, opts ...grpc.CallOption) (*GetSoldProductResponse, error) {
	out := new(GetSoldProductResponse)
	err := c.cc.Invoke(ctx, "/ecommerce.OrderService/GetSoldProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetCancelOrderByCustomer(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetHandledOrderByCustomerResponse, error) {
	out := new(GetHandledOrderByCustomerResponse)
	err := c.cc.Invoke(ctx, "/ecommerce.OrderService/GetCancelOrderByCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetCancelOrderBySupplier(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetHandledOrderBySupplierResponse, error) {
	out := new(GetHandledOrderBySupplierResponse)
	err := c.cc.Invoke(ctx, "/ecommerce.OrderService/GetCancelOrderBySupplier", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetAddressOrder(ctx context.Context, in *GetAddressOrderRequest, opts ...grpc.CallOption) (*GetAddressOrderResponse, error) {
	out := new(GetAddressOrderResponse)
	err := c.cc.Invoke(ctx, "/ecommerce.OrderService/GetAddressOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations must embed UnimplementedOrderServiceServer
// for forward compatibility
type OrderServiceServer interface {
	Ping(context.Context, *empty.Empty) (*Pong, error)
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error)
	DeleteOrder(context.Context, *DeleteOrderRequest) (*DeleteOrderResponse, error)
	UpdateOrder(context.Context, *UpdateOrderStatusRequest) (*UpdateOrderStatusResponse, error)
	HandleOrder(context.Context, *HandleOrderRequest) (*HandleOrderResponse, error)
	GetWaitingOrderBySupplier(context.Context, *GetWaitingOrderBySupplierRequest) (*GetWaitingOrderBySupplierResponse, error)
	GetWaitingOrderByCustomer(context.Context, *GetWaitingOrderByCustomerRequest) (*GetWaitingOrderByCustomerResponse, error)
	GetOrderByProductId(context.Context, *GetOrderByProductIdRequest) (*GetOrderByProductIdResponse, error)
	CheckOrderIsHandled(context.Context, *CheckOrderIsHandledRequest) (*CheckOrderIsHandledResponse, error)
	GetHandledOrderByCustomer(context.Context, *GetHandledOrderByCustomerRequest) (*GetHandledOrderByCustomerResponse, error)
	GetHandledOrderBySupllier(context.Context, *empty.Empty) (*GetHandledOrderBySupplierResponse, error)
	GetSoldProduct(context.Context, *GetSoldProductRequest) (*GetSoldProductResponse, error)
	GetCancelOrderByCustomer(context.Context, *empty.Empty) (*GetHandledOrderByCustomerResponse, error)
	GetCancelOrderBySupplier(context.Context, *empty.Empty) (*GetHandledOrderBySupplierResponse, error)
	GetAddressOrder(context.Context, *GetAddressOrderRequest) (*GetAddressOrderResponse, error)
	mustEmbedUnimplementedOrderServiceServer()
}

// UnimplementedOrderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (UnimplementedOrderServiceServer) Ping(context.Context, *empty.Empty) (*Pong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedOrderServiceServer) CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderServiceServer) DeleteOrder(context.Context, *DeleteOrderRequest) (*DeleteOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrder not implemented")
}
func (UnimplementedOrderServiceServer) UpdateOrder(context.Context, *UpdateOrderStatusRequest) (*UpdateOrderStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrder not implemented")
}
func (UnimplementedOrderServiceServer) HandleOrder(context.Context, *HandleOrderRequest) (*HandleOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleOrder not implemented")
}
func (UnimplementedOrderServiceServer) GetWaitingOrderBySupplier(context.Context, *GetWaitingOrderBySupplierRequest) (*GetWaitingOrderBySupplierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWaitingOrderBySupplier not implemented")
}
func (UnimplementedOrderServiceServer) GetWaitingOrderByCustomer(context.Context, *GetWaitingOrderByCustomerRequest) (*GetWaitingOrderByCustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWaitingOrderByCustomer not implemented")
}
func (UnimplementedOrderServiceServer) GetOrderByProductId(context.Context, *GetOrderByProductIdRequest) (*GetOrderByProductIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderByProductId not implemented")
}
func (UnimplementedOrderServiceServer) CheckOrderIsHandled(context.Context, *CheckOrderIsHandledRequest) (*CheckOrderIsHandledResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckOrderIsHandled not implemented")
}
func (UnimplementedOrderServiceServer) GetHandledOrderByCustomer(context.Context, *GetHandledOrderByCustomerRequest) (*GetHandledOrderByCustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHandledOrderByCustomer not implemented")
}
func (UnimplementedOrderServiceServer) GetHandledOrderBySupllier(context.Context, *empty.Empty) (*GetHandledOrderBySupplierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHandledOrderBySupllier not implemented")
}
func (UnimplementedOrderServiceServer) GetSoldProduct(context.Context, *GetSoldProductRequest) (*GetSoldProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSoldProduct not implemented")
}
func (UnimplementedOrderServiceServer) GetCancelOrderByCustomer(context.Context, *empty.Empty) (*GetHandledOrderByCustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCancelOrderByCustomer not implemented")
}
func (UnimplementedOrderServiceServer) GetCancelOrderBySupplier(context.Context, *empty.Empty) (*GetHandledOrderBySupplierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCancelOrderBySupplier not implemented")
}
func (UnimplementedOrderServiceServer) GetAddressOrder(context.Context, *GetAddressOrderRequest) (*GetAddressOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddressOrder not implemented")
}
func (UnimplementedOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ecommerce.OrderService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).Ping(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ecommerce.OrderService/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_DeleteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).DeleteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ecommerce.OrderService/DeleteOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).DeleteOrder(ctx, req.(*DeleteOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_UpdateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).UpdateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ecommerce.OrderService/UpdateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).UpdateOrder(ctx, req.(*UpdateOrderStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_HandleOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).HandleOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ecommerce.OrderService/HandleOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).HandleOrder(ctx, req.(*HandleOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetWaitingOrderBySupplier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWaitingOrderBySupplierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetWaitingOrderBySupplier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ecommerce.OrderService/GetWaitingOrderBySupplier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetWaitingOrderBySupplier(ctx, req.(*GetWaitingOrderBySupplierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetWaitingOrderByCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWaitingOrderByCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetWaitingOrderByCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ecommerce.OrderService/GetWaitingOrderByCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetWaitingOrderByCustomer(ctx, req.(*GetWaitingOrderByCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrderByProductId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderByProductIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrderByProductId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ecommerce.OrderService/GetOrderByProductId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrderByProductId(ctx, req.(*GetOrderByProductIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CheckOrderIsHandled_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckOrderIsHandledRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CheckOrderIsHandled(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ecommerce.OrderService/CheckOrderIsHandled",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CheckOrderIsHandled(ctx, req.(*CheckOrderIsHandledRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetHandledOrderByCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHandledOrderByCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetHandledOrderByCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ecommerce.OrderService/GetHandledOrderByCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetHandledOrderByCustomer(ctx, req.(*GetHandledOrderByCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetHandledOrderBySupllier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetHandledOrderBySupllier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ecommerce.OrderService/GetHandledOrderBySupllier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetHandledOrderBySupllier(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetSoldProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSoldProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetSoldProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ecommerce.OrderService/GetSoldProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetSoldProduct(ctx, req.(*GetSoldProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetCancelOrderByCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetCancelOrderByCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ecommerce.OrderService/GetCancelOrderByCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetCancelOrderByCustomer(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetCancelOrderBySupplier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetCancelOrderBySupplier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ecommerce.OrderService/GetCancelOrderBySupplier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetCancelOrderBySupplier(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetAddressOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAddressOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetAddressOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ecommerce.OrderService/GetAddressOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetAddressOrder(ctx, req.(*GetAddressOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ecommerce.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _OrderService_Ping_Handler,
		},
		{
			MethodName: "CreateOrder",
			Handler:    _OrderService_CreateOrder_Handler,
		},
		{
			MethodName: "DeleteOrder",
			Handler:    _OrderService_DeleteOrder_Handler,
		},
		{
			MethodName: "UpdateOrder",
			Handler:    _OrderService_UpdateOrder_Handler,
		},
		{
			MethodName: "HandleOrder",
			Handler:    _OrderService_HandleOrder_Handler,
		},
		{
			MethodName: "GetWaitingOrderBySupplier",
			Handler:    _OrderService_GetWaitingOrderBySupplier_Handler,
		},
		{
			MethodName: "GetWaitingOrderByCustomer",
			Handler:    _OrderService_GetWaitingOrderByCustomer_Handler,
		},
		{
			MethodName: "GetOrderByProductId",
			Handler:    _OrderService_GetOrderByProductId_Handler,
		},
		{
			MethodName: "CheckOrderIsHandled",
			Handler:    _OrderService_CheckOrderIsHandled_Handler,
		},
		{
			MethodName: "GetHandledOrderByCustomer",
			Handler:    _OrderService_GetHandledOrderByCustomer_Handler,
		},
		{
			MethodName: "GetHandledOrderBySupllier",
			Handler:    _OrderService_GetHandledOrderBySupllier_Handler,
		},
		{
			MethodName: "GetSoldProduct",
			Handler:    _OrderService_GetSoldProduct_Handler,
		},
		{
			MethodName: "GetCancelOrderByCustomer",
			Handler:    _OrderService_GetCancelOrderByCustomer_Handler,
		},
		{
			MethodName: "GetCancelOrderBySupplier",
			Handler:    _OrderService_GetCancelOrderBySupplier_Handler,
		},
		{
			MethodName: "GetAddressOrder",
			Handler:    _OrderService_GetAddressOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order_service.proto",
}
