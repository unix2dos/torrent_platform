package dht

import "time"

type Config struct {
	DhtAddr      string
	DhtDebugAddr string
	RedisDB      int
	RedisKey     string
	PingInterval time.Duration
}
