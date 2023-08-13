package ejercicios

import (
	"strconv"
	"log"
)

func DevuelveEnteros(val string) (int, string) {
	numero, err := strconv.Atoi(val)

	if err != nil {
		log.Fatal(err)
	}

	if (numero > 100) {
		return numero, "es mayor a 100"
	} else {
		return numero, "es menor a 100"
	}
}