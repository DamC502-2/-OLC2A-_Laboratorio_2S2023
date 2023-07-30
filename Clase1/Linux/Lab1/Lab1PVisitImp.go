package main

import (
	Lab1P "Lab1/Lab1P"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"log"
	"strconv"
)

type Lab1PVisitorImp struct {
	antlr.ParseTreeVisitor
}

func (l Lab1PVisitorImp) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		log.Fatal(val.GetText())
		return nil
	// lo que see
	case *Lab1P.SContext:
		fmt.Print("se ha reconocido S -> a")
		return tree.Accept(l)
	default:
		nodo := tree.Accept(l)
		if numero, ok := nodo.(int32); ok {
			return numero
		}
		if cadena, ok := nodo.(string); ok {
			return cadena
		}
		fmt.Print("tipo extraño")
		return nodo
	}
	return nil
}

func (l Lab1PVisitorImp) VisitS(ctx *Lab1P.SContext) interface{} {

	return ctx.A().Accept(l)
}

func (l Lab1PVisitorImp) VisitA0(ctx *Lab1P.A0Context) interface{} {
	return ctx.A().Accept(l)
}

func (l Lab1PVisitorImp) VisitAENTERO(ctx *Lab1P.AENTEROContext) interface{} {
	//TODO implement me
	numero, _ := strconv.Atoi(ctx.ENTERO().GetText())
	return numero
}

func (l Lab1PVisitorImp) VisitADECIMAL(ctx *Lab1P.ADECIMALContext) interface{} {
	//TODO implement me
	numero, _ := strconv.ParseFloat(ctx.DECIMAL().GetText(), 64)
	return numero
}

func (l Lab1PVisitorImp) VisitAID(ctx *Lab1P.AIDContext) interface{} {
	//TODO implement me
	return ctx.ID().GetText()
}

func (l Lab1PVisitorImp) VisitAEPSILUM(ctx *Lab1P.AEPSILUMContext) interface{} {
	//TODO implement me
	return "epsilum "
}

func NewLab1VisitorImp() Lab1P.Lab1_GVisitor {
	return &Lab1PVisitorImp{
		ParseTreeVisitor: &Lab1P.BaseLab1_GVisitor{},
	}
}
