package main

import (
	"godecero/middleware"
)

func main() {
	/*number, text := ejercicios.TwoValues("1234")

	fmt.Println(number)
	fmt.Println(text)*/

	/*teclado.IngresoNumeros()*/

	/*files.SumaTabla()*/
	/*canalST := make(chan bool)
	go goroutines.MiNombreLento("Ivan", canalST)
	defer func() {
		<-canalST
	}()
	fmt.Println("Estoy aca")*/

	middleware.MiMiddleware()
}
