# What is it


# Usage

```golang
package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
    "github.com/CorrectRoadH/echo-sse"
)
func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/events", func(c echo.Context) error {
        client := newSSEClint(c)
        defer client.Close()
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