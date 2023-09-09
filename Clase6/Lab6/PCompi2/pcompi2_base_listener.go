// Code generated from PCompi2.g4 by ANTLR 4.13.0. DO NOT EDIT.

package PCompi2 // PCompi2
import "github.com/antlr4-go/antlr/v4"

// BasePCompi2Listener is a complete listener for a parse tree produced by PCompi2Parser.
type BasePCompi2Listener struct{}

var _ PCompi2Listener = &BasePCompi2Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePCompi2Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePCompi2Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePCompi2Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePCompi2Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterS is called when production s is entered.
func (s *BasePCompi2Listener) EnterS(ctx *SContext) {}

// ExitS is called when production s is exited.
func (s *BasePCompi2Listener) ExitS(ctx *SContext) {}

// EnterCond is called when production cond is entered.
func (s *BasePCompi2Listener) EnterCond(ctx *CondContext) {}

// ExitCond is called when production cond is exited.
func (s *BasePCompi2Listener) ExitCond(ctx *CondContext) {}

// EnterExpr is called when production expr is entered.
func (s *BasePCompi2Listener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BasePCompi2Listener) ExitExpr(ctx *ExprContext) {}

// EnterOprel is called when production oprel is entered.
func (s *BasePCompi2Listener) EnterOprel(ctx *OprelContext) {}

// ExitOprel is called when production oprel is exited.
func (s *BasePCompi2Listener) ExitOprel(ctx *OprelContext) {}
