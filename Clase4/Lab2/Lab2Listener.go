package main

import (
	"Lab2/Lab2G"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
)

type Lab2ListenerI struct {
	Lab2G.BaseLab2_GListener
}

func (l Lab2ListenerI) VisitTerminal(node antlr.TerminalNode) {
	//TODO implement me

}

func (l Lab2ListenerI) VisitErrorNode(node antlr.ErrorNode) {
	//TODO implement me
	panic("implement me")
}

func (l Lab2ListenerI) EnterEveryRule(ctx antlr.ParserRuleContext) {
	//TODO implement me

}

func (l Lab2ListenerI) ExitEveryRule(ctx antlr.ParserRuleContext) {
	//TODO implement me

}

func (l Lab2ListenerI) EnterSLSentencias(c *Lab2G.SLSentenciasContext) {
	//TODO implement me
	panic("implement me")
}

func (l Lab2ListenerI) EnterL_Sentencia(c *Lab2G.L_SentenciaContext) {
	//TODO implement me
	panic("implement me")
}

func (l Lab2ListenerI) EnterSConsola(c *Lab2G.SConsolaContext) {
	//TODO implement me
	panic("implement me")
}

func (l Lab2ListenerI) EnterSDeclaracion(c *Lab2G.SDeclaracionContext) {
	//TODO implement me
	panic("implement me")
}

func (l Lab2ListenerI) EnterS_Asig(c *Lab2G.S_AsigContext) {
	//TODO implement me
	panic("implement me")
}

func (l Lab2ListenerI) EnterSumAsg(c *Lab2G.SumAsgContext) {
	//TODO implement me
	panic("implement me")
}

func (l Lab2ListenerI) EnterAsig(c *Lab2G.AsigContext) {
	//TODO implement me
	panic("implement me")
}

func (l Lab2ListenerI) EnterE_Decimal(c *Lab2G.E_DecimalContext) {
	//TODO implement me
	panic("implement me")
}

func (l Lab2ListenerI) EnterE_Entero(c *Lab2G.E_EnteroContext) {
	//TODO implement me
	fmt.Println("entrando en E -> entero ", c.ENTERO().GetText())

}

func (l Lab2ListenerI) EnterE_Id(c *Lab2G.E_IdContext) {
	//TODO implement me
	panic("implement me")
}

func (l Lab2ListenerI) EnterE_Neg(c *Lab2G.E_NegContext) {
	//TODO implement me
	panic("implement me")
}

func (l Lab2ListenerI) EnterE_par(c *Lab2G.E_parContext) {
	//TODO implement me
	panic("implement me")
}

func (l Lab2ListenerI) ExitSLSentencias(c *Lab2G.SLSentenciasContext) {
	//TODO implement me
	panic("implement me")
}

func (l Lab2ListenerI) ExitL_Sentencia(c *Lab2G.L_SentenciaContext) {
	//TODO implement me
	panic("implement me")
}

func (l Lab2ListenerI) ExitSConsola(c *Lab2G.SConsolaContext) {
	//TODO implement me
	panic("implement me")
}

func (l Lab2ListenerI) ExitSDeclaracion(c *Lab2G.SDeclaracionContext) {
	//TODO implement me
	panic("implement me")
}

func (l Lab2ListenerI) ExitS_Asig(c *Lab2G.S_AsigContext) {
	//TODO implement me
	panic("implement me")
}

func (l Lab2ListenerI) ExitSumAsg(c *Lab2G.SumAsgContext) {
	//TODO implement me
	panic("implement me")
}

func (l Lab2ListenerI) ExitAsig(c *Lab2G.AsigContext) {
	//TODO implement me
	panic("implement me")
}

func (l Lab2ListenerI) ExitE_Decimal(c *Lab2G.E_DecimalContext) {
	//TODO implement me
	panic("implement me")
}

func (l Lab2ListenerI) ExitE_Entero(c *Lab2G.E_EnteroContext) {
	//TODO implement me
	fmt.Println("saliendo de entero: ", c.ENTERO().GetText()) //reducción
}

func (l Lab2ListenerI) ExitE_muldiv(c *Lab2G.E_muldivContext) {
	//TODO implement me
	fmt.Println("saliendo de E -> E / * E ") //reducción

}

func (l Lab2ListenerI) ExitE_Id(c *Lab2G.E_IdContext) {
	//TODO implement me
	panic("implement me")
}

func (l Lab2ListenerI) ExitE_Neg(c *Lab2G.E_NegContext) {
	//TODO implement me
	panic("implement me")
}

func (l Lab2ListenerI) ExitE_par(c *Lab2G.E_parContext) {
	//TODO implement me
	panic("implement me")
}

func (l Lab2ListenerI) EnterE_muldiv(c *Lab2G.E_muldivContext) {
	//deplazamiento
	fmt.Println("Entrando en #E_Muldiv")
	fmt.Println("operacion: ", c.GetOp().GetText())
}

func (l Lab2ListenerI) EnterE_sumres(c *Lab2G.E_sumresContext) {
	//c.E(1) // E -> E + E
	//c.E(0) // E-> E + E
	// deplazamiento
	fmt.Println("Entrando a E -> E (+|-) E ")
	fmt.Println(c.GetOperacion().GetText())
}

func (l Lab2ListenerI) ExitE_sumres(c *Lab2G.E_sumresContext) {
	//reducción
	fmt.Println("Saliendo de E +|- E ")
	fmt.Println("operacion", c.GetOperacion().GetText())

}
