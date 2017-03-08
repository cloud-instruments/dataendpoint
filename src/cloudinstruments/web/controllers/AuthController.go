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
