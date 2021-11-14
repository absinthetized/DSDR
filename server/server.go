package main

import (
	"dsdr/services"
	"log"
	"net/http"

	"dsdr/data"

	"github.com/gin-gonic/gin"
)

func main() {
	// start default gin middleware
	r := gin.Default()

	// serve static resources from svelte
	r.Static("/svelte", "./static/build")
	r.StaticFile("/favicon.png", "./static/favicon.png")
	r.StaticFile("/global.css", "./static/global.css")
	r.StaticFile("/", "./static/index.html")

	// init the repo
	DB, err := data.NewDB()
	if err != nil {
		panic("Unable to connect to roles repository. Aborting.")
	}

	// serves a search to the front end - mockup for now
	r.GET("/search", func(c *gin.Context) {
		var searchString string
		err := c.ShouldBindQuery(&searchString)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
		}

		log.Print("query string is:", searchString)

		roles, err := services.SearchRole(searchString, DB)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err}) //c.JSON returns and ends the function
		}

		c.JSON(http.StatusOK, roles)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
