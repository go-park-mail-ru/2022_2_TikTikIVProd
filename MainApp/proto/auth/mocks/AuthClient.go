// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	__ "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/proto/auth"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// AuthClient is an autogenerated mock type for the AuthClient type
type AuthClient struct {
	mock.Mock
}

// CreateCookie provides a mock function with given fields: ctx, in, opts
func (_m *AuthClient) CreateCookie(ctx context.Context, in *__.Cookie, opts ...grpc.CallOption) (*__.Nothing, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *__.Nothing
	if rf, ok := ret.Get(0).(func(context.Context, *__.Cookie, ...grpc.CallOption) *__.Nothing); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.Nothing)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *__.Cookie, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCookie provides a mock function with given fields: ctx, in, opts
func (_m *AuthClient) DeleteCookie(ctx context.Context, in *__.ValueCookieRequest, opts ...grpc.CallOption) (*__.Nothing, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *__.Nothing
	if rf, ok := ret.Get(0).(func(context.Context, *__.ValueCookieRequest, ...grpc.CallOption) *__.Nothing); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.Nothing)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *__.ValueCookieRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCookie provides a mock function with given fields: ctx, in, opts
func (_m *AuthClient) GetCookie(ctx context.Context, in *__.ValueCookieRequest, opts ...grpc.CallOption) (*__.GetCookieResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *__.GetCookieResponse
	if rf, ok := ret.Get(0).(func(context.Context, *__.ValueCookieRequest, ...grpc.CallOption) *__.GetCookieResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.GetCookieResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *__.ValueCookieRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAuthClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthClient creates a new instance of AuthClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthClient(t mockConstructorTestingTNewAuthClient) *AuthClient {
	mock := &AuthClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
