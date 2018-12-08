package main

import (
	"fmt"

	"github.com/iamgoangle/go-mod-mod/a"
)

func SayB() {
	fmt.Println("Hey!! Say B")
}

func SayAll() {
	a.SayA()
	SayB()
}

func main() {
	SayAll()
}
