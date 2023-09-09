// Code generated from PCompi2.g4 by ANTLR 4.13.0. DO NOT EDIT.

package PCompi2 // PCompi2
import "github.com/antlr4-go/antlr/v4"

type BasePCompi2Visitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePCompi2Visitor) VisitS(ctx *SContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCompi2Visitor) VisitCond(ctx *CondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCompi2Visitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCompi2Visitor) VisitOprel(ctx *OprelContext) interface{} {
	return v.VisitChildren(ctx)
}
