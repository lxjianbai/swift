package xerr

const (
	OK                  = 200
	SERVER_COMMON_ERROR = 500
	REUQEST_PARAM_ERROR = 400
	USER_NOT_FOUND      = 40001
)

func MapErrMsg(errCode int) string {
	switch errCode {
	case OK:
		return "SUCCESS"
	case SERVER_COMMON_ERROR:
		return "服务器开小差了，请稍后再试"
	case REUQEST_PARAM_ERROR:
		return "参数错误"
	case USER_NOT_FOUND:
		return "用户不存在"
	default:
		return "服务器未知错误"
	}
}
