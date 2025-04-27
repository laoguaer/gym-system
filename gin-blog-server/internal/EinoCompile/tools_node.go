package EinoCompile

import (
	"context"
	"gin-blog/internal/model"

	utils2 "gin-blog/internal/utils"

	"github.com/cloudwego/eino-ext/components/tool/duckduckgo"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
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
	GetCourse CourseAction = "getCourse"
)

type CourseTaskRequest struct {
	Action            CourseAction `json:"action" jsonschema:"description=action to perform, enum:getCourse"`
	CourseName        string       `json:"courseName" jsonschema:"description=课程名称"`
	CourseDescription string       `json:"courseDescription" jsonschema:"description=模糊信息检索"`
}
type CourseTaskResponse struct {
	Status     string           `json:"status" jsonschema:"description=status of the response"`
	CouresList []model.CourseVO `json:"courseList" jsonschema:"description=查询到的课程列表"`
}

func (t *CourseToolImpl) ToEinoTool() (tool tool.BaseTool, err error) {
	return utils.InferTool("coures manager", "课程管理器，主要是查询课程信息", t.Invoke)
}

func (t *CourseToolImpl) Invoke(ctx context.Context, req *CourseTaskRequest) (resp *CourseTaskResponse, err error) {
	resp = &CourseTaskResponse{}
	switch req.Action {
	case GetCourse:
		resp.CouresList, _, err = model.GetCourseList(utils2.DB, 0, 100, req.CourseName, 0, 0, nil)
		if err != nil {
			resp.Status = "failed"
			return nil, err
		}
		resp.Status = "success"
	}
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
