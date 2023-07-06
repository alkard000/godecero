package main

import (
	"godecero/variables"
	"fmt"
)

func main() {
	estado, text := variables.ConviertoTexto(1234)

	fmt.Println(estado)
	fmt.Println(text)
}