package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"manageSystem/middleware"
)

func Run() {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery(), middleware.AuthLogin(), middleware.AuthLogin())

	apiRouter := router.Group("/api/v1")

	userRouter := apiRouter.Group("/user")
	{
		userRouter.GET("/getUserList", UserHandler.UserListHandler)
		userRouter.POST("/getUser", UserHandler.UserInfoHandler)
		userRouter.POST("/addUser", UserHandler.AddUserHandler)
		userRouter.POST("/editUser", UserHandler.EditUserHandler)
		userRouter.POST("/deleteUser", UserHandler.DeleteUserHandler)
	}

	videoRouter := apiRouter.Group("/videos")
	{
		videoRouter.GET("/getVideos")
	}

	router.Run(":" + viper.GetString("port"))
}
