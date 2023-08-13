// Code generated from Lab2_G.g4 by ANTLR 4.13.0. DO NOT EDIT.

package Lab2G // Lab2_G
import "github.com/antlr4-go/antlr/v4"

// BaseLab2_GListener is a complete listener for a parse tree produced by Lab2_GParser.
type BaseLab2_GListener struct{}

var _ Lab2_GListener = &BaseLab2_GListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLab2_GListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLab2_GListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLab2_GListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLab2_GListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSLSentencias is called when production SLSentencias is entered.
func (s *BaseLab2_GListener) EnterSLSentencias(ctx *SLSentenciasContext) {}

// ExitSLSentencias is called when production SLSentencias is exited.
func (s *BaseLab2_GListener) ExitSLSentencias(ctx *SLSentenciasContext) {}

// EnterL_Sentencia is called when production L_Sentencia is entered.
func (s *BaseLab2_GListener) EnterL_Sentencia(ctx *L_SentenciaContext) {}

// ExitL_Sentencia is called when production L_Sentencia is exited.
func (s *BaseLab2_GListener) ExitL_Sentencia(ctx *L_SentenciaContext) {}

// EnterSConsola is called when production SConsola is entered.
func (s *BaseLab2_GListener) EnterSConsola(ctx *SConsolaContext) {}

// ExitSConsola is called when production SConsola is exited.
func (s *BaseLab2_GListener) ExitSConsola(ctx *SConsolaContext) {}

// EnterSDeclaracion is called when production SDeclaracion is entered.
func (s *BaseLab2_GListener) EnterSDeclaracion(ctx *SDeclaracionContext) {}

// ExitSDeclaracion is called when production SDeclaracion is exited.
func (s *BaseLab2_GListener) ExitSDeclaracion(ctx *SDeclaracionContext) {}

// EnterS_Asig is called when production S_Asig is entered.
func (s *BaseLab2_GListener) EnterS_Asig(ctx *S_AsigContext) {}

// ExitS_Asig is called when production S_Asig is exited.
func (s *BaseLab2_GListener) ExitS_Asig(ctx *S_AsigContext) {}

// EnterSWhile is called when production SWhile is entered.
func (s *BaseLab2_GListener) EnterSWhile(ctx *SWhileContext) {}

// ExitSWhile is called when production SWhile is exited.
func (s *BaseLab2_GListener) ExitSWhile(ctx *SWhileContext) {}

// EnterSumAsg is called when production SumAsg is entered.
func (s *BaseLab2_GListener) EnterSumAsg(ctx *SumAsgContext) {}

// ExitSumAsg is called when production SumAsg is exited.
func (s *BaseLab2_GListener) ExitSumAsg(ctx *SumAsgContext) {}

// EnterAsig is called when production Asig is entered.
func (s *BaseLab2_GListener) EnterAsig(ctx *AsigContext) {}

// ExitAsig is called when production Asig is exited.
func (s *BaseLab2_GListener) ExitAsig(ctx *AsigContext) {}

// EnterE_Decimal is called when production E_Decimal is entered.
func (s *BaseLab2_GListener) EnterE_Decimal(ctx *E_DecimalContext) {}

// ExitE_Decimal is called when production E_Decimal is exited.
func (s *BaseLab2_GListener) ExitE_Decimal(ctx *E_DecimalContext) {}

// EnterE_Entero is called when production E_Entero is entered.
func (s *BaseLab2_GListener) EnterE_Entero(ctx *E_EnteroContext) {}

// ExitE_Entero is called when production E_Entero is exited.
func (s *BaseLab2_GListener) ExitE_Entero(ctx *E_EnteroContext) {}

// EnterE_muldiv is called when production E_muldiv is entered.
func (s *BaseLab2_GListener) EnterE_muldiv(ctx *E_muldivContext) {}

// ExitE_muldiv is called when production E_muldiv is exited.
func (s *BaseLab2_GListener) ExitE_muldiv(ctx *E_muldivContext) {}

// EnterE_Id is called when production E_Id is entered.
func (s *BaseLab2_GListener) EnterE_Id(ctx *E_IdContext) {}

// ExitE_Id is called when production E_Id is exited.
func (s *BaseLab2_GListener) ExitE_Id(ctx *E_IdContext) {}

// EnterE_sumres is called when production E_sumres is entered.
func (s *BaseLab2_GListener) EnterE_sumres(ctx *E_sumresContext) {}

// ExitE_sumres is called when production E_sumres is exited.
func (s *BaseLab2_GListener) ExitE_sumres(ctx *E_sumresContext) {}

// EnterE_Neg is called when production E_Neg is entered.
func (s *BaseLab2_GListener) EnterE_Neg(ctx *E_NegContext) {}

// ExitE_Neg is called when production E_Neg is exited.
func (s *BaseLab2_GListener) ExitE_Neg(ctx *E_NegContext) {}

// EnterE_par is called when production E_par is entered.
func (s *BaseLab2_GListener) EnterE_par(ctx *E_parContext) {}

// ExitE_par is called when production E_par is exited.
func (s *BaseLab2_GListener) ExitE_par(ctx *E_parContext) {}
