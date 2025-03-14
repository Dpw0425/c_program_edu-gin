package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type AuthClaims struct {
	Guard string `json:"guard"`
	jwt.RegisteredClaims
}

type Options jwt.RegisteredClaims

func NewNumericData(t time.Time) *jwt.NumericDate {
	return jwt.NewNumericDate(t)
}

// GenerateToken 生成 JWT 令牌
func GenerateToken(guard string, secret string, opts *Options) string {
	claims := &AuthClaims{
		Guard: guard,
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  opts.Audience,
			ExpiresAt: opts.ExpiresAt,
			ID:        opts.ID,
			IssuedAt:  opts.IssuedAt,
			Issuer:    opts.Issuer,
			NotBefore: opts.NotBefore,
			Subject:   opts.Subject,
		},
	}

	tokenString, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))

	return tokenString
}

// ParseToken 解析 JWT Token
func ParseToken(tokenString string, secret string) (*AuthClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})

	if token == nil {
		return nil, fmt.Errorf("token is nil")
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
