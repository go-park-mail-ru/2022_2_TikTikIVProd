// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.11
// source: attachment.proto

package __

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

// AttachmentsClient is the client API for Attachments service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AttachmentsClient interface {
	GetPostAttachments(ctx context.Context, in *GetPostAttachmentsRequest, opts ...grpc.CallOption) (*GetPostAttachmentsResponse, error)
	GetMessageAttachments(ctx context.Context, in *GetMessageAttachmentsRequest, opts ...grpc.CallOption) (*GetMessageAttachmentsResponse, error)
	GetAttachment(ctx context.Context, in *AttachmentId, opts ...grpc.CallOption) (*Attachment, error)
	CreateAttachment(ctx context.Context, in *Attachment, opts ...grpc.CallOption) (*AttachmentId, error)
	AddAttachmentsToMessage(ctx context.Context, in *AddAttachmentsToMessageRequest, opts ...grpc.CallOption) (*Nothing, error)
}

type attachmentsClient struct {
	cc grpc.ClientConnInterface
}

func NewAttachmentsClient(cc grpc.ClientConnInterface) AttachmentsClient {
	return &attachmentsClient{cc}
}

func (c *attachmentsClient) GetPostAttachments(ctx context.Context, in *GetPostAttachmentsRequest, opts ...grpc.CallOption) (*GetPostAttachmentsResponse, error) {
	out := new(GetPostAttachmentsResponse)
	err := c.cc.Invoke(ctx, "/attachment.Attachments/GetPostAttachments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachmentsClient) GetMessageAttachments(ctx context.Context, in *GetMessageAttachmentsRequest, opts ...grpc.CallOption) (*GetMessageAttachmentsResponse, error) {
	out := new(GetMessageAttachmentsResponse)
	err := c.cc.Invoke(ctx, "/attachment.Attachments/GetMessageAttachments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachmentsClient) GetAttachment(ctx context.Context, in *AttachmentId, opts ...grpc.CallOption) (*Attachment, error) {
	out := new(Attachment)
	err := c.cc.Invoke(ctx, "/attachment.Attachments/GetAttachment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachmentsClient) CreateAttachment(ctx context.Context, in *Attachment, opts ...grpc.CallOption) (*AttachmentId, error) {
	out := new(AttachmentId)
	err := c.cc.Invoke(ctx, "/attachment.Attachments/CreateAttachment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachmentsClient) AddAttachmentsToMessage(ctx context.Context, in *AddAttachmentsToMessageRequest, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/attachment.Attachments/AddAttachmentsToMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AttachmentsServer is the server API for Attachments service.
// All implementations must embed UnimplementedAttachmentsServer
// for forward compatibility
type AttachmentsServer interface {
	GetPostAttachments(context.Context, *GetPostAttachmentsRequest) (*GetPostAttachmentsResponse, error)
	GetMessageAttachments(context.Context, *GetMessageAttachmentsRequest) (*GetMessageAttachmentsResponse, error)
	GetAttachment(context.Context, *AttachmentId) (*Attachment, error)
	CreateAttachment(context.Context, *Attachment) (*AttachmentId, error)
	AddAttachmentsToMessage(context.Context, *AddAttachmentsToMessageRequest) (*Nothing, error)
	mustEmbedUnimplementedAttachmentsServer()
}

// UnimplementedAttachmentsServer must be embedded to have forward compatible implementations.
type UnimplementedAttachmentsServer struct {
}

func (UnimplementedAttachmentsServer) GetPostAttachments(context.Context, *GetPostAttachmentsRequest) (*GetPostAttachmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostAttachments not implemented")
}
func (UnimplementedAttachmentsServer) GetMessageAttachments(context.Context, *GetMessageAttachmentsRequest) (*GetMessageAttachmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessageAttachments not implemented")
}
func (UnimplementedAttachmentsServer) GetAttachment(context.Context, *AttachmentId) (*Attachment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttachment not implemented")
}
func (UnimplementedAttachmentsServer) CreateAttachment(context.Context, *Attachment) (*AttachmentId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAttachment not implemented")
}
func (UnimplementedAttachmentsServer) AddAttachmentsToMessage(context.Context, *AddAttachmentsToMessageRequest) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAttachmentsToMessage not implemented")
}
func (UnimplementedAttachmentsServer) mustEmbedUnimplementedAttachmentsServer() {}

// UnsafeAttachmentsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AttachmentsServer will
// result in compilation errors.
type UnsafeAttachmentsServer interface {
	mustEmbedUnimplementedAttachmentsServer()
}

func RegisterAttachmentsServer(s grpc.ServiceRegistrar, srv AttachmentsServer) {
	s.RegisterService(&Attachments_ServiceDesc, srv)
}

func _Attachments_GetPostAttachments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostAttachmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachmentsServer).GetPostAttachments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/attachment.Attachments/GetPostAttachments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachmentsServer).GetPostAttachments(ctx, req.(*GetPostAttachmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attachments_GetMessageAttachments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageAttachmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachmentsServer).GetMessageAttachments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/attachment.Attachments/GetMessageAttachments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachmentsServer).GetMessageAttachments(ctx, req.(*GetMessageAttachmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attachments_GetAttachment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttachmentId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachmentsServer).GetAttachment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/attachment.Attachments/GetAttachment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachmentsServer).GetAttachment(ctx, req.(*AttachmentId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attachments_CreateAttachment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Attachment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachmentsServer).CreateAttachment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/attachment.Attachments/CreateAttachment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachmentsServer).CreateAttachment(ctx, req.(*Attachment))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attachments_AddAttachmentsToMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAttachmentsToMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachmentsServer).AddAttachmentsToMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/attachment.Attachments/AddAttachmentsToMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachmentsServer).AddAttachmentsToMessage(ctx, req.(*AddAttachmentsToMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Attachments_ServiceDesc is the grpc.ServiceDesc for Attachments service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Attachments_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "attachment.Attachments",
	HandlerType: (*AttachmentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPostAttachments",
			Handler:    _Attachments_GetPostAttachments_Handler,
		},
		{
			MethodName: "GetMessageAttachments",
			Handler:    _Attachments_GetMessageAttachments_Handler,
		},
		{
			MethodName: "GetAttachment",
			Handler:    _Attachments_GetAttachment_Handler,
		},
		{
			MethodName: "CreateAttachment",
			Handler:    _Attachments_CreateAttachment_Handler,
		},
		{
			MethodName: "AddAttachmentsToMessage",
			Handler:    _Attachments_AddAttachmentsToMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "attachment.proto",
}
