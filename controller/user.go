package controller

import (
	"clean/dto"
	"clean/entity"
	"clean/service"
	"clean/utils"

	// "go/token"
	"net/http"

	"github.com/gin-gonic/gin"
)
	
type userController struct{
	userService service.UserService
	jwtService service.JWTService
}
type UserController interface{
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetMyInfo(ctx *gin.Context)
	UpdateMyName(ctx *gin.Context)
	DeleteMyAkun(ctx *gin.Context)
	GetAllAkun(ctx *gin.Context)
}
func NewUserController(us service.UserService, jwt service.JWTService) UserController{
	return &userController{
		userService: us,
		jwtService: jwt,
	}
}
//aman
func (c *userController) Register(ctx *gin.Context) {
	var userDTO dto.RegisterUser
	//mengambil
	errDTO := ctx.ShouldBind(&userDTO)
	//memasukkan kedalam userDTO
	if errDTO != nil{
		response := utils.BuildErrorResponse("Failed to process request",http.StatusBadRequest,utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}
	//memanggil layer service dengan user Controller, untuk membuat user yang telah dimasukkan kedalam userDTO
	user,err := c.userService.Register(ctx,userDTO)
	
	if err!=nil{
		response :=utils.BuildErrorResponse("Failed to register user",http.StatusBadRequest,utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}

	response := utils.BuildResponse("user created",http.StatusCreated,user)
	ctx.JSON(http.StatusCreated,response)
}

func (c *userController) Login(ctx *gin.Context){
	var userDTO dto.LoginUser
	errDTO := ctx.ShouldBind(&userDTO)

	if errDTO != nil{
		response := utils.BuildErrorResponse("Failed to process request",http.StatusBadRequest,utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}

	user,err := c.userService.Login(ctx,userDTO)
	if err!=nil {
		response := utils.BuildErrorResponse("Failed to login User",http.StatusBadRequest,utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}
	token,e := c.jwtService.GenerateToken(user.ID,user.Role)
	if e != nil {
		response := utils.BuildErrorResponse("Failed to login User",http.StatusBadRequest,utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}
	tok := entity.Authorization{
		Token: token,
		Role: user.Role,
	}
	response := utils.BuildResponse("Login berhasil",http.StatusCreated,tok)
	ctx.JSON(http.StatusCreated,response)
}
func (c *userController) GetMyInfo(ctx *gin.Context){
	id := ctx.MustGet("ID").(uint64)
	user,err := c.userService.GetUserInfo(ctx,id)
	
	if err !=nil {
		response := utils.BuildErrorResponse("Failed To get info",http.StatusBadRequest,utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}
	user.Password=" "
	response := utils.BuildResponse("berhasil mengambil data",http.StatusCreated,user)
	ctx.JSON(http.StatusOK,response)
}
func (c *userController) UpdateMyName(ctx *gin.Context){
	id := ctx.MustGet("ID").(uint64)
	user,err := c.userService.GetUserInfo(ctx,id)
	if err !=nil {
		response := utils.BuildErrorResponse("Failed To Update nama",http.StatusBadRequest,utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}
	var userDTO dto.UpdateName
	errDTO := ctx.ShouldBind(&userDTO)
	if errDTO != nil {
		response := utils.BuildErrorResponse("Gagal merubah nama",http.StatusBadRequest,utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}
	userNew,err := c.userService.UpdateNama(ctx,user.Nama,userDTO.Nama)
	if err !=nil {
		response := utils.BuildErrorResponse("Failed To get info",http.StatusBadRequest,utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}
	userNew.Password=" "
	response := utils.BuildResponse("berhasil mengubah data",http.StatusCreated,userNew)
	ctx.JSON(http.StatusOK,response)
}
func (c *userController) DeleteMyAkun(ctx *gin.Context){
	id := ctx.MustGet("ID").(uint64)
	user,err := c.userService.DeleteAkun(ctx,id)
	if err !=nil {
		response := utils.BuildErrorResponse("gagal menghapus akun",http.StatusBadRequest,utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}
	user.Password=" "
	response := utils.BuildResponse("berhasil menghapus akun",http.StatusCreated,user)
	ctx.JSON(http.StatusOK,response)
}
func (c *userController) GetAllAkun(ctx *gin.Context){
	users, err := c.userService.GetAllUser(ctx)
	if err != nil{
		response := utils.BuildErrorResponse("Failed To get info",http.StatusBadRequest,utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}
	response := utils.BuildResponse("succesfully get info",http.StatusCreated,users)
	ctx.JSON(http.StatusOK,response)
}