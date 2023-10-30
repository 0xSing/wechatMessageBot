package pkg

import "net/http"

type (
	Code struct {
		Status   bool   `json:"status"`
		Code     int    `json:"code"`
		Message  string `json:"message"`
		HttpCode int    `json:"-"`
	}

	Response struct {
		*Code
		Data interface{} `json:"data,omitempty"`
	}
)

func makeCode(httpCode, code int, status bool, msg string) *Code {
	return &Code{
		Status:   status,
		Code:     code,
		Message:  msg,
		HttpCode: httpCode,
	}
}

func (c *Code) SetMsg(msg string) *Code {
	newCode := new(Code)
	newCode.Status = c.Status
	newCode.Code = c.Code
	newCode.HttpCode = c.HttpCode
	newCode.Message = msg
	return newCode
}

func MakeResp(code *Code, data interface{}) *Response {
	return &Response{
		Code: code,
		Data: data,
	}
}

var (
	Success = makeCode(http.StatusOK, 0, true, "success")
	Created = makeCode(http.StatusCreated, 0, true, "created")

	ParamsError = makeCode(http.StatusBadRequest, 400006, false, "参数不合法")

	NotFoundInternalIp = makeCode(http.StatusUnauthorized, 401003, false, "无法获取客户端IP地址")
	NotInternalIp      = makeCode(http.StatusForbidden, 401004, false, "无法获取客户端IP地址")

	WechatSendTextErr   = makeCode(http.StatusBadRequest, 401200, false, "wechat send err")
	WechatSendImagesErr = makeCode(http.StatusBadRequest, 401201, false, "wechat send err")

	FileIsTooLarge = makeCode(http.StatusRequestEntityTooLarge, 413000, false, "上传文件不能超过100MB")

	InternalError = makeCode(http.StatusInternalServerError, 500001, false, "未知错误, 请稍后再试")
)
