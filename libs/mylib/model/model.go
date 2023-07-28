package model

import "fmt"

type Person struct {
	Name    string
	Address string
}

func (p Person) Test() {
	fmt.Println(p.Name)
}
