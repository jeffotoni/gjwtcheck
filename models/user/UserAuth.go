// Go Api server
// @jeffotoni
package muser

type UserAuth struct {
	User     string `json:"user,omitempty"`
	Key      string `json:"token,omitempty"`
	Time     int    `json:"time,omitempty"`
	Public   string `json:"public,omitempty"`
	Private  string `json:"private,omitempty"`
	Secret   string `json:"secret,omitempty"`
	Password string `json:"password,omitempty"`
	Expires  string `json:"expires,omitempty"`
	Typex    string `json:"typex,omitempty"`
	MyClaim  string `json:"claims,omitempty"`
}

type MyClaim struct {
	Audience  string `json:"aud,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	Id        string `json:"jti,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	Issuer    string `json:"iss,omitempty"`
	NotBefore int64  `json:"nbf,omitempty"`
	Subject   string `json:"sub,omitempty"`
}
