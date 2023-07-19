package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func MiNombreLento(nombre string, canalST chan bool) {
	letras := strings.Split(nombre, "")

	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}

	canalST <- true
}
