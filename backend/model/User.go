package model

type User struct {
	Username     string `form:"username" redis:"username"`
	Age          int    `form:"age" redis:"age"`
	Password     string `form:"password" redis:"password"`
	ProfilePhoto string `form:"profile_pic" redis:"profile_pic"`
}

type Profile struct {
	Username        string
	ProfilePhoto    string
	RecentlyWatched []string
	Playlist        map[string][]string
}

// type PlaylistMovies struct {
// 	Movies []string
// }

type Playlist struct {
	Username string
	Name     string
	Movies   string
}
