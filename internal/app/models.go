package app

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type Torrent struct {
	Model
	InfoHash string `json:"infoHash"`
	Title    string `json:"title"`
}

type TorrentService interface {
	Torrents() []Torrent
	CreateTorrent(Torrent) error
}
