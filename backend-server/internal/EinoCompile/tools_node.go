package EinoCompile

import (
	"context"
	"strconv"
	"time"

	model2 "gin-blog/internal/model"

	utils2 "gin-blog/internal/utils"

	"github.com/cloudwego/eino-ext/components/tool/duckduckgo"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"go.uber.org/zap"
)

func newTool(ctx context.Context) (bt tool.BaseTool, err error) {
	// TODO Modify component configuration here.
	config := &duckduckgo.Config{}
	bt, err = duckduckgo.NewTool(ctx, config)
	if err != nil {
		return nil, err
	}
	return bt, nil
}

// ------------------course tool----------------------
type CourseToolImpl struct {
	config *CourseToolConfig
}
type CourseToolConfig struct {
}
type CourseAction string

const (
	GetCourse CourseAction = "查询"
)

type CourseTaskRequest struct {
	Action     CourseAction `json:"action" jsonschema:"description=需要执行的动作，支持：查询"`
	CourseName string       `json:"courseName" jsonschema:"description=课程名称，没有提到则为空"`
	IsSingle   string       `json:"isSingle" jsonschema:"description=是否是私教课，没有提到则为空，如果是私教课，则为true，如果是团课，则为false"`
	CoachName  string       `json:"coachName" jsonschema:"description=教练名称，没有提到则为空，如果是赵教练，则为赵"`
}
type CourseTaskResponse struct {
	Status     string            `json:"status" jsonschema:"description=status of the response"`
	CouresList []model2.CourseVO `json:"courseList" jsonschema:"description=查询到的课程列表"`
}

func (t *CourseToolImpl) ToEinoTool() (tool tool.BaseTool, err error) {
	return utils.InferTool("coures manager", "课程管理器，主要是查询课程信息", t.Invoke)
}

func (t *CourseToolImpl) Invoke(ctx context.Context, req *CourseTaskRequest) (resp *CourseTaskResponse, err error) {
	resp = &CourseTaskResponse{}
	zap.L().Info("invoke course tool", zap.Any("req", req))
	switch req.Action {
	case GetCourse:
		coachIdList := []int{}
		if req.CoachName != "" {
			coachIdList, err = model2.GetCoachIdListByName(utils2.DB, req.CoachName)
			if err != nil {
				resp.Status = "failed"
				return nil, err
			}
		}
		var isSingle *bool
		if req.IsSingle != "" {
			parse, err := strconv.ParseBool(req.IsSingle)
			if err != nil {
				resp.Status = "failed"
				return nil, err
			}
			isSingle = &parse
		} else {
			isSingle = nil
		}
		resp.CouresList, err = model2.GetCourseListForAi(utils2.DB, req.CourseName, coachIdList, isSingle)
		if err != nil {
			resp.Status = "failed"
			zap.L().Error("invoke course tool, get course list failed", zap.Error(err))
			return nil, err
		}
		resp.Status = "success"
	}
	zap.L().Info("invoke course tool", zap.Any("resp", resp))
	return resp, nil
}

func newCourseTool() (bt tool.BaseTool, err error) {
	config := &CourseToolConfig{}
	t := &CourseToolImpl{
		config: config,
	}
	bt, err = t.ToEinoTool()
	if err != nil {
		return nil, err
	}
	return bt, nil
}

// ------------------course tool----------------------

//###################################################

// ------------------cocah tool----------------------

type CoachToolImpl struct {
}

type CoachToolRequest struct {
	Action    string `json:"action" jsonschema:"description=需要执行的动作，支持：查询"`
	CoachName string `json:"coachName" jsonschema:"description=教练名称，没有提到则为空，如果是赵教练，则为赵"`
}
type CoachToolResponse struct {
	Status  string         `json:"status" jsonschema:"description=status of the response"`
	CoachVO []model2.Coach `json:"coachVO" jsonschema:"description=查询到的教练信息列表"`
}

func (t *CoachToolImpl) ToEinoTool() (tool tool.BaseTool, err error) {
	return utils.InferTool("coach manager", "教练管理器，主要是查询教练信息", t.Invoke)
}
func (t *CoachToolImpl) Invoke(ctx context.Context, req *CoachToolRequest) (resp *CoachToolResponse, err error) {
	resp = &CoachToolResponse{}
	zap.L().Info("invoke coach tool", zap.Any("req", req))
	switch req.Action {
	case "查询":
		resp.CoachVO, err = model2.GetCoachForAi(utils2.DB, req.CoachName)
		if err != nil {
			resp.Status = "failed"
			zap.L().Error("invoke coach tool, get coach list failed", zap.Error(err))
			return nil, err
		}
		resp.Status = "success"
	}
	zap.L().Info("invoke coach tool", zap.Any("resp", resp))
	return resp, nil
}
func newCoachTool() (bt tool.BaseTool, err error) {
	t := &CoachToolImpl{}
	bt, err = t.ToEinoTool()
	if err != nil {
		return nil, err
	}
	return bt, nil
}

// ------------------coach tool----------------------

//###################################################

// ------------------reservation tool----------------------
type ReservationToolImpl struct {
}

type ReservationAction string

const (
	GetReservation    ReservationAction = "查询"
	CreateReservation ReservationAction = "创建"
	UpdateReservation ReservationAction = "更新"
	DeleteReservation ReservationAction = "删除"
)

type ReservationToolRequest struct {
	Action   ReservationAction `json:"action" jsonschema:"description=需要执行的动作，支持：查询、创建、更新、删除"`
	UserID   int               `json:"userId" jsonschema:"description=用户ID，没有提到则为0"`
	CourseID int               `json:"courseId" jsonschema:"description=课程ID，如果提到了课程名，则调用courseTool查询课程ID，没有提到则为0"`
	BuyCnt   int               `json:"buyCnt" jsonschema:"description=购买次数，没有提到则为0"`
	ID       int               `json:"id" jsonschema:"description=预约ID，没有提到则为0"`
}

type ReservationToolResponse struct {
	Status       string               `json:"status" jsonschema:"description=status of the response"`
	Reservations []model2.Reservation `json:"reservations" jsonschema:"description=查询到的预约信息列表"`
}

func (t *ReservationToolImpl) ToEinoTool() (tool tool.BaseTool, err error) {
	return utils.InferTool("reservation manager", "购买管理器，主要是查询、创建、更新、删除用户购买课程的记录", t.Invoke)
}

func (t *ReservationToolImpl) Invoke(ctx context.Context, req *ReservationToolRequest) (resp *ReservationToolResponse, err error) {
	resp = &ReservationToolResponse{}
	zap.L().Info("invoke reservation tool", zap.Any("req", req))
	switch req.Action {
	case GetReservation:
		if req.UserID != 0 && req.CourseID != 0 {
			// 查询特定用户的特定课程预约
			reservation, err := model2.GetReservationByUserIdAndCourseId(utils2.DB, req.UserID, req.CourseID)
			if err != nil {
				resp.Status = "failed"
				zap.L().Error("invoke reservation tool, get reservation failed", zap.Error(err))
				return nil, err
			}
			if reservation != nil {
				resp.Reservations = append(resp.Reservations, *reservation)
			}
		} else if req.UserID != 0 {
			// 查询用户的所有预约
			reservations, err := model2.GetReservationByUserId(utils2.DB, req.UserID)
			if err != nil {
				resp.Status = "failed"
				zap.L().Error("invoke reservation tool, get reservations failed", zap.Error(err))
				return nil, err
			}
			resp.Reservations = reservations
		}
		resp.Status = "success"
	case CreateReservation:
		if req.UserID == 0 || req.CourseID == 0 || req.BuyCnt == 0 {
			resp.Status = "failed"
			return nil, err
		}
		// 检查用户是否已经购买该课程
		reservation, err := model2.GetReservationByUserIdAndCourseId(utils2.DB, req.UserID, req.CourseID)
		if err != nil {
			resp.Status = "failed"
			zap.L().Error("invoke reservation tool, get reservation failed", zap.Error(err))
			return nil, err
		}
		if reservation != nil {
			// 更新购买数量
			reservation.BuyCnt += req.BuyCnt
			err = model2.UpdateReservation(utils2.DB, *reservation)
			if err != nil {
				resp.Status = "failed"
				zap.L().Error("invoke reservation tool, update reservation failed", zap.Error(err))
				return nil, err
			}
		} else {
			// 新增购买记录
			newReservation := model2.Reservation{
				UserID:    req.UserID,
				CourseID:  req.CourseID,
				BuyCnt:    req.BuyCnt,
				UseCnt:    0,
				CreatedAt: time.Now(),
			}
			err = model2.CreateReservation(utils2.DB, newReservation)
			if err != nil {
				resp.Status = "failed"
				zap.L().Error("invoke reservation tool, create reservation failed", zap.Error(err))
				return nil, err
			}
		}
		resp.Status = "success"
	case UpdateReservation:
		if req.UserID == 0 || req.CourseID == 0 || req.BuyCnt == 0 {
			resp.Status = "failed"
			return nil, err
		}
		// 更新预约信息
		reservation, err := model2.GetReservationByUserIdAndCourseId(utils2.DB, req.UserID, req.CourseID)
		if err != nil || reservation == nil {
			resp.Status = "failed"
			zap.L().Error("invoke reservation tool, get reservation failed", zap.Error(err))
			return nil, err
		}
		reservation.BuyCnt = req.BuyCnt
		err = model2.UpdateReservation(utils2.DB, *reservation)
		if err != nil {
			resp.Status = "failed"
			zap.L().Error("invoke reservation tool, update reservation failed", zap.Error(err))
			return nil, err
		}
		resp.Status = "success"
	}
	zap.L().Info("invoke reservation tool", zap.Any("resp", resp))
	return resp, nil
}

func newReservationTool() (bt tool.BaseTool, err error) {
	t := &ReservationToolImpl{}
	bt, err = t.ToEinoTool()
	if err != nil {
		return nil, err
	}
	return bt, nil
}

// ------------------reservation tool----------------------

//###################################################

// ------------------booking tool----------------------
type BookingToolImpl struct {
}

type BookingAction string

const (
	CreateBooking BookingAction = "预约"
	CancelBooking BookingAction = "取消预约"
	GetBooking    BookingAction = "查询"
)

type BookingToolRequest struct {
	Action    BookingAction `json:"action" jsonschema:"description=需要执行的动作，支持：预约、取消预约"`
	UserID    int           `json:"userId" jsonschema:"description=用户ID"`
	CourseID  int           `json:"courseId" jsonschema:"description=课程ID"`
	Status    int           `json:"status" jsonschema:"description=预约状态，预约时不需要提供，查询时没提到则为3,已预约为0,完成的预约则为1, 取消的预约则为2"`
	BookingID int           `json:"bookingId" jsonschema:"description=预约ID，取消预约时需要提供"`
	StartTime string        `json:"startTime" jsonschema:"description=预约开始时间，格式为2006-01-02 15:04:05，预约时需要提供"`
	EndTime   string        `json:"endTime" jsonschema:"description=预约结束时间，格式为2006-01-02 15:04:05，预约时需要提供"`
}

type BookingToolResponse struct {
	Status  string           `json:"status" jsonschema:"description=status of the response"`
	Message string           `json:"message" jsonschema:"description=响应消息"`
	Booking []model2.Booking `json:"booking" jsonschema:"description=预约信息列表"`
}

func (t *BookingToolImpl) ToEinoTool() (tool tool.BaseTool, err error) {
	return utils.InferTool("booking manager", "预约管理器，主要是进行课程预约和取消预约操作", t.Invoke)
}

func (t *BookingToolImpl) Invoke(ctx context.Context, req *BookingToolRequest) (resp *BookingToolResponse, err error) {
	resp = &BookingToolResponse{}
	zap.L().Info("invoke booking tool", zap.Any("req", req))
	switch req.Action {
	case GetBooking:
		if req.UserID == 0 {
			resp.Status = "failed"
			resp.Message = "用户ID不能为空"
			return resp, nil
		}
		if req.CourseID == 0 {
			// 查询用户的所有预约
			bookings, err := model2.GetBookingForAi(utils2.DB, req.UserID, req.Status)
			if err != nil {
				resp.Status = "failed"
				zap.L().Error("invoke booking tool, get bookings failed", zap.Error(err))
				return nil, err
			}
			resp.Booking = bookings
			resp.Status = "success"
			return resp, nil
		} else {
			// 查询用户的特定课程预约
			bookings, err := model2.GetBookingByUserIdAndCourseId(utils2.DB, req.UserID, req.CourseID)
			if err != nil {
				resp.Status = "failed"
				zap.L().Error("invoke booking tool, get bookings failed", zap.Error(err))
				return nil, err
			}
			resp.Booking = bookings
			resp.Status = "success"
			return resp, nil
		}

	case CreateBooking:
		if req.StartTime == "" || req.EndTime == "" {
			resp.Status = "failed"
			resp.Message = "预约时间不能为空"
			return resp, nil
		}

		startTime, err := time.Parse("2006-01-02 15:04:05", req.StartTime)
		if err != nil {
			resp.Status = "failed"
			resp.Message = "预约开始时间格式错误"
			return resp, nil
		}

		endTime, err := time.Parse("2006-01-02 15:04:05", req.EndTime)
		if err != nil {
			resp.Status = "failed"
			resp.Message = "预约结束时间格式错误"
			return resp, nil
		}

		// 查询课程信息
		course, err := model2.GetCourseById(utils2.DB, req.CourseID)
		if err != nil {
			resp.Status = "failed"
			resp.Message = "获取课程信息失败"
			zap.L().Error("invoke booking tool, get course failed", zap.Error(err))
			return resp, nil
		}

		if course == nil {
			resp.Status = "failed"
			resp.Message = "课程不存在"
			return resp, nil
		}
		// 检查用户是否有时间冲突
		hasConflict, err := model2.CheckUserTimeConflict(utils2.DB, req.UserID, startTime, endTime)
		if err != nil {
			resp.Status = "failed"
			resp.Message = "检查时间冲突失败"
			zap.L().Error("invoke booking tool, check time conflict failed", zap.Error(err))
			return resp, nil
		}

		if hasConflict {
			resp.Status = "failed"
			resp.Message = "用户在该时间段已有预约"
			return resp, nil
		}

		// 检查用户是否购买了该课程且有足够的次数
		reservation, err := model2.GetReservationByUserIdAndCourseId(utils2.DB, req.UserID, req.CourseID)
		if err != nil {
			resp.Status = "failed"
			resp.Message = "获取预约信息失败"
			zap.L().Error("invoke booking tool, get reservation failed", zap.Error(err))
			return resp, nil
		}

		if reservation == nil || reservation.BuyCnt <= reservation.UseCnt {
			resp.Status = "failed"
			resp.Message = "未购买该课程或可用次数不足"
			return resp, nil
		}

		// 根据课程类型进行不同的检查
		if course.IsSingle == 1 { // 私教课
			// 检测教练是否有时间冲突
			hasCoachConflict, err := model2.CheckCoachTimeConflict(utils2.DB, course.CoachID, startTime, endTime)
			if err != nil {
				resp.Status = "failed"
				resp.Message = "检查教练时间冲突失败"
				zap.L().Error("invoke booking tool, check coach time conflict failed", zap.Error(err))
				return resp, nil
			}

			if hasCoachConflict {
				resp.Status = "failed"
				resp.Message = "教练在该时间段已有预约"
				return resp, nil
			}
		} else { // 团课
			// 检查是否已达到最大人数
			isFull, remainingSlots, err := model2.CheckCourseCapacity(utils2.DB, req.CourseID, startTime, endTime)
			if err != nil {
				resp.Status = "failed"
				resp.Message = "检查课程容量失败"
				zap.L().Error("invoke booking tool, check course capacity failed", zap.Error(err))
				return resp, nil
			}

			if isFull {
				resp.Status = "failed"
				resp.Message = "该课程已达到最大人数"
				return resp, nil
			}

			zap.L().Debug("课程剩余名额", zap.Int("remainingSlots", remainingSlots))
		}

		// 插入booking表该记录
		booking := &model2.Booking{
			UserID:      req.UserID,
			CourseID:    req.CourseID,
			CoachID:     course.CoachID,
			StartTime:   startTime.Format("2006-01-02 15:04:05"),
			EndTime:     endTime.Format("2006-01-02 15:04:05"),
			Status:      0, // 0表示已预约
			CourseTitle: course.Title,
			CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		}

		err = model2.CreateBooking(utils2.DB, booking)
		if err != nil {
			resp.Status = "failed"
			resp.Message = "创建预约失败"
			zap.L().Error("invoke booking tool, create booking failed", zap.Error(err))
			return resp, nil
		}

		resp.Status = "success"
		resp.Message = "预约成功"
		resp.Booking = []model2.Booking{*booking}

	case CancelBooking:
		if req.BookingID == 0 {
			resp.Status = "failed"
			resp.Message = "预约ID不能为空"
			return resp, nil
		}

		// 获取预约信息
		booking, err := model2.GetBookingById(utils2.DB, req.BookingID)
		if err != nil {
			resp.Status = "failed"
			resp.Message = "获取预约信息失败"
			zap.L().Error("invoke booking tool, get booking failed", zap.Error(err))
			return resp, nil
		}

		if booking == nil {
			resp.Status = "failed"
			resp.Message = "预约不存在"
			return resp, nil
		}

		if booking.UserID != req.UserID {
			resp.Status = "failed"
			resp.Message = "用户ID不匹配"
			return resp, nil
		}

		// 取消预约
		err = model2.CancleBookingAndSubUseCnt(utils2.DB, req.BookingID)
		if err != nil {
			resp.Status = "failed"
			resp.Message = "取消预约失败"
			zap.L().Error("invoke booking tool, cancel booking failed", zap.Error(err))
			return resp, nil
		}

		resp.Status = "success"
		resp.Message = "取消预约成功"
	}

	zap.L().Info("invoke booking tool", zap.Any("resp", resp))
	return resp, nil
}

func newBookingTool() (bt tool.BaseTool, err error) {
	t := &BookingToolImpl{}
	bt, err = t.ToEinoTool()
	if err != nil {
		return nil, err
	}
	return bt, nil
}

// ------------------booking tool----------------------
