package services

import (
	data "dsdr/data"

	"testing"
)

func TestSearchRole(t *testing.T) {
	// default values
	var testSearchString string
	var testDB *data.DB

	roles, err := SearchRole(testSearchString, testDB)
	if roles != nil || err == nil {
		t.Fatalf("nil DB pointer doesn't return error")
	}

	// string custom but db default
	testSearchString = "compute"
	roles, err = SearchRole(testSearchString, testDB)
	if roles != nil || err == nil {
		t.Fatalf("nil DB pointer doesn't return error")
	}

	// string default, DB inited
	var testSearchString2 string
	testDB, _ = data.NewDB()
	roles, err = SearchRole(testSearchString2, testDB)
	if len(roles) != len(testDB.Roles) || err != nil {
		t.Fatalf("empty string should return no error and the entire DB")
	}

	//both properly inited - here we have to mock the repo...
	//TODO: this is a wrong test: depending on the string result could be empty
	roles, err = SearchRole(testSearchString, testDB)
	if len(roles) <= 0 || len(roles) >= len(testDB.Roles) || err != nil {
		t.Fatalf("inited string should return no error and a fraction of the DB")
	}
}
