package main

import "fmt"

func Hello(name string) string{
	message := fmt.Sprintf("Salve, %v. Saia daqui agora", name)
	return message
}

func main(){
	fmt.Println(Hello("Rhoan"))
}