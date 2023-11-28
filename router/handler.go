package router

import (
	"manageSystem/handler"
	"manageSystem/repository"
	"manageSystem/service"
)

var (
	UserHandler handler.UserHandler
)

func initHandler() {
	UserHandler = handler.UserHandler{
		UserSrv: &service.UserService{
			Repo: &repository.UserRepository{
				DB: nil,
			},
		}}
}
