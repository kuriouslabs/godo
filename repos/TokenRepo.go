package repos

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Error constants
var (
	ErrTokenIsInvalid     = errors.New("Token is invalid")
	ErrTokenIsMissing     = errors.New("Token is missing")
	ErrTokenIsExpired     = errors.New("Token is expired")
	ErrTokenParsingFailed = errors.New("An uknown error occurred when parsing token")
)

// TokenRepo a repo for tokens
type TokenRepo struct{}

// NewTokenRepo creates a new token repo
func NewTokenRepo() *TokenRepo {
	return &TokenRepo{}
}

var signingString = []byte("1EguuHf87tJO7Z0p91b439PGtsqyw12V")

func keyParser(t *jwt.Token) (interface{}, error) {
	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
	}
	return signingString, nil
}

// UserIDFromToken extracts the user-id from the token
func UserIDFromToken(token *jwt.Token) (string, bool) {
	uid, ok := token.Claims["sub"].(string)
	return uid, ok
}

// TokenFromRequest extracts a token from the request
func TokenFromRequest(r *http.Request) (*jwt.Token, error) {
	return doTokenValidation(func() (*jwt.Token, error) {
		return jwt.ParseFromRequest(r, keyParser)
	})
}

// TokenFromString extracts a token from the string
func TokenFromString(token string) (*jwt.Token, error) {
	return doTokenValidation(func() (*jwt.Token, error) {
		return jwt.Parse(token, keyParser)
	})
}

type tokenGetAction func() (*jwt.Token, error)

func doTokenValidation(tokenGetter tokenGetAction) (*jwt.Token, error) {
	token, err := tokenGetter()

	switch err.(type) {
	case nil:
		return validateToken(token)
	case *jwt.ValidationError:
		return nil, ErrTokenIsExpired
	}

	//TODO: Clean this up
	if err == jwt.ErrNoTokenInRequest {
		return nil, ErrTokenIsMissing
	}

	return nil, ErrTokenParsingFailed
}

func validateToken(token *jwt.Token) (*jwt.Token, error) {
	if token.Valid {
		return token, nil
	}
	return nil, ErrTokenIsInvalid
}

// GenerateTokenForUser Generates a token for the user
func GenerateTokenForUser(uid string, expiration time.Time) string {
	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims["exp"] = expiration.Unix()
	token.Claims["iat"] = time.Now().Unix()
	token.Claims["sub"] = uid

	tokenString, err := token.SignedString(signingString)
	if err != nil {
		panic(err)
	}

	return tokenString
}
