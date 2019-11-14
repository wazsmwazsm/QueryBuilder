package builder

import (
	"testing"
)

func TestSQLBuilderQuery(t *testing.T) {
	sb := NewSQLBuilder()

	sql, err := sb.Table("test").
		Select([]string{"name", "age", "school"}).
		Where("name", "=", "jack").
		Where("age", ">=", 18).
		OrderBy([]string{"age"}, "DESC").
		Limit(1, 10).
		GetQuerySQL()
	if err != nil {
		t.Error(err)
	}

	expectSQL := "SELECT `name`,`age`,`school` FROM test WHERE `name` = ? AND `age` >= ? ORDER BY `age` DESC LIMIT ? OFFSET ?"
	if sql != expectSQL {
		t.Error("sql gen err")
	}

	params := sb.GetParams()

	if params[0].(string) != "jack" ||
		params[1].(int) != 18 ||
		params[2].(int) != 10 ||
		params[3].(int) != 1 {
		t.Error("params gen err")
	}
}
