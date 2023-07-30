// Code generated from Lab1_G.g4 by ANTLR 4.13.0. DO NOT EDIT.

package Lab1P // Lab1_G
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by Lab1_GParser.
type Lab1_GVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by Lab1_GParser#s.
	VisitS(ctx *SContext) interface{}

	// Visit a parse tree produced by Lab1_GParser#A0.
	VisitA0(ctx *A0Context) interface{}

	// Visit a parse tree produced by Lab1_GParser#AENTERO.
	VisitAENTERO(ctx *AENTEROContext) interface{}

	// Visit a parse tree produced by Lab1_GParser#ADECIMAL.
	VisitADECIMAL(ctx *ADECIMALContext) interface{}

	// Visit a parse tree produced by Lab1_GParser#AID.
	VisitAID(ctx *AIDContext) interface{}

	// Visit a parse tree produced by Lab1_GParser#AEPSILUM.
	VisitAEPSILUM(ctx *AEPSILUMContext) interface{}
}
