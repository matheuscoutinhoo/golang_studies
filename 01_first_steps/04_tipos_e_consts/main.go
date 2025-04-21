package main

import (
	"fmt"
	"strconv"
)

/*
TIPOS BASICOS:
	- int int8 int16 int32 int64
	- uint uint8 uint16 uint32 uint64 - inteiros positivos
	- bool
	- byte - msm que uint8
	- rune - msm que int32
	- float32 float64
	- complex64 complex128
*/

func main() {
	// conversao de int - float
	x := 33
	f := float32(x)
	fmt.Println(f)

	//inteiro para string
	s := string(x) // converte para uma rune que representa um carcatere na tabela ASC
	fmt.Println(s)
	s = strconv.FormatInt(int64(x), 10) //forma correta de converter strings
	fmt.Println(s)

	// constante - variavel imutavel (caractere, strings. bools e numeros)
	const c = 10 // gofer (:=) nao permitido e tipo nao precisa ser explicito
}
