package Compilador

import (
	"fmt"
	"strconv"
)

// generadores de temporales
// impresion del código

type Contexto struct {
	tmp int
	lb  int
}

func NewContexto() *Contexto {
	return &Contexto{
		tmp: 0,
		lb: 0,
	}
}

func (ctx *Contexto) Gen(out string) {
	fmt.Println(out)
}

func (ctx *Contexto) NewTemp() string {
	ctx.tmp = ctx.tmp + 1
	return "t" + strconv.Itoa(ctx.tmp)
}

func (ctx *Contexto)NewEtq() string {
	ctx.lb = ctx.lb + 1
	return "L" + strconv.Itoa(ctx.lb)
}

func (ctx *Contexto)ImprimirEtiquetas(etiquetas []string) {
	for _, etiqueta := range etiquetas {
		ctx.Gen(etiqueta + ":")
	}
}

func (ctx *Contexto)Unir(etq1 []string, etq2 []string) []string {
	for _, etq := range etq2 {
		etq1 = append(etq1, etq)
	}
	return etq1
}
