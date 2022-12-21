// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	__ "github.com/go-park-mail-ru/2022_2_TikTikIVProd/AttachmentMicroservice/proto"
	mock "github.com/stretchr/testify/mock"
)

// UseCaseI is an autogenerated mock type for the UseCaseI type
type UseCaseI struct {
	mock.Mock
}

// CreateAttachment provides a mock function with given fields: _a0
func (_m *UseCaseI) CreateAttachment(_a0 *__.Attachment) (*__.Nothing, error) {
	ret := _m.Called(_a0)

	var r0 *__.Nothing
	if rf, ok := ret.Get(0).(func(*__.Attachment) *__.Nothing); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.Nothing)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*__.Attachment) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAttachment provides a mock function with given fields: _a0
func (_m *UseCaseI) GetAttachment(_a0 *__.AttachmentId) (*__.Attachment, error) {
	ret := _m.Called(_a0)

	var r0 *__.Attachment
	if rf, ok := ret.Get(0).(func(*__.AttachmentId) *__.Attachment); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.Attachment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*__.AttachmentId) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMessageAttachments provides a mock function with given fields: request
func (_m *UseCaseI) GetMessageAttachments(request *__.GetMessageAttachmentsRequest) (*__.GetMessageAttachmentsResponse, error) {
	ret := _m.Called(request)

	var r0 *__.GetMessageAttachmentsResponse
	if rf, ok := ret.Get(0).(func(*__.GetMessageAttachmentsRequest) *__.GetMessageAttachmentsResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.GetMessageAttachmentsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*__.GetMessageAttachmentsRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostAttachments provides a mock function with given fields: _a0
func (_m *UseCaseI) GetPostAttachments(_a0 *__.GetPostAttachmentsRequest) (*__.GetPostAttachmentsResponse, error) {
	ret := _m.Called(_a0)

	var r0 *__.GetPostAttachmentsResponse
	if rf, ok := ret.Get(0).(func(*__.GetPostAttachmentsRequest) *__.GetPostAttachmentsResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.GetPostAttachmentsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*__.GetPostAttachmentsRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUseCaseI interface {
	mock.TestingT
	Cleanup(func())
}

// NewUseCaseI creates a new instance of UseCaseI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUseCaseI(t mockConstructorTestingTNewUseCaseI) *UseCaseI {
	mock := &UseCaseI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}