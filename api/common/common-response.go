package common

// 响应体
type CommonResponse struct {
	ErrCode int         `json:"status"` // 请求状态
	Msg     string      `json:"msg"`    // 消息内容
	Data    interface{} `json:"data"`   // 数据信息
}

const (
	CODE_SUCCESS  = 1
	UPLOAD_FAILED = 2
)
