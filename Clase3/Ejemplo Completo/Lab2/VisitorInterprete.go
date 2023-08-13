package main

import (
	"Lab2/Interprete"
	NT "Lab2/Interprete/NoTerm"
	T "Lab2/Interprete/Terminales"

	"Lab2/Lab2G"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
)

type VisitorInterprete struct {
	antlr.ParseTreeVisitor
	Raiz Interprete.AbstrExpr
}

func NewVisitorInterprete() Lab2G.Lab2_GVisitor {
	return &VisitorInterprete{
		ParseTreeVisitor: &Lab2G.BaseLab2_GVisitor{},
		Raiz:             nil,
	}
}

func (v VisitorInterprete) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		//log.Fatal(val.GetText())
		return NT.NewNT_Error(val.GetText())
	default:
		//comprobamos el el nodo sea valido
		nodoIntepreter, ok := tree.Accept(v).(Interprete.AbstrExpr)
		if ok {
			return nodoIntepreter
		}
		fmt.Print("nodo desconocido")
		return NT.NewNT_Error("nodo desconocido")
	}
}

//func (v VisitorInterprete) VisitChildren(node antlr.RuleNode) interface{} {
//	//TODO implement me
//	panic("implement me")
//}

//func (v VisitorInterprete) VisitTerminal(node antlr.TerminalNode) interface{} {
//	//TODO implement me
//	panic("implement me")
//}

func (v VisitorInterprete) VisitErrorNode(node antlr.ErrorNode) interface{} {
	fmt.Print(node.GetText())
	return NT.NewNT_Error(node.GetText())
}

// S -> Lsentencias
func (v VisitorInterprete) VisitSLSentencias(ctx *Lab2G.SLSentenciasContext) interface{} {
	nodo := ctx.L_sentencias().Accept(v).(Interprete.AbstrExpr)
	v.Raiz = nodo
	return nodo
}

// LSetencia -> sentencia* #L_Sentencia
func (v VisitorInterprete) VisitL_Sentencia(ctx *Lab2G.L_SentenciaContext) interface{} {
	lsentencias := NT.NewLSetencias()
	sentencias := ctx.AllSentencia()
	for _, sentencia := range sentencias {
		nodoSentencia := sentencia.Accept(v).(Interprete.AbstrExpr)
		lsentencias.AddSentencia(nodoSentencia)
	}
	return lsentencias
}

func (v VisitorInterprete) VisitSConsola(ctx *Lab2G.SConsolaContext) interface{} {
	consola := NT.NewNT_Clog()
	lexpresiones := ctx.AllE()
	for _, lexpresione := range lexpresiones {
		nodoE := lexpresione.Accept(v).(Interprete.AbstrExpr)
		consola.AddExpresion(nodoE)
	}
	return consola
}

// sentencia ->   INT ID  (  ASIG e  )? #SDeclaracion
func (v VisitorInterprete) VisitSDeclaracion(ctx *Lab2G.SDeclaracionContext) interface{} {
	ID := ctx.ID().GetText()
	tipo := ctx.INT().GetText()

	var expr Interprete.AbstrExpr = nil
	if ctx.E() != nil {
		expr = ctx.E().Accept(v).(Interprete.AbstrExpr)
	}

	declaracion := NT.NewNT_DeclVar(ID, tipo, expr,
		ctx.ID().GetSymbol().GetLine(),
		ctx.ID().GetSymbol().GetColumn())
	return declaracion
}

// sentencia -> asignacion #S_Asig
func (v VisitorInterprete) VisitS_Asig(ctx *Lab2G.S_AsigContext) interface{} {
	return ctx.Asignacion().Accept(v).(Interprete.AbstrExpr)

}

// pendiente - opccional
func (v VisitorInterprete) VisitSumAsg(ctx *Lab2G.SumAsgContext) interface{} {
	//TODO implement me
	panic("implement me")
}

//  asignacion -> ID ASIG  e #Asig
func (v VisitorInterprete) VisitAsig(ctx *Lab2G.AsigContext) interface{} {
	ID := ctx.ID().GetText()
	expr := ctx.E().Accept(v).(Interprete.AbstrExpr)
	asig := NT.NewNT_AsigVar(ID, expr,
		ctx.ID().GetSymbol().GetLine(),
		ctx.ID().GetSymbol().GetColumn())
	return asig
}

// E -> Decimal
func (v VisitorInterprete) VisitE_Decimal(ctx *Lab2G.E_DecimalContext) interface{} {
	return T.NewT_Num(ctx.DECIMAL().GetText(),
		ctx.DECIMAL().GetSymbol().GetLine(),
		ctx.DECIMAL().GetSymbol().GetColumn())
}

// E -> Entero
func (v VisitorInterprete) VisitE_Entero(ctx *Lab2G.E_EnteroContext) interface{} {
	return T.NewT_Num(ctx.ENTERO().GetText(),
		ctx.ENTERO().GetSymbol().GetLine(),
		ctx.ENTERO().GetSymbol().GetColumn())
}

// E -> E ( * | /)=op E
func (v VisitorInterprete) VisitE_muldiv(ctx *Lab2G.E_muldivContext) interface{} {
	// izquierda
	e1 := ctx.E(0).Accept(v).(Interprete.AbstrExpr)
	// derecha
	e2 := ctx.E(1).Accept(v).(Interprete.AbstrExpr)
	if ctx.GetOp().GetText() == "*" {
		// multiplicacion
		return NT.NewNT_Mult(e1, e2)
	} else {
		// division
		return NT.NewNT_Error("division no implementada")
	}

}

func (v VisitorInterprete) VisitE_Id(ctx *Lab2G.E_IdContext) interface{} {
	return NT.NewNT_ID(ctx.ID().GetText(),
		ctx.ID().GetSymbol().GetLine(),
		ctx.ID().GetSymbol().GetLine())
}

// e ->  e operacion=('+'| '-') e  #E_sumres
func (v VisitorInterprete) VisitE_sumres(ctx *Lab2G.E_sumresContext) interface{} {
	// izquierda
	e1 := ctx.E(0).Accept(v).(Interprete.AbstrExpr)
	// derecha
	e2 := ctx.E(1).Accept(v).(Interprete.AbstrExpr)
	if ctx.GetOperacion().GetText() == "+" {
		return NT.NewNT_Suma(e1, e2)
	} else {
		return NT.NewNT_Error("division no implementada")
	}

}

func (v VisitorInterprete) VisitE_Neg(ctx *Lab2G.E_NegContext) interface{} {
	//TODO implement me
	panic("implement me")
}

// E -> ( E )
func (v VisitorInterprete) VisitE_par(ctx *Lab2G.E_parContext) interface{} {
	return ctx.E().Accept(v).(Interprete.AbstrExpr)
}
