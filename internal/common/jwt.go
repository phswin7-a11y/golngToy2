package common

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("toy_login!@#")

type Claims struct {
	UserIdx uint   `json:"user_idx"`
	UserID  string `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateToken 토큰 생성 (AccessToken: 1시간, RefreshToken: 3시간)
func GenerateToken(userIdx uint, userID string) (string, string, error) {
	// Access Token
	accessTokenClaims := &Claims{
		UserIdx: userIdx,
		UserID:  userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	at, err := accessToken.SignedString(secretKey)
	if err != nil {
		return "", "", err
	}

	// Refresh Token
	refreshTokenClaims := &Claims{
		UserIdx: userIdx,
		UserID:  userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	rt, err := refreshToken.SignedString(secretKey)
	if err != nil {
		return "", "", err
	}

	return at, rt, nil
}

// ValidateToken 토큰 검증
func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
