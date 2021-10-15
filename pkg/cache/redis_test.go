package cache

import (
	"context"
	"log"
	"testing"
	"time"
)

var ctx = context.Background()

func TestSetString(t *testing.T) {
	err := Rdb.Set(ctx, "test", "foo", 0).Err()
	if err != nil {
		panic(err)
	}
}

func TestGetString(t *testing.T) {
	val, err := Rdb.Get(ctx, "test").Result()
	if err != nil {
		panic(err)
	}
	log.Fatalf("value:%v\n", val)
}

func TestSAdd(t *testing.T) {
	err := Rdb.SAdd(ctx, "set", 2, 1).Err()
	if err != nil {
		panic(err)
	}
	err = Rdb.Expire(ctx, "set", 5*time.Second).Err()
	if err != nil {
		panic(err)
	}
}

func TestHSet(t *testing.T) {
	err := Rdb.HSet(ctx, "test001", "action", "add").Err()
	if err != nil {
		panic(err)
	}
}

func TestHGet(t *testing.T) {
	action, err := Rdb.HGet(ctx, "test001", "action").Result()
	if err != nil {
		panic(err)
	}
	log.Fatalf("action:%v", action)
}

func TestSGet(t *testing.T) {
	err := Rdb.SAdd(ctx, "set", "foo", 2, 1).Err()
	if err != nil {
		panic(err)
	}
	err = Rdb.Expire(ctx, "set", 5*time.Second).Err()
	if err != nil {
		panic(err)
	}
}

func TestSMember(t *testing.T) {
	es, err := Rdb.SMembers(ctx, "set").Result()
	if err != nil {
		panic(err)
	}
	log.Fatalf("%v\n", es)
}
