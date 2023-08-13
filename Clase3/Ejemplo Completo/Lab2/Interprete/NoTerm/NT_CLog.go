package NT

import (
	"Lab2/Interprete"
	Inteprete "Lab2/Interprete/Terminales"
	"strconv"
)

type NT_Clog struct {
	ListaExpr []Interprete.AbstrExpr
}

func (Clog *NT_Clog) Interpretar(ctx *Interprete.Contexto) *Interprete.Resultado {
	for _, expr := range Clog.ListaExpr {
		switch expr.(type) {
		case *Inteprete.T_Cad:
			//si el nodo es de tipo cadena se imprime la cadena (casteando el nodo)
			ctx.Print(expr.(*Inteprete.T_Cad).Valor)
			break
		default:
			//se imprime todos los numeros
			ctx.Print(strconv.Itoa(expr.Interpretar(ctx).Valor))
		}
	}
	ctx.Println("")
	return Interprete.NewNil()
}

func (clog *NT_Clog) AddExpresion(expr Interprete.AbstrExpr) {
	clog.ListaExpr = append(clog.ListaExpr, expr)
}

func NewNT_Clog() *NT_Clog {
	return &NT_Clog{
		ListaExpr: make([]Interprete.AbstrExpr, 0, 5),
	}
}
