package NT

import (
	"Lab2/Interprete"
	"strconv"
)

type NT_AsigVar struct {
	ID      string
	Expr    Interprete.AbstrExpr
	linea   int
	columna int
}

func (AVar *NT_AsigVar) Interpretar(ctx *Interprete.Contexto) *Interprete.Resultado {
	expr := AVar.Expr.Interpretar(ctx) // ejecutamos la expresión
	// obtenemos la variable de la memoria
	_, ok := ctx.GetVariable(AVar.ID)

	//si no exite la variable debe ser un error
	if !ok {
		ctx.AddError("No existe la variable: " + AVar.ID +
			" en la linea: " + strconv.Itoa(AVar.linea) +
			" en la columan" + strconv.Itoa(AVar.columna))
		return Interprete.NewNil()
	}
	//TODO verificacion de tipos
	//asignamos el nuevo valor
	ctx.AsigVariable(AVar.ID, expr)
	return Interprete.NewNil()
}

func NewNT_AsigVar(ID string, expr Interprete.AbstrExpr, linea int, columna int) *NT_AsigVar {
	return &NT_AsigVar{ID: ID, Expr: expr, linea: linea, columna: columna}
}
