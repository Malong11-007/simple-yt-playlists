package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PlaylistType string

const (
	Youtube    PlaylistType = "youtube"
	Spotify    PlaylistType = "spotify"
	SoundCloud PlaylistType = "SoundCloud"
)

type playlist struct {
	Id           string       `json:"id"`
	Name         string       `json:"name" validate:"required"`
	CreatedAt    time.Time    `json:"createdAt"`
	PlaylistType PlaylistType `json:"playlistType"`
	Links        []string     `json:"link"`
}

var p = []playlist{
	{
		Id:           "abc123",
		Name:         "Relaxing Vibes",
		CreatedAt:    time.Date(2022, 4, 15, 12, 30, 0, 0, time.UTC),
		PlaylistType: Youtube,
		Links: []string{
			"https://youtube.com/track1",
			"https://youtube.com/track2",
			"https://youtube.com/track3",
		},
	},
}

func getPlaylists(context *gin.Context) {
	fmt.Println("getPlaylists: First request....")
	context.IndentedJSON(http.StatusOK, p)
}

func addPlaylist(context *gin.Context) {
	var newPlaylist playlist
	if err := context.BindJSON(&newPlaylist); err != nil {
		context.IndentedJSON(http.StatusBadRequest, err)
	}

	p = append(p, newPlaylist)
	context.IndentedJSON(http.StatusCreated, newPlaylist)
}

func main() {
	router := gin.Default()
	router.GET("/playlists", getPlaylists)
	router.POST("/playlists", addPlaylist)
	router.Run("localhost:5000")
}
