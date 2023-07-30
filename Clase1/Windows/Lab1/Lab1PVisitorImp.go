package main

import (
	"Lab1/Lab1P"
	"fmt"
	"log"

	"github.com/antlr4-go/antlr/v4"
)

type Lab1PVisitorImp struct {
	antlr.ParseTreeVisitor
}

func NewLab1PVisitorImp() Lab1P.Lab1_GVisitor {
	return &Lab1PVisitorImp{
		ParseTreeVisitor: &Lab1P.BaseLab1_GVisitor{},
	}

}

func (l *Lab1PVisitorImp) VisitS(ctx *Lab1P.SContext) interface{} {
	ctx.A().Accept(l)
	return "Se ha completado el reconocimiento"
}

func (l *Lab1PVisitorImp) VisitA1(ctx *Lab1P.A1Context) interface{} {
	//ctx.A().Accept(l)
	return ctx.A().Accept(l)
}
func (l *Lab1PVisitorImp) VisitA2(ctx *Lab1P.A2Context) interface{} {
	fmt.Println("imprimiendo entero")
	fmt.Println(ctx.ENTERO().GetText())
	return nil
}
func (l *Lab1PVisitorImp) VisitA3(ctx *Lab1P.A3Context) interface{} {
	return nil
}
func (l *Lab1PVisitorImp) VisitA4(ctx *Lab1P.A4Context) interface{} {
	fmt.Println("ID en la linea: ")
	fmt.Println(ctx.ID().GetSymbol().GetLine())
	fmt.Println("ID en la columna: ")
	fmt.Println(ctx.ID().GetSymbol().GetColumn())
	fmt.Print(ctx.ID().GetText())
	return nil
}
func (l *Lab1PVisitorImp) VisitA5(ctx *Lab1P.A5Context) interface{} {

	return nil
}

// / funcion Visit
func (l *Lab1PVisitorImp) Visit(arbol antlr.ParseTree) interface{} {
	switch val := arbol.(type) {
	case *antlr.ErrorNodeImpl:
		log.Fatal(val.GetText())
		return nil
	default:
		nodo := arbol.Accept(l)
		return nodo
	}
}
