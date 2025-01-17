// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by mockery v2.30.16. DO NOT EDIT.
// Modified manually to add new method.
package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	client "go.temporal.io/sdk/client"
)

// ScheduleHandle is an autogenerated mock type for the ScheduleHandle type
type ScheduleHandle struct {
	mock.Mock
}

// Backfill provides a mock function with given fields: ctx, options
func (_m *ScheduleHandle) Backfill(ctx context.Context, options client.ScheduleBackfillOptions) error {
	ret := _m.Called(ctx, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, client.ScheduleBackfillOptions) error); ok {
		r0 = rf(ctx, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx
func (_m *ScheduleHandle) Delete(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Describe provides a mock function with given fields: ctx
func (_m *ScheduleHandle) Describe(ctx context.Context) (*client.ScheduleDescription, error) {
	ret := _m.Called(ctx)

	var r0 *client.ScheduleDescription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*client.ScheduleDescription, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *client.ScheduleDescription); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ScheduleDescription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetID provides a mock function with given fields:
func (_m *ScheduleHandle) GetID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Pause provides a mock function with given fields: ctx, options
func (_m *ScheduleHandle) Pause(ctx context.Context, options client.SchedulePauseOptions) error {
	ret := _m.Called(ctx, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, client.SchedulePauseOptions) error); ok {
		r0 = rf(ctx, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Trigger provides a mock function with given fields: ctx, options
func (_m *ScheduleHandle) Trigger(ctx context.Context, options client.ScheduleTriggerOptions) error {
	ret := _m.Called(ctx, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, client.ScheduleTriggerOptions) error); ok {
		r0 = rf(ctx, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Unpause provides a mock function with given fields: ctx, options
func (_m *ScheduleHandle) Unpause(ctx context.Context, options client.ScheduleUnpauseOptions) error {
	ret := _m.Called(ctx, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, client.ScheduleUnpauseOptions) error); ok {
		r0 = rf(ctx, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, options
func (_m *ScheduleHandle) Update(ctx context.Context, options client.ScheduleUpdateOptions) error {
	ret := _m.Called(ctx, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, client.ScheduleUpdateOptions) error); ok {
		r0 = rf(ctx, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewScheduleHandle creates a new instance of ScheduleHandle. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewScheduleHandle(t interface {
	mock.TestingT
	Cleanup(func())
}) *ScheduleHandle {
	mock := &ScheduleHandle{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
