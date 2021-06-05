package torrent

import (
	"encoding/json"
	"errors"

	"github.com/anacrolix/torrent"
)

type Torrent struct {
	torrent *torrent.Torrent
	meta    TorrentMeta
}

// sync is an idempotent function that brings a torrent into sync with its metadata
func (t Torrent) sync() {
	if t.meta.IsDownloading {
		t.torrent.DownloadAll()
	} else {
		t.torrent.CancelPieces(0, t.torrent.NumPieces())
	}
	if t.meta.IsUploading {
		t.torrent.AllowDataUpload()
	} else {
		t.torrent.DisallowDataUpload()
	}
}

func (t Torrent) MarshalJSON() ([]byte, error) {
	type Stats struct {
		BytesCompleted int64    `json:"bytesCompleted"`
		BytesMissing   int64    `json:"bytesMissing"`
		Files          []string `json:"files"`
		Peers          int      `json:"peers"`
	}

	type JSONOutput struct {
		TorrentMeta
		Stats Stats `json:"stats"`
	}

	files := []string{}
	for _, file := range t.torrent.Files() {
		files = append(files, file.DisplayPath())
	}

	return json.Marshal(JSONOutput{t.meta, Stats{
		BytesCompleted: t.torrent.BytesCompleted(),
		BytesMissing:   t.torrent.BytesMissing(),
		Files:          files,
		Peers:          len(t.torrent.PeerConns()),
	}})
}

func (t Torrent) GetFile(index int) (*torrent.File, error) {
	if index > len(t.torrent.Files()) {
		return nil, errors.New("index out of range")
	}
	return t.torrent.Files()[index], nil
}

func NewClient(clientConfig *torrent.ClientConfig) (*torrent.Client, error) {
	return torrent.NewClient(clientConfig)
}
