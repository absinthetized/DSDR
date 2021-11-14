package data

import (
	"regexp"

	models "dsdr/models"
)

// the roles repository mimiking an actual data layer (eg. a DB)
type RoleRepository struct {
	roles []models.Role
	db    *DB
}

// NewRoleRepository init a role repository
func NewRoleRepository(db *DB) *RoleRepository {
	repo := new(RoleRepository)
	repo.db = db

	return repo
}

// FindAll returns the whole dataset
func (r *RoleRepository) FindAll() (roles []models.Role) {
	for _, IAM := range r.db.Roles {
		r.roles = append(r.roles, *models.NewRoleFromIAM(IAM))
	}

	return r.roles
}

// FindPermissionsByRegexArray scans the repo for the reuired permissions and filters out the resulting IAMs
func (r *RoleRepository) FindPermissionsByRegexArray(terms []string) ([]models.Role, error) {
	var err error

	// this associates a IAM pointer to a Role so that we can
	// quickly check if a Role has been already discovered among IAMs
	var roleMap map[*models.BasicIAMRole]models.Role

	// search for all matches in all IAMs
	for _, term := range terms {
		err = r.searchSingleTerm(term, roleMap)
		if err != nil { // this is abit raw I could notify a wang rather than interrupt the loop
			break
		}
	}

	// extract matches and compute percentage of match againts the range of passed terms
	var roleMatches []models.Role
	for _, roleMatch := range roleMap {
		roleMatch.PercMatch =
			float32(roleMatch.Matches) / float32(len(roleMatch.IncludedPermissions))

		roleMatches = append(roleMatches, roleMatch)
	}

	return roleMatches, err
}

// searchSingleTerm is auxiliary and maches a single term against all the IAM DB
func (r *RoleRepository) searchSingleTerm(searchTerm string, roleMap map[*models.BasicIAMRole]models.Role) error {
	term, err := regexp.Compile(searchTerm)
	if err != nil {
		return err
	}

	// loop over all roles in the DB and search for permissions matching our term
	for _, IAM := range r.db.Roles {

		// no perms for this IAM just continue (a case at least have been hit during tests)
		if len(IAM.IncludedPermissions) == 0 {
			continue
		}

		// count matches

		var matchesNum int = 0

		for _, perm := range IAM.IncludedPermissions {
			matches := term.FindString(perm)
			matchesNum += len(matches)
		}

		if matchesNum == 0 { // nothing to do here
			continue
		}

		// if we match add/update this IAM in the Roles map
		role, roleExists := roleMap[&IAM]
		if !roleExists {
			// add a new Role to the map
			role = *models.NewRoleFromIAM(IAM)
			roleMap[&IAM] = role
		}

		role.Matches += matchesNum
		role.MatchedBy = append(role.MatchedBy, searchTerm)
	}

	return nil
}
