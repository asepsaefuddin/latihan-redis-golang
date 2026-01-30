package config

import (
	"context"
	"crypto/tls"
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func ConnectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	})
	// root context
	// Root context: context dasar sebagai awal dari semua context
	ctx := context.Background()              // context itu di butuhkan untuk lifecicle di golang
	_, err := RedisClient.Ping(ctx).Result() // dignakan untuk mengecek koneksi
	if err != nil {
		panic("koneksi gagal ke redis")
	}
}
