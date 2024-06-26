// Code generated by MockGen. DO NOT EDIT.
// Source: crypto.go
//
// Generated by this command:
//
//	mockgen -source crypto.go -destination ./crypto_mock.go -package crypto
//

// Package crypto is a generated GoMock package.
package crypto

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockEncryptionAlgorithm is a mock of EncryptionAlgorithm interface.
type MockEncryptionAlgorithm struct {
	ctrl     *gomock.Controller
	recorder *MockEncryptionAlgorithmMockRecorder
}

// MockEncryptionAlgorithmMockRecorder is the mock recorder for MockEncryptionAlgorithm.
type MockEncryptionAlgorithmMockRecorder struct {
	mock *MockEncryptionAlgorithm
}

// NewMockEncryptionAlgorithm creates a new mock instance.
func NewMockEncryptionAlgorithm(ctrl *gomock.Controller) *MockEncryptionAlgorithm {
	mock := &MockEncryptionAlgorithm{ctrl: ctrl}
	mock.recorder = &MockEncryptionAlgorithmMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEncryptionAlgorithm) EXPECT() *MockEncryptionAlgorithmMockRecorder {
	return m.recorder
}

// Algorithm mocks base method.
func (m *MockEncryptionAlgorithm) Algorithm() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Algorithm")
	ret0, _ := ret[0].(string)
	return ret0
}

// Algorithm indicates an expected call of Algorithm.
func (mr *MockEncryptionAlgorithmMockRecorder) Algorithm() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Algorithm", reflect.TypeOf((*MockEncryptionAlgorithm)(nil).Algorithm))
}

// Decrypt mocks base method.
func (m *MockEncryptionAlgorithm) Decrypt(hashed []byte, keyID string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decrypt", hashed, keyID)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decrypt indicates an expected call of Decrypt.
func (mr *MockEncryptionAlgorithmMockRecorder) Decrypt(hashed, keyID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decrypt", reflect.TypeOf((*MockEncryptionAlgorithm)(nil).Decrypt), hashed, keyID)
}

// DecryptString mocks base method.
func (m *MockEncryptionAlgorithm) DecryptString(hashed []byte, keyID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecryptString", hashed, keyID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecryptString indicates an expected call of DecryptString.
func (mr *MockEncryptionAlgorithmMockRecorder) DecryptString(hashed, keyID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecryptString", reflect.TypeOf((*MockEncryptionAlgorithm)(nil).DecryptString), hashed, keyID)
}

// DecryptionKeyIDs mocks base method.
func (m *MockEncryptionAlgorithm) DecryptionKeyIDs() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecryptionKeyIDs")
	ret0, _ := ret[0].([]string)
	return ret0
}

// DecryptionKeyIDs indicates an expected call of DecryptionKeyIDs.
func (mr *MockEncryptionAlgorithmMockRecorder) DecryptionKeyIDs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecryptionKeyIDs", reflect.TypeOf((*MockEncryptionAlgorithm)(nil).DecryptionKeyIDs))
}

// Encrypt mocks base method.
func (m *MockEncryptionAlgorithm) Encrypt(value []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encrypt", value)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encrypt indicates an expected call of Encrypt.
func (mr *MockEncryptionAlgorithmMockRecorder) Encrypt(value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encrypt", reflect.TypeOf((*MockEncryptionAlgorithm)(nil).Encrypt), value)
}

// EncryptionKeyID mocks base method.
func (m *MockEncryptionAlgorithm) EncryptionKeyID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EncryptionKeyID")
	ret0, _ := ret[0].(string)
	return ret0
}

// EncryptionKeyID indicates an expected call of EncryptionKeyID.
func (mr *MockEncryptionAlgorithmMockRecorder) EncryptionKeyID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EncryptionKeyID", reflect.TypeOf((*MockEncryptionAlgorithm)(nil).EncryptionKeyID))
}
