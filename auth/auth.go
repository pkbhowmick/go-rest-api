package auth

import (
	"encoding/base64"
	"net/http"
	"os"
	"strings"
)

func BasicAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get("Authorization")
		authStr := strings.Split(authHeader, " ")
		if len(authStr) < 2 || authStr[0] != "Basic" {
			http.Error(res, "Unauthorized", http.StatusUnauthorized)
			return
		}
		decodedStr, err := base64.StdEncoding.DecodeString(authStr[1])
		if err != nil {
			http.Error(res, "Unauthorized", http.StatusUnauthorized)
			return
		}
		temp := string(decodedStr)
		userAuth := strings.Split(temp, ":")
		if len(userAuth) != 2 {
			http.Error(res, "Unauthorized", http.StatusUnauthorized)
			return
		}
		username := userAuth[0]
		password := userAuth[1]

		if username != os.Getenv("ADMIN_USER") || password != os.Getenv("ADMIN_PASS") {
			http.Error(res, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(res, req)
	}
}
