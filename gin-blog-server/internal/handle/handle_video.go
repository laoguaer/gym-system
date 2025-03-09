package handle

import (
	g "gin-blog/internal/global"
	"gin-blog/internal/model"

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
	model.Video

	LikeCount    int `json:"like_count" gorm:"-"`
	ViewCount    int `json:"view_count" gorm:"-"`
	CommentCount int `json:"comment_count" gorm:"-"`
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
	rdb := GetRDB(c)
	_ = rdb

	// 调用model层的GetVideoList方法获取视频列表
	list, total, err := model.GetVideoList(db, query.Page, query.Size, query.Title, query.IsDelete, query.Status, query.CategoryId)
	if err != nil {
		ReturnError(c, g.ErrDbOp, err)
		return
	}

	// 将Video模型转换为VideoVO
	data := make([]VideoVO, 0)
	for _, video := range list {
		data = append(data, VideoVO{
			Video:        video,
			LikeCount:    0, // 实际项目中从Redis获取
			ViewCount:    0, // 实际项目中从Redis获取
			CommentCount: 0, // 实际项目中从评论表获取
		})
	}

	// 实际项目中，这里应该从Redis获取点赞数和观看数
	// likeCountMap := rdb.HGetAll(rctx, g.VIDEO_LIKE_COUNT).Val()
	// viewCountZ := rdb.ZRangeWithScores(rctx, g.VIDEO_VIEW_COUNT, 0, -1).Val()

	ReturnSuccess(c, PageResult[VideoVO]{
		Size:  query.Size,
		Page:  query.Page,
		Total: total,
		List:  data,
	})
}

func (*Video) GetDetail(c *gin.Context) {
	_, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ReturnError(c, g.ErrRequest, err)
		return
	}

	// article, err := model.GetArticle(GetDB(c), id)
	detail := VideoVO{
		// ID:           id,
		// Title:        "视频标题" + strconv.Itoa(id),
		// Desc:         "视频描述" + strconv.Itoa(id),
		// Cover:        "https://example.com/cover" + strconv.Itoa(id) + ".jpg",
		// Url:          "https://example.com/video" + strconv.Itoa(id) + ".mp4",
		// Status:       1,
		// CategoryId:   1,
		// CategoryName: "教程",
		// ViewCount:    100 + id,
		// LikeCount:    50 + id,
		// CommentCount: 20 + id,
		// CreatedAt:    "2023-01-0" + strconv.Itoa(id) + " 12:00:00",
		// UpdatedAt:    "2023-01-0" + strconv.Itoa(id) + " 12:00:00",
	}
	if err != nil {
		ReturnError(c, g.ErrDbOp, err)
		return
	}

	ReturnSuccess(c, detail)
}
