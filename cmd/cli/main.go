package main

import (
	"fmt"
	"log"

	"github.com/RhoanBarioni/GoLang_Tutorial/internal/jsonutil"
	"github.com/RhoanBarioni/GoLang_Tutorial/internal/service"
)

func main() {
	// alunos, err := jsonutil.Json()

	// if err != nil {
	// 	fmt.Println("Erro ao ler o alunos:", err)
	// 	return
	// }

	db, err := service.DbNew()

	if err != nil {
		panic(err)
	}

	var nota float64

	for {
		var aluno jsonutil.Aluno

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
		}

		// SalvarAluno
		err = service.SalvarAluno(db, &aluno)
		if err != nil {
			fmt.Println("Erro ao salvar no banco de dados")
			panic(err)
		}

		fmt.Print("Há mais alunos para serem inseridos? (s/n)")
		var response string
		fmt.Scan(&response)

		if response != "s" {
			break
		}
	}

	//
	fmt.Println(service.GetAlunos(db))

	fmt.Println("Deseja deletar algum aluno? n para encerrar e digite o nome para deletar")

	todosAlunos, err := service.GetAlunos(db)

	if err != nil {
		log.Println(err)
	}

	log.Println(todosAlunos)

	for {
		fmt.Println("Deseja deletar algum aluno? n para encerrar e digite o nome para deletar")

		for _, nomes := range todosAlunos {
			fmt.Println(nomes.Nome)
		}
		var alunoNome string
		fmt.Scan(&alunoNome)

		if alunoNome == "n" {
			break
		}

		service.DeleteAluno(db, alunoNome)

		// for i, a := range todosAlunos {
		// 	if a.Nome == alunoNome {
		// 		alunos = append(alunos[:i], alunos[i+1:]...)
		// 		break
		// 	}
		// }

		fmt.Println("Deseja deletar outro aluno? s/n")
		var response string
		fmt.Scan(&response)
		if response != "s" {
			break
		}
	}

	// novoJson, _ := json.Marshal(alunos)
	// fmt.Println(string(novoJson))
	// enderecoJson := "../../data/dados.json"
	// os.WriteFile(enderecoJson, novoJson, 0644)
}
