// Code generated from Lab2_G.g4 by ANTLR 4.13.0. DO NOT EDIT.

package Lab2G // Lab2_G
import "github.com/antlr4-go/antlr/v4"

type BaseLab2_GVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseLab2_GVisitor) VisitSLSentencias(ctx *SLSentenciasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLab2_GVisitor) VisitL_Sentencia(ctx *L_SentenciaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLab2_GVisitor) VisitSConsola(ctx *SConsolaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLab2_GVisitor) VisitSDeclaracion(ctx *SDeclaracionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLab2_GVisitor) VisitS_Asig(ctx *S_AsigContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLab2_GVisitor) VisitSumAsg(ctx *SumAsgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLab2_GVisitor) VisitAsig(ctx *AsigContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLab2_GVisitor) VisitE_Decimal(ctx *E_DecimalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLab2_GVisitor) VisitE_String(ctx *E_StringContext) interface{} {
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

func (v *BaseLab2_GVisitor) VisitE_Neg(ctx *E_NegContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLab2_GVisitor) VisitE_par(ctx *E_parContext) interface{} {
	return v.VisitChildren(ctx)
}
