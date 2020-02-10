// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/ARO-RP/pkg/util/azureclient/mgmt/redhatopenshift (interfaces: OpenShiftClustersClient,OperationsClient)

// Package mock_redhatopenshift is a generated GoMock package.
package mock_redhatopenshift

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	redhatopenshift "github.com/Azure/ARO-RP/pkg/client/services/preview/redhatopenshift/mgmt/2019-12-31-preview/redhatopenshift"
)

// MockOpenShiftClustersClient is a mock of OpenShiftClustersClient interface
type MockOpenShiftClustersClient struct {
	ctrl     *gomock.Controller
	recorder *MockOpenShiftClustersClientMockRecorder
}

// MockOpenShiftClustersClientMockRecorder is the mock recorder for MockOpenShiftClustersClient
type MockOpenShiftClustersClientMockRecorder struct {
	mock *MockOpenShiftClustersClient
}

// NewMockOpenShiftClustersClient creates a new mock instance
func NewMockOpenShiftClustersClient(ctrl *gomock.Controller) *MockOpenShiftClustersClient {
	mock := &MockOpenShiftClustersClient{ctrl: ctrl}
	mock.recorder = &MockOpenShiftClustersClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOpenShiftClustersClient) EXPECT() *MockOpenShiftClustersClientMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockOpenShiftClustersClient) Get(arg0 context.Context, arg1, arg2 string) (redhatopenshift.OpenShiftCluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(redhatopenshift.OpenShiftCluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockOpenShiftClustersClientMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockOpenShiftClustersClient)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method
func (m *MockOpenShiftClustersClient) List(arg0 context.Context) (redhatopenshift.OpenShiftClusterList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(redhatopenshift.OpenShiftClusterList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockOpenShiftClustersClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockOpenShiftClustersClient)(nil).List), arg0)
}

// ListByResourceGroup mocks base method
func (m *MockOpenShiftClustersClient) ListByResourceGroup(arg0 context.Context, arg1 string) (redhatopenshift.OpenShiftClusterList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByResourceGroup", arg0, arg1)
	ret0, _ := ret[0].(redhatopenshift.OpenShiftClusterList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByResourceGroup indicates an expected call of ListByResourceGroup
func (mr *MockOpenShiftClustersClientMockRecorder) ListByResourceGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByResourceGroup", reflect.TypeOf((*MockOpenShiftClustersClient)(nil).ListByResourceGroup), arg0, arg1)
}

// ListCredentials mocks base method
func (m *MockOpenShiftClustersClient) ListCredentials(arg0 context.Context, arg1, arg2 string) (redhatopenshift.OpenShiftClusterCredentials, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCredentials", arg0, arg1, arg2)
	ret0, _ := ret[0].(redhatopenshift.OpenShiftClusterCredentials)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCredentials indicates an expected call of ListCredentials
func (mr *MockOpenShiftClustersClientMockRecorder) ListCredentials(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCredentials", reflect.TypeOf((*MockOpenShiftClustersClient)(nil).ListCredentials), arg0, arg1, arg2)
}

// MockOperationsClient is a mock of OperationsClient interface
type MockOperationsClient struct {
	ctrl     *gomock.Controller
	recorder *MockOperationsClientMockRecorder
}

// MockOperationsClientMockRecorder is the mock recorder for MockOperationsClient
type MockOperationsClientMockRecorder struct {
	mock *MockOperationsClient
}

// NewMockOperationsClient creates a new mock instance
func NewMockOperationsClient(ctrl *gomock.Controller) *MockOperationsClient {
	mock := &MockOperationsClient{ctrl: ctrl}
	mock.recorder = &MockOperationsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOperationsClient) EXPECT() *MockOperationsClientMockRecorder {
	return m.recorder
}

// List mocks base method
func (m *MockOperationsClient) List(arg0 context.Context) (redhatopenshift.OperationList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(redhatopenshift.OperationList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockOperationsClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockOperationsClient)(nil).List), arg0)
}
