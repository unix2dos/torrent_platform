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

func (e *Engine) AddMagnet(hash string) (*torrent.Torrent, error) {
	url := "magnet:?xt=urn:btih:" + hash
	return e.client.AddMagnet(url)
}
