package main

import (
	"fmt"
)

func main() {
	var numero int;

	fmt.Print("Digite um número inteiro: ");
	fmt.Scanln(&numero);

	if numero >= 0 {
		for i := numero; i >= 0; i-- {
			fmt.Println(i);
		}
	} else {
		fmt.Println(numero);
	}

}