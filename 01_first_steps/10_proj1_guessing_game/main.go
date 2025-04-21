package main

import (
	"fmt"
	"math/rand"
)

func getGuesses(randNum int) {
	guesses := []int{}

	for i := 0; i < 10; i++ {
		var guess int
		fmt.Printf("Insira o %dº chute: ", i+1)
		fmt.Scanln(&guess)
		guesses = append(guesses, guess)

		if guess == randNum {
			fmt.Printf("Você acertou na %dª tentativa! A resposta era %d\n", i+1, randNum)
			fmt.Print("Chutes: ")
			for i, g := range guesses {
				if i > 0 {
					fmt.Print(", ")
				}
				fmt.Print(g)
			}
			fmt.Println()
			return
		}
	}

	fmt.Printf("Você falhou! A resposta era: %d\n", randNum) // %s trocado por %d
	fmt.Print("Chutes: ")
	for i, g := range guesses {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(g)
	}
	fmt.Println()
}

func main() {
	num := rand.Intn(100)
	getGuesses(num)
}
