package models

// BasicIAMRole models the json info of GCP IAM roles
type BasicIAMRole struct {
	Description         string   `json:"description"`
	Name                string   `json:"name"`
	Stage               string   `json:"stage"`
	Title               string   `json:"title"`
	IncludedPermissions []string `json:"includedPermissions"`
	Id                  int      `json:"id"`
}

// Role contains the IAM info of a Role plus some stats
type Role struct {
	BasicIAMRole

	Matches   int      `json:"matches"`
	MatchedBy []string `json:"matchedBy"`
	PercMatch float32  `json:"perc_match"`
}

// NewRole creates a new role pointing to a IAM resource
func NewRoleFromIAM(IAM BasicIAMRole) *Role {
	role := new(Role)
	role.BasicIAMRole = IAM
	return role
}
