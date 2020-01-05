package main

import (
	"github.com/wazsmwazsm/QueryBuilder/builder"
	"log"
)

func main() {
	sb := builder.NewSQLBuilder()

	sql, err := sb.Table("`test`").
		Select("*").
		GetQuerySQL()
	if err != nil {
		log.Fatal(err)
	}

	params := sb.GetQueryParams()

	log.Println(sql)    // SELECT * FROM `test`
	log.Println(params) // []
}
