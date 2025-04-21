package main

import "fmt"

func main() {
	// arrays tem tamanho constante, oq nao pode ser alterado
	arr := [3]int{}      //se vazio, os 3 valores serão 0
	arr2 := [3]int{2: 2} //definindo que na posição 2 há o elemento 2
	fmt.Println(arr, arr2)
}
