package files

import (
	"godecero/ejercicios"
	"fmt"
	"os"
	"bufio"
)

var filename string = "./files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.InputData()

	archivo, err := os.Create(filename)

	if err!= nil {
		fmt.Println("Error al crear el archivo"+err.Error())
		return
	}

	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.InputData()
	if !Append(filename, texto) {
		fmt.Println("Error al escribir en el archivo")
	}
}

func Append(filen string, texto string) bool {
	archivo, err := os.OpenFile(filen, os.O_APPEND|os.O_WRONLY, 0644)

	if err!= nil {
        fmt.Println("Error en el append"+err.Error())
        return false
    }

	_, err = archivo.WriteString(texto)

	if err!= nil {
        fmt.Println("Error en el Writestring"+err.Error())
        return false
    }

    archivo.Close()
    return true
}

func LeerTabla() {
	archivo, err := os.Open(filename)

    if err!= nil {
        fmt.Println("Error al abrir el archivo"+err.Error())
        return
    }

    scanner := bufio.NewScanner(archivo)

    for scanner.Scan() {
        registro := scanner.Text()
		fmt.Println("> ", registro)
    }

    archivo.Close()
}