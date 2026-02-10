---
sidebar_position: 2
---

# ⏳ Timeout

You can set timeout duration.

The default is 10s.

```go title="main.go"
package main

import (
	"time"

	"github.com/poteto0/go-nba-sdk/gns"
	"github.com/poteto0/go-nba-sdk/types"
)

func main() {
	config := types.GnsConfig{
		Timeout: 5 * time.Second,
	}

	client := gns.NewClient(&config)

	...
}
```
