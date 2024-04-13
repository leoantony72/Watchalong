package repo

import (
	"watchalong/model"
)

func RegisterUser(user model.User) {

	key := "user:" + user.Username
	r.HSet(ctx, key, &user)
}
