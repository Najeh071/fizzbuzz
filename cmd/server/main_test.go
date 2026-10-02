package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	router := setupRouter(nil)

	tests := []struct {
		name string
		method string
		path string
		expectedStat int
	} {
		{
			name:           "GET /api/v1/fizzbuzz endpoint bound",
			method:         http.MethodGet,
			path:           "/api/v1/fizzbuzz",
			expectedStat: http.StatusBadRequest,
		},
		{
			name:           "GET /api/v1/stats endpoint bound",
			method:         http.MethodGet,
			path:           "/api/v1/stats",
			expectedStat: http.StatusServiceUnavailable,
		},
		{
			name:           "Unknown endpoint returns 404",
			method:         http.MethodGet,
			path:           "/unknown",
			expectedStat: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func (t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)
			assert.Equal(t, tt.expectedStat, rr.Code)
		})
	}
}