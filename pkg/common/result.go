package common

type BusinessCode int

type Result struct {
	Code BusinessCode `json:"code"`
	Msg  string       `json:"msg"`
	Data any          `json:"data"`
}

func (r *Result) Success(data any) *Result {
	r.Code = 0
	r.Msg = "success"
	r.Data = data
	return r
}

func (r *Result) Error(code BusinessCode, msg string) *Result {
	r.Code = code
	r.Msg = msg
	return r
}

// 定义枚举异常码
const (
	// 通用异常码
	Success BusinessCode = 0
	// 参数异常
	ParamError BusinessCode = 101
	// 业务异常
	BusinessError BusinessCode = 102
	// 服务异常
	ServiceError BusinessCode = 103
	// 系统异常
	SystemError BusinessCode = 104
)