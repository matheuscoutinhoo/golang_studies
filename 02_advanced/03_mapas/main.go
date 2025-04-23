package main

import "fmt"

func main() {
	// map - é uma coleção de dados que relaciona chaves e valores
	m := map[string]string{"pessoa": "pedro"}
	fmt.Println(m)

	// para recuperar um valor de um map, basta fazer uma seleção por uma chave
	fmt.Println(m["pessoa"])  // retorna o valor associado a essa chave
	val, exist := m["pessoa"] //recupera o valor e true ou false caso exista ou vice-versa
	fmt.Println(val, exist)

	//a função delete() deleta uma chave de um mapa
	delete(m, "pessoa")
	fmt.Println(m["pessoa"])

	// a função clear() limpa um map mas mantém sua capacidade, ou seja, mantém o espaço alocado na memória
	clear(m)
	fmt.Println(m["pessoa"])

	//iterando em um mapa
	for key, value := range m {
		fmt.Println(key, value)
	}
}
