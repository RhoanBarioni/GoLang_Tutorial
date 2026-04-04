package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/RhoanBarioni/GoLang_Tutorial/internal/jsonutil"
	"github.com/RhoanBarioni/GoLang_Tutorial/internal/service"
)

func main() {
	alunos, err := jsonutil.Json()

	if err != nil {
		fmt.Println("Erro ao ler o alunos:", err)
		return
	}

	var nota float64

	for {
		var aluno jsonutil.Aluno
		var index int = -1

		fmt.Print("Digite o nome do Aluno: ")
		fmt.Scan(&aluno.Nome)

		for idx, a := range alunos {
			if a.Nome == aluno.Nome {
				aluno = a // Atribuir o aluno encontrado à variável "aluno"
				index = idx
				break
			}
		}

		for {
			fmt.Println("Para encerrar a adição das notas, colocar '-1'")
			fmt.Print("Nota: ")
			fmt.Scan(&nota)

			if nota == -1 {
				break
			}

			aluno.Notas = append(aluno.Notas, nota)
		}

		// verificar se ja existe e evitar dele duplicar no final
		if index != -1 {
			alunos[index] = aluno // Atualizar o aluno existente na lista
		} else {
			alunos = append(alunos, aluno)
		}

		fmt.Println(alunos)

		fmt.Print("Há mais alunos para serem inseridos? (s/n)")
		var response string
		fmt.Scan(&response)

		if response != "s" {
			break
		}
	}

	for {
		fmt.Println("Deseja deletar algum aluno? n para encerrar e digite o nome para deletar")
		for _, nomes := range alunos{
			fmt.Println(nomes.Nome)
		}
		var alunoNome string
		fmt.Scan(&alunoNome)

		if alunoNome == "n" {
			break
		}

		for i, a := range alunos {
			if a.Nome == alunoNome {
				alunos = append(alunos[:i], alunos[i+1:]...)
				break
			}
		}

		fmt.Println("Deseja deletar outro aluno? s/n")
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
	enderecoJson := "../../data/dados.json"
	os.WriteFile(enderecoJson, novoJson, 0644)
}
