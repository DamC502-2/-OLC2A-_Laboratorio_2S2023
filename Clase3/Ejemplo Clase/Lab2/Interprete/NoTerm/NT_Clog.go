package NT

import (
	interprete "Lab2/Interprete"
	"strconv"
)

type NT_Clog struct {
	ListaExpr []interprete.AbstrExpr
}

func (ntc *NT_Clog) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	for _, expr := range ntc.ListaExpr {
		exprV := expr.Interpretar(ctx)
		ctx.Print(strconv.Itoa(exprV.Valor))
	}
	ctx.Println("")
	return interprete.NewNil()
}

// añado expresiones a la lista
func (ntc *NT_Clog) AddExpresion(expr interprete.AbstrExpr) {
	ntc.ListaExpr = append(ntc.ListaExpr, expr)
}

// contructuro de go:
func NewNT_Clog() *NT_Clog {
	return &NT_Clog{
		//inicializo la lista
		ListaExpr: make([]interprete.AbstrExpr, 0, 5),
	}
}

// constructor 2
func NewNT_Clog2() *NT_Clog {
	var ntc = new(NT_Clog)
	//inicializando la lista
	ntc.ListaExpr = make([]interprete.AbstrExpr, 0, 5)
	return ntc
}
