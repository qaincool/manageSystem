package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"manageSystem/service"
)

var LoginSrv service.LoginService

func AuthLogin() gin.HandlerFunc {
	userPassMap, err := LoginSrv.Auth()
	if err != nil {
		log.Printf("读取数据库中用户账号密码信息失败：%s\n", err)
	}
	fmt.Println("success")
	return gin.BasicAuth(userPassMap)
}
