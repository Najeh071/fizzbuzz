package db_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"fizz-buzz/internal/db"
	"fizz-buzz/pkg/fizzbuzz"
)

func TestDatabase(t *testing.T) {
	database, err := db.InitDB(":memory:")
	assert.NoError(t, err)
	defer database.Close()

	ctx := context.Background()

	stat, err := database.GetMostFrequent(ctx)
	assert.NoError(t, err)
	assert.Nil(t, stat)

	req1 := fizzbuzz.Request{Int1: 3, Int2: 5, Limit: 15, Str1: "wiv", Str2: "iou"}
	req2 := fizzbuzz.Request{Int1: 2, Int2: 7, Limit: 20, Str1: "foo", Str2: "bar"}

	_ = database.TrackRequest(ctx, req2)
	_ = database.TrackRequest(ctx, req1)
	_ = database.TrackRequest(ctx, req1)

	top, err := database.GetMostFrequent(ctx)
	assert.NoError(t, err)
	assert.Equal(t, req1, top.Parameters)
	assert.Equal(t, int64(2), top.Hits)
}