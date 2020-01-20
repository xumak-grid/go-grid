package config

// User represents an user in aem
type User struct {
	Type        string                 `json:"type"`
	ID          string                 `json:"id"`
	LastName    string                 `json:"lastname,omitempty"`
	Password    string                 `json:"password,omitempty"`
	NewPassword string                 `json:"newPassword,omitempty"`
	Policy      string                 `json:"policy"`
	With        map[string]interface{} `json:"with,omitempty"`
	ACLS        *ACL                   `json:"acls,omitempty"`
	// The follwing keys are used to the response
	Ok    string `json:"ok,omitempty"`
	Error string `json:"error,omitempty"`
}

// ACL represents the ACLS of an user
type ACL struct {
	Deny  map[string]interface{} `json:"deny,omitempty"`
	Allow map[string]interface{} `json:"allow,omitempty"`
}

// user returns a new user with the configuration
func user(typ, id, password, policy string) *User {
	return &User{
		Type:        typ,
		ID:          id,
		Password:    password,
		NewPassword: password,
		Policy:      policy,
	}
}

// NewUser returns a new user with policy CREATE
func NewUser(id string, password string) *User {
	return user("internal", id, password, PolicyCreate)
}

// NewUserUpdate returns a new user with policy UPDATE
func NewUserUpdate(id string) *User {
	return user("internal", id, "", PolicyUpdate)
}

// NewUserChangePassword returns a new user with policy CHANGE_PASSWORD
func NewUserChangePassword(id string, newPassword string) *User {
	return user("root", id, newPassword, PolicyChangePassword)
}
