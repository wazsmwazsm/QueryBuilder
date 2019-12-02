# QueryBuilder
A sql query builder for golang

# mission

genarate SQL and bind params quikly

# Use

import 
```go
import (
	"github.com/wazsmwazsm/QueryBuilder/builder"
)
```

new sql builder
```go
sb := builder.NewSQLBuilder()
```

> Tips: sql builder is not reusable, one sql builder only for one SQL

## query

### where
```go
package main

import (
	"github.com/wazsmwazsm/QueryBuilder/builder"
	"log"
)

func main() {
    // new sql builder
	sb := builder.NewSQLBuilder()

    // build sql
	sql, err := sb.Table("`test`").
		Select("`name`", "`age`", "`school`").
		Where("`name`", "=", "jack").
		Where("`age`", ">=", 18).
		OrWhere("`name`", "like", "%admin%").
		GetQuerySQL()
	if err != nil {
		log.Fatal(err)
	}
    // get bind params
	params := sb.GetQueryParams()

	log.Println(sql)    // SELECT `name`,`age`,`school` FROM `test` WHERE `name` = ? AND `age` >= ? OR `name` like ?
    log.Println(params) // [jack 18 %admin%]
    
    // now you can use the sql and params to database/sql Query()\Exec() function
    // do someting...
}

```


### where in 

```go
package main

import (
	"github.com/wazsmwazsm/QueryBuilder/builder"
	"log"
)

func main() {
	sb := builder.NewSQLBuilder()

	sql, err := sb.Table("`test`").
		Select("`name`", "`age`", "`school`").
		WhereIn("`id`", 1, 2, 3).
		OrWhereNotIn("`uid`", 2, 4).
		GetQuerySQL()
	if err != nil {
		log.Fatal(err)
	}

	params := sb.GetQueryParams()

	log.Println(sql)    // SELECT `name`,`age`,`school` FROM `test` WHERE `id` IN (?,?,?) OR `uid` NOT IN (?,?)
	log.Println(params) // [1 2 3 2 4]
}
```

### where raw

sometimes you you have more needs for where conditions, you can use raw sql with WhereRaw()\OrWhereRaw() method

```go
package main

import (
	"github.com/wazsmwazsm/QueryBuilder/builder"
	"log"
)

func main() {
	sb := builder.NewSQLBuilder()

	sql, err := sb.Table("`test`").
		Select("`name`", "`age`", "`school`").
		WhereRaw("`title` = ?", "hello").
		Where("`name`", "=", "jack").
		OrWhereRaw("(`age` = ? OR `age` = ?) AND `class` = ?", 22, 25, "2-3").
		GetQuerySQL()
	if err != nil {
		log.Fatal(err)
	}

	params := sb.GetQueryParams()

	log.Println(sql)    // SELECT `name`,`age`,`school` FROM `test` WHERE `title` = ? AND `name` = ? OR (`age` = ? OR `age` = ?) AND `class` = ?
	log.Println(params) // [hello jack 22 25]
}

```

### aggregate func

also, you can use aggregate func with raw sql

```go
package main

import (
	"github.com/wazsmwazsm/QueryBuilder/builder"
	"log"
)

func main() {
	sb := builder.NewSQLBuilder()

	sql, err := sb.Table("`test`").
		Select("count(`age`)", "`username`").
		GetQuerySQL()
	if err != nil {
		log.Fatal(err)
	}

	params := sb.GetQueryParams()

	log.Println(sql)    // SELECT count(`age`), `username` FROM `test`
	log.Println(params) // []
}

```

### group by

```go
package main

import (
	"github.com/wazsmwazsm/QueryBuilder/builder"
	"log"
)

func main() {
	sb := builder.NewSQLBuilder()

	sql, err := sb.Table("`test`").
		Select("`school`", "`class`", "COUNT(*) as `ct`").
		GroupBy("`school`", "`class`").
		Having("`ct`", ">", "2").
		// Having("COUNT(*)", ">", "2"). // same as above
		GetQuerySQL()
	if err != nil {
		log.Fatal(err)
	}

	params := sb.GetQueryParams()

	log.Println(sql)    // SELECT `school`,`class`,COUNT(*) as `ct` FROM `test` GROUP BY `school`,`class` HAVING `ct` > ?
	log.Println(params) // [2]
}
```

such as where, having can also use raw sql with HavingRaw() method
```go
package main

import (
	"github.com/wazsmwazsm/QueryBuilder/builder"
	"log"
)

func main() {
	sb := builder.NewSQLBuilder()

	sql, err := sb.Table("test").
		SelectRaw("`school`, `class`, COUNT(*)").
		GroupBy("school", "class").
		HavingRaw("COUNT(*) > 2").
		GetQuerySQL()
	if err != nil {
		log.Fatal(err)
	}

	params := sb.GetQueryParams()

	log.Println(sql)    // SELECT `school`,`class`,COUNT(*) FROM `test` GROUP BY `school`,`class` HAVING COUNT(*) > 2
	log.Println(params) // []
}

```

### order by / limit

```go
package main

import (
	"github.com/wazsmwazsm/QueryBuilder/builder"
	"log"
)

func main() {
	sb := builder.NewSQLBuilder()

	sql, err := sb.Table("`test`").
		Select("`name`", "`age`", "`school`").
		Where("`name`", "=", "jack").
		Where("`age`", ">=", 18).
		OrderBy("DESC", "`age`", "`class`").
		Limit(1, 10).
		GetQuerySQL()
	if err != nil {
		log.Fatal(err)
	}

	params := sb.GetQueryParams()

	log.Println(sql)    // SELECT `name`,`age`,`school` FROM `test` WHERE `name` = ? AND `age` >= ? ORDER BY `age`,`class` DESC LIMIT ? OFFSET ?
	log.Println(params) // [jack 18 10 1]
}
```

### join

join only provide raw sql mode now

simple join
```go
package main

import (
	"github.com/wazsmwazsm/QueryBuilder/builder"
	"log"
)

func main() {
	sb := builder.NewSQLBuilder()

	sql, err := sb.Table("`test`").
		Select("`test`.`name`", "`test`.`age`", "`test2`.`teacher`").
		JoinRaw("LEFT JOIN `test2` ON `test`.`class` = `test2`.`class`").
		Where("`test`.`age`", ">=", 18).
		GetQuerySQL()
	if err != nil {
		log.Fatal(err)
	}

	params := sb.GetQueryParams()

	log.Println(sql)    // SELECT `test`.`name`,`test`.`age`,`test2`.`teacher` FROM `test` LEFT JOIN `test2` ON `test`.`class` = `test2`.`class` WHERE `test`.`age` >= ?
	log.Println(params) // [18]
}

```

join with params
```go
package main

import (
	"github.com/wazsmwazsm/QueryBuilder/builder"
	"log"
)

func main() {
	sb := builder.NewSQLBuilder()

	sql, err := sb.Table("`test`").
		Select("`test`.`name`", "`test`.`age`", "`test2`.`teacher`").
		JoinRaw("LEFT JOIN `test2` ON `test`.`class` = `test2`.`class` AND `test`.`num` = ?", 2333).
		Where("`test`.`age`", ">=", 18).
		GetQuerySQL()
	if err != nil {
		log.Fatal(err)
	}

	params := sb.GetQueryParams()

	log.Println(sql)    // SELECT `test`.`name`,`test`.`age`,`test2`.`teacher` FROM `test` LEFT JOIN `test2` ON `test`.`class` = `test2`.`class` AND `test`.`num` = ? WHERE `test`.`age` >= ?
	log.Println(params) // [2333 18]
}

```

more
```go
package main

import (
	"github.com/wazsmwazsm/QueryBuilder/builder"
	"log"
)

func main() {
	sb := builder.NewSQLBuilder()

	sql, err := sb.Table("`test` as t1").
		Select("`t1`.`name`", "`t1`.`age`", "`t2`.`teacher`", "`t3`.`address`").
		JoinRaw("LEFT JOIN `test2` as `t2` ON `t1`.`class` = `t2`.`class`").
		JoinRaw("INNER JOIN `test3` as t3 ON `t1`.`school` = `t3`.`school`").
		Where("`t1`.`age`", ">=", 18).
		GroupBy("`t1`.`age`").
		Having("COUNT(`t1`.`age`)", ">", 2).
		OrderBy("DESC", "`t1`.`age`").
		Limit(1, 10).
		GetQuerySQL()
	if err != nil {
		log.Fatal(err)
	}

	params := sb.GetQueryParams()

	log.Println(sql)    // SELECT `t1`.`name`,`t1`.`age`,`t2`.`teacher`,`t3`.`address` FROM `test` as t1 LEFT JOIN `test2` as `t2` ON `t1`.`class` = `t2`.`class` INNER JOIN `test3` as t3 ON `t1`.`school` = `t3`.`school` WHERE `t1`.`age` >= ? GROUP BY `t1`.`age` HAVING COUNT(`t1`.`age`) > ? ORDER BY `t1`.`age` DESC LIMIT ? OFFSET ?
	log.Println(params) // [18 2 10 1]
}

```

## insert

```go
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

```

## update
```go
package main

import (
	"github.com/wazsmwazsm/QueryBuilder/builder"
	"log"
)

func main() {
	sb := builder.NewSQLBuilder()

	sql, err := sb.Table("`test`").
		Update([]string{"`name`", "`age`"}, "jack", 18).
		Where("`id`", "=", 11).
		GetUpdateSQL()
	if err != nil {
		log.Fatal(err)
	}

	params := sb.GetUpdateParams()

	log.Println(sql)    // UPDATE `test` SET `name` = ?,`age` = ? WHERE `id` = ?
	log.Println(params) // [jack 18 11]
}

```

## delete
```go
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

```

# License

The QueryBuilder is open-sourced software licensed under the [MIT license](http://opensource.org/licenses/MIT).
