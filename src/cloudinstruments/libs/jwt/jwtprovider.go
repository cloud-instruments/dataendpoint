Copyright (c) <2018> <University of Washington>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

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
