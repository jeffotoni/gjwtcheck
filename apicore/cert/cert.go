package cert

// Go Api server
// @jeffotoni

import (
	"crypto/rsa"

	"log"

	"github.com/dgrijalva/jwt-go"
)

// openssl genrsa -out private.rsa 1024
// openssl rsa -in private.rsa -pubout > public.rsa.pub
const (
	RSA_PRIVATE_AUTH = `-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDtmjOWlksmNouSxY1V3anvteiL8W6F4l93FPadnFF48Fxcmq6X
B/X4DS5PLqYI1lpMgHpykJH7t/hAP+/coN8QcseKSHMFbgqtzvR4KkkohjtZ3qpt
ahzCEvZoWWdMxfwToTJnghPUJLS0Z3ybPJ6gf5SWA4wItqHB2ORAVI+PRQIDAQAB
AoGADDfDs1ijIKHpZ3C1JdcyJ9toy9lsX4IWep8fV1d0cQ2bEj/5/lYwdcLnwLab
XxT2q7xnj4CToSgCofsnsGatUafX+daWU2zYK9i6HDmzP5KN8xTZE7GRtkIfOf1h
i4ic2OJlR6McaxMF2EIyfoJLsMzyqZHv4Fcr2PMrZ2Ntg0ECQQD3cU+oHpILzMcI
eux4iJ6r6CuJidvHAF9aXEd4pfucyNDtpsez/2/5nH5Bq0Icb6HCH0SCx8NLbZcQ
3Ks4i96ZAkEA9dHFckhyiQOpBiLfubE/ugAahW21I7AS0bWeehYs/Wpy65g7WbKv
3w90SI1h4KmEsVhCuGooBFVuKRXSRge9jQJBAMLkwcv2QCwBH6dDQqvxV6CpHlmX
dr64QWJnmnehrJuh5EW6/kJ86BJfu2Y34LuLBYpjnE+xzdXEPJ8wX8ALEskCQQC2
RcRlWEjtRbSII0XHOCHrtnXz6AQ/oShjLYuENXtIsSzjeq7PdQqpmJj9zy/7WVdV
9P8MJVLk27/iYRbY7JZlAkB40JkcjYOCGAGbolyo7cwKzA7QwM6uLSStbzFXw4wT
UTJQW4p6+ixfkwh0yWBB9iEIwFNah5lHmDSSIgpDe39m
-----END RSA PRIVATE KEY-----
`

	RSA_PUBLIC_AUTH = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDtmjOWlksmNouSxY1V3anvteiL
8W6F4l93FPadnFF48Fxcmq6XB/X4DS5PLqYI1lpMgHpykJH7t/hAP+/coN8QcseK
SHMFbgqtzvR4KkkohjtZ3qptahzCEvZoWWdMxfwToTJnghPUJLS0Z3ybPJ6gf5SW
A4wItqHB2ORAVI+PRQIDAQAB
-----END PUBLIC KEY-----
`

	RSA_PRIVATE = `-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDtmjOWlksmNouSxY1V3anvteiL8W6F4l93FPadnFF48Fxcmq6X
B/X4DS5PLqYI1lpMgHpykJH7t/hAP+/coN8QcseKSHMFbgqtzvR4KkkohjtZ3qpt
ahzCEvZoWWdMxfwToTJnghPUJLS0Z3ybPJ6gf5SWA4wItqHB2ORAVI+PRQIDAQAB
AoGADDfDs1ijIKHpZ3C1JdcyJ9toy9lsX4IWep8fV1d0cQ2bEj/5/lYwdcLnwLab
XxT2q7xnj4CToSgCofsnsGatUafX+daWU2zYK9i6HDmzP5KN8xTZE7GRtkIfOf1h
i4ic2OJlR6McaxMF2EIyfoJLsMzyqZHv4Fcr2PMrZ2Ntg0ECQQD3cU+oHpILzMcI
eux4iJ6r6CuJidvHAF9aXEd4pfucyNDtpsez/2/5nH5Bq0Icb6HCH0SCx8NLbZcQ
3Ks4i96ZAkEA9dHFckhyiQOpBiLfubE/ugAahW21I7AS0bWeehYs/Wpy65g7WbKv
3w90SI1h4KmEsVhCuGooBFVuKRXSRge9jQJBAMLkwcv2QCwBH6dDQqvxV6CpHlmX
dr64QWJnmnehrJuh5EW6/kJ86BJfu2Y34LuLBYpjnE+xzdXEPJ8wX8ALEskCQQC2
RcRlWEjtRbSII0XHOCHrtnXz6AQ/oShjLYuENXtIsSzjeq7PdQqpmJj9zy/7WVdV
9P8MJVLk27/iYRbY7JZlAkB40JkcjYOCGAGbolyo7cwKzA7QwM6uLSStbzFXw4wT
UTJQW4p6+ixfkwh0yWBB9iEIwFNah5lHmDSSIgpDe39m
-----END RSA PRIVATE KEY-----
`

	RSA_PUBLIC = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDtmjOWlksmNouSxY1V3anvteiL
8W6F4l93FPadnFF48Fxcmq6XB/X4DS5PLqYI1lpMgHpykJH7t/hAP+/coN8QcseK
SHMFbgqtzvR4KkkohjtZ3qptahzCEvZoWWdMxfwToTJnghPUJLS0Z3ybPJ6gf5SW
A4wItqHB2ORAVI+PRQIDAQAB
-----END PUBLIC KEY-----
`
)

var (
	//PrivateKey rsa
	PrivateKeyAuth *rsa.PrivateKey
	//PublicKey rsa
	PublicKeyAuth *rsa.PublicKey

	//PrivateKey rsa
	PrivateKey *rsa.PrivateKey
	//PublicKey rsa
	PublicKey *rsa.PublicKey
)

func init() {
	var err error
	publicByte := []byte(RSA_PUBLIC_AUTH)
	PublicKeyAuth, err = jwt.ParseRSAPublicKeyFromPEM(publicByte)
	if err != nil {
		log.Println("error ParseRSAPublicKeyFromPEM:", err.Error())
		return
	}
	PublicKey = PublicKeyAuth
	privateByte := []byte(RSA_PRIVATE_AUTH)
	PrivateKeyAuth, err = jwt.ParseRSAPrivateKeyFromPEM(privateByte)
	if err != nil {
		log.Println("error ParseRSAPrivateKeyFromPEM:", err.Error())
		return
	}
	PrivateKey = PrivateKeyAuth
}
