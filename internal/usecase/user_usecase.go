package usecase

import (
	"blogging-platform-api/internal/entity"
	"context"
	"log"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	repo    entity.UserRepository
	timeout time.Duration
	env     *entity.Config
}

func NewUserUsecase(repo entity.UserRepository, contextTimeout time.Duration, env *entity.Config) entity.UserUsecase {
	return &userUsecase{
		repo:    repo,
		timeout: contextTimeout,
		env:     env,
	}
}

func (u *userUsecase) Register(req *entity.UserRegisterReq) error {
	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	parsedCost, err := strconv.Atoi(u.env.HASH_COST)
	if err != nil {
		log.Println(err)
		return entity.ErrGlobalServerErr
	}

	hasdPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), parsedCost)
	if err != nil {
		log.Println(err)
		return entity.ErrGlobalServerErr
	}

	req.HashedPassword = string(hasdPassword)

	return u.repo.Register(ctx, req)
}
