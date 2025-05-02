package handle

import (
	g "gin-blog/internal/global"
	"gin-blog/internal/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Course struct{}

func (*Course) GetList(c *gin.Context) {
	var query model.CourseQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		ReturnError(c, g.ErrRequest, err)
		return
	}

	list, total, err := model.GetCourseList(GetDB(c), query.Page, query.Size, query.Title, query.CoachID, query.TagID, query.IsSingle)
	if err != nil {
		ReturnError(c, g.ErrDbOp, err)
		return
	}

	ReturnSuccess(c, PageResult[model.CourseVO]{
		Total: total,
		List:  list,
		Size:  query.Size,
		Page:  query.Page,
	})
}

type UserCourseVo struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	StartTime   time.Time    `json:"start_time"`
	EndTime     time.Time    `json:"end_time"`
	CoachID     int          `json:"coach_id"`
	CoachName   string       `json:"coach_name"`
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
			CoachName:   course.CoachName,
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

type AddOrUpdateCourseBody struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	StartTime   time.Time `json:"start_time" binding:"required"`
	EndTime     time.Time `json:"end_time" binding:"required"`
	CoachID     int       `json:"coach_id" binding:"required,min=1"`
	CoachName   string    `json:"coach_name" binding:"required"`
	MaxCapacity int       `json:"max_capacity"`
	IsSingle    int       `json:"is_single"`
}

// AddCourse 添加课程
func (*Course) AddOrUpdateCourse(c *gin.Context) {
	var body AddOrUpdateCourseBody
	if err := c.ShouldBindJSON(&body); err != nil {
		ReturnError(c, g.ErrRequest, err)
		return
	}

	// 记录请求信息
	zap.L().Debug("add course body", zap.Any("body", body))

	// 对于私教课，不需要设置最大上课容量
	if body.IsSingle == 1 {
		// 私教课默认容量为1
		body.MaxCapacity = 1
	} else if body.MaxCapacity <= 0 {
		// 团体课必须设置最大上课容量
		ReturnError(c, g.ErrRequest, nil)
		return
	}

	// 创建课程对象
	course := &model.Course{
		Title:       body.Title,
		Description: body.Description,
		Tags:        "",
		StartTime:   body.StartTime,
		EndTime:     body.EndTime,
		CoachID:     body.CoachID,
		CoachName:   body.CoachName,
		MaxCapacity: body.MaxCapacity,
		IsSingle:    body.IsSingle,
	}

	// 保存到数据库
	db := GetDB(c)
	if body.ID != 0 {
		// 更新课程
		course.ID = body.ID
		err := model.UpdateCourse(db, course)
		if err != nil {
			ReturnError(c, g.ErrDbOp, err)
			return
		}
	} else {
		// 创建课程
		err := model.CreateCourse(db, course)
		if err != nil {
			ReturnError(c, g.ErrDbOp, err)
			return
		}
	}

	ReturnSuccess(c, nil)
}

// DeleteCourse 删除课程
// 请求参数：[6]
func (*Course) DeleteCourse(c *gin.Context) {
	bodyBytes, _ := c.GetRawData()
	idStr := string(bodyBytes)
	// 去除首尾的中括号
	idStr = idStr[1 : len(idStr)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		c.JSON(400, gin.H{"error": "无效ID"})
		return
	}
	// 从数据库中删除课程
	db := GetDB(c)
	err = model.DeleteCourse(db, id)
	if err != nil {
		ReturnSuccess(c, nil)
		return
	}
	ReturnSuccess(c, nil)
}
