package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	file, err := os.Open("dados.json") // ele abre o arquivo e fica lendo ela

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(file)

	defer file.Close() // pede para o go finalizar a função/tarefa depois que tudo for feito

	byteValue, err := io.ReadAll(file) // vai dar um get e organizar tudo em []bytes (slice em byte)
	// essa função percorre o json

	fmt.Println(byteValue)

	if err != nil {
		log.Fatal(err)
	}

	var student []string // criar var para que ele receba todo o json em forma de array
	err = json.Unmarshal(byteValue, &student) // faz o err pegar a função pq a função vai sempre tentar retornar um erro, então fica mais fácil e clean deixar assim
	fmt.Println(student)

	if err != nil {
		log.Fatal("Erro no Unmarshal: ", err)
	}

	for _, aluno := range student {
		fmt.Println(aluno)
	}
}
