package dto

// ErrorResponse defines the standard error response structure.
type ErrorResponse struct {
	Error string `json:"error"`
}

// 统一响应格式
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
