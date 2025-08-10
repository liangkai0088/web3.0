package errorx

type CodeError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewCodeError(code int, message string) *CodeError {
	return &CodeError{Code: code, Message: message}
}

func (e *CodeError) Error() string {
	return e.Message
}

var (
	ErrSystem             = NewCodeError(10001, "系统异常，请稍后再试")
	ErrUserNotFound       = NewCodeError(10002, "用户不存在")
	ErrInvalidCredentials = NewCodeError(10003, "认证失败")
	ErrUnauthorized       = NewCodeError(10004, "权限不足")
	ErrInvalidParams      = NewCodeError(10005, "请求参数错误")
	ErrPostNotFound       = NewCodeError(10006, "文章不存在")
	ErrCommentNotFound    = NewCodeError(10007, "评论不存在")
)
