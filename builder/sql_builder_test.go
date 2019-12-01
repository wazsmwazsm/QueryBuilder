package builder

import (
	"testing"
)

func TestGetSQLErr(t *testing.T) {
	sb := NewSQLBuilder()

	_, err := sb.GetQuerySQL()
	if err != ErrTableEmpty {
		t.Error("check err")
	}

	sb = NewSQLBuilder()

	_, err = sb.GetInsertSQL()
	if err != ErrTableEmpty {
		t.Error("check err")
	}

	sb = NewSQLBuilder()

	_, err = sb.GetUpdateSQL()
	if err != ErrTableEmpty {
		t.Error("check err")
	}

	sb = NewSQLBuilder()

	_, err = sb.GetDeleteSQL()
	if err != ErrTableEmpty {
		t.Error("check err")
	}

	sb = NewSQLBuilder()

	_, err = sb.Table("test").GetInsertSQL()
	if err != ErrInsertEmpty {
		t.Error("check err")
	}

	sb = NewSQLBuilder()

	_, err = sb.Table("test").GetUpdateSQL()
	if err != ErrUpdateEmpty {
		t.Error("check err")
	}
}
func TestSQLBuilderSelect(t *testing.T) {
	sb := NewSQLBuilder()

	sql, err := sb.Table("test").
		Select("name", "age", "school").
		GetQuerySQL()
	if err != nil {
		t.Error(err)
	}

	expectSQL := "SELECT `name`,`age`,`school` FROM `test`"
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

	expectSQL := "SELECT * FROM `test`"
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
	expectSQL := "SELECT count(`age`), username FROM `test`"
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

	expectSQL := "SELECT `name`,`age`,`school` FROM `test` WHERE `name` = ? AND `age` >= ? OR `name` like ?"
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
	expectSQL := "SELECT `name`,`age`,`school` FROM `test` WHERE `title` = ? AND `name` = ? OR `age` = ? OR `age` = ?"
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
	expectSQL := "SELECT `name`,`age`,`school` FROM `test` WHERE `id` IN (?,?,?) OR `uid` NOT IN (?,?)"
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

func TestSQLBuilderOrWhereIn(t *testing.T) {
	sb := NewSQLBuilder()

	sql, err := sb.Table("test").
		Select("name", "age", "school").
		OrWhereIn("id", 1, 2, 3).
		WhereNotIn("uid", 2, 4).
		GetQuerySQL()
	if err != nil {
		t.Error(err)
	}
	expectSQL := "SELECT `name`,`age`,`school` FROM `test` WHERE `id` IN (?,?,?) AND `uid` NOT IN (?,?)"
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

	expectSQL := "SELECT `name`,`age`,`school` FROM `test` GROUP BY `school`,`class`"
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

	expectSQL := "SELECT `name`,`age`,`school` FROM `test` GROUP BY `school`,`class` HAVING `name` = ?"
	if sql != expectSQL {
		t.Error("sql gen err")
	}

	params := sb.GetQueryParams()

	if params[0].(string) != "a" {
		t.Error("params gen err")
	}
}

func TestSQLBuilderOrHaving(t *testing.T) {
	sb := NewSQLBuilder()

	sql, err := sb.Table("test").
		Select("name", "age", "school").
		GroupBy("school", "class").
		Having("name", "=", "a").
		OrHaving("age", "=", 12).
		GetQuerySQL()
	if err != nil {
		t.Error(err)
	}

	expectSQL := "SELECT `name`,`age`,`school` FROM `test` GROUP BY `school`,`class` HAVING `name` = ? OR `age` = ?"
	if sql != expectSQL {
		t.Error("sql gen err")
	}

	params := sb.GetQueryParams()

	if params[0].(string) != "a" {
		t.Error("params gen err")
	}

	if params[1].(int) != 12 {
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

	expectSQL := "SELECT `name`,`age`,`school` FROM `test`"
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

	expectSQL := "SELECT `name`,`age`,`school` FROM `test` GROUP BY `school`,`class` HAVING `name` = ? AND count(`school`) <= ?"
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

func TestSQLBuilderOrHavingRaw(t *testing.T) {
	sb := NewSQLBuilder()

	sql, err := sb.Table("test").
		Select("name", "age", "school").
		GroupBy("school", "class").
		OrHavingRaw("count(`school`) <= ?", 22).
		GetQuerySQL()
	if err != nil {
		t.Error(err)
	}

	expectSQL := "SELECT `name`,`age`,`school` FROM `test` GROUP BY `school`,`class` HAVING count(`school`) <= ?"
	if sql != expectSQL {
		t.Error("sql gen err")
	}

	params := sb.GetQueryParams()

	if params[0].(int) != 22 {
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
	expectSQL := "SELECT `name`,`age`,`school` FROM `test` ORDER BY `age` ASC"
	if sql != expectSQL {
		t.Error("sql gen err")
	}

}

func TestSQLBuilderOrderBy2(t *testing.T) {
	sb := NewSQLBuilder()

	sql, err := sb.Table("test").
		Select("name", "age", "school").
		OrderBy("ASC", "age", "class").
		GetQuerySQL()
	if err != nil {
		t.Error(err)
	}
	expectSQL := "SELECT `name`,`age`,`school` FROM `test` ORDER BY `age`,`class` ASC"
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
	expectSQL := "SELECT `name`,`age`,`school` FROM `test` LIMIT ? OFFSET ?"
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

	expectSQL := "SELECT `name`,`age`,`school` FROM `test` WHERE `name` = ? AND `age` >= ? ORDER BY `age` DESC LIMIT ? OFFSET ?"
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

func TestSQLJoin(t *testing.T) {
	sb := NewSQLBuilder()

	sql, err := sb.Table("test").
		SelectRaw("`test`.`name`, `test`.`age`, `test2`.`teacher`").
		JoinRaw("LEFT JOIN `test2` ON `test`.`class` = `test2`.`class`").
		Where("age", ">=", 18).
		GetQuerySQL()
	if err != nil {
		t.Error(err)
	}
	expectSQL := "SELECT `test`.`name`, `test`.`age`, `test2`.`teacher` FROM `test`" +
		" LEFT JOIN `test2` ON `test`.`class` = `test2`.`class` WHERE `age` >= ?"
	if sql != expectSQL {
		t.Error("sql gen err")
	}

	params := sb.GetQueryParams()

	if params[0].(int) != 18 {
		t.Error("params gen err")
	}
}

func TestSQLJoinWithParams(t *testing.T) {
	sb := NewSQLBuilder()

	sql, err := sb.Table("test").
		SelectRaw("`test`.`name`, `test`.`age`, `test2`.`teacher`").
		JoinRaw("LEFT JOIN `test2` ON `test`.`class` = `test2`.`class` AND `test`.`num` = ?", 2333).
		Where("age", ">=", 18).
		GetQuerySQL()
	if err != nil {
		t.Error(err)
	}
	expectSQL := "SELECT `test`.`name`, `test`.`age`, `test2`.`teacher` FROM `test`" +
		" LEFT JOIN `test2` ON `test`.`class` = `test2`.`class` AND `test`.`num` = ? WHERE `age` >= ?"
	if sql != expectSQL {
		t.Error("sql gen err")
	}

	params := sb.GetQueryParams()

	if params[0].(int) != 2333 || params[1].(int) != 18 {
		t.Error("params gen err")
	}
}

func TestSQLJoin2(t *testing.T) {
	sb := NewSQLBuilder()

	sql, err := sb.TableRaw("`test` as t1").
		SelectRaw("`t1`.`name`, `t1`.`age`, `t2`.`teacher`, `t3`.`address`").
		JoinRaw("LEFT JOIN `test2` as `t2` ON `t1`.`class` = `t2`.`class`").
		JoinRaw("LEFT JOIN `test3` as t3 ON `t1`.`school` = `t3`.`school`").
		Where("age", ">=", 18).
		GetQuerySQL()
	if err != nil {
		t.Error(err)
	}
	expectSQL := "SELECT `t1`.`name`, `t1`.`age`, `t2`.`teacher`, `t3`.`address` FROM `test` as t1" +
		" LEFT JOIN `test2` as `t2` ON `t1`.`class` = `t2`.`class`" +
		" LEFT JOIN `test3` as t3 ON `t1`.`school` = `t3`.`school` WHERE `age` >= ?"
	if sql != expectSQL {
		t.Error("sql gen err")
	}

	params := sb.GetQueryParams()

	if params[0].(int) != 18 {
		t.Error("params gen err")
	}
}

func TestSQLBuilderInsert(t *testing.T) {
	sb := NewSQLBuilder()

	sql, err := sb.Table("test").
		Insert([]string{"name", "age"}, "jack", 18).
		GetInsertSQL()
	if err != nil {
		t.Error(err)
	}

	expectSQL := "INSERT INTO `test` (`name`,`age`) VALUES (?,?)"
	if sql != expectSQL {
		t.Error("sql gen err")
	}

	params := sb.GetInsertParams()

	if params[0].(string) != "jack" ||
		params[1].(int) != 18 {
		t.Error("params gen err")
	}
}

func TestSQLBuilderUpdate(t *testing.T) {
	sb := NewSQLBuilder()

	sql, err := sb.Table("test").
		Update([]string{"name", "age"}, "jack", 18).
		Where("id", "=", 11).
		GetUpdateSQL()
	if err != nil {
		t.Error(err)
	}

	expectSQL := "UPDATE `test` SET `name` = ?,`age` = ? WHERE `id` = ?"
	if sql != expectSQL {
		t.Error("sql gen err")
	}

	params := sb.GetUpdateParams()

	if params[0].(string) != "jack" ||
		params[1].(int) != 18 ||
		params[2].(int) != 11 {
		t.Error("params gen err")
	}
}

func TestSQLBuilderDelete(t *testing.T) {
	sb := NewSQLBuilder()

	sql, err := sb.Table("test").
		Where("id", "=", 11).
		GetDeleteSQL()
	if err != nil {
		t.Error(err)
	}

	expectSQL := "DELETE FROM `test` WHERE `id` = ?"
	if sql != expectSQL {
		t.Error("sql gen err")
	}

	params := sb.GetDeleteParams()

	if params[0].(int) != 11 {
		t.Error("params gen err")
	}
}

func TestGenPlaceholders(t *testing.T) {
	pss := []string{
		GenPlaceholders(5),
		GenPlaceholders(3),
		GenPlaceholders(1),
		GenPlaceholders(0),
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

func BenchmarkQuery(b *testing.B) {

	for i := 0; i < b.N; i++ {
		sb := NewSQLBuilder()
		_, err := sb.Table("test").
			Select("name", "age", "school").
			Where("name", "=", "jack").
			Where("age", ">=", 18).
			OrderBy("DESC", "age").
			Limit(1, 10).
			GetQuerySQL()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSelect(b *testing.B) {

	for i := 0; i < b.N; i++ {
		sb := NewSQLBuilder()
		_, err := sb.Table("test").
			Select("name", "age", "school").
			GetQuerySQL()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkWhere(b *testing.B) {

	for i := 0; i < b.N; i++ {
		sb := NewSQLBuilder()
		_, err := sb.Table("test").
			Where("age", ">=", 18).
			GetQuerySQL()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkWhereIn(b *testing.B) {

	for i := 0; i < b.N; i++ {
		sb := NewSQLBuilder()
		_, err := sb.Table("test").
			WhereIn("age", 18, 19, 20, 31, 22, 33, 24, 45).
			GetQuerySQL()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkWhereRaw(b *testing.B) {

	for i := 0; i < b.N; i++ {
		sb := NewSQLBuilder()
		_, err := sb.Table("test").
			WhereRaw("`age` >= ?", 18).
			GetQuerySQL()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGroupBy(b *testing.B) {

	for i := 0; i < b.N; i++ {
		sb := NewSQLBuilder()
		_, err := sb.Table("test").
			GroupBy("age").
			GetQuerySQL()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkHaving(b *testing.B) {

	for i := 0; i < b.N; i++ {
		sb := NewSQLBuilder()
		_, err := sb.Table("test").
			SelectRaw("`school`, `class`, COUNT(*) as `ct`").
			GroupBy("school", "class").
			Having("ct", ">", "2").
			GetQuerySQL()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkHavingRaw(b *testing.B) {

	for i := 0; i < b.N; i++ {
		sb := NewSQLBuilder()
		_, err := sb.Table("test").
			SelectRaw("`school`, `class`, COUNT(*)").
			GroupBy("school", "class").
			HavingRaw("COUNT(*) > 2").
			GetQuerySQL()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkInsert(b *testing.B) {

	for i := 0; i < b.N; i++ {
		sb := NewSQLBuilder()
		_, err := sb.Table("test").
			Insert([]string{"name", "class"}, "bob", "2-3").
			GetInsertSQL()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUpdate(b *testing.B) {

	for i := 0; i < b.N; i++ {
		sb := NewSQLBuilder()
		_, err := sb.Table("test").
			Update([]string{"name", "class"}, "bob", "2-3").
			GetUpdateSQL()
		if err != nil {
			b.Fatal(err)
		}
	}
}
