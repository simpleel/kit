package slice

// Prepend 追加到 slice 最前面
func Prepend(dst []interface{}, src interface{}) []interface{} {
	args := make([]interface{}, len(dst)+1)

	args[0] = src
	for i, v := range dst {
		args[i+1] = v
	}

	return args
}
