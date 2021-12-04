package data

import (
	"testing"
)

func TestDBParser(t *testing.T) {
	roles, err := db_parser(".")
	if len(roles) == 0 || err != nil {
		t.Fatalf("DB should load roles and return no error")
	}
}

func TestDBParserFailOnFolderAbsPath(t *testing.T) { //this doesn't fail in the right place
	roles, err := db_parser("@pippo//\\")
	if len(roles) != 0 || err == nil {
		t.Fatalf("DB should fail with error on wrong path input")
	}
}

func TestDBParserFailOnReadDir(t *testing.T) {
	roles, err := db_parser("/")
	if len(roles) != 0 || err == nil {
		t.Fatalf("DB should fail with error on unauthorized or non readable folder")
	}
}
