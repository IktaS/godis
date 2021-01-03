package short

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// RedisRepo defines what's needed in a redis repository
type RedisRepo struct {
	rdb *redis.Client
	ctx context.Context
}

// Init initialize RedisRepo
func (r *RedisRepo) Init() error {
	r.ctx = context.Background()
	r.rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return nil
}

// Save saves a link data to redis database
func (r *RedisRepo) Save(l *Link) error {
	err := r.rdb.Set(r.ctx, l.key, l.val, 0).Err()
	return err
}

// Get gets a data from redis database
func (r *RedisRepo) Get(key string) (string, error) {
	val, err := r.rdb.Get(r.ctx, key).Result()
	return val, err
}
