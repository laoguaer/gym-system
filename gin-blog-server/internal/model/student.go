package model

import (
	"time"

	"gorm.io/gorm"
)

type StudentVO struct {
	ID        int       `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `gorm:"unique;type:varchar(50)" json:"username"`
	Email     string    `json:"email" gorm:"type:varchar(30)"`
	Nickname  string    `json:"nickname" gorm:"unique;type:varchar(30);not null"`
	Avatar    string    `json:"avatar" gorm:"type:varchar(1024);not null"`
	Intro     string    `json:"intro" gorm:"type:varchar(255)"`
}

func GetStudentList(db *gorm.DB, page, size int, name, nickname string) (list []StudentVO, total int64, err error) {
	if name != "" {
		db = db.Where("username LIKE ?", "%"+name+"%")
	}
	ret := []UserAuth{}
	result := db.Model(&UserAuth{}).
		Joins("LEFT JOIN user_info ON user_info.id = user_auth.user_info_id").
		Where("user_info.nickname LIKE ?", "%"+nickname+"%").
		Preload("UserInfo").
		Preload("Roles").
		Count(&total).
		Scopes(Paginate(page, size)).
		Find(&ret)
	for _, v := range ret {
		if v.Roles[0].Name == "普通用户" {
			student := StudentVO{
				ID:        v.UserInfo.ID,
				CreatedAt: v.UserInfo.CreatedAt,
				UpdatedAt: v.UserInfo.UpdatedAt,
				Username:  v.Username,
				Email:     v.UserInfo.Email,
				Nickname:  v.UserInfo.Nickname,
				Avatar:    v.UserInfo.Avatar,
				Intro:     v.UserInfo.Intro,
			}
			list = append(list, student)
		}
	}
	return list, total, result.Error
}
