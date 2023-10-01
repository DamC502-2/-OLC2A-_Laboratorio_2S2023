package NT

import "Lab6/Compilador"

type NT_s struct {
	nodo Compilador.CAbstrExpr
}

func (N NT_s) Compilar(ctx *Compilador.Contexto) *Compilador.Atributos {
	atributos := N.nodo.Compilar(ctx)
	ctx.Gen("//etiquetas verdaderas")
	ctx.ImprimirEtiquetas(atributos.EV)
	ctx.Gen("//etiquetas falsas")
	ctx.ImprimirEtiquetas(atributos.EF)
	return atributos
}

func NewNT_s(nodo Compilador.CAbstrExpr) *NT_s {
	return &NT_s{nodo: nodo}
}
