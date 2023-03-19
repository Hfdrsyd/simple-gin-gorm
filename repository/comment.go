package repository

import (
	"clean/entity"
	"context"
	"gorm.io/gorm"
)
type commentRepository struct {
	db *gorm.DB
}
type CommentRepository interface {
	CreateComment(ctx context.Context, tx *gorm.DB, komen entity.Comment) (entity.Comment, error)
}
func NewCommentRepository(dbv *gorm.DB) CommentRepository {
	return &commentRepository{
		db: dbv,
	}
}
func (rb *commentRepository)CreateComment(ctx context.Context, tx *gorm.DB, komen entity.Comment) (entity.Comment, error){
	var err error
	if tx == nil {
		tx =rb.db.WithContext(ctx).Debug().Create(&komen)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Create(&komen).Error
	}

	if err != nil {
		return entity.Comment{}, err
	}
	return komen, nil
}