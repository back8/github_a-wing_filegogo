package share

import (
	"net/url"
	"strings"

	"filegogo/lightcable"
)

func IsShareInit(addr string) bool {
	if u, err := url.Parse(addr); err != nil {
		return false
	} else {
		if arr := strings.Split(u.Path, "/"); len(arr) > 2 {
			if arr[2] != "" {
				return true
			}
			return false
		}
		return false
	}
}

// http://localhost:8033/<PrefixShort>/1024"
// To:
// ws://localhost:8033/<PrefixShare>/1024"
func ShareToWebSocket(addr string) string {
	if u, err := url.Parse(addr); err != nil {
		return addr
	} else {
		if u.Scheme == "https" {
			u.Scheme = "wss"
		} else {
			u.Scheme = "ws"
		}

		if arr := strings.Split(u.Path, "/"); len(arr) > 2 {
			u.Path = "/" + lightcable.PrefixShare + "/" + arr[2]
		}
		return u.String()
	}
}

// ws://localhost:8033/<PrefixShare>/1024"
// To:
// http://localhost:8033/<PrefixShort>/1024"
func WebSocketToShare(addr string) string {
	if u, err := url.Parse(addr); err != nil {
		return addr
	} else {
		if u.Scheme == "wss" {
			u.Scheme = "https"
		} else {
			u.Scheme = "http"
		}

		if arr := strings.Split(u.Path, "/"); len(arr) > 2 {
			u.Path = arr[2]
		}
		return u.String()
	}
}
