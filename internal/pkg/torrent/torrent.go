package torrent

import (
	"encoding/json"
	"log"

	"github.com/anacrolix/torrent"
)

type Torrent struct {
	torrent *torrent.Torrent
	meta    TorrentMeta
}

func (t Torrent) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"infoHash": t.torrent.InfoHash().HexString(),
	})
}

type TorrentService struct {
	TorrentMetaService TorrentMetaService
	TorrentClient      *torrent.Client
}

func NewClient(clientConfig *torrent.ClientConfig) (*torrent.Client, error) {
	return torrent.NewClient(clientConfig)
}

func (s TorrentService) AddFromFile(path string) (Torrent, error) {
	// Add to client
	torrent, err := s.TorrentClient.AddTorrentFromFile(path)
	if err != nil {
		return Torrent{}, err
	}
	<-torrent.GotInfo()

	// Get metadata
	meta, err := s.TorrentMetaService.GetOrCreateTorrentMeta(torrent.InfoHash().HexString())
	if err != nil {
		return Torrent{}, err
	}

	return Torrent{torrent: torrent, meta: meta}, nil
}

func (s TorrentService) AddFromMagnet(magnet string) (Torrent, error) {
	// Add to client
	torrent, err := s.TorrentClient.AddMagnet(magnet)
	if err != nil {
		return Torrent{}, err
	}
	<-torrent.GotInfo()

	// Get metadata
	meta, err := s.TorrentMetaService.GetOrCreateTorrentMeta(torrent.InfoHash().HexString())
	if err != nil {
		return Torrent{}, err
	}

	return Torrent{torrent: torrent, meta: meta}, nil
}

func (s TorrentService) GetOne(infoHash string) (Torrent, error) {
	log.Fatal("unimplemented")
	return Torrent{}, nil
}

func (s TorrentService) GetAll() ([]Torrent, error) {
	torrents := []Torrent{}

	rawTorrents := s.TorrentClient.Torrents()

	for _, rawTorrent := range rawTorrents {
		torrentMeta, err := s.TorrentMetaService.TorrentMeta(rawTorrent.InfoHash().HexString())
		if err != nil {
			return nil, err
		}
		torrents = append(torrents, Torrent{
			torrent: rawTorrent,
			meta:    torrentMeta,
		})
	}

	return torrents, nil
}
