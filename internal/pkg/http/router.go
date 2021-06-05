package http

import (
	"net/http"

	"github.com/bevy-stream/bevy-torrent-client/internal/pkg/torrent"
	"github.com/gin-gonic/gin"
)

func NewRouter(torrentService torrent.TorrentService) *gin.Engine {
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
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := form.validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.Magnet != "" {
			torrent, err := torrentService.AddFromMagnet(form.Magnet)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, torrent)
			return
		}

		if form.File != "" {
			torrent, err := torrentService.AddFromFile(form.File)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, torrent)
			return
		}

	})

	// Get a single torrent
	r.GET("/torrents/:id", func(c *gin.Context) {
		id := c.Param("id")
		torrent, err := torrentService.GetOne(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, torrent)
	})

	// Update a single torrent
	r.PUT("/torrents/:id", func(c *gin.Context) {
		id := c.Param("id")
		meta := torrent.TorrentMeta{}
		if err := c.ShouldBind(&meta); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		meta.InfoHash = id

		torrent, err := torrentService.Update(meta)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, torrent)
	})

	// Delete a single torrent
	r.DELETE("/torrents/:id", func(c *gin.Context) {
		c.String(http.StatusNotImplemented, "Not implemented")
	})

	return r
}
