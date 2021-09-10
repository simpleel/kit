package query

import (
	"fmt"
	"strings"
)

type condition struct {
	params []string
	values []interface{}
}

// NewCondition 返回 SQL 条件构造对象
func NewCondition() *condition {
	return &condition{}
}

// Add 添加一个查询参数和值
func (c *condition) Add(param string, value interface{}) {
	c.params = append(c.params, param)
	c.values = append(c.values, value)
}

// SQLParams 返回 SQL 查询参数字符串
func (c *condition) SQLParams() string {
	if len(c.params) == 0 {
		return ""
	}

	return strings.Join(c.params, " and ")
}

// SQLWhere 返回有 where 字符串的查询参数字符串
func (c *condition) SQLWhere() string {
	var s = c.SQLParams()
	if len(s) == 0 {
		return s
	}
	return fmt.Sprintf(" where %s", s)
}

// SQLValues 返回 SQL 查询参数
func (c *condition) SQLValues() []interface{} {
	return c.values
}
