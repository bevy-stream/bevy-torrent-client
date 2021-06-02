package sqlite

import (
	"github.com/bevy-stream/bevy-torrent-client/internal/app"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TorrentService struct {
	DB *gorm.DB
}

func InitDB(path string, autoMigrate bool) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate schemas
	if autoMigrate {
		db.AutoMigrate(&app.Torrent{})
	}

	return db
}

func (s *TorrentService) Torrents() ([]app.Torrent, error) {
	var torrents []app.Torrent
	result := s.DB.Find(&torrents)
	return torrents, result.Error
}

func (s *TorrentService) CreateTorrent(torrent app.Torrent) (app.Torrent, error) {
	result := s.DB.Create(&torrent)
	return torrent, result.Error
}
