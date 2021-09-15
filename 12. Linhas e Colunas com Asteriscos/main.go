//Escreva um algoritmo que receba um valor para linhas e um para colunas e retorne...
//...uma sequência de asteriscos cujos números sejam correspondentes às linhas e colunas recebidas

package main

import (
	"fmt"
)

func main() {
	var line int;
	var column int;

	fmt.Print("Digite um valor para as linhas: ");
	fmt.Scanln(&line);
	fmt.Print("Digite um valor para as colunas: ");
	fmt.Scanln(&column);

	var strAux = "";

	for i := 0; i < column; i++ {
		strAux += "*";
	}

	for i := 0; i < line; i++ {
		fmt.Println(strAux);
	}

}