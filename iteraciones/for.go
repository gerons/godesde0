package iteraciones

import (
	"fmt"
)

func Iterar() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func TipoWhile() {
	i := 0

	for i < 10 {
		fmt.Println(i)

		i++
	}
}

func TipoWhile2() {
	i := 0
	sigue := true

	for sigue {
		fmt.Println(i)
		i++

		if i == 100 {
			sigue = false
		}
	}
}