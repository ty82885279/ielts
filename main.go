package main

import (
	"fmt"
	"ielts/dao"
	"ielts/models"
	"ielts/routers"
)

func main() {

	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close()

	dao.DB.AutoMigrate(&models.User{})
	dao.DB.AutoMigrate(&models.Record{})

	r := routers.SetupRouters()
	fmt.Println("启动服务")
	r.Run(":9090")
}
