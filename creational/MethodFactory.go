package creational

import (
	"errors"
)

var lastId int = 0

type Object struct {
	Id      int
	Quatity int
	Address string
}

func ObjectFactory(Quatity int, Address string) (*Object, error) {
	if Quatity <= 0 {
		return nil, errors.New("invalid quantity")
	}
	if Address == "" {
		return nil, errors.New("invalid address")
	}
	lastId++
	return &Object{
		Id:      lastId,
		Quatity: Quatity,
		Address: Address,
	}, nil
}

// func main() {
// 	obj, err := ObjectFactory(10, "address 1")
// 	if err != nil {
// 		panic(err)
// 	}
// 	obj1, err := ObjectFactory(10, "address 1")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("result %#v, \n", obj)
// 	fmt.Printf("result %#v, \n", obj1)
// }
