// Code generated by mockery v2.49.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UrlHandler is an autogenerated mock type for the UrlHandler type
type UrlHandler struct {
	mock.Mock
}

// DeleteURL provides a mock function with given fields: alias
func (_m *UrlHandler) DeleteURL(alias string) error {
	ret := _m.Called(alias)

	if len(ret) == 0 {
		panic("no return value specified for DeleteURL")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(alias)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUrl provides a mock function with given fields: alias
func (_m *UrlHandler) GetUrl(alias string) (string, error) {
	ret := _m.Called(alias)

	if len(ret) == 0 {
		panic("no return value specified for GetUrl")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(alias)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(alias)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(alias)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveUrl provides a mock function with given fields: urlToSave, alias
func (_m *UrlHandler) SaveUrl(urlToSave string, alias string) error {
	ret := _m.Called(urlToSave, alias)

	if len(ret) == 0 {
		panic("no return value specified for SaveUrl")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(urlToSave, alias)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUrlHandler creates a new instance of UrlHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUrlHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *UrlHandler {
	mock := &UrlHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
