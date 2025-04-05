package model

import (
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	id        int `gorm:"primaryKey"`
	UserID    int `gorm:"not null"`
	CourseID  int `gorm:"not null"`
	CreatedAt time.Time
	UseCnt    int `gorm:"not null;default:0"`
}

// GetUserCourseList 根据用户ID获取已购课程列表
func GetReservationByUserId(db *gorm.DB, userID int) ([]Reservation, error) {
	var reservations []Reservation
	if err := db.Where("user_id = ?", userID).Find(&reservations).Error; err != nil {
		return nil, err
	}
	return reservations, nil
}
