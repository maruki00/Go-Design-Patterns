package structural_test

import (
	"fmt"
	"testing"
)

type Data struct {
	Name string
	Age  int
}

type Adapter interface {
	GetData() string
}

type JsonAdapter struct {
	d Data
}

func NewJsonAdapter(d Data) *JsonAdapter {
	return &JsonAdapter{
		d: d,
	}
}

func (j *JsonAdapter) GetData() string {
	return fmt.Sprintf(`{"name": "%s", "age": %d}`, j.d.Name, j.d.Age)
}

type XMLAdapter struct {
	d Data
}

func NewXMLAdapter(d Data) *XMLAdapter {
	return &XMLAdapter{
		d: d,
	}
}

func (x *XMLAdapter) GetData() string {
	return fmt.Sprintf(`<data><name>%s</name><age>%d</age></data>`, x.d.Name, x.d.Age)
}

// Test Package
func TestJsonAdapter_GetData(t *testing.T) {
	data := Data{Name: "Alice", Age: 30}
	adapter := NewJsonAdapter(data)
	expected := `{"name": "Alice", "age": 30}`
	if result := adapter.GetData(); result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestXMLAdapter_GetData(t *testing.T) {
	data := Data{Name: "Bob", Age: 25}
	adapter := NewXMLAdapter(data)
	expected := `<data><name>Bob</name><age>25</age></data>`
	if result := adapter.GetData(); result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
