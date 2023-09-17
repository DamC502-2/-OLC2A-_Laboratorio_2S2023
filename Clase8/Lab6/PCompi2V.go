package main

import (
	"Lab6/Compilador"
	"Lab6/Compilador/NT"
	"Lab6/Compilador/Terminal"
	"Lab6/PCompi2N"
	"github.com/antlr4-go/antlr/v4"
)

type PCompi2V struct {
	antlr.ParseTreeVisitor
}

func NewPCompi2V() PCompi2N.PCompi2NVisitor {
	return &PCompi2V{
		ParseTreeVisitor: &PCompi2N.BasePCompi2NVisitor{},
	}
}

/// funcion visit generica para iniciar el análisis
func (P PCompi2V) Visit(tree antlr.ParseTree) interface{} {
	nodo := tree.Accept(P)
	return nodo
}

func (P PCompi2V) VisitS(ctx *PCompi2N.SContext) interface{} {
	//TODO implement me
	nodo := ctx.Cond().Accept(P).(Compilador.CAbstrExpr)
	return NT.NewNT_s(nodo)
}

func (P PCompi2V) VisitCondOr(ctx *PCompi2N.CondOrContext) interface{} {

	nodoIzq := ctx.GetC1().Accept(P).(Compilador.CAbstrExpr)
	nodoDer := ctx.GetC2().Accept(P).(Compilador.CAbstrExpr)
	return NT.NewNT_CondOr(nodoIzq, nodoDer)
}

func (P PCompi2V) VisitCondAnd(ctx *PCompi2N.CondAndContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (P PCompi2V) VisitCondI(ctx *PCompi2N.CondIContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (P PCompi2V) VisitCondBool(ctx *PCompi2N.CondBoolContext) interface{} {

	return Terminal.NewT_Bool(ctx.GetBool_().GetText())
}

func (P PCompi2V) VisitCondPar(ctx *PCompi2N.CondParContext) interface{} {
	//TODO implement me
	return ctx.Cond().Accept(P).(Compilador.CAbstrExpr)
}

func (P PCompi2V) VisitCondOper(ctx *PCompi2N.CondOperContext) interface{} {
	//TODO implement me
	nodoIzq := ctx.GetE1().Accept(P).(Compilador.CAbstrExpr)
	nodoDer := ctx.GetE2().Accept(P).(Compilador.CAbstrExpr)

	return NT.NewNT_CondOper(nodoIzq, nodoDer, ctx.GetOpr().GetText())

}

func (P PCompi2V) VisitEnum(ctx *PCompi2N.EnumContext) interface{} {
	//TODO implement me
	return Terminal.NewT_Numero(ctx.GetN().GetText())
}

func (P PCompi2V) VisitEid(ctx *PCompi2N.EidContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (P PCompi2V) VisitEpa(ctx *PCompi2N.EpaContext) interface{} {
	//TODO implement me
	return ctx.GetE1().Accept(P).(Compilador.CAbstrExpr)
}

func (P PCompi2V) VisitENeg(ctx *PCompi2N.ENegContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (P PCompi2V) VisitEmd(ctx *PCompi2N.EmdContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (P PCompi2V) VisitEsr(ctx *PCompi2N.EsrContext) interface{} {
	nodoIzq := ctx.GetE1().Accept(P).(Compilador.CAbstrExpr)
	nodoDer := ctx.GetE2().Accept(P).(Compilador.CAbstrExpr)

	return NT.NewNT_Esr(nodoIzq, nodoDer, ctx.GetOp().GetText())

}
