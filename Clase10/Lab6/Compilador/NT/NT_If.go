package NT

import "Lab6/Compilador"

type NT_If struct {
	Cond        Compilador.CAbstrExpr
	Bloque      Compilador.CAbstrExpr
	BloqueFalso Compilador.CAbstrExpr
}

func (N NT_If) Compilar(ctx *Compilador.Contexto) *Compilador.Atributos {
	// el if
	ctx.Gen("// codigo cond")
	atributos := N.Cond.Compilar(ctx) // se va a generar el código
	ctx.Gen("// EV cond")
	Lsalida := ctx.NewEtq()
	ctx.ImprimirEtiquetas(atributos.EV)
	ctx.Gen("// codigo de bloque")
	N.Bloque.Compilar(ctx)
	ctx.Gen("goto " + Lsalida + ":")
	ctx.ImprimirEtiquetas(atributos.EF)
	if N.BloqueFalso != nil {
		N.BloqueFalso.Compilar(ctx)
	}
	ctx.Gen(Lsalida + ":")
	return &Compilador.Atributos{
		EV:  make([]string, 0),
		EF:  make([]string, 0),
		Dir: "",
	}

}

func NewNT_If(cond Compilador.CAbstrExpr, bloque Compilador.CAbstrExpr, bloqueFalso Compilador.CAbstrExpr) *NT_If {
	return &NT_If{Cond: cond, Bloque: bloque, BloqueFalso: bloqueFalso}
}
