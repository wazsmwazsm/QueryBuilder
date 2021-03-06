package builder

import (
	"errors"
	"strings"
)

var (
	// ErrTableEmpty table not set
	ErrTableEmpty = errors.New("table empty")
	// ErrInsertEmpty insert content not set
	ErrInsertEmpty = errors.New("insert content empty")
	// ErrUpdateEmpty update content not set
	ErrUpdateEmpty = errors.New("update content empty")
)

// SQLBuilder sql builder
type SQLBuilder struct {
	_select       string
	_insert       string
	_update       string
	_delete       string
	_table        string
	_join         string
	_where        string
	_groupBy      string
	_having       string
	_orderBy      string
	_limit        string
	_insertParams []interface{}
	_updateParams []interface{}
	_whereParams  []interface{}
	_havingParams []interface{}
	_limitParams  []interface{}
	_joinParams   []interface{}
}

// NewSQLBuilder init sql builder
func NewSQLBuilder() *SQLBuilder {
	return &SQLBuilder{}
}

// GetQuerySQL get sql
func (sb *SQLBuilder) GetQuerySQL() (string, error) {
	if sb._table == "" {
		return "", ErrTableEmpty
	}
	var buf strings.Builder

	buf.WriteString("SELECT ")
	if sb._select != "" {
		buf.WriteString(sb._select)
	} else {
		buf.WriteString("*")
	}
	buf.WriteString(" FROM ")
	buf.WriteString(sb._table)
	if sb._join != "" {
		buf.WriteString(" ")
		buf.WriteString(sb._join)
	}
	if sb._where != "" {
		buf.WriteString(" ")
		buf.WriteString(sb._where)
	}
	if sb._groupBy != "" {
		buf.WriteString(" ")
		buf.WriteString(sb._groupBy)
	}
	if sb._having != "" {
		buf.WriteString(" ")
		buf.WriteString(sb._having)
	}
	if sb._orderBy != "" {
		buf.WriteString(" ")
		buf.WriteString(sb._orderBy)
	}
	if sb._limit != "" {
		buf.WriteString(" ")
		buf.WriteString(sb._limit)
	}

	return buf.String(), nil
}

// GetInsertSQL get sql
func (sb *SQLBuilder) GetInsertSQL() (string, error) {
	if sb._table == "" {
		return "", ErrTableEmpty
	}
	if sb._insert == "" {
		return "", ErrInsertEmpty
	}

	var buf strings.Builder

	buf.WriteString("INSERT INTO ")
	buf.WriteString(sb._table)
	buf.WriteString(" ")
	buf.WriteString(sb._insert)

	return buf.String(), nil
}

// GetUpdateSQL get sql
func (sb *SQLBuilder) GetUpdateSQL() (string, error) {
	if sb._table == "" {
		return "", ErrTableEmpty
	}

	if sb._update == "" {
		return "", ErrUpdateEmpty
	}

	var buf strings.Builder

	buf.WriteString("UPDATE ")
	buf.WriteString(sb._table)
	buf.WriteString(" ")
	buf.WriteString(sb._update)
	if sb._where != "" {
		buf.WriteString(" ")
		buf.WriteString(sb._where)
	}

	return buf.String(), nil
}

// GetDeleteSQL get sql
func (sb *SQLBuilder) GetDeleteSQL() (string, error) {
	if sb._table == "" {
		return "", ErrTableEmpty
	}

	var buf strings.Builder

	buf.WriteString("DELETE FROM ")
	buf.WriteString(sb._table)
	if sb._where != "" {
		buf.WriteString(" ")
		buf.WriteString(sb._where)
	}

	return buf.String(), nil
}

// GetQueryParams get params
func (sb *SQLBuilder) GetQueryParams() []interface{} {
	params := []interface{}{}
	params = append(params, sb._joinParams...)
	params = append(params, sb._whereParams...)
	params = append(params, sb._havingParams...)
	params = append(params, sb._limitParams...)
	return params
}

// GetInsertParams get params
func (sb *SQLBuilder) GetInsertParams() []interface{} {
	params := []interface{}{}
	params = append(params, sb._insertParams...)
	return params
}

// GetUpdateParams get params
func (sb *SQLBuilder) GetUpdateParams() []interface{} {
	params := []interface{}{}
	params = append(params, sb._updateParams...)
	params = append(params, sb._whereParams...)
	return params
}

// GetDeleteParams get params
func (sb *SQLBuilder) GetDeleteParams() []interface{} {
	params := []interface{}{}
	params = append(params, sb._whereParams...)
	return params
}

// Table set table
func (sb *SQLBuilder) Table(table string) *SQLBuilder {

	sb._table = table

	return sb
}

// Select set select cols
func (sb *SQLBuilder) Select(cols ...string) *SQLBuilder {
	var buf strings.Builder

	for k, col := range cols {

		buf.WriteString(col)

		if k != len(cols)-1 {
			buf.WriteString(",")
		}
	}

	sb._select = buf.String()

	return sb
}

// Insert set Insert
func (sb *SQLBuilder) Insert(cols []string, values ...interface{}) *SQLBuilder {
	var buf strings.Builder

	buf.WriteString("(")
	for k, col := range cols {

		buf.WriteString(col)

		if k != len(cols)-1 {
			buf.WriteString(",")
		}
	}
	buf.WriteString(") VALUES (")

	for k := range cols {
		buf.WriteString("?")
		if k != len(cols)-1 {
			buf.WriteString(",")
		}
	}
	buf.WriteString(")")

	sb._insert = buf.String()

	for _, value := range values {
		sb._insertParams = append(sb._insertParams, value)
	}

	return sb
}

// Update set update
func (sb *SQLBuilder) Update(cols []string, values ...interface{}) *SQLBuilder {
	var buf strings.Builder

	buf.WriteString("SET ")

	for k, col := range cols {

		buf.WriteString(col)

		buf.WriteString(" = ?")
		if k != len(cols)-1 {
			buf.WriteString(",")
		}
	}

	sb._update = buf.String()

	for _, value := range values {
		sb._updateParams = append(sb._updateParams, value)
	}

	return sb
}

// WhereRaw set where raw string
func (sb *SQLBuilder) WhereRaw(s string, values ...interface{}) *SQLBuilder {
	return sb.whereRaw("AND", s, values)
}

// OrWhereRaw set where raw string
func (sb *SQLBuilder) OrWhereRaw(s string, values ...interface{}) *SQLBuilder {
	return sb.whereRaw("OR", s, values)
}

func (sb *SQLBuilder) whereRaw(operator string, s string, values []interface{}) *SQLBuilder {
	var buf strings.Builder

	buf.WriteString(sb._where) // append

	if buf.Len() == 0 {
		buf.WriteString("WHERE ")
	} else {
		buf.WriteString(" ")
		buf.WriteString(operator)
		buf.WriteString(" ")
	}

	buf.WriteString(s)
	sb._where = buf.String()

	for _, value := range values {
		sb._whereParams = append(sb._whereParams, value)
	}

	return sb
}

// Where set where cond
func (sb *SQLBuilder) Where(field string, condition string, value interface{}) *SQLBuilder {
	return sb.where("AND", condition, field, value)
}

// OrWhere set or where cond
func (sb *SQLBuilder) OrWhere(field string, condition string, value interface{}) *SQLBuilder {
	return sb.where("OR", condition, field, value)
}

func (sb *SQLBuilder) where(operator string, condition string, field string, value interface{}) *SQLBuilder {
	var buf strings.Builder

	buf.WriteString(sb._where) // append

	if buf.Len() == 0 {
		buf.WriteString("WHERE ")
	} else {
		buf.WriteString(" ")
		buf.WriteString(operator)
		buf.WriteString(" ")
	}

	buf.WriteString(field)

	buf.WriteString(" ")
	buf.WriteString(condition)
	buf.WriteString(" ")
	buf.WriteString("?")

	sb._where = buf.String()

	sb._whereParams = append(sb._whereParams, value)

	return sb
}

// WhereIn set where in cond
func (sb *SQLBuilder) WhereIn(field string, values ...interface{}) *SQLBuilder {
	return sb.whereIn("AND", "IN", field, values)
}

// OrWhereIn set or where in cond
func (sb *SQLBuilder) OrWhereIn(field string, values ...interface{}) *SQLBuilder {
	return sb.whereIn("OR", "IN", field, values)
}

// WhereNotIn set where not in cond
func (sb *SQLBuilder) WhereNotIn(field string, values ...interface{}) *SQLBuilder {
	return sb.whereIn("AND", "NOT IN", field, values)
}

// OrWhereNotIn set or where not in cond
func (sb *SQLBuilder) OrWhereNotIn(field string, values ...interface{}) *SQLBuilder {
	return sb.whereIn("OR", "NOT IN", field, values)
}

func (sb *SQLBuilder) whereIn(operator string, condition string, field string, values []interface{}) *SQLBuilder {
	var buf strings.Builder

	buf.WriteString(sb._where) // append

	if buf.Len() == 0 {
		buf.WriteString("WHERE ")
	} else {
		buf.WriteString(" ")
		buf.WriteString(operator)
		buf.WriteString(" ")
	}

	buf.WriteString(field)

	plhs := GenPlaceholders(len(values))
	buf.WriteString(" ")
	buf.WriteString(condition)
	buf.WriteString(" ")
	buf.WriteString("(")
	buf.WriteString(plhs)
	buf.WriteString(")")

	sb._where = buf.String()

	for _, value := range values {
		sb._whereParams = append(sb._whereParams, value)
	}

	return sb
}

// GroupBy set group by fields
func (sb *SQLBuilder) GroupBy(fields ...string) *SQLBuilder {
	var buf strings.Builder

	buf.WriteString("GROUP BY ")

	for k, field := range fields {

		buf.WriteString(field)

		if k != len(fields)-1 {
			buf.WriteString(",")
		}
	}

	sb._groupBy = buf.String()

	return sb
}

// HavingRaw set having raw string
func (sb *SQLBuilder) HavingRaw(s string, values ...interface{}) *SQLBuilder {
	return sb.havingRaw("AND", s, values)
}

// OrHavingRaw set having raw string
func (sb *SQLBuilder) OrHavingRaw(s string, values ...interface{}) *SQLBuilder {
	return sb.havingRaw("OR", s, values)
}

func (sb *SQLBuilder) havingRaw(operator string, s string, values []interface{}) *SQLBuilder {
	var buf strings.Builder

	buf.WriteString(sb._having) // append

	if buf.Len() == 0 {
		buf.WriteString("HAVING ")
	} else {
		buf.WriteString(" ")
		buf.WriteString(operator)
		buf.WriteString(" ")
	}

	buf.WriteString(s)
	sb._having = buf.String()

	for _, value := range values {
		sb._havingParams = append(sb._havingParams, value)
	}

	return sb
}

// Having set having cond
func (sb *SQLBuilder) Having(field string, condition string, value interface{}) *SQLBuilder {
	return sb.having("AND", condition, field, value)
}

// OrHaving set or having cond
func (sb *SQLBuilder) OrHaving(field string, condition string, value interface{}) *SQLBuilder {
	return sb.having("OR", condition, field, value)
}

func (sb *SQLBuilder) having(operator string, condition string, field string, value interface{}) *SQLBuilder {
	if sb._groupBy == "" { // group by not set
		return sb
	}

	var buf strings.Builder

	buf.WriteString(sb._having) // append

	if buf.Len() == 0 {
		buf.WriteString("HAVING ")
	} else {
		buf.WriteString(" ")
		buf.WriteString(operator)
		buf.WriteString(" ")
	}

	buf.WriteString(field)

	buf.WriteString(" ")
	buf.WriteString(condition)
	buf.WriteString(" ")
	buf.WriteString("?")

	sb._having = buf.String()

	sb._havingParams = append(sb._havingParams, value)

	return sb
}

// OrderBy set order by fields
func (sb *SQLBuilder) OrderBy(operator string, fields ...string) *SQLBuilder {
	var buf strings.Builder

	buf.WriteString("ORDER BY ")

	for k, field := range fields {

		buf.WriteString(field)

		if k != len(fields)-1 {
			buf.WriteString(",")
		}
	}

	buf.WriteString(" ")
	buf.WriteString(operator)

	sb._orderBy = buf.String()

	return sb
}

// Limit set limit
func (sb *SQLBuilder) Limit(offset, num interface{}) *SQLBuilder {
	var buf strings.Builder

	buf.WriteString("LIMIT ? OFFSET ?")

	sb._limit = buf.String()

	sb._limitParams = append(sb._limitParams, num, offset)

	return sb
}

// JoinRaw join with raw sql
func (sb *SQLBuilder) JoinRaw(join string, values ...interface{}) *SQLBuilder {
	var buf strings.Builder

	buf.WriteString(sb._join)
	if buf.Len() != 0 {
		buf.WriteString(" ")
	}
	buf.WriteString(join)

	sb._join = buf.String()

	for _, value := range values {
		sb._joinParams = append(sb._joinParams, value)
	}

	return sb
}

// GenPlaceholders generate placeholders
func GenPlaceholders(n int) string {
	var buf strings.Builder

	for i := 0; i < n-1; i++ {
		buf.WriteString("?,")
	}

	if n > 0 {
		buf.WriteString("?")
	}

	return buf.String()
}
