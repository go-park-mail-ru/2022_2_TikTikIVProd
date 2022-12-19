// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	__ "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/proto/attachment"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// AttachmentsClient is an autogenerated mock type for the AttachmentsClient type
type AttachmentsClient struct {
	mock.Mock
}

// CreateAttachment provides a mock function with given fields: ctx, in, opts
func (_m *AttachmentsClient) CreateAttachment(ctx context.Context, in *__.Attachment, opts ...grpc.CallOption) (*__.AttachmentId, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *__.AttachmentId
	if rf, ok := ret.Get(0).(func(context.Context, *__.Attachment, ...grpc.CallOption) *__.AttachmentId); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.AttachmentId)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *__.Attachment, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAttachment provides a mock function with given fields: ctx, in, opts
func (_m *AttachmentsClient) GetAttachment(ctx context.Context, in *__.AttachmentId, opts ...grpc.CallOption) (*__.Attachment, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *__.Attachment
	if rf, ok := ret.Get(0).(func(context.Context, *__.AttachmentId, ...grpc.CallOption) *__.Attachment); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.Attachment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *__.AttachmentId, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostAttachments provides a mock function with given fields: ctx, in, opts
func (_m *AttachmentsClient) GetPostAttachments(ctx context.Context, in *__.GetPostAttachmentsRequest, opts ...grpc.CallOption) (*__.GetPostAttachmentsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *__.GetPostAttachmentsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *__.GetPostAttachmentsRequest, ...grpc.CallOption) *__.GetPostAttachmentsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.GetPostAttachmentsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *__.GetPostAttachmentsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAttachmentsClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewAttachmentsClient creates a new instance of AttachmentsClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAttachmentsClient(t mockConstructorTestingTNewAttachmentsClient) *AttachmentsClient {
	mock := &AttachmentsClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
