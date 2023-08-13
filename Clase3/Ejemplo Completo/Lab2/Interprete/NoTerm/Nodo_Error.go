package NT

import "Lab2/Interprete"

type NT_Error struct {
	errorS string
}

func NewNT_Error(errorS string) *NT_Error {
	return &NT_Error{errorS: errorS}
}

func (N *NT_Error) Interpretar(ctx *Interprete.Contexto) *Interprete.Resultado {
	ctx.AddError(N.errorS)
	return Interprete.NewNil()
}
