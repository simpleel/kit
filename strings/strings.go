package strings

import (
	"strings"
	"unicode"
)

// 合并字符串
func Concat(s ...string) string {
	return strings.Join(s, "")
}

// 驼峰转换为蛇形（下划线）
func ToSnakeCase(s string) string {
	if s == "" {
		return s
	}
	if len(s) == 1 {
		return strings.ToLower(s)
	}

	source := []rune(s)
	dist := strings.Builder{}
	dist.Grow(len(s) + len(s)/3) // avoid reallocation memory, 33% ~ 50% is recommended
	skipNext := false
	for i := 0; i < len(source); i++ {
		cur := source[i]
		switch cur {
		case '-', '_':
			dist.WriteRune('_')
			skipNext = true
			continue
		}
		if unicode.IsLower(cur) || unicode.IsDigit(cur) {
			dist.WriteRune(cur)
			continue
		}

		if i == 0 {
			dist.WriteRune(unicode.ToLower(cur))
			continue
		}

		last := source[i-1]
		if (!unicode.IsLetter(last)) || unicode.IsLower(last) {
			if skipNext {
				skipNext = false
			} else {
				dist.WriteRune('_')
			}
			dist.WriteRune(unicode.ToLower(cur))
			continue
		}
		dist.WriteRune(unicode.ToLower(cur))
	}

	return dist.String()
}
