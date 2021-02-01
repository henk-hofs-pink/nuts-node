package network

import (
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"strings"
)

// Config holds the config for Network
type Config struct {
	// Socket address for gRPC to listen on
	GrpcAddr     string
	DatabaseFile string
	// EnableTLS specifies whether to enable TLS for incoming connections.
	EnableTLS bool
	// Public address of this nodes other nodes can use to connect to this node.
	PublicAddr     string
	BootstrapNodes string
	CertFile       string
	CertKeyFile    string
	TrustStoreFile string

	// AdvertHashesInterval specifies how often (in milliseconds) the node should broadcasts its last hashes so
	// other nodes can compare and synchronize.
	AdvertHashesInterval int
}

// DefaultConfig returns the default NetworkEngine configuration.
func DefaultConfig() Config {
	return Config{
		GrpcAddr:             ":5555",
		DatabaseFile:         "network.db",
		EnableTLS:            true,
		AdvertHashesInterval: 2000,
	}
}

func (c Config) parseBootstrapNodes() []string {
	var result []string
	for _, addr := range strings.Split(c.BootstrapNodes, " ") {
		trimmedAddr := strings.TrimSpace(addr)
		if trimmedAddr != "" {
			result = append(result, trimmedAddr)
		}
	}
	return result
}

func (c Config) loadTrustStore() (*x509.CertPool, error) {
	trustStore := x509.NewCertPool()
	data, err := ioutil.ReadFile(c.TrustStoreFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read trust store (file=%s): %w", c.TrustStoreFile, err)
	}
	if ok := trustStore.AppendCertsFromPEM(data); !ok {
		return nil, fmt.Errorf("unable to load one or more certificates from trust store (file=%s)", c.TrustStoreFile)
	}
	return trustStore, nil
}