// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// NodeClassApplier is an autogenerated mock type for the NodeClassApplier type
type NodeClassApplier struct {
	mock.Mock
}

type NodeClassApplier_Expecter struct {
	mock *mock.Mock
}

func (_m *NodeClassApplier) EXPECT() *NodeClassApplier_Expecter {
	return &NodeClassApplier_Expecter{mock: &_m.Mock}
}

// PatchSubnets provides a mock function with given fields: ctx, subnetIDs
func (_m *NodeClassApplier) PatchSubnets(ctx context.Context, subnetIDs []string) error {
	ret := _m.Called(ctx, subnetIDs)

	if len(ret) == 0 {
		panic("no return value specified for PatchSubnets")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) error); ok {
		r0 = rf(ctx, subnetIDs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NodeClassApplier_PatchSubnets_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PatchSubnets'
type NodeClassApplier_PatchSubnets_Call struct {
	*mock.Call
}

// PatchSubnets is a helper method to define mock.On call
//   - ctx context.Context
//   - subnetIDs []string
func (_e *NodeClassApplier_Expecter) PatchSubnets(ctx interface{}, subnetIDs interface{}) *NodeClassApplier_PatchSubnets_Call {
	return &NodeClassApplier_PatchSubnets_Call{Call: _e.mock.On("PatchSubnets", ctx, subnetIDs)}
}

func (_c *NodeClassApplier_PatchSubnets_Call) Run(run func(ctx context.Context, subnetIDs []string)) *NodeClassApplier_PatchSubnets_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string))
	})
	return _c
}

func (_c *NodeClassApplier_PatchSubnets_Call) Return(_a0 error) *NodeClassApplier_PatchSubnets_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *NodeClassApplier_PatchSubnets_Call) RunAndReturn(run func(context.Context, []string) error) *NodeClassApplier_PatchSubnets_Call {
	_c.Call.Return(run)
	return _c
}

// NewNodeClassApplier creates a new instance of NodeClassApplier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNodeClassApplier(t interface {
	mock.TestingT
	Cleanup(func())
}) *NodeClassApplier {
	mock := &NodeClassApplier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
