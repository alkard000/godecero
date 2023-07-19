package middleware

import (
	"fmt"
)

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func dividir(a, b int) int {
	return a / b
}

func MiMiddleware() {
	fmt.Println("Mi Middleware")

	result := operacionesMiddleware(sumar)(2, 3)
	fmt.Println(result)

	result = operacionesMiddleware(restar)(10, 6)
	fmt.Println(result)

	result = operacionesMiddleware(multiplicar)(2, 4)
	fmt.Println(result)

	result = operacionesMiddleware(dividir)(10, 20)
	fmt.Println(result)
}

func operacionesMiddleware(operacion func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de Operaciones Middleware")
		return operacion(a, b)
	}
}
