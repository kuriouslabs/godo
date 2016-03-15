package repos

import (
	"testing"
	"time"
)

const testUserID = "user-id"
const testUserName = "user-name"

func TestGenerateTokenForUserReturnsValidToken(t *testing.T) {
	exp := time.Now().Add(time.Second * time.Duration(60))

	tokenStr := GenerateTokenForUser(testUserID, exp)
	if _, err := TokenFromString(tokenStr); err != nil {
		t.Error("Should generate a valid token")
	}
}

func makeTokenRepo() *TokenRepo {
	repo := NewTokenRepo()
	return repo
}
