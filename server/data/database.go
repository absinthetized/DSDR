package data

import (
	"dsdr/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// the roles repository mimiking an actual data layer (eg. a DB)
type DB struct {
	Roles []models.BasicIAMRole
}

type DBMethods interface {
	Connect(string) ([]models.BasicIAMRole, error)
}

//Connect loads IAM info from the fake DB
func (d *DB) Connect(folder string) error {
	this_dir, pathErr := filepath.Abs(folder)
	if pathErr != nil {
		return pathErr
	}

	role_dir := filepath.Dir(this_dir) + string(os.PathSeparator) + "roles"

	files, err := ioutil.ReadDir(role_dir)

	if err != nil {
		log.Print(err)
		return err
	}

	for id, file := range files {
		// read file
		data, err := ioutil.ReadFile(role_dir + string(os.PathSeparator) + file.Name())
		if err != nil {
			log.Print(err)
			return err
		}

		var role models.BasicIAMRole
		err = json.Unmarshal(data, &role)
		if err != nil {
			log.Print(err)
			return err
		}

		role.Id = id
		d.Roles = append(d.Roles, role)
	}

	return nil
}
