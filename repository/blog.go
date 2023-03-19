package repository

import (
	"clean/entity"
	"context"
	"fmt"
	"gorm.io/gorm"
	"errors"
)
type blogRepository struct {
	db *gorm.DB
}
type BlogRepository interface {
	BuatBlog(ctx context.Context, tx *gorm.DB, blog entity.Blog) (entity.Blog, error)
	GetAllBlog(ctx context.Context) ([]entity.Blog, error)
	GetBlogByID(ctx context.Context, tx *gorm.DB, userID uint64) ([]entity.Blog, error)
}
func NewBlogRepository(dbv *gorm.DB) BlogRepository {
	return &blogRepository{
		db: dbv,
	}
}
func (rb *blogRepository)BuatBlog(ctx context.Context, tx *gorm.DB, blog entity.Blog) (entity.Blog, error){
	var err error
	if tx == nil {
		tx =rb.db.WithContext(ctx).Debug().Create(&blog)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Create(&blog).Error
	}

	if err != nil {
		return entity.Blog{}, err
	}
	return blog, nil
}
func (rb *blogRepository)GetAllBlog(ctx context.Context) ([]entity.Blog, error){
	var err error
	var blogs []entity.Blog
	tx := rb.db.WithContext(ctx).Debug().Preload("Likes").Preload("Comments").Find(&blogs)
	if tx.Error != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []entity.Blog{}, fmt.Errorf("blog not found")
		}
		return []entity.Blog{}, err
	}
	return blogs, nil
}
func (rb *blogRepository) GetBlogByID(ctx context.Context, tx *gorm.DB, userID uint64) ([]entity.Blog, error){
	var blog []entity.Blog
	var err error
	if tx == nil {
		tx = rb.db.WithContext(ctx).Debug().Where("user_id = ?", userID).Preload("Likes").Preload("Comments").Find(&blog)
		err = tx.Error
	} else {
		err = rb.db.WithContext(ctx).Debug().Where("user_id = ?", userID).Preload("Likes").Preload("Comments").Find(&blog).Error
	}
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []entity.Blog{}, fmt.Errorf("user with ID %d not found", userID)
		}
		return []entity.Blog{}, err
	}
	return blog, nil
}