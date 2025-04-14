package main

import (
	"fmt"
)

// funcoes podem receber parametros e tambem retorná-los, assim como podem retornar mais de um valor
func dividir(a, b int) (int, int) {
	return a / b, a % b
}

// Higher Order Function - função que retorna outra função
func somarHof(a int) func(int) int {
	// a func abaixo é um closure, uma função que captura uma var do escopo acima
	return func(b int) int {
		return a + b
	}
}

// função variãtica: recebe um numero variavel de parametros
func somarVariatica(nums ...int) int {
	var out int
	for _, n := range nums {
		out += n + 1
	}
	return out
}

func main() {
	divisao, resto := dividir(3, 2)
	fmt.Println(divisao, resto)
	hofResult := somarHof(3)
	fmt.Println(hofResult(2))
	fmt.Println(somarVariatica(1, 2, 3, 4))
}
