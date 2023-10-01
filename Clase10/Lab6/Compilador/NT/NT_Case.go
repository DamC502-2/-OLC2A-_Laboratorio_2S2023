package NT

import "Lab6/Compilador"

type NT_Case struct {
	Expr   Compilador.CAbstrExpr
	Bloque Compilador.CAbstrExpr
}

func (N NT_Case) Compilar(ctx *Compilador.Contexto) *Compilador.Atributos {
	/*
		<cod sentencias c3d>
		<goto Lsalida>
	*/
	N.Bloque.Compilar(ctx)
	Lsalida := ctx.PeekDislay()
	ctx.Gen("goto " + Lsalida + "//goto Lsalida case ")
	return &Compilador.Atributos{
		EV:  make([]string, 0, 0),
		EF:  make([]string, 0, 0),
		Dir: "",
	}
}

func NewNT_Case(expr Compilador.CAbstrExpr, blockque Compilador.CAbstrExpr) *NT_Case {
	return &NT_Case{Expr: expr, Bloque: blockque}
}
