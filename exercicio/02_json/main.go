package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Student struct {
	Name []map[string]float64 `json:"name"`
}

func main() {
	file, err := os.Open("dados.json")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close() // pede para o go finalizar a função/tarefa depois que tudo for feito

	byteValue, err := io.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	var student Student
    err = json.Unmarshal(byteValue, &student)

    if err != nil {
        log.Fatal("Erro no Unmarshal: ", err)
    }

    fmt.Printf("Nome: %v\n\n", student.Name)

	for _, data := range student.Name{

		fmt.Println(len(data))

		for nome, nota := range data{
			fmt.Printf("Aluno: %v \n Nota: %1.f \n\n", nome, nota)
		}
	}
}
