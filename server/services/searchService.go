package services

import (
	data "dsdr/data"
	models "dsdr/models"
)

// SearchRole returns the roles whose permissions match the provided searchString
// searchString is a string of space separated search terms
func SearchRole(searchString string) ([]models.Role, error) {
	//just fake the search for now and simply load the full database
	return data.DB_parser()
}
