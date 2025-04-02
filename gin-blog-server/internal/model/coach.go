package model

import (
	"time"

	"gorm.io/gorm"
)

// Coach 教练信息模型
type Coach struct {
	ID         int            `json:"id" gorm:"primaryKey"`
	Name       string         `json:"name" gorm:"type:varchar(50);not null"`
	Phone      string         `json:"phone" gorm:"type:varchar(15);not null;uniqueIndex"`
	Desc       string         `json:"desc" gorm:"type:text;not null"`
	Avatar     string         `json:"avatar" gorm:"type:varchar(255);not null"`
	Occupation string         `json:"occupation" gorm:"type:varchar(1024);not null"`
	CreatedAt  time.Time      `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP;ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// CoachQuery 教练查询参数
type CoachQuery struct {
	Page       int    `form:"page" binding:"required,min=1"`
	Size       int    `form:"size" binding:"required,min=1,max=50"`
	Name       string `form:"name"`
	Occupation string `form:"occupation"`
}

// GetCoachList 获取教练列表
func GetCoachList(db *gorm.DB, page, size int, name, occupation string) (list []Coach, total int64, err error) {
	if name != "" {
		db = db.Where("name LIKE ?", "%"+name+"%")
	}

	if occupation != "" {
		db = db.Where("occupation LIKE ?", "%"+occupation+"%")
	}

	// 获取总数
	err = db.Model(&Coach{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	err = db.Model(&Coach{}).Limit(size).Offset((page - 1) * size).Find(&list).Error
	return list, total, err
}

// GetCoachById 根据ID获取教练信息
func GetCoachById(db *gorm.DB, id int) (*Coach, error) {
	var coach Coach
	result := db.Model(&coach).Where("id = ?", id).First(&coach)
	return &coach, result.Error
}

func GetCoachWithUserIds(db *gorm.DB, ids []int) ([]Coach, error) {
	var coach []Coach
	err := db.Where("id IN ?", ids).Find(&coach).Error
	return coach, err
}
