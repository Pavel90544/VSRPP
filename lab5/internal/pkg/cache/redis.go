package cache

import (
    "context"
    "encoding/json"
    "errors"
    "time"

    "github.com/redis/go-redis/v9"
)

type RedisCache struct {
    client *redis.Client
    ctx    context.Context
    ttl    time.Duration
}

func NewRedisCache(addr, password string, db int, ttl time.Duration) (*RedisCache, error) {
    client := redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: password,
        DB:       db,
    })

    ctx := context.Background()

    if err := client.Ping(ctx).Err(); err != nil {
        return nil, errors.New("failed to connect to Redis: " + err.Error())
    }

    return &RedisCache{
        client: client,
        ctx:    ctx,
        ttl:    ttl,
    }, nil
}

func (r *RedisCache) Set(key string, value interface{}, ttl time.Duration) error {
    if ttl == 0 {
        ttl = r.ttl
    }

    data, err := json.Marshal(value)
    if err != nil {
        return err
    }

    return r.client.Set(r.ctx, key, data, ttl).Err()
}

func (r *RedisCache) Get(key string, value interface{}) error {
    data, err := r.client.Get(r.ctx, key).Bytes()
    if err != nil {
        if err == redis.Nil {
            return errors.New("key not found")
        }
        return err
    }

    return json.Unmarshal(data, value)
}

func (r *RedisCache) Delete(key string) error {
    return r.client.Del(r.ctx, key).Err()
}

func (r *RedisCache) Clear() error {
    return r.client.FlushDB(r.ctx).Err()
}

func (r *RedisCache) Exists(key string) bool {
    result := r.client.Exists(r.ctx, key)
    return result.Val() > 0
}

func (r *RedisCache) Close() error {
    return r.client.Close()
}
