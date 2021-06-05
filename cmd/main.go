package main

import (
	"log"
	"net/http"

	anacrolix "github.com/anacrolix/torrent"
	"github.com/bevy-stream/bevy-torrent-client/internal/pkg/torrent"
	"github.com/gin-gonic/gin"
)

type createTorrentForm struct {
	Magnet string `json:"magnet"`
	File   string `json:"file"`
}

func setupRouter(torrentService torrent.TorrentService) *gin.Engine {
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// List all torrents
	r.GET("/torrents", func(c *gin.Context) {
		torrents, err := torrentService.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"torrents": torrents})
	})

	// Create new torrent
	r.POST("/torrents", func(c *gin.Context) {
		form := createTorrentForm{}
		if err := c.ShouldBind(&form); err != nil {
			log.Printf("Invalid Body: \n%s", err)
			c.String(http.StatusBadRequest, "Invalid body")
			return
		}
		if form.File == "" && form.Magnet == "" {
			log.Println("Invalid Body: file or magnet must exist")
			c.String(http.StatusBadRequest, "Invalid body")
			return
		}

		if form.Magnet != "" {
			torrent, err := torrentService.AddFromMagnet(form.Magnet)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"torrent": torrent})
			return
		}

		if form.File != "" {
			torrent, err := torrentService.AddFromFile(form.File)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"torrent": torrent})
			return
		}

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
	db, err := torrent.InitDB("torrents/bevy.db", true)
	if err != nil {
		panic("failed to connect database")
	}
	torrentMetaService := torrent.TorrentMetaService{
		DB: db,
	}

	clientConfig := anacrolix.NewDefaultClientConfig()
	clientConfig.DataDir = "torrents"
	c, err := torrent.NewClient(clientConfig)
	defer c.Close()

	torrentService := torrent.TorrentService{
		TorrentMetaService: torrentMetaService,
		TorrentClient:      c,
	}

	r := setupRouter(torrentService)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
