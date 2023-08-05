// Code generated from Lab2_G.g4 by ANTLR 4.13.0. DO NOT EDIT.

package Lab2G // Lab2_G
import "github.com/antlr4-go/antlr/v4"

type BaseLab2_GVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseLab2_GVisitor) VisitS_LSentencias(ctx *S_LSentenciasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLab2_GVisitor) VisitL_Sentencias(ctx *L_SentenciasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLab2_GVisitor) VisitS_Consola(ctx *S_ConsolaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLab2_GVisitor) VisitS_Declaracion(ctx *S_DeclaracionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLab2_GVisitor) VisitS_Asig(ctx *S_AsigContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLab2_GVisitor) VisitAsig(ctx *AsigContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLab2_GVisitor) VisitE_Decimal(ctx *E_DecimalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLab2_GVisitor) VisitE_Entero(ctx *E_EnteroContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLab2_GVisitor) VisitE_muldiv(ctx *E_muldivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLab2_GVisitor) VisitE_Id(ctx *E_IdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLab2_GVisitor) VisitE_sumres(ctx *E_sumresContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLab2_GVisitor) VisitE_par(ctx *E_parContext) interface{} {
	return v.VisitChildren(ctx)
}
