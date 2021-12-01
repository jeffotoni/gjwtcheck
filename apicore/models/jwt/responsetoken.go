// Go Api server
// @jeffotoni
// 2021-01-04

package models

//
// ResponseToken
//
type ResponseToken struct {

	//
	// token
	//
	Token string `json:"token,omitempty"`

	Expires string `json:"expires,omitempty"`

	Message string `json:"message"`
}
