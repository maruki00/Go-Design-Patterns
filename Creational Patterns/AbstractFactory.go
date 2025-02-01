package main

import "fmt"

type IFactory interface {
	MakeShoe() any
	MakeShirt() any
}

type IProduct interface {
	GetLogo() string
	GetSize() int
}

func GetInstance(brand string) IFactory {
	switch brand {
	case "nike":
		return &Nike{}
	case "addidas":
		return &Addidas{}
	default:
		panic("invalidd type")
	}
}

type AddidasShoe struct {
	Shoe
}

type AddidasTshirt struct {
	Tshirt
}

type NikeShoe struct {
	Shoe
}

type NikeTshirt struct {
	Tshirt
}

type Addidas struct{}

func (a *Addidas) MakeShoe() any {
	return &AddidasShoe{
		Shoe: Shoe{
			Logo: "addidas shoe",
			Size: 41,
		},
	}
}
func (a *Addidas) MakeShirt() any {
	return &AddidasTshirt{
		Tshirt: Tshirt{
			Logo: "addidas tshirt",
			Size: 5,
		},
	}
}

type Nike struct{}

func (a *Nike) MakeShoe() any {
	return &NikeShoe{
		Shoe: Shoe{
			Logo: "addidas shoe",
			Size: 41,
		},
	}
}
func (a *Nike) MakeShirt() any {
	return &NikeTshirt{
		Tshirt: Tshirt{
			Logo: "nike tshirt",
			Size: 8,
		},
	}
}

type Shoe struct {
	Logo string
	// Color string
	Size int
}

func (s *Shoe) GetLogo() string {
	return s.Logo
}
func (s *Shoe) GetSize() int {
	return s.Size
}

type Tshirt struct {
	Logo string
	// Color string
	Size int
}

func (s *Tshirt) GetLogo() string {
	return s.Logo
}
func (s *Tshirt) GetSize() int {
	return s.Size
}

func main() {
	addidas := GetInstance("addidas")
	nike := GetInstance("nike")

	fmt.Println("addidas : ", addidas.MakeShirt())
	fmt.Println("nike : ", nike.MakeShoe())

}
