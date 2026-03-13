package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	file, err := os.Open("dados.json")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close() // pede para o go finalizar a função/tarefa depois que tudo for feito

	byteValue, err := io.ReadAll(file) // vai dar um get e organizar tudo em []bytes (slice em byte)

	if err != nil {
		log.Fatal(err)
	}

	var student []string
	err = json.Unmarshal(byteValue, &student) // faz o err pegar a função pq a função vai sempre tentar retornar um erro, então fica mais fácil e clean deixar assim
	fmt.Println(student)

	if err != nil {
		log.Fatal("Erro no Unmarshal: ", err)
	}

	for _, aluno := range student {
		fmt.Println(aluno)
	}
}
