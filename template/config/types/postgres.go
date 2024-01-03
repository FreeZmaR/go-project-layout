package types

type Postgres struct {
	Host                   string `yaml:"host"`
	User                   string `yaml:"user"`
	Password               string `yaml:"password"`
	Database               string `yaml:"database"`
	DefaultQueryTimeoutSec int    `yaml:"default_query_timeout_sec"`
	CloseConnectTimeoutSec int    `yaml:"close_connect_timeout_sec"`
	MakeConnectTimeoutSec  int    `yaml:"make_connect_timeout_sec"`
	Port                   uint16 `yaml:"port"`
}
