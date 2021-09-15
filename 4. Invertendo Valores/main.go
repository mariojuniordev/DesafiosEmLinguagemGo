package main

import (
	"fmt"
)

func main() {
	var value bool = false;
	
	if value == true {
		value = false;
		fmt.Println(value);
	} else {
		value = true;
		fmt.Println(value);
	}

}