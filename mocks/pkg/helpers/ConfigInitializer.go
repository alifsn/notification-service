// Code generated by mockery v2.39.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ConfigInitializer is an autogenerated mock type for the ConfigInitializer type
type ConfigInitializer struct {
	mock.Mock
}

// InitConfig provides a mock function with given fields: privateKeyConf, publicKeyConf, privateKeyRefConf, publicKeyRefConf
func (_m *ConfigInitializer) InitConfig(privateKeyConf string, publicKeyConf string, privateKeyRefConf string, publicKeyRefConf string) {
	_m.Called(privateKeyConf, publicKeyConf, privateKeyRefConf, publicKeyRefConf)
}

// NewConfigInitializer creates a new instance of ConfigInitializer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConfigInitializer(t interface {
	mock.TestingT
	Cleanup(func())
}) *ConfigInitializer {
	mock := &ConfigInitializer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
