package response

import (
	"net/http"
	"reflect"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type behavior interface {
	Code() int
	Message() string
}

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"` // 去掉 omitempty，保证结构一致
}

func HttpJson(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err == nil {
		body.Code = 200
		body.Msg = "OK"
		body.Data = fixNilData(resp)
	} else {
		// 默认脱敏配置
		body.Code = 500
		body.Msg = "服务器内部错误"
		body.Data = struct{}{}

		if b, ok := err.(behavior); ok {
			// 业务主动抛出的错误，安全输出
			body.Code = b.Code()
			body.Msg = b.Message()
		} else {
			// 意外错误（DB、系统、Panic），记录日志并隐藏细节
			logx.WithContext(r.Context()).Errorf("【系统异常】详细错误: %+v", err)
		}
	}
	httpx.WriteJson(w, http.StatusOK, body)
}

func fixNilData(data interface{}) interface{} {
	if data == nil {
		return struct{}{}
	}
	v := reflect.ValueOf(data)
	kind := v.Kind()
	if (kind == reflect.Slice || kind == reflect.Array) && v.IsNil() {
		return make([]interface{}, 0)
	}
	return data
}
