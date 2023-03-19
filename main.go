package main

import (
	"clean/config"
	"clean/controller"
	"clean/repository"
	routes "clean/route"
	"clean/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}

	db := config.SetupDatabaseConnection()
	jwtService := service.NewJWTService()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService,jwtService)

	blogRepo := repository.NewBlogRepository(db)
	blogservice := service.NewBlogService(blogRepo)
	blogController := controller.NewBlogController(blogservice,jwtService)

	likeRepo := repository.NewLikeRepository(db)
	likeService := service.NewLikeService(likeRepo)
	likeController := controller.NewLikeController(likeService,jwtService)

	commentRepo := repository.NewCommentRepository(db)
	commentService := service.NewCommentService(commentRepo)
	commentController := controller.NewCommentController(commentService,jwtService)
	
	defer config.CloseDatabaseConnection(db)
	port := os.Getenv("PORT")
	server := gin.Default()
	routes.Router(server, userController, jwtService)
	routes.BRouter(server,blogController,likeController,commentController,jwtService)
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)
}
