package controller

import (
	"clean/dto"
	"clean/service"
	"clean/utils"

	// "go/token"
	"net/http"

	"github.com/gin-gonic/gin"
)
type blogController struct{
	blogService service.BlogService
	jwtService service.JWTService
}
type BlogController interface{
	AddingBlog(ctx *gin.Context)
	GetAllBlog(ctx *gin.Context)
	GetMyBlog(ctx *gin.Context)
}
func NewBlogController(bs service.BlogService, jwt service.JWTService) BlogController {
	return &blogController{
		blogService: bs,
		jwtService: jwt,
	}
}
//ama
func (bc *blogController) AddingBlog(ctx *gin.Context){
	var blogDTO dto.BlogCreate
	id := ctx.MustGet("ID").(uint64)
	err := ctx.ShouldBind(&blogDTO)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to upload blog",http.StatusBadRequest,utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}
	blog,err := bc.blogService.AddBlog(ctx,blogDTO,id)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to upload blog",http.StatusBadRequest,utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}
	response := utils.BuildResponse("success to upload blog",http.StatusCreated,blog)
	ctx.JSON(http.StatusCreated,response)
}
func (bc *blogController) GetAllBlog(ctx *gin.Context){
	blogs, err := bc.blogService.GetAllBlog(ctx)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to get the blog",http.StatusBadRequest,utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}
	response := utils.BuildResponse("success to get the blog",http.StatusCreated,blogs)
	ctx.JSON(http.StatusOK,response)
}
func (bc *blogController) GetMyBlog(ctx *gin.Context){
	idAkun := ctx.MustGet("ID").(uint64)
	blog, err := bc.blogService.GetBlogByID(ctx,idAkun)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to take info",http.StatusBadRequest,utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}
	response := utils.BuildResponse("success to take info",http.StatusCreated,blog)
	ctx.JSON(http.StatusOK,response)
}