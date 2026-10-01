package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"fizz-buzz/internal/handler"
)

func TestFizzBuzzHandler_Success(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/fizzbuzz?int1=3&int2=5&limit=15&str1=wi&str2=viou", nil)
	rr := httptest.NewRecorder()

	h := handler.NewHTTPHandler(nil)
	h.HandleFizzBuzzQuery(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var resp map[string][]string
	err := json.NewDecoder(rr.Body).Decode(&resp)
	assert.NoError(t, err)

	expected := []string{
		"1", "2", "wi", "4", "viou",
		"wi", "7", "8", "wi", "viou",
		"11", "wi", "13", "14", "wiviou",
	}
	assert.Equal(t, expected, resp["data"])
}

func TestFizzBuzzHandler_InvalidParams(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/fizzbuzz?int1=0&int2=5&limit=15&str1=fizz&str2=buzz", nil)
	rr := httptest.NewRecorder()

	h := handler.NewHTTPHandler(nil)
	h.HandleFizzBuzzQuery(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}