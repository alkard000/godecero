package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "BUenos Aires"

	fmt.Println(paises)
	fmt.Println(paises["Mexico"])

	campeonato := map[string]int{
		"Barcelona": 39,
		"Real Madrid": 40,
		"Chivasco": 41,
		"Atletico Madrid": 42,
        "Sevilla": 43,
	}

	fmt.Println(campeonato)

	/*for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo %s, tiene un puntaje de %d\n", equipo, puntaje)
	}*/

	delete(campeonato, "Atletico Madrid")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Atletico Madrid"]

	fmt.Printf("El equipo Atletico Madrid tiene un puntaje de %d, y el equipo existe es %t\n", puntaje, existe)
}