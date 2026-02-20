package mediaschool

import (
	"fmt"
)
func Calc(notas float64)(float64){
	mediaNota := notas/5

	if mediaNota < 7 {
		fmt.Println("Reprovado, seu burro")
	} else {
		fmt.Println("Aprovado, meu nego bom")
	}
	return mediaNota
}