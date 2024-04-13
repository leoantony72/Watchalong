package repo

import "watchalong/model"

func Login(user model.User) (bool, error) {
	key := "user:" + user.Username
	result, err := r.HGetAll(ctx, key).Result()
	if err != nil {
		return false, err
	}

	if result["password"] == user.Password {
		return true, nil
	}

	return false, nil
}