package funciones

import (
	"fmt"
)

func Calculos() {
	
	sum := func(numero1 int, numero2 int) int {
		return numero1 + numero2
	}

	fmt.Println(sum(10, 25))
}