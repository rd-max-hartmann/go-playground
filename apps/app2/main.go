package main

import (
	"fmt"
	"undefined/libs/mylib"
)

func Hello(name string) string {
	result := "Hello " + mylib.Mylib("Max")
	return result
}

func main() {
	fmt.Println(Hello("app2"))
}
