package serror

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseError struct {
	Message string      // 错误消息
	Status  int         // 响应状态码
	Data    interface{} //data
}

func NewResponseError(status int, msg string) *ResponseError {
	return &ResponseError{
		Status:  status,
		Message: msg,
	}
}

func (resp *ResponseError) Error() error {
	return fmt.Errorf("code :%d\n,msg :%s\n", resp.Status, resp.Message)
}
func (resp *ResponseError) Response(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status:": resp.Status,
		"msg":     resp.Message,
		"data":    resp.Data},
	)
}
