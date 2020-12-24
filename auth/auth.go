package auth

import (
	"encoding/base64"
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
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

func GenerateToken(userID string) (string, error) {
	signingKey := []byte(os.Getenv("SIGNING_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		Id:        userID,
		Issuer:    "Admin",
	})
	signedToken, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func JwtAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get("Authorization")
		if len(authHeader) == 0 {
			http.Error(res, "Access token is missing or invalid", http.StatusUnauthorized)
			return
		}
		authStr := strings.Split(authHeader, " ")
		if len(authStr) < 2 || authStr[0] != "Bearer" {
			http.Error(res, "Access token is missing or invalid", http.StatusUnauthorized)
			return
		}
		tokenString := authStr[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("error")
			}
			return []byte(os.Getenv("SIGNING_KEY")), nil
		})
		if err != nil {
			http.Error(res, "Access token is missing or invalid", http.StatusUnauthorized)
		} else if _, ok := err.(*jwt.ValidationError); ok {
			http.Error(res, "Access token is missing or invalid", http.StatusUnauthorized)
		} else if token.Valid {
			next.ServeHTTP(res, req)
		} else {
			http.Error(res, "Access token is missing or invalid", http.StatusUnauthorized)
		}

	}
}
