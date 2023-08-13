package NT

import (
	"Lab2/Interprete"
	"strconv"
)

type NT_Suma struct {
	ExprIzq Interprete.AbstrExpr
	ExprDer Interprete.AbstrExpr
}

type NT_Mult struct {
	ExprIzq Interprete.AbstrExpr
	ExprDer Interprete.AbstrExpr
}

type NT_ID struct {
	ID      string
	linea   int
	columna int
}

func (N_ID NT_ID) Interpretar(ctx *Interprete.Contexto) *Interprete.Resultado {
	expr, ok := ctx.GetVariable(N_ID.ID)
	if !ok {
		ctx.AddError("Error, la variable: " + N_ID.ID +
			"no existe en la linea: " + strconv.Itoa(N_ID.linea) +
			"en la columna: " + strconv.Itoa(N_ID.columna))
		return Interprete.NewNil()
	}
	return expr
}

func NewNT_ID(ID string, linea int, columna int) *NT_ID {
	return &NT_ID{ID: ID, linea: linea, columna: columna}
}

func (N NT_Mult) Interpretar(ctx *Interprete.Contexto) *Interprete.Resultado {
	//ejecutamos los nodos hijos
	resultadoIzq := N.ExprIzq.Interpretar(ctx)
	resultadoDer := N.ExprDer.Interpretar(ctx)
	//TODO verificación de tipos
	return Interprete.NewResultado(resultadoIzq.Valor*resultadoDer.Valor,
		resultadoIzq.ValorS+resultadoDer.ValorS)

}

func NewNT_Mult(exprIzq Interprete.AbstrExpr, exprDer Interprete.AbstrExpr) *NT_Mult {
	return &NT_Mult{ExprIzq: exprIzq, ExprDer: exprDer}
}

func (N *NT_Suma) Interpretar(ctx *Interprete.Contexto) *Interprete.Resultado {
	// ejecutamos los nodos hijos
	resultadoIzq := N.ExprIzq.Interpretar(ctx)
	resultadoDer := N.ExprDer.Interpretar(ctx)

	//TODO verificar tipos
	//ooperamos lo resultados
	return Interprete.NewResultado(resultadoIzq.Valor+resultadoDer.Valor,
		resultadoIzq.ValorS+resultadoDer.ValorS)
}

func NewNT_Suma(ExprIzq Interprete.AbstrExpr, ExprDer Interprete.AbstrExpr) *NT_Suma {
	return &NT_Suma{
		ExprDer: ExprDer,
		ExprIzq: ExprIzq,
	}
}
