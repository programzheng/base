package cache

import (
	_ "base/config"
	"log"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var Rdb = redis.NewClient(&redis.Options{
	Addr:     viper.Get("REDIS_ADDR").(string),
	Password: viper.Get("REDIS_PASSWORD").(string),
	DB:       getDb(),
})

func getDb() int {
	db, err := strconv.Atoi(viper.Get("REDIS_DB").(string))
	if err != nil {
		log.Fatalf("cache redis getDb error:%v", err)
	}
	return db
}
