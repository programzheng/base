package cache

import (
	"crypto/tls"
	"log"
	"strconv"

	_ "github.com/programzheng/base/config"
	"github.com/programzheng/base/pkg/helper"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func GetRedisClient() *redis.Client {
	var client = redis.NewClient(&redis.Options{
		Addr:      viper.Get("REDIS_ADDR").(string),
		TLSConfig: getTLSConfig(),
		Password:  viper.Get("REDIS_PASSWORD").(string),
		DB:        getDb(),
	})

	return client
}

func getTLSConfig() *tls.Config {
	tlsBool := helper.ConvertToBool(viper.Get("REDIS_TLS").(string))
	if tlsBool {
		tlsConfig := &tls.Config{}
		if helper.ConvertToBool(viper.Get("REDIS_TLS_SKIP_VERIFY").(string)) {
			tlsConfig.InsecureSkipVerify = true
			return tlsConfig
		}
		tlsConfig.MinVersion = tls.VersionTLS12
		return tlsConfig
	}
	return nil
}

func getDb() int {
	db, err := strconv.Atoi(viper.Get("REDIS_DB").(string))
	if err != nil {
		log.Fatalf("cache redis getDb error:%v", err)
	}
	return db
}
