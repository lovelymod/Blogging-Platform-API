package repository

import (
	"blogging-platform-api/internal/entity"
	"context"

	"gorm.io/gorm"
)

type blogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) entity.BlogRepository {
	return &blogRepository{
		db: db,
	}
}

func (repo *blogRepository) Create(ctx context.Context, blog *entity.Blog) error {
	return repo.db.WithContext(ctx).Create(blog).Error
}
