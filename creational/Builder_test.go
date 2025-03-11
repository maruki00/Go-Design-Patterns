package creational

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

type mockBuilder struct {
	query  *Query
	buffer *bytes.Buffer
}

func (m *mockBuilder) Table(table string) IBuilder {
	m.query._table = table
	return m
}

func (m *mockBuilder) Where(where string) IBuilder {
	m.query._where = where
	return m
}

func (m *mockBuilder) Limit(limit int) IBuilder {
	m.query._limit = limit
	return m
}

func (m *mockBuilder) Offset(offset int) IBuilder {
	m.query._offset = offset
	return m
}

func (m *mockBuilder) Select(columns []string) IBuilder {
	m.query._select = columns
	return m
}

func (m *mockBuilder) Get() {
	query := " select "
	query += strings.Join(m.query._select, ", ")
	query += " From " + m.query._table
	query += " Where " + m.query._where
	query += fmt.Sprintf(" Limit %d Offset %d ", m.query._limit, m.query._offset)
	m.buffer.WriteString(query)
}

func TestQueryBuilder(t *testing.T) {
	buffer := &bytes.Buffer{}
	query := &Query{}
	builder := &mockBuilder{query: query, buffer: buffer}

	builder.Table("users").
		Where("age > 30").
		Select([]string{"id", "name", "email"}).
		Limit(10).
		Offset(5).
		Get()

	expected := " select id, name, email From users Where age > 30 Limit 10 Offset 5 "
	result := buffer.String()

	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}
