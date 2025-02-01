package main

import (
	"fmt"
	"strings"
)

type IBuilder interface {
	Table(string) IBuilder
	Where(string) IBuilder
	Limit(int) IBuilder
	Offset(int) IBuilder
	Select([]string) IBuilder
	Get()
}

type Query struct {
	_table  string
	_where  string
	_limit  int
	_offset int
	_select []string
}

func (obj *Query) Table(table string) IBuilder {
	obj._table = table
	return obj
}

func (obj *Query) Where(where string) IBuilder {
	obj._where = where
	return obj
}

func (obj *Query) Limit(limit int) IBuilder {
	obj._limit = limit
	return obj
}

func (obj *Query) Offset(offset int) IBuilder {
	obj._offset = offset
	return obj
}

func (obj *Query) Select(columns []string) IBuilder {
	obj._select = columns
	return obj
}

func (obj *Query) Get() {

	query := " select "
	query += strings.Join(obj._select, ", ")
	query += " From " + obj._table
	query += " Where " + obj._where
	query += fmt.Sprintf(" Limit %d Offset %d ", obj._limit, obj._offset)
	fmt.Println(query)
}

func main() {

	q := &Query{}
	q.Table("tab1").Select([]string{"id", "name"}).Where("id = 1").Limit(10).Offset(10).Get()
}

func helloword() string {
	return "hello_world"
}
