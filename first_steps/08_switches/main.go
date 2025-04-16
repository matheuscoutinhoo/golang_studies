package main

import (
	"fmt"
	"time"
)

func main() {
	// Em go, os switches não são contínuos, eles não passam de um para o outro
	// entre switch e if, a performance será a mesma. Escolha por legibilidade.
	do(1)
	do(2)
	fmt.Println(isWeekend(6))
}

func do(x int) {
	switch x {
	case 1:
		fmt.Println(1)
		fallthrough // faz com que o proximo case seja executado
	case 2:
		fmt.Println(2)
	}
}

func isWeekend(x time.Weekday) bool {
	//um switch case pode ter cases com validações múltiplas
	switch x {
	case time.Sunday, time.Saturday:
		return true
	default:
		return false
	}
}
