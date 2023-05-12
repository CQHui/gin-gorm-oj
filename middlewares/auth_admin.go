package middlewares

import (
	"gin-gorm-oj/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthAdminCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		userClaims, err := helper.ParseToken(auth)
		if err != nil {
			ctx.Abort()
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized",
			})
			return
		}

		if userClaims.IsAdmin != 1 {
			ctx.Abort()
			ctx.JSON(http.StatusOK, gin.H{
				"code":    http.StatusForbidden,
				"message": "Unauthorized",
			})
			return
		}
		ctx.Next()
	}
}
