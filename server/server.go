package main

import (
	"dsdr/services"
	"net/http"

	"dsdr/data"

	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
)

type SearchQuery struct {
	Query string `form:"query"`
}

func main() {
	// start default gin middleware
	r := gin.Default()
	p := bluemonday.StrictPolicy()

	// serve static resources from svelte
	r.Static("/svelte", "./static/build")
	r.StaticFile("/favicon.png", "./static/favicon.png")
	r.StaticFile("/global.css", "./static/global.css")
	r.StaticFile("/", "./static/index.html")

	// init the repo
	var DB data.FileSystemDB
	err := DB.Connect("roles")
	if err != nil {
		panic("Unable to connect to roles repository. Aborting.")
	}

	// serves a search to the front end - mockup for now
	r.GET("/search", func(c *gin.Context) {
		var sq SearchQuery
		err := c.ShouldBindQuery(&sq)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
		}

		searchString := p.Sanitize(sq.Query)

		roles, err := services.SearchRole(searchString, &DB)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err}) //c.JSON returns and ends the function
		}

		c.JSON(http.StatusOK, roles)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
