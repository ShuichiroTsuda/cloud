package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ShuichiroTsuda/cloud/db"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello world")
	})
	db.Init()
	r.Run(":8080")
	db.Close()
}
