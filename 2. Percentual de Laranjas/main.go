package main

import (
	"fmt"
)

func main() {

	var totalLaranjas float32 = 4;
	var qtdLaranjas float32 = 1;
	var percentualRestante float32 = (1 -(qtdLaranjas/totalLaranjas))*100;

	fmt.Printf("O percentual restante de Laranjas é: %f %%", percentualRestante);

}