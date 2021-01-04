package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IktaS/godis/internal/short"
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Test_homeHandler(t *testing.T) {
	tests := []struct {
		name         string
		setup        func() *http.Request
		expectedCode int
		expectedBody string
	}{
		// TODO: Add test cases.
		{
			name: "string parameter get request",
			setup: func() *http.Request {
				return httptest.NewRequest("GET", "/", nil)
			},
			expectedCode: 200,
			expectedBody: "Got Home!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := tt.setup()
			rec := httptest.NewRecorder()
			router := mux.NewRouter()
			router.HandleFunc("/", homeHandler)
			router.ServeHTTP(rec, req)

			res := rec.Result()
			defer res.Body.Close()
			assert.Equal(t, tt.expectedCode, res.StatusCode, "Status code should be the same as expected")

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				assert.Error(t, err, "Could not read response")
			}

			assert.Equal(t, tt.expectedBody, string(b))
		})
	}
}

func Test_saveHandler(t *testing.T) {
	tests := []struct {
		name         string
		setup        func(*testing.T, *miniredis.Miniredis) (*short.Repo, *http.Request)
		teardown     func(*testing.T, *miniredis.Miniredis)
		expectedCode int
		expectedBody string
	}{
		// TODO: Add test cases.
		{
			name: "valid request",
			setup: func(t *testing.T, s *miniredis.Miniredis) (*short.Repo, *http.Request) {
				repo := short.NewRedisRepo(context.Background(), &redis.Options{
					Addr: s.Addr(),
				})
				(*repo).Init()
				data := &short.Link{
					Key: "key",
					Val: "val",
				}
				body, err := json.Marshal(data)
				if err != nil {
					t.Fatal(err)
				}
				req := httptest.NewRequest("POST", "/save", bytes.NewBuffer(body))
				return repo, req
			},
			teardown: func(t *testing.T, s *miniredis.Miniredis) {
				s.Close()
				return
			},
			expectedCode: 200,
			expectedBody: "Data saved!",
		},
		{
			name: "valid request",
			setup: func(t *testing.T, s *miniredis.Miniredis) (*short.Repo, *http.Request) {
				repo := short.NewRedisRepo(context.Background(), &redis.Options{
					Addr: s.Addr(),
				})
				(*repo).Init()
				data := &short.Link{
					Key: "",
					Val: "val",
				}
				body, err := json.Marshal(data)
				if err != nil {
					t.Fatal(err)
				}
				req := httptest.NewRequest("POST", "/save", bytes.NewBuffer(body))
				return repo, req
			},
			teardown: func(t *testing.T, s *miniredis.Miniredis) {
				s.FlushAll()
				s.Close()
				return
			},
			expectedCode: 500,
			expectedBody: "Invalid Key\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv, err := miniredis.Run()
			if err != nil {
				t.Fatal(err)
			}
			repo, req := tt.setup(t, srv)
			rec := httptest.NewRecorder()
			router := mux.NewRouter()
			router.HandleFunc("/save", useRepo(saveHandler, repo))
			router.ServeHTTP(rec, req)

			res := rec.Result()
			defer res.Body.Close()
			assert.Equal(t, tt.expectedCode, res.StatusCode, "Status code should be the same as expected")

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.expectedBody, string(b))
			tt.teardown(t, srv)
		})
	}
}

func Test_getHandler(t *testing.T) {
	tests := []struct {
		name         string
		setup        func(*testing.T, *miniredis.Miniredis) (*short.Repo, *http.Request)
		teardown     func(*testing.T, *miniredis.Miniredis)
		expectedCode int
		expectedBody string
		wantErr      bool
	}{
		{
			name: "valid request",
			setup: func(t *testing.T, s *miniredis.Miniredis) (*short.Repo, *http.Request) {
				repo := short.NewRedisRepo(context.Background(), &redis.Options{
					Addr: s.Addr(),
				})
				(*repo).Init()
				s.Set("key", "val")
				req := httptest.NewRequest("GET", "/key", nil)
				return repo, req
			},
			teardown: func(t *testing.T, s *miniredis.Miniredis) {
				s.FlushAll()
				s.Close()
				return
			},
			expectedCode: 200,
			expectedBody: "val",
			wantErr:      false,
		},
		{
			name: "non existent request",
			setup: func(t *testing.T, s *miniredis.Miniredis) (*short.Repo, *http.Request) {
				repo := short.NewRedisRepo(context.Background(), &redis.Options{
					Addr: s.Addr(),
				})
				(*repo).Init()
				s.Set("key", "val")
				req := httptest.NewRequest("GET", "/key2", nil)
				return repo, req
			},
			teardown: func(t *testing.T, s *miniredis.Miniredis) {
				s.FlushAll()
				s.Close()
				return
			},
			expectedCode: 204,
			expectedBody: "",
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv, err := miniredis.Run()
			if err != nil {
				t.Fatal(err)
			}
			repo, req := tt.setup(t, srv)
			rec := httptest.NewRecorder()
			router := mux.NewRouter()
			router.HandleFunc("/{key}", useRepo(getHandler, repo))
			router.ServeHTTP(rec, req)

			res := rec.Result()
			defer res.Body.Close()
			assert.Equal(t, tt.expectedCode, res.StatusCode, "Status code should be the same as expected")

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}
			var got string
			err = json.Unmarshal(b, &got)
			if tt.wantErr {
				assert.Error(t, err)
			}
			assert.Equal(t, tt.expectedBody, got)
			tt.teardown(t, srv)
		})
	}
}
