package main

import (
	mediaschool "calc"
	jsonutil "jsonutil"
	"fmt"
)

type Aluno struct {
	Nome  string
	Notas []float64
}

func main() {
	jsonutil.Json()

	var alunos []Aluno
	var nota float64
	// var notas []float64
	// var soma float64
	// var nome string

	// pegar o aluno e notas e jogar num arr
	// ay

	for {
		var aluno Aluno

		fmt.Print("Digite o nome do Aluno: ")
		fmt.Scan(&aluno.Nome)

		for {
			fmt.Println("Para encerrar a adição das notas, colocar '-1'")
			fmt.Print("Nota: ")
			fmt.Scan(&nota)

			if nota == -1 {
				break
			}

			aluno.Notas = append(aluno.Notas, nota)
			// soma += nota

		}

		alunos = append(alunos, aluno)
		fmt.Print(alunos)

		fmt.Print("Há mais alunos para serem inseridos? (s/n)")
		var response string
		fmt.Scan(&response)

		if response != "s" {
			break
		}
	}

	for i := 0; i < len(alunos); i++ {
		fmt.Println(alunos[i])
		mediaschool.Calc(alunos[i].Nome, alunos[i].Notas)
	}
	// for i := 0; i < len(alunos); i++{
	// 	fmt.Println(alunos[i].Notas)

	// 	var soma float64
	// 	for v := 0; v < len(alunos[i].Notas); v++{
	// 		// fmt.Println(alunos[i].Notas[v])
	// 		soma += alunos[i].Notas[v]
	// 	}
	// 	fmt.Println("AQ A CONTA PORRA")
	// 	fmt.Println(soma)
	// }
}
