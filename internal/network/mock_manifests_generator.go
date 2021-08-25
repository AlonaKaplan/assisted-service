// Code generated by MockGen. DO NOT EDIT.
// Source: manifests_generator.go

// Package network is a generated GoMock package.
package network

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	common "github.com/openshift/assisted-service/internal/common"
	logrus "github.com/sirupsen/logrus"
)

// MockManifestsGeneratorAPI is a mock of ManifestsGeneratorAPI interface.
type MockManifestsGeneratorAPI struct {
	ctrl     *gomock.Controller
	recorder *MockManifestsGeneratorAPIMockRecorder
}

// MockManifestsGeneratorAPIMockRecorder is the mock recorder for MockManifestsGeneratorAPI.
type MockManifestsGeneratorAPIMockRecorder struct {
	mock *MockManifestsGeneratorAPI
}

// NewMockManifestsGeneratorAPI creates a new mock instance.
func NewMockManifestsGeneratorAPI(ctrl *gomock.Controller) *MockManifestsGeneratorAPI {
	mock := &MockManifestsGeneratorAPI{ctrl: ctrl}
	mock.recorder = &MockManifestsGeneratorAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManifestsGeneratorAPI) EXPECT() *MockManifestsGeneratorAPIMockRecorder {
	return m.recorder
}

// AddChronyManifest mocks base method.
func (m *MockManifestsGeneratorAPI) AddChronyManifest(ctx context.Context, log logrus.FieldLogger, c *common.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddChronyManifest", ctx, log, c)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddChronyManifest indicates an expected call of AddChronyManifest.
func (mr *MockManifestsGeneratorAPIMockRecorder) AddChronyManifest(ctx, log, c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddChronyManifest", reflect.TypeOf((*MockManifestsGeneratorAPI)(nil).AddChronyManifest), ctx, log, c)
}

// AddDiskEncryptionManifest mocks base method.
func (m *MockManifestsGeneratorAPI) AddDiskEncryptionManifest(ctx context.Context, log logrus.FieldLogger, c *common.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDiskEncryptionManifest", ctx, log, c)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddDiskEncryptionManifest indicates an expected call of AddDiskEncryptionManifest.
func (mr *MockManifestsGeneratorAPIMockRecorder) AddDiskEncryptionManifest(ctx, log, c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDiskEncryptionManifest", reflect.TypeOf((*MockManifestsGeneratorAPI)(nil).AddDiskEncryptionManifest), ctx, log, c)
}

// AddDnsmasqForSingleNode mocks base method.
func (m *MockManifestsGeneratorAPI) AddDnsmasqForSingleNode(ctx context.Context, log logrus.FieldLogger, c *common.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDnsmasqForSingleNode", ctx, log, c)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddDnsmasqForSingleNode indicates an expected call of AddDnsmasqForSingleNode.
func (mr *MockManifestsGeneratorAPIMockRecorder) AddDnsmasqForSingleNode(ctx, log, c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDnsmasqForSingleNode", reflect.TypeOf((*MockManifestsGeneratorAPI)(nil).AddDnsmasqForSingleNode), ctx, log, c)
}

// AddSchedulableMastersManifest mocks base method.
func (m *MockManifestsGeneratorAPI) AddSchedulableMastersManifest(ctx context.Context, log logrus.FieldLogger, c *common.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSchedulableMastersManifest", ctx, log, c)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSchedulableMastersManifest indicates an expected call of AddSchedulableMastersManifest.
func (mr *MockManifestsGeneratorAPIMockRecorder) AddSchedulableMastersManifest(ctx, log, c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSchedulableMastersManifest", reflect.TypeOf((*MockManifestsGeneratorAPI)(nil).AddSchedulableMastersManifest), ctx, log, c)
}

// AddTelemeterManifest mocks base method.
func (m *MockManifestsGeneratorAPI) AddTelemeterManifest(ctx context.Context, log logrus.FieldLogger, c *common.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTelemeterManifest", ctx, log, c)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTelemeterManifest indicates an expected call of AddTelemeterManifest.
func (mr *MockManifestsGeneratorAPIMockRecorder) AddTelemeterManifest(ctx, log, c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTelemeterManifest", reflect.TypeOf((*MockManifestsGeneratorAPI)(nil).AddTelemeterManifest), ctx, log, c)
}
