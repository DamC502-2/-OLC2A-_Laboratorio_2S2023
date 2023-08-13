package NT

import interprete "Lab2/Interprete"

type NT_LSentencias struct {
	Sentencias []interprete.AbstrExpr
}

// se encarga de añadir sentencias
func (ls *NT_LSentencias) AddSentencia(sentencia interprete.AbstrExpr) {
	ls.Sentencias = append(ls.Sentencias, sentencia)
}

func NewLSentencias() *NT_LSentencias {
	return &NT_LSentencias{
		Sentencias: make([]interprete.AbstrExpr, 0),
	}
}
func (ls *NT_LSentencias) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	for _, sentencia := range ls.Sentencias {
		sentencia.Interpretar(ctx)
	}
	return interprete.NewNil()
}
