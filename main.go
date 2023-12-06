package main

import (
	"manageSystem/router"
	"manageSystem/utils"
)

func init() {
	utils.InitViper()
	utils.InitDB()
	router.InitHandler()
	router.InitMiddleware()
}
func main() {
	router.Run()
}
