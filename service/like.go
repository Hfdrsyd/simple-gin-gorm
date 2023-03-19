package service

import (
	"clean/entity"
	"clean/repository"
	"context"
)

type likeService struct{
	likeRepo repository.LikeRepository
}
type LikeService interface{
	Addlike(ctx context.Context, blogID uint64,userID uint64) (entity.Like, error)
	// GetAlllike(ctx context.Context) ([]entity.Like, error)
}
func NewLikeService(ls repository.LikeRepository) LikeService{
	return &likeService{
		likeRepo: ls,
	}
}

func (l *likeService) Addlike(ctx context.Context,blogID uint64,userID uint64) (entity.Like, error){
	var like entity.Like
	like.BlogID = blogID
	like.UserID = userID
	likeN, err := l.likeRepo.CreateLike(ctx,nil,like)
	if err != nil {
		return entity.Like{},err
	}
	return likeN,nil
}