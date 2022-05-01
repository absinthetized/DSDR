package models

import "strings"

// basic type to be derived/extended
type abstractIAMRole struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	Stage       string `json:"stage"`
	Title       string `json:"title"`
	Id          int    `json:"id"`
}

// BqIAMRole maps the info of GCP IAM roles stored in big query
type BqIAMRole struct {
	abstractIAMRole

	Included_permissions string // bq doesn't store arrays in a field
}

// BasicIAMRole models the json info of GCP IAM roles
// it is expected to be a method in place able to map the BqIAMRole Included_permissions string
// of comma separeted roles into an array of strings as the one in BasicIAMRole
type BasicIAMRole struct {
	abstractIAMRole

	IncludedPermissions []string `json:"includedPermissions"`
}

func NewIAMfromBq(bq BqIAMRole) *BasicIAMRole {
	role := new(BasicIAMRole)
	role.abstractIAMRole = bq.abstractIAMRole
	role.IncludedPermissions = append(role.IncludedPermissions, strings.Split(bq.Included_permissions, ",")...)
	return role
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
