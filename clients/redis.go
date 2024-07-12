package clients

import (
	"context"
	"encoding/json"
	"time"

	redisotel "github.com/go-redis/redis/extra/redisotel/v8"
	redis "github.com/go-redis/redis/v8"
)

type Ints []int

func (s Ints) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s Ints) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &s)
}

// Global
var RedisClient *redis.Client

type RedisConfig struct {
	Address  string
	MaxRetry int
	Password string
}

func NewRedisClient(config *RedisConfig) (*redis.Client, error) {

	redisClient := redis.NewClient(&redis.Options{
		Addr:       config.Address,
		MaxRetries: config.MaxRetry,
		Password:   config.Password,
	})
	redisClient.AddHook(redisotel.NewTracingHook())

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*5))
	defer cancel()

	err := redisClient.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	return redisClient, nil
}
