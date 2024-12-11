package provider

import (
	"c_program_edu-gin/internal/config"
	"c_program_edu-gin/pkg/logger"
	"context"
	"fmt"
	"github.com/go-redis/redis"
)

func NewRedisClient(conf *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:        conf.Redis.Host + ":" + conf.Redis.Port,
		Password:    conf.Redis.Auth,
		DB:          conf.Redis.Database,
		ReadTimeout: 0,
	})

	if _, err := client.WithContext(context.TODO()).Ping().Result(); err != nil {
		logger.Panicf("Redis Client Error: %v!", err)
		fmt.Println("Redis Client Error: ", err)
		return nil
	}

	return client
}
