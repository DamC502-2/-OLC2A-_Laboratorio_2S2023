package NT

import interprete "Lab2/Interprete"

type NT_Suma struct {
	ExprIzq interprete.AbstrExpr
	ExprDer interprete.AbstrExpr
}

func (ntsum *NT_Suma) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	//TODO manejo de tipos
	exprIzq := ntsum.ExprIzq.Interpretar(ctx)
	exprDer := ntsum.ExprDer.Interpretar(ctx)

	resultado := interprete.NewResultado(
		exprIzq.Valor+exprDer.Valor,
		"pediente")
	return resultado
}

// Constructor for NT_Suma
func NewNT_Suma(exprIzq interprete.AbstrExpr,
	exprDer interprete.AbstrExpr) *NT_Suma {
	o := new(NT_Suma)
	o.ExprIzq = exprIzq
	o.ExprDer = exprDer
	return o
}
