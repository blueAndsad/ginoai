package routers

import (
	"awesomeProject/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1Group := r.Group("v1")
	{
		v1Group.GET("test", controller.Hello)
	}

	return r
}
