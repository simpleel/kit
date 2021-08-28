package errcode

const (
	// ErrInvalidParamCode 参数错误
	ErrInvalidParamCode = -1
	// SuccessCode 操作成功
	SuccessCode = 0
	// 操作失败
	FailedCode = 40001

	// 数据不存在
	ErrNotExistsCode = 40100
	// 已存在相同数据
	ErrExistsCode = 40101

	// 管理员被禁用
	ErrAdminDenyCode = 43001
	ErrAdminDenyMsg  = "管理员被禁用"

	ErrUserNotExistsCode         = 43100
	ErrUserNotExistsMsg          = "用户不存在"
	ErrUserExistsCode            = 43101
	ErrUserExistsMsg             = "用户已存在"
	ErrUserInvalidPasswordCode   = 43102
	ErrUserInvalidPasswordMsg    = "密码错误"
	ErrInvalidFormatPasswordCode = 43103
	ErrInvalidFormatPasswordMsg  = "密码格式不正确"
	ErrUserDenyCode              = 43104
	ErrUserDenyMsg               = "用户被禁用"
	ErrUserInvalidProviderCode   = 43107
	ErrUserInvalidProviderMsg    = "不支持的登录方式"
	ErrUserEmailNotVerifyCode    = 43108

	// 访问被拒绝
	ErrAccessDenyCode = 43003

	// 验证失败
	ErrAuthConfigCode = 47001
	ErrAuthFaildCode  = 47003

	// 接口校验失败
	ErrAccessTokenInvalidCode = 43000
	ErrAccessTokenInvalidMsg  = "接口校验失败"

	// 服务器内部错误
	ErrServerCode = 50000
	ErrServerMsg  = "服务器内部错误"
)
