// Code generated by mockery v2.42.2. DO NOT EDIT.

package task

import (
	context "context"

	q "github.com/goharbor/harbor/src/lib/q"
	mock "github.com/stretchr/testify/mock"
)

// mockTaskManager is an autogenerated mock type for the Manager type
type mockTaskManager struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx, query
func (_m *mockTaskManager) Count(ctx context.Context, query *q.Query) (int64, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for Count")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) (int64, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) int64); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, executionID, job, extraAttrs
func (_m *mockTaskManager) Create(ctx context.Context, executionID int64, job *Job, extraAttrs ...map[string]interface{}) (int64, error) {
	_va := make([]interface{}, len(extraAttrs))
	for _i := range extraAttrs {
		_va[_i] = extraAttrs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, executionID, job)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, *Job, ...map[string]interface{}) (int64, error)); ok {
		return rf(ctx, executionID, job, extraAttrs...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, *Job, ...map[string]interface{}) int64); ok {
		r0 = rf(ctx, executionID, job, extraAttrs...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, *Job, ...map[string]interface{}) error); ok {
		r1 = rf(ctx, executionID, job, extraAttrs...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecutionIDsByVendorAndStatus provides a mock function with given fields: ctx, vendorType, status
func (_m *mockTaskManager) ExecutionIDsByVendorAndStatus(ctx context.Context, vendorType string, status string) ([]int64, error) {
	ret := _m.Called(ctx, vendorType, status)

	if len(ret) == 0 {
		panic("no return value specified for ExecutionIDsByVendorAndStatus")
	}

	var r0 []int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) ([]int64, error)); ok {
		return rf(ctx, vendorType, status)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []int64); ok {
		r0 = rf(ctx, vendorType, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, vendorType, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, id
func (_m *mockTaskManager) Get(ctx context.Context, id int64) (*Task, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*Task, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *Task); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Task)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLog provides a mock function with given fields: ctx, id
func (_m *mockTaskManager) GetLog(ctx context.Context, id int64) ([]byte, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetLog")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]byte, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []byte); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLogByJobID provides a mock function with given fields: ctx, jobID
func (_m *mockTaskManager) GetLogByJobID(ctx context.Context, jobID string) ([]byte, error) {
	ret := _m.Called(ctx, jobID)

	if len(ret) == 0 {
		panic("no return value specified for GetLogByJobID")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]byte, error)); ok {
		return rf(ctx, jobID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []byte); ok {
		r0 = rf(ctx, jobID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, jobID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsTaskFinished provides a mock function with given fields: ctx, reportID
func (_m *mockTaskManager) IsTaskFinished(ctx context.Context, reportID string) bool {
	ret := _m.Called(ctx, reportID)

	if len(ret) == 0 {
		panic("no return value specified for IsTaskFinished")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, reportID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// List provides a mock function with given fields: ctx, query
func (_m *mockTaskManager) List(ctx context.Context, query *q.Query) ([]*Task, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []*Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) ([]*Task, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) []*Task); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Task)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListScanTasksByReportUUID provides a mock function with given fields: ctx, uuid
func (_m *mockTaskManager) ListScanTasksByReportUUID(ctx context.Context, uuid string) ([]*Task, error) {
	ret := _m.Called(ctx, uuid)

	if len(ret) == 0 {
		panic("no return value specified for ListScanTasksByReportUUID")
	}

	var r0 []*Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*Task, error)); ok {
		return rf(ctx, uuid)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*Task); ok {
		r0 = rf(ctx, uuid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Task)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveStatusFromTask provides a mock function with given fields: ctx, reportID
func (_m *mockTaskManager) RetrieveStatusFromTask(ctx context.Context, reportID string) string {
	ret := _m.Called(ctx, reportID)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveStatusFromTask")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, reportID)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Stop provides a mock function with given fields: ctx, id
func (_m *mockTaskManager) Stop(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Stop")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, task, props
func (_m *mockTaskManager) Update(ctx context.Context, task *Task, props ...string) error {
	_va := make([]interface{}, len(props))
	for _i := range props {
		_va[_i] = props[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, task)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *Task, ...string) error); ok {
		r0 = rf(ctx, task, props...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateExtraAttrs provides a mock function with given fields: ctx, id, extraAttrs
func (_m *mockTaskManager) UpdateExtraAttrs(ctx context.Context, id int64, extraAttrs map[string]interface{}) error {
	ret := _m.Called(ctx, id, extraAttrs)

	if len(ret) == 0 {
		panic("no return value specified for UpdateExtraAttrs")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, map[string]interface{}) error); ok {
		r0 = rf(ctx, id, extraAttrs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateStatusInBatch provides a mock function with given fields: ctx, jobIDs, status, batchSize
func (_m *mockTaskManager) UpdateStatusInBatch(ctx context.Context, jobIDs []string, status string, batchSize int) error {
	ret := _m.Called(ctx, jobIDs, status, batchSize)

	if len(ret) == 0 {
		panic("no return value specified for UpdateStatusInBatch")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []string, string, int) error); ok {
		r0 = rf(ctx, jobIDs, status, batchSize)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// newMockTaskManager creates a new instance of mockTaskManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockTaskManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockTaskManager {
	mock := &mockTaskManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
