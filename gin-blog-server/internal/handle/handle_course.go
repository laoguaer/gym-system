package handle

import (
	g "gin-blog/internal/global"
	"gin-blog/internal/model"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Course struct{}
type MyCourseVo struct {
	Courses []model.CourseVO `json:"courses"`
}
type UserCourseVo struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	StartTime   time.Time    `json:"start_time"`
	EndTime     time.Time    `json:"end_time"`
	CoachID     int          `json:"coach_id"`
	IsSingle    int          `json:"is_single"`
	Coach       *model.Coach `json:"coach,omitempty"`
	TagList     []string     `json:"tag_list,omitempty"`
	UseCnt      int          `json:"use_cnt"`
	BuyCnt      int          `json:"buy_cnt"`
	MaxCapacity int          `json:"max_capacity"`
}

type GetUserCourseListQuery struct {
	UserID int `form:"user_id" binding:"required,min=1"`
}

func (*Course) GetUserCourseList(c *gin.Context) {
	var query GetUserCourseListQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		ReturnError(c, g.ErrRequest, err)
		return
	}
	reservations, err := model.GetReservationByUserId(GetDB(c), query.UserID)
	if err != nil {
		ReturnError(c, g.ErrDbOp, err)
		return
	}
	courseList := make([]UserCourseVo, 0)
	for _, reservation := range reservations {
		// 记录请求和响应信息
		zap.L().Debug("reservation", zap.Any("reservation", reservation))
		course, err := model.GetCourseById(GetDB(c), reservation.CourseID)
		if err != nil {
			ReturnError(c, g.ErrDbOp, err)
			return
		}
		courseList = append(courseList, UserCourseVo{
			ID:          course.ID,
			Title:       course.Title,
			Description: course.Description,
			StartTime:   course.StartTime,
			EndTime:     course.EndTime,
			CoachID:     course.CoachID,
			IsSingle:    course.IsSingle,
			TagList:     course.TagList,
			Coach:       course.Coach,
			UseCnt:      reservation.UseCnt,
			BuyCnt:      reservation.BuyCnt,
			MaxCapacity: course.MaxCapacity,
		})
	}
	ReturnSuccess(c, courseList)
}

type BuyCourseBody struct {
	UserID   int `json:"user_id" binding:"required,min=1"`
	CourseID int `json:"course_id" binding:"required,min=1"`
	BuyCnt   int `json:"count" binding:"required,min=1"`
}

func (*Course) BuyCourse(c *gin.Context) {
	var body BuyCourseBody
	if err := c.ShouldBindJSON(&body); err != nil {
		ReturnError(c, g.ErrRequest, err)
		return
	}
	// 记录请求和响应信息
	zap.L().Debug("body", zap.Any("body", body))
	db := GetDB(c)
	// 事务
	if err := db.Transaction(func(tx *gorm.DB) error {
		// 检查用户是否已经购买该课程
		reservation, err := model.GetReservationByUserIdAndCourseId(tx, body.UserID, body.CourseID)
		if err != nil {
			zap.L().Error("get reservation by user id and course id", zap.Error(err))
			return err
		}
		if reservation != nil {
			zap.L().Debug("用户存在reservation", zap.Any("reservation", reservation))
			// 更新购买数量
			reservation.BuyCnt += body.BuyCnt
			err = model.UpdateReservation(tx, *reservation)
			if err != nil {
				zap.L().Error("update reservation", zap.Error(err))
				return err
			}
		} else {
			// 新增购买记录
			reservation = &model.Reservation{
				UserID:    body.UserID,
				CourseID:  body.CourseID,
				BuyCnt:    body.BuyCnt,
				UseCnt:    0,
				CreatedAt: time.Now(),
			}
			err = model.CreateReservation(tx, *reservation)
			if err != nil {
				zap.L().Error("create reservation", zap.Error(err))
				return err
			}
		}
		return nil
	}); err != nil {
		ReturnError(c, g.ErrDbOp, err)
		return
	}
	ReturnSuccess(c, nil)
}
