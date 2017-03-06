package handlers

import (
	jwtprovider "cloudinstruments/libs/jwt"
	"fmt"
	"net/http"
	"os"
)

var JwtSecret []byte

func init() {
	JwtSecret = []byte(os.Getenv("JWTSECRET"))
}

func AuthHandler(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		provider := jwtprovider.NewJwtProvider(JwtSecret)
		authHeader := r.Header.Get("Authorization")
		fmt.Println(authHeader)
		if !provider.IsValidToken(authHeader) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
