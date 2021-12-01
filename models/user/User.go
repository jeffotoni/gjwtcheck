// Go Api server
// @jeffotoni

package muser

import mToken "github.com/jeffotoni/gjwtcheck/apicore/models/jwt"

// User structure
type User struct {
	Name      string `json:"name,omitempty"`
	User      string `json:"user,omitempty"`
	Id        string `json:"id,omitempty"`
	Iss       string `json:"iss,omitempty"`
	AvatarURL string `json:"avatar,omitempty"`
	Message   string `json:"message"`
	mToken.ResponseToken
}

type User2 struct {
	UserToken string `json:"user_token,omitempty"`
	Message   string `json:"message"`
	mToken.ResponseToken
}
