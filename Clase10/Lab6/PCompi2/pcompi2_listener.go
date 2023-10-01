// Code generated from PCompi2.g4 by ANTLR 4.13.0. DO NOT EDIT.

package PCompi2 // PCompi2
import "github.com/antlr4-go/antlr/v4"

// PCompi2Listener is a complete listener for a parse tree produced by PCompi2Parser.
type PCompi2Listener interface {
	antlr.ParseTreeListener

	// EnterS is called when entering the s production.
	EnterS(c *SContext)

	// EnterCond is called when entering the cond production.
	EnterCond(c *CondContext)

	// EnterMarcador is called when entering the marcador production.
	EnterMarcador(c *MarcadorContext)

	// EnterMarcador2 is called when entering the marcador2 production.
	EnterMarcador2(c *Marcador2Context)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterOprel is called when entering the oprel production.
	EnterOprel(c *OprelContext)

	// ExitS is called when exiting the s production.
	ExitS(c *SContext)

	// ExitCond is called when exiting the cond production.
	ExitCond(c *CondContext)

	// ExitMarcador is called when exiting the marcador production.
	ExitMarcador(c *MarcadorContext)

	// ExitMarcador2 is called when exiting the marcador2 production.
	ExitMarcador2(c *Marcador2Context)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitOprel is called when exiting the oprel production.
	ExitOprel(c *OprelContext)
}
