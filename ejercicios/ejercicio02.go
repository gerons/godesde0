package ejercicios

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func TabladeMultiplicar() string {
	var numero int
	var err error
	var texto string

	scanner := bufio.NewScanner((os.Stdin))

	for {
		fmt.Println("Ingrese un número:")
		
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())

			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i < 10; i++ {
		texto += fmt.Sprintf( "%d x %d = %d \n", numero, i, numero*i )
	}

	return texto
}