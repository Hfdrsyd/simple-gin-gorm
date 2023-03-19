package service

import (
	"clean/dto"
	"clean/entity"
	"clean/repository"
	"context"

)

type blogService struct{
	blogRepo repository.BlogRepository
}
type BlogService interface{
	AddBlog(ctx context.Context, blogDTO dto.BlogCreate ,userID uint64) (entity.Blog, error)
	GetAllBlog(ctx context.Context) ([]entity.Blog, error)
	GetBlogByID(ctx context.Context,userID uint64) ([]entity.Blog, error)
}
func NewBlogService(bs repository.BlogRepository) BlogService{
	return &blogService{
		blogRepo: bs,
	}
}
func (bs *blogService) AddBlog(ctx context.Context, blogDTO dto.BlogCreate ,userID uint64) (entity.Blog, error){
	var blog entity.Blog
	blog.ID=blogDTO.ID
	blog.Isi_blog=blogDTO.Isi_blog
	blog.Judul=blogDTO.Judul
	blog.User_ID=userID

	nBlog,err := bs.blogRepo.BuatBlog(ctx,nil,blog)
	if err != nil {
		return entity.Blog{},err
	}
	return nBlog,nil
}
func (bs *blogService) GetAllBlog(ctx context.Context) ([]entity.Blog, error){
	blogs,err := bs.blogRepo.GetAllBlog(ctx)
	if err != nil {
		return []entity.Blog{},err
	}
	return blogs,nil
}
func (bs *blogService) GetBlogByID(ctx context.Context,userID uint64) ([]entity.Blog, error){
	blog, err := bs.blogRepo.GetBlogByID(ctx,nil,userID)
	if err !=nil {
		return []entity.Blog{},err
	}
	return blog,nil
}