package routes

import (
	"clean/controller"
	"clean/service"
	"clean/middleware"
	"github.com/gin-gonic/gin"
)

func BRouter(route *gin.Engine, BlogController controller.BlogController,LikeController controller.LikeController,komenController controller.CommentController, jwtService service.JWTService){
	routes := route.Group("/blogs")
	{
		routes.POST("/upload",middleware.Authenticate(service.NewJWTService(),"user"),BlogController.AddingBlog)
		routes.GET("/get", BlogController.GetAllBlog)
		routes.GET("/my",middleware.Authenticate(service.NewJWTService(),"user"),BlogController.GetMyBlog)
		routes.POST("/like/:Blog_id", middleware.Authenticate(service.NewJWTService(),"user"), LikeController.Like)
		routes.POST("/comment/:Bid",middleware.Authenticate(service.NewJWTService(),"user"),komenController.Addingcomment)
	}
}