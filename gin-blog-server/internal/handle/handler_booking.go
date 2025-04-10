package handle

import (
	g "gin-blog/internal/global"
	"gin-blog/internal/model"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Booking struct{}

type GetUserBookingWithDayQuery struct {
	UserID int
	Day    string
}

func (*Booking) GetUserBookingWithDay(c *gin.Context) {
	var query GetUserBookingWithDayQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dayTime, err := time.Parse("2006-01-02", query.Day)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	zap.L().Debug("dayTime", zap.Any("dayTime", dayTime))
	bookings, err := model.GetBookingByUserIdAndTime(GetDB(c), query.UserID, dayTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ReturnSuccess(c, bookings)
}

type CancleBookingBody struct {
	BookingID int `json:"booking_id" binding:"required,min=1"`
	UserID    int `json:"user_id" binding:"required,min=1"`
}

func (*Booking) CancleBooking(c *gin.Context) {
	var body CancleBookingBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := GetDB(c).Transaction(func(tx *gorm.DB) error {
		booking, err := model.GetBookingById(GetDB(c), body.BookingID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return err
		}
		if booking == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "booking not found"})
			return err
		}
		if booking.UserID != body.UserID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user id not match"})
			return err
		}
		// 取消预约
		if err := model.CancleBookingAndSubUseCnt(GetDB(c), body.BookingID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return err
		}
		// 减少useCnt

		return nil
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ReturnError(c, g.ErrDbOp, err)
		return
	}

	ReturnSuccess(c, nil)
}

type GetAvailableBookingTimeQuery struct {
	Day      string `form:"date" binding:"required"`
	CourseID int    `form:"course_id" binding:"required"`
	UserID   int    `form:"user_id" binding:"required,min=1"`
}

func (*Booking) GetAvailableBookingTime(c *gin.Context) {
	var query GetAvailableBookingTimeQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dayTime, err := time.Parse("2006-01-02", query.Day)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	course, err := model.GetCourseById(GetDB(c), query.CourseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	zap.L().Debug("course ", zap.Any("course", course))
	if course == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "course not found"})
		return
	}

	// 计算可用时间
	availableTime := []string{}

	// 根据课程类型返回不同的可用时间
	if course.IsSingle == 1 { // 私教课
		// 对于私教课，返回10-22点中教练和用户都没有预约的时间段
		for i := 10; i < 22; i++ {
			startTime := time.Date(dayTime.Year(), dayTime.Month(), dayTime.Day(), i, 0, 0, 0, time.Local)
			endTime := time.Date(dayTime.Year(), dayTime.Month(), dayTime.Day(), i+1, 0, 0, 0, time.Local)

			// 检查用户在该时间段是否有冲突
			userConflict, err := model.CheckUserTimeConflict(GetDB(c), query.UserID, startTime, endTime)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			// 检查教练在该时间段是否有冲突
			coachConflict, err := model.CheckCoachTimeConflict(GetDB(c), course.CoachID, startTime, endTime)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			// 如果用户和教练都没有冲突，则添加到可用时间
			if !userConflict && !coachConflict {
				availableTime = append(availableTime, startTime.Format("2006-01-02 15:04:05"))
			}
		}
	} else { // 团体课
		hour := course.StartTime.Hour()

		// 对于团体课，只返回start_time，并且如果该时间段用户有其它预约，则不返回
		startTime := time.Date(dayTime.Year(), dayTime.Month(), dayTime.Day(), hour, 0, 0, 0, time.Local)
		endTime := time.Date(dayTime.Year(), dayTime.Month(), dayTime.Day(), hour+1, 0, 0, 0, time.Local)

		// 检查用户在该时间段是否有冲突
		userConflict, err := model.CheckUserTimeConflict(GetDB(c), query.UserID, startTime, endTime)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 如果用户没有冲突，则添加到可用时间
		if !userConflict {
			availableTime = append(availableTime, startTime.Format("2006-01-02 15:04:05"))
		}
	}

	ReturnSuccess(c, availableTime)
}

type BookingBody struct {
	UserID    int    `json:"user_id" binding:"required,min=1"`
	CourseID  int    `json:"course_id" binding:"required,min=1"`
	StartTime string `json:"start_time" binding:"required"`
	EndTime   string `json:"end_time" binding:"required"`
}

func (*Booking) Booking(c *gin.Context) {
	var body BookingBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	startTime, err := time.Parse("2006-01-02 15:04:05", body.StartTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	endTime, err := time.Parse("2006-01-02 15:04:05", body.EndTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := GetDB(c).Transaction(func(tx *gorm.DB) error {
		// step1. 检测该用户是否有时间冲突
		hasConflict, err := model.CheckUserTimeConflict(tx, body.UserID, startTime, endTime)
		if err != nil {
			return err
		}
		if hasConflict {
			c.JSON(http.StatusBadRequest, gin.H{"error": "用户在该时间段已有预约"})
			return err
		}

		// step2. 查询课程信息
		course, err := model.GetCourseById(tx, body.CourseID)
		if err != nil {
			return err
		}
		if course == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "课程不存在"})
			return err
		}

		// 检查用户是否购买了该课程且有足够的次数
		reservation, err := model.GetReservationByUserIdAndCourseId(tx, body.UserID, body.CourseID)
		if err != nil {
			return err
		}
		if reservation == nil || reservation.BuyCnt <= reservation.UseCnt {
			c.JSON(http.StatusBadRequest, gin.H{"error": "未购买该课程或可用次数不足"})
			return err
		}

		// 根据课程类型进行不同的检查
		if course.IsSingle == 1 { // 私教课
			// 检测教练是否有时间冲突
			hasCoachConflict, err := model.CheckCoachTimeConflict(tx, course.CoachID, startTime, endTime)
			if err != nil {
				return err
			}
			if hasCoachConflict {
				c.JSON(http.StatusBadRequest, gin.H{"error": "教练在该时间段已有预约"})
				return err
			}
		} else { // 团课
			// 检查是否已达到最大人数
			isFull, remainingSlots, err := model.CheckCourseCapacity(tx, body.CourseID, startTime, endTime)
			if err != nil {
				return err
			}
			if isFull {
				c.JSON(http.StatusBadRequest, gin.H{"error": "该课程已达到最大人数"})
				return err
			}
			zap.L().Debug("课程剩余名额", zap.Int("remainingSlots", remainingSlots))
		}

		// step3. 插入booking表该记录
		booking := &model.Booking{
			UserID:      body.UserID,
			CourseID:    body.CourseID,
			CoachID:     course.CoachID,
			StartTime:   startTime.Format("2006-01-02 15:04:05"),
			EndTime:     endTime.Format("2006-01-02 15:04:05"),
			Status:      0, // 0表示已预约
			CourseTitle: course.Title,
			CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		}
		zap.L().Debug("booking ", zap.Any("startTime", startTime.String()))
		if err := model.CreateBooking(tx, booking); err != nil {
			return err
		}

		return nil
	}); err != nil {
		ReturnError(c, g.ErrDbOp, err)
		return
	}
	ReturnSuccess(c, nil)
}
