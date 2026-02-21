package entity

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthRegisterReq struct {
	FirstName      string `json:"firstName" binding:"required"`
	LastName       string `json:"lastName" binding:"required"`
	Username       string `json:"username"`
	Email          string `json:"email" binding:"required,email"`
	Password       string `json:"password" binding:"required"`
	HashedPassword string `json:"hashed_password"`
}

type AuthLoginReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthLoginResp struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	User         *User  `json:"user"`
}

type AuthRepository interface {
	Register(ctx context.Context, user *User) error
	Login(ctx context.Context, email string) (*User, error)
	SaveRefreshToken(ctx context.Context, rtk *RefreshToken) error
	RefreshToken(ctx context.Context, claim *jwt.RegisteredClaims) (*User, *RefreshToken, error)
	UpdateRefreshToken(ctx context.Context, rtk *RefreshToken) error
}

type AuthUsecase interface {
	Register(req *AuthRegisterReq) error
	Login(req *AuthLoginReq) (*AuthLoginResp, error)
	RefreshToken(rtk string) (string, string, error)
}

type AuthHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	RefreshToken(c *gin.Context)
}
