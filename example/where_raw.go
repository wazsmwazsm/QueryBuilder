package main

import (
	"github.com/wazsmwazsm/QueryBuilder/builder"
	"log"
)

func main() {
	sb := builder.NewSQLBuilder()

	sql, err := sb.Table("test").
		Select("name", "age", "school").
		WhereRaw("`title` = ?", "hello").
		Where("name", "=", "jack").
		OrWhereRaw("`age` = ? OR `age` = ?", 22, 25).
		GetQuerySQL()
	if err != nil {
		log.Fatal(err)
	}

	params := sb.GetQueryParams()

	log.Println(sql)    // SELECT `name`,`age`,`school` FROM test WHERE `title` = ? AND `name` = ? OR `age` = ? OR `age` = ?
	log.Println(params) // [hello jack 22 25]
}
