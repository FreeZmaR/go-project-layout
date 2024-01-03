package types

import "crypto/tls"

type HTTPServer struct {
	Host            string   `yaml:"host"`
	Port            string   `yaml:"port"`
	ReadTimeoutSec  uint8    `yaml:"read_timeout_sec"`
	WriteTimeoutSec uint8    `yaml:"write_timeout_sec"`
	IdleTimeoutSec  uint8    `yaml:"idle_timeout_sec"`
	CSRFKey         string   `yaml:"csrf_key"`
	CSRFMaxAgeSec   int      `yaml:"csrf_max_age_sec"`
	TLS             *HTTPTLS `yaml:"tls"`
}

func (s HTTPServer) Addr() string {
	return s.Host + ":" + s.Port
}

type HTTPTLS struct {
	KeyFilepath  string         `yaml:"key_filepath"`
	CertFilepath string         `yaml:"cert_filepath"`
	MinVersion   *uint16        `yaml:"min_version"`
	MaxVersion   *uint16        `yaml:"max_version"`
	Curves       *[]tls.CurveID `yaml:"curves"`
	Ciphers      *[]uint16      `yaml:"ciphers"`
}
