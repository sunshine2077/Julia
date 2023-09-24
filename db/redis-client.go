package db

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

// Redis Client
var RedisClient *redis.Client

// Connect to redis
func ClientToRedis(v *viper.Viper) error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:    fmt.Sprintf("%s:%d", (*v).GetString("redis1.host"), (*v).GetInt("redis1.port")),
		DB:      (*v).GetInt("redis1.db"),
		Network: (*v).GetString("redis1.network"),
		// 最大连接数：4倍CPU核心数
		PoolSize:     15,
		MinIdleConns: 10,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
