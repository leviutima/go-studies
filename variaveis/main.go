package main

import "fmt"

func main() {

	// Essa é a maneira explícita de declarar uma variavel (declaramos var e o tipo daquela var)
	// var name string = "levi"

	// Já essa é uma maneira curta e abstrata de declarar outra variavel (aqui o go descobre se é int ou string sozinho)
	// idade := 21

	// fmt.Println(name, idade, "anos")

	// nota-se que sem a utilização dos ":" podemos mudar o valor da variável, caso tente usar essa variavel idade
	// novamente com ":=" vai dar erro, go identifica que você está tentando recriar essa váriavel novamente
	// idade = 22
	// fmt.Println(name, idade, "anos")

	// os tipos do go são: string, int, float64 e bool

	// var a int = 10
	// var b float64 = 3.0
	// resultado := a + b    ERRO! não pode somar int com float64 direto
	// No JS 10 + 3.0 funciona liso. Em Go você tem que converter explicitamente:

	// No JS, variável sem valor é undefined. Em Go não existe undefined. Toda variável declarada sem valor
	// recebe um "valor zero" do seu tipo:
	// var nome string    vira "" (string vazia)
	// var idade int      vira 0
	// var ativo bool    vira false
	// var preco float64  vira 0.0

	// Isso é ótimo: nunca tem undefined causando bug. A variável sempre tem um valor válido do tipo dela.

	// Constantes
	// Pra valores que nunca mudam, usa const:
	// const Pi = 3.14
	// const AppName = "go-studies"
	// Igual o const do JS, mas em Go é pra valores realmente fixos em tempo de compilação (números, strings, bools).

	name := "Levi"
	age := 20

	var height float64 = 1.75
	var isDeveloper bool = true

	var city string 
	var savings int 

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("Is developer:", isDeveloper)
	fmt.Println("City (empty):", city)
	fmt.Println("Savings (zero):", savings)

	city = "São Paulo"
	fmt.Println("City now:", city)

	years := 5
	average := float64(savings) / float64(years)
	fmt.Println("Average:", average)
}
