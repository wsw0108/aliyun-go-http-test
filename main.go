package main

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StripFcBasePath() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			req := c.Request()
			url := req.URL
			basePath := req.Header.Get("x-fc-base-path")
			if basePath != "" {
				replacer := strings.NewReplacer(basePath, "")
				uri := replacer.Replace(req.URL.Path)
				req.RequestURI = uri
				url.Path = uri
			}
			return next(c)
		}
	}
}

func main() {
	// Echo instance
	e := echo.New()

	e.Pre(StripFcBasePath())

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/ping", ping)

	// Start server
	e.Logger.Fatal(e.Start(":9000"))
}

// Handler
func ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong!")
}
