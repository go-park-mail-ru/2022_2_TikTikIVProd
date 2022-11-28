// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	__ "github.com/go-park-mail-ru/2022_2_TikTikIVProd/ChatMicroservice/proto"
	mock "github.com/stretchr/testify/mock"
)

// UseCaseI is an autogenerated mock type for the UseCaseI type
type UseCaseI struct {
	mock.Mock
}

// CreateDialog provides a mock function with given fields: _a0
func (_m *UseCaseI) CreateDialog(_a0 *__.Dialog) (*__.Nothing, error) {
	ret := _m.Called(_a0)

	var r0 *__.Nothing
	if rf, ok := ret.Get(0).(func(*__.Dialog) *__.Nothing); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.Nothing)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*__.Dialog) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateMessage provides a mock function with given fields: _a0
func (_m *UseCaseI) CreateMessage(_a0 *__.Message) (*__.Nothing, error) {
	ret := _m.Called(_a0)

	var r0 *__.Nothing
	if rf, ok := ret.Get(0).(func(*__.Message) *__.Nothing); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.Nothing)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*__.Message) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectAllDialogs provides a mock function with given fields: _a0
func (_m *UseCaseI) SelectAllDialogs(_a0 *__.SelectAllDialogsRequest) (*__.SelectAllDialogsResponse, error) {
	ret := _m.Called(_a0)

	var r0 *__.SelectAllDialogsResponse
	if rf, ok := ret.Get(0).(func(*__.SelectAllDialogsRequest) *__.SelectAllDialogsResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.SelectAllDialogsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*__.SelectAllDialogsRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectDialog provides a mock function with given fields: _a0
func (_m *UseCaseI) SelectDialog(_a0 *__.DialogId) (*__.Dialog, error) {
	ret := _m.Called(_a0)

	var r0 *__.Dialog
	if rf, ok := ret.Get(0).(func(*__.DialogId) *__.Dialog); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.Dialog)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*__.DialogId) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectDialogByUsers provides a mock function with given fields: _a0
func (_m *UseCaseI) SelectDialogByUsers(_a0 *__.SelectDialogByUsersRequest) (*__.Dialog, error) {
	ret := _m.Called(_a0)

	var r0 *__.Dialog
	if rf, ok := ret.Get(0).(func(*__.SelectDialogByUsersRequest) *__.Dialog); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.Dialog)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*__.SelectDialogByUsersRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectMessages provides a mock function with given fields: _a0
func (_m *UseCaseI) SelectMessages(_a0 *__.DialogId) (*__.SelectMessagesResponse, error) {
	ret := _m.Called(_a0)

	var r0 *__.SelectMessagesResponse
	if rf, ok := ret.Get(0).(func(*__.DialogId) *__.SelectMessagesResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.SelectMessagesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*__.DialogId) error); ok {
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
