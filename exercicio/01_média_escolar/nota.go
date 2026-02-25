package main

import (
	mediaschool "calc"
	"fmt"
)

func main() {
	var nota float64
	var notas []float64
	var soma float64

	for{
		fmt.Println("Para encerrar a adição das notas, colocar -1")
		fmt.Print("Nota: ")
		fmt.Scan(&nota)

		if (nota == -1){
			break
		}

		notas = append(notas, nota)
		soma += nota
	}

	fmt.Println(mediaschool.Calc(notas, soma))
}
