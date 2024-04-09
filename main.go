package main

import (
	"fmt"
)

type meuTipo int
var x meuTipo
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("x is %T\n", x)
	x = 42
	fmt.Println(x)

	y = int(x)

	fmt.Println(y)
	fmt.Printf("y is %T\n", y)
}
