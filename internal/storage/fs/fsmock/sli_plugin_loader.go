// Code generated by mockery v2.53.3. DO NOT EDIT.

package fsmock

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	sli "github.com/slok/sloth/internal/plugin/sli"
)

// SLIPluginLoader is an autogenerated mock type for the SLIPluginLoader type
type SLIPluginLoader struct {
	mock.Mock
}

// LoadRawSLIPlugin provides a mock function with given fields: ctx, src
func (_m *SLIPluginLoader) LoadRawSLIPlugin(ctx context.Context, src string) (*sli.SLIPlugin, error) {
	ret := _m.Called(ctx, src)

	if len(ret) == 0 {
		panic("no return value specified for LoadRawSLIPlugin")
	}

	var r0 *sli.SLIPlugin
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*sli.SLIPlugin, error)); ok {
		return rf(ctx, src)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *sli.SLIPlugin); ok {
		r0 = rf(ctx, src)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sli.SLIPlugin)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, src)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSLIPluginLoader creates a new instance of SLIPluginLoader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSLIPluginLoader(t interface {
	mock.TestingT
	Cleanup(func())
}) *SLIPluginLoader {
	mock := &SLIPluginLoader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
