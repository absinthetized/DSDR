package services

import (
	data "dsdr/data"
	models "dsdr/models"
	"sort"
	"strings"
)

// SearchRole returns the roles whose permissions match the provided searchString
// searchString is a string of space separated search terms
func SearchRole(searchString string, db *data.DB) ([]models.Role, error) {
	repo := data.NewRoleRepository(db)

	// no search term has been passed, just return the whole DB!
	if len(searchString) == 0 {
		return repo.FindAll(), nil
	}

	// separate the search terms in a slice
	// sort search so that, if not all the criteria are matched we can order
	// alphabetically by first maching term later
	items := strings.Split(searchString, " ")
	sort.Strings(items)
	return repo.FindPermissionsByRegexArray(items)
}
