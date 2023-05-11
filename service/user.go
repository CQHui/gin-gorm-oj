package service

import (
	"gin-gorm-oj/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUserDetail
// @Tags 公共方法
// @Summary 用户详情
// @Param identity query string false "problem identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user/detail [get]
func GetUserDetail(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "问题标识不能为空",
		})
		return
	}

	userBasic := new(models.UserBasic)
	err := models.DB.Omit("password").Where("identity = ?", identity).Find(&userBasic).Error

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get User Detail By Identity:" + identity + " Error:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": userBasic,
	})
}
