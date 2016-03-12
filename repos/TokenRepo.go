package repos

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var signingString = []byte("hello")

func keyParser(t *jwt.Token) (interface{}, error) {
	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
	}
	return signingString, nil
}

func UserIDFromToken(token *jwt.Token) (string, bool) {
	uid, ok := token.Claims["sub"].(string)
	return uid, ok
}

func ValidateTokenFromRequest(r *http.Request) (int, *jwt.Token) {
	token, err := jwt.ParseFromRequest(r, keyParser)

	switch err.(type) {
	case nil:
		return validateToken(token)
	case *jwt.ValidationError:
		return http.StatusUnauthorized, nil
	}

	//TODO: Clean this up
	if err == jwt.ErrNoTokenInRequest {
		return http.StatusUnauthorized, nil
	}

	return http.StatusInternalServerError, nil // unknown error
}

func validateToken(token *jwt.Token) (int, *jwt.Token) {
	if token.Valid {
		return http.StatusOK, token
	}
	return http.StatusUnauthorized, nil
}

////////// TEMP
func GetToken(uuid string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	// t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	// token.Claims["exp"] = t.Add(time.Hour * time.Duration(72)).Unix()

	token.Claims["exp"] = time.Now().Add(time.Second * time.Duration(60)).Unix()
	token.Claims["iat"] = time.Now().Unix()
	token.Claims["sub"] = uuid

	tokenString, err := token.SignedString(signingString)
	if err != nil {
		panic(err)
	}

	return tokenString, nil
}
