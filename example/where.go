package main

import (
	"github.com/wazsmwazsm/QueryBuilder/builder"
	"log"
)

func main() {
	sb := builder.NewSQLBuilder()

	sql, err := sb.Table("test").
		Select("name", "age", "school").
		Where("name", "=", "jack").
		Where("age", ">=", 18).
		OrWhere("name", "like", "%admin%").
		GetQuerySQL()
	if err != nil {
		log.Fatal(err)
	}

	params := sb.GetQueryParams()

	log.Println(sql)    // SELECT `name`,`age`,`school` FROM test WHERE `name` = ? AND `age` >= ? OR `name` like ?
	log.Println(params) // [jack 18 %admin%]
}
