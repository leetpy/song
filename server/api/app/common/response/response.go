package response

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 统一响应结构体
type Body struct {
	Code int    `json:"code"`           // 状态码
	Msg  string `json:"msg"`            // 消息
	Data any    `json:"data,omitempty"` // 数据
}

func Success(w http.ResponseWriter, data any) {
	httpx.WriteJson(w, http.StatusOK, &Body{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}

// 失败响应
func Error(w http.ResponseWriter, code int, msg string) {
	httpx.WriteJson(w, http.StatusOK, &Body{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

// 带数据的失败响应
func ErrorWithData(w http.ResponseWriter, code int, msg string, data any) {
	httpx.WriteJson(w, http.StatusOK, &Body{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// 分页响应
type PageData struct {
	List     any   `json:"list"`
	Total    int64 `json:"total"`
	Page     int   `json:"page"`
	PageSize int   `json:"pageSize"`
}

func SuccessWithPage(w http.ResponseWriter, list any, total int64, page, pageSize int) {
	Success(w, &PageData{
		List:     list,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	})
}
