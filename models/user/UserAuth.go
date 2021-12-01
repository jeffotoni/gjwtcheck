// Go Api server
// @jeffotoni
package muser

type UserAuth struct {
	User     string `json:"user,omitempty"`
	Key      string `json:"key,omitempty"`
	Public       string `json:"public,omitempty"`
	Password string `json:"password,omitempty"`
	Expires  string `json:"expires,omitempty"`
	Typex    string `json:"typex,omitempty"`
}
