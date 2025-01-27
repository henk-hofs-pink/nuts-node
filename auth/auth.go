/*
 * Copyright (C) 2023 Nuts community
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 *
 */

package auth

import (
	"errors"
	"path"
	"time"

	"github.com/nuts-foundation/nuts-node/auth/services"
	"github.com/nuts-foundation/nuts-node/auth/services/notary"
	"github.com/nuts-foundation/nuts-node/auth/services/oauth"
	"github.com/nuts-foundation/nuts-node/core"
	"github.com/nuts-foundation/nuts-node/crypto"
	"github.com/nuts-foundation/nuts-node/didman"
	"github.com/nuts-foundation/nuts-node/jsonld"
	"github.com/nuts-foundation/nuts-node/pki"
	"github.com/nuts-foundation/nuts-node/vcr"
	"github.com/nuts-foundation/nuts-node/vdr/didservice"
	"github.com/nuts-foundation/nuts-node/vdr/didstore"
)

// ErrMissingPublicURL is returned when the publicUrl is missing from the config
var ErrMissingPublicURL = errors.New("auth.publicurl must be set in strictmode")

const contractValidity = 60 * time.Minute

var _ AuthenticationServices = (*Auth)(nil)

// Auth is the main struct of the Auth service
type Auth struct {
	config          Config
	jsonldManager   jsonld.JSONLD
	authzServer     oauth.AuthorizationServer
	relyingParty    oauth.RelyingParty
	contractNotary  services.ContractNotary
	serviceResolver didman.CompoundServiceResolver
	keyStore        crypto.KeyStore
	registry        didstore.Store
	vcr             vcr.VCR
	pkiProvider     pki.Provider
	shutdownFunc    func()
}

// Name returns the name of the module.
func (auth *Auth) Name() string {
	return ModuleName
}

// Config returns the actual config of the module.
func (auth *Auth) Config() interface{} {
	return &auth.config
}

// ContractNotary returns an implementation of the ContractNotary interface.
func (auth *Auth) ContractNotary() services.ContractNotary {
	return auth.contractNotary
}

// NewAuthInstance accepts a Config with several Nuts Engines and returns an instance of Auth
func NewAuthInstance(config Config, registry didstore.Store, vcr vcr.VCR, keyStore crypto.KeyStore, serviceResolver didman.CompoundServiceResolver, jsonldManager jsonld.JSONLD, pkiProvider pki.Provider) *Auth {
	return &Auth{
		config:          config,
		jsonldManager:   jsonldManager,
		registry:        registry,
		keyStore:        keyStore,
		vcr:             vcr,
		pkiProvider:     pkiProvider,
		serviceResolver: serviceResolver,
		shutdownFunc:    func() {},
	}
}

// AuthzServer returns the oauth.AuthorizationServer
func (auth *Auth) AuthzServer() oauth.AuthorizationServer {
	return auth.authzServer
}

// RelyingParty returns the oauth.RelyingParty
func (auth *Auth) RelyingParty() oauth.RelyingParty {
	return auth.relyingParty
}

// Configure the Auth struct by creating a validator and create an Irma server
func (auth *Auth) Configure(config core.ServerConfig) error {
	if auth.config.Irma.SchemeManager == "" {
		return errors.New("IRMA SchemeManager must be set")
	}

	if config.Strictmode && auth.config.Irma.SchemeManager != "pbdf" {
		return errors.New("in strictmode the only valid irma-scheme-manager is 'pbdf'")
	}

	// TODO: this is verifier/signer specific
	if auth.config.PublicURL == "" && config.Strictmode {
		return ErrMissingPublicURL
	}

	auth.contractNotary = notary.NewNotary(notary.Config{
		PublicURL:             auth.config.PublicURL,
		IrmaConfigPath:        path.Join(config.Datadir, "irma"),
		IrmaSchemeManager:     auth.config.Irma.SchemeManager,
		AutoUpdateIrmaSchemas: auth.config.Irma.AutoUpdateSchemas,
		ContractValidators:    auth.config.ContractValidators,
		ContractValidity:      contractValidity,
		StrictMode:            config.Strictmode,
	}, auth.vcr, didservice.KeyResolver{Store: auth.registry}, auth.keyStore, auth.jsonldManager, auth.pkiProvider)

	tlsEnabled := config.TLS.Enabled()
	if config.Strictmode && !tlsEnabled {
		return errors.New("in strictmode TLS must be enabled")
	}

	tlsConfig, err := auth.pkiProvider.CreateTLSConfig(config.TLS) // returns nil if TLS is disabled
	if err != nil {
		return err
	}

	if err := auth.contractNotary.Configure(); err != nil {
		return err
	}

	accessTokenLifeSpan := time.Duration(auth.config.AccessTokenLifeSpan) * time.Second
	auth.authzServer = oauth.NewAuthorizationServer(auth.registry, auth.vcr, auth.vcr.Verifier(), auth.serviceResolver,
		auth.keyStore, auth.contractNotary, auth.jsonldManager, accessTokenLifeSpan)
	auth.relyingParty = oauth.NewRelyingParty(auth.registry, auth.serviceResolver,
		auth.keyStore, time.Duration(auth.config.HTTPTimeout)*time.Second, tlsConfig)

	if err := auth.authzServer.Configure(auth.config.ClockSkew, config.Strictmode); err != nil {
		return err
	}

	return nil
}

// Start starts the Auth engine (Noop)
func (auth *Auth) Start() error {
	return nil
}

// Shutdown stops the Auth engine
func (auth *Auth) Shutdown() error {
	return nil
}
