// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/caos/zitadel/pkg/grpc/admin (interfaces: AdminServiceClient)

// Package api is a generated GoMock package.
package api

import (
	context "context"
	admin "github.com/caos/zitadel/pkg/grpc/admin"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
)

// MockAdminServiceClient is a mock of AdminServiceClient interface
type MockAdminServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockAdminServiceClientMockRecorder
}

// MockAdminServiceClientMockRecorder is the mock recorder for MockAdminServiceClient
type MockAdminServiceClientMockRecorder struct {
	mock *MockAdminServiceClient
}

// NewMockAdminServiceClient creates a new mock instance
func NewMockAdminServiceClient(ctrl *gomock.Controller) *MockAdminServiceClient {
	mock := &MockAdminServiceClient{ctrl: ctrl}
	mock.recorder = &MockAdminServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAdminServiceClient) EXPECT() *MockAdminServiceClientMockRecorder {
	return m.recorder
}

// AddIamMember mocks base method
func (m *MockAdminServiceClient) AddIamMember(arg0 context.Context, arg1 *admin.AddIamMemberRequest, arg2 ...grpc.CallOption) (*admin.IamMember, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddIamMember", varargs...)
	ret0, _ := ret[0].(*admin.IamMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddIamMember indicates an expected call of AddIamMember
func (mr *MockAdminServiceClientMockRecorder) AddIamMember(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddIamMember", reflect.TypeOf((*MockAdminServiceClient)(nil).AddIamMember), varargs...)
}

// ChangeIamMember mocks base method
func (m *MockAdminServiceClient) ChangeIamMember(arg0 context.Context, arg1 *admin.ChangeIamMemberRequest, arg2 ...grpc.CallOption) (*admin.IamMember, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChangeIamMember", varargs...)
	ret0, _ := ret[0].(*admin.IamMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeIamMember indicates an expected call of ChangeIamMember
func (mr *MockAdminServiceClientMockRecorder) ChangeIamMember(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeIamMember", reflect.TypeOf((*MockAdminServiceClient)(nil).ChangeIamMember), varargs...)
}

// ClearView mocks base method
func (m *MockAdminServiceClient) ClearView(arg0 context.Context, arg1 *admin.ViewID, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ClearView", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClearView indicates an expected call of ClearView
func (mr *MockAdminServiceClientMockRecorder) ClearView(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearView", reflect.TypeOf((*MockAdminServiceClient)(nil).ClearView), varargs...)
}

// CreateOidcIdp mocks base method
func (m *MockAdminServiceClient) CreateOidcIdp(arg0 context.Context, arg1 *admin.OidcIdpConfigCreate, arg2 ...grpc.CallOption) (*admin.Idp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateOidcIdp", varargs...)
	ret0, _ := ret[0].(*admin.Idp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOidcIdp indicates an expected call of CreateOidcIdp
func (mr *MockAdminServiceClientMockRecorder) CreateOidcIdp(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOidcIdp", reflect.TypeOf((*MockAdminServiceClient)(nil).CreateOidcIdp), varargs...)
}

// CreateOrgIamPolicy mocks base method
func (m *MockAdminServiceClient) CreateOrgIamPolicy(arg0 context.Context, arg1 *admin.OrgIamPolicyRequest, arg2 ...grpc.CallOption) (*admin.OrgIamPolicy, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateOrgIamPolicy", varargs...)
	ret0, _ := ret[0].(*admin.OrgIamPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrgIamPolicy indicates an expected call of CreateOrgIamPolicy
func (mr *MockAdminServiceClientMockRecorder) CreateOrgIamPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrgIamPolicy", reflect.TypeOf((*MockAdminServiceClient)(nil).CreateOrgIamPolicy), varargs...)
}

// DeactivateIdpConfig mocks base method
func (m *MockAdminServiceClient) DeactivateIdpConfig(arg0 context.Context, arg1 *admin.IdpID, arg2 ...grpc.CallOption) (*admin.Idp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeactivateIdpConfig", varargs...)
	ret0, _ := ret[0].(*admin.Idp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeactivateIdpConfig indicates an expected call of DeactivateIdpConfig
func (mr *MockAdminServiceClientMockRecorder) DeactivateIdpConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeactivateIdpConfig", reflect.TypeOf((*MockAdminServiceClient)(nil).DeactivateIdpConfig), varargs...)
}

// DeleteOrgIamPolicy mocks base method
func (m *MockAdminServiceClient) DeleteOrgIamPolicy(arg0 context.Context, arg1 *admin.OrgIamPolicyID, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteOrgIamPolicy", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteOrgIamPolicy indicates an expected call of DeleteOrgIamPolicy
func (mr *MockAdminServiceClientMockRecorder) DeleteOrgIamPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrgIamPolicy", reflect.TypeOf((*MockAdminServiceClient)(nil).DeleteOrgIamPolicy), varargs...)
}

// GetFailedEvents mocks base method
func (m *MockAdminServiceClient) GetFailedEvents(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*admin.FailedEvents, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFailedEvents", varargs...)
	ret0, _ := ret[0].(*admin.FailedEvents)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFailedEvents indicates an expected call of GetFailedEvents
func (mr *MockAdminServiceClientMockRecorder) GetFailedEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFailedEvents", reflect.TypeOf((*MockAdminServiceClient)(nil).GetFailedEvents), varargs...)
}

// GetIamMemberRoles mocks base method
func (m *MockAdminServiceClient) GetIamMemberRoles(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*admin.IamMemberRoles, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIamMemberRoles", varargs...)
	ret0, _ := ret[0].(*admin.IamMemberRoles)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIamMemberRoles indicates an expected call of GetIamMemberRoles
func (mr *MockAdminServiceClientMockRecorder) GetIamMemberRoles(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIamMemberRoles", reflect.TypeOf((*MockAdminServiceClient)(nil).GetIamMemberRoles), varargs...)
}

// GetOrgByID mocks base method
func (m *MockAdminServiceClient) GetOrgByID(arg0 context.Context, arg1 *admin.OrgID, arg2 ...grpc.CallOption) (*admin.Org, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOrgByID", varargs...)
	ret0, _ := ret[0].(*admin.Org)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgByID indicates an expected call of GetOrgByID
func (mr *MockAdminServiceClientMockRecorder) GetOrgByID(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgByID", reflect.TypeOf((*MockAdminServiceClient)(nil).GetOrgByID), varargs...)
}

// GetOrgIamPolicy mocks base method
func (m *MockAdminServiceClient) GetOrgIamPolicy(arg0 context.Context, arg1 *admin.OrgIamPolicyID, arg2 ...grpc.CallOption) (*admin.OrgIamPolicy, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOrgIamPolicy", varargs...)
	ret0, _ := ret[0].(*admin.OrgIamPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgIamPolicy indicates an expected call of GetOrgIamPolicy
func (mr *MockAdminServiceClientMockRecorder) GetOrgIamPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgIamPolicy", reflect.TypeOf((*MockAdminServiceClient)(nil).GetOrgIamPolicy), varargs...)
}

// GetViews mocks base method
func (m *MockAdminServiceClient) GetViews(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*admin.Views, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetViews", varargs...)
	ret0, _ := ret[0].(*admin.Views)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetViews indicates an expected call of GetViews
func (mr *MockAdminServiceClientMockRecorder) GetViews(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetViews", reflect.TypeOf((*MockAdminServiceClient)(nil).GetViews), varargs...)
}

// Healthz mocks base method
func (m *MockAdminServiceClient) Healthz(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Healthz", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Healthz indicates an expected call of Healthz
func (mr *MockAdminServiceClientMockRecorder) Healthz(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Healthz", reflect.TypeOf((*MockAdminServiceClient)(nil).Healthz), varargs...)
}

// IsOrgUnique mocks base method
func (m *MockAdminServiceClient) IsOrgUnique(arg0 context.Context, arg1 *admin.UniqueOrgRequest, arg2 ...grpc.CallOption) (*admin.UniqueOrgResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IsOrgUnique", varargs...)
	ret0, _ := ret[0].(*admin.UniqueOrgResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsOrgUnique indicates an expected call of IsOrgUnique
func (mr *MockAdminServiceClientMockRecorder) IsOrgUnique(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOrgUnique", reflect.TypeOf((*MockAdminServiceClient)(nil).IsOrgUnique), varargs...)
}

// ReactivateIdpConfig mocks base method
func (m *MockAdminServiceClient) ReactivateIdpConfig(arg0 context.Context, arg1 *admin.IdpID, arg2 ...grpc.CallOption) (*admin.Idp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReactivateIdpConfig", varargs...)
	ret0, _ := ret[0].(*admin.Idp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReactivateIdpConfig indicates an expected call of ReactivateIdpConfig
func (mr *MockAdminServiceClientMockRecorder) ReactivateIdpConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReactivateIdpConfig", reflect.TypeOf((*MockAdminServiceClient)(nil).ReactivateIdpConfig), varargs...)
}

// Ready mocks base method
func (m *MockAdminServiceClient) Ready(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Ready", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Ready indicates an expected call of Ready
func (mr *MockAdminServiceClientMockRecorder) Ready(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ready", reflect.TypeOf((*MockAdminServiceClient)(nil).Ready), varargs...)
}

// RemoveFailedEvent mocks base method
func (m *MockAdminServiceClient) RemoveFailedEvent(arg0 context.Context, arg1 *admin.FailedEventID, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveFailedEvent", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveFailedEvent indicates an expected call of RemoveFailedEvent
func (mr *MockAdminServiceClientMockRecorder) RemoveFailedEvent(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFailedEvent", reflect.TypeOf((*MockAdminServiceClient)(nil).RemoveFailedEvent), varargs...)
}

// RemoveIamMember mocks base method
func (m *MockAdminServiceClient) RemoveIamMember(arg0 context.Context, arg1 *admin.RemoveIamMemberRequest, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveIamMember", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveIamMember indicates an expected call of RemoveIamMember
func (mr *MockAdminServiceClientMockRecorder) RemoveIamMember(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveIamMember", reflect.TypeOf((*MockAdminServiceClient)(nil).RemoveIamMember), varargs...)
}

// RemoveIdpConfig mocks base method
func (m *MockAdminServiceClient) RemoveIdpConfig(arg0 context.Context, arg1 *admin.IdpID, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveIdpConfig", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveIdpConfig indicates an expected call of RemoveIdpConfig
func (mr *MockAdminServiceClientMockRecorder) RemoveIdpConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveIdpConfig", reflect.TypeOf((*MockAdminServiceClient)(nil).RemoveIdpConfig), varargs...)
}

// SearchIamMembers mocks base method
func (m *MockAdminServiceClient) SearchIamMembers(arg0 context.Context, arg1 *admin.IamMemberSearchRequest, arg2 ...grpc.CallOption) (*admin.IamMemberSearchResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchIamMembers", varargs...)
	ret0, _ := ret[0].(*admin.IamMemberSearchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchIamMembers indicates an expected call of SearchIamMembers
func (mr *MockAdminServiceClientMockRecorder) SearchIamMembers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchIamMembers", reflect.TypeOf((*MockAdminServiceClient)(nil).SearchIamMembers), varargs...)
}

// SearchOrgs mocks base method
func (m *MockAdminServiceClient) SearchOrgs(arg0 context.Context, arg1 *admin.OrgSearchRequest, arg2 ...grpc.CallOption) (*admin.OrgSearchResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchOrgs", varargs...)
	ret0, _ := ret[0].(*admin.OrgSearchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchOrgs indicates an expected call of SearchOrgs
func (mr *MockAdminServiceClientMockRecorder) SearchOrgs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchOrgs", reflect.TypeOf((*MockAdminServiceClient)(nil).SearchOrgs), varargs...)
}

// SetUpOrg mocks base method
func (m *MockAdminServiceClient) SetUpOrg(arg0 context.Context, arg1 *admin.OrgSetUpRequest, arg2 ...grpc.CallOption) (*admin.OrgSetUpResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetUpOrg", varargs...)
	ret0, _ := ret[0].(*admin.OrgSetUpResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetUpOrg indicates an expected call of SetUpOrg
func (mr *MockAdminServiceClientMockRecorder) SetUpOrg(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUpOrg", reflect.TypeOf((*MockAdminServiceClient)(nil).SetUpOrg), varargs...)
}

// UpdateIdpConfig mocks base method
func (m *MockAdminServiceClient) UpdateIdpConfig(arg0 context.Context, arg1 *admin.IdpUpdate, arg2 ...grpc.CallOption) (*admin.Idp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateIdpConfig", varargs...)
	ret0, _ := ret[0].(*admin.Idp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateIdpConfig indicates an expected call of UpdateIdpConfig
func (mr *MockAdminServiceClientMockRecorder) UpdateIdpConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIdpConfig", reflect.TypeOf((*MockAdminServiceClient)(nil).UpdateIdpConfig), varargs...)
}

// UpdateOidcIdpConfig mocks base method
func (m *MockAdminServiceClient) UpdateOidcIdpConfig(arg0 context.Context, arg1 *admin.OidcIdpConfigUpdate, arg2 ...grpc.CallOption) (*admin.OidcIdpConfig, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateOidcIdpConfig", varargs...)
	ret0, _ := ret[0].(*admin.OidcIdpConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOidcIdpConfig indicates an expected call of UpdateOidcIdpConfig
func (mr *MockAdminServiceClientMockRecorder) UpdateOidcIdpConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOidcIdpConfig", reflect.TypeOf((*MockAdminServiceClient)(nil).UpdateOidcIdpConfig), varargs...)
}

// UpdateOrgIamPolicy mocks base method
func (m *MockAdminServiceClient) UpdateOrgIamPolicy(arg0 context.Context, arg1 *admin.OrgIamPolicyRequest, arg2 ...grpc.CallOption) (*admin.OrgIamPolicy, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateOrgIamPolicy", varargs...)
	ret0, _ := ret[0].(*admin.OrgIamPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrgIamPolicy indicates an expected call of UpdateOrgIamPolicy
func (mr *MockAdminServiceClientMockRecorder) UpdateOrgIamPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrgIamPolicy", reflect.TypeOf((*MockAdminServiceClient)(nil).UpdateOrgIamPolicy), varargs...)
}

// Validate mocks base method
func (m *MockAdminServiceClient) Validate(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*structpb.Struct, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Validate", varargs...)
	ret0, _ := ret[0].(*structpb.Struct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Validate indicates an expected call of Validate
func (mr *MockAdminServiceClientMockRecorder) Validate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockAdminServiceClient)(nil).Validate), varargs...)
}
