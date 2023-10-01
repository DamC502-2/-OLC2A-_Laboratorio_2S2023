package NT

import (
	"Lab6/Compilador"
)

type NT_CondOper struct {
	ExprIzq Compilador.CAbstrExpr
	ExprDer Compilador.CAbstrExpr
	oper    string
}

type NT_CondOr struct {
	CondIzq Compilador.CAbstrExpr
	CondDer Compilador.CAbstrExpr
}

/// a < b OR d == e
func (N NT_CondOr) Compilar(ctx *Compilador.Contexto) *Compilador.Atributos {

	//exprDer := N.CondDer.Compilar(ctx)
	//exprIzq := N.CondIzq.Compilar(ctx)
	//exprIzq3 := N.CondIzq.Compilar(ctx)
	/* salida
		 if d == e goto L1
		goto L2
	//
		if a > b goto L3
		goto L4
	//
		if a > b goto L5
			goto L6
	*/

	condIzq := N.CondIzq.Compilar(ctx)
	/// marcador==========
	ctx.ImprimirEtiquetas(condIzq.EF)
	/// marcador=======
	condDer := N.CondDer.Compilar(ctx)

	return &Compilador.Atributos{
		// unimos etiquetas
		EV: ctx.Unir(condIzq.EV, condDer.EV),
		// sintetizamos EF de cond2
		EF:  condDer.EF,
		Dir: "",
	}
}

func NewNT_CondOr(condIzq Compilador.CAbstrExpr, condDer Compilador.CAbstrExpr) *NT_CondOr {
	return &NT_CondOr{CondIzq: condIzq, CondDer: condDer}
}

func (N NT_CondOper) Compilar(ctx *Compilador.Contexto) *Compilador.Atributos {
	exprIzq := N.ExprIzq.Compilar(ctx)
	exprDer := N.ExprDer.Compilar(ctx)

	EV := ctx.NewEtq()
	EF := ctx.NewEtq()
	EVs := make([]string, 0, 0)
	EFs := make([]string, 0, 0)
	EVs = append(EVs, EV)
	EFs = append(EFs, EF)

	ctx.Gen("if " + exprIzq.Dir + " " + N.oper + " " +
		exprDer.Dir + " then goto " + EV)
	ctx.Gen("goto " + EF)

	return &Compilador.Atributos{
		EV:  EVs,
		EF:  EFs,
		Dir: "",
	}
}

func NewNT_CondOper(exprIzq Compilador.CAbstrExpr, exprDer Compilador.CAbstrExpr, oper string) *NT_CondOper {
	return &NT_CondOper{ExprIzq: exprIzq, ExprDer: exprDer, oper: oper}
}
