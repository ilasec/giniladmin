package errorcode

type ErrorCode int

const (
	ERROR_CODE_SUCCESS   ErrorCode = 200 //成功
	ERROR_CODE_FAILED    ErrorCode = 400 // 失败
	ERROR_CODE_USEREXIST ErrorCode = 502 //用户已存在
	ERROR_CODE_NOTEXIST  ErrorCode = 503
)
