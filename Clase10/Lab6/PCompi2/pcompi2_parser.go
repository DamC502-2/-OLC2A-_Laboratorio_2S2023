// Code generated from PCompi2.g4 by ANTLR 4.13.0. DO NOT EDIT.

package PCompi2 // PCompi2
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

//El código que incrusta al principio
// importaciones
import "Lab6/gen"

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type PCompi2Parser struct {
	*antlr.BaseParser
}

var PCompi2ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func pcompi2ParserInit() {
	staticData := &PCompi2ParserStaticData
	staticData.LiteralNames = []string{
		"", "'!'", "'AND'", "'OR'", "'true'", "'false'", "'-'", "'*'", "'/'",
		"'+'", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='", "')'", "'('",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "PARDER",
		"PARIZQ", "ENTERO", "DECIMAL", "ENBLANCO", "ID",
	}
	staticData.RuleNames = []string{
		"s", "cond", "marcador", "marcador2", "expr", "oprel",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 21, 104, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 3, 1, 36, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 50, 8, 1, 10, 1, 12, 1, 53, 9, 1, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 73, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 85, 8, 4, 10, 4, 12, 4, 88, 9, 4, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3,
		5, 102, 8, 5, 1, 5, 0, 2, 2, 8, 6, 0, 2, 4, 6, 8, 10, 0, 2, 1, 0, 7, 8,
		2, 0, 6, 6, 9, 9, 113, 0, 12, 1, 0, 0, 0, 2, 35, 1, 0, 0, 0, 4, 54, 1,
		0, 0, 0, 6, 56, 1, 0, 0, 0, 8, 72, 1, 0, 0, 0, 10, 101, 1, 0, 0, 0, 12,
		13, 3, 2, 1, 0, 13, 14, 5, 0, 0, 1, 14, 15, 6, 0, -1, 0, 15, 1, 1, 0, 0,
		0, 16, 17, 6, 1, -1, 0, 17, 18, 5, 1, 0, 0, 18, 19, 3, 2, 1, 7, 19, 20,
		6, 1, -1, 0, 20, 36, 1, 0, 0, 0, 21, 22, 3, 8, 4, 0, 22, 23, 3, 10, 5,
		0, 23, 24, 3, 8, 4, 0, 24, 25, 6, 1, -1, 0, 25, 36, 1, 0, 0, 0, 26, 27,
		5, 17, 0, 0, 27, 28, 3, 2, 1, 0, 28, 29, 5, 16, 0, 0, 29, 30, 6, 1, -1,
		0, 30, 36, 1, 0, 0, 0, 31, 32, 5, 4, 0, 0, 32, 36, 6, 1, -1, 0, 33, 34,
		5, 5, 0, 0, 34, 36, 6, 1, -1, 0, 35, 16, 1, 0, 0, 0, 35, 21, 1, 0, 0, 0,
		35, 26, 1, 0, 0, 0, 35, 31, 1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 36, 51, 1,
		0, 0, 0, 37, 38, 10, 6, 0, 0, 38, 39, 5, 2, 0, 0, 39, 40, 3, 4, 2, 0, 40,
		41, 3, 2, 1, 7, 41, 42, 6, 1, -1, 0, 42, 50, 1, 0, 0, 0, 43, 44, 10, 5,
		0, 0, 44, 45, 5, 3, 0, 0, 45, 46, 3, 6, 3, 0, 46, 47, 3, 2, 1, 6, 47, 48,
		6, 1, -1, 0, 48, 50, 1, 0, 0, 0, 49, 37, 1, 0, 0, 0, 49, 43, 1, 0, 0, 0,
		50, 53, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 3, 1, 0,
		0, 0, 53, 51, 1, 0, 0, 0, 54, 55, 6, 2, -1, 0, 55, 5, 1, 0, 0, 0, 56, 57,
		6, 3, -1, 0, 57, 7, 1, 0, 0, 0, 58, 59, 6, 4, -1, 0, 59, 60, 5, 6, 0, 0,
		60, 61, 3, 8, 4, 6, 61, 62, 6, 4, -1, 0, 62, 73, 1, 0, 0, 0, 63, 64, 5,
		17, 0, 0, 64, 65, 3, 8, 4, 0, 65, 66, 5, 16, 0, 0, 66, 67, 6, 4, -1, 0,
		67, 73, 1, 0, 0, 0, 68, 69, 5, 18, 0, 0, 69, 73, 6, 4, -1, 0, 70, 71, 5,
		21, 0, 0, 71, 73, 6, 4, -1, 0, 72, 58, 1, 0, 0, 0, 72, 63, 1, 0, 0, 0,
		72, 68, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 73, 86, 1, 0, 0, 0, 74, 75, 10,
		5, 0, 0, 75, 76, 7, 0, 0, 0, 76, 77, 3, 8, 4, 6, 77, 78, 6, 4, -1, 0, 78,
		85, 1, 0, 0, 0, 79, 80, 10, 4, 0, 0, 80, 81, 7, 1, 0, 0, 81, 82, 3, 8,
		4, 5, 82, 83, 6, 4, -1, 0, 83, 85, 1, 0, 0, 0, 84, 74, 1, 0, 0, 0, 84,
		79, 1, 0, 0, 0, 85, 88, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 86, 87, 1, 0, 0,
		0, 87, 9, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 89, 90, 5, 10, 0, 0, 90, 102,
		6, 5, -1, 0, 91, 92, 5, 11, 0, 0, 92, 102, 6, 5, -1, 0, 93, 94, 5, 12,
		0, 0, 94, 102, 6, 5, -1, 0, 95, 96, 5, 13, 0, 0, 96, 102, 6, 5, -1, 0,
		97, 98, 5, 14, 0, 0, 98, 102, 6, 5, -1, 0, 99, 100, 5, 15, 0, 0, 100, 102,
		6, 5, -1, 0, 101, 89, 1, 0, 0, 0, 101, 91, 1, 0, 0, 0, 101, 93, 1, 0, 0,
		0, 101, 95, 1, 0, 0, 0, 101, 97, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 102,
		11, 1, 0, 0, 0, 7, 35, 49, 51, 72, 84, 86, 101,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// PCompi2ParserInit initializes any static state used to implement PCompi2Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPCompi2Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PCompi2ParserInit() {
	staticData := &PCompi2ParserStaticData
	staticData.once.Do(pcompi2ParserInit)
}

// NewPCompi2Parser produces a new parser instance for the optional input antlr.TokenStream.
func NewPCompi2Parser(input antlr.TokenStream) *PCompi2Parser {
	PCompi2ParserInit()
	this := new(PCompi2Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &PCompi2ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "PCompi2.g4"

	return this
}

// Note that '@members' cannot be changed now, but this should have been 'globals'
// If you are looking to have variables for each instance, use '@structmembers'

// instancias de variables
// contador de temporales
var temp int = 1

// PCompi2Parser tokens.
const (
	PCompi2ParserEOF      = antlr.TokenEOF
	PCompi2ParserT__0     = 1
	PCompi2ParserT__1     = 2
	PCompi2ParserT__2     = 3
	PCompi2ParserT__3     = 4
	PCompi2ParserT__4     = 5
	PCompi2ParserT__5     = 6
	PCompi2ParserT__6     = 7
	PCompi2ParserT__7     = 8
	PCompi2ParserT__8     = 9
	PCompi2ParserT__9     = 10
	PCompi2ParserT__10    = 11
	PCompi2ParserT__11    = 12
	PCompi2ParserT__12    = 13
	PCompi2ParserT__13    = 14
	PCompi2ParserT__14    = 15
	PCompi2ParserPARDER   = 16
	PCompi2ParserPARIZQ   = 17
	PCompi2ParserENTERO   = 18
	PCompi2ParserDECIMAL  = 19
	PCompi2ParserENBLANCO = 20
	PCompi2ParserID       = 21
)

// PCompi2Parser rules.
const (
	PCompi2ParserRULE_s         = 0
	PCompi2ParserRULE_cond      = 1
	PCompi2ParserRULE_marcador  = 2
	PCompi2ParserRULE_marcador2 = 3
	PCompi2ParserRULE_expr      = 4
	PCompi2ParserRULE_oprel     = 5
)

// ISContext is an interface to support dynamic dispatch.
type ISContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetC1 returns the c1 rule contexts.
	GetC1() ICondContext

	// SetC1 sets the c1 rule contexts.
	SetC1(ICondContext)

	// Getter signatures
	EOF() antlr.TerminalNode
	Cond() ICondContext

	// IsSContext differentiates from other interfaces.
	IsSContext()
}

type SContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	c1     ICondContext
}

func NewEmptySContext() *SContext {
	var p = new(SContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PCompi2ParserRULE_s
	return p
}

func InitEmptySContext(p *SContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PCompi2ParserRULE_s
}

func (*SContext) IsSContext() {}

func NewSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SContext {
	var p = new(SContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PCompi2ParserRULE_s

	return p
}

func (s *SContext) GetParser() antlr.Parser { return s.parser }

func (s *SContext) GetC1() ICondContext { return s.c1 }

func (s *SContext) SetC1(v ICondContext) { s.c1 = v }

func (s *SContext) EOF() antlr.TerminalNode {
	return s.GetToken(PCompi2ParserEOF, 0)
}

func (s *SContext) Cond() ICondContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICondContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICondContext)
}

func (s *SContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2Listener); ok {
		listenerT.EnterS(s)
	}
}

func (s *SContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2Listener); ok {
		listenerT.ExitS(s)
	}
}

func (s *SContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PCompi2Visitor:
		return t.VisitS(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PCompi2Parser) S() (localctx ISContext) {
	localctx = NewSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PCompi2ParserRULE_s)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(12)

		var _x = p.cond(0)

		localctx.(*SContext).c1 = _x
	}
	{
		p.SetState(13)
		p.Match(PCompi2ParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	gen.Gen("//etiquetas verdaderas")
	gen.ImprimirEtiquetas(localctx.(*SContext).GetC1().GetEV())
	gen.Gen("//etiquetas falsas")
	gen.ImprimirEtiquetas(localctx.(*SContext).GetC1().GetEF())

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICondContext is an interface to support dynamic dispatch.
type ICondContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetC1 returns the c1 rule contexts.
	GetC1() ICondContext

	// GetC returns the c rule contexts.
	GetC() ICondContext

	// GetE1 returns the e1 rule contexts.
	GetE1() IExprContext

	// GetOpr returns the opr rule contexts.
	GetOpr() IOprelContext

	// GetE2 returns the e2 rule contexts.
	GetE2() IExprContext

	// GetC2 returns the c2 rule contexts.
	GetC2() ICondContext

	// SetC1 sets the c1 rule contexts.
	SetC1(ICondContext)

	// SetC sets the c rule contexts.
	SetC(ICondContext)

	// SetE1 sets the e1 rule contexts.
	SetE1(IExprContext)

	// SetOpr sets the opr rule contexts.
	SetOpr(IOprelContext)

	// SetE2 sets the e2 rule contexts.
	SetE2(IExprContext)

	// SetC2 sets the c2 rule contexts.
	SetC2(ICondContext)

	// GetEV returns the EV attribute.
	GetEV() []string

	// GetEF returns the EF attribute.
	GetEF() []string

	// SetEV sets the EV attribute.
	SetEV([]string)

	// SetEF sets the EF attribute.
	SetEF([]string)

	// Getter signatures
	AllCond() []ICondContext
	Cond(i int) ICondContext
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	Oprel() IOprelContext
	PARIZQ() antlr.TerminalNode
	PARDER() antlr.TerminalNode
	Marcador() IMarcadorContext
	Marcador2() IMarcador2Context

	// IsCondContext differentiates from other interfaces.
	IsCondContext()
}

type CondContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	EV     []string
	EF     []string
	c1     ICondContext
	op     antlr.Token
	c      ICondContext
	e1     IExprContext
	opr    IOprelContext
	e2     IExprContext
	c2     ICondContext
}

func NewEmptyCondContext() *CondContext {
	var p = new(CondContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PCompi2ParserRULE_cond
	return p
}

func InitEmptyCondContext(p *CondContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PCompi2ParserRULE_cond
}

func (*CondContext) IsCondContext() {}

func NewCondContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CondContext {
	var p = new(CondContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PCompi2ParserRULE_cond

	return p
}

func (s *CondContext) GetParser() antlr.Parser { return s.parser }

func (s *CondContext) GetOp() antlr.Token { return s.op }

func (s *CondContext) SetOp(v antlr.Token) { s.op = v }

func (s *CondContext) GetC1() ICondContext { return s.c1 }

func (s *CondContext) GetC() ICondContext { return s.c }

func (s *CondContext) GetE1() IExprContext { return s.e1 }

func (s *CondContext) GetOpr() IOprelContext { return s.opr }

func (s *CondContext) GetE2() IExprContext { return s.e2 }

func (s *CondContext) GetC2() ICondContext { return s.c2 }

func (s *CondContext) SetC1(v ICondContext) { s.c1 = v }

func (s *CondContext) SetC(v ICondContext) { s.c = v }

func (s *CondContext) SetE1(v IExprContext) { s.e1 = v }

func (s *CondContext) SetOpr(v IOprelContext) { s.opr = v }

func (s *CondContext) SetE2(v IExprContext) { s.e2 = v }

func (s *CondContext) SetC2(v ICondContext) { s.c2 = v }

func (s *CondContext) GetEV() []string { return s.EV }

func (s *CondContext) GetEF() []string { return s.EF }

func (s *CondContext) SetEV(v []string) { s.EV = v }

func (s *CondContext) SetEF(v []string) { s.EF = v }

func (s *CondContext) AllCond() []ICondContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICondContext); ok {
			len++
		}
	}

	tst := make([]ICondContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICondContext); ok {
			tst[i] = t.(ICondContext)
			i++
		}
	}

	return tst
}

func (s *CondContext) Cond(i int) ICondContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICondContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICondContext)
}

func (s *CondContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *CondContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CondContext) Oprel() IOprelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOprelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOprelContext)
}

func (s *CondContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(PCompi2ParserPARIZQ, 0)
}

func (s *CondContext) PARDER() antlr.TerminalNode {
	return s.GetToken(PCompi2ParserPARDER, 0)
}

func (s *CondContext) Marcador() IMarcadorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMarcadorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMarcadorContext)
}

func (s *CondContext) Marcador2() IMarcador2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMarcador2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMarcador2Context)
}

func (s *CondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2Listener); ok {
		listenerT.EnterCond(s)
	}
}

func (s *CondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2Listener); ok {
		listenerT.ExitCond(s)
	}
}

func (s *CondContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PCompi2Visitor:
		return t.VisitCond(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PCompi2Parser) Cond() (localctx ICondContext) {
	return p.cond(0)
}

func (p *PCompi2Parser) cond(_p int) (localctx ICondContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewCondContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ICondContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, PCompi2ParserRULE_cond, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(17)

			var _m = p.Match(PCompi2ParserT__0)

			localctx.(*CondContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(18)

			var _x = p.cond(7)

			localctx.(*CondContext).c = _x
		}

		localctx.(*CondContext).EV = localctx.(*CondContext).GetC().GetEF()
		localctx.(*CondContext).EF = localctx.(*CondContext).GetC().GetEV()

	case 2:
		{
			p.SetState(21)

			var _x = p.expr(0)

			localctx.(*CondContext).e1 = _x
		}
		{
			p.SetState(22)

			var _x = p.Oprel()

			localctx.(*CondContext).opr = _x
		}
		{
			p.SetState(23)

			var _x = p.expr(0)

			localctx.(*CondContext).e2 = _x
		}

		localctx.(*CondContext).EV = append(localctx.(*CondContext).EV, gen.NewEtq())
		localctx.(*CondContext).EF = append(localctx.(*CondContext).EF, gen.NewEtq())
		var cad string
		cad = localctx.(*CondContext).GetE1().GetDir() + " " + localctx.(*CondContext).GetOpr().GetOp() + " " + localctx.(*CondContext).GetE2().GetDir()
		gen.Gen("if " + cad + " then goto " + localctx.(*CondContext).EV[0])
		gen.Gen("goto " + localctx.(*CondContext).EF[0])

	case 3:
		{
			p.SetState(26)
			p.Match(PCompi2ParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(27)

			var _x = p.cond(0)

			localctx.(*CondContext).c = _x
		}
		{
			p.SetState(28)
			p.Match(PCompi2ParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*CondContext).EV = localctx.(*CondContext).GetC().GetEV()
		localctx.(*CondContext).EF = localctx.(*CondContext).GetC().GetEF()

	case 4:
		{
			p.SetState(31)
			p.Match(PCompi2ParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*CondContext).EV = append(localctx.(*CondContext).EV, gen.NewEtq())
		localctx.(*CondContext).EF = append(localctx.(*CondContext).EF, gen.NewEtq())
		gen.Gen("goto " + localctx.(*CondContext).EV[0])

	case 5:
		{
			p.SetState(33)
			p.Match(PCompi2ParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*CondContext).EV = append(localctx.(*CondContext).EV, gen.NewEtq())
		localctx.(*CondContext).EF = append(localctx.(*CondContext).EF, gen.NewEtq())
		gen.Gen("goto " + localctx.(*CondContext).EF[0])

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(49)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewCondContext(p, _parentctx, _parentState)
				localctx.(*CondContext).c1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, PCompi2ParserRULE_cond)
				p.SetState(37)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(38)

					var _m = p.Match(PCompi2ParserT__1)

					localctx.(*CondContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(39)
					p.Marcador(localctx.(*CondContext).GetC1().GetEV())
				}
				{
					p.SetState(40)

					var _x = p.cond(7)

					localctx.(*CondContext).c2 = _x
				}

				localctx.(*CondContext).EV = localctx.(*CondContext).GetC2().GetEV()
				localctx.(*CondContext).EF = gen.Unir(localctx.(*CondContext).GetC1().GetEF(), localctx.(*CondContext).GetC2().GetEF())

			case 2:
				localctx = NewCondContext(p, _parentctx, _parentState)
				localctx.(*CondContext).c1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, PCompi2ParserRULE_cond)
				p.SetState(43)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(44)

					var _m = p.Match(PCompi2ParserT__2)

					localctx.(*CondContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(45)
					p.Marcador2(localctx.(*CondContext).GetC1().GetEF())
				}
				{
					p.SetState(46)

					var _x = p.cond(6)

					localctx.(*CondContext).c2 = _x
				}

				localctx.(*CondContext).EF = localctx.(*CondContext).GetC2().GetEF()
				localctx.(*CondContext).EV = gen.Unir(localctx.(*CondContext).GetC1().GetEV(), localctx.(*CondContext).GetC2().GetEV())

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMarcadorContext is an interface to support dynamic dispatch.
type IMarcadorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetEV returns the EV attribute.
	GetEV() []string

	// SetEV sets the EV attribute.
	SetEV([]string)

	// IsMarcadorContext differentiates from other interfaces.
	IsMarcadorContext()
}

type MarcadorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	EV     []string
}

func NewEmptyMarcadorContext() *MarcadorContext {
	var p = new(MarcadorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PCompi2ParserRULE_marcador
	return p
}

func InitEmptyMarcadorContext(p *MarcadorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PCompi2ParserRULE_marcador
}

func (*MarcadorContext) IsMarcadorContext() {}

func NewMarcadorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int, EV []string) *MarcadorContext {
	var p = new(MarcadorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PCompi2ParserRULE_marcador

	p.EV = EV

	return p
}

func (s *MarcadorContext) GetParser() antlr.Parser { return s.parser }

func (s *MarcadorContext) GetEV() []string { return s.EV }

func (s *MarcadorContext) SetEV(v []string) { s.EV = v }

func (s *MarcadorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MarcadorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MarcadorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2Listener); ok {
		listenerT.EnterMarcador(s)
	}
}

func (s *MarcadorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2Listener); ok {
		listenerT.ExitMarcador(s)
	}
}

func (s *MarcadorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PCompi2Visitor:
		return t.VisitMarcador(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PCompi2Parser) Marcador(EV []string) (localctx IMarcadorContext) {
	localctx = NewMarcadorContext(p, p.GetParserRuleContext(), p.GetState(), EV)
	p.EnterRule(localctx, 4, PCompi2ParserRULE_marcador)
	p.EnterOuterAlt(localctx, 1)

	// imprimir cond.EV
	gen.ImprimirEtiquetas(EV)

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMarcador2Context is an interface to support dynamic dispatch.
type IMarcador2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetEF returns the EF attribute.
	GetEF() []string

	// SetEF sets the EF attribute.
	SetEF([]string)

	// IsMarcador2Context differentiates from other interfaces.
	IsMarcador2Context()
}

type Marcador2Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	EF     []string
}

func NewEmptyMarcador2Context() *Marcador2Context {
	var p = new(Marcador2Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PCompi2ParserRULE_marcador2
	return p
}

func InitEmptyMarcador2Context(p *Marcador2Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PCompi2ParserRULE_marcador2
}

func (*Marcador2Context) IsMarcador2Context() {}

func NewMarcador2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int, EF []string) *Marcador2Context {
	var p = new(Marcador2Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PCompi2ParserRULE_marcador2

	p.EF = EF

	return p
}

func (s *Marcador2Context) GetParser() antlr.Parser { return s.parser }

func (s *Marcador2Context) GetEF() []string { return s.EF }

func (s *Marcador2Context) SetEF(v []string) { s.EF = v }

func (s *Marcador2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Marcador2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Marcador2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2Listener); ok {
		listenerT.EnterMarcador2(s)
	}
}

func (s *Marcador2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2Listener); ok {
		listenerT.ExitMarcador2(s)
	}
}

func (s *Marcador2Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PCompi2Visitor:
		return t.VisitMarcador2(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PCompi2Parser) Marcador2(EF []string) (localctx IMarcador2Context) {
	localctx = NewMarcador2Context(p, p.GetParserRuleContext(), p.GetState(), EF)
	p.EnterRule(localctx, 6, PCompi2ParserRULE_marcador2)
	p.EnterOuterAlt(localctx, 1)

	// imprimir cond.EF
	gen.ImprimirEtiquetas(EF)

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// GetN returns the n token.
	GetN() antlr.Token

	// GetId returns the id token.
	GetId() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// SetN sets the n token.
	SetN(antlr.Token)

	// SetId sets the id token.
	SetId(antlr.Token)

	// GetE1 returns the e1 rule contexts.
	GetE1() IExprContext

	// GetE2 returns the e2 rule contexts.
	GetE2() IExprContext

	// SetE1 sets the e1 rule contexts.
	SetE1(IExprContext)

	// SetE2 sets the e2 rule contexts.
	SetE2(IExprContext)

	// GetDir returns the dir attribute.
	GetDir() string

	// SetDir sets the dir attribute.
	SetDir(string)

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	PARIZQ() antlr.TerminalNode
	PARDER() antlr.TerminalNode
	ENTERO() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	dir    string
	e1     IExprContext
	op     antlr.Token
	n      antlr.Token
	id     antlr.Token
	e2     IExprContext
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PCompi2ParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PCompi2ParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PCompi2ParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) GetOp() antlr.Token { return s.op }

func (s *ExprContext) GetN() antlr.Token { return s.n }

func (s *ExprContext) GetId() antlr.Token { return s.id }

func (s *ExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprContext) SetN(v antlr.Token) { s.n = v }

func (s *ExprContext) SetId(v antlr.Token) { s.id = v }

func (s *ExprContext) GetE1() IExprContext { return s.e1 }

func (s *ExprContext) GetE2() IExprContext { return s.e2 }

func (s *ExprContext) SetE1(v IExprContext) { s.e1 = v }

func (s *ExprContext) SetE2(v IExprContext) { s.e2 = v }

func (s *ExprContext) GetDir() string { return s.dir }

func (s *ExprContext) SetDir(v string) { s.dir = v }

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(PCompi2ParserPARIZQ, 0)
}

func (s *ExprContext) PARDER() antlr.TerminalNode {
	return s.GetToken(PCompi2ParserPARDER, 0)
}

func (s *ExprContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(PCompi2ParserENTERO, 0)
}

func (s *ExprContext) ID() antlr.TerminalNode {
	return s.GetToken(PCompi2ParserID, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2Listener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2Listener); ok {
		listenerT.ExitExpr(s)
	}
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PCompi2Visitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PCompi2Parser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *PCompi2Parser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, PCompi2ParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PCompi2ParserT__5:
		{
			p.SetState(59)

			var _m = p.Match(PCompi2ParserT__5)

			localctx.(*ExprContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(60)

			var _x = p.expr(6)

			localctx.(*ExprContext).e1 = _x
		}
		localctx.(*ExprContext).dir = gen.NewTemp()
		gen.Gen(localctx.(*ExprContext).dir + " = -1 ")
		gen.Gen(localctx.(*ExprContext).dir + " = " + localctx.(*ExprContext).dir + " * " + localctx.(*ExprContext).GetE1().GetDir())

	case PCompi2ParserPARIZQ:
		{
			p.SetState(63)
			p.Match(PCompi2ParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(64)

			var _x = p.expr(0)

			localctx.(*ExprContext).e1 = _x
		}
		{
			p.SetState(65)
			p.Match(PCompi2ParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*ExprContext).dir = localctx.(*ExprContext).GetE1().GetDir()
		// E.dir = E1.dir

	case PCompi2ParserENTERO:
		{
			p.SetState(68)

			var _m = p.Match(PCompi2ParserENTERO)

			localctx.(*ExprContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).dir = (func() string {
			if localctx.(*ExprContext).GetN() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).GetN().GetText()
			}
		}())

	case PCompi2ParserID:
		{
			p.SetState(70)

			var _m = p.Match(PCompi2ParserID)

			localctx.(*ExprContext).id = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).dir = (func() string {
			if localctx.(*ExprContext).GetId() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).GetId().GetText()
			}
		}())

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(84)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, PCompi2ParserRULE_expr)
				p.SetState(74)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(75)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == PCompi2ParserT__6 || _la == PCompi2ParserT__7) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(76)

					var _x = p.expr(6)

					localctx.(*ExprContext).e2 = _x
				}

				localctx.(*ExprContext).dir = gen.NewTemp()
				gen.Gen(localctx.(*ExprContext).dir + " = " + localctx.(*ExprContext).GetE1().GetDir() + " " + (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()) + " " + localctx.(*ExprContext).GetE2().GetDir())

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, PCompi2ParserRULE_expr)
				p.SetState(79)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(80)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == PCompi2ParserT__5 || _la == PCompi2ParserT__8) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(81)

					var _x = p.expr(5)

					localctx.(*ExprContext).e2 = _x
				}

				localctx.(*ExprContext).dir = gen.NewTemp()
				gen.Gen(localctx.(*ExprContext).dir + " = " + localctx.(*ExprContext).GetE1().GetDir() + " " + (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()) + " " + localctx.(*ExprContext).GetE2().GetDir())

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOprelContext is an interface to support dynamic dispatch.
type IOprelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOpe returns the ope token.
	GetOpe() antlr.Token

	// SetOpe sets the ope token.
	SetOpe(antlr.Token)

	// GetOp returns the op attribute.
	GetOp() string

	// SetOp sets the op attribute.
	SetOp(string)

	// IsOprelContext differentiates from other interfaces.
	IsOprelContext()
}

type OprelContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     string
	ope    antlr.Token
}

func NewEmptyOprelContext() *OprelContext {
	var p = new(OprelContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PCompi2ParserRULE_oprel
	return p
}

func InitEmptyOprelContext(p *OprelContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PCompi2ParserRULE_oprel
}

func (*OprelContext) IsOprelContext() {}

func NewOprelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OprelContext {
	var p = new(OprelContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PCompi2ParserRULE_oprel

	return p
}

func (s *OprelContext) GetParser() antlr.Parser { return s.parser }

func (s *OprelContext) GetOpe() antlr.Token { return s.ope }

func (s *OprelContext) SetOpe(v antlr.Token) { s.ope = v }

func (s *OprelContext) GetOp() string { return s.op }

func (s *OprelContext) SetOp(v string) { s.op = v }

func (s *OprelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OprelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OprelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2Listener); ok {
		listenerT.EnterOprel(s)
	}
}

func (s *OprelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2Listener); ok {
		listenerT.ExitOprel(s)
	}
}

func (s *OprelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PCompi2Visitor:
		return t.VisitOprel(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PCompi2Parser) Oprel() (localctx IOprelContext) {
	localctx = NewOprelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PCompi2ParserRULE_oprel)
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PCompi2ParserT__9:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(89)

			var _m = p.Match(PCompi2ParserT__9)

			localctx.(*OprelContext).ope = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*OprelContext).op = (func() string {
			if localctx.(*OprelContext).GetOpe() == nil {
				return ""
			} else {
				return localctx.(*OprelContext).GetOpe().GetText()
			}
		}())

	case PCompi2ParserT__10:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(91)

			var _m = p.Match(PCompi2ParserT__10)

			localctx.(*OprelContext).ope = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*OprelContext).op = (func() string {
			if localctx.(*OprelContext).GetOpe() == nil {
				return ""
			} else {
				return localctx.(*OprelContext).GetOpe().GetText()
			}
		}())

	case PCompi2ParserT__11:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(93)

			var _m = p.Match(PCompi2ParserT__11)

			localctx.(*OprelContext).ope = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*OprelContext).op = (func() string {
			if localctx.(*OprelContext).GetOpe() == nil {
				return ""
			} else {
				return localctx.(*OprelContext).GetOpe().GetText()
			}
		}())

	case PCompi2ParserT__12:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(95)

			var _m = p.Match(PCompi2ParserT__12)

			localctx.(*OprelContext).ope = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*OprelContext).op = (func() string {
			if localctx.(*OprelContext).GetOpe() == nil {
				return ""
			} else {
				return localctx.(*OprelContext).GetOpe().GetText()
			}
		}())

	case PCompi2ParserT__13:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(97)

			var _m = p.Match(PCompi2ParserT__13)

			localctx.(*OprelContext).ope = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*OprelContext).op = (func() string {
			if localctx.(*OprelContext).GetOpe() == nil {
				return ""
			} else {
				return localctx.(*OprelContext).GetOpe().GetText()
			}
		}())

	case PCompi2ParserT__14:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(99)

			var _m = p.Match(PCompi2ParserT__14)

			localctx.(*OprelContext).ope = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*OprelContext).op = (func() string {
			if localctx.(*OprelContext).GetOpe() == nil {
				return ""
			} else {
				return localctx.(*OprelContext).GetOpe().GetText()
			}
		}())

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *PCompi2Parser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *CondContext = nil
		if localctx != nil {
			t = localctx.(*CondContext)
		}
		return p.Cond_Sempred(t, predIndex)

	case 4:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *PCompi2Parser) Cond_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PCompi2Parser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
