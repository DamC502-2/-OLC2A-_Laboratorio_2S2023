// Code generated from PCompi2N.g4 by ANTLR 4.13.0. DO NOT EDIT.

package PCompi2N // PCompi2N
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by PCompi2NParser.
type PCompi2NVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by PCompi2NParser#s.
	VisitS(ctx *SContext) interface{}

	// Visit a parse tree produced by PCompi2NParser#IFsen.
	VisitIFsen(ctx *IFsenContext) interface{}

	// Visit a parse tree produced by PCompi2NParser#switch.
	VisitSwitch(ctx *SwitchContext) interface{}

	// Visit a parse tree produced by PCompi2NParser#case.
	VisitCase(ctx *CaseContext) interface{}

	// Visit a parse tree produced by PCompi2NParser#defaultc.
	VisitDefaultc(ctx *DefaultcContext) interface{}

	// Visit a parse tree produced by PCompi2NParser#Block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by PCompi2NParser#SentenciaEj.
	VisitSentenciaEj(ctx *SentenciaEjContext) interface{}

	// Visit a parse tree produced by PCompi2NParser#SentenciaSwitch.
	VisitSentenciaSwitch(ctx *SentenciaSwitchContext) interface{}

	// Visit a parse tree produced by PCompi2NParser#CondOr.
	VisitCondOr(ctx *CondOrContext) interface{}

	// Visit a parse tree produced by PCompi2NParser#CondAnd.
	VisitCondAnd(ctx *CondAndContext) interface{}

	// Visit a parse tree produced by PCompi2NParser#CondI.
	VisitCondI(ctx *CondIContext) interface{}

	// Visit a parse tree produced by PCompi2NParser#CondBool.
	VisitCondBool(ctx *CondBoolContext) interface{}

	// Visit a parse tree produced by PCompi2NParser#CondPar.
	VisitCondPar(ctx *CondParContext) interface{}

	// Visit a parse tree produced by PCompi2NParser#CondOper.
	VisitCondOper(ctx *CondOperContext) interface{}

	// Visit a parse tree produced by PCompi2NParser#Enum.
	VisitEnum(ctx *EnumContext) interface{}

	// Visit a parse tree produced by PCompi2NParser#Eid.
	VisitEid(ctx *EidContext) interface{}

	// Visit a parse tree produced by PCompi2NParser#Esr.
	VisitEsr(ctx *EsrContext) interface{}

	// Visit a parse tree produced by PCompi2NParser#Epa.
	VisitEpa(ctx *EpaContext) interface{}

	// Visit a parse tree produced by PCompi2NParser#ENeg.
	VisitENeg(ctx *ENegContext) interface{}

	// Visit a parse tree produced by PCompi2NParser#Emd.
	VisitEmd(ctx *EmdContext) interface{}
}
