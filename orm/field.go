package orm

import (
	"fmt"
	"go/ast"
	"reflect"
	"strings"

	"github.com/tal-tech/go-zero/core/stringx"
)

// Field 表示一个数据库字段的名称和值对象
type Field struct {
	// 原属性名
	Name string
	// 属性对应的字段名
	FieldName string
	// 获取字段的值的方法
	ValueOf func(reflect.Value) (value interface{}, zero bool)
}

// GetUpdateFieldValues 获取 data 对象的字段和值，ignoreFields 表示要忽略的字段，常用在排除主键字段
func GetUpdateFieldValues(data interface{}, ignoreFields ...string) (string, []interface{}) {
	var (
		fields []string
		values []interface{}
	)

	model := reflect.ValueOf(data)
	fieldValues := GetFields(data)
	for _, f := range fieldValues {
		v, isZero := f.ValueOf(model)
		if !isZero &&
			(len(ignoreFields) == 0 ||
				!stringx.Contains(ignoreFields, f.FieldName) ||
				!stringx.Contains(ignoreFields, fmt.Sprintf("`%s`", f.FieldName))) {
			fields = append(fields, fmt.Sprintf("`%s`=?", f.FieldName))
			values = append(values, v)
		}
	}

	updateSQL := ""
	if len(fields) > 0 {
		updateSQL = strings.Join(fields, ", ")
	}
	return updateSQL, values
}

// GetFields 获取 data 对象的字段
func GetFields(data interface{}) []Field {
	var fields []Field

	modelType := reflect.ValueOf(data).Type()
	for modelType.Kind() == reflect.Slice || modelType.Kind() == reflect.Array || modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}

	for i := 0; i < modelType.NumField(); i++ {
		if fieldStruct := modelType.Field(i); ast.IsExported(fieldStruct.Name) {
			field := Field{
				Name:      fieldStruct.Name,
				FieldName: parseTag(fieldStruct.Tag.Get("db")),
			}
			field.ValueOf = func(value reflect.Value) (interface{}, bool) {
				fieldValue := reflect.Indirect(value).Field(fieldStruct.Index[0])
				return fieldValue.Interface(), fieldValue.IsZero()
			}
			fields = append(fields, field)
		}
	}

	return fields
}

func parseTag(tag string) string {
	if len(tag) > 0 && tag[0] != '-' {
		return tag
	}
	return ""
}
