package exceptions

import "errors"

// InternalServerError 内部服务错误
var (
	ErrInternalServer = errors.New("内部服务错误")
)
