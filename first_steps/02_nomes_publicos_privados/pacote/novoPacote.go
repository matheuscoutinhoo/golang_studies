package pacote

import (
	"fmt"
	"pacoteo/pacote/internal" // por ser pai do internal, consegue acessar
)

var Foo string = "ddd"  // variavel publica (pode ser importada)
var foo2 string = "ddd" // variavel privada (nao pode ser importada)

func mostrarInterna() {
	fmt.Println(internal.VarInterna)
	fmt.Println(foo2)
}
