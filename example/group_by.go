package main

import (
	"github.com/wazsmwazsm/QueryBuilder/builder"
	"log"
)

func main() {
	sb := builder.NewSQLBuilder()

	sql, err := sb.Table("test").
		SelectRaw("`school`, `class`, COUNT(*) as `ct`").
		GroupBy("school", "class").
		Having("ct", ">", "2").
		GetQuerySQL()
	if err != nil {
		log.Fatal(err)
	}

	params := sb.GetQueryParams()

	log.Println(sql)    // SELECT `school`, `class`, COUNT(*) as `ct` FROM test GROUP BY `school`,`class` HAVING `ct` > ?
	log.Println(params) // [2]
}
