package router

import (
	_ "gin-gorm-oj/docs"
	"gin-gorm-oj/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/ping", service.Ping)

	r.GET("/problem/list", service.GetProblemBasicList)
	r.GET("/problem/detail", service.GetProblemDetail)

	return r
}
