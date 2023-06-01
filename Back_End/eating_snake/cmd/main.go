package main

import (
	"eating_snake/api"
	"eating_snake/dao"
	"fmt"
)

//cd到cmd后go run main.go启动
func main() {
	fmt.Println("main start")

	dao.InitDB()
	if dao.DB != nil {
		defer dao.DB.Close()
	}

	router := api.InitRouter()
	fmt.Println("router init")

	router.Run("127.0.0.1:8000")
}
