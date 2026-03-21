package mediaschool

import (
	"fmt"
)

func Calc(nome string, notas []float64) (float64) {
	var soma float64
	for i := 0; i < len(notas); i++{
		soma += notas[i]
	}
	mediaNota := soma / float64(len(notas))

	if mediaNota > 7 {
		fmt.Printf("O %v passou de ano \n", nome)
	} else if mediaNota <= 7 && mediaNota >= 5 {
		fmt.Printf("O %v tá de recuperação \n", nome)
	} else{
		fmt.Printf("Não tem jeito %v repetiu \n", nome)
	}
	return mediaNota
}