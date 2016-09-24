package main

import (
	"fmt"
	"testing"
)

func TestRandom(t *testing.T) {
	num := random(1, 10)
	fmt.Println(num)
	if num > 9 || num < 1 {
		t.Error("Number out of range")
	}
}
