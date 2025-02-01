package main 

import (
	"strings"
)


type IFactory struct {
	MakeShoe() any
	MakeShirt() any
}


func GetInstance(brand string) IFactory {
	switch  brand {
	case "nike":
		return &Nike{}	
	case "addidas":
		return &Addidas{}
	default :
		return nil
	}
}

type IProduct interface {
	GetLogo() string
	GetSize() int	
}

type Tshirt struct {
	logo string
	size int
}
func (t *Tshirt) getLogo() string{
	return t.logo
}  
func (t *Tshirt) getSize() int	 {
	return t.size
}  


type Addids struct {
	
}




type 
