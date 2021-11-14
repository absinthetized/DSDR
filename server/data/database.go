package data

import (
	"dsdr/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// the roles repository mimiking an actual data layer (eg. a DB)
type DB struct {
	Roles []models.BasicIAMRole
}

// NewRoleRepository init a role repository
func NewDB() (*DB, error) {
	db := new(DB)
	var err error

	db.Roles, err = db_parser()
	return db, err
}

//aux function. db_parser loads IAM info from the fake DB
func db_parser() ([]models.BasicIAMRole, error) {

	// TODO: use BasicIAMRole to load info
	role_dir := "./roles"
	files, err := ioutil.ReadDir(role_dir)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	var roles []models.BasicIAMRole

	for id, file := range files {
		// read file
		data, err := ioutil.ReadFile(role_dir + string(os.PathSeparator) + file.Name())
		if err != nil {
			log.Print(err)
			return nil, err
		}

		var role models.BasicIAMRole
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
