package main

import (
	"fmt"
)

func main() {
	//Escreva um código que recebe a idade do usuário e retorne se ele é maior ou menor de idade
	var idade int;
		fmt.Printf("Digite a sua idade: ");
		fmt.Scanf("%d", &idade);
		if idade >= 18 {
			fmt.Println("Você é maior de idade!!!");
		} else {
			fmt.Println("Você é menor de idade!!!");
		}
}