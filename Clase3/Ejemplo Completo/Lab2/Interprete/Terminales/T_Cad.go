package T

import (
	"Lab2/Interprete"
)

type T_Cad struct {
	Valor   string
	linea   int
	columan int
}

func (tcad *T_Cad) Interpretar(ctx *Interprete.Contexto) *Interprete.Resultado {
	return Interprete.NewResultado(0, tcad.Valor)

}

func NewT_Cad(valor string, linea int, columan int) *T_Cad {
	return &T_Cad{Valor: valor, linea: linea, columan: columan}
}
