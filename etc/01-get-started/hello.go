// package main

// import "fmt"

// func main() {
//     fmt.Println("Hi,  Welcome!")
// }

package main

import (
    "fmt"

    "greetings"
)

func main() {
    message := greetings.Hello("Gladys")
	// usar o nome importado no IMPORT para usar uma função de outro modulo
	// importei o greetings no import e usa greetings.FUNÇÃO
    fmt.Println(message)
}