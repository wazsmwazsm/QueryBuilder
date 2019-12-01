package main

import (
	"github.com/wazsmwazsm/QueryBuilder/builder"
	"log"
)

func main() {
	sb := builder.NewSQLBuilder()

	sql, err := sb.Table("test").
		Update([]string{"name", "age"}, "jack", 18).
		Where("id", "=", 11).
		GetUpdateSQL()
	if err != nil {
		log.Fatal(err)
	}

	params := sb.GetUpdateParams()

	log.Println(sql)    // UPDATE test SET `name` = ?,`age` = ? WHERE `id` = ?
	log.Println(params) // [jack 18 11]
}
