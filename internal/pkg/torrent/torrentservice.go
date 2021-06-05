package torrent

import (
	"errors"

	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/metainfo"
)

type TorrentService struct {
	TorrentMetaService TorrentMetaService
	TorrentClient      *torrent.Client
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

	t := Torrent{torrent: torrent, meta: meta}
	t.sync()

	return t, nil
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

	t := Torrent{torrent: torrent, meta: meta}
	t.sync()

	return t, nil
}

func (s TorrentService) GetOne(infoHash string) (Torrent, error) {
	torrent, found := s.TorrentClient.Torrent(metainfo.NewHashFromHex(infoHash))
	if !found {
		return Torrent{}, errors.New("torrent not found")
	}
	meta, err := s.TorrentMetaService.TorrentMeta(infoHash)
	if err != nil {
		return Torrent{}, err
	}
	return Torrent{torrent, meta}, nil
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

func (s TorrentService) Update(meta TorrentMeta) (Torrent, error) {
	torrent, found := s.TorrentClient.Torrent(metainfo.NewHashFromHex(meta.InfoHash))
	if !found {
		return Torrent{}, errors.New("torrent not found")
	}

	meta, err := s.TorrentMetaService.Update(meta)
	if err != nil {
		return Torrent{}, err
	}

	t := Torrent{torrent: torrent, meta: meta}
	t.sync()

	return t, nil
}
