package router

import (
	"manageSystem/handler"
	"manageSystem/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Run() {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	apiRouter := router.Group("/api/v1")

	userRouter := apiRouter.Group("/user")
	{
		userRouter.Use(middleware.AuthToken())
		userRouter.GET("/getUserList", UserHandler.UserListHandler)
		userRouter.POST("/getUser", UserHandler.UserInfoHandler)
		userRouter.POST("/addUser", UserHandler.AddUserHandler)
		userRouter.POST("/editUser", UserHandler.EditUserHandler)
		userRouter.POST("/deleteUser", UserHandler.DeleteUserHandler)
	}

	videoRouter := apiRouter.Group("/video")
	{
		videoRouter.Use(middleware.AuthToken())
		videoRouter.GET("/getVideoList", VideoHandler.VideoListHandler)
		videoRouter.POST("/getVideo", VideoHandler.VideoInfoHandler)
		videoRouter.POST("/addVideo", VideoHandler.AddVideoHandler)
		videoRouter.POST("/editVideo", VideoHandler.EditVideoHandler)
		videoRouter.POST("/deleteVideo", VideoHandler.DeleteVideoHandler)
	}

	// 登录
	apiRouter.POST("/login", LoginHandler.LoginHandler)
	// 权限
	apiRouter.Use(middleware.AuthToken())
	apiRouter.GET("/permission", PermissionHandler.PermissionHandler)

	// 上传视频
	apiRouter.POST("/upload", handler.UploadFileHandler)

	router.Run(":" + viper.GetString("port"))
}
