package files

import (
	//"bufio"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/yosan-aly/e-comerce/ejercicios"
)

 var fileName string = "./files/txt/tabla.txt"

func SaveTable() {
	var texto string = ejercicios.TablaNumerica()
	archivo, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()

}


func SumarTable() {
	var texto string = ejercicios.TablaNumerica()
	if !Append(fileName, texto) {
		fmt.Println("No se pudo agregar al archivo")
		return
	}
}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(filen, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el append " + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el WriteStrig " + err.Error())
		return false
	}
	arch.Close()
	return true
}

func ReadFileioutil() {
	archivo, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}

	fmt.Println(string(archivo))
}

func ReadFile() {
	archivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error al leer el archivo:", err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}
	archivo.Close()
}
