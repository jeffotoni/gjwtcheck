package jwt

// @jeffotoni

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	//"github.com/golang-jwt/jwt/v4"
	"github.com/gofiber/utils"
	cert "github.com/jeffotoni/gjwtcheck/core/cert"
	mJwt "github.com/jeffotoni/gjwtcheck/core/models/jwt"

	"log"

	"github.com/jeffotoni/gjwtcheck/core/pkg/fmts"
)

var (
	expires = int64(0)
	nbf     = int64(0)
)

// iat set to now
// nbf set to tomorrow 12:00pm
// exp set to tomorrow 1:00pm

func SetExpires(second int) {
	duration, _ := time.ParseDuration("-1.0h")
	nbf = time.Now().Add(duration).Unix()
	expires = time.Now().Add(time.Second * time.Duration(second)).Unix()
}

func Token(user string, IP string) (string, string, string, string, error) {
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
			NotBefore: nbf,
			IssuedAt:  time.Now().Unix(),
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
		return "", "", "", "", err
	}

	//zerar
	expires = int64(0)

	// return token string
	// rsaP := cert.PrivateKey.Public()
	// rsaP.(string)
	return cert.RSA_PRIVATE, cert.RSA_PUBLIC, tokenString, expiresData, nil
}

func TokenHS256(user string, IP string) (string, string, string, error) {
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
			NotBefore: nbf,
			IssuedAt:  time.Now().Unix(),
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
		return "", "", "", err
	}

	//zerar
	expires = int64(0)

	// return token string
	return cert.SecretSH256, tokenString, expiresData, nil
}
