package types

type Outbox struct {
	Server   *HTTPServer `yaml:"http_server"`
	Postgres *Postgres   `yaml:"postgres"`
	Redis    *Redis      `yaml:"redis"`
}
