package handle

import (
	g "gin-blog/internal/global"
	"gin-blog/internal/model"
	"time"

	"github.com/gin-gonic/gin"
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
}

type GetUserCourseListQuery struct {
	UserID int `form:"user_id" binding:"required,min=1"`
}

func(*Course) GetUserCourseList(c *gin.Context) {
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
			UseCnt:      reservation.UseCnt,
			TagList:     course.TagList,
			Coach:       course.Coach,
		})
	}
	ReturnSuccess(c, courseList)
}
