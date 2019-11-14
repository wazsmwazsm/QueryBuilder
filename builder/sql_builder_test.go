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

	t.Error(sql)

}
