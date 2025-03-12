package structural

import "fmt"

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
	return fmt.Sprintf(`{"name": %s, "age": %d}`, j.d.Name, j.d.Age)
}

type XMLAdapter struct {
	d Data
}

func NewXMLAdapter(d Data) *JsonAdapter {
	return &JsonAdapter{
		d: d,
	}
}

func (x *XMLAdapter) GetData() string {
	return fmt.Sprintf(`<data><name>%s</name><age>%d</age></data>`, x.d.Name, x.d.Age)
}
