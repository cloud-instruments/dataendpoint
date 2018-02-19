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


package controllers

import (
	"bytes"
	"cloudinstruments/libs/jwt"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
)

var AuthSecret, JwtSecret []byte

func init() {
	AuthSecret = []byte(os.Getenv("BASICAUTHSECRET"))
	JwtSecret = []byte(os.Getenv("JWTSECRET"))
	if AuthSecret == nil || JwtSecret == nil {
		fmt.Println("Failed to retrieve environment variables")
	}
}

var GetTokenHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		decodedSecret := r.Header.Get("Authentication")
		fmt.Println(decodedSecret)
		defer r.Body.Close()
		secret, errDecoding := base64.StdEncoding.DecodeString(decodedSecret)
		if errDecoding != nil {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		fmt.Println(secret)
		if !bytes.Equal(AuthSecret, secret) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		provider := jwt.NewJwtProvider(JwtSecret)
		signedToken, errGettingToken := provider.New()
		if errGettingToken != nil {
			http.Error(w, errGettingToken.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println(signedToken)
		w.Write([]byte(signedToken))
	})
