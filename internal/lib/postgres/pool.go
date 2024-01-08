package postgres

import (
	"context"
	"fmt"
	"github.com/FreeZmaR/go-service-structure/template/config/types"
	"net/url"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Pool struct {
	instance     *pgxpool.Pool
	queryTimeout time.Duration
}

func NewPool(cfg *types.Postgres) (*Pool, error) {
	poolCFG, err := makePoolConfig(cfg)
	if err != nil {
		return nil, err
	}

	defaultQueryTimeout := time.Duration(cfg.DefaultQueryTimeoutSec) * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), defaultQueryTimeout)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, poolCFG)
	if err != nil {
		return nil, err
	}

	instance := &Pool{
		instance:     pool,
		queryTimeout: defaultQueryTimeout,
	}

	if err = instance.ping(); err != nil {
		pool.Close()

		return nil, err
	}

	return instance, nil
}

func (p *Pool) Exec(ctx context.Context, sql string, arguments ...any) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	if nil == ctx || ctx == context.Background() {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), p.queryTimeout)
		defer cancel()
	}

	_, err := p.instance.Exec(ctx, sql, arguments...)

	return err
}

func (p *Pool) Query(ctx context.Context, sql string, optionsAndArgs ...any) (Rows, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	if ctx != nil && ctx != context.Background() {
		return p.instance.Query(ctx, sql, optionsAndArgs...)
	}

	ctx, cancel := context.WithTimeout(context.Background(), p.queryTimeout)
	rows, err := p.instance.Query(ctx, sql, optionsAndArgs...)
	if err != nil {
		cancel()

		return nil, err
	}

	return &RowsInstance{
		Rows:     rows,
		cancelFN: cancel,
	}, nil
}

func (p *Pool) QueryRow(ctx context.Context, sql string, optionsAndArgs ...any) Row {
	if ctx.Err() != nil {
		return &RowInstance{err: ctx.Err()}
	}

	if ctx != nil && ctx != context.Background() {
		return p.instance.QueryRow(ctx, sql, optionsAndArgs...)
	}

	ctx, cancel := context.WithTimeout(context.Background(), p.queryTimeout)

	return &RowInstance{
		Row:      p.instance.QueryRow(ctx, sql, optionsAndArgs...),
		cancelFN: cancel,
	}
}

func (p *Pool) Close() error {
	p.instance.Close()

	return nil
}

func (p *Pool) ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), p.queryTimeout)
	defer cancel()

	return p.instance.Ping(ctx)
}

func makePoolConfig(cfg *types.Postgres) (*pgxpool.Config, error) {
	connStr := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(cfg.User, cfg.Password),
		Host:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Path:     cfg.Database,
		RawQuery: "sslmode=prefer",
	}

	poolCFG, err := pgxpool.ParseConfig(connStr.String())
	if err != nil {
		return nil, err
	}

	poolCFG.MinConns = 1
	if cfg.PoolSize > 0 {
		poolCFG.MaxConns = int32(cfg.PoolSize)
	}

	poolCFG.MaxConnIdleTime = time.Duration(cfg.IdleConnectTimeoutSec) * time.Second

	return poolCFG, nil
}
