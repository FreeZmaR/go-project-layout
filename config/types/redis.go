package types

type Redis struct {
	TLS          *RedisTLS `yaml:"tls"`
	Network      string    `yaml:"network"`
	Host         string    `yaml:"host"`
	Port         string    `yaml:"port"`
	User         string    `yaml:"user"`
	Password     string    `yaml:"password"`
	Database     int       `yaml:"database"`
	QueryTimeSer int       `yaml:"query_timeout_sec"`
}

type RedisTLS struct {
	PEMFilePath string `yaml:"pem_filepath"`
	KeyFilePath string `yaml:"key_filepath"`
	CAFilePath  string `yaml:"ca_filepath"`
	SNI         string `yaml:"sni"`
}
