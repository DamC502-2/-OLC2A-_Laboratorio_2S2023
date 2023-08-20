package Interprete

import (
	"fmt"
	"strconv"
)

/*
Conversor: ayuda a convertir un objeto de tipo Result a
un tipo específico, utilizable cuando sea necesario
*/
type Conversor struct {
	ctx *Contexto
}

func NewConversor(ctx *Contexto) *Conversor {
	return &Conversor{
		ctx: ctx,
	}
}

// 10 + NIL
func (c *Conversor) Ampliar(res *Resultado, tipo TipoE) *Resultado {
	// TODO hacer deep copy
	switch tipo {
	case Nil:
		c.ctx.AddError("Conversion a a nil Ilegal.....")
		return NewNil()
	case Bool:
		panic("Conversion no soportada")
	case Integer:
		// tipo origen
		switch res.Tipo {
		case Nil:
			c.ctx.AddError("No se puede convertir de Nil a Integer")
			return NewNil()
		case Bool:
			if res.ValorB == true {
				res.Valor = 1
			} else {
				res.Valor = 0
			}
			res.Tipo = Integer
			return res
		case Integer:
			return res
		case Float:
			res.Valor = int(res.ValorF) // se trunca
			res.Tipo = Integer
			return res
		case String:
			c.ctx.AddError("No se puede convertir de String a Int")
			return NewNil()
		default:
		}
	case Float:
		//TODO
	case String:
		switch res.Tipo {
		case Bool:
			if res.ValorB {
				res.ValorS = "true"
			} else {
				res.ValorS = "false"
			}
			res.Tipo = String
			return res
		case Integer:
			res.ValorS = strconv.Itoa(res.Valor)
			res.Tipo = String
			return res
		case Float:
			res.ValorS = fmt.Sprintf("%f", res.ValorF)
			res.Tipo = String
			return res
		case String:
			return res
		default:
		}
	default:
	}
	c.ctx.AddError("Conversion illegal o no soportada de: " +
		tipo.String() + " y " + res.Tipo.String())
	return NewNil()
}

/*
case Nil:
	case Bool:
	case Integer:
	case Float:
	case String:
	default:
*/
