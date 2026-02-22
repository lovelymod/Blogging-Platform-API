package repository

import (
	"blogging-platform-api/internal/entity"
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) entity.AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (repo *authRepository) Register(ctx context.Context, registerUser *entity.User) error {
	var count int64

	if err := repo.db.WithContext(ctx).Model(&entity.User{}).Where(&entity.User{Email: registerUser.Email}).Or(&entity.User{Username: registerUser.Username}).Count(&count).Error; err != nil {
		log.Println(err)
		return entity.ErrGlobalServerErr
	}

	if count != 0 {
		return entity.ErrAuthThisEmailOrUsernameIsAlreadyUsed
	}

	if err := repo.db.WithContext(ctx).Create(registerUser).Error; err != nil {
		log.Println(err)
		return entity.ErrGlobalServerErr
	}

	return nil
}

func (repo *authRepository) Login(ctx context.Context, email string) (*entity.User, error) {
	var existUser entity.User

	if err := repo.db.WithContext(ctx).Where(&entity.User{Email: email}).First(&existUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, entity.ErrAuthWrongEmailOrPassword
		}

		log.Println(err)
		return nil, entity.ErrGlobalServerErr
	}

	return &existUser, nil
}

func (repo *authRepository) RefreshToken(ctx context.Context, claims *jwt.RegisteredClaims) (*entity.User, *entity.RefreshToken, error) {
	var existUser entity.User
	var existRtk entity.RefreshToken

	userID, _ := strconv.ParseUint(claims.Subject, 10, 64)

	if err := repo.db.WithContext(ctx).Where(&entity.RefreshToken{Jti: claims.ID, UserID: uint(userID)}).First(&existRtk).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, entity.ErrGlobalNotFound
		}
		log.Println(err)
		return nil, nil, entity.ErrGlobalServerErr
	}

	if err := repo.db.WithContext(ctx).Where(&entity.User{ID: existRtk.UserID}).First(&existUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, entity.ErrGlobalNotFound
		}
		log.Println(err)
		return nil, nil, entity.ErrGlobalServerErr
	}

	return &existUser, &existRtk, nil
}

func (repo *authRepository) SaveRefreshToken(ctx context.Context, rtk *entity.RefreshToken) error {
	if err := repo.db.WithContext(ctx).Create(rtk).Error; err != nil {
		log.Println(err)
		return entity.ErrGlobalServerErr
	}

	return nil
}

func (repo *authRepository) UpdateRefreshToken(ctx context.Context, rtk *entity.RefreshToken) error {
	if err := repo.db.Where(&entity.RefreshToken{Jti: rtk.Jti}).Updates(rtk).Error; err != nil {
		log.Println(err)
		return entity.ErrGlobalServerErr
	}

	return nil
}
