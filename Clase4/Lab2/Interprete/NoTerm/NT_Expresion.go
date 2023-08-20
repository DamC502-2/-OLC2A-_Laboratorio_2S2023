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
	return &Interprete.Resultado{
		Tipo:  Interprete.Integer,
		Valor: resultadoIzq.Valor * resultadoDer.Valor,
	}

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
	//return Interprete.NewResultado(resultadoIzq.Valor+resultadoDer.Valor,
	//resultadoIzq.ValorS+resultadoDer.ValorS)
	// sigue la ponderación de los tipos

	if resultadoIzq.Tipo != resultadoDer.Tipo {
		if resultadoIzq.Tipo > resultadoDer.Tipo {
			resultadoDer = ctx.Conversor.Ampliar(resultadoDer, resultadoIzq.Tipo)
		} else {
			resultadoIzq = ctx.Conversor.Ampliar(resultadoIzq, resultadoDer.Tipo)
		}
	}

	if resultadoDer.Nil || resultadoIzq.Nil {
		ctx.AddError("No se puede realizar una suma de tipos incompatibles")
		return Interprete.NewNil()
	}

	switch resultadoIzq.Tipo {
	case Interprete.Bool:
		ctx.AddError("No se pueden sumar expresiones de tipo Bool")
		return Interprete.NewNil()
	case Interprete.Integer:
		return Interprete.NewIntLiteral(resultadoDer.Valor + resultadoDer.Valor)
	case Interprete.Float:
		//return Interprete.NewFloatLiteral(resultadoIzq.ValorF + resultadoDer.ValorF )
		return &Interprete.Resultado{
			ValorF: resultadoDer.ValorF + resultadoIzq.ValorF,
			Tipo:   Interprete.Float,
		}
	case Interprete.String:
		return Interprete.NewStringLiteral(resultadoIzq.ValorS + resultadoDer.ValorS)
	default:
		ctx.AddError("Error de tipos no valido" + resultadoIzq.Tipo.String() +
			" y " + resultadoDer.Tipo.String())
	}
	return Interprete.NewNil()

}

func NewNT_Suma(ExprIzq Interprete.AbstrExpr, ExprDer Interprete.AbstrExpr) *NT_Suma {
	return &NT_Suma{
		ExprDer: ExprDer,
		ExprIzq: ExprIzq,
	}
}
