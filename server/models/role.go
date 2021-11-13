package models

// Role models the json info of GCP IAM roles
type Role struct {
	Description         string   `json:"description"`
	Name                string   `json:"name"`
	Stage               string   `json:"stage"`
	Title               string   `json:"title"`
	IncludedPermissions []string `json:"includedPermissions"`
	Id                  int      `json:"id"`
}
