package fizzbuzz_test

import (
	"fizz-buzz/pkg/fizzbuzz"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCompute_Success(t *testing.T) {
	req := fizzbuzz.Request{
		Int1:  3,
		Int2:  5,
		Limit: 15,
		Str1:  "fizz",
		Str2:  "buzz",
	}

	expected := []string{
		"1", "2", "fizz", "4", "buzz",
		"fizz", "7", "8", "fizz", "buzz",
		"11", "fizz", "13", "14", "fizzbuzz",
	}

	res, err := fizzbuzz.Compute(req)

	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}

func TestCompute_ValidationErrors(t *testing.T) {
	tests := []struct {
		name string
		req  fizzbuzz.Request
	}{
		{
			name: "zero or negative int1",
			req:  fizzbuzz.Request{Int1: 0, Int2: 5, Limit: 10, Str1: "fizz", Str2: "buzz"},
		},
		{
			name: "negative int2",
			req:  fizzbuzz.Request{Int1: 3, Int2: -2, Limit: 10, Str1: "fizz", Str2: "buzz"},
		},
		{
			name: "zero limit",
			req:  fizzbuzz.Request{Int1: 3, Int2: 5, Limit: 0, Str1: "fizz", Str2: "buzz"},
		},
		{
			name: "limit exceeding max threshold",
			req:  fizzbuzz.Request{Int1: 3, Int2: 5, Limit: 100001, Str1: "fizz", Str2: "buzz"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := fizzbuzz.Compute(tt.req)
			assert.Error(t, err)
			assert.Nil(t, res)
		})
	}
}
