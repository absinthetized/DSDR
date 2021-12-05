package services

import (
	mocks "dsdr/mocks/data"
	models "dsdr/models"

	"testing"
)

func TestSearchRoleDBInited(t *testing.T) {
	// string default, DB inited
	var testSearchString string

	var testDB mocks.DB
	testDB.On("Connect", "../roles").Return(nil)
	testDB.On("Roles").Return([]models.BasicIAMRole{
		*new(models.BasicIAMRole),
		*new(models.BasicIAMRole),
	})

	roles, err := SearchRole(testSearchString, &testDB)
	if len(roles) != len(testDB.Roles()) || err != nil {
		t.Fatalf("empty string should return no error and the entire DB")
	}
}

func TestSearchRoleAllWellInited(t *testing.T) {
	//both properly inited - here we have to mock the repo...
	//TODO: this is a wrong test: depending on the string result could be empty
	var testSearchString = "compute network storage"

	pippo := *new(models.BasicIAMRole)
	pippo.IncludedPermissions = []string{"computeUser", "networkUser"}
	var testDB mocks.DB
	testDB.On("Connect", "../roles").Return(nil)
	testDB.On("Roles").Return([]models.BasicIAMRole{
		*new(models.BasicIAMRole),
		pippo,
	})

	roles, err := SearchRole(testSearchString, &testDB)
	if len(roles) <= 0 || len(roles) > len(testDB.Roles()) || err != nil {
		t.Fatalf("inited string should return no error and a fraction of the DB")
	}
}
