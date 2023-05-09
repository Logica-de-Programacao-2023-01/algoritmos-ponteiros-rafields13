// Implemente uma função que receba um ponteiro para uma struct "Livro" com campos "Título", "Autor" e "Preço", e que adicione um desconto de 10% no preço do livro.

package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Price  float64
}

func tenPercentDiscount(book *Book) {
	book.Price = book.Price - (book.Price * 0.1)
}

func main() {
	book := Book{
		Title:  "Hamlet",
		Author: "William Shakespeare",
		Price:  38,
	}
	fmt.Println(book)

	tenPercentDiscount(&book)
	fmt.Print(book)
}
