package e

const (
	Success       = 200
	Error         = 500
	InvalidParams = 400

	// user模块错误
	ErrorExistUser             = 30001 //用户已经存在
	ErrorFailEncryption        = 30002
	ErrorExistUserNotFound     = 30003
	ErrorNotCompare            = 30004
	ErrorAuthToken             = 30005
	ErrorAuthCheckTokenTImeOut = 30006 //token过期
	ErrorUploadFail            = 30007
	ErrorSendEmail             = 30008

	// product 模块错误
)
