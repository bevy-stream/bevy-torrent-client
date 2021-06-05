package main

import (
	anacrolix "github.com/anacrolix/torrent"
	"github.com/bevy-stream/bevy-torrent-client/internal/pkg/http"
	"github.com/bevy-stream/bevy-torrent-client/internal/pkg/torrent"
)

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

	r := http.NewRouter(torrentService)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
