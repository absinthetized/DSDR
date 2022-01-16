package data

import (
	mocks "dsdr/mocks/data"
	"dsdr/models"
	"sort"
	"testing"
)

func TestNewRoleRepository(t *testing.T) {
	var testDB mocks.DB
	testDB.On("Connect", "../roles").Return(nil)

	NewRoleRepository(&testDB) //this should not fail eve if db is nil
}

func TestFindAll(t *testing.T) {
	var testDB mocks.DB
	testDB.On("Connect", "../roles").Return(nil)
	testDB.On("Roles").Return([]models.BasicIAMRole{
		*new(models.BasicIAMRole),
		*new(models.BasicIAMRole),
	})

	repo := NewRoleRepository(&testDB)
	roles := repo.FindAll()
	if len(roles) == 0 {
		t.Fatalf("find all should return non empty list of roles")
	}
}

func TestFindPermissionsByRegexArrayEmptyArray(t *testing.T) {
	var emptyInput []string

	var testDB mocks.DB
	testDB.On("Connect", "../roles").Return(nil)
	testDB.On("Roles").Return([]models.BasicIAMRole{
		*new(models.BasicIAMRole),
		*new(models.BasicIAMRole),
	})

	repo := NewRoleRepository(&testDB)
	roles, err := repo.FindPermissionsByRegexArray(emptyInput)

	if len(roles) != len(testDB.Roles()) || err != nil {
		t.Fatalf("find by permissions should return all results and no error if an empty array is passed")
	}
}

func TestFindPermissionsByRegexArrayWithArrayFilledWithExistentValues(t *testing.T) {
	input := []string{"compute", "network"}

	role1 := *new(models.BasicIAMRole)
	role1.IncludedPermissions = []string{"computeUser", "networkUser"}
	role2 := *new(models.BasicIAMRole)
	role2.IncludedPermissions = []string{"computeUser"}
	var testDB mocks.DB

	testDB.On("Connect", "../roles").Return(nil)
	testDB.On("Roles").Return([]models.BasicIAMRole{
		*new(models.BasicIAMRole),
		role1,
		role2,
	})

	repo := NewRoleRepository(&testDB)
	roles, err := repo.FindPermissionsByRegexArray(input)

	// that's where the previous impl. of the test failed: it should be exactly 2 match and no err...
	if len(roles) != 2 || err != nil {
		t.Fatalf("find by permissions should return exact number of results and no error if a well inited array is passed")
	}

	/* also check for contents of returned permissions */

	target := "computeUser"

	isRoleOk := func(role models.Role) bool {
		sort.Strings(role.IncludedPermissions)
		i := sort.SearchStrings(role.IncludedPermissions, target)
		return (i < len(role1.IncludedPermissions) && role1.IncludedPermissions[i] == target)
	}

	if !isRoleOk(roles[0]) || !isRoleOk(roles[1]) {
		t.Fatalf("find by permissions should return the expected permission array")
	}
}
