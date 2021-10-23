package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// start default gin middleware
	r := gin.Default()

	// map static resources from svelte
	r.Static("/svelte", "./static/build")
	r.StaticFile("/favicon.png", "./static/favicon.png")
	r.StaticFile("/global.css", "./static/global.css")
	r.StaticFile("/", "./static/index.html")

	// just a debug end point to be removed
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
