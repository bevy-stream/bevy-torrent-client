package torrent

import (
	"encoding/json"

	"github.com/anacrolix/torrent"
)

type Torrent struct {
	torrent *torrent.Torrent
	meta    TorrentMeta
}

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
		BytesCompleted int64 `json:"bytesCompleted"`
		BytesMissing   int64 `json:"bytesMissing"`
	}

	type JSONOutput struct {
		TorrentMeta
		Stats Stats `json:"stats"`
	}

	return json.Marshal(JSONOutput{t.meta, Stats{
		BytesCompleted: t.torrent.BytesCompleted(),
		BytesMissing:   t.torrent.BytesMissing(),
	}})
}

func NewClient(clientConfig *torrent.ClientConfig) (*torrent.Client, error) {
	return torrent.NewClient(clientConfig)
}
