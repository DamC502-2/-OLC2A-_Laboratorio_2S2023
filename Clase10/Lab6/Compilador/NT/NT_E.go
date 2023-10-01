package NT

import "Lab6/Compilador"

type NT_Esr struct {
	ExprIzq Compilador.CAbstrExpr
	ExprDer Compilador.CAbstrExpr
	oper    string
}

func (N NT_Esr) Compilar(ctx *Compilador.Contexto) *Compilador.Atributos {
	exprIzq := N.ExprIzq.Compilar(ctx)
	exprDer := N.ExprDer.Compilar(ctx)
	/// validacion de tipos
	temp := ctx.NewTemp()
	ctx.Gen( temp + " = " + exprIzq.Dir +" " + N.oper + " " + exprDer.Dir)
	return &Compilador.Atributos{
		EV:  make([]string, 0,0),
		EF:  make([]string, 0,0),
		Dir: temp,
	}
}

func NewNT_Esr(exprIzq Compilador.CAbstrExpr, exprDer Compilador.CAbstrExpr, oper string) *NT_Esr {
	return &NT_Esr{ExprIzq: exprIzq, ExprDer: exprDer, oper: oper}
}

type NT_Emd struct {
	ExprIzq Compilador.CAbstrExpr
	ExprDer Compilador.CAbstrExpr
	oper    string
}

func NewNT_Emd(exprIzq Compilador.CAbstrExpr, exprDer Compilador.CAbstrExpr, oper string) *NT_Emd {
	return &NT_Emd{ExprIzq: exprIzq, ExprDer: exprDer, oper: oper}
}

type NT_ENeg struct {
	ExprDer Compilador.CAbstrExpr
}

func NewNT_ENeg(exprDer Compilador.CAbstrExpr) *NT_ENeg {
	return &NT_ENeg{ExprDer: exprDer}
}
