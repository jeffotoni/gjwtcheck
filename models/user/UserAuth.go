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
}
