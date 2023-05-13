package router

import (
	_ "gin-gorm-oj/docs"
	"gin-gorm-oj/middlewares"
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

	r.GET("/user/detail", service.GetUserDetail)

	r.GET("/submit/list", service.GetSubmitList)

	r.POST("/login", service.Login)
	r.POST("/email/code", service.SendCode)
	r.POST("/user", service.Register)

	r.GET("/rank/list", service.GetRankList)

	authAdmin := r.Group("/admin", middlewares.AuthAdminCheck())
	authAdmin.POST("/problem", service.GetRankList)

	authAdmin.POST("/category", service.CategoryCreate)
	authAdmin.PUT("/category", service.CategoryModify)
	authAdmin.DELETE("/category", service.CategoryDelete)

	return r
}
