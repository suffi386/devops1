// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/caos/zitadel/pkg/grpc/auth (interfaces: AuthServiceClient)

// Package api is a generated GoMock package.
package api

import (
	context "context"
	auth "github.com/caos/zitadel/pkg/grpc/auth"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
)

// MockAuthServiceClient is a mock of AuthServiceClient interface
type MockAuthServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceClientMockRecorder
}

// MockAuthServiceClientMockRecorder is the mock recorder for MockAuthServiceClient
type MockAuthServiceClientMockRecorder struct {
	mock *MockAuthServiceClient
}

// NewMockAuthServiceClient creates a new mock instance
func NewMockAuthServiceClient(ctrl *gomock.Controller) *MockAuthServiceClient {
	mock := &MockAuthServiceClient{ctrl: ctrl}
	mock.recorder = &MockAuthServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthServiceClient) EXPECT() *MockAuthServiceClientMockRecorder {
	return m.recorder
}

// AddMfaOTP mocks base method
func (m *MockAuthServiceClient) AddMfaOTP(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*auth.MfaOtpResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddMfaOTP", varargs...)
	ret0, _ := ret[0].(*auth.MfaOtpResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddMfaOTP indicates an expected call of AddMfaOTP
func (mr *MockAuthServiceClientMockRecorder) AddMfaOTP(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMfaOTP", reflect.TypeOf((*MockAuthServiceClient)(nil).AddMfaOTP), varargs...)
}

// AddMyExternalIDP mocks base method
func (m *MockAuthServiceClient) AddMyExternalIDP(arg0 context.Context, arg1 *auth.ExternalIDPAddRequest, arg2 ...grpc.CallOption) (*auth.ExternalIDPResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddMyExternalIDP", varargs...)
	ret0, _ := ret[0].(*auth.ExternalIDPResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddMyExternalIDP indicates an expected call of AddMyExternalIDP
func (mr *MockAuthServiceClientMockRecorder) AddMyExternalIDP(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMyExternalIDP", reflect.TypeOf((*MockAuthServiceClient)(nil).AddMyExternalIDP), varargs...)
}

// ChangeMyPassword mocks base method
func (m *MockAuthServiceClient) ChangeMyPassword(arg0 context.Context, arg1 *auth.PasswordChange, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChangeMyPassword", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeMyPassword indicates an expected call of ChangeMyPassword
func (mr *MockAuthServiceClientMockRecorder) ChangeMyPassword(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeMyPassword", reflect.TypeOf((*MockAuthServiceClient)(nil).ChangeMyPassword), varargs...)
}

// ChangeMyUserEmail mocks base method
func (m *MockAuthServiceClient) ChangeMyUserEmail(arg0 context.Context, arg1 *auth.UpdateUserEmailRequest, arg2 ...grpc.CallOption) (*auth.UserEmail, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChangeMyUserEmail", varargs...)
	ret0, _ := ret[0].(*auth.UserEmail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeMyUserEmail indicates an expected call of ChangeMyUserEmail
func (mr *MockAuthServiceClientMockRecorder) ChangeMyUserEmail(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeMyUserEmail", reflect.TypeOf((*MockAuthServiceClient)(nil).ChangeMyUserEmail), varargs...)
}

// ChangeMyUserName mocks base method
func (m *MockAuthServiceClient) ChangeMyUserName(arg0 context.Context, arg1 *auth.ChangeUserNameRequest, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChangeMyUserName", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeMyUserName indicates an expected call of ChangeMyUserName
func (mr *MockAuthServiceClientMockRecorder) ChangeMyUserName(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeMyUserName", reflect.TypeOf((*MockAuthServiceClient)(nil).ChangeMyUserName), varargs...)
}

// ChangeMyUserPhone mocks base method
func (m *MockAuthServiceClient) ChangeMyUserPhone(arg0 context.Context, arg1 *auth.UpdateUserPhoneRequest, arg2 ...grpc.CallOption) (*auth.UserPhone, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChangeMyUserPhone", varargs...)
	ret0, _ := ret[0].(*auth.UserPhone)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeMyUserPhone indicates an expected call of ChangeMyUserPhone
func (mr *MockAuthServiceClientMockRecorder) ChangeMyUserPhone(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeMyUserPhone", reflect.TypeOf((*MockAuthServiceClient)(nil).ChangeMyUserPhone), varargs...)
}

// GetMyMfas mocks base method
func (m *MockAuthServiceClient) GetMyMfas(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*auth.MultiFactors, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMyMfas", varargs...)
	ret0, _ := ret[0].(*auth.MultiFactors)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMyMfas indicates an expected call of GetMyMfas
func (mr *MockAuthServiceClientMockRecorder) GetMyMfas(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMyMfas", reflect.TypeOf((*MockAuthServiceClient)(nil).GetMyMfas), varargs...)
}

// GetMyPasswordComplexityPolicy mocks base method
func (m *MockAuthServiceClient) GetMyPasswordComplexityPolicy(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*auth.PasswordComplexityPolicy, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMyPasswordComplexityPolicy", varargs...)
	ret0, _ := ret[0].(*auth.PasswordComplexityPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMyPasswordComplexityPolicy indicates an expected call of GetMyPasswordComplexityPolicy
func (mr *MockAuthServiceClientMockRecorder) GetMyPasswordComplexityPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMyPasswordComplexityPolicy", reflect.TypeOf((*MockAuthServiceClient)(nil).GetMyPasswordComplexityPolicy), varargs...)
}

// GetMyProjectPermissions mocks base method
func (m *MockAuthServiceClient) GetMyProjectPermissions(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*auth.MyPermissions, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMyProjectPermissions", varargs...)
	ret0, _ := ret[0].(*auth.MyPermissions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMyProjectPermissions indicates an expected call of GetMyProjectPermissions
func (mr *MockAuthServiceClientMockRecorder) GetMyProjectPermissions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMyProjectPermissions", reflect.TypeOf((*MockAuthServiceClient)(nil).GetMyProjectPermissions), varargs...)
}

// GetMyUser mocks base method
func (m *MockAuthServiceClient) GetMyUser(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*auth.UserView, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMyUser", varargs...)
	ret0, _ := ret[0].(*auth.UserView)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMyUser indicates an expected call of GetMyUser
func (mr *MockAuthServiceClientMockRecorder) GetMyUser(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMyUser", reflect.TypeOf((*MockAuthServiceClient)(nil).GetMyUser), varargs...)
}

// GetMyUserAddress mocks base method
func (m *MockAuthServiceClient) GetMyUserAddress(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*auth.UserAddressView, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMyUserAddress", varargs...)
	ret0, _ := ret[0].(*auth.UserAddressView)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMyUserAddress indicates an expected call of GetMyUserAddress
func (mr *MockAuthServiceClientMockRecorder) GetMyUserAddress(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMyUserAddress", reflect.TypeOf((*MockAuthServiceClient)(nil).GetMyUserAddress), varargs...)
}

// GetMyUserChanges mocks base method
func (m *MockAuthServiceClient) GetMyUserChanges(arg0 context.Context, arg1 *auth.ChangesRequest, arg2 ...grpc.CallOption) (*auth.Changes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMyUserChanges", varargs...)
	ret0, _ := ret[0].(*auth.Changes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMyUserChanges indicates an expected call of GetMyUserChanges
func (mr *MockAuthServiceClientMockRecorder) GetMyUserChanges(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMyUserChanges", reflect.TypeOf((*MockAuthServiceClient)(nil).GetMyUserChanges), varargs...)
}

// GetMyUserEmail mocks base method
func (m *MockAuthServiceClient) GetMyUserEmail(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*auth.UserEmailView, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMyUserEmail", varargs...)
	ret0, _ := ret[0].(*auth.UserEmailView)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMyUserEmail indicates an expected call of GetMyUserEmail
func (mr *MockAuthServiceClientMockRecorder) GetMyUserEmail(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMyUserEmail", reflect.TypeOf((*MockAuthServiceClient)(nil).GetMyUserEmail), varargs...)
}

// GetMyUserPhone mocks base method
func (m *MockAuthServiceClient) GetMyUserPhone(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*auth.UserPhoneView, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMyUserPhone", varargs...)
	ret0, _ := ret[0].(*auth.UserPhoneView)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMyUserPhone indicates an expected call of GetMyUserPhone
func (mr *MockAuthServiceClientMockRecorder) GetMyUserPhone(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMyUserPhone", reflect.TypeOf((*MockAuthServiceClient)(nil).GetMyUserPhone), varargs...)
}

// GetMyUserProfile mocks base method
func (m *MockAuthServiceClient) GetMyUserProfile(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*auth.UserProfileView, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMyUserProfile", varargs...)
	ret0, _ := ret[0].(*auth.UserProfileView)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMyUserProfile indicates an expected call of GetMyUserProfile
func (mr *MockAuthServiceClientMockRecorder) GetMyUserProfile(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMyUserProfile", reflect.TypeOf((*MockAuthServiceClient)(nil).GetMyUserProfile), varargs...)
}

// GetMyUserSessions mocks base method
func (m *MockAuthServiceClient) GetMyUserSessions(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*auth.UserSessionViews, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMyUserSessions", varargs...)
	ret0, _ := ret[0].(*auth.UserSessionViews)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMyUserSessions indicates an expected call of GetMyUserSessions
func (mr *MockAuthServiceClientMockRecorder) GetMyUserSessions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMyUserSessions", reflect.TypeOf((*MockAuthServiceClient)(nil).GetMyUserSessions), varargs...)
}

// GetMyZitadelPermissions mocks base method
func (m *MockAuthServiceClient) GetMyZitadelPermissions(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*auth.MyPermissions, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMyZitadelPermissions", varargs...)
	ret0, _ := ret[0].(*auth.MyPermissions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMyZitadelPermissions indicates an expected call of GetMyZitadelPermissions
func (mr *MockAuthServiceClientMockRecorder) GetMyZitadelPermissions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMyZitadelPermissions", reflect.TypeOf((*MockAuthServiceClient)(nil).GetMyZitadelPermissions), varargs...)
}

// Healthz mocks base method
func (m *MockAuthServiceClient) Healthz(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
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
func (mr *MockAuthServiceClientMockRecorder) Healthz(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Healthz", reflect.TypeOf((*MockAuthServiceClient)(nil).Healthz), varargs...)
}

// Ready mocks base method
func (m *MockAuthServiceClient) Ready(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
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
func (mr *MockAuthServiceClientMockRecorder) Ready(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ready", reflect.TypeOf((*MockAuthServiceClient)(nil).Ready), varargs...)
}

// RemoveMfaOTP mocks base method
func (m *MockAuthServiceClient) RemoveMfaOTP(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveMfaOTP", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveMfaOTP indicates an expected call of RemoveMfaOTP
func (mr *MockAuthServiceClientMockRecorder) RemoveMfaOTP(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveMfaOTP", reflect.TypeOf((*MockAuthServiceClient)(nil).RemoveMfaOTP), varargs...)
}

// RemoveMyExternalIDP mocks base method
func (m *MockAuthServiceClient) RemoveMyExternalIDP(arg0 context.Context, arg1 *auth.ExternalIDPRemoveRequest, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveMyExternalIDP", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveMyExternalIDP indicates an expected call of RemoveMyExternalIDP
func (mr *MockAuthServiceClientMockRecorder) RemoveMyExternalIDP(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveMyExternalIDP", reflect.TypeOf((*MockAuthServiceClient)(nil).RemoveMyExternalIDP), varargs...)
}

// RemoveMyUserPhone mocks base method
func (m *MockAuthServiceClient) RemoveMyUserPhone(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveMyUserPhone", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveMyUserPhone indicates an expected call of RemoveMyUserPhone
func (mr *MockAuthServiceClientMockRecorder) RemoveMyUserPhone(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveMyUserPhone", reflect.TypeOf((*MockAuthServiceClient)(nil).RemoveMyUserPhone), varargs...)
}

// ResendMyEmailVerificationMail mocks base method
func (m *MockAuthServiceClient) ResendMyEmailVerificationMail(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ResendMyEmailVerificationMail", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResendMyEmailVerificationMail indicates an expected call of ResendMyEmailVerificationMail
func (mr *MockAuthServiceClientMockRecorder) ResendMyEmailVerificationMail(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResendMyEmailVerificationMail", reflect.TypeOf((*MockAuthServiceClient)(nil).ResendMyEmailVerificationMail), varargs...)
}

// ResendMyPhoneVerificationCode mocks base method
func (m *MockAuthServiceClient) ResendMyPhoneVerificationCode(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ResendMyPhoneVerificationCode", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResendMyPhoneVerificationCode indicates an expected call of ResendMyPhoneVerificationCode
func (mr *MockAuthServiceClientMockRecorder) ResendMyPhoneVerificationCode(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResendMyPhoneVerificationCode", reflect.TypeOf((*MockAuthServiceClient)(nil).ResendMyPhoneVerificationCode), varargs...)
}

// SearchMyExternalIDPs mocks base method
func (m *MockAuthServiceClient) SearchMyExternalIDPs(arg0 context.Context, arg1 *auth.ExternalIDPSearchRequest, arg2 ...grpc.CallOption) (*auth.ExternalIDPSearchResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchMyExternalIDPs", varargs...)
	ret0, _ := ret[0].(*auth.ExternalIDPSearchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchMyExternalIDPs indicates an expected call of SearchMyExternalIDPs
func (mr *MockAuthServiceClientMockRecorder) SearchMyExternalIDPs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchMyExternalIDPs", reflect.TypeOf((*MockAuthServiceClient)(nil).SearchMyExternalIDPs), varargs...)
}

// SearchMyProjectOrgs mocks base method
func (m *MockAuthServiceClient) SearchMyProjectOrgs(arg0 context.Context, arg1 *auth.MyProjectOrgSearchRequest, arg2 ...grpc.CallOption) (*auth.MyProjectOrgSearchResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchMyProjectOrgs", varargs...)
	ret0, _ := ret[0].(*auth.MyProjectOrgSearchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchMyProjectOrgs indicates an expected call of SearchMyProjectOrgs
func (mr *MockAuthServiceClientMockRecorder) SearchMyProjectOrgs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchMyProjectOrgs", reflect.TypeOf((*MockAuthServiceClient)(nil).SearchMyProjectOrgs), varargs...)
}

// SearchMyUserGrant mocks base method
func (m *MockAuthServiceClient) SearchMyUserGrant(arg0 context.Context, arg1 *auth.UserGrantSearchRequest, arg2 ...grpc.CallOption) (*auth.UserGrantSearchResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchMyUserGrant", varargs...)
	ret0, _ := ret[0].(*auth.UserGrantSearchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchMyUserGrant indicates an expected call of SearchMyUserGrant
func (mr *MockAuthServiceClientMockRecorder) SearchMyUserGrant(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchMyUserGrant", reflect.TypeOf((*MockAuthServiceClient)(nil).SearchMyUserGrant), varargs...)
}

// UpdateMyUserAddress mocks base method
func (m *MockAuthServiceClient) UpdateMyUserAddress(arg0 context.Context, arg1 *auth.UpdateUserAddressRequest, arg2 ...grpc.CallOption) (*auth.UserAddress, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateMyUserAddress", varargs...)
	ret0, _ := ret[0].(*auth.UserAddress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMyUserAddress indicates an expected call of UpdateMyUserAddress
func (mr *MockAuthServiceClientMockRecorder) UpdateMyUserAddress(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMyUserAddress", reflect.TypeOf((*MockAuthServiceClient)(nil).UpdateMyUserAddress), varargs...)
}

// UpdateMyUserProfile mocks base method
func (m *MockAuthServiceClient) UpdateMyUserProfile(arg0 context.Context, arg1 *auth.UpdateUserProfileRequest, arg2 ...grpc.CallOption) (*auth.UserProfile, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateMyUserProfile", varargs...)
	ret0, _ := ret[0].(*auth.UserProfile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMyUserProfile indicates an expected call of UpdateMyUserProfile
func (mr *MockAuthServiceClientMockRecorder) UpdateMyUserProfile(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMyUserProfile", reflect.TypeOf((*MockAuthServiceClient)(nil).UpdateMyUserProfile), varargs...)
}

// Validate mocks base method
func (m *MockAuthServiceClient) Validate(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*structpb.Struct, error) {
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
func (mr *MockAuthServiceClientMockRecorder) Validate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockAuthServiceClient)(nil).Validate), varargs...)
}

// VerifyMfaOTP mocks base method
func (m *MockAuthServiceClient) VerifyMfaOTP(arg0 context.Context, arg1 *auth.VerifyMfaOtp, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "VerifyMfaOTP", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyMfaOTP indicates an expected call of VerifyMfaOTP
func (mr *MockAuthServiceClientMockRecorder) VerifyMfaOTP(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyMfaOTP", reflect.TypeOf((*MockAuthServiceClient)(nil).VerifyMfaOTP), varargs...)
}

// VerifyMyUserEmail mocks base method
func (m *MockAuthServiceClient) VerifyMyUserEmail(arg0 context.Context, arg1 *auth.VerifyMyUserEmailRequest, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "VerifyMyUserEmail", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyMyUserEmail indicates an expected call of VerifyMyUserEmail
func (mr *MockAuthServiceClientMockRecorder) VerifyMyUserEmail(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyMyUserEmail", reflect.TypeOf((*MockAuthServiceClient)(nil).VerifyMyUserEmail), varargs...)
}

// VerifyMyUserPhone mocks base method
func (m *MockAuthServiceClient) VerifyMyUserPhone(arg0 context.Context, arg1 *auth.VerifyUserPhoneRequest, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "VerifyMyUserPhone", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyMyUserPhone indicates an expected call of VerifyMyUserPhone
func (mr *MockAuthServiceClientMockRecorder) VerifyMyUserPhone(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyMyUserPhone", reflect.TypeOf((*MockAuthServiceClient)(nil).VerifyMyUserPhone), varargs...)
}
