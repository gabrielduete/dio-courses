package main

import "fmt"

func main() {
	// or not to declare type in variables
	var nome string = "foo"
	var idade int = 24
	var versao float32 = 3.2
	fmt.Println("Olá", nome)
	fmt.Println("Sua idade é", idade)
	fmt.Println("A versão do Go é", versao)
}