package jwt

// @jeffotoni

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	//"github.com/golang-jwt/jwt/v4"
	"github.com/gofiber/utils"
	cert "github.com/jeffotoni/gjwtcheck/apicore/cert"
	mJwt "github.com/jeffotoni/gjwtcheck/apicore/models/jwt"

	"log"

	"github.com/jeffotoni/gjwtcheck/apicore/pkg/fmts"
)

var (
	expires = int64(0)
)

func SetExpires(second int) {
	expires = time.Now().Add(time.Second * time.Duration(second)).Unix()
}

func Token(user string, IP string) (string, string, error) {
	if expires == 0 {
		expires = time.Now().Add(time.Minute * 6).Unix()
	}

	layout := "2006-01-02 15:04:05"
	var err error

	t := time.Unix(expires, 0)
	expiresData := t.Format(layout)

	claims := mJwt.Claim{
		User: user,
		Id:   utils.UUID(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expires,
			Issuer: fmts.ConcatStr("gjwtcheck - created in:", time.Now().Format("2006-01-02 15:04:05"),
				" expires:", expiresData),
		},
	}

	// Generating token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	// Transforming into string
	tokenString, err := token.SignedString(cert.PrivateKey)
	if err != nil {
		log.Println("err:", err.Error())
		return "", "", err
	}

	//zerar
	expires = int64(0)

	// return token string
	return tokenString, expiresData, nil
}

func TokenHS256(user string, IP string) (string, string, error) {
	if expires == 0 {
		expires = time.Now().Add(time.Minute * 4).Unix()
	}

	layout := "2006-01-02 15:04:05"
	var err error

	t := time.Unix(expires, 0)
	expiresData := t.Format(layout)

	claims := mJwt.Claim{
		User: user,
		Id:   utils.UUID(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expires,
			Issuer: fmts.ConcatStr("gjwtcheck - created in:", time.Now().Format("2006-01-02 15:04:05"),
				" expires:", expiresData),
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	tokenString, err := token.SignedString([]byte(cert.SecretSH256))
	if err != nil {
		log.Println("Error SignedString Token:", err.Error())
		return tokenString, expiresData, err
	}

	//zerar
	expires = int64(0)

	// return token string
	return tokenString, expiresData, nil
}
