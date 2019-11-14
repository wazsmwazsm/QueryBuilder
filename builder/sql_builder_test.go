package builder

import (
	"testing"
)

func TestSQLBuilderSelect(t *testing.T) {
	sb := NewSQLBuilder()

	sql, err := sb.Table("test").
		Select("name", "age", "school").
		GetQuerySQL()
	if err != nil {
		t.Error(err)
	}

	expectSQL := "SELECT `name`,`age`,`school` FROM test"
	if sql != expectSQL {
		t.Error("sql gen err")
	}
}

func TestSQLBuilderSelectAll(t *testing.T) {
	sb := NewSQLBuilder()

	sql, err := sb.Table("test").GetQuerySQL()
	if err != nil {
		t.Error(err)
	}

	expectSQL := "SELECT * FROM test"
	if sql != expectSQL {
		t.Error("sql gen err")
	}
}

func TestSQLBuilderSelectRaw(t *testing.T) {
	sb := NewSQLBuilder()

	sql, err := sb.Table("test").
		SelectRaw("count(`age`), username").
		GetQuerySQL()
	if err != nil {
		t.Error(err)
	}
	expectSQL := "SELECT count(`age`), username FROM test"
	if sql != expectSQL {
		t.Error("sql gen err")
	}
}

func TestSQLBuilderWhere(t *testing.T) {
	sb := NewSQLBuilder()

	sql, err := sb.Table("test").
		Select("name", "age", "school").
		Where("name", "=", "jack").
		Where("age", ">=", 18).
		OrWhere("name", "like", "%admin%").
		GetQuerySQL()
	if err != nil {
		t.Error(err)
	}

	expectSQL := "SELECT `name`,`age`,`school` FROM test WHERE `name` = ? AND `age` >= ? OR `name` like ?"
	if sql != expectSQL {
		t.Error("sql gen err")
	}

	params := sb.GetQueryParams()

	if params[0].(string) != "jack" ||
		params[1].(int) != 18 ||
		params[2].(string) != "%admin%" {
		t.Error("params gen err")
	}
}

func TestSQLBuilderWhereRaw(t *testing.T) {
	sb := NewSQLBuilder()

	sql, err := sb.Table("test").
		Select("name", "age", "school").
		WhereRaw("`title` = ?", "hello").
		Where("name", "=", "jack").
		OrWhereRaw("`age` = ? OR `age` = ?", 22, 25).
		GetQuerySQL()
	if err != nil {
		t.Error(err)
	}
	expectSQL := "SELECT `name`,`age`,`school` FROM test WHERE `title` = ? AND `name` = ? OR `age` = ? OR `age` = ?"
	if sql != expectSQL {
		t.Error("sql gen err")
	}

	params := sb.GetQueryParams()

	if params[0].(string) != "hello" {
		t.Error("params gen err")
	}
	if params[1].(string) != "jack" {
		t.Error("params gen err")
	}
	if params[2].(int) != 22 {
		t.Error("params gen err")
	}
	if params[3].(int) != 25 {
		t.Error("params gen err")
	}
}

func TestSQLBuilderWhereIn(t *testing.T) {
	sb := NewSQLBuilder()

	sql, err := sb.Table("test").
		Select("name", "age", "school").
		WhereIn("id", 1, 2, 3).
		OrWhereNotIn("uid", 2, 4).
		GetQuerySQL()
	if err != nil {
		t.Error(err)
	}
	expectSQL := "SELECT `name`,`age`,`school` FROM test WHERE `id` IN (?,?,?) OR `uid` NOT IN (?,?)"
	if sql != expectSQL {
		t.Error("sql gen err")
	}

	params := sb.GetQueryParams()

	if params[0].(int) != 1 {
		t.Error("params gen err")
	}
	if params[1].(int) != 2 {
		t.Error("params gen err")
	}
	if params[2].(int) != 3 {
		t.Error("params gen err")
	}
	if params[3].(int) != 2 {
		t.Error("params gen err")
	}
	if params[4].(int) != 4 {
		t.Error("params gen err")
	}
}

func TestSQLBuilderGroupBy(t *testing.T) {
	sb := NewSQLBuilder()

	sql, err := sb.Table("test").
		Select("name", "age", "school").
		GroupBy("school", "class").
		GetQuerySQL()
	if err != nil {
		t.Error(err)
	}

	expectSQL := "SELECT `name`,`age`,`school` FROM test GROUP BY `school`,`class`"
	if sql != expectSQL {
		t.Error("sql gen err")
	}
}

func TestSQLBuilderHaving(t *testing.T) {
	sb := NewSQLBuilder()

	sql, err := sb.Table("test").
		Select("name", "age", "school").
		GroupBy("school", "class").
		Having("name", "=", "a").
		GetQuerySQL()
	if err != nil {
		t.Error(err)
	}

	expectSQL := "SELECT `name`,`age`,`school` FROM test GROUP BY `school`,`class` HAVING `name` = ?"
	if sql != expectSQL {
		t.Error("sql gen err")
	}

	params := sb.GetQueryParams()

	if params[0].(string) != "a" {
		t.Error("params gen err")
	}
}

func TestSQLBuilderHavingNotGen(t *testing.T) {
	sb := NewSQLBuilder()

	sql, err := sb.Table("test").
		Select("name", "age", "school").
		Having("name", "=", "a").
		GetQuerySQL()
	if err != nil {
		t.Error(err)
	}

	expectSQL := "SELECT `name`,`age`,`school` FROM test"
	if sql != expectSQL {
		t.Error("sql gen err")
	}

	params := sb.GetQueryParams()

	if len(params) != 0 {
		t.Error("params gen err")
	}
}

func TestSQLBuilderHavingRaw(t *testing.T) {
	sb := NewSQLBuilder()

	sql, err := sb.Table("test").
		Select("name", "age", "school").
		GroupBy("school", "class").
		Having("name", "=", "a").
		HavingRaw("count(`school`) <= ?", 22).
		GetQuerySQL()
	if err != nil {
		t.Error(err)
	}

	expectSQL := "SELECT `name`,`age`,`school` FROM test GROUP BY `school`,`class` HAVING `name` = ? AND count(`school`) <= ?"
	if sql != expectSQL {
		t.Error("sql gen err")
	}

	params := sb.GetQueryParams()

	if params[0].(string) != "a" {
		t.Error("params gen err")
	}
	if params[1].(int) != 22 {
		t.Error("params gen err")
	}
}

func TestSQLBuilderOrderBy(t *testing.T) {
	sb := NewSQLBuilder()

	sql, err := sb.Table("test").
		Select("name", "age", "school").
		OrderBy("ASC", "age").
		GetQuerySQL()
	if err != nil {
		t.Error(err)
	}
	expectSQL := "SELECT `name`,`age`,`school` FROM test ORDER BY `age` ASC"
	if sql != expectSQL {
		t.Error("sql gen err")
	}

}

func TestSQLBuilderLimit(t *testing.T) {
	sb := NewSQLBuilder()

	sql, err := sb.Table("test").
		Select("name", "age", "school").
		Limit(1, 10).
		GetQuerySQL()
	if err != nil {
		t.Error(err)
	}
	expectSQL := "SELECT `name`,`age`,`school` FROM test LIMIT ? OFFSET ?"
	if sql != expectSQL {
		t.Error("sql gen err")
	}

	params := sb.GetQueryParams()
	if params[0].(int) != 10 {
		t.Error("params gen err")
	}
	if params[1].(int) != 1 {
		t.Error("params gen err")
	}
}

func TestSQLBuilderQuery(t *testing.T) {
	sb := NewSQLBuilder()

	sql, err := sb.Table("test").
		Select("name", "age", "school").
		Where("name", "=", "jack").
		Where("age", ">=", 18).
		OrderBy("DESC", "age").
		Limit(1, 10).
		GetQuerySQL()
	if err != nil {
		t.Error(err)
	}

	expectSQL := "SELECT `name`,`age`,`school` FROM test WHERE `name` = ? AND `age` >= ? ORDER BY `age` DESC LIMIT ? OFFSET ?"
	if sql != expectSQL {
		t.Error("sql gen err")
	}

	params := sb.GetQueryParams()

	if params[0].(string) != "jack" ||
		params[1].(int) != 18 ||
		params[2].(int) != 10 ||
		params[3].(int) != 1 {
		t.Error("params gen err")
	}
}

func TestGeneratePlaceholders(t *testing.T) {
	pss := []string{
		generatePlaceholders(5),
		generatePlaceholders(3),
		generatePlaceholders(1),
		generatePlaceholders(0),
	}
	results := []string{
		"?,?,?,?,?",
		"?,?,?",
		"?",
		"",
	}

	for k, ps := range pss {
		if ps != results[k] {
			t.Errorf("%s not equal to %s\n", ps, results[k])
		}
	}

}
