package cache

import (
	"context"
	"fmt"
	"github.com/TotrazOrstO/Imapro/pkg/config"
	"github.com/go-redis/redis/v8"
)

func InitRedis(cfg config.RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("error connecting to Redis: %v", err)
	}

	return client, nil
}
