package config

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func StartRedis() *redis.Client {
	//redis://default:NJyAwKlsk74nEexqpAHRnCsXLoJmH2Uq@redis-15751.c100.us-east-1-4.ec2.cloud.redislabs.com:15751
	redis := redis.NewClient(&redis.Options{
		Addr:     "redis-15751.c100.us-east-1-4.ec2.cloud.redislabs.com:15751",
		Password: "NJyAwKlsk74nEexqpAHRnCsXLoJmH2Uq",
		DB:       0,
	})

	g := redis.Ping(ctx)
	fmt.Println("redis client:", g)
	return redis
}
