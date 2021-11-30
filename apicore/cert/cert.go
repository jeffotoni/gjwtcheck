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
	MIICXAIBAAKBgQCkT6K9UpeDfAIM9OOayQnHQHs7yGh05Gm8Ha7/R1lSOk22duFm
	O0xwh+gtcFGxBJF/o2Y12tCVHFxfqtjzGkSkX5gGaPmJmSrz2NOi/ANrnjGxNhgg
	tZM+nLJfRCt7/v9RVjh+ZR049+CCOqX2YSn/2GILufNMWhZArs3b8x1aBwIDAQAB
	AoGAT7t9IxXDBEDW1ea309KedhvqMPZtCdwVGv3kf2fmBirYryQ8+FjqHOk2V2Zb
	neSznhZycMu/z7u1b47+O+mRSNMxzlS93uEB0GG8it48vX7rRE43+RbCOrsNmU9s
	qAkMtq6Wqkhp7raSaRGNMLguUYwd5YpA1ii6CtXrXidESdECQQDPy/bzj/+sFENS
	rPK0qoFs9Ndls8fRAtph/O0ZIWudLwLcbpfnHN6gHLf8txratL1Bcs81PPLD7f4d
	48bQfGNLAkEAym1FCKRB0zHi//xPPXL/k0tYw35zRjKVSmc7KeFiddXpT7Wj2OHJ
	aPdcJy6kW137dT4UOEjOCeEzKLdNxMOytQJAWjV+wN/7q5WkYGjqeJoo08c7F1DD
	5y3o5m8p8yX6FOPKxy1PzqpOz42IJjLLerTKEHaqE7+g2IQiNJGkxu+pHwJBAMG6
	m8Pud7To3IgC788ufNx50sSeAzKefHRNobiuJG4DwBtyChIp3HlhqscxA0kSA/Mr
	62wGeXHJHg4MiBgU9rUCQGuRsuI88zTNxDgTp9FVzQeMWBRDCGPUisIOmHqGJCJc
	rUvBtS3WTHtYySILiv+TA1vrtXno3FzgWf/zsnilemQ=
	-----END RSA PRIVATE KEY-----
	`

	RSA_PUBLIC_AUTH = `-----BEGIN PUBLIC KEY-----
	MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCkT6K9UpeDfAIM9OOayQnHQHs7
	yGh05Gm8Ha7/R1lSOk22duFmO0xwh+gtcFGxBJF/o2Y12tCVHFxfqtjzGkSkX5gG
	aPmJmSrz2NOi/ANrnjGxNhggtZM+nLJfRCt7/v9RVjh+ZR049+CCOqX2YSn/2GIL
	ufNMWhZArs3b8x1aBwIDAQAB
	-----END PUBLIC KEY-----
	`

	RSA_PRIVATE = `-----BEGIN RSA PRIVATE KEY-----
	MIICXAIBAAKBgQCkT6K9UpeDfAIM9OOayQnHQHs7yGh05Gm8Ha7/R1lSOk22duFm
	O0xwh+gtcFGxBJF/o2Y12tCVHFxfqtjzGkSkX5gGaPmJmSrz2NOi/ANrnjGxNhgg
	tZM+nLJfRCt7/v9RVjh+ZR049+CCOqX2YSn/2GILufNMWhZArs3b8x1aBwIDAQAB
	AoGAT7t9IxXDBEDW1ea309KedhvqMPZtCdwVGv3kf2fmBirYryQ8+FjqHOk2V2Zb
	neSznhZycMu/z7u1b47+O+mRSNMxzlS93uEB0GG8it48vX7rRE43+RbCOrsNmU9s
	qAkMtq6Wqkhp7raSaRGNMLguUYwd5YpA1ii6CtXrXidESdECQQDPy/bzj/+sFENS
	rPK0qoFs9Ndls8fRAtph/O0ZIWudLwLcbpfnHN6gHLf8txratL1Bcs81PPLD7f4d
	48bQfGNLAkEAym1FCKRB0zHi//xPPXL/k0tYw35zRjKVSmc7KeFiddXpT7Wj2OHJ
	aPdcJy6kW137dT4UOEjOCeEzKLdNxMOytQJAWjV+wN/7q5WkYGjqeJoo08c7F1DD
	5y3o5m8p8yX6FOPKxy1PzqpOz42IJjLLerTKEHaqE7+g2IQiNJGkxu+pHwJBAMG6
	m8Pud7To3IgC788ufNx50sSeAzKefHRNobiuJG4DwBtyChIp3HlhqscxA0kSA/Mr
	62wGeXHJHg4MiBgU9rUCQGuRsuI88zTNxDgTp9FVzQeMWBRDCGPUisIOmHqGJCJc
	rUvBtS3WTHtYySILiv+TA1vrtXno3FzgWf/zsnilemQ=
	-----END RSA PRIVATE KEY-----`

	RSA_PUBLIC = `-----BEGIN PUBLIC KEY-----
	MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCkT6K9UpeDfAIM9OOayQnHQHs7
	yGh05Gm8Ha7/R1lSOk22duFmO0xwh+gtcFGxBJF/o2Y12tCVHFxfqtjzGkSkX5gG
	aPmJmSrz2NOi/ANrnjGxNhggtZM+nLJfRCt7/v9RVjh+ZR049+CCOqX2YSn/2GIL
	ufNMWhZArs3b8x1aBwIDAQAB
	-----END PUBLIC KEY-----`
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
