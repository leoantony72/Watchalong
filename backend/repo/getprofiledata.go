package repo

import (
	"fmt"
	"watchalong/model"
)

func GetProfileData(username string) (*model.Profile, error) {

	key := "user:" + username
	field := "profile_pic"
	profile_pic, err := r.HGet(ctx, key, field).Result()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	key = "user:" + username + ":recentlywatched"
	result, err := r.ZRange(ctx, key, 0, -1).Result()
	if err != nil {
		fmt.Println("err hiiii: ", err.Error())
		return nil, err
	}

	key = "user:" + username + ":playlist"
	rplaylist, _ := r.ZRange(ctx, key, 0, -1).Result()

	// profiledata := model.Profile{}
	profiledata := model.Profile{
		Playlist: make(map[string][]string),
	}
	for _, str := range rplaylist {
		fmt.Println("keyvalue: ", key)
		key := "user:" + username + ":playlist:" + str
		data, _ := r.ZRange(ctx, key, 0, -1).Result()
		fmt.Println("keyvalue: ", key, data)
		// profiledata.Playlist

		profiledata.Playlist[str] = data
	}

	imgpath := "localhost:8080/img/" + profile_pic
	profiledata.Username = username
	profiledata.ProfilePhoto = imgpath

	profiledata.RecentlyWatched = result
	return &profiledata, nil
}
