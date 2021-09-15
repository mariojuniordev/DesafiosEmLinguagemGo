package main

import (
	"fmt"
)

func main() {
	var inicio int;
	var fim int;

	fmt.Print("Digite um número de início: ");
	fmt.Scanln(&inicio);
	fmt.Print("Digite um número de fim: ");
	fmt.Scanln(&fim);

	if inicio < fim {
		for i := inicio; i <= fim; i++ {
			fmt.Println(i);
		}
	} else {
		for i := inicio; i >= fim; i-- {
			fmt.Println(i);
		}
	}

}