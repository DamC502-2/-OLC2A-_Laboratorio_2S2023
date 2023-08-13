package Interprete



type Resultado struct {
	Valor  int
	ValorS string
	Nil    bool
}

func NewResultado(valor int, valorS string) *Resultado {
	return &Resultado{Valor: valor, ValorS: valorS, Nil: false}
}
func NewNil() *Resultado {

	return &Resultado{Valor: 0, ValorS: "", Nil: true}
}

// intefaz en Go
type AbstrExpr interface {
	Interpretar(ctx *Contexto) *Resultado
}

