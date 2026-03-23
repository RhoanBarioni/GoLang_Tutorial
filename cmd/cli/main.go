package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/RhoanBarioni/GoLang_Tutorial/internal/jsonutil"
	"github.com/RhoanBarioni/GoLang_Tutorial/internal/service"
)

type Aluno struct {
	Nome  string
	Notas []float64
}

func main() {
	jsonutil.Json()

	// Realiza a leitura arquivo
	// Faz um "Unmarshal" para a struct "Aluno"

	var alunos []Aluno
	var nota float64

	for {
		var aluno Aluno
		var nomeAluno string

		fmt.Print("Digite o nome do Aluno: ")
		fmt.Scan(&nomeAluno)

		//Verificar se esse aluno já está presente no slice de Alunos

		// Se sim, você deve atualizar as notas desse aluno
		// Se não, você deve criar um novo aluno

		//Obs: Utilizar a variável "aluno" para armazenar o nome e as notas.

		for {
			fmt.Println("Para encerrar a adição das notas, colocar '-1'")
			fmt.Print("Nota: ")
			fmt.Scan(&nota)

			if nota == -1 {
				break
			}

			aluno.Notas = append(aluno.Notas, nota)

			// aluno.Notas = notas (novo slice aqui)
		}

		alunos = append(alunos, aluno)

		fmt.Println(alunos)

		fmt.Print("Há mais alunos para serem inseridos? (s/n)")
		var response string
		fmt.Scan(&response)

		if response != "s" {
			break
		}
	}

	for i := 0; i < len(alunos); i++ {
		fmt.Println(alunos[i])
		service.Calc(alunos[i].Nome, alunos[i].Notas)
	}

	novoJson, _ := json.Marshal(alunos)
	fmt.Println(string(novoJson))
	enderecoJson := "data/dados.json"
	os.WriteFile(enderecoJson, novoJson, 0644)
	jsonutil.Json()
}
