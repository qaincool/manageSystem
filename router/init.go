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
	VideoHandler      handler.VideoHandler
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

	VideoHandler = handler.VideoHandler{
		VideoSrv: &service.VideoService{
			Repo: &repository.VideoRepository{
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
