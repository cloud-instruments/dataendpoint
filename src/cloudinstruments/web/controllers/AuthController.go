package controllers

import (
	"bytes"
	"cloudinstruments/libs/jwt"
	"encoding/base64"
	"fmt"
	"io/ioutil"
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

		defer r.Body.Close()
		d := base64.NewDecoder(base64.StdEncoding, r.Body)
		content, err := ioutil.ReadAll(d)
		if err != nil {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		if !bytes.Equal(AuthSecret, content) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		provider := jwt.NewJwtProvider(JwtSecret)
		signedToken, errGettingToken := provider.New()
		if errGettingToken != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println(signedToken)
		w.Write([]byte(signedToken))
	})
