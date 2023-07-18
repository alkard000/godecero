package ejercicios

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

var numero1 int
var err error

func InputData() {
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
		fmt.Println(i*numero1)
	}
}