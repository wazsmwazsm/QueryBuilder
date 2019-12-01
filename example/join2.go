package main

import (
	"github.com/wazsmwazsm/QueryBuilder/builder"
	"log"
)

func main() {
	sb := builder.NewSQLBuilder()

	sql, err := sb.TableRaw("`test` as t1").
		SelectRaw("`t1`.`name`, `t1`.`age`, `t2`.`teacher`, `t3`.`address`").
		JoinRaw("LEFT JOIN `test2` as `t2` ON `t1`.`class` = `t2`.`class`").
		JoinRaw("INNER JOIN `test3` as t3 ON `t1`.`school` = `t3`.`school`").
		WhereRaw("`t1`.`age` >= ?", 18).
		GetQuerySQL()
	if err != nil {
		log.Fatal(err)
	}

	params := sb.GetQueryParams()

	log.Println(sql)    // SELECT `t1`.`name`, `t1`.`age`, `t2`.`teacher`, `t3`.`address` FROM `test` as t1 LEFT JOIN `test2` as `t2` ON `t1`.`class` = `t2`.`class` INNER JOIN `test3` as t3 ON `t1`.`school` = `t3`.`school` WHERE `t1`.`age` >= ?
	log.Println(params) // [18]
}
