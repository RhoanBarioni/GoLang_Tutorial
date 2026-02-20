package mediaschool

import (
	"fmt"
)
func Calc(nota1, nota2, nota3, nota4, nota5 float64)(float64){
	mediaNota := (nota1 + nota2 + nota3 + nota4 + nota5)/5

	if mediaNota < 7 {
		fmt.Println("Reprovado, seu burro")
	} else {fmt.Println("Aprovado, meu nego bom")}
	return mediaNota
}