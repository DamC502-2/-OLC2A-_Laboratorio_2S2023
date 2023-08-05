// Code generated from Lab2_G.g4 by ANTLR 4.13.0. DO NOT EDIT.

package Lab2G // Lab2_G
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by Lab2_GParser.
type Lab2_GVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by Lab2_GParser#S_LSentencias.
	VisitS_LSentencias(ctx *S_LSentenciasContext) interface{}

	// Visit a parse tree produced by Lab2_GParser#L_Sentencias.
	VisitL_Sentencias(ctx *L_SentenciasContext) interface{}

	// Visit a parse tree produced by Lab2_GParser#S_Consola.
	VisitS_Consola(ctx *S_ConsolaContext) interface{}

	// Visit a parse tree produced by Lab2_GParser#S_Declaracion.
	VisitS_Declaracion(ctx *S_DeclaracionContext) interface{}

	// Visit a parse tree produced by Lab2_GParser#S_Asig.
	VisitS_Asig(ctx *S_AsigContext) interface{}

	// Visit a parse tree produced by Lab2_GParser#Asig.
	VisitAsig(ctx *AsigContext) interface{}

	// Visit a parse tree produced by Lab2_GParser#E_Decimal.
	VisitE_Decimal(ctx *E_DecimalContext) interface{}

	// Visit a parse tree produced by Lab2_GParser#E_Entero.
	VisitE_Entero(ctx *E_EnteroContext) interface{}

	// Visit a parse tree produced by Lab2_GParser#E_muldiv.
	VisitE_muldiv(ctx *E_muldivContext) interface{}

	// Visit a parse tree produced by Lab2_GParser#E_Id.
	VisitE_Id(ctx *E_IdContext) interface{}

	// Visit a parse tree produced by Lab2_GParser#E_sumres.
	VisitE_sumres(ctx *E_sumresContext) interface{}

	// Visit a parse tree produced by Lab2_GParser#E_par.
	VisitE_par(ctx *E_parContext) interface{}
}
