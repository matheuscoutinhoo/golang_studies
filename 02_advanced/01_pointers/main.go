package main

import "fmt"

func main() {
	// ponteiro - variavel que aponta para um endereço de memória
	x := 10
	pointer := &x
	fmt.Printf("Endereço de x: %v | valor de x: %d", pointer, *pointer)

	// passagem por referencia: como o endereço é passado como param, a funcao acessa a variavel em si e não um copia
	num := 1
	take(&num)
	fmt.Println(num)

	point := create()
	fmt.Print(*point, point)
}

func take(num *int) {
	*num = 100
}

func create() *int { // em go, ponteiros podem ser retornos de funçoes
	x := 10
	return &x
}
