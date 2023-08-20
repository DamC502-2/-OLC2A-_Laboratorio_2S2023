package Interprete

// intefaz en Go
type AbstrExpr interface {
	Interpretar(ctx *Contexto) *Resultado
}
