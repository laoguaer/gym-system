package handle

import (
	g "gin-blog/internal/global"
	"gin-blog/internal/model"

	"github.com/gin-gonic/gin"
)

// Student 学员模块处理函数
type Student struct{}

// StudentQuery 学员查询参数
type StudentQuery struct {
	Page     int    `form:"page_num"`  // 当前页数（从1开始）
	Size     int    `form:"page_size"` // 每页条数
	Name     string `form:"name"`
	NickName string `form:"nickname"`
}

// GetList 获取学员列表
func (*Student) GetList(c *gin.Context) {
	var query StudentQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		ReturnError(c, g.ErrRequest, err)
		return
	}

	// 调用model层获取学员列表
	list, total, err := model.GetStudentList(GetDB(c), query.Page, query.Size, query.Name, query.NickName)
	if err != nil {
		ReturnError(c, g.ErrDbOp, err)
		return
	}

	ReturnSuccess(c, PageResult[model.StudentVO]{
		Total: total,
		List:  list,
		Size:  query.Size,
		Page:  query.Page,
	})
}
