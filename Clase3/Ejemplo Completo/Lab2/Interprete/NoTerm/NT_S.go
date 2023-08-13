package NT

import "Lab2/Interprete"

type NT_S struct {
	lSentencias Interprete.AbstrExpr
}

func (s *NT_S) Interpretar(ctx *Interprete.Contexto) *Interprete.Resultado {
	// se inicia la ejecución
	return s.lSentencias.Interpretar(ctx)
}

func NewNT_S(lSentencias Interprete.AbstrExpr) *NT_S {
	return &NT_S{lSentencias: lSentencias}
}
