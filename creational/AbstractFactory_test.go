package creational_test

import (
	"creational"
	"fmt"
	"testing"
)

func TestGetInstance(t *testing.T) {
	addidas := creational.GetInstance("addidas")
	nike := creational.GetInstance("nike")

	fmt.Println("addidas : ", addidas.MakeShirt())
	fmt.Println("nike : ", nike.MakeShoe())

}
