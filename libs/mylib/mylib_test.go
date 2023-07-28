package mylib

import (
	"testing"
)

func TestMylib(t *testing.T) {
	result := Mylib("works")
	if result != "Mylib works" {
		t.Error("Expected Mylib to append 'works'")
	}
}
