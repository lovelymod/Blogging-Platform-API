package usecase

import (
	"blogging-platform-api/internal/entity"
	"blogging-platform-api/pkg/utils"
	"context"
	"log"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type authUsercase struct {
	repo               entity.AuthRepository
	timeout            time.Duration
	hashCost           int
	accessTokenSecret  string
	refreshTokenSecret string
}

func NewAuthUsecase(repo entity.AuthRepository, contextTimeout time.Duration, config *entity.Config) entity.AuthUsecase {
	cost, _ := strconv.Atoi(config.HASH_COST)
	if cost < 12 {
		cost = 12
	}

	return &authUsercase{
		repo:               repo,
		timeout:            contextTimeout,
		hashCost:           cost,
		accessTokenSecret:  config.ACCESS_TOKEN_SECRET,
		refreshTokenSecret: config.REFRESH_TOKEN_SECRET,
	}
}

func (u *authUsercase) Register(req *entity.AuthRegisterReq) error {
	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	hasdPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), u.hashCost)
	if err != nil {
		log.Println(err)
		return entity.ErrGlobalServerErr
	}

	user := &entity.User{
		Email:          req.Email,
		HashedPassword: string(hasdPassword),
		FirstName:      req.FirstName,
		LastName:       req.LastName,
	}

	if req.Username != "" {
		user.Username = req.Username
	} else {
		user.Username = "user" + strconv.FormatInt(time.Now().Unix(), 10)
	}

	return u.repo.CreateUser(ctx, user)
}

func (u *authUsercase) Login(req *entity.AuthLoginReq) (*entity.AuthLoginResp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	existUser, err := u.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existUser.HashedPassword), []byte(req.Password)); err != nil {
		log.Println(err)
		return nil, entity.ErrAuthWrongEmailOrPassword
	}

	atk, err := utils.SignAccessToken(existUser, u.accessTokenSecret)
	if err != nil {
		return nil, entity.ErrGlobalServerErr
	}

	claims, rtk, err := utils.SignRefreshToken(existUser, u.refreshTokenSecret)
	if err != nil {
		log.Println(err)
		return nil, entity.ErrGlobalServerErr
	}

	savedRtk := &entity.RefreshToken{
		UserID:    existUser.ID,
		Token:     rtk,
		ExpiresAt: claims.ExpiresAt.Time,
		Jti:       claims.ID,
	}

	if err := u.repo.CreateRefreshToken(ctx, savedRtk); err != nil {
		return nil, err
	}

	return &entity.AuthLoginResp{
		User:         existUser,
		AccessToken:  atk,
		RefreshToken: rtk,
	}, nil
}

func (u *authUsercase) Logout(refreshToken string) error {
	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	cliams, err := utils.ParseRefreshToken(refreshToken, u.refreshTokenSecret)
	if err != nil {
		log.Println(err)
		return err
	}

	updatedRefreshTokon := &entity.RefreshToken{
		Jti:       cliams.ID,
		IsRevoked: true,
	}

	return u.repo.UpdateRefreshToken(ctx, updatedRefreshTokon)
}

func (u *authUsercase) RefreshToken(rtk string) (string, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	// Get refreshToken in db
	oldClaims, err := utils.ParseRefreshToken(rtk, u.refreshTokenSecret)
	if err != nil {
		return "", "", err
	}

	// Get refreshToken in db
	existRtk, err := u.repo.GetRefreshToken(ctx, oldClaims)
	if err != nil {
		return "", "", err
	}

	if existRtk.IsRevoked {
		return "", "", entity.ErrAuthTokenExpired
	}

	if time.Now().After(existRtk.ExpiresAt) {
		return "", "", entity.ErrAuthTokenExpired
	}

	// Sign new accessToken
	newAtk, err := utils.SignAccessToken(&existRtk.User, u.accessTokenSecret)
	if err != nil {
		return "", "", entity.ErrGlobalServerErr
	}

	// Sign new refreshToken
	newClaims, newRtk, err := utils.SignRefreshToken(&existRtk.User, u.refreshTokenSecret)
	if err != nil {
		log.Println(err)
		return "", "", entity.ErrGlobalServerErr
	}

	savedRtk := &entity.RefreshToken{
		UserID:    existRtk.UserID,
		Token:     newRtk,
		ExpiresAt: newClaims.ExpiresAt.Time,
		Jti:       newClaims.ID,
	}

	if err := u.repo.CreateRefreshToken(ctx, savedRtk); err != nil {
		return "", "", err
	}

	// Revok old refreshToken
	existRtk.IsRevoked = true
	if err := u.repo.UpdateRefreshToken(ctx, existRtk); err != nil {
		return "", "", err
	}

	return newAtk, newRtk, nil
}
