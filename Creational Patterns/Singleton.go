package main

import "fmt"

type Object struct {
	id int
}

var singleTon *Object

func GetObject() *Object {
	if singleTon == nil {
		singleTon = &Object{}
	}
	return singleTon
}

func (o *Object) GetId() int {
	return o.id
}
func (o *Object) SetId(id int) {
	o.id = id
}

func main() {
	obj := GetObject()
	obj.SetId(10)
	obj = GetObject()
	fmt.Println(obj.GetId())

}
