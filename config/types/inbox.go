package types

type Inbox struct {
	Server   *HTTPServer `yaml:"http_server"`
	Postgres *Postgres   `yaml:"postgres"`
}
