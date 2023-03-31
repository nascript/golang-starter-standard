package token

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// HOW TO USE THIS PACKAGE/FUNCTION?:
//
// 	jwt := token.JSONWebToken{
//		Issuer:    "skilledin.io",
//		SecretKey: []byte(config.Instance.AccessSecret),
//		IssuedAt:  time.Now(),
//		ExpiredAt: time.Now().Add(1 * time.Minute),
//	}
//  payload := map[string]any{}{
//    "tenant_id": "123-456-tenant-id",
//    "user_id": "123-456-user-id",
//    ... and so on.
// }
//  accessToken, err := jwt.Claim(payload)
//

type IJSONWebToken interface {
	Claim(payload interface{}) (string, error)
}

type JSONWebTokenClaim struct {
	jwt.RegisteredClaims
	Payload interface{} `json:"payload"`
}

type JSONWebToken struct {
	Issuer    string
	SecretKey []byte
	IssuedAt  time.Time
	ExpiredAt time.Time
}

func (j *JSONWebToken) Claim(payload interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JSONWebTokenClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.Issuer,
			IssuedAt:  &jwt.NumericDate{Time: j.IssuedAt},
			ExpiresAt: &jwt.NumericDate{Time: j.ExpiredAt},
		},
		Payload: payload,
	})

	return token.SignedString(j.SecretKey)
}
