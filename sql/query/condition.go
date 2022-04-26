package query

import (
	"fmt"
	"strings"

	kitstrings "go.simpleel.com/kit/strings"
)

type condition struct {
	params []string
	values []interface{}
}

// 返回 SQL 条件构造对象。
// 同时只支持一种条件 and/or 判断，如果需要 and/or 拼接，可以创建两个不同的条件对象。
func NewCondition() *condition {
	return &condition{}
}

// 添加一个查询参数和值
func (c *condition) Add(param string, value interface{}) {
	c.params = append(c.params, param)
	c.values = append(c.values, value)
}

// 返回 and 查询条件
func (c *condition) AndParams() string {
	if len(c.params) == 0 {
		return ""
	}

	return kitstrings.Concat("(", strings.Join(c.params, " and "), ")")
}

// 返回 or 查询条件
func (c *condition) OrParams() string {
	if len(c.params) == 0 {
		return ""
	}

	return kitstrings.Concat("(", strings.Join(c.params, " or "), ")")
}

// 返回有 where 字符串的查询参数字符串
func (c *condition) SQLWhere() string {
	var s = c.AndParams()
	if len(s) == 0 {
		return s
	}
	return fmt.Sprintf(" where %s", s)
}

// 返回 SQL 查询参数
func (c *condition) SQLValues() []interface{} {
	return c.values
}
