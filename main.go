package main

import "fmt"

func main() {
	// em go, é possivel declarar variaveis em if statements
	if x := 10; x == 10 {
		fmt.Println("condition ok")
	} else if x > 0 {
		fmt.Println("condition else if ok")
	} else {
		fmt.Println("else ok")
	}

	//short statement
	if meuBool := true; meuBool {
	}
}
