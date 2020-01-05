package main

import (
	"github.com/wazsmwazsm/QueryBuilder/builder"
	"log"
)

func main() {
	sb := builder.NewSQLBuilder()

	sql, err := sb.Table("`test`").
		Select("`age`", "COUNT(age)").
		GetQuerySQL()
	if err != nil {
		log.Fatal(err)
	}

	params := sb.GetQueryParams()

	log.Println(sql)    // SELECT `age`,COUNT(age) FROM `test`
	log.Println(params) // []
}
