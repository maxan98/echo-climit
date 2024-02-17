# echo-climit

Echo-climit is a concurrency limiter middleware for [echo](https://github.com/labstack/echo/).

This version is working with echo v3. Please checkout v2 branch if you want use session with echo v2.

## Installation
`go get github.com/maxan98/echo-climit`

## Usage
```go
package main

import (
	"time"

	"github.com/labstack/echo/v4"
	climit "github.com/maxan98/echo-climit"
)

func main() {
	// Via Use, either global, or appy to existing group (when using multiple middlewares)
	e := echo.New()
	e.Use(climit.New(1, time.Minute))

	// On group creation
	g := e.Group("/group", climit.New(1, time.Minute))
}

```