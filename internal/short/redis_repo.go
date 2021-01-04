package short

import (
	"context"
	"fmt"
	"regexp"

	"github.com/go-redis/redis/v8"
)

// redisRepo defines what's needed in a redis repository
type redisRepo struct {
	rdb *redis.Client
	rop *redis.Options
	ctx context.Context
}

// NewRedisRepo creates a new redis repository
func NewRedisRepo(ctx context.Context, rop *redis.Options) *Repo {
	var repo Repo
	repo = &redisRepo{
		rdb: nil,
		rop: rop,
		ctx: ctx,
	}
	return &repo
}

// Init initialize RedisRepo
func (r *redisRepo) Init() error {
	r.rdb = redis.NewClient(r.rop)
	return nil
}

// Save saves a link data to redis database
func (r *redisRepo) Save(l *Link) error {
	if isValid, _ := regexp.Match("([^\\|\\[\\]])", []byte(l.Key)); !isValid {
		return fmt.Errorf("Invalid Key")
	}
	if isValid, _ := regexp.Match("([^\\|\\[\\]])", []byte(l.Val)); !isValid {
		return fmt.Errorf("Invalid Value")
	}
	err := r.rdb.Set(r.ctx, l.Key, l.Val, 0).Err()
	return err
}

// Get gets a data from redis database
func (r *redisRepo) Get(key string) (string, error) {
	val, err := r.rdb.Get(r.ctx, key).Result()
	return val, err
}
