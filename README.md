# checkproxy
 

[![GoDoc](https://godoc.org/github.com/gomjw/checkproxy?status.svg)](https://godoc.org/github.com/gomjw/checkproxy)

> Package checkproxy makes checking proxies easy!

## Install

```console
go get -u github.com/gomjw/checkproxy/...
```

## Import

```go
import "github.com/gomjw/checkproxy"
```

## Usage

#### func  CheckSocksProxy

```go
func CheckSocksProxy(proxyAddress string, timeout time.Duration) bool
```
CheckSocksProxy checks if a SOCKS proxy is online and usable.



---

> Made by the awesome contributors of [@gomjw](https://github.com/gomjw) &nbsp;&middot;&nbsp;
> Maintainer [@MarvinJWendt](https://github.com/MarvinJWendt) &nbsp;&middot;&nbsp;
> Readme template version: 1.1.0
