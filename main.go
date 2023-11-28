package main

import (
	"manageSystem/router"
	"manageSystem/utils"
)

func init() {
	utils.InitViper()
	utils.InitDB()
}
func main() {
	router.Run()
}
