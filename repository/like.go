package repository

import (
	"clean/entity"
	"context"
	"gorm.io/gorm"
)
type likeRepository struct {
	db *gorm.DB
}
type LikeRepository interface {
	CreateLike(ctx context.Context, tx *gorm.DB, blog entity.Like) (entity.Like, error)
}
func NewLikeRepository(dbv *gorm.DB) LikeRepository {
	return &likeRepository{
		db: dbv,
	}
}
func (rb *likeRepository)CreateLike(ctx context.Context, tx *gorm.DB, like entity.Like) (entity.Like, error){
	var err error
	if tx == nil {
		tx =rb.db.WithContext(ctx).Debug().Create(&like)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Create(&like).Error
	}

	if err != nil {
		return entity.Like{}, err
	}
	return like, nil
}