package T

import (
	interprete "Lab2/Interprete"
	"strconv"
)

type T_Num struct {
	Num     string
	linea   int
	columna int
}

func (T_N *T_Num) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	num, _ := strconv.Atoi(T_N.Num)
	return interprete.NewResultado(num, T_N.Num)
}

// Constructor for T_Num
func NewT_Num(Num string, linea int, columna int) *T_Num {
	o := new(T_Num)
	o.Num = Num
	o.linea = linea
	o.columna = columna
	return o
}
