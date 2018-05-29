package base

import "fmt"

type UrlType int

const (
	UtHash UrlType = iota
)

func (s *Client) UrlFor(ut UrlType, args ...interface{}) string {
	return UrlFor(ut, s.addr, args...)
}

func UrlFor(ut UrlType, host string, args ...interface{}) string {
	switch ut {
	case UtHash:
		return fmt.Sprintf("http://%s/hash", host)
		// case UtBoxTaskAct:
		// 	return fmt.Sprintf("https://%s/box/%s/task/%s/%s", append([]interface{}{host, ver}, args...)...)
	}
	return fmt.Sprintf("http://%s/", host)
}
