package data

import (
	"testing"
)

func TestNewRoleRepository(t *testing.T) {
	var db *DB
	NewRoleRepository(db) //this should not fail eve if db is nil
}

func TestFindAll(t *testing.T) {
	//I need to mock DB somewhat...
	db, _ := NewDB()
	repo := NewRoleRepository(db)
	roles := repo.FindAll()
	if len(roles) == 0 {
		t.Fatalf("find all shloud return non empty list of roles")
	}
}

func TestFindPermissionsByRegexArrayEmptyArray(t *testing.T) {
	var emptyInput []string

	db, _ := NewDB()
	repo := NewRoleRepository(db)
	roles, err := repo.FindPermissionsByRegexArray(emptyInput)

	if len(roles) != len(db.Roles) || err != nil {
		t.Fatalf("find by permissions should return all results and no error if an empty array is passed")
	}
}

func TestFindPermissionsByRegexArrayFilledArray(t *testing.T) {
	input := []string{"compute", "network"}

	db, _ := NewDB()
	repo := NewRoleRepository(db)
	roles, err := repo.FindPermissionsByRegexArray(input)

	if len(roles) == 0 || err != nil {
		t.Fatalf("find by permissions should return some results and no error if a well inited array is passed")
	}
}

// 	return roleMatches, err
// }

// // searchSingleTerm is auxiliary and maches a single term against all the IAM DB
// func (r *RoleRepository) searchSingleTerm(searchTerm string, roleMap map[*models.BasicIAMRole]models.Role) error {
// 	term, err := regexp.Compile(searchTerm)
// 	if err != nil {
// 		return err
// 	}

// 	// loop over all roles in the DB and search for permissions matching our term
// 	for _, IAM := range r.db.Roles {

// 		// no perms for this IAM just continue (a case at least have been hit during tests)
// 		if len(IAM.IncludedPermissions) == 0 {
// 			continue
// 		}

// 		// count matches

// 		var matchesNum int = 0

// 		for _, perm := range IAM.IncludedPermissions {
// 			matches := term.FindString(perm)
// 			matchesNum += len(matches)
// 		}

// 		if matchesNum == 0 { // nothing to do here
// 			continue
// 		}

// 		// if we match add/update this IAM in the Roles map
// 		role, roleExists := roleMap[&IAM]
// 		if !roleExists {
// 			// add a new Role to the map
// 			role = *models.NewRoleFromIAM(IAM)
// 			roleMap[&IAM] = role
// 		}

// 		role.Matches += matchesNum
// 		role.MatchedBy = append(role.MatchedBy, searchTerm)
// 	}

// 	return nil
// }
