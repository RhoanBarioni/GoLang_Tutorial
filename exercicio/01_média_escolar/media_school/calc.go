package mediaschool

import (
	"fmt"
)

func Calc(notas []float64, soma float64) (float64) {
	mediaNota := soma / float64(len(notas))

	if mediaNota > 7 {
		fmt.Println("Passou de ano, boa")
	} else if mediaNota <= 7 && mediaNota >= 5 {
		fmt.Println("Tá de recuperação, zé")
	} else{
		fmt.Println("Não tem jeito, repetiu")
	}
	return mediaNota
}
