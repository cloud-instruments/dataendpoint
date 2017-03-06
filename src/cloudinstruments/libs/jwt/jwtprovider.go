package jwt

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

type JwtProvider struct {
	secret []byte
}

func NewJwtProvider(secretVal []byte) *JwtProvider {
	return &JwtProvider{secret: secretVal}
}

func getExpirationDateTime() int64 {
	return time.Now().Add(time.Minute * 1).Unix()
}

func (j *JwtProvider) New() (string, error) {
	claims := &jwt.StandardClaims{
		Issuer:    "CloudInstruments",
		ExpiresAt: getExpirationDateTime(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, errSigning := token.SignedString(j.secret)
	if errSigning != nil {
		return "", errSigning
	}

	return signedToken, nil
}

func (j *JwtProvider) IsValidToken(tokenToCheck string) bool {
	if tokenToCheck == "" {
		return false
	}

	token, err := jwt.Parse(tokenToCheck, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		fmt.Println(j.secret)
		return j.secret, nil
	})

	return token != nil && token.Valid && err == nil
}
