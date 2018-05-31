package base

import (
	"errors"
	"log"
	"net"

	"github.com/anacrolix/dht"
)

var (
	GlobalBootstrapAddrs = []string{"207.246.80.69:16181"} //外网
	// GlobalBootstrapAddrs = []string{"172.24.120.28:16181"}
)

func BootstrapAddrs() (addrs []dht.Addr, err error) {
	for _, s := range GlobalBootstrapAddrs {
		host, port, err := net.SplitHostPort(s)
		if err != nil {
			panic(err)
		}
		hostAddrs, err := net.LookupHost(host)
		if err != nil {
			log.Printf("error looking up %q: %v", s, err)
			continue
		}
		for _, a := range hostAddrs {
			ua, err := net.ResolveUDPAddr("udp", net.JoinHostPort(a, port))
			if err != nil {
				log.Printf("error resolving %q: %v", a, err)
				continue
			}
			addrs = append(addrs, dht.NewAddr(ua))
		}
	}
	if len(addrs) == 0 {
		err = errors.New("nothing resolved")
	}
	return
}
