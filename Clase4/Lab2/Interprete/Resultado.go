package Interprete

type TipoE uint

const (
	Nil     TipoE = 0
	Bool    TipoE = 1
	Integer TipoE = 2
	Float   TipoE = 3
	String  TipoE = 4
)

func (t TipoE) String() string {
	switch t {
	case Nil:
		return "NIL"
	case Integer:
		return "Integer"
	case Float:
		return "Float"
	case String:
		return "String"
	case Bool:
		return "Bool"
	default:
		return "Desconocido"
	}
}

type Resultado struct {
	Valor  int
	ValorF float64
	ValorS string
	ValorB bool
	Nil    bool
	Tipo   TipoE
}

func NewResultado(valor int, valorS string) *Resultado {
	return &Resultado{Valor: valor, ValorS: valorS, Nil: false}
}

func NewFloatLiteral(valor float64) *Resultado {
	return &Resultado{
		ValorF: valor,
		Tipo:   Float,
		Valor:  0,
		ValorS: "",
		Nil:    false,
	}
}

func NewIntLiteral(valor int) *Resultado {
	return &Resultado{
		ValorF: 0.0,
		Tipo:   Integer,
		Valor:  valor,
		ValorS: "",
		Nil:    false,
	}
}

func NewStringLiteral(valor string) *Resultado {
	return &Resultado{
		ValorF: 0.0,
		Tipo:   String,
		Valor:  0,
		ValorS: valor,
		Nil:    false,
	}
}

func NewNil() *Resultado {

	return &Resultado{Valor: 0, ValorS: "", Nil: true,
		Tipo: Nil}
}
