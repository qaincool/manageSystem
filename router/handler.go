package router

import (
	"manageSystem/handler"
	"manageSystem/repository"
	"manageSystem/service"
	"manageSystem/utils"
)

var (
	UserHandler handler.UserHandler
)

func init() {
	initHandler()
}

func initHandler() {
	UserHandler = handler.UserHandler{
		UserSrv: &service.UserService{
			Repo: &repository.UserRepository{
				DB: utils.DB,
			},
		}}
}
