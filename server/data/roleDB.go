package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	models "dsdr/models" // temporary patch
)

//DB_parser loads IAM info from the fake DB
func DB_parser() ([]models.Role, error) {
	role_dir := "./roles"
	files, err := ioutil.ReadDir(role_dir)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	var roles []models.Role

	for id, file := range files {
		// read file
		data, err := ioutil.ReadFile(role_dir + string(os.PathSeparator) + file.Name())
		if err != nil {
			log.Print(err)
			return nil, err
		}

		var role models.Role
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
