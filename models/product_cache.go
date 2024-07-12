package models

import (
	"context"
	"fmt"
	"time"

	redis "github.com/go-redis/redis/v8"
	"github.com/huynhtrongtien/dove/clients"
	"github.com/huynhtrongtien/dove/entities"
)

type ICachedProduct interface {
	Set(ctx context.Context, data *entities.Product) error
	Get(ctx context.Context, id int64) (*entities.Product, error)
	GetByUUID(ctx context.Context, uuid string) (*entities.Product, error)
	Delete(ctx context.Context, id int64) error
}

type CachedProduct struct {
	Prefix     string
	MaxRetry   int
	Expiration time.Duration
}

func (c *CachedProduct) getProductInfoKey(id int64) string {
	return fmt.Sprintf("%s:product:%d:info", c.Prefix, id)
}

func (c *CachedProduct) getProductIDKey(uuid string) string {
	return fmt.Sprintf("%s:product:%s:id", c.Prefix, uuid)
}

// https://redis.uptrace.dev/guide/go-redis-pipelines.html#transactions
func (c CachedProduct) Set(ctx context.Context, data *entities.Product) error {
	infoKey := c.getProductInfoKey(data.ID)
	idKey := c.getProductIDKey(data.UUID)
	txf := func(tx *redis.Tx) error {
		_, err := tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.HMSet(ctx, infoKey, "uuid", data.UUID, "code", data.Code, "fullname", data.FullName, "category_id", data.CategoryID)
			pipe.Expire(ctx, infoKey, c.Expiration)
			pipe.Set(ctx, idKey, data.ID, c.Expiration)
			return nil
		})
		return err
	}

	// retry one time
	for i := 0; i < c.MaxRetry; i++ {
		err := clients.RedisClient.Watch(ctx, txf, infoKey, idKey)
		if err == nil {
			return nil
		}
		if err == redis.TxFailedErr {
			continue
		}
		// Return any other error.
		return err
	}

	return nil
}

func (c CachedProduct) Get(ctx context.Context, id int64) (*entities.Product, error) {

	cmd := clients.RedisClient.HGetAll(ctx, c.getProductInfoKey(id))
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}

	// if conected with server, redis allways return empty
	if len(cmd.Val()) < 1 {
		return nil, redis.Nil
	}

	result := &entities.Product{}
	err := cmd.Scan(result)
	if err != nil {
		return nil, err
	}

	result.ID = id
	return result, nil
}

func (c CachedProduct) GetByUUID(ctx context.Context, uuid string) (*entities.Product, error) {
	result := clients.RedisClient.Get(ctx, c.getProductIDKey(uuid))
	if result.Err() != nil {
		return nil, result.Err()
	}

	id, err := result.Int64()
	if err != nil {
		return nil, err
	}

	return c.Get(ctx, id)
}

func (c CachedProduct) Delete(ctx context.Context, id int64) error {
	var uuidKey string
	infoKey := c.getProductInfoKey(id)

	txf := func(tx *redis.Tx) error {
		uuidValue := tx.HGet(ctx, infoKey, FieldUUID)
		if uuidValue.Err() != nil {
			return uuidValue.Err()
		}

		uuidKey = c.getProductIDKey(uuidValue.Val())
		_, err := tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.Del(ctx, infoKey, uuidKey)
			return nil
		})
		return err
	}

	// watch and retry if need
	for i := 0; i < c.MaxRetry; i++ {
		err := clients.RedisClient.Watch(ctx, txf, infoKey, uuidKey)
		if err == nil {
			return nil
		}
		if err == redis.TxFailedErr {
			continue
		}
		// Return any other error.
		return err
	}

	return nil
}
