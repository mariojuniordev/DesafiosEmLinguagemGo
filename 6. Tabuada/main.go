package main

import (
	"fmt"
)

func main() {

	var value int;

	fmt.Print("Digite um valor inteiro: ");
	fmt.Scanln(&value);

	fmt.Println("A TABUADA DO NÚMERO DIGITADO É: ")

	for i := 0; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", value, i, value*i);
	}

}