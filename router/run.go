package router

import "github.com/gin-gonic/gin"

func Run() {
	router := gin.New()
	router.Use()

	apiRouter := router.Group("/api/v1")

	userRouter := apiRouter.Group("/user")
	{
		userRouter.GET("/getUser", UserHandler.UserInfoHandler)
		userRouter.GET("/getUserList", UserHandler.UserListHandler)
		userRouter.POST("/addUser", UserHandler.AddUserHandler)
		userRouter.POST("/editUser", UserHandler.EditUserHandler)
		userRouter.POST("/deleteUser", UserHandler.DeleteUserHandler)
	}
}
