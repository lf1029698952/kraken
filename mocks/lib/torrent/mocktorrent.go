// Code generated by MockGen. DO NOT EDIT.
// Source: code.uber.internal/infra/kraken/lib/torrent (interfaces: Client)

// Package mocktorrent is a generated GoMock package.
package mocktorrent

import (
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockClient) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockClientMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClient)(nil).Close))
}

// DownloadTorrent mocks base method
func (m *MockClient) DownloadTorrent(arg0 string) (io.ReadCloser, error) {
	ret := m.ctrl.Call(m, "DownloadTorrent", arg0)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadTorrent indicates an expected call of DownloadTorrent
func (mr *MockClientMockRecorder) DownloadTorrent(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadTorrent", reflect.TypeOf((*MockClient)(nil).DownloadTorrent), arg0)
}

// GetManifest mocks base method
func (m *MockClient) GetManifest(arg0, arg1 string) (io.ReadCloser, error) {
	ret := m.ctrl.Call(m, "GetManifest", arg0, arg1)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetManifest indicates an expected call of GetManifest
func (mr *MockClientMockRecorder) GetManifest(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetManifest", reflect.TypeOf((*MockClient)(nil).GetManifest), arg0, arg1)
}

// PostManifest mocks base method
func (m *MockClient) PostManifest(arg0, arg1, arg2 string, arg3 io.Reader) error {
	ret := m.ctrl.Call(m, "PostManifest", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostManifest indicates an expected call of PostManifest
func (mr *MockClientMockRecorder) PostManifest(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostManifest", reflect.TypeOf((*MockClient)(nil).PostManifest), arg0, arg1, arg2, arg3)
}