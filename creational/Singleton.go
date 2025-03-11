package creational

type singleTonObject struct {
	id int
}

var singleTon *singleTonObject

func GetObject() *singleTonObject {
	if singleTon == nil {
		singleTon = &singleTonObject{}
	}
	return singleTon
}

func (o *singleTonObject) GetId() int {
	return o.id
}
func (o *singleTonObject) SetId(id int) {
	o.id = id
}
