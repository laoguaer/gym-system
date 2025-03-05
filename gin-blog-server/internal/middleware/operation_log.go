package middleware

import (
	"bytes"
	g "gin-blog/internal/global"
	"gin-blog/internal/handle"
	"gin-blog/internal/model"
	"gin-blog/internal/utils"
	"io"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// TODO: 优化 API 路径格式
var optMap = map[string]string{
	"Article":      "文章",
	"BlogInfo":     "博客信息",
	"Category":     "分类",
	"Comment":      "评论",
	"FriendLink":   "友链",
	"Menu":         "菜单",
	"Message":      "留言",
	"OperationLog": "操作日志",
	"Resource":     "资源权限",
	"Role":         "角色",
	"Tag":          "标签",
	"User":         "用户",
	"Page":         "页面",
	// "Login":        "登录",

	"POST":   "新增或修改",
	"PUT":    "修改",
	"DELETE": "删除",
}

func GetOptString(key string) string {
	return optMap[key]
}

// 记录操作日志中间件
func OperationLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: 记录文件上传
		// 不记录 GET 请求操作记录 (太多了) 和 文件上传操作记录 (请求体太长)
		if c.Request.Method != "GET" && !strings.Contains(c.Request.RequestURI, "upload") {
			blw := &CustomResponseWriter{
				body:           bytes.NewBufferString(""),
				ResponseWriter: c.Writer,
			}
			c.Writer = blw

			auth, _ := handle.CurrentUserAuth(c)

			body, _ := io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

			ipAddress := utils.IP.GetIpAddress(c)
			ipSource := utils.IP.GetIpSource(ipAddress)

			moduleName := getOptResource(c.HandlerName())
			operationLog := model.OperationLog{
				OptModule:     moduleName, // TODO: 优化
				OptType:       GetOptString(c.Request.Method),
				OptUrl:        c.Request.RequestURI,
				OptMethod:     c.HandlerName(),
				OptDesc:       GetOptString(c.Request.Method) + moduleName, // TODO: 优化
				RequestParam:  string(body),
				RequestMethod: c.Request.Method,
				UserId:        auth.UserInfoId,
				Nickname:      auth.UserInfo.Nickname,
				IpAddress:     ipAddress,
				IpSource:      ipSource,
			}
			c.Next()
			operationLog.ResponseData = blw.body.String() // 从缓存中获取响应体内容

			db := c.MustGet(g.CTX_DB).(*gorm.DB)
			if err := db.Create(&operationLog).Error; err != nil {
				zap.L().Error("操作日志记录失败", zap.Error(err))
				handle.ReturnError(c, g.ErrDbOp, err)
				return
			}
		} else {
			c.Next()
		}
	}
}

// "gin-blog/api/v1.(*Resource).Delete-fm" => "Resource"
func getOptResource(handlerName string) string {
	s := strings.Split(handlerName, ".")[1]
	return s[2 : len(s)-1]
}
