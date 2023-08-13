package main

import (
	"fmt"
	"godesde0/ejercicios"
	"godesde0/variables"
	"runtime"
)

func main () {
	estado, texto := variables.ConviertoaTexto(1588)

	fmt.Println(estado)
	fmt.Println(texto)

	if os := runtime.GOOS; os == "darwin" {
		fmt.Println("esto no es Windows")
	} else {
		fmt.Println("es widows!!")
	}

	switch os := runtime.GOOS; os {
		case "linux":
			fmt.Println("Linux")
		case "darwin":
			fmt.Println("Mac")
		default:
			fmt.Printf("%s \n", os) 
	}

	num, descripcion := ejercicios.DevuelveEnteros("150")
	fmt.Printf("el número %d es %s", num, descripcion)
}