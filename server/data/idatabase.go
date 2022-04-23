package data

import (
	"dsdr/models"
)

// methods of DB objects
type DB interface {
	Connect(string) error
	Roles() []models.BasicIAMRole
}
