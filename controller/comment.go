package controller

import (
	"clean/dto"
	"clean/service"
	"clean/utils"
	"strconv"
	"net/http"

	"github.com/gin-gonic/gin"
)
type commentController struct{
	commentService service.CommentService
	jwtService service.JWTService
}
type CommentController interface{
	Addingcomment(ctx *gin.Context)
}
func NewCommentController(cc service.CommentService, jwt service.JWTService) CommentController {
	return &commentController{
		commentService: cc,
		jwtService: jwt,
	}
}
func (cc *commentController) Addingcomment(ctx *gin.Context) {
	idAkun := ctx.MustGet("ID").(uint64)
	idBlog,err := strconv.ParseUint(ctx.Param("Bid"), 10, 64)
	if err != nil {
		response := utils.BuildErrorResponse("failed to do komen",http.StatusBadRequest,utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}
	var komenDTO dto.Comment
	er := ctx.ShouldBind(&komenDTO)
	if er != nil {
		response := utils.BuildErrorResponse("failed to do comment",http.StatusBadRequest,utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}
	komen ,e := cc.commentService.AddComment(ctx,komenDTO,idBlog,idAkun)
	if e != nil {
		response := utils.BuildErrorResponse("failed to do comment",http.StatusBadRequest,utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}
	response := utils.BuildResponse("Comment succesfully",http.StatusCreated,komen)
	ctx.JSON(http.StatusOK,response)
}