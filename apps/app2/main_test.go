package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	result := Hello("world")
	if result != "Hello Mylib world" {
		t.Error("Expected Hello to append 'world'")
	}
}
