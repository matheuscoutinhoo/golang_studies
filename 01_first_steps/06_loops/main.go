package main

import "fmt"

func main() {
	// for loop basico
	var res int
	for i := 0; i < 10; i++ {
		res += i
	}

	// loop eterno
	for res < 100 { // eterno ate que essa condicao seja true
		res += 1
		fmt.Println(res)

	}

	// loop range - itera em um intervalo de valores
	arr := [5]int{1, 2, 3, 4, 5}
	for _, el := range arr {
		fmt.Println(el)
	}

	for range 2 {
		fmt.Println("ola")
	}

}
