package NT

import "Lab2/Interprete"

//TODO añadir tipos
type NT_DeclVar struct {
	ID        string
	Tipo      string
	Expresion Interprete.AbstrExpr
	linea     int
	columan   int
}

func NewNT_DeclVar(ID string, tipo string, expresion Interprete.AbstrExpr, linea int, columan int) *NT_DeclVar {
	return &NT_DeclVar{ID: ID, Tipo: tipo, Expresion: expresion, linea: linea, columan: columan}
}

func (NT_DV *NT_DeclVar) Interpretar(ctx *Interprete.Contexto) *Interprete.Resultado {
	var expr = Interprete.NewNil()
	if NT_DV.Expresion != nil {
		expr = NT_DV.Expresion.Interpretar(ctx)
	}
	ctx.AddVariable(NT_DV.ID, expr, NT_DV.linea, NT_DV.columan)
	return Interprete.NewNil()
}
