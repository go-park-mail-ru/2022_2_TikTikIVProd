// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UnsafeUsersServer is an autogenerated mock type for the UnsafeUsersServer type
type UnsafeUsersServer struct {
	mock.Mock
}

// mustEmbedUnimplementedUsersServer provides a mock function with given fields:
func (_m *UnsafeUsersServer) mustEmbedUnimplementedUsersServer() {
	_m.Called()
}

type mockConstructorTestingTNewUnsafeUsersServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewUnsafeUsersServer creates a new instance of UnsafeUsersServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUnsafeUsersServer(t mockConstructorTestingTNewUnsafeUsersServer) *UnsafeUsersServer {
	mock := &UnsafeUsersServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
