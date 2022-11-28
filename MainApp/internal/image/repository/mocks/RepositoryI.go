// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	models "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	mock "github.com/stretchr/testify/mock"
)

// RepositoryI is an autogenerated mock type for the RepositoryI type
type RepositoryI struct {
	mock.Mock
}

// CreateImage provides a mock function with given fields: image
func (_m *RepositoryI) CreateImage(image *models.Image) error {
	ret := _m.Called(image)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Image) error); ok {
		r0 = rf(image)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetImage provides a mock function with given fields: imageID
func (_m *RepositoryI) GetImage(imageID uint64) (*models.Image, error) {
	ret := _m.Called(imageID)

	var r0 *models.Image
	if rf, ok := ret.Get(0).(func(uint64) *models.Image); ok {
		r0 = rf(imageID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Image)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(imageID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostImages provides a mock function with given fields: postID
func (_m *RepositoryI) GetPostImages(postID uint64) ([]*models.Image, error) {
	ret := _m.Called(postID)

	var r0 []*models.Image
	if rf, ok := ret.Get(0).(func(uint64) []*models.Image); ok {
		r0 = rf(postID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Image)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(postID)
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
