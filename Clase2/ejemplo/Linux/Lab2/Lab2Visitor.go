package main

import (
	"Lab2/Lab2G"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"log"
)

type Lab2GVisitorI struct {
	antlr.ParseTreeVisitor
}

// funcion para

func (l Lab2GVisitorI) VisitChildren(node antlr.RuleNode) interface{} {
	//TODO implement me
	panic("implement me")
}

func (l Lab2GVisitorI) VisitTerminal(node antlr.TerminalNode) interface{} {
	//TODO implement me

	panic("implement me")
}

func (l Lab2GVisitorI) VisitErrorNode(node antlr.ErrorNode) interface{} {
	//TODO implement me
	panic("implement me")
}

func (l Lab2GVisitorI) VisitSLSentencias(ctx *Lab2G.SLSentenciasContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (l Lab2GVisitorI) VisitL_Sentencia(ctx *Lab2G.L_SentenciaContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (l Lab2GVisitorI) VisitSConsola(ctx *Lab2G.SConsolaContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (l Lab2GVisitorI) VisitSDeclaracion(ctx *Lab2G.SDeclaracionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (l Lab2GVisitorI) VisitS_Asig(ctx *Lab2G.S_AsigContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (l Lab2GVisitorI) VisitSumAsg(ctx *Lab2G.SumAsgContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (l Lab2GVisitorI) VisitAsig(ctx *Lab2G.AsigContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (l Lab2GVisitorI) VisitE_Decimal(ctx *Lab2G.E_DecimalContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (l Lab2GVisitorI) VisitE_Entero(ctx *Lab2G.E_EnteroContext) interface{} {
	fmt.Println("entrando en E -> entero ", ctx.ENTERO().GetText())
	//
	fmt.Println("saliendo de entero: ", ctx.ENTERO().GetText())
	return nil
}

func (l Lab2GVisitorI) VisitE_muldiv(ctx *Lab2G.E_muldivContext) interface{} {
	//TODO implement me
	fmt.Println("Entrando en #E_Muldiv")
	fmt.Println("operacion: ", ctx.GetOp().GetText())

	ctx.E(0).Accept(l) //desplazamos en E1
	ctx.E(1).Accept(l) //desplazmaos en E2
	//reducción
	fmt.Println("saliendo de E -> E / * E ")
	return nil
}

func (l Lab2GVisitorI) VisitE_Id(ctx *Lab2G.E_IdContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (l Lab2GVisitorI) VisitE_sumres(ctx *Lab2G.E_sumresContext) interface{} {
	//TODO implement me

	fmt.Println("Entrando a E -> E (+|-) E ")
	fmt.Println(ctx.GetOperacion().GetText())

	ctx.E(0).Accept(l) //deplazamos en E1
	ctx.E(1).Accept(l) //desplazamos en E2
	//reducción
	fmt.Println("Saliendo de E +|- E ")
	fmt.Println("operacion", ctx.GetOperacion().GetText())

	return nil

}

func (l Lab2GVisitorI) VisitE_Neg(ctx *Lab2G.E_NegContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (l Lab2GVisitorI) VisitE_par(ctx *Lab2G.E_parContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (l Lab2GVisitorI) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		log.Fatal(val.GetText())
		return nil
	// indica que se ha reconocido un nodo S
	case *Lab2G.SContext:
		fmt.Print("se ha reconocido S -> a")
		return tree.Accept(l)

	default:
		nodo := tree.Accept(l)
		return nodo
	}
	return nil
}

// construcctor para hacer un nuevo Lab2GVisitorI
func NewLab2VisitorI() Lab2G.Lab2_GVisitor {
	return &Lab2GVisitorI{
		ParseTreeVisitor: &Lab2G.BaseLab2_GVisitor{},
	}
}
