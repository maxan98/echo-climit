package climit

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// New creates a concurrency limiter middleware
func New(limit int64, wait time.Duration) echo.MiddlewareFunc {
	// if limit <= 0, returns an empty middleware
	if limit <= 0 {
		return func(next echo.HandlerFunc) echo.HandlerFunc {
			return next
		}
	}

	windows := make(chan struct{}, limit)
	// the middleware
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			release := func() {
				<-windows
			}
			select {
			case <-time.After(wait):
				return c.JSON(http.StatusTooManyRequests, "concurrency limit exceeded")
			case windows <- struct{}{}:
				defer release()
				return next(c)
			}
		}
	}
}
