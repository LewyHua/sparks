package utils

const (
	CodeSuccess int32 = 200 + iota
	CodeServerBusy
	CodeInvalidParam
	CodeInvalidToken
	CodeDBError

	CodeFollowMyself

	CodeInvalidFileType
	CodeInvalidFileSize
	CodeUploadFileError

	CodeNotFriend
	CodeFollowRepeat
	CodeCancelFollowRepeat
	CodeFavoriteRepeat

	CodeWrongLoginCredentials
	CodeUsernameNotFound
	CodeUserNotFound
	CodeInvalidRegisterUsername
	CodeInvalidRegisterPassword
	CodeUsernameAlreadyExists

	CodeInvalidCommentAction
	CodeLimiterCount
)

var message map[int32]string

func init() {
	message = make(map[int32]string)
	message[CodeSuccess] = "success"
	message[CodeServerBusy] = "服务器开小差啦,稍后再来试一试"
	message[CodeInvalidParam] = "参数错误"
	message[CodeInvalidToken] = "请退出重新登陆"
	message[CodeDBError] = "数据库繁忙,请稍后再试"
	message[CodeFollowMyself] = "不能关注自己哦"
	message[CodeFollowRepeat] = "请勿重复关注"
	message[CodeCancelFollowRepeat] = "请勿重复取关"
	message[CodeFavoriteRepeat] = "请勿重复点赞/取消点赞"
	message[CodeInvalidFileType] = "无效的文件类型"
	message[CodeInvalidFileSize] = "文件过大或过小"
	message[CodeUploadFileError] = "文件上传失败"
	message[CodeNotFriend] = "对方并不是您的好友"
	message[CodeWrongLoginCredentials] = "用户名或密码错误"
	message[CodeUsernameNotFound] = "用户名不存在"
	message[CodeUserNotFound] = "用户不存在"
	message[CodeInvalidRegisterUsername] = "用户名不合规"
	message[CodeInvalidRegisterPassword] = "密码不合规"
	message[CodeUsernameAlreadyExists] = "用户名已存在"
	message[CodeInvalidCommentAction] = "这不是您的评论"
	message[CodeLimiterCount] = "请求次数过多，已被限制，稍后再试"
}

func MapErrMsg(errCode int32) string {
	if msg, ok := message[errCode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errCode int32) bool {
	if _, ok := message[errCode]; ok {
		return true
	} else {
		return false
	}
}
