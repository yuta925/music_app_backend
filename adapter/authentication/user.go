package authentication

import (
	"fmt"
	"music-app/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var userTokenJwt = &hs256jwt{
	sigKey: []byte(config.SigKey()),
	createClaims: func() jwt.Claims {
		return &userClaims{}
	},
}

type userClaims struct {
	ID        string   `json:"jti"`
	Subject   string   `json:"sub"`
	IssuedAt  int      `json:"iat"`
	ExpiredAt int      `json:"exp"`
}

func IssueUserToken(userID string) (string, error) {
	id := uuid.New()


	expiredAt := time.Date(9999,11,1,1,1,1,1,time.UTC)
	claims := &userClaims{
		ID:        id.String(),
		Subject:   userID,
		IssuedAt:  int(time.Now().Unix()),
		ExpiredAt: int(expiredAt.Unix()),
	}

	return userTokenJwt.issueToken(claims)
}

func (c *userClaims) Valid() error {

	_, err := uuid.Parse(c.ID)
	if err != nil {
		return fmt.Errorf("invalid id=%s: %w", c.ID, err)
	}

	return nil
}

func VerifyUserToken(token string) (string, error) {
	claims, err := userTokenJwt.verifyToken(token)
	if err != nil {
		return "", err
	}

	return claims.(*userClaims).Subject, nil
}

// 配列a が 配列b を包含しているかどうかを返す
func isInclusive(a, b []string) bool {
	for _, v := range b {
		if !contains(a, v) {
			return false
		}
	}
	return true
}

func contains(s []string, e string) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
