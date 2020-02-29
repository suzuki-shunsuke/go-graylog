package graylog

type (
	// UserToken is an access token for a user
	UserToken struct {
		Name       string `json:"name"`
		Token      string `json:"token"`
		LastAccess string `json:"last_access,omitempty"`
	}
)
