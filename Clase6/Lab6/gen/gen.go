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
