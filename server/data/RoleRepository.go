package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"regexp"

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

// FindAll returns the whole dataset
func (r *RoleRepository) FindAll() (roles []models.Role) {
	return r.roles
}

// FindPermissionsByRegexArray scans the repo for the passed items and concats the results
func (r *RoleRepository) FindPermissionsByRegexArray(terms []string) ([]models.Role, error) {
	// TODO ,maybe not here, reset roles stats

	var roles []models.Role

	for _, term := range terms {
		err := r.searchSingleTerm(term)
		if err != nil {
			return nil, err
		}
	}

	// TODO: perform filtering and return matching roles

	return roles, nil
}

// searchSingleTerm is auxiliary and maches a single terms against all the DB
func (r *RoleRepository) searchSingleTerm(searchTerm string) error {
	term, err := regexp.Compile(searchTerm)
	if err != nil {
		return err
	}

	//
	// roles.forEach(role => {
	//    //console.log(role)
	//    if (role.includedPermissions === undefined)
	// 	  return

	//    // let see if we match this search term
	//    let matchingPerms = role.includedPermissions.filter(perm => {
	// 	  return term.test(perm) ? true : false
	//    })

	for _, role := range r.roles {
		if role.IncludedPermissions != nil { //or have to check len?!
			var matchingPerms []string

			for _, perm := range role.IncludedPermissions {
				match := term.FindString(perm)
				if len(match) > 0 {
					matchingPerms = append(matchingPerms, match)
				}
			}

			//    //add the number of matches for this search term
			//    if (matchingPerms.length > 0) {
			// 	  role.matches += matchingPerms.length
			// 	  role.matchedBy.push(searchTerm)
			//    }
			// })

			if len(matchingPerms) > 0 {
				role.Matches += len(matchingPerms)
				role.MatchedBy = append(role.MatchedBy, matchingPerms...)
			}
		}
	}

	return nil
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
