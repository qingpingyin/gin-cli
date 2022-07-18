package ginx

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"template/pkg/errorx"
)

type Response struct {
	ctx  *gin.Context
	err  error
	Data interface{}
}

type ResponseData struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func (r *Response) Response() {
	codes := errorx.Cause(r.err)
	r.ctx.JSON(http.StatusOK, ResponseData{
		Code: codes.Code(),
		Msg:  codes.Message(),
		Data: r.Data,
	})
}
