package files

import (
	"godesde0/ejercicios"
	"os"
	// "bufio"
	// "ioutil"
	"fmt"
)

var filneName string = "./files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.TabladeMultiplicar()

	archivo, err := os.Create(filneName)

	if err != nil {
		fmt.Println("Error al crear el archivo"+err.Error())
		return
	}

	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.TabladeMultiplicar()
	if !Append(filneName, texto) {
		fmt.Println("Erro al concatenar contenido")
	}
}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(filen, os.O_WRONLY | os.O_APPEND, 0655)

	if err != nil {
		fmt.Println("Error durante el append"+err.Error())
		return false
	}

	_, err = arch.WriteString(texto)

	arch.Close()
	return true
}