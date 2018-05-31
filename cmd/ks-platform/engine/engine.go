package engine

import (
	"io"

	"github.com/anacrolix/torrent"

	"torrent_platform/base"
)

type Engine struct {
	client *torrent.Client
}

func New() *Engine {
	return &Engine{}
}

func (e *Engine) Configure() error {

	tc := torrent.Config{
		Debug:            false,
		Seed:             true,
		DisableIPv6:      true,
		DhtStartingNodes: base.BootstrapAddrs,
		DefaultStorage:   base.NewFile("", ""),
	}

	client, err := torrent.NewClient(&tc)
	if err != nil {
		return err
	}
	e.client = client

	return nil
}

func (e *Engine) ShowDebug(w io.Writer) {
	e.client.WriteStatus(w)
}

func (e *Engine) AddFileSeed(path string) (infohash string, err error) {

	mi, err := base.GenerateMetaInfo(path)
	if err != nil {
		return
	}
	infohash = mi.HashInfoBytes().String()

	//check exist
	t, ok := e.client.Torrent(mi.HashInfoBytes())
	if ok {
		return infohash, nil
	}

	t, err = e.client.AddTorrent(mi)
	if err != nil {
		return
	}

	t.Info().Name = path
	go func() {
		<-t.GotInfo()
		t.DownloadAll()
	}()

	return
}

func (e *Engine) DelFileSeed(path string) (infohash string, err error) {

	mi, err := base.GenerateMetaInfo(path)
	if err != nil {
		return
	}
	infohash = mi.HashInfoBytes().String()

	//check exist
	t, ok := e.client.Torrent(mi.HashInfoBytes())
	if ok {
		t.Drop()
	}

	return
}

func (e *Engine) ListFileSeed() (res []string) {

	torrents := e.client.Torrents()
	for _, v := range torrents {
		res = append(res, v.Name())
	}
	return
}
