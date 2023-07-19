package arreglos_slice

import (
	"fmt"
)

var tablaS []int = []int{1, 2, 3, 4, 5, 6}
var arreglo [10]int = [10]int{1, 2, 3, 4, 5,6 ,7 ,8, 9, 10}

func MuestroSlice() {
	fmt.Println(tablaS)
	porcion := arreglo[3:] // 3 es el indice de inicio
	porcion2 := arreglo[:5] // 5 es el indice de fin
	porcion3 := arreglo[2:5] // 5 es el indice de fin

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad () {
    elementos := make([]int, 5, 20)
	
	fmt.Printf("Largo %d, Capacidad %d\n", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)

	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("Largo %d, Capacidad %d\n", len(nums), cap(nums))
}