// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	types "github.com/cosmos/cosmos-sdk/types"
	mock "github.com/stretchr/testify/mock"
)

// BankKeeper is an autogenerated mock type for the BankKeeper type
type BankKeeper struct {
	mock.Mock
}

type BankKeeper_Expecter struct {
	mock *mock.Mock
}

func (_m *BankKeeper) EXPECT() *BankKeeper_Expecter {
	return &BankKeeper_Expecter{mock: &_m.Mock}
}

// BurnCoins provides a mock function with given fields: ctx, name, amt
func (_m *BankKeeper) BurnCoins(ctx types.Context, name string, amt types.Coins) error {
	ret := _m.Called(ctx, name, amt)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, string, types.Coins) error); ok {
		r0 = rf(ctx, name, amt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BankKeeper_BurnCoins_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BurnCoins'
type BankKeeper_BurnCoins_Call struct {
	*mock.Call
}

// BurnCoins is a helper method to define mock.On call
//   - ctx types.Context
//   - name string
//   - amt types.Coins
func (_e *BankKeeper_Expecter) BurnCoins(ctx interface{}, name interface{}, amt interface{}) *BankKeeper_BurnCoins_Call {
	return &BankKeeper_BurnCoins_Call{Call: _e.mock.On("BurnCoins", ctx, name, amt)}
}

func (_c *BankKeeper_BurnCoins_Call) Run(run func(ctx types.Context, name string, amt types.Coins)) *BankKeeper_BurnCoins_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(string), args[2].(types.Coins))
	})
	return _c
}

func (_c *BankKeeper_BurnCoins_Call) Return(_a0 error) *BankKeeper_BurnCoins_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BankKeeper_BurnCoins_Call) RunAndReturn(run func(types.Context, string, types.Coins) error) *BankKeeper_BurnCoins_Call {
	_c.Call.Return(run)
	return _c
}

// GetBalance provides a mock function with given fields: ctx, addr, denom
func (_m *BankKeeper) GetBalance(ctx types.Context, addr types.AccAddress, denom string) types.Coin {
	ret := _m.Called(ctx, addr, denom)

	var r0 types.Coin
	if rf, ok := ret.Get(0).(func(types.Context, types.AccAddress, string) types.Coin); ok {
		r0 = rf(ctx, addr, denom)
	} else {
		r0 = ret.Get(0).(types.Coin)
	}

	return r0
}

// BankKeeper_GetBalance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBalance'
type BankKeeper_GetBalance_Call struct {
	*mock.Call
}

// GetBalance is a helper method to define mock.On call
//   - ctx types.Context
//   - addr types.AccAddress
//   - denom string
func (_e *BankKeeper_Expecter) GetBalance(ctx interface{}, addr interface{}, denom interface{}) *BankKeeper_GetBalance_Call {
	return &BankKeeper_GetBalance_Call{Call: _e.mock.On("GetBalance", ctx, addr, denom)}
}

func (_c *BankKeeper_GetBalance_Call) Run(run func(ctx types.Context, addr types.AccAddress, denom string)) *BankKeeper_GetBalance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(types.AccAddress), args[2].(string))
	})
	return _c
}

func (_c *BankKeeper_GetBalance_Call) Return(_a0 types.Coin) *BankKeeper_GetBalance_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BankKeeper_GetBalance_Call) RunAndReturn(run func(types.Context, types.AccAddress, string) types.Coin) *BankKeeper_GetBalance_Call {
	_c.Call.Return(run)
	return _c
}

// HasBalance provides a mock function with given fields: ctx, addr, amt
func (_m *BankKeeper) HasBalance(ctx types.Context, addr types.AccAddress, amt types.Coin) bool {
	ret := _m.Called(ctx, addr, amt)

	var r0 bool
	if rf, ok := ret.Get(0).(func(types.Context, types.AccAddress, types.Coin) bool); ok {
		r0 = rf(ctx, addr, amt)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// BankKeeper_HasBalance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HasBalance'
type BankKeeper_HasBalance_Call struct {
	*mock.Call
}

// HasBalance is a helper method to define mock.On call
//   - ctx types.Context
//   - addr types.AccAddress
//   - amt types.Coin
func (_e *BankKeeper_Expecter) HasBalance(ctx interface{}, addr interface{}, amt interface{}) *BankKeeper_HasBalance_Call {
	return &BankKeeper_HasBalance_Call{Call: _e.mock.On("HasBalance", ctx, addr, amt)}
}

func (_c *BankKeeper_HasBalance_Call) Run(run func(ctx types.Context, addr types.AccAddress, amt types.Coin)) *BankKeeper_HasBalance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(types.AccAddress), args[2].(types.Coin))
	})
	return _c
}

func (_c *BankKeeper_HasBalance_Call) Return(_a0 bool) *BankKeeper_HasBalance_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BankKeeper_HasBalance_Call) RunAndReturn(run func(types.Context, types.AccAddress, types.Coin) bool) *BankKeeper_HasBalance_Call {
	_c.Call.Return(run)
	return _c
}

// MintCoins provides a mock function with given fields: ctx, name, amt
func (_m *BankKeeper) MintCoins(ctx types.Context, name string, amt types.Coins) error {
	ret := _m.Called(ctx, name, amt)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, string, types.Coins) error); ok {
		r0 = rf(ctx, name, amt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BankKeeper_MintCoins_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MintCoins'
type BankKeeper_MintCoins_Call struct {
	*mock.Call
}

// MintCoins is a helper method to define mock.On call
//   - ctx types.Context
//   - name string
//   - amt types.Coins
func (_e *BankKeeper_Expecter) MintCoins(ctx interface{}, name interface{}, amt interface{}) *BankKeeper_MintCoins_Call {
	return &BankKeeper_MintCoins_Call{Call: _e.mock.On("MintCoins", ctx, name, amt)}
}

func (_c *BankKeeper_MintCoins_Call) Run(run func(ctx types.Context, name string, amt types.Coins)) *BankKeeper_MintCoins_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(string), args[2].(types.Coins))
	})
	return _c
}

func (_c *BankKeeper_MintCoins_Call) Return(_a0 error) *BankKeeper_MintCoins_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BankKeeper_MintCoins_Call) RunAndReturn(run func(types.Context, string, types.Coins) error) *BankKeeper_MintCoins_Call {
	_c.Call.Return(run)
	return _c
}

// SendCoins provides a mock function with given fields: ctx, fromAddr, toAddr, amt
func (_m *BankKeeper) SendCoins(ctx types.Context, fromAddr types.AccAddress, toAddr types.AccAddress, amt types.Coins) error {
	ret := _m.Called(ctx, fromAddr, toAddr, amt)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, types.AccAddress, types.AccAddress, types.Coins) error); ok {
		r0 = rf(ctx, fromAddr, toAddr, amt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BankKeeper_SendCoins_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendCoins'
type BankKeeper_SendCoins_Call struct {
	*mock.Call
}

// SendCoins is a helper method to define mock.On call
//   - ctx types.Context
//   - fromAddr types.AccAddress
//   - toAddr types.AccAddress
//   - amt types.Coins
func (_e *BankKeeper_Expecter) SendCoins(ctx interface{}, fromAddr interface{}, toAddr interface{}, amt interface{}) *BankKeeper_SendCoins_Call {
	return &BankKeeper_SendCoins_Call{Call: _e.mock.On("SendCoins", ctx, fromAddr, toAddr, amt)}
}

func (_c *BankKeeper_SendCoins_Call) Run(run func(ctx types.Context, fromAddr types.AccAddress, toAddr types.AccAddress, amt types.Coins)) *BankKeeper_SendCoins_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(types.AccAddress), args[2].(types.AccAddress), args[3].(types.Coins))
	})
	return _c
}

func (_c *BankKeeper_SendCoins_Call) Return(_a0 error) *BankKeeper_SendCoins_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BankKeeper_SendCoins_Call) RunAndReturn(run func(types.Context, types.AccAddress, types.AccAddress, types.Coins) error) *BankKeeper_SendCoins_Call {
	_c.Call.Return(run)
	return _c
}

// SendCoinsFromAccountToModule provides a mock function with given fields: ctx, senderAddr, recipientModule, amt
func (_m *BankKeeper) SendCoinsFromAccountToModule(ctx types.Context, senderAddr types.AccAddress, recipientModule string, amt types.Coins) error {
	ret := _m.Called(ctx, senderAddr, recipientModule, amt)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, types.AccAddress, string, types.Coins) error); ok {
		r0 = rf(ctx, senderAddr, recipientModule, amt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BankKeeper_SendCoinsFromAccountToModule_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendCoinsFromAccountToModule'
type BankKeeper_SendCoinsFromAccountToModule_Call struct {
	*mock.Call
}

// SendCoinsFromAccountToModule is a helper method to define mock.On call
//   - ctx types.Context
//   - senderAddr types.AccAddress
//   - recipientModule string
//   - amt types.Coins
func (_e *BankKeeper_Expecter) SendCoinsFromAccountToModule(ctx interface{}, senderAddr interface{}, recipientModule interface{}, amt interface{}) *BankKeeper_SendCoinsFromAccountToModule_Call {
	return &BankKeeper_SendCoinsFromAccountToModule_Call{Call: _e.mock.On("SendCoinsFromAccountToModule", ctx, senderAddr, recipientModule, amt)}
}

func (_c *BankKeeper_SendCoinsFromAccountToModule_Call) Run(run func(ctx types.Context, senderAddr types.AccAddress, recipientModule string, amt types.Coins)) *BankKeeper_SendCoinsFromAccountToModule_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(types.AccAddress), args[2].(string), args[3].(types.Coins))
	})
	return _c
}

func (_c *BankKeeper_SendCoinsFromAccountToModule_Call) Return(_a0 error) *BankKeeper_SendCoinsFromAccountToModule_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BankKeeper_SendCoinsFromAccountToModule_Call) RunAndReturn(run func(types.Context, types.AccAddress, string, types.Coins) error) *BankKeeper_SendCoinsFromAccountToModule_Call {
	_c.Call.Return(run)
	return _c
}

// SendCoinsFromModuleToAccount provides a mock function with given fields: ctx, senderModule, recipientAddr, amt
func (_m *BankKeeper) SendCoinsFromModuleToAccount(ctx types.Context, senderModule string, recipientAddr types.AccAddress, amt types.Coins) error {
	ret := _m.Called(ctx, senderModule, recipientAddr, amt)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, string, types.AccAddress, types.Coins) error); ok {
		r0 = rf(ctx, senderModule, recipientAddr, amt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BankKeeper_SendCoinsFromModuleToAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendCoinsFromModuleToAccount'
type BankKeeper_SendCoinsFromModuleToAccount_Call struct {
	*mock.Call
}

// SendCoinsFromModuleToAccount is a helper method to define mock.On call
//   - ctx types.Context
//   - senderModule string
//   - recipientAddr types.AccAddress
//   - amt types.Coins
func (_e *BankKeeper_Expecter) SendCoinsFromModuleToAccount(ctx interface{}, senderModule interface{}, recipientAddr interface{}, amt interface{}) *BankKeeper_SendCoinsFromModuleToAccount_Call {
	return &BankKeeper_SendCoinsFromModuleToAccount_Call{Call: _e.mock.On("SendCoinsFromModuleToAccount", ctx, senderModule, recipientAddr, amt)}
}

func (_c *BankKeeper_SendCoinsFromModuleToAccount_Call) Run(run func(ctx types.Context, senderModule string, recipientAddr types.AccAddress, amt types.Coins)) *BankKeeper_SendCoinsFromModuleToAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(string), args[2].(types.AccAddress), args[3].(types.Coins))
	})
	return _c
}

func (_c *BankKeeper_SendCoinsFromModuleToAccount_Call) Return(_a0 error) *BankKeeper_SendCoinsFromModuleToAccount_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BankKeeper_SendCoinsFromModuleToAccount_Call) RunAndReturn(run func(types.Context, string, types.AccAddress, types.Coins) error) *BankKeeper_SendCoinsFromModuleToAccount_Call {
	_c.Call.Return(run)
	return _c
}

// NewBankKeeper creates a new instance of BankKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBankKeeper(t interface {
	mock.TestingT
	Cleanup(func())
}) *BankKeeper {
	mock := &BankKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
