package handle

import (
	"fmt"
	g "gin-blog/internal/global"
	"gin-blog/internal/model"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
)

// Coach 教练模块处理函数
type Coach struct{}

// CoachVO 教练前端展示数据结构
type CoachVO struct {
	ID         int     `json:"id"`
	Avatar     string  `json:"avatar"`
	Name       string  `json:"name"`
	Specialty  string  `json:"specialty"`
	Rating     float64 `json:"rating"`
	Experience string  `json:"experience"`
	Intro      string  `json:"intro"`
	Students   int     `json:"students"`
	Courses    int     `json:"courses"`
}

// CoachQuery 教练查询参数
type CoachQuery struct {
	Page int    `form:"page" binding:"required,min=1"`
	Size int    `form:"size" binding:"required,min=1,max=50"`
	Name string `form:"name"`
}

// CoachVO 教练前端展示数据结构
type AdminCoachVO struct {
	ID       int       `json:"id"`
	Name     string    `json:"name" `
	Phone    string    `json:"phone" `
	Desc     string    `json:"desc" `
	Avatar   string    `json:"avatar"`
	JoinTime time.Time `json:"join_time" `
}

func (*Coach) GetList(c *gin.Context) {
	var query CoachQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		ReturnError(c, g.ErrRequest, err)
		return
	}

	coachs, total, err := model.GetAdminCoachList(GetDB(c), query.Page, query.Size, query.Name)
	if err != nil {
		ReturnError(c, g.ErrDbOp, err)
		return
	}
	list := make([]AdminCoachVO, 0, len(coachs))
	for _, coach := range coachs {
		list = append(list, AdminCoachVO{
			ID:       coach.ID,
			Name:     coach.Name,
			Phone:    coach.Phone,
			Desc:     coach.Desc,
			Avatar:   coach.Avatar,
			JoinTime: coach.CreatedAt,
		})
	}

	ReturnSuccess(c, PageResult[AdminCoachVO]{
		Total: total,
		List:  list,
		Size:  query.Size,
		Page:  query.Page,
	})
}

// GetCoachList 获取前台教练列表
// @Summary 获取前台教练列表
// @Description 获取前台教练列表
// @Tags 前台-教练模块
// @Accept application/json
// @Produce application/json
// @Param page_num query int true "页码" minimum(1)
// @Param page_size query int true "每页数量" minimum(1) maximum(50)
// @Param name query string false "教练姓名"
// @Param occupation query string false "职业/专长"
// @Success 200 {object} Response{data=PageResult{list=[]CoachVO}}
// @Router /api/front/coach/list [get]

func (*Front) GetCoachList(c *gin.Context) {
	var query CoachQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		ReturnError(c, g.ErrRequest, err)
		return
	}

	list, total, err := model.GetCoachList(GetDB(c), query.Page, query.Size, query.Name)
	if err != nil {
		ReturnError(c, g.ErrDbOp, err)
		return
	}

	// 转换为前端需要的数据结构
	coachVOList := make([]CoachVO, 0, len(list))
	for _, coach := range list {
		// 计算教练经验（根据创建时间）
		experience := "新手"
		years := time.Now().Year() - coach.CreatedAt.Year()
		if years > 0 {
			experience = fmt.Sprintf("%d年", years)
		}

		// 创建CoachVO对象
		coachVO := CoachVO{
			ID:         coach.ID,
			Avatar:     coach.Avatar,
			Name:       coach.Name,
			Specialty:  coach.Occupation,
			Rating:     5.0, // 默认评分5.0
			Experience: experience,
			Intro:      coach.Desc,
			Students:   128, // 模拟数据，实际项目中应该查询reservation表
			Courses:    15,  // 模拟数据，实际项目中应该查询course表
		}
		coachVOList = append(coachVOList, coachVO)
	}

	ReturnSuccess(c, PageResult[CoachVO]{
		Total: total,
		List:  coachVOList,
		Size:  query.Size,
		Page:  query.Page,
	})
}

type GetCoachWithIdsReq struct {
	UserId []int `json:"user_id"`
}

func (*Coach) GetCoachWithUserIds(c *gin.Context) {
	var req GetCoachWithIdsReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ReturnError(c, g.ErrRequest, err)
		return
	}
	// userId去重 先排序 再去重
	sort.Slice(req.UserId, func(i, j int) bool {
		return req.UserId[i] < req.UserId[j]
	})
	size := 0
	for i := 0; i < len(req.UserId); i++ {
		if i > 0 && req.UserId[i] == req.UserId[i-1] {
			continue
		}
		req.UserId[size] = req.UserId[i]
		size++
	}
	req.UserId = req.UserId[:size]

	list, err := model.GetCoachWithUserIds(GetDB(c), req.UserId)
	if err != nil {
		ReturnError(c, g.ErrDbOp, err)
		return
	}
	// 转换为前端需要的数据结构
	coachVOList := make([]CoachVO, 0, len(list))
	for _, coach := range list {
		// 计算教练经验（根据创建时间）
		experience := "新手"
		years := time.Now().Year() - coach.CreatedAt.Year()
		if years > 0 {
			experience = fmt.Sprintf("%d年", years)
		}

		// 创建CoachVO对象
		coachVO := CoachVO{
			ID:         coach.ID,
			Avatar:     coach.Avatar,
			Name:       coach.Name,
			Specialty:  coach.Occupation,
			Rating:     5.0, // 默认评分5.0
			Experience: experience,
			Intro:      coach.Desc,
			Students:   128, // 模拟数据，实际项目中应该查询reservation表
			Courses:    15,  // 模拟数据，实际项目中应该查询course表
		}
		coachVOList = append(coachVOList, coachVO)
	}

	ReturnSuccess(c, coachVOList)
}
