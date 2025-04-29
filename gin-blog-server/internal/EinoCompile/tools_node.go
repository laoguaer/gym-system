package EinoCompile

import (
	"context"
	"strconv"

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
	CourseName *string      `json:"courseName" jsonschema:"description=课程名称，没有提到则为nil"`
	IsSingle   *string      `json:"isSingle" jsonschema:"description=是否是私教课，没有提到则为nil，如果是私教课，则为true，如果是团课，则为false"`
	CoachName  *string      `json:"coachName" jsonschema:"description=教练名称，没有提到则为nil，如果是赵教练，则为赵"`
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
	// resp.Status = "success"
	// resp.CouresList = "ai体态课程-教练赵宇锋-结合3D动作捕捉技术实时分析体态数据，自动生成矫正训练方案，含肩颈平衡/骨盆矫正等12种评估维度"
	zap.L().Info("invoke course tool", zap.Any("req", req))
	switch req.Action {
	case GetCourse:
		coachIdList := []int{}
		if req.CoachName != nil && *req.CoachName != "nil" {
			coachIdList, err = model2.GetCoachIdListByName(utils2.DB, *req.CoachName)
			if err != nil {
				resp.Status = "failed"
				return nil, err
			}
		}
		var isSingle *bool
		if req.IsSingle != nil && *req.IsSingle!= "nil" {
			parse, err := strconv.ParseBool(*req.IsSingle)
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
