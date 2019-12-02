package main

import (
	"github.com/wazsmwazsm/QueryBuilder/builder"
	"log"
)

func main() {
	sb := builder.NewSQLBuilder()

	sql, err := sb.Table("`test`").
		Select("`school`", "`class`", "COUNT(*)").
		GroupBy("`school`", "`class`").
		HavingRaw("COUNT(*) > 2").
		GetQuerySQL()
	if err != nil {
		log.Fatal(err)
	}

	params := sb.GetQueryParams()

	log.Println(sql)    // SELECT `school`,`class`,COUNT(*) FROM `test` GROUP BY `school`,`class` HAVING COUNT(*) > 2
	log.Println(params) // []
}
