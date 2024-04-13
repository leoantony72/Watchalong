package repo

import (
	"context"
	"watchalong/config"

	"github.com/redis/go-redis/v9"
)

var r *redis.Client = config.StartRedis()
var ctx = context.Background()
