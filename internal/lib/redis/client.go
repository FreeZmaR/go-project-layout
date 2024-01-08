package redis

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/FreeZmaR/go-service-structure/template/config/types"
	"time"

	rd "github.com/go-redis/redis/v8"
)

const DefaultTimeout = 5 * time.Second

type Client interface {
	Connect
	Close() error
}

type clientInstance struct {
	client  *rd.Client
	timeout time.Duration
}

func NewClient(cfg *types.Redis) (Client, error) {
	options := &rd.Options{
		Network:  cfg.Network,
		Addr:     cfg.Host + ":" + cfg.Port,
		DB:       cfg.Database,
		Username: cfg.User,
		Password: cfg.Password,
	}

	if cfg.TLS != nil {
		tlsCfg, err := makeTLSConfig(cfg.TLS)
		if err != nil {
			return nil, err
		}

		options.TLSConfig = tlsCfg
	}

	timeout := DefaultTimeout
	if cfg.QueryTimeSer > 0 {
		timeout = time.Duration(cfg.QueryTimeSer) * time.Second
	}

	client := rd.NewClient(options)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	status := client.Ping(ctx)
	if status.Err() != nil {
		return nil, status.Err()
	}

	return &clientInstance{client: client, timeout: timeout}, nil
}

func (c *clientInstance) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	if nil == ctx || ctx == context.Background() {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), c.timeout)
		defer cancel()
	}

	_, err := c.client.WithContext(ctx).Set(ctx, key, value, expiration).Result()
	if err != nil {
		return newBadRequest(err)
	}

	return nil
}

func (c *clientInstance) SetJSON(ctx context.Context, key string, dst any, expiration time.Duration) error {
	buf := bytes.Buffer{}
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)

	err := encoder.Encode(dst)
	if err != nil {
		return newJSONError(err)
	}

	return c.Set(ctx, key, buf.Bytes(), expiration)
}

func (c *clientInstance) Get(ctx context.Context, key string) (string, error) {
	if nil == ctx || ctx == context.Background() {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), c.timeout)
		defer cancel()
	}

	data, err := c.client.WithContext(ctx).Get(ctx, key).Result()
	if err != nil {
		return "", newBadRequest(err)
	}

	return data, nil
}

func (c *clientInstance) Close() error {
	return c.client.Close()
}
