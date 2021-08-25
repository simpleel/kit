package setting

import "encoding/json"

// Settings 表示一个应用配置的实体对象
type Settings interface {
	Name() string
}

// Unmarshal 将 JSON 字符串转换为实体对象，in 必须是 JSON 格式的字符串。
func Unmarshal(settings Settings, in string) error {
	return json.Unmarshal([]byte(in), settings)
}

// Marshal 将实体对象转换为 JSON 字符串
func Marshal(settings Settings) string {
	b, err := json.Marshal(settings)
	if err != nil {
		return ""
	}
	return string(b)
}
