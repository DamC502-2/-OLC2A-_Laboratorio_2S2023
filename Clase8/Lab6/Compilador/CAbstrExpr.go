package Compilador

type CAbstrExpr interface {
	Compilar(ctx *Contexto) *Atributos
}
