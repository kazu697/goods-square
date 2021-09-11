// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/domain/repository/product.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	model "goods-square/internal/domain/model"
	repository "goods-square/internal/domain/repository"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "github.com/jinzhu/gorm"
)

// MockProduct is a mock of Product interface.
type MockProduct struct {
	ctrl     *gomock.Controller
	recorder *MockProductMockRecorder
}

// MockProductMockRecorder is the mock recorder for MockProduct.
type MockProductMockRecorder struct {
	mock *MockProduct
}

// NewMockProduct creates a new mock instance.
func NewMockProduct(ctrl *gomock.Controller) *MockProduct {
	mock := &MockProduct{ctrl: ctrl}
	mock.recorder = &MockProductMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProduct) EXPECT() *MockProductMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockProduct) Get(ctx context.Context, db *gorm.DB, query *repository.ProductGetQuery) (*model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, db, query)
	ret0, _ := ret[0].(*model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockProductMockRecorder) Get(ctx, db, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockProduct)(nil).Get), ctx, db, query)
}

// List mocks base method.
func (m *MockProduct) List(ctx context.Context, db *gorm.DB, query *repository.ProductListQuery) (model.Products, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, db, query)
	ret0, _ := ret[0].(model.Products)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockProductMockRecorder) List(ctx, db, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockProduct)(nil).List), ctx, db, query)
}
