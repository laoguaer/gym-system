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

func GetAdminCoachList(db *gorm.DB, page, size int, name string) (AdminCoachVO []Coach, total int64, err error) {
	return GetCoachList(db, page, size, name)
}

// GetCoachList 获取教练列表
func GetCoachList(db *gorm.DB, page, size int, name string) (list []Coach, total int64, err error) {
	if name != "" {
		db = db.Where("name LIKE ?", "%"+name+"%")
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

func GetCoachIdListByName(db *gorm.DB, name string) ([]int, error) {
	if name == "" {
		return []int{}, nil
	}
	var res []int
	err := db.Model(&Coach{}).Select("id").Where("name LIKE?", "%"+name+"%").Find(&res).Error
	return res, err
}
