package main

import (
	"app/modules"
	"app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	mod := modules.Get()
	routes.Api(r.Group("/api"), mod)

	r.Run(":8080")
}
