// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/dao.go

// Package repository is a generated GoMock package.
package repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/nht1206/pricetracker/internal/model"
)

// MockDAO is a mock of DAO interface.
type MockDAO struct {
	ctrl     *gomock.Controller
	recorder *MockDAOMockRecorder
}

// MockDAOMockRecorder is the mock recorder for MockDAO.
type MockDAOMockRecorder struct {
	mock *MockDAO
}

// NewMockDAO creates a new mock instance.
func NewMockDAO(ctrl *gomock.Controller) *MockDAO {
	mock := &MockDAO{ctrl: ctrl}
	mock.recorder = &MockDAOMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDAO) EXPECT() *MockDAOMockRecorder {
	return m.recorder
}

// FindTargetTrackingProduct mocks base method.
func (m *MockDAO) FindTargetTrackingProduct() ([]model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindTargetTrackingProduct")
	ret0, _ := ret[0].([]model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindTargetTrackingProduct indicates an expected call of FindTargetTrackingProduct.
func (mr *MockDAOMockRecorder) FindTargetTrackingProduct() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindTargetTrackingProduct", reflect.TypeOf((*MockDAO)(nil).FindTargetTrackingProduct))
}

// GetAllUserFollowed mocks base method.
func (m *MockDAO) GetAllUserFollowed(productId uint64) ([]model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUserFollowed", productId)
	ret0, _ := ret[0].([]model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUserFollowed indicates an expected call of GetAllUserFollowed.
func (mr *MockDAOMockRecorder) GetAllUserFollowed(productId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUserFollowed", reflect.TypeOf((*MockDAO)(nil).GetAllUserFollowed), productId)
}

// GetProductPrice mocks base method.
func (m *MockDAO) GetProductPrice(productId uint64) (*model.Price, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductPrice", productId)
	ret0, _ := ret[0].(*model.Price)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductPrice indicates an expected call of GetProductPrice.
func (mr *MockDAOMockRecorder) GetProductPrice(productId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductPrice", reflect.TypeOf((*MockDAO)(nil).GetProductPrice), productId)
}

// LockProductToTrackPrice mocks base method.
func (m *MockDAO) LockProductToTrackPrice(productId uint64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LockProductToTrackPrice", productId)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LockProductToTrackPrice indicates an expected call of LockProductToTrackPrice.
func (mr *MockDAOMockRecorder) LockProductToTrackPrice(productId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockProductToTrackPrice", reflect.TypeOf((*MockDAO)(nil).LockProductToTrackPrice), productId)
}

// UnlockProduct mocks base method.
func (m *MockDAO) UnlockProduct(productId uint64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnlockProduct", productId)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnlockProduct indicates an expected call of UnlockProduct.
func (mr *MockDAOMockRecorder) UnlockProduct(productId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnlockProduct", reflect.TypeOf((*MockDAO)(nil).UnlockProduct), productId)
}

// UpdateProductPrice mocks base method.
func (m *MockDAO) UpdateProductPrice(productId uint64, newPrice string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProductPrice", productId, newPrice)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProductPrice indicates an expected call of UpdateProductPrice.
func (mr *MockDAOMockRecorder) UpdateProductPrice(productId, newPrice interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProductPrice", reflect.TypeOf((*MockDAO)(nil).UpdateProductPrice), productId, newPrice)
}

// UpdateProductStatusToFailed mocks base method.
func (m *MockDAO) UpdateProductStatusToFailed(productId uint64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProductStatusToFailed", productId)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProductStatusToFailed indicates an expected call of UpdateProductStatusToFailed.
func (mr *MockDAOMockRecorder) UpdateProductStatusToFailed(productId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProductStatusToFailed", reflect.TypeOf((*MockDAO)(nil).UpdateProductStatusToFailed), productId)
}
