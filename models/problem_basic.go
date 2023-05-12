package models

import (
	"gorm.io/gorm"
)

type ProblemBasic struct {
	gorm.Model
	Identity          string             `gorm:"column:identity;type:varchar(36);" json:"identity"`                  // 问题表的唯一标识
	ProblemCategories []*ProblemCategory `gorm:"foreignKey:problem_id;references:id" json:"problem_categories"`      // 关联问题分类表
	Title             string             `gorm:"column:title;type:varchar(255);" json:"title"`                       // 文章标题
	Content           string             `gorm:"column:content;type:text;" json:"content"`                           // 文章正文
	MaxRuntime        int                `gorm:"column:max_runtime;type:int(11);" json:"max_runtime"`                // 最大运行时长
	MaxMem            int                `gorm:"column:max_mem;type:int(11);" json:"max_mem"`                        // 最大运行内存
	TestCases         []*TestCase        `gorm:"foreignKey:problem_identity;references:identity;" json:"test_cases"` // 关联测试用例表

}

func (table *ProblemBasic) TableName() string {
	return "problem_basic"
}

func GetProblemBasicList(keyword string, categoryIdentity string) *gorm.DB {

	tx := DB.Model(new(ProblemBasic)).Preload("ProblemCategories").Preload("CategoryBasic").
		Where("title like ? OR content like ? ", "%"+keyword+"%", "%"+keyword+"%")
	if categoryIdentity != "" {
		tx.Joins("RIGHT JOIN problem_category pc on pc.problem_id = problem_basic.id").
			Where("pc.category_id = (select cb.id from category_basic cb where cb.identity = ?)")
	}
	return tx
}
