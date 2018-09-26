package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Renan"
	versao := 1.1
	fmt.Println("Hello, sr.", nome)
	fmt.Println("Este programa esta na versao", versao)

	fmt.Println("O tipo da variavel versao eh", reflect.TypeOf(versao))

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do Programa")

	var comando int
	fmt.Scanf("%d", &comando)
	fmt.Scan(&comando)

	fmt.Println("Digitado: ", comando)
	fmt.Println("Address: ", &comando)

	var ad = &comando
	fmt.Println("Type of ad", reflect.TypeOf(ad))
}
