// Code generated by mockery v2.39.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Producer is an autogenerated mock type for the Producer type
type Producer struct {
	mock.Mock
}

// Close provides a mock function with given fields: ctx
func (_m *Producer) Close(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Publish provides a mock function with given fields: topic, message, kafkaPartition
func (_m *Producer) Publish(topic string, message []byte, kafkaPartition *int32) {
	_m.Called(topic, message, kafkaPartition)
}

// NewProducer creates a new instance of Producer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProducer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Producer {
	mock := &Producer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
