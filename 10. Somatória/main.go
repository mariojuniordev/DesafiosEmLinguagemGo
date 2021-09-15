package main

import (
	"fmt"
)

func main() {
	var numero int;

	fmt.Print("Digite um número inteiro e maior que 0: ");
	fmt.Scanln(&numero);

	var somatoria int = 0;

	for i := numero; i > 0; i-- {
		somatoria += i;
	}

	fmt.Printf("A somatória dos números entre 0 e o número digitado é: %d", somatoria);

}