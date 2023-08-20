package T

import (
	"Lab2/Interprete"
	"strconv"
)

type T_Num struct {
	Num     string
	linea   int
	columna int
}

func (T_N T_Num) Interpretar(ctx *Interprete.Contexto) *Interprete.Resultado {
	num, _ := strconv.Atoi(T_N.Num)
	return Interprete.NewIntLiteral(num)
}

func NewT_Num(num string, linea int, columna int) *T_Num {
	return &T_Num{Num: num, linea: linea, columna: columna}
}
