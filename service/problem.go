package service

import (
	"gin-gorm-oj/define"
	"gin-gorm-oj/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

// GetProblemBasicList
// @Summary 问题列表
// @Description 获取问题列表
// @Tags 公共方法
// @Param page query  int false "输入当前页，默认第一页"
// @Param size query  int false "输入页数"
// @Param keyword query  string false "keyword"
// @Param category_identity query  string false "category_identity"
// @Success 200 {string} json
// @Router /problem/list [get]
func GetProblemBasicList(c *gin.Context) {
	size, _ := strconv.Atoi(c.DefaultQuery("page", define.DefaultSize))
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("get problem basci list error")
		return
	}
	page = (page - 1) * size
	var count int64

	keyword := c.Query("keyword")
	categoryIdentity := c.Query("category_identity")

	list := make([]*models.ProblemBasic, 0)

	tx := models.GetProblemBasicList(keyword, categoryIdentity)
	err = tx.Count(&count).Omit("content").Offset(page).Limit(size).Find(&list).Error

	if err != nil {
		log.Println("get problem basic list failed")
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"list":  list,
			"count": count,
		},
	})
}

// GetProblemDetail
// @Summary 问题详情
// @Description 获取问题
// @Tags 公共方法
// @Param identity query  string false "problem_identity"
// @Success 200 {string} json
// @Router /problem/detail [get]
func GetProblemDetail(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "问题标识不能为空",
		})
		return
	}
	problemBasic := new(models.ProblemBasic)
	err := models.DB.Where("identity = ?", identity).
		Preload("ProblemCategories").
		Preload("ProblemCategories.CategoryBasic").
		First(&problemBasic).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "问题不存在",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "get problem error" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": problemBasic,
	})
}
