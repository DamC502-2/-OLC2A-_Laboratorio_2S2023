package NT

import (
	"Lab2/Interprete"
)

type NT_Clog struct {
	ListaExpr []Interprete.AbstrExpr
}

func (Clog *NT_Clog) Interpretar(ctx *Interprete.Contexto) *Interprete.Resultado {
	var expresion *Interprete.Resultado
	for _, expr := range Clog.ListaExpr {
		/*	switch expr.(type) {
			case *Inteprete.T_Cad:
				//si el nodo es de tipo cadena se imprime la cadena (casteando el nodo)
				ctx.Print(expr.(*Inteprete.T_Cad).Valor)
				break
			default:
				//se imprime todos los numeros
				ctx.Print(strconv.Itoa(expr.Interpretar(ctx).Valor))
			}*/
		expresion = expr.Interpretar(ctx)
		switch expresion.Tipo {
		case Interprete.Nil:
			ctx.Print("NIL")
		case Interprete.Bool:
			expresion = ctx.Conversor.Ampliar(expresion, Interprete.String)
			ctx.Print(expresion.ValorS)
		case Interprete.Integer:
			expresion = ctx.Conversor.Ampliar(expresion, Interprete.String)
			ctx.Print(expresion.ValorS)
		case Interprete.Float:
			expresion = ctx.Conversor.Ampliar(expresion, Interprete.String)
			ctx.Print(expresion.ValorS)
		case Interprete.String:
			ctx.Print(expresion.ValorS)
		default:
			ctx.AddError("tipo no esperado en la consola")
			ctx.Print("ERROR!")
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
