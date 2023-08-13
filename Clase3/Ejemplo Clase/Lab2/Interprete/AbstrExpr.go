package interprete

type Resultado struct {
	Valor  int
	valorS string
	Nil    bool
}

// Constructor for Resultado
func NewResultado(valor int, valorS string) *Resultado {
	o := new(Resultado)
	o.Valor = valor
	o.valorS = valorS
	o.Nil = false
	return o
}
func NewNil() *Resultado {
	return &Resultado{
		Valor:  0,
		valorS: "0",
		Nil:    true}

}

// interfaz de patrón Interpreter
type AbstrExpr interface {
	Interpretar(ctx *Contexto) *Resultado
}
