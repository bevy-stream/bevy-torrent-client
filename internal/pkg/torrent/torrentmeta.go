package torrent

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TorrentMeta struct {
	gorm.Model
	InfoHash      string `json:"infoHash" gorm:"uniqueIndex"`
	IsUploading   bool   `json:"isUploading"`
	IsDownloading bool   `json:"isDownloading"`
}

func defaultTorrentMeta() TorrentMeta {
	return TorrentMeta{
		Model:         gorm.Model{},
		InfoHash:      "",
		IsUploading:   false,
		IsDownloading: false,
	}
}

type TorrentMetaService struct {
	DB *gorm.DB
}

func InitDB(path string, autoMigrate bool) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate schemas
	if autoMigrate {
		db.AutoMigrate(&TorrentMeta{})
	}

	return db, nil
}

func (s *TorrentMetaService) TorrentMetas() ([]TorrentMeta, error) {
	var torrents []TorrentMeta
	result := s.DB.Find(&torrents)
	return torrents, result.Error
}

func (s *TorrentMetaService) TorrentMeta(infoHash string) (TorrentMeta, error) {
	var torrentMeta TorrentMeta
	result := s.DB.First(&torrentMeta, "info_hash = ?", infoHash)
	return torrentMeta, result.Error
}

func (s *TorrentMetaService) GetOrCreateTorrentMeta(infoHash string) (TorrentMeta, error) {
	var meta TorrentMeta
	meta, err := s.TorrentMeta(infoHash)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			meta := defaultTorrentMeta()
			meta.InfoHash = infoHash
			meta, err := s.CreateTorrentMeta(meta)
			if err != nil {
				return meta, err
			}
		} else {
			return meta, err
		}
		return meta, nil
	}
	return meta, nil
}

func (s *TorrentMetaService) CreateTorrentMeta(torrentMeta TorrentMeta) (TorrentMeta, error) {
	result := s.DB.Create(&torrentMeta)
	return torrentMeta, result.Error
}

func (s *TorrentMetaService) Update(meta TorrentMeta) (TorrentMeta, error) {
	result := s.DB.Model(&meta).Where("info_hash = ?", meta.InfoHash).Updates(meta)
	return meta, result.Error
}
