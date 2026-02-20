package main

import (
	"fmt"
	"calc"
)

func main(){
	var nota1, nota2, nota3, nota4, nota5 float64
	fmt.Println("Primeira Nota: ")
	fmt.Scanln(&nota1)
	fmt.Println("Segunda Nota: ")
	fmt.Scanln(&nota2)
	fmt.Println("Terceira Nota: ")
	fmt.Scanln(&nota3)
	fmt.Println("Quarta Nota: ")
	fmt.Scanln(&nota4)
	fmt.Println("Quinta Nota: ")
	fmt.Scanln(&nota5)

	notas := nota1 + nota2 + nota3 + nota4 + nota5

	fmt.Println(mediaschool.Calc(notas))
}