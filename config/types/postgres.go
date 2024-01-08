package types

type Postgres struct {
	Host                   string `yaml:"host"`
	User                   string `yaml:"user"`
	Password               string `yaml:"password"`
	Database               string `yaml:"database"`
	DefaultQueryTimeoutSec int    `yaml:"default_query_timeout_sec"`
	IdleConnectTimeoutSec  int    `yaml:"idle_connect_timeout_sec"`
	PoolSize               int    `yaml:"pool_size"`
	Port                   uint16 `yaml:"port"`
}
