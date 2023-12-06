package router

import (
	"manageSystem/handler"
	"manageSystem/middleware"
	"manageSystem/repository"
	"manageSystem/service"
	"manageSystem/utils"
)

var (
	UserHandler handler.UserHandler
)

func InitHandler() {
	UserHandler = handler.UserHandler{
		UserSrv: &service.UserService{
			Repo: &repository.UserRepository{
				DB: utils.DB,
			},
		}}

}

func InitMiddleware() {
	middleware.LoginSrv = service.LoginService{
		Repo: repository.LoginRepository{
			DB: utils.DB,
		},
	}
}
