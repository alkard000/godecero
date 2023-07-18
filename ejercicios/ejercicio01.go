package ejercicios

import "strconv"

func TwoValues(valor string) (int, string) {
	if formated, _ := strconv.Atoi(valor); formated > 100{
		return formated, "Es mayor a 100"
	} else {
		return formated, "Es menor a 100"
	}
}