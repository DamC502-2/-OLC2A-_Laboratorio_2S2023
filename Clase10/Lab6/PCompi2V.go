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

func (P PCompi2V) VisitSentenciaSwitch(ctx *PCompi2N.SentenciaSwitchContext) interface{} {
	//TODO implement me
	return ctx.Switch_().Accept(P).(Compilador.CAbstrExpr)
}

//switch: 'switch' expr '{'  case* defaultc?    '}' ;
func (P PCompi2V) VisitSwitch(ctx *PCompi2N.SwitchContext) interface{} {
	//TODO implement me
	exprN := ctx.Expr().Accept(P).(Compilador.CAbstrExpr)
	var defaultN Compilador.CAbstrExpr = nil
	if ctx.Defaultc() != nil {
		defaultN = ctx.Defaultc().Accept(P).(Compilador.CAbstrExpr)
	}
	var listaCaseN = make([]Compilador.CAbstrExpr, 0, 0)

	listaCase := ctx.AllCase_()
	for _, caseV := range listaCase {
		listaCaseN = append(listaCaseN,
			//visitor para cada case
			caseV.Accept(P).(Compilador.CAbstrExpr))
	}

	nodoSwitch := NT.NewNT_Switch(exprN)
	nodoSwitch.Deafultc = defaultN
	nodoSwitch.Cases = listaCaseN
	return nodoSwitch
}

func (P PCompi2V) VisitCase(ctx *PCompi2N.CaseContext) interface{} {
	//TODO implement me
	exprN := ctx.Expr().Accept(P).(Compilador.CAbstrExpr)
	block := ctx.BlockS().Accept(P).(Compilador.CAbstrExpr)
	return NT.NewNT_Case(exprN, block)
}

func (P PCompi2V) VisitDefaultc(ctx *PCompi2N.DefaultcContext) interface{} {
	//TODO implement me
	return ctx.BlockS().Accept(P).(Compilador.CAbstrExpr)
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

/*
s:
    c1=cond c2=cond c3=cond EOF
    | if
    | blockS
    | sentencia
    ;

*/
// S -> cond
func (P PCompi2V) VisitS(ctx *PCompi2N.SContext) interface{} {
	//TODO implement me
	//if ctx.Sentencia() != nil {
	//	//ctx.
	//}
	//conds := ctx.AllCond()
	//ctx.GetC1()
	//ctx.GetC2()
	//ctx.GetC3()
	if ctx.If_() != nil {
		nodo := ctx.If_().Accept(P).(Compilador.CAbstrExpr)
		return NT.NewNT_s(nodo)
	}
	if ctx.Cond() != nil {
		return NT.NewNT_s(
			ctx.Cond().Accept(P).(Compilador.CAbstrExpr),
		)
	}
	if ctx.Switch_() != nil {
		return NT.NewNT_s(
			ctx.Switch_().Accept(P).(Compilador.CAbstrExpr),
		)
	}

	return nil

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

func (P PCompi2V) VisitIFsen(ctx *PCompi2N.IFsenContext) interface{} {
	//TODO implement me

	var bloque1 Compilador.CAbstrExpr
	var bloqueFalso Compilador.CAbstrExpr = nil

	cond := ctx.Cond().Accept(P).(Compilador.CAbstrExpr)

	listaBloques := ctx.AllBlockS()
	if listaBloques != nil {
		if len(listaBloques) == 2 {
			bloque1 = ctx.BlockS(0).Accept(P).(Compilador.CAbstrExpr)
			bloqueFalso = ctx.BlockS(1).Accept(P).(Compilador.CAbstrExpr)
		} else {
			// viene un bloque
			bloque1 = ctx.BlockS(0).Accept(P).(Compilador.CAbstrExpr)
		}
	}
	if ctx.If_() != nil {
		bloqueFalso = ctx.If_().Accept(P).(Compilador.CAbstrExpr)
	}
	return NT.NewNT_If(cond, bloque1, bloqueFalso)
}

func (P PCompi2V) VisitBlock(ctx *PCompi2N.BlockContext) interface{} {
	//TODO implement me

	if ctx.Sentencia() != nil {
		setencias := ctx.Sentencia().Accept(P).(Compilador.CAbstrExpr)
		return setencias

	} else {
		return NT.NewNT_Lsentencias()
	}

}

func (P PCompi2V) VisitSentenciaEj(ctx *PCompi2N.SentenciaEjContext) interface{} {
	//TODO implement me
	setencias := ctx.GetChildCount()
	nodoSentencias := NT.NewNT_Lsentencias()
	for i := 0; i < setencias; i++ {
		nodoSentencias.Sentencias =
			append(nodoSentencias.Sentencias, NT.NewNT_Sentencia())
	}
	return nodoSentencias
}
