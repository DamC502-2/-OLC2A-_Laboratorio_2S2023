package main

import (
	"Lab1/Lab1P"
	"github.com/antlr4-go/antlr/v4"
)

type Lab1PListenerImp struct {
	Lab1P.BaseLab1_GListener
}


func (l *Lab1PListenerImp) VisitTerminal(node antlr.TerminalNode)      {

}
func (l *Lab1PListenerImp) VisitErrorNode(node antlr.ErrorNode)        {

}
func (l *Lab1PListenerImp) EnterEveryRule(ctx antlr.ParserRuleContext) {

}
func (l *Lab1PListenerImp) ExitEveryRule(ctx antlr.ParserRuleContext)  {

}
func (l *Lab1PListenerImp) EnterS(ctx *Lab1P.SContext)                 {

}
func (l *Lab1PListenerImp) ExitS(ctx *Lab1P.SContext)                  {

}
func (l *Lab1PListenerImp) EnterA0(ctx *Lab1P.A0Context)               {

}
func (l *Lab1PListenerImp) ExitA0(ctx *Lab1P.A0Context)                {

}
func (l *Lab1PListenerImp) EnterAENTERO(ctx *Lab1P.AENTEROContext)     {

}
func (l *Lab1PListenerImp) ExitAENTERO(ctx *Lab1P.AENTEROContext)      {

}
func (l *Lab1PListenerImp) EnterADECIMAL(ctx *Lab1P.ADECIMALContext)   {

}
func (l *Lab1PListenerImp) ExitADECIMAL(ctx *Lab1P.ADECIMALContext)    {

}
func (l *Lab1PListenerImp) EnterAID(ctx *Lab1P.AIDContext)             {

}
func (l *Lab1PListenerImp) ExitAID(ctx *Lab1P.AIDContext)              {

}
func (l *Lab1PListenerImp) EnterAEPSILUM(ctx *Lab1P.AEPSILUMContext)   {

}
func (l *Lab1PListenerImp) ExitAEPSILUM(ctx *Lab1P.AEPSILUMContext)    {

}
