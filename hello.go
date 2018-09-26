package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome = "Renan"
	idade := 24 //:= indica que eh uma declaracao de variavel
	var versao float32 = 1.1
	fmt.Println("Hello, sr.", nome, "Sua idade eh", idade)
	fmt.Println("Este programa esta na versao", versao)

	fmt.Println("O tipo da variavel versao eh", reflect.TypeOf(versao))
}
