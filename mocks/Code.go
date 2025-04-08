// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/huabtc/maroto/v2/pkg/core/entity"
	mock "github.com/stretchr/testify/mock"

	props "github.com/huabtc/maroto/v2/pkg/props"
)

// Code is an autogenerated mock type for the Code type
type Code struct {
	mock.Mock
}

type Code_Expecter struct {
	mock *mock.Mock
}

func (_m *Code) EXPECT() *Code_Expecter {
	return &Code_Expecter{mock: &_m.Mock}
}

// GenBar provides a mock function with given fields: code, cell, prop
func (_m *Code) GenBar(code string, cell *entity.Cell, prop *props.Barcode) (*entity.Image, error) {
	ret := _m.Called(code, cell, prop)

	if len(ret) == 0 {
		panic("no return value specified for GenBar")
	}

	var r0 *entity.Image
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *entity.Cell, *props.Barcode) (*entity.Image, error)); ok {
		return rf(code, cell, prop)
	}
	if rf, ok := ret.Get(0).(func(string, *entity.Cell, *props.Barcode) *entity.Image); ok {
		r0 = rf(code, cell, prop)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Image)
		}
	}

	if rf, ok := ret.Get(1).(func(string, *entity.Cell, *props.Barcode) error); ok {
		r1 = rf(code, cell, prop)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Code_GenBar_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenBar'
type Code_GenBar_Call struct {
	*mock.Call
}

// GenBar is a helper method to define mock.On call
//   - code string
//   - cell *entity.Cell
//   - prop *props.Barcode
func (_e *Code_Expecter) GenBar(code interface{}, cell interface{}, prop interface{}) *Code_GenBar_Call {
	return &Code_GenBar_Call{Call: _e.mock.On("GenBar", code, cell, prop)}
}

func (_c *Code_GenBar_Call) Run(run func(code string, cell *entity.Cell, prop *props.Barcode)) *Code_GenBar_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*entity.Cell), args[2].(*props.Barcode))
	})
	return _c
}

func (_c *Code_GenBar_Call) Return(_a0 *entity.Image, _a1 error) *Code_GenBar_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Code_GenBar_Call) RunAndReturn(run func(string, *entity.Cell, *props.Barcode) (*entity.Image, error)) *Code_GenBar_Call {
	_c.Call.Return(run)
	return _c
}

// GenDataMatrix provides a mock function with given fields: code
func (_m *Code) GenDataMatrix(code string) (*entity.Image, error) {
	ret := _m.Called(code)

	if len(ret) == 0 {
		panic("no return value specified for GenDataMatrix")
	}

	var r0 *entity.Image
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entity.Image, error)); ok {
		return rf(code)
	}
	if rf, ok := ret.Get(0).(func(string) *entity.Image); ok {
		r0 = rf(code)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Image)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Code_GenDataMatrix_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenDataMatrix'
type Code_GenDataMatrix_Call struct {
	*mock.Call
}

// GenDataMatrix is a helper method to define mock.On call
//   - code string
func (_e *Code_Expecter) GenDataMatrix(code interface{}) *Code_GenDataMatrix_Call {
	return &Code_GenDataMatrix_Call{Call: _e.mock.On("GenDataMatrix", code)}
}

func (_c *Code_GenDataMatrix_Call) Run(run func(code string)) *Code_GenDataMatrix_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Code_GenDataMatrix_Call) Return(_a0 *entity.Image, _a1 error) *Code_GenDataMatrix_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Code_GenDataMatrix_Call) RunAndReturn(run func(string) (*entity.Image, error)) *Code_GenDataMatrix_Call {
	_c.Call.Return(run)
	return _c
}

// GenQr provides a mock function with given fields: code
func (_m *Code) GenQr(code string) (*entity.Image, error) {
	ret := _m.Called(code)

	if len(ret) == 0 {
		panic("no return value specified for GenQr")
	}

	var r0 *entity.Image
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entity.Image, error)); ok {
		return rf(code)
	}
	if rf, ok := ret.Get(0).(func(string) *entity.Image); ok {
		r0 = rf(code)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Image)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Code_GenQr_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenQr'
type Code_GenQr_Call struct {
	*mock.Call
}

// GenQr is a helper method to define mock.On call
//   - code string
func (_e *Code_Expecter) GenQr(code interface{}) *Code_GenQr_Call {
	return &Code_GenQr_Call{Call: _e.mock.On("GenQr", code)}
}

func (_c *Code_GenQr_Call) Run(run func(code string)) *Code_GenQr_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Code_GenQr_Call) Return(_a0 *entity.Image, _a1 error) *Code_GenQr_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Code_GenQr_Call) RunAndReturn(run func(string) (*entity.Image, error)) *Code_GenQr_Call {
	_c.Call.Return(run)
	return _c
}

// NewCode creates a new instance of Code. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCode(t interface {
	mock.TestingT
	Cleanup(func())
},
) *Code {
	mock := &Code{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
