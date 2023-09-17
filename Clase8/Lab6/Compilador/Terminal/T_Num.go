package Terminal

import "Lab6/Compilador"

type T_Numero struct {
	num string
}

func (t T_Numero) Compilar(ctx *Compilador.Contexto) *Compilador.Atributos {
	//TODO implement me
	return &Compilador.Atributos{
		EV:  make([]string, 0, 0),
		EF:  make([]string, 0, 0),
		Dir: t.num,
	}
}

func NewT_Numero(num string) *T_Numero {
	return &T_Numero{num: num}
}

