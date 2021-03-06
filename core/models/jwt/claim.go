// Go Api server
// @jeffotoni
// 2021-01-04

package models

import jwt "github.com/dgrijalva/jwt-go"

// Claim structure, where we will use
// to validate our token with jwt

//
// jwt
//
type Claim struct {

	//
	//
	//
	User string `json:"user"`

	Id string `json:"id"`

	Key string `json:"key"`

	//
	//
	//
	jwt.StandardClaims
}
