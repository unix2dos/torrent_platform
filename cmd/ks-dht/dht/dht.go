package dht

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/anacrolix/dht"

	"torrent_platform/base"
)

type DHT struct {
	cf     Config
	server *dht.Server
}

func New(cf Config) *DHT {
	return &DHT{cf: cf}
}

func (d *DHT) Start() {

	//dht server
	conn, err := net.ListenPacket("udp4", d.cf.DhtAddr)
	if err != nil {
		return
	}
	defer conn.Close()

	d.server, err = dht.NewServer(&dht.ServerConfig{
		Conn:          conn,
		StartingNodes: base.BootstrapAddrs,
		NoSecurity:    true,
	})
	if err != nil {
		return
	}

	//load data
	err = d.loadTable()
	if err != nil {
		return
	}
	log.Printf("dht server on %s, ID is %x", d.server.Addr(), d.server.ID())

	//bootstrap
	go func() {
		if tried, err := d.server.Bootstrap(); err != nil {
			log.Printf("error bootstrapping: %s", err)
		} else {
			log.Printf("finished bootstrapping: crawled %d addrs", tried)
		}
	}()

	//pingNodes
	go func() {
		for {
			if tried, err := d.server.PingNodes(); err != nil {
				log.Printf("error ping nodes: %s", err)
			} else {
				log.Printf("finished ping nodes: crawled %d addrs", tried)
			}
			time.Sleep(d.cf.PingInterval) //server定时ping nodes, 1. 让新加入的节点response, 变成good节点 2. 清理无效节点
		}
	}()

	//http debug info
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			d.server.WriteStatus(w)
		})
		http.ListenAndServe(d.cf.DhtDebugAddr, mux)
	}()

	//block
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		ch := make(chan os.Signal)
		signal.Notify(ch)
		<-ch
		signal.Stop(ch)
		cancel()
	}()

	<-ctx.Done()
	d.server.Close()

	if err := d.saveTable(); err != nil {
		log.Printf("error saving node table: %s", err)
	}
}

func (d *DHT) loadTable() (err error) {

	// str, err := redis.Redis(d.cf.RedisDB).Get(d.cf.RedisKey)
	// if err != nil {
	// 	return
	// }
	//
	// var ns krpc.CompactIPv6NodeInfo
	// err = ns.UnmarshalBinary([]byte(str))
	// if err != nil {
	// 	return
	// }
	//
	// var added int64
	// for _, n := range ns {
	// 	if d.server.AddNode(n) == nil {
	// 		added++
	// 	}
	// }
	// log.Printf("loaded %d nodes from table file", added)
	return
}

func (d *DHT) saveTable() (err error) {
	// nodes := d.server.Nodes()
	// b, err := krpc.CompactIPv6NodeInfo(nodes).MarshalBinary()
	// if err != nil {
	// 	return
	// }
	// err = redis.Redis(d.cf.RedisDB).Set(d.cf.RedisKey, b)
	return
}
