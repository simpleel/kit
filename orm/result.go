package orm

// Result 表示 SQL 执行结果
type Result struct {
	ID int64
}

// LastInsertId 返回 SQL 执行插入后的 ID
func (r Result) LastInsertId() (int64, error) {
	return r.ID, nil
}

// RowsAffected 返回 SQL 影响的行数
func (r Result) RowsAffected() (int64, error) {
	var c int64 = 0

	if r.ID > 0 {
		c = 1
	}

	return c, nil
}
