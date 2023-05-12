package service

import (
	"gin-gorm-oj/help"
	"gin-gorm-oj/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
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

// Login
// @Tags 公共方法
// @Summary 用户登录
// @Param username formData string false "username"
// @Param password formData string false "password"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /login [post]
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "user information is empty",
		})
		return
	}
	password = help.GetMD5(password)
	log.Print("username = {}, password = {}", username, password)

	data := new(models.UserBasic)
	err := models.DB.Where("name = ? and password = ? ", username, password).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "username or password is wrong",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get user failed, error = " + err.Error(),
		})
		return
	}

	token, err := help.GenerateToken(data.Identity, data.Name)

	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "creating token failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"token": token,
		},
	})
}
