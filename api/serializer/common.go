package serializer

import (
	"github.com/gin-gonic/gin"
	"time"
)

// 定义错误码, 保证前端可以快速接入
const (
	CodeDatabaseError  = 50001
	CodeParamaterError = 40001
)

// Response 基础序列化器
type Response struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Error     string      `json:"error,omitempty"`
	Timestamp int64       `json:"timestamp"`
}

func (response Response) Result() *Response {
	response.Timestamp = time.Now().Unix()

	return &response
}

// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}

func (response TrackedErrorResponse) Result() *TrackedErrorResponse {
	response.Timestamp = time.Now().Unix()

	return &response
}

// 标准错误
func Error(errorCode int, message string, err error) Response {
	response := Response{
		Code:    errorCode,
		Message: message,
	}

	// 生产环境隐藏底层报错
	if err != nil && gin.Mode() != gin.ReleaseMode {
		response.Error = err.Error()
	}

	return response
}

// 数据库错误
func DatabaseError(message string, err error) Response {
	if message == "" {
		message = "数据库操作失败."
	}

	return Error(CodeDatabaseError, message, err)
}

// 参数错误
func ParamaterError(message string, err error) Response {
	if message == "" {
		message = "输入参数错误."
	}

	return Error(CodeParamaterError, message, err)
}
