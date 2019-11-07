// Package checkproxy makes checking proxies easy!
package checkproxy

import (
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"golang.org/x/net/proxy"
)

// CheckSocksProxy checks if a SOCKS proxy is online and usable.
func CheckSocksProxy(proxyAddress string, timeout time.Duration) bool {

	d := net.Dialer{
		Timeout:   timeout,
		KeepAlive: timeout,
	}

	dialer, _ := proxy.SOCKS5("tcp", proxyAddress, nil, &d)

	httpClient := &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			DisableKeepAlives: true,
			Dial:              dialer.Dial,
		},
	}

	response, err := httpClient.Get("https://api.ipify.org?format=json")

	if err != nil {
		return false
	}

	defer response.Body.Close()

	io.Copy(ioutil.Discard, response.Body)

	return true
}
