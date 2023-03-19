package service

import (
	"clean/dto"
	"clean/entity"
	"clean/repository"
	"context"
)

type commentService struct{
	commentRepo repository.CommentRepository
}
type CommentService interface{
	AddComment(ctx context.Context, userDTO dto.Comment, blogID uint64,userID uint64) (entity.Comment, error)
	// GetAlllike(ctx context.Context) ([]entity.Like, error)
}
func NewCommentService(cr repository.CommentRepository) CommentService{
	return &commentService{
		commentRepo: cr,
	}
}
func (c *commentService) AddComment(ctx context.Context, komenDTO dto.Comment ,blogID uint64, userID uint64) (entity.Comment, error){
	var komen entity.Comment
	komen.ID=komenDTO.ID
	komen.BlogID=blogID
	komen.Text=komenDTO.Text
	komen.UserID = userID

	komen ,err := c.commentRepo.CreateComment(ctx,nil,komen)
	if err != nil {
		return entity.Comment{},err
	}
	return komen,nil
}