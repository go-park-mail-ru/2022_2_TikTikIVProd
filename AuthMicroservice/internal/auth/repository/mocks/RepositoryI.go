// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	models "github.com/go-park-mail-ru/2022_2_TikTikIVProd/AuthMicroservice/models"
	mock "github.com/stretchr/testify/mock"
)

// RepositoryI is an autogenerated mock type for the RepositoryI type
type RepositoryI struct {
	mock.Mock
}

// CreateCookie provides a mock function with given fields: cookie
func (_m *RepositoryI) CreateCookie(cookie *models.Cookie) error {
	ret := _m.Called(cookie)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Cookie) error); ok {
		r0 = rf(cookie)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCookie provides a mock function with given fields: value
func (_m *RepositoryI) DeleteCookie(value string) error {
	ret := _m.Called(value)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCookie provides a mock function with given fields: value
func (_m *RepositoryI) GetCookie(value string) (string, error) {
	ret := _m.Called(value)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepositoryI interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepositoryI creates a new instance of RepositoryI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepositoryI(t mockConstructorTestingTNewRepositoryI) *RepositoryI {
	mock := &RepositoryI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
