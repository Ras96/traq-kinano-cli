// Code generated by MockGen. DO NOT EDIT.
// Source: alias.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	ent "github.com/Ras96/traq-kinano-cli/ent"
	repository "github.com/Ras96/traq-kinano-cli/usecases/repository"
	gomock "github.com/golang/mock/gomock"
)

// MockAliasRepository is a mock of AliasRepository interface.
type MockAliasRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAliasRepositoryMockRecorder
}

// MockAliasRepositoryMockRecorder is the mock recorder for MockAliasRepository.
type MockAliasRepositoryMockRecorder struct {
	mock *MockAliasRepository
}

// NewMockAliasRepository creates a new mock instance.
func NewMockAliasRepository(ctrl *gomock.Controller) *MockAliasRepository {
	mock := &MockAliasRepository{ctrl: ctrl}
	mock.recorder = &MockAliasRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAliasRepository) EXPECT() *MockAliasRepositoryMockRecorder {
	return m.recorder
}

// AddAlias mocks base method.
func (m *MockAliasRepository) AddAlias(args *repository.AddAliasArgs) (*ent.Alias, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAlias", args)
	ret0, _ := ret[0].(*ent.Alias)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAlias indicates an expected call of AddAlias.
func (mr *MockAliasRepositoryMockRecorder) AddAlias(args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAlias", reflect.TypeOf((*MockAliasRepository)(nil).AddAlias), args)
}

// CallAlias mocks base method.
func (m *MockAliasRepository) CallAlias(short string) (*ent.Alias, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallAlias", short)
	ret0, _ := ret[0].(*ent.Alias)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallAlias indicates an expected call of CallAlias.
func (mr *MockAliasRepositoryMockRecorder) CallAlias(short interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallAlias", reflect.TypeOf((*MockAliasRepository)(nil).CallAlias), short)
}
