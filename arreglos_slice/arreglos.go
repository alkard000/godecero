package arreglos_slice

import (
	"fmt"
)

var tabla [10]int = [10]int{1, 2, 3, 4, 5, 6, 7}
var matriz [30][30]int

func MuestroArreglos() {
	tabla[7] = 33
	tabla[2] = 54
	tabla2 := [10]string{"Hola", "Mundo", "Como", "Estan"}
	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[7][24] = 15
	fmt.Println(matriz)
}