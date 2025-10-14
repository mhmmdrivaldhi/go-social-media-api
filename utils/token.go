package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mhmmdrivaldhi/go-social-media-api/config"
)

type JwtToken interface {
	GenerateToken(id int, email string) (string, error)
	ValidateToken(tokenString string) (*Claims, error)
}

type jwtToken struct{
	cfg config.JWTConfig
}

type Claims struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func (t *jwtToken) GenerateToken(id int, email string) (string, error) {
	claims := &Claims{
		ID:    id,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(t.cfg.AccessTokenLifeTime) * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Issuer: t.cfg.JWTSignatureKey,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(t.cfg.JWTSignatureKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return signedToken, nil
}

func (t *jwtToken) ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(jt *jwt.Token) (interface{}, error) {
		_, ok := jt.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", jt.Header["alg"])
		}

		return []byte(t.cfg.JWTSignatureKey), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, fmt.Errorf("token expired: %w", err)
		}

		if errors.Is(err, jwt.ErrTokenNotValidYet) {
			return nil, fmt.Errorf("token not active yet: %w", err)
		}

		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func NewJwtToken(cfg config.JWTConfig) JwtToken {
	return &jwtToken{cfg: cfg}
}
