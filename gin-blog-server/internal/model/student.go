package model

import (
	"time"

	"go.uber.org/zap"
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
	defer func() {
		if err := recover(); err != nil {
			zap.L().Error("GetStudentList panic", zap.Any("err", err))
		}
	}()
	if name != "" {
		db = db.Where("username LIKE ?", "%"+name+"%")
	}
	ret := []UserAuth{}
	result := db.Debug().Model(&UserAuth{}).
		Joins("LEFT JOIN user_info ON user_info.id = user_auth.user_info_id")
	if nickname != "" {
		result = result.Where("nickname LIKE?", "%"+nickname+"%")
	}
	err = result.Preload("UserInfo").
		Preload("Roles").
		Count(&total).
		Scopes(Paginate(page, size)).
		Find(&ret).Error
	if err != nil {
		zap.L().Error(err.Error())
		return nil, 0, err
	}
	for _, v := range ret {
		zap.L().Debug("user", zap.Any("user", v))
		if len(v.Roles) != 0 && v.Roles[0].ID == 2 {
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
