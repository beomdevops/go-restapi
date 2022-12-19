package service

import (
	"encoding/base64"
	"encoding/binary"
	"io/ioutil"

	jwt "github.com/golang-jwt/jwt/v4"
)

const (
	PrivKeyPath = "key/private.pem" // openssl genrsa -out app.rsa keysize
	PubKeyPath  = "key/public.pem"  // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

type MyCustomClaims struct {
	*jwt.RegisteredClaims
	FooClaim int
	BarClaim string
}

func GenJwt() string {
	//expiration := time.Now().Add(time.Second * 3600)
	claims := MyCustomClaims{
		RegisteredClaims: &jwt.RegisteredClaims{
			Issuer:  "ISSUER",
			Subject: "JWT Creation",
		},
		FooClaim: 123,
		BarClaim: "bar",
	}
	var pri_pem, err = ioutil.ReadFile(PrivKeyPath)
	if err != nil {
		panic(err)
	}
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(pri_pem)
	if err != nil {
		panic(err)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	myJwt, _ := token.SignedString(signKey)

	return myJwt
}

func GenJwk() (string, string) {
	//expiration := time.Now().Add(time.Second * 3600)

	var pubpem, err = ioutil.ReadFile(PubKeyPath)
	if err != nil {
		panic(err)
	}
	signKey, err := jwt.ParseRSAPublicKeyFromPEM(pubpem)
	if err != nil {
		panic(err)
	}

	n := base64.RawURLEncoding.EncodeToString(signKey.N.Bytes())
	e := base64.RawURLEncoding.EncodeToString(convertEtoBytes(signKey.E))

	return n, e
}

func convertEtoBytes(e int) []byte {
	e_byte := make([]byte, 4)

	binary.BigEndian.PutUint32(e_byte, uint32(e))

	e_byte = e_byte[1:]
	return e_byte
}
