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

// EnterS_LSentencias is called when production S_LSentencias is entered.
func (s *BaseLab2_GListener) EnterS_LSentencias(ctx *S_LSentenciasContext) {}

// ExitS_LSentencias is called when production S_LSentencias is exited.
func (s *BaseLab2_GListener) ExitS_LSentencias(ctx *S_LSentenciasContext) {}

// EnterL_Sentencias is called when production L_Sentencias is entered.
func (s *BaseLab2_GListener) EnterL_Sentencias(ctx *L_SentenciasContext) {}

// ExitL_Sentencias is called when production L_Sentencias is exited.
func (s *BaseLab2_GListener) ExitL_Sentencias(ctx *L_SentenciasContext) {}

// EnterS_Consola is called when production S_Consola is entered.
func (s *BaseLab2_GListener) EnterS_Consola(ctx *S_ConsolaContext) {}

// ExitS_Consola is called when production S_Consola is exited.
func (s *BaseLab2_GListener) ExitS_Consola(ctx *S_ConsolaContext) {}

// EnterS_Declaracion is called when production S_Declaracion is entered.
func (s *BaseLab2_GListener) EnterS_Declaracion(ctx *S_DeclaracionContext) {}

// ExitS_Declaracion is called when production S_Declaracion is exited.
func (s *BaseLab2_GListener) ExitS_Declaracion(ctx *S_DeclaracionContext) {}

// EnterS_Asig is called when production S_Asig is entered.
func (s *BaseLab2_GListener) EnterS_Asig(ctx *S_AsigContext) {}

// ExitS_Asig is called when production S_Asig is exited.
func (s *BaseLab2_GListener) ExitS_Asig(ctx *S_AsigContext) {}

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

// EnterE_par is called when production E_par is entered.
func (s *BaseLab2_GListener) EnterE_par(ctx *E_parContext) {}

// ExitE_par is called when production E_par is exited.
func (s *BaseLab2_GListener) ExitE_par(ctx *E_parContext) {}
