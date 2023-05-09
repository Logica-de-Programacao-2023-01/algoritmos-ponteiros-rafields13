// Escreva uma função em Go que receba um ponteiro para um struct Livro com campos título e autor, e altere o título do livro para "Desconhecido" se o autor for "Anônimo".

package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
}

func setUnknown(livro *Livro) {
	if livro.Autor == "Anônimo" {
		livro.Titulo = "Desconhecido"
	}
}

func main() {
	livro := Livro{
		Titulo: "Dom Casmurro",
		Autor:  "Anônimo", // Machado de Assis
	}
	fmt.Println(livro)

	setUnknown(&livro)
	fmt.Print(livro)
}
