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
	BuyCnt    int `gorm:"not null;default:0"`
}

// GetUserCourseList 根据用户ID获取已购课程列表
func GetReservationByUserId(db *gorm.DB, userID int) ([]Reservation, error) {
	var reservations []Reservation
	if err := db.Where("user_id = ?", userID).Find(&reservations).Error; err != nil {
		return nil, err
	}
	return reservations, nil
}

// CREATE TABLE `booking` (
//     `id` int NOT NULL AUTO_INCREMENT,
//     `user_id` int NOT NULL COMMENT '用户ID',
//     `course_id` int NOT NULL COMMENT '课程ID',
//     `start_time` datetime NOT NULL COMMENT '预约开始时间',
//     `end_time` datetime NOT NULL COMMENT '预约结束时间',
//     `status` int DEFAULT '0' COMMENT '预约状态，0表示已预约，1表示已完成，2表示已取消',
//     `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
//     PRIMARY KEY (`id`),
//     KEY `idx_user_time` (`user_id`, `start_time`),
//     CONSTRAINT `check_time` CHECK ((`start_time` < `end_time`))
// ) ENGINE = InnoDB AUTO_INCREMENT = 2 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci

type Booking struct {
	id        int `gorm:"primaryKey"`
	UserID    int `gorm:"not null"`
	CourseID  int `gorm:"not null"`
	StartTime time.Time
	EndTime   time.Time
	Status    int `gorm:"not null;default:0"`
	CreatedAt time.Time
}

func GetBookingByUserId(db *gorm.DB, userID int) ([]Booking, error) {
	var bookings []Booking
	if err := db.Where("user_id =?", userID).Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}

func GetBookingByUserIdAndTime(db *gorm.DB, userID int, dayTime time.Time) ([]Booking, error) {
	var bookings []Booking
	if err := db.Where("user_id =? AND DATE(start_time) = ?", userID, dayTime).Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}
