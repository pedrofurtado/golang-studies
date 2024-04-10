package main

import (
	"fmt"
	"runtime"
	"time"
)

type IPAddr [4]byte
func(i IPAddr) String() string {
  return fmt.Sprintf("IP3 -> %d.%d.%d.%d", i[0], i[1], i[2], i[3])
}

type meuTipo int
var x meuTipo
var y int

type Vertex struct {
	X int
	Y int
}

type MyType struct {
	X, Y float64
}

func (v MyType) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *MyType) ScaleWithPointer(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type Person struct {
	Name string
	Age  int
}

type AnotherPerson struct {
	Name string
	Age  int
}

// String() is a method from interface Stringer (defined in fmt package)
func (p Person) String() string {
	return fmt.Sprintf("esse metodo é chamado qdo o fmt vai printar algo na tela. Valor: %v (%v years)", p.Name, p.Age)
}

func soma_mais(y int) func(int) int {
	counter := 0
	fmt.Println("initial counter from counter_of_method_calls", counter)
	return func(x int) int {
	counter += 1
	fmt.Println("updated counter from counter_of_method_calls to", counter)
		return y + x
	}
}

func numero_impar() func() int {
	impar := -1
	return func() int {
		impar += 2
		return impar
	}
}

func say(s string) {
	for i := 0; i < 25; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, i)
	}
}

func main() {
	fmt.Println(x)
	fmt.Printf("x is %T\n", x)
	x = 42
	fmt.Println(x)

	y = int(x)

	fmt.Println(y)
	fmt.Printf("y is %T\n", y)

	////

	i := 0
	count := 15
	for {
		fmt.Println(i)
		i += 2
		if !(i < count) {
			break
		}
	}

	////

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	////

	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

	////

	v := Vertex{1, 2}
	pp := v //faz um clone do ponteiro. a mudança em pp nao afeta o v
	pp.X = 15
	fmt.Println(v)
	fmt.Println(pp)


	r := Vertex{1, 2}
	d := &r //aponta pro msm ponteiro de r. a mudança em d afeta agora o r
	d.X = 15
	fmt.Println(r)
	fmt.Println(d)

	////
	fmt.Println("////")

	names := []string{}
	fmt.Println(names, len(names), cap(names))
	names = append(names, "John")
	fmt.Println(names, len(names), cap(names))
	names = append(names, "Mary")
	fmt.Println(names, len(names), cap(names))
	names = append(names, "Jane")
	fmt.Println(names, len(names), cap(names))

	////

	fmt.Println("////")
	arr := []int{1, 2, 4, 8, 16, 32, 64, 128}

	arr = append(arr, 888)
	arr = append(arr, 999)

	for i, item := range arr {
		fmt.Printf("The item number %d is: %d\n", i, item)
	}

	////

	fmt.Println("////")

	map_vertex := make(map[string]Vertex)

	map_vertex["Bell Labs"] = Vertex{11, 21}
	map_vertex["Otavio James"] = Vertex{0, 9}

	fmt.Println("map inteiro", map_vertex)
	fmt.Println("Bell Labs", map_vertex["Bell Labs"])
	fmt.Println("chave q nao existe", map_vertex["chave q nao existe"])

	////

	fmt.Println("////")
	mm := soma_mais(3)
	fmt.Println("mm(0) ->", mm(0))
	fmt.Println("mm(2) ->", mm(2))
	fmt.Println("mm(12) ->", mm(12))

	////

	fmt.Println("////")
	gerar_numero_impar := numero_impar()
	fmt.Println("numero impar", gerar_numero_impar())
	fmt.Println("numero impar", gerar_numero_impar())
	fmt.Println("numero impar", gerar_numero_impar())
	fmt.Println("numero impar", gerar_numero_impar())
	fmt.Println("numero impar", gerar_numero_impar())

	////

	fmt.Println("////")
	my_type := MyType{3, 4}
	my_type.Scale(10)
	fmt.Println("X:", my_type.X, "| Y:", my_type.Y)
	fmt.Println("---")
	my_type.ScaleWithPointer(2)
	fmt.Println("X:", my_type.X, "| Y:", my_type.Y)

	////

	fmt.Println("////")
	person := Person{Age: 10, Name: "John"}
	another_person := AnotherPerson{Age: 12, Name: "Mary"}
	other_type := MyType{1, 2}
	fmt.Println("Lets print the structs with fields names below")
	fmt.Printf("%+v\n", person)
	fmt.Printf("%+v\n", other_type)
	fmt.Printf("%+v\n", another_person)
	fmt.Println("===")
	fmt.Println("Lets print the structs without fields names below")
	fmt.Printf("%v\n", person)
	fmt.Printf("%v\n", other_type)
	fmt.Printf("%v\n", another_person)

	////

	fmt.Println("////")
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	////

	fmt.Println("////")
	fmt.Println("INICIO GO ROUTINES")
	go say("R")
	fmt.Println("MEIO GO ROUTINES")
	say("S")
	fmt.Println("FIM GO ROUTINES")

	////

	fmt.Println("////")
}
