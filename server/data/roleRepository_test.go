package data

import (
	mocks "dsdr/mocks/data"
	"dsdr/models"
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

	pippo := *new(models.BasicIAMRole)
	pippo.IncludedPermissions = []string{"computeUser", "networkUser"}
	var testDB mocks.DB
	testDB.On("Connect", "../roles").Return(nil)
	testDB.On("Roles").Return([]models.BasicIAMRole{
		*new(models.BasicIAMRole),
		pippo,
	})

	repo := NewRoleRepository(&testDB)
	roles, err := repo.FindPermissionsByRegexArray(input)

	if len(roles) == 0 || err != nil {
		t.Fatalf("find by permissions should return some results and no error if a well inited array is passed")
	}
}
