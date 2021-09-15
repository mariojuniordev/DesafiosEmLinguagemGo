//Escrever um algoritmo que receba um número e retorne o fatorial desse número

package main

import (
	"fmt"
)

func main() {

	var number int;

	fmt.Print("Digite um número inteiro: ");
	fmt.Scanln(&number);

	var fat int = 1;

	for i := number; i > 0; i-- {
		fat *= i;
	}

	fmt.Printf("O fatorial do número digitado é: %d", fat);

}