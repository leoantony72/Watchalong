package repo

import (
	"fmt"
	"watchalong/model"

	"github.com/redis/go-redis/v9"
)

func CreatePlaylists(playlist model.Playlist) {

	fmt.Println("playlist", playlist)

	key := "user:" + playlist.Username + ":playlist"
	r.ZAdd(ctx, key, redis.Z{Member: playlist.Name})

	key = "user:" + playlist.Username + ":playlist:" + playlist.Name
	r.ZAdd(ctx, key, redis.Z{Member: playlist.Movies})

}
