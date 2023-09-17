// Code generated from PCompi2N.g4 by ANTLR 4.13.0. DO NOT EDIT.

package PCompi2N // PCompi2N
import "github.com/antlr4-go/antlr/v4"

// BasePCompi2NListener is a complete listener for a parse tree produced by PCompi2NParser.
type BasePCompi2NListener struct{}

var _ PCompi2NListener = &BasePCompi2NListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePCompi2NListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePCompi2NListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePCompi2NListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePCompi2NListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterS is called when production s is entered.
func (s *BasePCompi2NListener) EnterS(ctx *SContext) {}

// ExitS is called when production s is exited.
func (s *BasePCompi2NListener) ExitS(ctx *SContext) {}

// EnterCondOr is called when production CondOr is entered.
func (s *BasePCompi2NListener) EnterCondOr(ctx *CondOrContext) {}

// ExitCondOr is called when production CondOr is exited.
func (s *BasePCompi2NListener) ExitCondOr(ctx *CondOrContext) {}

// EnterCondAnd is called when production CondAnd is entered.
func (s *BasePCompi2NListener) EnterCondAnd(ctx *CondAndContext) {}

// ExitCondAnd is called when production CondAnd is exited.
func (s *BasePCompi2NListener) ExitCondAnd(ctx *CondAndContext) {}

// EnterCondI is called when production CondI is entered.
func (s *BasePCompi2NListener) EnterCondI(ctx *CondIContext) {}

// ExitCondI is called when production CondI is exited.
func (s *BasePCompi2NListener) ExitCondI(ctx *CondIContext) {}

// EnterCondBool is called when production CondBool is entered.
func (s *BasePCompi2NListener) EnterCondBool(ctx *CondBoolContext) {}

// ExitCondBool is called when production CondBool is exited.
func (s *BasePCompi2NListener) ExitCondBool(ctx *CondBoolContext) {}

// EnterCondPar is called when production CondPar is entered.
func (s *BasePCompi2NListener) EnterCondPar(ctx *CondParContext) {}

// ExitCondPar is called when production CondPar is exited.
func (s *BasePCompi2NListener) ExitCondPar(ctx *CondParContext) {}

// EnterCondOper is called when production CondOper is entered.
func (s *BasePCompi2NListener) EnterCondOper(ctx *CondOperContext) {}

// ExitCondOper is called when production CondOper is exited.
func (s *BasePCompi2NListener) ExitCondOper(ctx *CondOperContext) {}

// EnterEnum is called when production Enum is entered.
func (s *BasePCompi2NListener) EnterEnum(ctx *EnumContext) {}

// ExitEnum is called when production Enum is exited.
func (s *BasePCompi2NListener) ExitEnum(ctx *EnumContext) {}

// EnterEid is called when production Eid is entered.
func (s *BasePCompi2NListener) EnterEid(ctx *EidContext) {}

// ExitEid is called when production Eid is exited.
func (s *BasePCompi2NListener) ExitEid(ctx *EidContext) {}

// EnterEsr is called when production Esr is entered.
func (s *BasePCompi2NListener) EnterEsr(ctx *EsrContext) {}

// ExitEsr is called when production Esr is exited.
func (s *BasePCompi2NListener) ExitEsr(ctx *EsrContext) {}

// EnterEpa is called when production Epa is entered.
func (s *BasePCompi2NListener) EnterEpa(ctx *EpaContext) {}

// ExitEpa is called when production Epa is exited.
func (s *BasePCompi2NListener) ExitEpa(ctx *EpaContext) {}

// EnterENeg is called when production ENeg is entered.
func (s *BasePCompi2NListener) EnterENeg(ctx *ENegContext) {}

// ExitENeg is called when production ENeg is exited.
func (s *BasePCompi2NListener) ExitENeg(ctx *ENegContext) {}

// EnterEmd is called when production Emd is entered.
func (s *BasePCompi2NListener) EnterEmd(ctx *EmdContext) {}

// ExitEmd is called when production Emd is exited.
func (s *BasePCompi2NListener) ExitEmd(ctx *EmdContext) {}
