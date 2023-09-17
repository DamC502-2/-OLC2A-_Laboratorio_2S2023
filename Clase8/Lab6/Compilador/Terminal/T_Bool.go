package Terminal

import "Lab6/Compilador"

type T_Bool struct {
	valor string //true ó false

}

func (t T_Bool) Compilar(ctx *Compilador.Contexto) *Compilador.Atributos {
	EV := ctx.NewEtq()
	EF := ctx.NewEtq()

	if (t.valor == "falso") {
		ctx.Gen("goto " + EF)
	} else {
		ctx.Gen("goto " + EV)
	}
	EVs := make([]string, 1, 1)
	EVs = append(EVs, EV)

	EFs := make([]string, 1, 1)
	EFs = append(EFs, EF)

	return &Compilador.Atributos{
		EV:  EVs,
		EF:  EFs,
		Dir: "",
	}
}

func NewT_Bool(valor string) *T_Bool {
	return &T_Bool{valor: valor}
}


