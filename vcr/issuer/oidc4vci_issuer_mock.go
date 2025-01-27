// Code generated by MockGen. DO NOT EDIT.
// Source: vcr/issuer/oidc4vci_issuer.go

// Package issuer is a generated GoMock package.
package issuer

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	did "github.com/nuts-foundation/go-did/did"
	vc "github.com/nuts-foundation/go-did/vc"
	oidc4vci "github.com/nuts-foundation/nuts-node/vcr/oidc4vci"
)

// MockOIDCIssuer is a mock of OIDCIssuer interface.
type MockOIDCIssuer struct {
	ctrl     *gomock.Controller
	recorder *MockOIDCIssuerMockRecorder
}

// MockOIDCIssuerMockRecorder is the mock recorder for MockOIDCIssuer.
type MockOIDCIssuerMockRecorder struct {
	mock *MockOIDCIssuer
}

// NewMockOIDCIssuer creates a new mock instance.
func NewMockOIDCIssuer(ctrl *gomock.Controller) *MockOIDCIssuer {
	mock := &MockOIDCIssuer{ctrl: ctrl}
	mock.recorder = &MockOIDCIssuerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOIDCIssuer) EXPECT() *MockOIDCIssuerMockRecorder {
	return m.recorder
}

// HandleAccessTokenRequest mocks base method.
func (m *MockOIDCIssuer) HandleAccessTokenRequest(ctx context.Context, issuer did.DID, preAuthorizedCode string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleAccessTokenRequest", ctx, issuer, preAuthorizedCode)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleAccessTokenRequest indicates an expected call of HandleAccessTokenRequest.
func (mr *MockOIDCIssuerMockRecorder) HandleAccessTokenRequest(ctx, issuer, preAuthorizedCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleAccessTokenRequest", reflect.TypeOf((*MockOIDCIssuer)(nil).HandleAccessTokenRequest), ctx, issuer, preAuthorizedCode)
}

// HandleCredentialRequest mocks base method.
func (m *MockOIDCIssuer) HandleCredentialRequest(ctx context.Context, issuer did.DID, request oidc4vci.CredentialRequest, accessToken string) (*vc.VerifiableCredential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleCredentialRequest", ctx, issuer, request, accessToken)
	ret0, _ := ret[0].(*vc.VerifiableCredential)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleCredentialRequest indicates an expected call of HandleCredentialRequest.
func (mr *MockOIDCIssuerMockRecorder) HandleCredentialRequest(ctx, issuer, request, accessToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleCredentialRequest", reflect.TypeOf((*MockOIDCIssuer)(nil).HandleCredentialRequest), ctx, issuer, request, accessToken)
}

// Metadata mocks base method.
func (m *MockOIDCIssuer) Metadata(issuer did.DID) (oidc4vci.CredentialIssuerMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Metadata", issuer)
	ret0, _ := ret[0].(oidc4vci.CredentialIssuerMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Metadata indicates an expected call of Metadata.
func (mr *MockOIDCIssuerMockRecorder) Metadata(issuer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Metadata", reflect.TypeOf((*MockOIDCIssuer)(nil).Metadata), issuer)
}

// OfferCredential mocks base method.
func (m *MockOIDCIssuer) OfferCredential(ctx context.Context, credential vc.VerifiableCredential, walletURL string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OfferCredential", ctx, credential, walletURL)
	ret0, _ := ret[0].(error)
	return ret0
}

// OfferCredential indicates an expected call of OfferCredential.
func (mr *MockOIDCIssuerMockRecorder) OfferCredential(ctx, credential, walletURL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OfferCredential", reflect.TypeOf((*MockOIDCIssuer)(nil).OfferCredential), ctx, credential, walletURL)
}

// ProviderMetadata mocks base method.
func (m *MockOIDCIssuer) ProviderMetadata(issuer did.DID) (oidc4vci.ProviderMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProviderMetadata", issuer)
	ret0, _ := ret[0].(oidc4vci.ProviderMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProviderMetadata indicates an expected call of ProviderMetadata.
func (mr *MockOIDCIssuerMockRecorder) ProviderMetadata(issuer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProviderMetadata", reflect.TypeOf((*MockOIDCIssuer)(nil).ProviderMetadata), issuer)
}
