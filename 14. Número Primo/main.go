//Escreva um algoritmo que recebe um número inteiro e retorne se o número digitado é primo ou não

package main

import (
	"fmt"
)

func main() {
	var number int;
	var qttDividers int = 0;

	fmt.Print("Digite um número inteiro: ");
	fmt.Scanln(&number);

	for i := 1; i <= number; i++ {
		if number % i == 0 {
			qttDividers++;
		}
	}

	if qttDividers == 2 {
		fmt.Printf("%d é primo!", number);
	} else {
		fmt.Printf("%d não é primo!", number);
	}

}