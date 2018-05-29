package base

import (
	"log"
	"os"
	"path/filepath"

	"github.com/anacrolix/torrent/storage"
)

func cleanDB(dir string) {
	p := filepath.Join(dir, ".torrent.bolt.db")
	err := os.Remove(p)
	if err != nil {
		log.Printf("rm %s error %s", p, err)
	}
}

func pieceCompletionForDir(dir string) (ret storage.PieceCompletion) {
	cleanDB(dir)
	ret, err := storage.NewBoltPieceCompletion(dir)
	if err != nil {
		log.Printf("couldn't open piece completion db in %q: %s", dir, err)
		ret = storage.NewMapPieceCompletion()
	}
	return
}

// All Torrent data stored in this baseDir
func NewFile(baseDir string, dbDir string) storage.ClientImpl {
	return storage.NewFileWithCompletion(baseDir, pieceCompletionForDir(dbDir))
}
