// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	__ "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/proto/image"

	mock "github.com/stretchr/testify/mock"
)

// ImagesServer is an autogenerated mock type for the ImagesServer type
type ImagesServer struct {
	mock.Mock
}

// CreateImage provides a mock function with given fields: _a0, _a1
func (_m *ImagesServer) CreateImage(_a0 context.Context, _a1 *__.Image) (*__.ImageId, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *__.ImageId
	if rf, ok := ret.Get(0).(func(context.Context, *__.Image) *__.ImageId); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.ImageId)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *__.Image) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetImage provides a mock function with given fields: _a0, _a1
func (_m *ImagesServer) GetImage(_a0 context.Context, _a1 *__.ImageId) (*__.Image, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *__.Image
	if rf, ok := ret.Get(0).(func(context.Context, *__.ImageId) *__.Image); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.Image)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *__.ImageId) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostImages provides a mock function with given fields: _a0, _a1
func (_m *ImagesServer) GetPostImages(_a0 context.Context, _a1 *__.GetPostImagesRequest) (*__.GetPostImagesResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *__.GetPostImagesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *__.GetPostImagesRequest) *__.GetPostImagesResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.GetPostImagesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *__.GetPostImagesRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedImagesServer provides a mock function with given fields:
func (_m *ImagesServer) mustEmbedUnimplementedImagesServer() {
	_m.Called()
}

type mockConstructorTestingTNewImagesServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewImagesServer creates a new instance of ImagesServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewImagesServer(t mockConstructorTestingTNewImagesServer) *ImagesServer {
	mock := &ImagesServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
