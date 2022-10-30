// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	models "github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
	mock "github.com/stretchr/testify/mock"
)

// UseCaseI is an autogenerated mock type for the UseCaseI type
type UseCaseI struct {
	mock.Mock
}

// Auth provides a mock function with given fields: cookie
func (_m *UseCaseI) Auth(cookie string) (*models.User, error) {
	ret := _m.Called(cookie)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(string) *models.User); ok {
		r0 = rf(cookie)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(cookie)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCookie provides a mock function with given fields: value
func (_m *UseCaseI) DeleteCookie(value string) error {
	ret := _m.Called(value)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SignIn provides a mock function with given fields: user
func (_m *UseCaseI) SignIn(user models.UserSignIn) (*models.User, *models.Cookie, error) {
	ret := _m.Called(user)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(models.UserSignIn) *models.User); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 *models.Cookie
	if rf, ok := ret.Get(1).(func(models.UserSignIn) *models.Cookie); ok {
		r1 = rf(user)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*models.Cookie)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(models.UserSignIn) error); ok {
		r2 = rf(user)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SignUp provides a mock function with given fields: user
func (_m *UseCaseI) SignUp(user *models.User) (*models.Cookie, error) {
	ret := _m.Called(user)

	var r0 *models.Cookie
	if rf, ok := ret.Get(0).(func(*models.User) *models.Cookie); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Cookie)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.User) error); ok {
		r1 = rf(user)
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
