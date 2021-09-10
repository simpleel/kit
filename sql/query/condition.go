package query

import (
	"fmt"
	"strings"
)

type condition struct {
	params []string
	values []interface{}
}

func NewCondition() *condition {
	return &condition{}
}

func (c *condition) Add(param string, value interface{}) {
	c.params = append(c.params, param)
	c.values = append(c.values, value)
}

func (c *condition) SQLParams() string {
	if len(c.params) == 0 {
		return ""
	}

	return strings.Join(c.params, " and ")
}

func (c *condition) SQLWhere() string {
	var s = c.SQLParams()
	if len(s) == 0 {
		return s
	}
	return fmt.Sprintf(" where %s", s)
}

func (c *condition) SQLValues() []interface{} {
	return c.values
}
