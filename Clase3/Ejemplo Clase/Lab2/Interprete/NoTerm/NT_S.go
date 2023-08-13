package NT

import interprete "Lab2/Interprete"

func (NTs *NT_S) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	return NTs.sentencias.Interpretar(ctx)
}

type NT_S struct {
	sentencias interprete.AbstrExpr
}

// Constructor for NT_S
func NewNT_S(sentencias interprete.AbstrExpr) *NT_S {
	o := new(NT_S)
	o.sentencias = sentencias
	return o
}
