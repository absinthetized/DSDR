package data

import (
	"testing"
)

func TestDBParser(t *testing.T) {
	var DB DB
	err := DB.Connect(".")
	if len(DB.Roles) == 0 || err != nil {
		t.Fatalf("DB should load roles and return no error")
	}
}

func TestDBParserFailOnFolderAbsPath(t *testing.T) { //this doesn't fail in the right place
	var DB DB
	err := DB.Connect("@pippo//\\")
	if len(DB.Roles) != 0 || err == nil {
		t.Fatalf("DB should fail with error on wrong path input")
	}
}

func TestDBParserFailOnReadDir(t *testing.T) {
	var DB DB
	err := DB.Connect("/")
	if len(DB.Roles) != 0 || err == nil {
		t.Fatalf("DB should fail with error on unauthorized or non readable folder")
	}
}
