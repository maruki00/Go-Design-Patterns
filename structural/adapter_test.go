package structural_test

import (
	"structural"
	"testing"
)

// Test Package
func TestJsonAdapter_GetData(t *testing.T) {
	data := structural.Data{Name: "Alice", Age: 30}
	adapter := structural.NewJsonAdapter(data)
	expected := `{"name": "Alice", "age": 30}`
	if result := adapter.GetData(); result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestXMLAdapter_GetData(t *testing.T) {
	data := structural.Data{Name: "Bob", Age: 25}
	adapter := structural.NewXMLAdapter(data)
	expected := `<data><name>Bob</name><age>25</age></data>`
	if result := adapter.GetData(); result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
