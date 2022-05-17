package cache

import (
	"crypto/tls"
	"log"
	"strconv"

	"github.com/programzheng/base/config"
	_ "github.com/programzheng/base/config"
	"github.com/programzheng/base/pkg/helper"

	"github.com/go-redis/redis/v8"
)

func GetRedisClient() *redis.Client {
	var client = redis.NewClient(&redis.Options{
		Addr:      config.Cfg.GetString("REDIS_ADDR"),
		TLSConfig: getTLSConfig(),
		Password:  config.Cfg.GetString("REDIS_PASSWORD"),
		DB:        getDb(),
	})

	return client
}

func getTLSConfig() *tls.Config {
	tlsBool := helper.ConvertToBool(config.Cfg.GetString("REDIS_TLS"))
	if tlsBool {
		tlsConfig := &tls.Config{}
		if helper.ConvertToBool(config.Cfg.GetString("REDIS_TLS_SKIP_VERIFY")) {
			tlsConfig.InsecureSkipVerify = true
			return tlsConfig
		}
		tlsConfig.MinVersion = tls.VersionTLS12
		return tlsConfig
	}
	return nil
}

func getDb() int {
	db, err := strconv.Atoi(config.Cfg.GetString("REDIS_DB"))
	if err != nil {
		log.Fatalf("cache redis getDb error:%v", err)
	}
	return db
}
