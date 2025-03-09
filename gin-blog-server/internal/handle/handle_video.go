package handle

import (
	g "gin-blog/internal/global"
	// "gin-blog/internal/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Video struct{}

// 视频查询参数
type VideoQuery struct {
	PageQuery
	Title      string `form:"title"`
	CategoryId int    `form:"category_id"`
	Status     int    `form:"status"`
	IsDelete   *bool  `form:"is_delete"`
}

// 视频数据传输对象
type VideoVO struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Desc         string `json:"desc"`
	Cover        string `json:"cover"`
	Url          string `json:"video_url"`
	Status       int    `json:"status"`
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
	ViewCount    int    `json:"view_count"`
	LikeCount    int    `json:"like_count"`
	CommentCount int    `json:"comment_count"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

// @Summary 获取视频列表
// @Description 根据条件查询获取视频列表
// @Tags Video
// @Param title query string false "标题"
// @Param category_id query int false "分类ID"
// @Param status query int false "状态"
// @Param is_delete query bool false "是否删除"
// @Param page_num query int false "页码"
// @Param page_size query int false "每页数量"
// @Accept json
// @Produce json
// @Success 0 {object} Response[PageResult[VideoVO]]
// @Router /front/video/list [get]
func (*Video) GetList(c *gin.Context) {
	var query VideoQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		ReturnError(c, g.ErrRequest, err)
		return
	}

	db := GetDB(c)
	// rdb := GetRDB(c)

	// 这里应该调用model层的GetVideoList方法，但由于模型尚未定义，这里先模拟返回数据
	// 实际开发中，需要先在model包中定义Video模型和相关方法
	// list, total, err := model.GetVideoList(db, query.Page, query.Size, query.Title, query.IsDelete, query.Status, query.CategoryId)
	_ = db // 临时使用db变量，避免未使用导入的错误

	// 模拟数据
	total := int64(10)
	list := make([]VideoVO, 0)
	for i := 1; i <= 5; i++ {
		list = append(list, VideoVO{
			ID:           i,
			Title:        "视频标题" + strconv.Itoa(i),
			Desc:         "视频描述" + strconv.Itoa(i),
			Cover:        "https://example.com/cover" + strconv.Itoa(i) + ".jpg",
			Url:          "https://example.com/video" + strconv.Itoa(i) + ".mp4",
			Status:       1,
			CategoryId:   1,
			CategoryName: "教程",
			ViewCount:    100 + i,
			LikeCount:    50 + i,
			CommentCount: 20 + i,
			CreatedAt:    "2023-01-0" + strconv.Itoa(i) + " 12:00:00",
			UpdatedAt:    "2023-01-0" + strconv.Itoa(i) + " 12:00:00",
		})
	}

	// 实际项目中，这里应该从Redis获取点赞数和观看数
	// likeCountMap := rdb.HGetAll(rctx, g.VIDEO_LIKE_COUNT).Val()
	// viewCountZ := rdb.ZRangeWithScores(rctx, g.VIDEO_VIEW_COUNT, 0, -1).Val()

	ReturnSuccess(c, PageResult[VideoVO]{
		Size:  query.Size,
		Page:  query.Page,
		Total: total,
		List:  list,
	})
}
