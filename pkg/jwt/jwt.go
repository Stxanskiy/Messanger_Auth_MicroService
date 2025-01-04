package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken = errors.New("invalid token")
)

type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

type JWTManager struct {
	SecretKey       string
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

func NewJWTManager(secretKey string, accessTTL, refreshTTL time.Duration) *JWTManager {
	return &JWTManager{
		SecretKey:       secretKey,
		AccessTokenTTL:  accessTTL,
		RefreshTokenTTL: refreshTTL,
	}
}

func (j *JWTManager) GenerateTokens(userID int) (*Claims, *Claims, error) {
	now := time.Now()

	accessClaims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(j.AccessTokenTTL)),
		},
	}

	refreshClaims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(j.RefreshTokenTTL)),
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString([]byte(j.SecretKey))
	if err != nil {
		return nil, nil, err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(j.SecretKey))
	if err != nil {
		return nil, nil, err
	}

	accessClaims.ID = accessToken
	refreshClaims.ID = refreshToken

	return accessClaims, refreshClaims, nil
}

func (j *JWTManager) ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
