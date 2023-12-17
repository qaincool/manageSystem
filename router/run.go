package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"manageSystem/middleware"
)

func Run() {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	apiRouter := router.Group("/api/v1")

	// 登录
	apiRouter.POST("/login", LoginHandler.LoginHandler)

	userRouter := apiRouter.Group("/user")
	{
		userRouter.Use(middleware.AuthToken())
		userRouter.GET("/getUserList", UserHandler.UserListHandler)
		userRouter.POST("/getUser", UserHandler.UserInfoHandler)
		userRouter.POST("/addUser", UserHandler.AddUserHandler)
		userRouter.POST("/editUser", UserHandler.EditUserHandler)
		userRouter.POST("/deleteUser", UserHandler.DeleteUserHandler)
	}

	permissionRouter := apiRouter
	permissionRouter.Use(middleware.AuthToken())
	permissionRouter.GET("/permission", PermissionHandler.PermissionHandler)

	videoRouter := apiRouter.Group("/videos")
	{
		videoRouter.Use(middleware.AuthToken())
		videoRouter.GET("/getVideos")
	}

	router.Run(":" + viper.GetString("port"))
}
