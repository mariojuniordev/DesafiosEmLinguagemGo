package main

import (
	"fmt"
)

func main() {
	var idadeMaria int = 25;
	var alturaMaria float32 = 1.53;
	
	var suaIdade int;
	var suaAltura float32;

	fmt.Print("Digite a sua idade: ");
	fmt.Scanln(&suaIdade);
	fmt.Print("Digite a sua altura: ");
	fmt.Scanln(&suaAltura);

	if suaIdade >= idadeMaria && suaAltura < alturaMaria {
		fmt.Println("Sim!");
	} else {
		fmt.Println("Não!");
	}
}