package utils

import (
	"blogging-platform-api/internal/entity"
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func SignAccessToken(user *entity.User, secret string) (*jwt.RegisteredClaims, string, error) {
	secretKey := []byte(secret)

	claims := &jwt.RegisteredClaims{
		Issuer:    "blogging-platform-api",
		Subject:   strconv.FormatUint(uint64(user.ID), 10),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 1)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	sign, err := token.SignedString(secretKey)

	return claims, sign, err
}

func ParseAccessToken(tokenString string, secret string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, entity.ErrAuthTokenInvalid
		}
		return []byte(secret), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, entity.ErrAuthTokenExpired
		}
		return nil, entity.ErrAuthTokenInvalid
	}

	claim, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, entity.ErrAuthTokenInvalid
	}

	return claim, nil
}

func SignRefreshToken(user *entity.User, secret string) (*jwt.RegisteredClaims, string, error) {
	secretKey := []byte(secret)

	claims := &jwt.RegisteredClaims{
		Issuer:    "blogging-platform-api",
		Subject:   strconv.FormatUint(uint64(user.ID), 10),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
		ID:        uuid.NewString(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	sign, err := token.SignedString(secretKey)

	return claims, sign, err
}

func ParseRefreshToken(tokenString string, secret string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, entity.ErrAuthTokenInvalid
		}
		return []byte(secret), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, entity.ErrAuthTokenExpired
		}
		return nil, entity.ErrAuthTokenInvalid
	}

	claim, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, entity.ErrAuthTokenInvalid
	}

	return claim, nil
}
