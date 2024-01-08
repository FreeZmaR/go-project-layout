package postgres

import (
	"context"
)

type Client interface {
	Connect
	Close() error
}

type Connect interface {
	Exec(ctx context.Context, sql string, arguments ...any) error
	Query(ctx context.Context, sql string, optionsAndArgs ...any) (Rows, error)
	QueryRow(ctx context.Context, sql string, optionsAndArgs ...any) Row
}

type Rows interface {
	Row
	Next() bool
	Close()
	Err() error
}

type Row interface {
	Scan(dest ...any) error
}
