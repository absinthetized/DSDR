package services

import (
	data "dsdr/data"
	models "dsdr/models"
	"sort"
	"strings"
)

// SearchRole returns the roles whose permissions match the provided searchString
// searchString is a string of space separated search terms
func SearchRole(searchString string, repo *data.RoleRepository) ([]models.Role, error) {
	// no search term has been passed, just return the whole DB!
	if len(searchString) == 0 {
		return repo.FindAll(), nil
	}

	// separate the search terms in a slice
	items := strings.Split(searchString, " ")

	// sort search so that, if not all the criteria are matched we can order
	// alphabetically by first maching term later
	sort.Strings(items)

	return repo.FindPermissionsByRegexArray(items), nil //check if err is of any usage in this context..
}