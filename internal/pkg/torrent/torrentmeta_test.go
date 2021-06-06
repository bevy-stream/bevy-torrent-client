package torrent

import (
	"os/exec"
	"reflect"
	"testing"
	"time"

	"gorm.io/gorm"
)

func cleanup() {
	cmd := exec.Command("rm", ".test_bevy.db")
	cmd.Run()
}

func TestUpdate(t *testing.T) {
	defer cleanup()

	db, err := InitDB(".test_bevy.db", true)
	if err != nil {
		panic("failed to connect database")
	}

	s := TorrentMetaService{
		DB: db,
	}

	createInput := TorrentMeta{
		InfoHash: "hash1",
		IsPaused: true,
	}
	createOutput := TorrentMeta{
		Model: gorm.Model{
			ID: 1,
		},
		InfoHash: "hash1",
		IsPaused: true,
	}

	updateInput := TorrentMeta{
		Model:    gorm.Model{},
		InfoHash: "hash1",
		IsPaused: false,
	}

	getInput := "hash1"
	getOutput := TorrentMeta{
		Model: gorm.Model{
			ID: 1,
		},
		InfoHash: "hash1",
		IsPaused: false,
	}

	meta, err := s.CreateTorrentMeta(createInput)
	if err != nil {
		t.Fatalf("Unexpected error %s", err)
	}
	meta.CreatedAt = time.Time{}
	meta.UpdatedAt = time.Time{}
	if !reflect.DeepEqual(meta, createOutput) {
		t.Fatalf("Expected:\n %+v\nGot\n%+v", createOutput, meta)
	}

	_, err = s.Update(updateInput)
	if err != nil {
		t.Fatalf("Unexpected error %s", err)
	}

	meta, err = s.TorrentMeta(getInput)
	if err != nil {
		t.Fatalf("Unexpected error %s", err)
	}
	meta.CreatedAt = time.Time{}
	meta.UpdatedAt = time.Time{}
	if !reflect.DeepEqual(meta, getOutput) {
		t.Fatalf("Discrepancy in Get. Expected:\n %+v\nGot\n%+v", getOutput, meta)
	}
}
