package main

import (
	interprete "Lab2/Interprete"
	NT "Lab2/Interprete/NoTerm"
	T "Lab2/Interprete/Terminales"
	"Lab2/Lab2G"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type VisitorInterprete struct {
	antlr.ParseTreeVisitor
	Raiz interprete.AbstrExpr
}

func NewVisitorInterpreter() Lab2G.Lab2_GVisitor {
	return &VisitorInterprete{
		ParseTreeVisitor: &Lab2G.BaseLab2_GVisitor{},
		Raiz:             nil,
	}
}

// ctrl + shift + p
// Visit a parse tree produced by Lab2_GParser#SWhile.
func (vI *VisitorInterprete) VisitSWhile(ctx *Lab2G.SWhileContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// funcion para manejar errores
func (vI *VisitorInterprete) Visit(tree antlr.ParseTree) interface{} {
	nodoInterprete, ok := tree.Accept(vI).(interprete.AbstrExpr)
	if !ok {
		fmt.Println("error")
		return nil //devolver un error
	}
	return nodoInterprete
}

func (vI *VisitorInterprete) VisitChildren(node antlr.RuleNode) interface{} {
	panic("not implemented") // TODO: Implement
}

func (vI *VisitorInterprete) VisitTerminal(node antlr.TerminalNode) interface{} {
	panic("not implemented") // TODO: Implement
}

func (vI *VisitorInterprete) VisitErrorNode(node antlr.ErrorNode) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by Lab2_GParser#SLSentencias.
func (vI *VisitorInterprete) VisitSLSentencias(ctx *Lab2G.SLSentenciasContext) interface{} {
	lsetencias := ctx.L_sentencias().Accept(vI).(interprete.AbstrExpr)
	nodoS := NT.NewNT_S(lsetencias)
	vI.Raiz = nodoS
	return nodoS
}

// Visit a parse tree produced by Lab2_GParser#L_Sentencia.
func (vI *VisitorInterprete) VisitL_Sentencia(ctx *Lab2G.L_SentenciaContext) interface{} {
	lSentencias := NT.NewLSentencias()
	sentenciasAntlr := ctx.AllSentencia()
	for _, sentencia := range sentenciasAntlr {
		nodoSentencia := sentencia.Accept(vI).(interprete.AbstrExpr)
		lSentencias.AddSentencia(nodoSentencia)
	}
	return lSentencias

}

// Visit a parse tree produced by Lab2_GParser#SConsola.
func (vI *VisitorInterprete) VisitSConsola(ctx *Lab2G.SConsolaContext) interface{} {
	consola := NT.NewNT_Clog()
	lExpresiones := ctx.AllE()

	for _, expr := range lExpresiones {
		nodoExpr := expr.Accept(vI).(interprete.AbstrExpr)
		consola.AddExpresion(nodoExpr)
	}
	return consola
}

// Visit a parse tree produced by Lab2_GParser#SDeclaracion.
func (vI *VisitorInterprete) VisitSDeclaracion(ctx *Lab2G.SDeclaracionContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by Lab2_GParser#S_Asig.
func (vI *VisitorInterprete) VisitS_Asig(ctx *Lab2G.S_AsigContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by Lab2_GParser#SumAsg.
func (vI *VisitorInterprete) VisitSumAsg(ctx *Lab2G.SumAsgContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by Lab2_GParser#Asig.
func (vI *VisitorInterprete) VisitAsig(ctx *Lab2G.AsigContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by Lab2_GParser#E_Decimal.
func (vI *VisitorInterprete) VisitE_Decimal(ctx *Lab2G.E_DecimalContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by Lab2_GParser#E_Entero.
func (vI *VisitorInterprete) VisitE_Entero(ctx *Lab2G.E_EnteroContext) interface{} {
	nodoEntero := T.NewT_Num(
		ctx.ENTERO().GetText(),
		ctx.ENTERO().GetSymbol().GetLine(),
		ctx.ENTERO().GetSymbol().GetColumn())
	return nodoEntero
}

// Visit a parse tree produced by Lab2_GParser#E_muldiv.
func (vI *VisitorInterprete) VisitE_muldiv(ctx *Lab2G.E_muldivContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by Lab2_GParser#E_Id.
func (vI *VisitorInterprete) VisitE_Id(ctx *Lab2G.E_IdContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by Lab2_GParser#E_sumres.
func (vI *VisitorInterprete) VisitE_sumres(ctx *Lab2G.E_sumresContext) interface{} {
	/// izquierda
	exprIzq := ctx.E(0).Accept(vI).(interprete.AbstrExpr)
	// derecha
	exprDer := ctx.E(1).Accept(vI).(interprete.AbstrExpr)

	if ctx.GetOperacion().GetText() == "+" {
		return NT.NewNT_Suma(exprIzq, exprDer)

	} else {
		panic("resta no implementada")
	}

}

// Visit a parse tree produced by Lab2_GParser#E_Neg.
func (vI *VisitorInterprete) VisitE_Neg(ctx *Lab2G.E_NegContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by Lab2_GParser#E_par.
func (vI *VisitorInterprete) VisitE_par(ctx *Lab2G.E_parContext) interface{} {
	panic("not implemented") // TODO: Implement
}
