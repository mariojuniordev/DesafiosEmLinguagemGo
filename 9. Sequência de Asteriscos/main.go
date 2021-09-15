package main

import (
	"fmt"
)

func main() {
	var numero int;

	fmt.Print("Digite um número inteiro: ");
	fmt.Scanln(&numero);

	var aux string = "";

	for i := 1; i <= numero; i++ {
		aux += "*";
		fmt.Println(aux);
	}

}