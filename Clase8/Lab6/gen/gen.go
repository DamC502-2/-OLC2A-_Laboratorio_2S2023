package gen

import (
	"fmt"
	"strconv"
)

// generadores de temporales
// impresion del código

var tmp = 0
var lb = 0

func Gen(out string) {
	fmt.Println(out)
}

func NewTemp() string {
	tmp = tmp + 1
	return "t" + strconv.Itoa(tmp)
}

func NewEtq() string {
	lb = lb + 1
	return "L" + strconv.Itoa(lb)
}

func ImprimirEtiquetas(etiquetas []string) {
	for _, etiqueta := range etiquetas {
		Gen(etiqueta + ":")
	}
}

func Unir(etq1 []string, etq2 []string) []string {
	for _, etq := range etq2 {
		etq1 = append(etq1, etq)
	}
	return etq1
}
