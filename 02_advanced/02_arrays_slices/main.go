package main

import "fmt"

// array x slice: slice é um array porém com tamanho dinâmico e sem elementos alocados sequencialmente na memória
func main() {
	// um slice é basicamente um ponteiro para um array
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4]
	arr[2] = 4 // altera o slice e o array
	fmt.Println(slice)

	//slice literal: slice declarado
	slice1 := []int{1, 2, 3} // o compilador cria o array
	fmt.Println(slice1)

	//len - tamanho do slice/array | cap - capacidade do array/slice
	fmt.Println(len(slice1), cap(slice1))

	//append - retorna uma copia de um slice, mas adicionando um novo elemento ao final
	newSlice := append(slice1, 5)
	fmt.Println(newSlice)

	//make(type, len, cap) - evita realocaçÕes de memória em um slice, criando um slice com uma capacidade pré-determinada
	allocSlice := make([]int, 0, 10)
	fmt.Println(len(allocSlice), cap(allocSlice))
}
