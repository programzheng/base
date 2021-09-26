package cache

import (
	"context"
	"fmt"
	"testing"
)

var ctx = context.Background()

func TestSet(t *testing.T) {
	err := Rdb.Set(ctx, "test", "foo", 0).Err()
	if err != nil {
		panic(err)
	}
}

func TestGet(t *testing.T) {
	val, err := Rdb.Get(ctx, "test").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("value:", val)
}
