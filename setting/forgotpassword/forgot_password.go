package forgotpassword

import "encoding/json"

const SettingName string = "forgot_password_mail"

// MailSettings 表示找回密码的邮件配置
type MailSettings struct {
	// 发送人邮箱地址
	MailFrom string
	// 发送人别名
	MailFromAlias string
	// 邮件标题。可以使用 %VERIFY_CODE% 表示验证码
	MailTitle string
	// 邮件内容。可以使用的变量标识符：
	// 1、%VERIFY_CODE% 表示验证码
	// 2、%DISPLAY_NAME% 表示用户名
	// 3、%USER_EMAIL% 表示用户的 Email
	MailTemplate string
}

// GetSettingsFromString 将 JSON 字符串转换为实体对象。in 必须是 JSON 格式的字符串，否则返回 nil。
func GetSettingsFromString(in string) *MailSettings {
	s := &MailSettings{}
	if err := json.Unmarshal([]byte(in), s); err != nil {
		return nil
	}
	return s
}

// String 将实体对象转换为 JSON 字符串
func (s *MailSettings) String() string {
	b, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	return string(b)
}
