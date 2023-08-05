// Code generated from Lab2_G.g4 by ANTLR 4.13.0. DO NOT EDIT.

package Lab2G // Lab2_G
import "github.com/antlr4-go/antlr/v4"

// Lab2_GListener is a complete listener for a parse tree produced by Lab2_GParser.
type Lab2_GListener interface {
	antlr.ParseTreeListener

	// EnterS_LSentencias is called when entering the S_LSentencias production.
	EnterS_LSentencias(c *S_LSentenciasContext)

	// EnterL_Sentencias is called when entering the L_Sentencias production.
	EnterL_Sentencias(c *L_SentenciasContext)

	// EnterS_Consola is called when entering the S_Consola production.
	EnterS_Consola(c *S_ConsolaContext)

	// EnterS_Declaracion is called when entering the S_Declaracion production.
	EnterS_Declaracion(c *S_DeclaracionContext)

	// EnterS_Asig is called when entering the S_Asig production.
	EnterS_Asig(c *S_AsigContext)

	// EnterAsig is called when entering the Asig production.
	EnterAsig(c *AsigContext)

	// EnterE_Decimal is called when entering the E_Decimal production.
	EnterE_Decimal(c *E_DecimalContext)

	// EnterE_Entero is called when entering the E_Entero production.
	EnterE_Entero(c *E_EnteroContext)

	// EnterE_muldiv is called when entering the E_muldiv production.
	EnterE_muldiv(c *E_muldivContext)

	// EnterE_Id is called when entering the E_Id production.
	EnterE_Id(c *E_IdContext)

	// EnterE_sumres is called when entering the E_sumres production.
	EnterE_sumres(c *E_sumresContext)

	// EnterE_par is called when entering the E_par production.
	EnterE_par(c *E_parContext)

	// ExitS_LSentencias is called when exiting the S_LSentencias production.
	ExitS_LSentencias(c *S_LSentenciasContext)

	// ExitL_Sentencias is called when exiting the L_Sentencias production.
	ExitL_Sentencias(c *L_SentenciasContext)

	// ExitS_Consola is called when exiting the S_Consola production.
	ExitS_Consola(c *S_ConsolaContext)

	// ExitS_Declaracion is called when exiting the S_Declaracion production.
	ExitS_Declaracion(c *S_DeclaracionContext)

	// ExitS_Asig is called when exiting the S_Asig production.
	ExitS_Asig(c *S_AsigContext)

	// ExitAsig is called when exiting the Asig production.
	ExitAsig(c *AsigContext)

	// ExitE_Decimal is called when exiting the E_Decimal production.
	ExitE_Decimal(c *E_DecimalContext)

	// ExitE_Entero is called when exiting the E_Entero production.
	ExitE_Entero(c *E_EnteroContext)

	// ExitE_muldiv is called when exiting the E_muldiv production.
	ExitE_muldiv(c *E_muldivContext)

	// ExitE_Id is called when exiting the E_Id production.
	ExitE_Id(c *E_IdContext)

	// ExitE_sumres is called when exiting the E_sumres production.
	ExitE_sumres(c *E_sumresContext)

	// ExitE_par is called when exiting the E_par production.
	ExitE_par(c *E_parContext)
}
