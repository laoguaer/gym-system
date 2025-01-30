package handle

import (
	"bytes"
	"encoding/json"
	"fmt"
	g "gin-blog/internal/global"
	"gin-blog/internal/model"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
curl -X POST 'https://api.coze.cn/v1/workflow/stream_run' \
-H "Authorization: Bearer pat_Ozo3WylcPI18kUqTY7wYNucGrjUtQziKsUy73dALUhRTfu0VZmzLktyabPs132hP" \
-H "Content-Type: application/json" \

	-d '{
	  "parameters": {
	    "start": "默认"
	  },
	  "workflow_id": "7454209939809320995"
	}'

{"code":0,"cost":"0","data":"{\"content_type\":1,\"data\":\"{\\\"output\\\":\\\"{\\\\\\\"回答\\\\\\\":\\\\\\\"这是默认消息\\\\\\\",\\\\\\\"问题\\\\\\\":\\\\\\\"默认\\\\\\\"}\\\"}\",\"original_result\":null,\"type_for_model\":2}","debug_url":"https://www.coze.cn/work_flow?execute_id=7464449491292700682\u0026space_id=7452369603411329059\u0026workflow_id=7454209939809320995","msg":"Success","token":269}
*/

type Chat struct{}
type ChatReq struct {
	ChatString string `json:"chat_string" binding:"required"`
}

type ChatResp struct {
	Code     int    `json:"code"`
	Cost     string `json:"cost"`
	Data     string `json:"data"`
	DebugUrl string `json:"debug_url"`
	Msg      string `json:"msg"`
	Token    int    `json:"token"`
}

type CozeReq struct {
	Parameters map[string]string `json:"parameters" binding:"required"`
	WorkflowId string            `json:"workflow_id" binding:"required"`
}
type CozeResp struct {
	Code     int      `json:"code"`
	Cost     string   `json:"cost"`
	Data     string `json:"data"`
	DebugUrl string   `json:"debug_url"`
	Msg      string   `json:"msg"`
	Token    int      `json:"token"`
}
type CozeData struct {
	ContentType    int    `json:"content_type"`
	Data           string `json:"data"`
	OriginalResult string `json:"original_result"`
	TypeForModel   int    `json:"type_for_model"`
}

func (*Chat) Send(c *gin.Context) {
	var req ChatReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	url := g.GetConfig().Coze.Url
	workflowId := g.GetConfig().Coze.WorkflowId
	configMap, err := model.GetConfigMap(GetDB(c))
	if err != nil {
		ReturnError(c, g.ErrDbOp, err)
		return
	}
	token := configMap["coze_token"]
	cozeReq := CozeReq{
		Parameters: map[string]string{"start": req.ChatString},
		WorkflowId: workflowId,
	}
	requestBody, _ := json.Marshal(cozeReq)
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "请求创建失败"})
		return
	}
	httpReq.Header.Set("Authorization", "Bearer "+token)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "请求发送失败"})
		return
	}
	defer resp.Body.Close()
	log.Println(resp.Body)
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "响应读取失败"})
		return
	}
	var apiResp CozeResp
	if err := json.Unmarshal(bodyBytes, &apiResp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("响应：%v 响应解析失败: %v", string(bodyBytes), err)})
		return
	}
	var cozeData CozeData
	if err := json.Unmarshal([]byte(apiResp.Data), &cozeData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("响应：%v 响应解析失败: %v", string(bodyBytes), err)})
		return
	}
	log.Printf("chat Request: %v\n response: %v", string(requestBody), apiResp)
	c.JSON(http.StatusOK, apiResp) }
