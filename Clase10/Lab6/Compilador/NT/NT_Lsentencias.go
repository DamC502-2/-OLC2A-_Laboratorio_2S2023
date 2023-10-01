package NT

import "Lab6/Compilador"

type NT_Lsentencias struct {
	Sentencias []Compilador.CAbstrExpr
}

func (nts *NT_Lsentencias) Compilar(ctx *Compilador.Contexto) *Compilador.Atributos {
	//TODO implement me
	for _, sentencia := range nts.Sentencias {
		sentencia.Compilar(ctx)
	}
	return &Compilador.Atributos{
		EV:  make([]string, 0),
		EF:  make([]string, 0),
		Dir: "",
	}
}

func NewNT_Lsentencias() *NT_Lsentencias {
	return &NT_Lsentencias{
		Sentencias: make([]Compilador.CAbstrExpr, 0),
	}
}

func (nts *NT_Lsentencias) push(nodo Compilador.CAbstrExpr) {
	nts.Sentencias = append(nts.Sentencias, nodo)
}
