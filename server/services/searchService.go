package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Role models the json info of GCP IAM roles
type Role struct {
	Description         string   `json:"description"`
	Name                string   `json:"name"`
	Stage               string   `json:"stage"`
	Title               string   `json:"title"`
	IncludedPermissions []string `json:"includedPermissions"`
	Id                  int      `json:"id"`
}

// SearchRole returns the roles whose permissions match the provided searchString
// searchString is a string of space separated search terms
func SearchRole(searchString string) ([]Role, error) {
	//just fake the search for now and simply load the full database
	return db_parser()
}

//db_parser loads IAM info from the fake DB
func db_parser() ([]Role, error) {
	role_dir := "./roles"
	files, err := ioutil.ReadDir(role_dir)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	var roles []Role

	for id, file := range files {
		// read file
		data, err := ioutil.ReadFile(role_dir + string(os.PathSeparator) + file.Name())
		if err != nil {
			log.Print(err)
			return nil, err
		}

		var role Role
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
