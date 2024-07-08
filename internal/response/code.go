package response

import "fmt"

type Code int // 状态码

var (
	SuccessCode       Code = 0
	InputInvalidError Code = 9020
	ParamError        Code = 9021
	JsonMarshalError  Code = 9022
	CreateTimerErr    Code = 9023
	EnableTimerError  Code = 9024
)

var codeMsgDict = map[Code]string{
	SuccessCode:       "ok",
	InputInvalidError: "input invalid",
	ParamError:        "param failed",
	JsonMarshalError:  "json marshal failed",
	CreateTimerErr:    "create timer failed",
	EnableTimerError:  "enable timer failed",
}

func (c *Code) Message() string {
	if msg, ok := codeMsgDict[*c]; ok {
		return msg
	}
	return fmt.Sprintf("unknown error code %d", *c)
}

type ErrorCode struct {
	Code Code `json:"code"`
}

func (e *ErrorCode) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Code.Message())
}
