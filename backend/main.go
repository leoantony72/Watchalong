package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"watchalong/config"
	"watchalong/model"
	"watchalong/repo"

	// "watchalong/webrtc"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {

	r := gin.Default()
	r.GET("/", serveHome)
	r.LoadHTMLGlob("*.html")
	r.POST("/login", login)
	r.POST("/signup", signup)

	r.GET("/ws", test)

	r.GET("/profile", profile)

	r.POST("/playlist", createPlaylist)

	r.POST("/recentwatched", addToRecentlyPlayed)

	// r.POST("/previews", moviePreview)

	r.Static("/img", "./img/profile")
	r.Static("/video", "./video")
	go config.Send()

	r.Run(":8080")
}

func test(c *gin.Context) {
	// ctx := context.Background()
	config.Wshandler(c.Writer, c.Request, c)

}

func serveHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index1.html", nil)
}

func signup(c *gin.Context) {
	user := model.User{}
	err := c.Bind(&user)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": "enter valid parameters"})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": "file upload error"})
		return
	}

	extension := filepath.Ext(file.Filename)
	filename := uuid.New().String() + extension
	var dst string = "img/profile/"
	err = c.SaveUploadedFile(file, dst+filename)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": "file upload error"})
		return
	}

	fmt.Println(user, file.Filename)

	user.ProfilePhoto = filename

	repo.RegisterUser(user)
	c.JSON(200, gin.H{"message": "Register success"})

}

func login(c *gin.Context) {
	user := model.User{}
	err := c.Bind(&user)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": "enter valid parameters"})
		return
	}

	if user.Username == "" || user.Password == "" {
		c.JSON(400, gin.H{"message": "Invalid username/password"})
		return
	}
	valid, err := repo.Login(user)
	if err != nil {
		fmt.Println("login_err: ", err.Error())
		c.JSON(400, gin.H{"message": "something went wrong"})
		return
	}

	if !valid {
		c.JSON(400, gin.H{"message": "invalid username/password"})
		return
	}

	c.SetCookie("session", user.Username, 3600*24, "/", "localhost", false, true)

	c.JSON(200, gin.H{"message": "Login successful"})

}

func profile(c *gin.Context) {
	username := c.Query("username")

	if username == "" {
		c.JSON(400, gin.H{"message": "Username must be specified"})
		return
	}
	data, err := repo.GetProfileData(username)
	if err != nil {
		fmt.Println("profile_err", err)
		c.JSON(400, gin.H{"message": "something went wrong"})
		return
	}
	c.JSON(200, gin.H{"message": data})
}

func addToRecentlyPlayed(c *gin.Context) {
	movie := c.Query("movie")
	username, err := c.Cookie("session")

	if username == "" || err != nil {
		c.JSON(400, gin.H{"message": "Authentication is required"})
	}
	if movie == "" {
		c.JSON(400, gin.H{"message": "movie name is required to add to recently played"})
		return
	}

	repo.RecentlyPlayed(username, movie)
	c.JSON(200, gin.H{"message": "added to recently played"})
}

func createPlaylist(c *gin.Context) {
	playlistName := c.Query("playlist")
	movie := c.Query("movie")

	fmt.Println(movie, playlistName)
	if movie == "" {
		c.JSON(400, gin.H{"message": "movie name not specified"})
		return
	}

	username, err := c.Cookie("session")
	if err != nil {
		c.JSON(400, gin.H{"message": "Authorization required"})
		return
	}

	if username == "" {
		c.JSON(200, gin.H{"message": "username not available"})
		return
	}

	playlists := model.Playlist{}
	playlists.Username = username

	playlists.Name = playlistName
	playlists.Movies = movie
	repo.CreatePlaylists(playlists)

	c.JSON(200, gin.H{"message": "Movie added to playlist"})
}
func home(c *gin.Context) {
	c.JSON(200, gin.H{"message": "lloyd"})
}

// func moviePreview(c *gin.Context) {
// 	movie := c.Query("movie")



// }
