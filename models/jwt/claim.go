package mjwt

// Go Api server
// @jeffotoni
// 2021-01-04
import jwt "github.com/dgrijalva/jwt-go"

// Claim structure, where we will use
// to validate our token with jwt

//Claim struct
type Claim struct {

	//
	//
	//
	User string `json:"user"`

	//
	//
	//
	jwt.StandardClaims
}
