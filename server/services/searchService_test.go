package services

import (
	data "dsdr/data"

	"testing"
)

func TestSearchRoleDefaultValues(t *testing.T) {
	// default values
	var testSearchString string
	var testDB *data.DB

	roles, err := SearchRole(testSearchString, testDB)
	if roles != nil || err == nil {
		t.Fatalf("nil DB pointer doesn't return error")
	}
}

func TestSearchRoleStringInited(t *testing.T) {
	// string custom but db default
	var testSearchString = "compute"
	var testDB *data.DB

	roles, err := SearchRole(testSearchString, testDB)
	if roles != nil || err == nil {
		t.Fatalf("nil DB pointer doesn't return error")
	}
}

func TestSearchRoleDBInited(t *testing.T) {
	// string default, DB inited
	var testSearchString string
	testDB, _ := data.NewDB()

	roles, err := SearchRole(testSearchString, testDB)
	if len(roles) != len(testDB.Roles) || err != nil {
		t.Fatalf("empty string should return no error and the entire DB")
	}
}

func TestSearchRoleAllInited(t *testing.T) {
	//both properly inited - here we have to mock the repo...
	//TODO: this is a wrong test: depending on the string result could be empty
	var testSearchString = "compute network storage"
	testDB, _ := data.NewDB()

	roles, err := SearchRole(testSearchString, testDB)
	if len(roles) <= 0 || len(roles) >= len(testDB.Roles) || err != nil {
		t.Fatalf("inited string should return no error and a fraction of the DB")
	}
}
