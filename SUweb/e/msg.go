package e

var MsgFlags = map[int]string {
	SUCCESS : "ok",
	ERROR : "fail",
	INVALID_PARAMS : "请求参数错误",
	ERROR_EXIST_NAME : "已存在该用户名",
	ERROR_NOT_EXIST_USER : "该用户不存在",
	ERROR_NOT_LOGIN : "未登录",
	ERROR_NOT_EXIST_PROJECT: "不存在该项目",
	ERROR_PASSWORD: "密码错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}