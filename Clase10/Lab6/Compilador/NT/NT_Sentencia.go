package NT

import "Lab6/Compilador"

type NT_Sentencia struct {
}

func NewNT_Sentencia() *NT_Sentencia {
	return &NT_Sentencia{}
}

func (N *NT_Sentencia) Compilar(ctx *Compilador.Contexto) *Compilador.Atributos {
	//TODO implement me
	ctx.Gen("setencia")
	return &Compilador.Atributos{
		EV:  make([]string, 0),
		EF:  make([]string, 0),
		Dir: "",
	}

}
