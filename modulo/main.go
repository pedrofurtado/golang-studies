package main

import (
	"modulo/visibilidade"
	"fmt"
)

func main() {
	p := visibilidade.Pessoa{Name: "Joao", Idade: 23}
	fmt.Printf("%+v\n", p)
}
