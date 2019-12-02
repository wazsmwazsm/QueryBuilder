package main

import (
	"github.com/wazsmwazsm/QueryBuilder/builder"
	"log"
)

func main() {
	sb := builder.NewSQLBuilder()

	sql, err := sb.Table("`test`").
		Insert([]string{"`name`", "`age`"}, "jack", 18).
		GetInsertSQL()
	if err != nil {
		log.Fatal(err)
	}

	params := sb.GetInsertParams()

	log.Println(sql)    // INSERT INTO `test` (`name`,`age`) VALUES (?,?)
	log.Println(params) // [jack 18]
}
