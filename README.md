# echo-climit

Echo-climit is a concurrency limiter middleware for [echo](https://github.com/labstack/echo/).

It allows exactly N requests (globally, per endpoint or per group) to be handled at given time.
Can be useful when used with cache middleware and you want other request to wait for the one who will populate the cache so that you don't overflow db with heavy requests.

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
	e := echo.New()

	// Via Use, either global, or appy to existing group (when using multiple middlewares)
	// 1 - number of parallel requests, time.Minute - wait time before cancelling the request by timeout	
	e.Use(climit.New(1, time.Minute))

	// On group creation
	// 1 - number of parallel requests, time.Minute - wait time before cancelling the request by timeout	
	g := e.Group("/group", climit.New(1, time.Minute))
}

```
