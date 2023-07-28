package main

import (
	"fmt"
	"playground/libs/mylib"
)

func Hello(name string) string {
	result := "Hello " + mylib.Mylib(name)
	return result
}

func main() {
	fmt.Println(Hello("app2"))
}
