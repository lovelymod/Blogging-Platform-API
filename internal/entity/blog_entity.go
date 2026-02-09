package entity

import (
	"context"

	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
}

type BlogRepository interface {
	Create(ctx context.Context, blog *Blog) error
}

type BlogUsecase interface {
	Create(ctx context.Context, blog *Blog) error
}
