package main

import (
	"dsdr/models"
	"dsdr/services"
	"log"
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

	// init the repo
	var DB data.BqDB

	err := DB.Connect("")
	if err != nil {
		panic("Unable to connect to roles repository. Aborting.")
	}

	defer DB.Close()

	// just a debug line here...
	IamRoleMapper := data.NewDataMapper[models.BqIAMRole](&DB, "roles_dataset.roles_table")

	name := "matteo"
	filter := new(models.BqIAMRoleFilter)
	filter.Name = &name
	simpleQuery := IamRoleMapper.FindAll().Where(*filter)
	iams, err1 := IamRoleMapper.Run(simpleQuery)
	if err1 != nil {
		log.Println("datamapper FindAll failed:", err)

	} else {
		log.Println("----------------")
		log.Println("-- DM result ---")
		log.Println("----------------")
		for _, iam := range iams {
			log.Println(iam)
		}
		log.Println("----------------")
	}

	// serves a search to the front end - mockup for now
	r.GET("/api/v1/search", func(c *gin.Context) {
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

	r.Run(":8000") // listen and serve on 0.0.0.0:8080
}
