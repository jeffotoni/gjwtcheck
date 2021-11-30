// Go Api server
// @jeffotoni

package muser

import mToken "github.com/jeffotoni/gjwtcheck/models/jwt"

// User structure
type User struct {
	Name      string `json:"name"`
	AvatarURL string `json:"user_avatar,omitempty"`
	Message   string `json:"message"`
	mToken.Response
}

type User2 struct {
	UserToken string `json:"user_token,omitempty"`
	Message   string `json:"message"`
	mToken.Response
}
