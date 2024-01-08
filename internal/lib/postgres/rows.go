package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type RowsInstance struct {
	pgx.Rows
	cancelFN context.CancelFunc
}

func (r *RowsInstance) Close() {
	if r.cancelFN != nil {
		defer r.cancelFN()
	}

	r.Rows.Close()
}

type RowInstance struct {
	pgx.Row
	cancelFN context.CancelFunc
	err      error
}

func (r *RowInstance) Scan(dest ...any) error {
	if r.cancelFN != nil {
		defer r.cancelFN()
	}

	if r.err != nil {
		return r.err
	}

	return r.Row.Scan(dest...)
}
