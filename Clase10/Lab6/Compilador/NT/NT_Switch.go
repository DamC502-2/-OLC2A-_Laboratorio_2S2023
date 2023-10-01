package NT

import "Lab6/Compilador"

type NT_Switch struct {
	Expr     Compilador.CAbstrExpr
	Cases    []Compilador.CAbstrExpr
	Deafultc Compilador.CAbstrExpr
}

func (N NT_Switch) Compilar(ctx *Compilador.Contexto) *Compilador.Atributos {
	// Cod ed E
	expr := N.Expr.Compilar(ctx)
	Lprueba := ctx.NewEtq()
	ctx.Gen("goto " + Lprueba + "// goto Lprueba")

	// agrego una etiqueta al display
	LSalida := ctx.PushDisplay()
	// lista de expresiones
	listaExpr := make([]Compilador.CAbstrExpr, 0, 0)
	listaEtq := make([]string, 0, 0)
	for _, caseN := range N.Cases {
		caseNT := caseN.(*NT_Case)
		listaExpr = append(listaExpr, caseNT.Expr)
		etq := ctx.NewEtq()
		listaEtq = append(listaEtq, etq)
		ctx.Gen(etq + ": //case ====") //
		caseN.Compilar(ctx)
	}
	var Ln string = ""
	if N.Deafultc != nil {
		Ln = ctx.NewEtq()
		ctx.Gen(Ln + ": //Ln - Default")
		N.Deafultc.Compilar(ctx)
		//goto Lsalida
		ctx.Gen("goto " + LSalida + ":" + "// goto Lsalida")
	}

	/// Las pruebas
	ctx.Gen(Lprueba + ": // Lprueba ")

	for i, exprC := range listaExpr {
		//TODO manejo de tipos
		atrib := exprC.Compilar(ctx)
		ctx.Gen("if " + expr.Dir + " == " +
			atrib.Dir + " " + listaEtq[i])

	}
	if N.Deafultc != nil {
		ctx.Gen("goto" + Ln + " // Ln")
	}

	ctx.Gen(LSalida + ":")
	ctx.PopDisplay()
	return &Compilador.Atributos{
		EV:  make([]string, 0, 0),
		EF:  make([]string, 0, 0),
		Dir: "",
	}
}

func NewNT_Switch(expr Compilador.CAbstrExpr) *NT_Switch {
	return &NT_Switch{
		Expr:     expr,
		Cases:    make([]Compilador.CAbstrExpr, 0, 0),
		Deafultc: nil}
}
