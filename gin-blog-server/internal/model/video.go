package model

import (
	"gorm.io/gorm"
)

const (
	VIDEO_STATUS_PUBLIC = iota + 1 // 公开
	VIDEO_STATUS_SECRET            // 私密
	VIDEO_STATUS_DRAFT             // 草稿
)

// belongTo: 一个视频 属于 一个分类
// belongTo: 一个视频 属于 一个用户
type Video struct {
	Model

	Title    string `gorm:"type:varchar(100);not null" json:"title"`
	Desc     string `json:"desc"`
	Cover    string `json:"cover"`
	Url      string `gorm:"type:varchar(255)" json:"video_url"`
	Status   int    `gorm:"type:tinyint;comment:状态(1-公开 2-私密 3-草稿)" json:"status"` // 1-公开 2-私密 3-草稿
	IsDelete bool   `json:"is_delete"`

	CategoryId int `json:"category_id"`
	UserId     int `json:"-"` // user_auth_id

	Category *Category `gorm:"foreignkey:CategoryId" json:"category"`
	User     *UserAuth `gorm:"foreignkey:UserId" json:"user"`
}

// 获取视频列表
func GetVideoList(db *gorm.DB, page, size int, title string, isDelete *bool, status, categoryId int) (data []Video, total int64, err error) {
	db = db.Model(Video{})

	if title != "" {
		db = db.Where("title LIKE ?", "%"+title+"%")
	}
	if isDelete != nil {
		db = db.Where("is_delete = ?", *isDelete)
	}
	if status != 0 {
		db = db.Where("status = ?", status)
	}
	if categoryId != 0 {
		db = db.Where("category_id = ?", categoryId)
	}

	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	db = db.Preload("Category")
	result := db.Offset((page - 1) * size).Limit(size).Find(&data)
	return data, total, result.Error
}

// 获取视频详情
func GetVideo(db *gorm.DB, id int) (data *Video, err error) {
	result := db.Preload("Category").
		Where(Video{Model: Model{ID: id}}).
		First(&data)
	return data, result.Error
}
