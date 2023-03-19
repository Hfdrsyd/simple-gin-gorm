package controller

import (
	"clean/service"
	"clean/utils"
	"strconv"

	// "go/token"
	"net/http"

	"github.com/gin-gonic/gin"
)
	
type likeController struct{
	likeService service.LikeService
	jwtService service.JWTService
}
type LikeController interface{
	Like(ctx *gin.Context)
}
func NewLikeController(ls service.LikeService, jwt service.JWTService) LikeController{
	return &likeController{
		likeService: ls,
		jwtService: jwt,
	}
}
func (l *likeController)Like(ctx *gin.Context){
	idAkun := ctx.MustGet("ID").(uint64)
	idBlog,err := strconv.ParseUint(ctx.Param("Blog_id"), 10, 64)
	if err != nil {
		response := utils.BuildErrorResponse("failed to do like",http.StatusBadRequest,utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}
	like ,er := l.likeService.Addlike(ctx,idBlog,idAkun)
	if er != nil {
		response := utils.BuildErrorResponse("failed to do like",http.StatusBadRequest,utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}
	response := utils.BuildResponse("like berhasil dilakukan",http.StatusCreated,like)
	ctx.JSON(http.StatusOK,response)
}