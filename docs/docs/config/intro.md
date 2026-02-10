---
sidebar_position: 1
---

# 📌 Summary

This is your go-nba-sdk config.

```go title="config.go"
type GnsConfig struct {
	// default is 10s
	Timeout time.Duration

	// default is empty string
	ProxyUrl string
}
```

If you didn't want to set, you should set nil.

```go title="main.go"
package main

import (
	"github.com/poteto0/go-nba-sdk/gns"
)

func main() {
	client := gns.NewClient(nil)
}
```
