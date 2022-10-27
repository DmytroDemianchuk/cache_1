package cache

import "fmt"

type cache interface {
	Print()
}

type Set struct {
	name string
}

type Get struct {
	name string
}

type Delete struct {
	name string
}

func (s Set) Print() {
	fmt.Printf("Name: name")
}

func (g Get) Print() {
	fmt.Printf("Name: name")
}

func (d Delete) Print() {
	fmt.Printf("Name: name")
}
