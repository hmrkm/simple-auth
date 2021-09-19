// Code generated by MockGen. DO NOT EDIT.
// Source: gorm_conn.go

// Package io is a generated GoMock package.
package io

import (
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockDBConn is a mock of DBConn interface.
type MockDBConn struct {
	ctrl     *gomock.Controller
	recorder *MockDBConnMockRecorder
}

// MockDBConnMockRecorder is the mock recorder for MockDBConn.
type MockDBConnMockRecorder struct {
	mock *MockDBConn
}

// NewMockDBConn creates a new mock instance.
func NewMockDBConn(ctrl *gomock.Controller) *MockDBConn {
	mock := &MockDBConn{ctrl: ctrl}
	mock.recorder = &MockDBConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBConn) EXPECT() *MockDBConnMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockDBConn) Create(value interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", value)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockDBConnMockRecorder) Create(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDBConn)(nil).Create), value)
}

// DB mocks base method.
func (m *MockDBConn) DB() (*sql.DB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DB")
	ret0, _ := ret[0].(*sql.DB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DB indicates an expected call of DB.
func (mr *MockDBConnMockRecorder) DB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DB", reflect.TypeOf((*MockDBConn)(nil).DB))
}

// Find mocks base method.
func (m *MockDBConn) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{dest}
	for _, a := range conds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Find", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Find indicates an expected call of Find.
func (mr *MockDBConnMockRecorder) Find(dest interface{}, conds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dest}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockDBConn)(nil).Find), varargs...)
}

// First mocks base method.
func (m *MockDBConn) First(dest interface{}, conds ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{dest}
	for _, a := range conds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "First", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// First indicates an expected call of First.
func (mr *MockDBConnMockRecorder) First(dest interface{}, conds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dest}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "First", reflect.TypeOf((*MockDBConn)(nil).First), varargs...)
}
