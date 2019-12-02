package main

import (
	"github.com/wazsmwazsm/QueryBuilder/builder"
	"log"
)

func main() {
	sb := builder.NewSQLBuilder()

	sql, err := sb.Table("`test`").
		Where("`id`", "=", 11).
		GetDeleteSQL()
	if err != nil {
		log.Fatal(err)
	}

	params := sb.GetDeleteParams()

	log.Println(sql)    // DELETE FROM `test` WHERE `id` = ?
	log.Println(params) // [11]
}
