---
sidebar_position: 3
---

# 🌐 Proxy

You can use a proxy server to access the NBA API if you are in a restricted environment.

The request route will be:
`[Your Client] ---> [Your Proxy] ---> [NBA API]`

The default is an empty string (no proxy).

```go title="main.go"
package main

import (
	"github.com/poteto0/go-nba-sdk/gns"
	"github.com/poteto0/go-nba-sdk/types"
)

func main() {
	config := types.GnsConfig{
		ProxyUrl: "http://your-proxy-server.com:8080",
	}

	client := gns.NewClient(&config)

	...
}
```
