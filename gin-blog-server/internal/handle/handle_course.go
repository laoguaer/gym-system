package handle

import (
	g "gin-blog/internal/global"
	"gin-blog/internal/model"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
		})
	}
	ReturnSuccess(c, courseList)
}
