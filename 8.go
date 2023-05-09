// Crie uma função que receba um ponteiro para uma variável como argumento e modifique o valor da variável dentro da função. Certifique-se de que o ponteiro aponte para uma área de memória válida e que a memória seja liberada após o uso.

package main

import "fmt"

func modifyValue(pointer *int, variable int) {
	*pointer = variable
}

func main() {
	variable := 73
	pointer := &variable
	fmt.Println(variable)

	modifyValue(pointer, 37)
	fmt.Print(variable)

	pointer = nil // Liberar o espaço da memória

}
