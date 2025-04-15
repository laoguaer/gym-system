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
	CoachID     int `gorm:"not null"`
	StartTime   string
	EndTime     string
	Status      int    `gorm:"not null;default:0"`
	CourseTitle string `gorm:"not null"`
	CreatedAt   string
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

// CheckUserTimeConflict 检测用户在指定时间段是否有时间冲突
func CheckUserTimeConflict(db *gorm.DB, userID int, startTime, endTime time.Time) (bool, error) {
	var count int64
	// 使用SQL查询逻辑检测时间冲突
	// 如果存在满足条件的记录，则表示有冲突
	err := db.Model(&Booking{}).Where(
		"user_id = ? AND start_time < ? AND end_time > ? AND status = 0",
		userID, endTime, startTime,
	).Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// CheckCoachTimeConflict 检测教练在指定时间段是否有时间冲突（针对私教课）
func CheckCoachTimeConflict(db *gorm.DB, coachID int, startTime, endTime time.Time) (bool, error) {
	var count int64
	// 使用SQL查询逻辑检测时间冲突
	// 如果存在满足条件的记录，则表示有冲突
	err := db.Model(&Booking{}).Where(
		"coach_id = ? AND start_time < ? AND end_time > ? AND status = 0",
		coachID, endTime, startTime,
	).Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// CheckCourseCapacity 检查团课是否已达到最大人数
func CheckCourseCapacity(db *gorm.DB, courseID int, startTime, endTime time.Time) (bool, int, error) {
	// 1. 获取课程信息，查询最大容量
	var course Course
	if err := db.Model(&Course{}).Where("id = ?", courseID).First(&course).Error; err != nil {
		return false, 0, err
	}

	// 2. 查询当前已预约人数
	var count int64
	if err := db.Model(&Booking{}).Where(
		"course_id = ? AND start_time = ? AND end_time = ? AND status = 0",
		courseID, startTime, endTime,
	).Count(&count).Error; err != nil {
		return false, 0, err
	}

	// 3. 判断是否已达到最大容量
	return int(count) >= course.MaxCapacity, course.MaxCapacity - int(count), nil
}

// UpdateBookingStatus 更新预约状态
func UpdateBookingStatus(db *gorm.DB, id int, status int) error {
	// 检查是否提供了更新条件
	if id == 0 {
		return gorm.ErrMissingWhereClause
	}

	// 使用Model和Where来指定更新条件
	if err := db.Model(&Booking{}).Where("id = ?", id).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}

// CreateBooking 创建预约记录
func CreateBooking(db *gorm.DB, booking *Booking) error {
	// 创建预约记录
	if err := db.Debug().Create(booking).Error; err != nil {
		return err
	}
	// 更新用户课程使用次数
	if err := db.Model(&Reservation{}).Where(
		"user_id = ? AND course_id = ?",
		booking.UserID, booking.CourseID,
	).Update("use_cnt", gorm.Expr("use_cnt + ?", 1)).Error; err != nil {
		return err
	}

	return nil
}
