package main

import (
	"flag"
	"log"
	"strings"
	"time"

	envcfg "github.com/wealthworks/envflagset"

	"torrent_platform/base"
	"torrent_platform/cmd/ks-dht/dht"
	"torrent_platform/cmd/ks-dht/server"
)

var (
	Version  = "dev"
	httpAddr = ":26180"

	dhtAddr        = ":16181"
	dhtDebugAddr   = ":16182"
	bootstrapAddrs string
	redisDB        = 0
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	envcfg.New("ks-dht", Version)
	flag.StringVar(&httpAddr, "http-addr", httpAddr, "http addr")

	flag.StringVar(&dhtAddr, "dht-addr", dhtAddr, "dht server address")
	flag.StringVar(&dhtDebugAddr, "dht-debug-addr", dhtDebugAddr, "dht debug info address")
	flag.StringVar(&bootstrapAddrs, "bootstrap-addrs", "", "bootstrap addrs, separate with comma")
	flag.IntVar(&redisDB, "redis-db", 0, "redis db index")
}

func main() {

	envcfg.Parse()
	if bootstrapAddrs != "" {
		base.GlobalBootstrapAddrs = strings.Split(bootstrapAddrs, ",")
	}

	dht := dht.New(dht.Config{
		DhtAddr:      dhtAddr,
		DhtDebugAddr: dhtDebugAddr,
		RedisDB:      redisDB,
		RedisKey:     "dhtServerNodes",
		PingInterval: 15 * time.Second,
	})

	server := &server.Server{DHT: dht}
	server.Run(httpAddr)
}
