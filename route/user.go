package routes

import (
	"clean/controller"
	"clean/service"
	"clean/middleware"
	"github.com/gin-gonic/gin"
)

func Router(route *gin.Engine, UserController controller.UserController, jwtService service.JWTService){
	routes := route.Group("/users")
	{
		routes.POST("/register", UserController.Register)
		routes.POST("/login", UserController.Login)
		routes.GET("/my", middleware.Authenticate(service.NewJWTService(),"user"), UserController.GetMyInfo)
		routes.GET("/get/admin", middleware.Authenticate(service.NewJWTService(),"admin"), UserController.GetAllAkun)
		routes.PUT("/my/update", middleware.Authenticate(service.NewJWTService(),"user"), UserController.UpdateMyName)
		routes.DELETE("/my/delete", middleware.Authenticate(service.NewJWTService(),"user"), UserController.DeleteMyAkun)
	}
}