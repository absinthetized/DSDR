package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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

	// serves a search to the front end - mockup for now
	r.GET("/search", func(c *gin.Context) {
		var searchString string

		roles, err := db_parser()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err}) //c.JSON returns and ends the function
		}

		// just return the whole DB for now - client will perform the search in this mockup
		if c.ShouldBindQuery(&searchString) == nil {
			c.JSON(http.StatusOK, roles) //to be changed with proper search

		} else {
			c.JSON(http.StatusOK, roles)
		}

		log.Print("query string is:", searchString)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}

func db_parser() ([]role, error) {
	role_dir := "./roles"
	files, err := ioutil.ReadDir(role_dir)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	var roles []role

	for id, file := range files {
		// read file
		data, err := ioutil.ReadFile(role_dir + string(os.PathSeparator) + file.Name())
		if err != nil {
			log.Print(err)
			return nil, err
		}

		var role role
		err = json.Unmarshal(data, &role)
		if err != nil {
			log.Print(err)
			return nil, err
		}

		role.Id = id
		roles = append(roles, role)
	}

	return roles, nil
}

type role struct {
	Description         string   `json:"description"`
	Name                string   `json:"name"`
	Stage               string   `json:"stage"`
	Title               string   `json:"title"`
	IncludedPermissions []string `json:"includedPermissions"`
	Id                  int      `json:"id"`
}
