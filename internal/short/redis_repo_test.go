package short

import (
	"context"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func Test_redisRepo_Save(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*testing.T, *miniredis.Miniredis) *redisRepo
		teardown func(*testing.T, *miniredis.Miniredis)
		link     *Link
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "valid key pair",
			setup: func(t *testing.T, s *miniredis.Miniredis) *redisRepo {
				repo := &redisRepo{
					rdb: nil,
					rop: &redis.Options{
						Addr: s.Addr(),
					},
					ctx: context.Background(),
				}
				repo.rdb = redis.NewClient(repo.rop)
				return repo
			},
			teardown: func(t *testing.T, s *miniredis.Miniredis) {
				s.FlushAll()
				s.Close()
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
			srv, err := miniredis.Run()
			if err != nil {
				t.Fatal(err)
			}
			r := tt.setup(t, srv)
			err = r.Save(tt.link)
			if tt.wantErr {
				assert.Error(t, err)
			}
			tt.teardown(t, srv)
		})
	}
}

func Test_redisRepo_Get(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*testing.T, *miniredis.Miniredis) *redisRepo
		teardown func(*testing.T, *miniredis.Miniredis)
		key      string
		want     string
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "valid key pair",
			setup: func(t *testing.T, s *miniredis.Miniredis) *redisRepo {
				repo := &redisRepo{
					rdb: nil,
					rop: &redis.Options{
						Addr: s.Addr(),
					},
					ctx: context.Background(),
				}
				repo.rdb = redis.NewClient(repo.rop)
				s.Set("key", "val")
				return repo
			},
			teardown: func(t *testing.T, s *miniredis.Miniredis) {
				s.FlushAll()
				s.Close()
				return
			},
			key:     "key",
			want:    "val",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv, err := miniredis.Run()
			if err != nil {
				t.Fatal(err)
			}
			r := tt.setup(t, srv)
			got, err := r.Get(tt.key)
			if tt.wantErr {
				assert.Error(t, err)
			}
			assert.Equal(t, tt.want, got, tt.name)
			tt.teardown(t, srv)
		})
	}
}
