package echosse

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type sseClient struct {
	c echo.Context
}

func newSSEClint(c echo.Context) *sseClient {
	c.Response().Header().Set(echo.HeaderContentType, "text/event-stream")
	c.Response().Header().Set(echo.HeaderCacheControl, "no-cache")
	c.Response().Header().Set(echo.HeaderConnection, "keep-alive")

	c.Response().WriteHeader(http.StatusOK)
	return &sseClient{
		c: c,
	}
}

func (c *sseClient) SendEvent(text string) error {
	w := c.c.Response().Writer
	data := fmt.Sprintf("data: %s\n\n", string(text))
	w.Write([]byte(data))
	c.c.Response().Flush()
	return nil
}

func (c *sseClient) Close() error {
	w := c.c.Response().Writer
	data := []byte("data: [DONE]\n\n")
	w.Write(data)
	return nil
}
