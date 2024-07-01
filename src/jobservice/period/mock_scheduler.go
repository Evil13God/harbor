// Code generated by mockery v2.42.2. DO NOT EDIT.

package period

import mock "github.com/stretchr/testify/mock"

// MockScheduler is an autogenerated mock type for the Scheduler type
type MockScheduler struct {
	mock.Mock
}

// Schedule provides a mock function with given fields: policy
func (_m *MockScheduler) Schedule(policy *Policy) (int64, error) {
	ret := _m.Called(policy)

	if len(ret) == 0 {
		panic("no return value specified for Schedule")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(*Policy) (int64, error)); ok {
		return rf(policy)
	}
	if rf, ok := ret.Get(0).(func(*Policy) int64); ok {
		r0 = rf(policy)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(*Policy) error); ok {
		r1 = rf(policy)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Start provides a mock function with given fields:
func (_m *MockScheduler) Start() {
	_m.Called()
}

// UnSchedule provides a mock function with given fields: policyID
func (_m *MockScheduler) UnSchedule(policyID string) error {
	ret := _m.Called(policyID)

	if len(ret) == 0 {
		panic("no return value specified for UnSchedule")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(policyID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockScheduler creates a new instance of MockScheduler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockScheduler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockScheduler {
	mock := &MockScheduler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
