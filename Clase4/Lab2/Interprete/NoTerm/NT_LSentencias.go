package NT

import "Lab2/Interprete"

type NT_LSentencias struct {
	Sentencias []Interprete.AbstrExpr
}

func (ls *NT_LSentencias) Interpretar(ctx *Interprete.Contexto) *Interprete.Resultado {
	//recorremos todas las sentencias
	for _, sentencia := range ls.Sentencias {
		sentencia.Interpretar(ctx)
	}
	//devolvemos nil
	return Interprete.NewNil()

}
func (ls *NT_LSentencias) AddSentencia(sentencia Interprete.AbstrExpr) {
	ls.Sentencias = append(ls.Sentencias, sentencia)
}

func NewLSetencias() *NT_LSentencias {
	return &NT_LSentencias{
		Sentencias: make([]Interprete.AbstrExpr, 0),
	}
}
