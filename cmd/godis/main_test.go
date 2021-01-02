package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

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
				return httptest.NewRequest("GET", "/imam", nil)
			},
			expectedCode: 200,
			expectedBody: "Welcome home imam",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := tt.setup()
			rec := httptest.NewRecorder()
			router := mux.NewRouter()
			router.HandleFunc("/{name}", homeHandler)
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
