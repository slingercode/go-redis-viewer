package pkg

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var ctx context.Context

func ConnectRedis(URI string, PORT string) *redis.Client {
	ctx = context.Background()

	rclient := redis.NewClient(&redis.Options{
		Addr: URI + ":" + PORT,
		Password: "",
	})

	err := rclient.Ping(ctx).Err(); if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return rclient
}

func GetAllKeys(rclient *redis.Client) []string {
	res, err := rclient.Keys(ctx, "*").Result(); if err != nil {
		log.Fatal(err)
	}

	return res
}

func GetKeyValue(rclient *redis.Client, key string) string {
	res, err := rclient.Get(ctx, key).Result(); if err != nil {
		log.Fatal(err)
	}

	return res
}
