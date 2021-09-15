package main

import (
	"fmt"
)

func main() {

	var g1 float32;
	var g2 float32;
	var g3 float32;
	var g4 float32;

	fmt.Print("Digite a sua primeira nota: ");
	fmt.Scanln(&g1);
	fmt.Print("Digite a sua segunda nota: ");
	fmt.Scanln(&g2);
	fmt.Print("Digite a sua terceira nota: ");
	fmt.Scanln(&g3);
	fmt.Print("Digite a sua quarta nota: ");
	fmt.Scanln(&g4);

	average := (g1 + g2 + g3 + g4)/4;


	fmt.Printf("Sua média é %f\n", average)

	if average >= 9 && average <= 10 {
		fmt.Println("A");
	} else if average >= 7.5 && average <= 8.9 {
		fmt.Println("B");
	} else if average >= 6 && average <= 7.4 {
		fmt.Println("C");
	} else if average >= 4 && average <= 5.9 {
		fmt.Println("D");
	} else {
		fmt.Println("F");
	}

}