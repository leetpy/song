package errorx

const (
	// 系统级别错误 1xxx
	ErrCodeSystem   = 1000
	ErrCodeParam    = 1001
	ErrCodeDatabase = 1002
	ErrCodeCache    = 1003

	// 业务级别错误 2xxx
	ErrCodeUser          = 2000
	ErrCodeUserNotFound  = 2001
	ErrCodeUserExist     = 2002
	ErrCodePasswordWrong = 2003

	// 认证错误 3xxx
	ErrCodeAuth         = 3000
	ErrCodeTokenInvalid = 3001
	ErrCodeTokenExpired = 3002

	// 业务模块错误 4xxx
	ErrCodeCharacter = 4000
	ErrCodeRelation  = 5000
	ErrCodeMap       = 6000
)

var errMsg = map[int]string{
	ErrCodeSystem:        "系统错误",
	ErrCodeParam:         "参数错误",
	ErrCodeDatabase:      "数据库错误",
	ErrCodeCache:         "缓存错误",
	ErrCodeUser:          "用户错误",
	ErrCodeUserNotFound:  "用户不存在",
	ErrCodeUserExist:     "用户已存在",
	ErrCodePasswordWrong: "密码错误",
	ErrCodeAuth:          "认证失败",
	ErrCodeTokenInvalid:  "Token 无效",
	ErrCodeTokenExpired:  "Token 已过期",
}

func GetErrMsg(code int) string {
	if msg, ok := errMsg[code]; ok {
		return msg
	}
	return "未知错误"
}
