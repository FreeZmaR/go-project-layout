package redis

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"github.com/FreeZmaR/go-service-structure/template/config/types"
	"os"
)

func makeTLSConfig(cfg *types.RedisTLS) (*tls.Config, error) {
	if cfg.PEMFilePath == "" {
		return nil, errors.New("pem file path not provided")
	}

	if cfg.SNI == "" {
		return nil, errors.New("SNI not provided")
	}

	cert, err := tls.LoadX509KeyPair(cfg.PEMFilePath, cfg.KeyFilePath)
	if err != nil {
		return nil, err
	}

	tlsCfg := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   cfg.SNI,
		MinVersion:   tls.VersionTLS12,
	}

	if cfg.CAFilePath == "" {
		return tlsCfg, nil
	}

	caCertPool := x509.NewCertPool()

	caCert, err := os.ReadFile(cfg.CAFilePath)
	if err != nil {
		return nil, err
	}

	if !caCertPool.AppendCertsFromPEM(caCert) {
		return nil, errors.New("unable to add CA to cert pool")
	}

	tlsCfg.RootCAs = caCertPool
	tlsCfg.ClientCAs = caCertPool

	return tlsCfg, nil
}
