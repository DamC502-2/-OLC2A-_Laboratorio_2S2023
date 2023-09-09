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
		"", "'*'", "'/'", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='", "')'",
		"'('",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "PARDER", "PARIZQ", "ENTERO", "DECIMAL",
		"ENBLANCO", "ID",
	}
	staticData.RuleNames = []string{
		"s", "cond", "expr", "oprel",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 14, 52, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 19, 8, 1, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 26, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5,
		2, 33, 8, 2, 10, 2, 12, 2, 36, 9, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 50, 8, 3, 1, 3, 0, 1, 4, 4, 0,
		2, 4, 6, 0, 1, 1, 0, 1, 2, 55, 0, 8, 1, 0, 0, 0, 2, 18, 1, 0, 0, 0, 4,
		25, 1, 0, 0, 0, 6, 49, 1, 0, 0, 0, 8, 9, 3, 2, 1, 0, 9, 10, 5, 0, 0, 1,
		10, 11, 6, 0, -1, 0, 11, 1, 1, 0, 0, 0, 12, 19, 1, 0, 0, 0, 13, 14, 3,
		4, 2, 0, 14, 15, 3, 6, 3, 0, 15, 16, 3, 4, 2, 0, 16, 17, 6, 1, -1, 0, 17,
		19, 1, 0, 0, 0, 18, 12, 1, 0, 0, 0, 18, 13, 1, 0, 0, 0, 19, 3, 1, 0, 0,
		0, 20, 21, 6, 2, -1, 0, 21, 22, 5, 11, 0, 0, 22, 26, 6, 2, -1, 0, 23, 24,
		5, 14, 0, 0, 24, 26, 6, 2, -1, 0, 25, 20, 1, 0, 0, 0, 25, 23, 1, 0, 0,
		0, 26, 34, 1, 0, 0, 0, 27, 28, 10, 3, 0, 0, 28, 29, 7, 0, 0, 0, 29, 30,
		3, 4, 2, 4, 30, 31, 6, 2, -1, 0, 31, 33, 1, 0, 0, 0, 32, 27, 1, 0, 0, 0,
		33, 36, 1, 0, 0, 0, 34, 32, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 5, 1, 0,
		0, 0, 36, 34, 1, 0, 0, 0, 37, 38, 5, 3, 0, 0, 38, 50, 6, 3, -1, 0, 39,
		40, 5, 4, 0, 0, 40, 50, 6, 3, -1, 0, 41, 42, 5, 5, 0, 0, 42, 50, 6, 3,
		-1, 0, 43, 44, 5, 6, 0, 0, 44, 50, 6, 3, -1, 0, 45, 46, 5, 7, 0, 0, 46,
		50, 6, 3, -1, 0, 47, 48, 5, 8, 0, 0, 48, 50, 6, 3, -1, 0, 49, 37, 1, 0,
		0, 0, 49, 39, 1, 0, 0, 0, 49, 41, 1, 0, 0, 0, 49, 43, 1, 0, 0, 0, 49, 45,
		1, 0, 0, 0, 49, 47, 1, 0, 0, 0, 50, 7, 1, 0, 0, 0, 4, 18, 25, 34, 49,
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
	PCompi2ParserPARDER   = 9
	PCompi2ParserPARIZQ   = 10
	PCompi2ParserENTERO   = 11
	PCompi2ParserDECIMAL  = 12
	PCompi2ParserENBLANCO = 13
	PCompi2ParserID       = 14
)

// PCompi2Parser rules.
const (
	PCompi2ParserRULE_s     = 0
	PCompi2ParserRULE_cond  = 1
	PCompi2ParserRULE_expr  = 2
	PCompi2ParserRULE_oprel = 3
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
		p.SetState(8)

		var _x = p.Cond()

		localctx.(*SContext).c1 = _x
	}
	{
		p.SetState(9)
		p.Match(PCompi2ParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICondContext is an interface to support dynamic dispatch.
type ICondContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetE1 returns the e1 rule contexts.
	GetE1() IExprContext

	// GetOpr returns the opr rule contexts.
	GetOpr() IOprelContext

	// GetE2 returns the e2 rule contexts.
	GetE2() IExprContext

	// SetE1 sets the e1 rule contexts.
	SetE1(IExprContext)

	// SetOpr sets the opr rule contexts.
	SetOpr(IOprelContext)

	// SetE2 sets the e2 rule contexts.
	SetE2(IExprContext)

	// GetEV returns the EV attribute.
	GetEV() []string

	// GetEF returns the EF attribute.
	GetEF() []string

	// SetEV sets the EV attribute.
	SetEV([]string)

	// SetEF sets the EF attribute.
	SetEF([]string)

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	Oprel() IOprelContext

	// IsCondContext differentiates from other interfaces.
	IsCondContext()
}

type CondContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	EV     []string
	EF     []string
	e1     IExprContext
	opr    IOprelContext
	e2     IExprContext
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

func (s *CondContext) GetE1() IExprContext { return s.e1 }

func (s *CondContext) GetOpr() IOprelContext { return s.opr }

func (s *CondContext) GetE2() IExprContext { return s.e2 }

func (s *CondContext) SetE1(v IExprContext) { s.e1 = v }

func (s *CondContext) SetOpr(v IOprelContext) { s.opr = v }

func (s *CondContext) SetE2(v IExprContext) { s.e2 = v }

func (s *CondContext) GetEV() []string { return s.EV }

func (s *CondContext) GetEF() []string { return s.EF }

func (s *CondContext) SetEV(v []string) { s.EV = v }

func (s *CondContext) SetEF(v []string) { s.EF = v }

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
	localctx = NewCondContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PCompi2ParserRULE_cond)
	p.SetState(18)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PCompi2ParserEOF:
		p.EnterOuterAlt(localctx, 1)

	case PCompi2ParserENTERO, PCompi2ParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(13)

			var _x = p.expr(0)

			localctx.(*CondContext).e1 = _x
		}
		{
			p.SetState(14)

			var _x = p.Oprel()

			localctx.(*CondContext).opr = _x
		}
		{
			p.SetState(15)

			var _x = p.expr(0)

			localctx.(*CondContext).e2 = _x
		}

		localctx.(*CondContext).EV = append(localctx.(*CondContext).EV, gen.NewEtq())
		localctx.(*CondContext).EF = append(localctx.(*CondContext).EF, gen.NewEtq())
		var cad string
		cad = localctx.(*CondContext).GetE1().GetDir() + " " + localctx.(*CondContext).GetOpr().GetOp() + " " + localctx.(*CondContext).GetE2().GetDir()
		gen.Gen("if" + cad + " then goto " + localctx.(*CondContext).EV[0])
		gen.Gen("goto " + localctx.(*CondContext).EF[0])

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

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetN returns the n token.
	GetN() antlr.Token

	// GetId returns the id token.
	GetId() antlr.Token

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetN sets the n token.
	SetN(antlr.Token)

	// SetId sets the id token.
	SetId(antlr.Token)

	// SetOp sets the op token.
	SetOp(antlr.Token)

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
	ENTERO() antlr.TerminalNode
	ID() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	dir    string
	e1     IExprContext
	n      antlr.Token
	id     antlr.Token
	op     antlr.Token
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

func (s *ExprContext) GetN() antlr.Token { return s.n }

func (s *ExprContext) GetId() antlr.Token { return s.id }

func (s *ExprContext) GetOp() antlr.Token { return s.op }

func (s *ExprContext) SetN(v antlr.Token) { s.n = v }

func (s *ExprContext) SetId(v antlr.Token) { s.id = v }

func (s *ExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprContext) GetE1() IExprContext { return s.e1 }

func (s *ExprContext) GetE2() IExprContext { return s.e2 }

func (s *ExprContext) SetE1(v IExprContext) { s.e1 = v }

func (s *ExprContext) SetE2(v IExprContext) { s.e2 = v }

func (s *ExprContext) GetDir() string { return s.dir }

func (s *ExprContext) SetDir(v string) { s.dir = v }

func (s *ExprContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(PCompi2ParserENTERO, 0)
}

func (s *ExprContext) ID() antlr.TerminalNode {
	return s.GetToken(PCompi2ParserID, 0)
}

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
	_startState := 4
	p.EnterRecursionRule(localctx, 4, PCompi2ParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PCompi2ParserENTERO:
		{
			p.SetState(21)

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
			p.SetState(23)

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
	p.SetState(34)
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
			localctx = NewExprContext(p, _parentctx, _parentState)
			localctx.(*ExprContext).e1 = _prevctx
			p.PushNewRecursionContext(localctx, _startState, PCompi2ParserRULE_expr)
			p.SetState(27)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				goto errorExit
			}
			{
				p.SetState(28)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*ExprContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == PCompi2ParserT__0 || _la == PCompi2ParserT__1) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*ExprContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(29)

				var _x = p.expr(4)

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

		}
		p.SetState(36)
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
	p.EnterRule(localctx, 6, PCompi2ParserRULE_oprel)
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PCompi2ParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(37)

			var _m = p.Match(PCompi2ParserT__2)

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

	case PCompi2ParserT__3:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(39)

			var _m = p.Match(PCompi2ParserT__3)

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

	case PCompi2ParserT__4:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(41)

			var _m = p.Match(PCompi2ParserT__4)

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

	case PCompi2ParserT__5:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(43)

			var _m = p.Match(PCompi2ParserT__5)

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

	case PCompi2ParserT__6:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(45)

			var _m = p.Match(PCompi2ParserT__6)

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

	case PCompi2ParserT__7:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(47)

			var _m = p.Match(PCompi2ParserT__7)

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
	case 2:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *PCompi2Parser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
