package router

import (
	"manageSystem/handler"
	"manageSystem/middleware"
	"manageSystem/repository"
	"manageSystem/service"
	"manageSystem/utils"
)

var (
	LoginHandler      handler.LoginHandler
	UserHandler       handler.UserHandler
	PermissionHandler handler.PermissionHandler
)

func InitHandler() {
	LoginHandler = handler.LoginHandler{
		TokenSrv: &service.TokenService{
			Repo: &repository.TokenRepository{
				DB: utils.DB,
			},
		},
	}

	UserHandler = handler.UserHandler{
		UserSrv: &service.UserService{
			Repo: &repository.UserRepository{
				DB: utils.DB,
			},
		},
	}

	PermissionHandler = handler.PermissionHandler{
		PermissionSrv: &service.PermissionService{
			Repo: &repository.PermissionRepository{
				DB: utils.DB,
			},
		},
	}

}

func InitMiddleware() {
	middleware.TokenSrv = service.TokenService{
		Repo: &repository.TokenRepository{
			DB: utils.DB,
		},
	}
}
