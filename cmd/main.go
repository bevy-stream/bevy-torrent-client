package main

import (
	"net/http"

	"github.com/bevy-stream/bevy-torrent-client/internal/app"
	"github.com/bevy-stream/bevy-torrent-client/internal/pkg/sqlite"
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	db := sqlite.InitDB("/tmp/tmp.db", true)
	torrentService := sqlite.TorrentService{
		DB: db,
	}

	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// List all torrents
	r.GET("/torrents", func(c *gin.Context) {
		torrents, err := torrentService.Torrents()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"torrents": torrents})
	})

	// Create new torrent
	r.POST("/torrents", func(c *gin.Context) {
		torrent := app.Torrent{}
		if err := c.ShouldBind(&torrent); err != nil {
			c.String(http.StatusBadRequest, `the body should be a torrent`)
			return
		}

		torrent, err := torrentService.CreateTorrent(torrent)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"torrent": torrent})
	})

	// Get a single torrent
	r.GET("/torrents/:id", func(c *gin.Context) {
		c.String(http.StatusNotImplemented, "Not implemented")
	})

	// Update a single torrent
	r.PUT("/torrents/:id", func(c *gin.Context) {
		c.String(http.StatusNotImplemented, "Not implemented")
	})

	// Delete a single torrent
	r.DELETE("/torrents/:id", func(c *gin.Context) {
		c.String(http.StatusNotImplemented, "Not implemented")
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
