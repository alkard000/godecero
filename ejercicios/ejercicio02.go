package ejercicios

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

var numero1 int
var err error
var texto string

func InputData() string {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese el primer numero: ")
		if scanner.Scan() {
			numero1, err = strconv.Atoi(scanner.Text())
			if err!= nil {
				continue
			}else{
				break
			}
		}		
	}

	for	i := 0; i <= 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d \n", numero1, i, numero1*i)
	}

	fmt.Println(texto)
	return texto
}