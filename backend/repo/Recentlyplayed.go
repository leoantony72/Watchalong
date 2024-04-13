package repo

import "github.com/redis/go-redis/v9"

func RecentlyPlayed(username, movie string) {
	key := "user:" + username + ":recentlywatched"
	r.ZAdd(ctx, key, redis.Z{Member: movie})
}
