package db_conn

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"graduation_project/config"
	"sync"
	"time"
)

var redisClient *redis.Client
var redisOnce sync.Once

func InitRedis(cfg config.RedisConfig) (*redis.Client, error) {
	var err error
	redisOnce.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			Password: cfg.Password,
			DB:       cfg.DB,
		})

		// 测试连接
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_, err = redisClient.Ping(ctx).Result()
	})
	return redisClient, err
}

func GetRedis() *redis.Client {
	return redisClient
}
