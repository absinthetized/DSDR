package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	models "dsdr/models" // temporary patch
)

// the roles repository mimiking an actual data layer (eg. a DB)
type RoleRepository struct {
	roles []models.Role
}

// NewRoleRepository init a role repository
func NewRoleRepository() (*RoleRepository, error) {
	repo := new(RoleRepository)
	var err error

	repo.roles, err = db_parser()
	return repo, err
}

// SearchFor scans the repo for the passed items and concats the results
func (r *RoleRepository) SearchFor(terms []string) []models.Role {
	// no search term has been passed, just return the whole DB!
	if len(terms) == 0 {
		return r.roles
	}

	// here we perform an actual search

	var roles []models.Role
	for _, term := range terms {
		roles = r.searchSingleTerm(term)
	}

	return roles
}

// searchSingleTerm is auxiliary and maches a single terms against all the DB
func (r *RoleRepository) searchSingleTerm(term string) []models.Role {
	//let term = new RegExp(searchTerm)
	//
	// roles.forEach(role => {
	//    //console.log(role)
	//    if (role.includedPermissions === undefined)
	// 	  return

	//    // let see if we match this search term
	//    let matchingPerms = role.includedPermissions.filter(perm => {
	// 	  return term.test(perm) ? true : false
	//    })

	//    //add the number of matches for this search term
	//    if (matchingPerms.length > 0) {
	// 	  role.matches += matchingPerms.length
	// 	  role.matchedBy.push(searchTerm)
	//    }
	// })

	// just a mockup
	return r.roles
}

//aux function. db_parser loads IAM info from the fake DB
func db_parser() ([]models.Role, error) {
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
