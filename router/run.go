package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"manageSystem/middleware"
)

func Run() {
	router := gin.New()
	router.Use(gin.Recovery(), middleware.AuthLogin(), gin.BasicAuth())

	apiRouter := router.Group("/api/v1")

	userRouter := apiRouter.Group("/user")
	{
		userRouter.GET("/getUser", UserHandler.UserInfoHandler)
		userRouter.GET("/getUserList", UserHandler.UserListHandler)
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
