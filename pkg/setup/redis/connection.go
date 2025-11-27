package redis

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

type Config struct {
	Addr        string
	Password    string
	Database    int
	PoolSize    int
	PoolTimeout time.Duration
}

type Client struct {
	Conn *redis.Client
}

var (
	redisInstance *Client
	redisOnce     sync.Once
)

func NewRedis(cfg Config) (*Client, error) {
	var initError error

	redisOnce.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:         cfg.Addr,
			Password:     cfg.Password,
			DB:           cfg.Database,
			PoolSize:     100, // 500 connections per instance
			PoolTimeout:  120 * time.Second,
			MinIdleConns: 10,
		})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := client.Ping(ctx).Err(); err != nil {
			initError = fmt.Errorf("[REDIS] failed to connect to Redis: %w", err)
			return
		}

		log.Println("[REDIS] connected")
		redisInstance = &Client{Conn: client}
	})

	if initError != nil {
		return nil, initError
	}

	return redisInstance, nil
}

func (c *Client) Close() error {
	if c.Conn == nil {
		return nil
	}

	if err := c.Conn.Close(); err != nil {
		return fmt.Errorf("[REDIS] failed to close connection: %w", err)
	}

	log.Println("[REDIS] connection closed.")
	return nil
}
