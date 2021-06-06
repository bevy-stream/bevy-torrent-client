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
	if t.meta.IsPaused {
		t.torrent.CancelPieces(0, t.torrent.NumPieces())
		t.torrent.DisallowDataUpload()
		t.torrent.DisallowDataDownload()
	} else {
		t.torrent.DownloadAll()
		t.torrent.AllowDataUpload()
		t.torrent.AllowDataDownload()
	}
}

func (t Torrent) MarshalJSON() ([]byte, error) {
	type Info struct {
		BytesCompleted int64    `json:"bytesCompleted"`
		BytesMissing   int64    `json:"bytesMissing"`
		Files          []string `json:"files"`
		Peers          int      `json:"peers"`
		Name           string   `json:"name"`
		Length         int64    `json:"length"`
	}

	type JSONOutput struct {
		TorrentMeta
		Stats Info `json:"info"`
	}

	files := []string{}
	for _, file := range t.torrent.Files() {
		files = append(files, file.DisplayPath())
	}

	return json.Marshal(JSONOutput{t.meta, Info{
		BytesCompleted: t.torrent.BytesCompleted(),
		BytesMissing:   t.torrent.BytesMissing(),
		Files:          files,
		Peers:          len(t.torrent.PeerConns()),
		Name:           t.torrent.Name(),
		Length:         t.torrent.Length(),
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
