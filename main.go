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

type Pessoa struct {
	Nome string
	Idade int
}

type Emprego struct {
	Cargo string
	Empresa string
}

type Profissional struct {
	Pessoa
	Empregos []Emprego
	CarteiraTrabalho string
	Salario int
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

func somaDeVariosNumeros(numbers ...int) (int, int) {
	sum := 0

	for _, n := range numbers {
		sum += n
	}

	return sum, len(numbers)
}

func rodarComDefer() {
	fmt.Println("rodando algo no inicio")
	defer fmt.Println("Rodando defer ...")
	fmt.Println("vai dar pau")
	//panic("erro") // descomente essa linha pra testar o defer
	fmt.Println("deu pau")
	fmt.Println("rodando algo no fim")
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
	pessoa1 := Pessoa{Nome: "João", Idade: 13}
	pessoa2 := Pessoa{
		Nome: "Maria",
		Idade: 44,
	}
	fmt.Println(pessoa1)
	fmt.Println(pessoa2)
	profissional1 := Profissional{
		CarteiraTrabalho: "123.456.79-0",
		Salario: 12500,
		Pessoa: Pessoa{
			Nome: "Claudio",
			Idade: 23,
		},
		Empregos: []Emprego{
			Emprego{
				Cargo: "Desenvolvedor",
				Empresa: "TI S.A",
			},
			Emprego{
				Cargo: "Freelancer",
				Empresa: "MEI",
			},
		},
	}

	profissional1.Empregos = append(profissional1.Empregos, Emprego{Cargo: "Diretor", Empresa: "LG"})
	fmt.Printf("%+v\n", profissional1)
	fmt.Println(profissional1.Pessoa.Nome)
	fmt.Println(profissional1.Nome) // Promoted fields in struct (Nome está dentro de Pessoa. Mas se Profissional nao tiver um atributo chamado Nome, que sobrescreva o valor, o golang "infere" e redireciona pro struct interno de pessoa)
	fmt.Println(profissional1.Empregos)
	fmt.Println(profissional1.Empregos[0])
	fmt.Println(profissional1.Empregos[1].Empresa)
	for i, emprego := range profissional1.Empregos {
		fmt.Printf("O %dº emprego é %s na %s\n", i + 1, emprego.Cargo, emprego.Empresa)
	}

	profissional2 := Profissional{}
	fmt.Println(profissional2)

	////

	fmt.Println("////Structs anonimos")
	meuStruct := struct {
		MeuCampo1 string
		MeuCampo2 int
	}{
		MeuCampo1: "bla",
		MeuCampo2: 44,
	}

	fmt.Println(meuStruct)
	fmt.Println(meuStruct.MeuCampo1)
	fmt.Println(meuStruct.MeuCampo2)

	////

	fmt.Println("////Maps")
	meuMapaPessoas := map[string]Pessoa{}

	meuMapaPessoas["Joao"] = Pessoa{
		Nome: "Joao Silveira",
		Idade: 23,
	}

	meuMapaPessoas["Alexandre"] = Pessoa{
		Nome: "Alexandre Oliveira",
		Idade: 88,
	}

	meuMapaPessoas["Rogerio"] = Pessoa{
		Nome: "Rogerio Smith",
		Idade: 12,
	}
	delete(meuMapaPessoas, "Alexandre")
	fmt.Println(meuMapaPessoas)


	////

	fmt.Println("////Funcao com parametro variatico (variadic function)")
	fmt.Println(somaDeVariosNumeros(1,10,9,10,4))
	listaNumeros := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(somaDeVariosNumeros(listaNumeros...))
	fmt.Println(somaDeVariosNumeros())

	////

	fmt.Println("////Defer")
	rodarComDefer()
}
