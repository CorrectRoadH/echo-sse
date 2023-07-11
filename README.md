# What is it
This a simple library for sending Server-Sent Events (SSE) with Echo framework.

# Usage

```golang
package main

import (
	"math/rand"
	"time"

	echosse "github.com/CorrectRoadH/echo-sse"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/events", func(c echo.Context) error {
		client := echosse.NewSSEClint(c)
		for {
			const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
			seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

			b := make([]byte, 10)
			for i := range b {
				b[i] = charset[seededRand.Intn(len(charset))]
			}

			client.SendEvent(string(b))
			time.Sleep(1 * time.Second)
		}
	})
	e.Logger.Fatal(e.Start(":8080"))
}
```
