package mediaschool

import (
	"fmt"
)

func Calc(nome string, notas []float64, soma float64) (float64) {
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