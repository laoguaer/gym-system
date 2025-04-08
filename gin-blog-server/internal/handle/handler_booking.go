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
		if err := model.CancleBooking(GetDB(c), body.BookingID); err != nil {
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
