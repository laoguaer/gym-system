package handle

import (
	"gin-blog/internal/model"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
