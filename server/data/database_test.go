package data

import (
	"testing"
)

func TestDBParser(t *testing.T) {
	roles, err := db_parser()
	if len(roles) == 0 || err != nil {
		t.Fatalf("DB should load roles and return no error")
	}
}
