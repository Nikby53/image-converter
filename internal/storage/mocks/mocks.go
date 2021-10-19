package mocks

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStorageInterface is a mock of StorageInterface interface.
type MockStorageInterface struct {
	ctrl     *gomock.Controller
	recorder *MockStorageInterfaceMockRecorder
}

// MockStorageInterfaceMockRecorder is the mock recorder for MockStorageInterface.
type MockStorageInterfaceMockRecorder struct {
	mock *MockStorageInterface
}

// NewMockStorageInterface creates a new mock instance.
func NewMockStorageInterface(ctrl *gomock.Controller) *MockStorageInterface {
	mock := &MockStorageInterface{ctrl: ctrl}
	mock.recorder = &MockStorageInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageInterface) EXPECT() *MockStorageInterfaceMockRecorder {
	return m.recorder
}

// DownloadFile mocks base method.
func (m *MockStorageInterface) DownloadFile(fileID string) (io.ReadSeeker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadFile", fileID)
	ret0, _ := ret[0].(io.ReadSeeker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadFile indicates an expected call of DownloadFile.
func (mr *MockStorageInterfaceMockRecorder) DownloadFile(fileID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadFile", reflect.TypeOf((*MockStorageInterface)(nil).DownloadFile), fileID)
}

// DownloadImageFromID mocks base method.
func (m *MockStorageInterface) DownloadImageFromID(fileID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadImageFromID", fileID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadImageFromID indicates an expected call of DownloadImageFromID.
func (mr *MockStorageInterfaceMockRecorder) DownloadImageFromID(fileID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadImageFromID", reflect.TypeOf((*MockStorageInterface)(nil).DownloadImageFromID), fileID)
}

// UploadFile mocks base method.
func (m *MockStorageInterface) UploadFile(image io.ReadSeeker, fileID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFile", image, fileID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadFile indicates an expected call of UploadFile.
func (mr *MockStorageInterfaceMockRecorder) UploadFile(image, fileID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFile", reflect.TypeOf((*MockStorageInterface)(nil).UploadFile), image, fileID)
}

// UploadTargetFile mocks base method.
func (m *MockStorageInterface) UploadTargetFile(filename, fileID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadTargetFile", filename, fileID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadTargetFile indicates an expected call of UploadTargetFile.
func (mr *MockStorageInterfaceMockRecorder) UploadTargetFile(filename, fileID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadTargetFile", reflect.TypeOf((*MockStorageInterface)(nil).UploadTargetFile), filename, fileID)
}
