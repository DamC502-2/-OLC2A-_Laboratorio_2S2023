// Code generated from Lab2_G.g4 by ANTLR 4.13.0. DO NOT EDIT.

package Lab2G // Lab2_G
import "github.com/antlr4-go/antlr/v4"

// Lab2_GListener is a complete listener for a parse tree produced by Lab2_GParser.
type Lab2_GListener interface {
	antlr.ParseTreeListener

	// EnterSLSentencias is called when entering the SLSentencias production.
	EnterSLSentencias(c *SLSentenciasContext)

	// EnterL_Sentencia is called when entering the L_Sentencia production.
	EnterL_Sentencia(c *L_SentenciaContext)

	// EnterSConsola is called when entering the SConsola production.
	EnterSConsola(c *SConsolaContext)

	// EnterSDeclaracion is called when entering the SDeclaracion production.
	EnterSDeclaracion(c *SDeclaracionContext)

	// EnterS_Asig is called when entering the S_Asig production.
	EnterS_Asig(c *S_AsigContext)

	// EnterSumAsg is called when entering the SumAsg production.
	EnterSumAsg(c *SumAsgContext)

	// EnterAsig is called when entering the Asig production.
	EnterAsig(c *AsigContext)

	// EnterE_Decimal is called when entering the E_Decimal production.
	EnterE_Decimal(c *E_DecimalContext)

	// EnterE_String is called when entering the E_String production.
	EnterE_String(c *E_StringContext)

	// EnterE_Entero is called when entering the E_Entero production.
	EnterE_Entero(c *E_EnteroContext)

	// EnterE_muldiv is called when entering the E_muldiv production.
	EnterE_muldiv(c *E_muldivContext)

	// EnterE_Id is called when entering the E_Id production.
	EnterE_Id(c *E_IdContext)

	// EnterE_sumres is called when entering the E_sumres production.
	EnterE_sumres(c *E_sumresContext)

	// EnterE_Neg is called when entering the E_Neg production.
	EnterE_Neg(c *E_NegContext)

	// EnterE_par is called when entering the E_par production.
	EnterE_par(c *E_parContext)

	// ExitSLSentencias is called when exiting the SLSentencias production.
	ExitSLSentencias(c *SLSentenciasContext)

	// ExitL_Sentencia is called when exiting the L_Sentencia production.
	ExitL_Sentencia(c *L_SentenciaContext)

	// ExitSConsola is called when exiting the SConsola production.
	ExitSConsola(c *SConsolaContext)

	// ExitSDeclaracion is called when exiting the SDeclaracion production.
	ExitSDeclaracion(c *SDeclaracionContext)

	// ExitS_Asig is called when exiting the S_Asig production.
	ExitS_Asig(c *S_AsigContext)

	// ExitSumAsg is called when exiting the SumAsg production.
	ExitSumAsg(c *SumAsgContext)

	// ExitAsig is called when exiting the Asig production.
	ExitAsig(c *AsigContext)

	// ExitE_Decimal is called when exiting the E_Decimal production.
	ExitE_Decimal(c *E_DecimalContext)

	// ExitE_String is called when exiting the E_String production.
	ExitE_String(c *E_StringContext)

	// ExitE_Entero is called when exiting the E_Entero production.
	ExitE_Entero(c *E_EnteroContext)

	// ExitE_muldiv is called when exiting the E_muldiv production.
	ExitE_muldiv(c *E_muldivContext)

	// ExitE_Id is called when exiting the E_Id production.
	ExitE_Id(c *E_IdContext)

	// ExitE_sumres is called when exiting the E_sumres production.
	ExitE_sumres(c *E_sumresContext)

	// ExitE_Neg is called when exiting the E_Neg production.
	ExitE_Neg(c *E_NegContext)

	// ExitE_par is called when exiting the E_par production.
	ExitE_par(c *E_parContext)
}
