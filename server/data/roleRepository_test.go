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
	var testDB DB
	testDB.Connect("../roles")
	repo := NewRoleRepository(&testDB)
	roles := repo.FindAll()
	if len(roles) == 0 {
		t.Fatalf("find all shloud return non empty list of roles")
	}
}

func TestFindPermissionsByRegexArrayEmptyArray(t *testing.T) {
	var emptyInput []string

	var testDB DB
	testDB.Connect("../roles")
	repo := NewRoleRepository(&testDB)
	roles, err := repo.FindPermissionsByRegexArray(emptyInput)

	if len(roles) != len(testDB.Roles) || err != nil {
		t.Fatalf("find by permissions should return all results and no error if an empty array is passed")
	}
}

func TestFindPermissionsByRegexArrayFilledArray(t *testing.T) {
	input := []string{"compute", "network"}

	var testDB DB
	testDB.Connect("../roles")
	repo := NewRoleRepository(&testDB)
	roles, err := repo.FindPermissionsByRegexArray(input)

	if len(roles) == 0 || err != nil {
		t.Fatalf("find by permissions should return some results and no error if a well inited array is passed")
	}
}
