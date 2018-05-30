package engine

import (
	"log"
	"os"
	"path/filepath"

	"github.com/anacrolix/torrent/bencode"
	"github.com/anacrolix/torrent/metainfo"
)

func GenerateMetaInfo(dir string) (mi *metainfo.MetaInfo, err error) {

	baseSize := int64(1 << 14)       //16Kb
	multi := SizeSum(dir) / baseSize //16kb倍数
	if multi == 0 {
		multi = 1
	} else if multi > 2200 { //如果文件过大, 那么块数量最好在1200-2200之间
		multi = multi / 1700
	}
	pieceLength := baseSize * multi

	mi = &metainfo.MetaInfo{}
	mi.SetDefaults()
	info := metainfo.Info{
		PieceLength: pieceLength,
	}
	err = info.BuildFromFilePath(dir)
	log.Printf("PieceLength=%d mediaSize=%d multi=%d numPieces=%d\n", pieceLength, SizeSum(dir), multi, info.NumPieces())
	if err != nil {
		return nil, err
	}
	mi.InfoBytes, err = bencode.Marshal(info)
	if err != nil {
		return nil, err
	}

	return
}

func SizeSum(folder string) (sizeSum int64) {
	filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsRegular() {
			sizeSum += info.Size()
		}
		return nil
	})
	return
}
