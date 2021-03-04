package main

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func StripFcBasePath(c *gin.Context) {
	req := c.Request
	basePath := req.Header.Get("x-fc-base-path")
	if basePath != "" {
		replacer := strings.NewReplacer(basePath, "")
		uri := replacer.Replace(req.URL.Path)
		req.RequestURI = uri
		req.URL.Path = uri
	}
	c.Next()
}

func main() {
	r := gin.Default()
	r.Use(StripFcBasePath)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("0.0.0.0:9000")
}
