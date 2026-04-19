package jsonutil

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Aluno struct {
	Id    int
	Nome  string
	Media float64
	Notas []float64
}

func Json() ([]Aluno, error) {

	file, err := os.Open("../../data/dados.json") // ele abre o arquivo e fica lendo ela

	if err != nil {
		return nil, err
	}

	// fmt.Println(file)

	defer file.Close() // pede para o go finalizar a função/tarefa depois que tudo for feito

	byteValue, err := io.ReadAll(file) // vai dar um get e organizar tudo em []bytes (slice em byte)
	// essa função percorre o json

	// fmt.Println(byteValue)

	if err != nil {
		return nil, err
	}

	var alunos []Aluno                       // criar var para que ele receba todo o json em forma de array
	err = json.Unmarshal(byteValue, &alunos) // faz o err pegar a função pq a função vai sempre tentar retornar um erro, então fica mais fácil e clean deixar assim
	// fmt.Println(student)

	if err != nil {
		return nil, err
	}

	for _, aluno := range alunos {
		fmt.Printf("Nome: %v \nNota: %v\n", aluno.Nome, aluno.Notas)
		fmt.Println("------------------")
	}

	return alunos, nil
}
