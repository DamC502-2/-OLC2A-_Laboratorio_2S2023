package Compilador

import (
	"fmt"
	"strconv"
)

// generadores de temporales
// impresion del código

type Contexto struct {
	tmp     int
	lb      int
	display []string //GODS
}

func NewContexto() *Contexto {
	return &Contexto{
		tmp: 0,
		lb:  0,
		display: make([]string, 0,0),
	}
}

func (ctx *Contexto) PushDisplay() string {
	nuevaEtiqueta := ctx.NewEtq()
	ctx.display = append(ctx.display, nuevaEtiqueta)
	return nuevaEtiqueta
}

func (ctx *Contexto) PeekDislay() string {
	if len(ctx.display) >= 1 {
		return ctx.display[len(ctx.display)-1]
	} else {
		panic("error display vacío")
	}
}

func (ctx *Contexto) PopDisplay() string {
	etiqueta := ctx.PeekDislay()
	if len(ctx.display) > 0 {
		ctx.display = ctx.display[:len(ctx.display)-1]
	}
	return etiqueta
}

func (ctx *Contexto) Gen(out string) {
	fmt.Println(out)
}

func (ctx *Contexto) NewTemp() string {
	ctx.tmp = ctx.tmp + 1
	return "t" + strconv.Itoa(ctx.tmp)
}

func (ctx *Contexto) NewEtq() string {
	ctx.lb = ctx.lb + 1
	return "L" + strconv.Itoa(ctx.lb)
}

func (ctx *Contexto) ImprimirEtiquetas(etiquetas []string) {
	for _, etiqueta := range etiquetas {
		ctx.Gen(etiqueta + ":")
	}
}

func (ctx *Contexto) Unir(etq1 []string, etq2 []string) []string {
	for _, etq := range etq2 {
		etq1 = append(etq1, etq)
	}
	return etq1
}
