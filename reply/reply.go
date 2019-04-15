package reply

// reply codes
const (
	CodeSuccess uint = 1
	CodeFailure uint = 9999
)

// Reply basic API reply struct
type Reply struct {
	Code    uint        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// New Reply
func New(code uint, msg string) *Reply {
	return &Reply{Code: code, Message: msg}
}

// OK create success reply
func OK() *Reply {
	return New(CodeSuccess, "ok")
}

// Data create success reply with data
func Data(data interface{}) *Reply {
	return &Reply{Code: CodeSuccess, Message: "ok", Data: data}
}

// Error create failure reply base on error
func Error(err error) *Reply {
	return New(CodeFailure, err.Error())
}
