package strings

import "strings"

// 合并字符串
func Concat(s ...string) string {
	return strings.Join(s, "")
}
