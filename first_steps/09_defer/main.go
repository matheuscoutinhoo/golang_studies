package main

import "fmt"

func main() {
	// defer adia a chamada de uma função até que a main a retorne
	// todo o fluxo é executado, só então, antes voltar para a main o defer é executado
	doDefer()
}

func doDefer() {
	defer fmt.Println("Hello")
	fmt.Println("World")
}

// quando há mais de um defer, eles são executados em ordem LIFO
