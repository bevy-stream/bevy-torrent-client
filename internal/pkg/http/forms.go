package http

import "errors"

type createTorrentForm struct {
	Magnet string `json:"magnet"`
	File   string `json:"file"`
}

func (f createTorrentForm) validate() error {
	if f.File == "" && f.Magnet == "" {
		return errors.New("file or magnet must exist")
	}
	return nil
}
