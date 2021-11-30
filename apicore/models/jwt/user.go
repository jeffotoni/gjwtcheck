// Go Api server
// @jeffotoni
// 2021-01-04

package models

//
// User structure
//
type User struct {

	//
	//
	//
	Login string `json:"login"`

	//
	//
	//
	Password string `json:"password,omitempty"`

	//
	//
	//
	Role string `json:"role"`
}
