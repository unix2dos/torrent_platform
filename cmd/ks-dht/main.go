package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/anacrolix/dht"
	"github.com/gin-gonic/gin"
	envcfg "github.com/wealthworks/envflagset"

	"torrent_platform/base"
)

var (
	Version   = "dev"
	udpAddr   = ":16181"
	debugAddr = ":16182"
	dataAddr  = ":26180"

	bootstrapAddrs string
	server         *dht.Server
	redisDB        = 0
	redisKey       = "dhtServerNodes"

	PingInterval = 15 * time.Second
)

func loadTable() (err error) {

	// str, err := redis.Redis(redisDB).Get(redisKey)
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
	// 	if server.AddNode(n) == nil {
	// 		added++
	// 	}
	// }
	// log.Printf("loaded %d nodes from table file", added)

	return
}

func saveTable() (err error) {

	// nodes := server.Nodes()
	// b, err := krpc.CompactIPv6NodeInfo(nodes).MarshalBinary()
	// if err != nil {
	// 	return
	// }
	// err = redis.Redis(redisDB).Set(redisKey, b)

	return
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	envcfg.New("ks-dht", Version)

	flag.StringVar(&udpAddr, "udp-addr", udpAddr, "local UDP address")
	flag.StringVar(&debugAddr, "debug-addr", debugAddr, "dht debug info address")
	flag.StringVar(&bootstrapAddrs, "bootstrap-addrs", "", "bootstrap addrs, separate with comma")
	flag.IntVar(&redisDB, "redis-db", 0, "redis db index")

	envcfg.Parse()

	if bootstrapAddrs != "" {
		base.GlobalBootstrapAddrs = strings.Split(bootstrapAddrs, ",")
	}

	//udp server
	conn, err := net.ListenPacket("udp4", udpAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	server, err = dht.NewServer(&dht.ServerConfig{
		Conn:          conn,
		StartingNodes: base.BootstrapAddrs,
		NoSecurity:    true,
	})
	if err != nil {
		log.Fatal(err)
	}

	//http debug info
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			server.WriteStatus(w)
		})
		http.ListenAndServe(debugAddr, mux)
	}()

	//listen data
	go func() {
		router := gin.New()
		router.PUT("/hash", putHash)
		router.DELETE("/hash", delHash)
		router.GET("/hash", getHash)
		router.Run(dataAddr)
	}()

	//load data
	err = loadTable()
	if err != nil {
		log.Printf("error loading table: %s\n", err)
	}
	log.Printf("dht server on %s, ID is %x", server.Addr(), server.ID())

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		ch := make(chan os.Signal)
		signal.Notify(ch)
		<-ch
		cancel()
	}()

	go func() {
		if tried, err := server.Bootstrap(); err != nil {
			log.Printf("error bootstrapping: %s", err)
		} else {
			log.Printf("finished bootstrapping: crawled %d addrs", tried)
		}
	}()

	go func() {
		for {
			if tried, err := server.PingNodes(); err != nil {
				log.Printf("error ping nodes: %s", err)
			} else {
				log.Printf("finished ping nodes: crawled %d addrs", tried)
			}
			time.Sleep(PingInterval) //server定时ping nodes, 1. 让新加入的节点response, 变成good节点 2. 清理无效节点
		}
	}()

	<-ctx.Done()
	server.Close()

	if err := saveTable(); err != nil {
		log.Printf("error saving node table: %s", err)
	}
}
