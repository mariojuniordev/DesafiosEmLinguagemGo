package main

import (
	"fmt"
)

func main() {
	var size int;

	fmt.Print("Digite o tamanho da Matriz Quadrada (use números inteiros): ");
	fmt.Scanln(&size);

	for line:=0; line < size; line++ {
		lineArray := "";
		
		for column := 0; column < size; column++ {
			checkLine := "0";
			if line == column {
				checkLine = "1";
			} else {
				checkLine = "0";
			}
			lineArray += checkLine;
		}

		fmt.Println(lineArray);

	}

}