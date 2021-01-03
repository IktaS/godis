package short

import (
	"context"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func Test_redisRepo_Save(t *testing.T) {
	srv, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}
	defer srv.Close()
	tests := []struct {
		name     string
		setup    func(*testing.T) *redisRepo
		teardown func(*testing.T, *redisRepo)
		link     *Link
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "valid key pair",
			setup: func(*testing.T) *redisRepo {
				repo := &redisRepo{
					rdb: nil,
					rop: &redis.Options{
						Addr: srv.Addr(),
					},
					ctx: context.Background(),
				}
				repo.rdb = redis.NewClient(repo.rop)
				return repo
			},
			teardown: func(t *testing.T, f *redisRepo) {
				return
			},
			link: &Link{
				Key: "key",
				Val: "val",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.setup(t)
			err := r.Save(tt.link)
			if tt.wantErr {
				assert.Error(t, err)
			}
			tt.teardown(t, r)
		})
	}
}

func Test_redisRepo_Get(t *testing.T) {
	srv, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}
	defer srv.Close()
	tests := []struct {
		name     string
		setup    func(*testing.T) *redisRepo
		teardown func(*testing.T, *redisRepo)
		key      string
		want     string
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "valid key pair",
			setup: func(*testing.T) *redisRepo {
				repo := &redisRepo{
					rdb: nil,
					rop: &redis.Options{
						Addr: srv.Addr(),
					},
					ctx: context.Background(),
				}
				repo.rdb = redis.NewClient(repo.rop)
				srv.Set("key", "val")
				return repo
			},
			teardown: func(t *testing.T, f *redisRepo) {
				srv.Del("key")
				return
			},
			key:     "key",
			want:    "val",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.setup(t)
			got, err := r.Get(tt.key)
			if tt.wantErr {
				assert.Error(t, err)
			}
			assert.Equal(t, tt.want, got, tt.name)
			tt.teardown(t, r)
		})
	}
}
