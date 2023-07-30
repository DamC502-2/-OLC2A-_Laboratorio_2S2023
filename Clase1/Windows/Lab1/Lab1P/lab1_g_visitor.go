// Code generated from Lab1_G.g4 by ANTLR 4.13.0. DO NOT EDIT.

package Lab1P // Lab1_G
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by Lab1_GParser.
type Lab1_GVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by Lab1_GParser#s.
	VisitS(ctx *SContext) interface{}

	// Visit a parse tree produced by Lab1_GParser#A1.
	VisitA1(ctx *A1Context) interface{}

	// Visit a parse tree produced by Lab1_GParser#A2.
	VisitA2(ctx *A2Context) interface{}

	// Visit a parse tree produced by Lab1_GParser#A3.
	VisitA3(ctx *A3Context) interface{}

	// Visit a parse tree produced by Lab1_GParser#A4.
	VisitA4(ctx *A4Context) interface{}

	// Visit a parse tree produced by Lab1_GParser#A5.
	VisitA5(ctx *A5Context) interface{}
}
