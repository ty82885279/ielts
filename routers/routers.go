package routers

import (
	"ielts/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouters() *gin.Engine {

	r := gin.Default()
	group := r.Group("/v1")
	{
		group.GET("/index", controller.Index)
		group.POST("/register", controller.RegisterUser)
		group.POST("/updataUser", controller.UpdateUser)
		group.GET("/login", controller.Login)

		//
		group.POST("/record", controller.CreatRecord)
		group.GET("/allRecord", controller.GetRecords)
		group.POST("/deleteRecord", controller.DeleteARecord)

		//
		group.POST("/uploadAvatar", controller.UploadAvatar)

	}
	return r
}
