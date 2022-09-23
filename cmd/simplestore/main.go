package main

import (
	"graphdemo/api/router"
	"graphdemo/pkg/util"
)

func main() {
	util.InitDB()
	router.ServeRouter()
}
