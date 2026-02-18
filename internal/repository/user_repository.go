package repository

import (
	"blogging-platform-api/internal/entity"
	"context"
	"log"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) entity.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (repo *userRepository) Register(ctx context.Context, req *entity.UserRegisterReq) error {
	var count int64

	if err := repo.db.WithContext(ctx).Model(&entity.User{}).Where(&entity.User{Email: req.Email}).Count(&count).Error; err != nil {
		log.Println(err)
		return entity.ErrGlobalServerErr
	}

	if count != 0 {
		return entity.ErrUserThisEmailIsAlreadyUsed
	}

	newUser := &entity.User{
		Email:          req.Email,
		HashedPassword: req.HashedPassword,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
	}

	if err := repo.db.WithContext(ctx).Create(newUser).Error; err != nil {
		log.Println(err)
		return entity.ErrGlobalServerErr
	}

	return nil
}
