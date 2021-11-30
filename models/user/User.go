// Go Api server
// @jeffotoni

package muser

import mToken "github.com/jeffotoni/gjwtcheck/models/jwt"

// User structure
type User struct {
	RoleName  string `json:"role_name"`
	UserToken string `json:"user_token"`
	RoleValue string `json:"role_value"`
	MetaKey   string `json:"meta_key"`
	MetaValue string `json:"meta_value"`
	AvatarURL string `json:"user_avatar,omitempty"`
	Message   string `json:"message"`
	mToken.Response
}

type User2 struct {
	UserToken string `json:"user_token,omitempty"`
	Message   string `json:"message"`
	mToken.Response
}
