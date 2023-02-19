package utils

const(
	SUCCESS int32 = 0
	FAIL int32= 1
	USER_SUCCESS_REGISTER int32 = 2100	//USER 类 2xxx
	USER_SUCCESS_LOGIN int32 = 2101
	USER_FAIL_REGISTER int32 = 2000
	USER_FAIL_LOGIN int32 = 2001
	USER_NOT_EXIT int32 = 2002
	USER_PASSWORD_IS_NOT_CORRECT int32 = 2003

	ERROR_TOKEN_EXIST     int32 = 1004  // TOKEN 类  1xxx
	ERROR_TOKEN_RUNTIME    int32 = 1005
	ERROR_TOKEN_WRONG     int32 = 1006
	ERROR_TOKEN_TYPE_WRONG int32 = 1007

	VIDEO_PUSH_SUCCESS int32 = 3000  //VIDEO 类 3xxx
	VIDEO_PUSH_FAIL int32 = 3001
	VIDEO_GET_SUCCESS int32 = 3002
)

var StatusMsg = map[int32]string{
	SUCCESS: "成功",
	FAIL: "失败",
	USER_SUCCESS_REGISTER: "用户注册成功",
	USER_SUCCESS_LOGIN: "用户登录成功",
	USER_FAIL_REGISTER: "用户注册失败",
	USER_FAIL_LOGIN: "用户登录失败",
	USER_NOT_EXIT: "用户不存在",
	USER_PASSWORD_IS_NOT_CORRECT:"密码输入错误，请重新输入密码",
	ERROR_TOKEN_EXIST:      "TOKEN不存在,请重新登陆",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期,请重新登陆",
	ERROR_TOKEN_WRONG:      "TOKEN不正确,请重新登陆",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误,请重新登陆",
	VIDEO_PUSH_SUCCESS:     "视频上传成功",
	VIDEO_PUSH_FAIL:        "视频上传失败",
	VIDEO_GET_SUCCESS:		"视频获取成功",
}

func GetStatusMsg(code int32) string {
	return StatusMsg[code]
}