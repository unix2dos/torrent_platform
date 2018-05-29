package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/bencode"
	"github.com/anacrolix/torrent/metainfo"

	"torrent_platform/base"
)

func FileSeed(path string, client *torrent.Client) (infohash string, err error) {

	mi, err := GenerateMetaInfo(path)
	if err != nil {
		return "", err
	}

	// check infohash is exist
	_, new := client.AddTorrentInfoHashWithStorage(mi.HashInfoBytes(), nil)
	if !new {
		return "", base.ErrTorrentAlreadyExist
	}

	// add torrent
	tTorrent, err := client.AddTorrent(mi)
	if err != nil {
		return "", err
	}

	tTorrent.Info().Name = path
	go func() {
		<-tTorrent.GotInfo()
		tTorrent.DownloadAll()
	}()

	return mi.HashInfoBytes().String(), nil
}

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
