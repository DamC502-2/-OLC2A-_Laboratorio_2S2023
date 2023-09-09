// Code generated from PCompi2.g4 by ANTLR 4.13.0. DO NOT EDIT.

package PCompi2 // PCompi2
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by PCompi2Parser.
type PCompi2Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by PCompi2Parser#s.
	VisitS(ctx *SContext) interface{}

	// Visit a parse tree produced by PCompi2Parser#cond.
	VisitCond(ctx *CondContext) interface{}

	// Visit a parse tree produced by PCompi2Parser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by PCompi2Parser#oprel.
	VisitOprel(ctx *OprelContext) interface{}
}
