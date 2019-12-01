package main

import (
	"github.com/wazsmwazsm/QueryBuilder/builder"
	"log"
)

func main() {
	sb := builder.NewSQLBuilder()

	sql, err := sb.Table("test").
		Select("name", "age", "school").
		WhereIn("id", 1, 2, 3).
		OrWhereNotIn("uid", 2, 4).
		GetQuerySQL()
	if err != nil {
		log.Fatal(err)
	}

	params := sb.GetQueryParams()

	log.Println(sql)    // SELECT `name`,`age`,`school` FROM test WHERE `id` IN (?,?,?) OR `uid` NOT IN (?,?)
	log.Println(params) // [1 2 3 2 4]
}
