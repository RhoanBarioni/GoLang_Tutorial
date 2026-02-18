package greetings

import "fmt"

func Hello(name string) string {
	// func em MAIUSCULO é func publica ou exportavel para outros modulos
	// cria uma var e coloca seu tipo de valor depois o tipo de valor recebido de saída
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
	// criar uma var local private (funciona apenas dentro da func) e joga um valor nela
	// o := faz o Go autodeclarar a variavel
	// Usar o \n para dar a quebra de linha
	// usa o % e seu tipos para jogar uma var dentro dele de forma dinamica sem ter q colocar valor fixo
	// usa o % e depois da virgula sempre colocar na ordem q está dentro dos ""
    return message
}