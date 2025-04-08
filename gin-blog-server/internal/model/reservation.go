package model

import (
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	Id        int `gorm:"primaryKey"`
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

func GetReservationByUserIdAndCourseId(db *gorm.DB, userID int, courseID int) (*Reservation, error) {
	var reservation Reservation
	if err := db.Where("user_id =? AND course_id =?", userID, courseID).First(&reservation).Error; err != nil {
		return nil, err
	}
	return &reservation, nil
}

func UpdateReservation(db *gorm.DB, reservation Reservation) error {
	// 检查是否提供了更新条件
	if reservation.Id == 0 {
		return gorm.ErrMissingWhereClause
	}

	// 使用Model和Where来指定更新条件，避免更新零值字段
	if err := db.Model(&Reservation{}).Where("id = ?", reservation.Id).Updates(map[string]interface{}{
		"user_id":    reservation.UserID,
		"course_id":  reservation.CourseID,
		"use_cnt":    reservation.UseCnt,
		"buy_cnt":    reservation.BuyCnt,
		"created_at": reservation.CreatedAt,
	}).Error; err != nil {
		return err
	}
	return nil
}

func CreateReservation(db *gorm.DB, reservation Reservation) error {
	if err := db.Create(&reservation).Error; err != nil {
		return err
	}
	return nil
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
	Id          int `gorm:"primaryKey"`
	UserID      int `gorm:"not null"`
	CourseID    int `gorm:"not null"`
	StartTime   time.Time
	EndTime     time.Time
	Status      int    `gorm:"not null;default:0"`
	CourseTitle string `gorm:"not null"`
	CreatedAt   time.Time
}

func GetBookingByUserId(db *gorm.DB, userID int) ([]Booking, error) {
	var bookings []Booking
	if err := db.Where("user_id =?", userID).Order("start_time").Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}

func GetBookingByUserIdAndTime(db *gorm.DB, userID int, dayTime time.Time) ([]Booking, error) {
	var bookings []Booking
	if err := db.Debug().Where("user_id =? AND DATE(start_time) = DATE(?)", userID, dayTime).Order("start_time").Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}

func GetBookingById(db *gorm.DB, id int) (*Booking, error) {
	var booking Booking
	if err := db.Where("id =?", id).First(&booking).Error; err != nil {
		return nil, err
	}
	return &booking, nil
}

// TODO booking表增加reservation_id字段
func CancleBookingAndSubUseCnt(db *gorm.DB, id int) error {
	if err := db.Transaction(func(tx *gorm.DB) error {
		// 1. 取消预约
		if err := db.Model(&Booking{}).Where("id =?", id).Update("status", 2).Error; err != nil {
			return err
		}
		// 2. 减少useCnt
		if err := db.Model(&Reservation{}).Where("course_id = (SELECT course_id FROM booking WHERE id =?)", id).Update("use_cnt", gorm.Expr("use_cnt - ?", 1)).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
