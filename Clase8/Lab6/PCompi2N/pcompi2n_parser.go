// Code generated from PCompi2N.g4 by ANTLR 4.13.0. DO NOT EDIT.

package PCompi2N // PCompi2N
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type PCompi2NParser struct {
	*antlr.BaseParser
}

var PCompi2NParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func pcompi2nParserInit() {
	staticData := &PCompi2NParserStaticData
	staticData.LiteralNames = []string{
		"", "'!'", "'true'", "'false'", "'<'", "'>'", "'=='", "'!='", "'AND'",
		"'OR'", "'-'", "'*'", "'/'", "'+'", "')'", "'('",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "PARDER", "PARIZQ",
		"ENTERO", "DECIMAL", "ENBLANCO", "ID",
	}
	staticData.RuleNames = []string{
		"s", "cond", "expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 19, 57, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 1, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3,
		1, 22, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 30, 8, 1, 10, 1,
		12, 1, 33, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		3, 2, 44, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 52, 8, 2, 10,
		2, 12, 2, 55, 9, 2, 1, 2, 0, 2, 2, 4, 3, 0, 2, 4, 0, 4, 1, 0, 2, 3, 1,
		0, 4, 7, 1, 0, 11, 12, 2, 0, 10, 10, 13, 13, 63, 0, 6, 1, 0, 0, 0, 2, 21,
		1, 0, 0, 0, 4, 43, 1, 0, 0, 0, 6, 7, 3, 2, 1, 0, 7, 8, 5, 0, 0, 1, 8, 1,
		1, 0, 0, 0, 9, 10, 6, 1, -1, 0, 10, 11, 5, 1, 0, 0, 11, 22, 3, 2, 1, 6,
		12, 22, 7, 0, 0, 0, 13, 14, 3, 4, 2, 0, 14, 15, 7, 1, 0, 0, 15, 16, 3,
		4, 2, 0, 16, 22, 1, 0, 0, 0, 17, 18, 5, 15, 0, 0, 18, 19, 3, 2, 1, 0, 19,
		20, 5, 14, 0, 0, 20, 22, 1, 0, 0, 0, 21, 9, 1, 0, 0, 0, 21, 12, 1, 0, 0,
		0, 21, 13, 1, 0, 0, 0, 21, 17, 1, 0, 0, 0, 22, 31, 1, 0, 0, 0, 23, 24,
		10, 3, 0, 0, 24, 25, 5, 8, 0, 0, 25, 30, 3, 2, 1, 4, 26, 27, 10, 2, 0,
		0, 27, 28, 5, 9, 0, 0, 28, 30, 3, 2, 1, 3, 29, 23, 1, 0, 0, 0, 29, 26,
		1, 0, 0, 0, 30, 33, 1, 0, 0, 0, 31, 29, 1, 0, 0, 0, 31, 32, 1, 0, 0, 0,
		32, 3, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 34, 35, 6, 2, -1, 0, 35, 36, 5,
		10, 0, 0, 36, 44, 3, 4, 2, 6, 37, 38, 5, 15, 0, 0, 38, 39, 3, 4, 2, 0,
		39, 40, 5, 14, 0, 0, 40, 44, 1, 0, 0, 0, 41, 44, 5, 16, 0, 0, 42, 44, 5,
		19, 0, 0, 43, 34, 1, 0, 0, 0, 43, 37, 1, 0, 0, 0, 43, 41, 1, 0, 0, 0, 43,
		42, 1, 0, 0, 0, 44, 53, 1, 0, 0, 0, 45, 46, 10, 5, 0, 0, 46, 47, 7, 2,
		0, 0, 47, 52, 3, 4, 2, 6, 48, 49, 10, 4, 0, 0, 49, 50, 7, 3, 0, 0, 50,
		52, 3, 4, 2, 5, 51, 45, 1, 0, 0, 0, 51, 48, 1, 0, 0, 0, 52, 55, 1, 0, 0,
		0, 53, 51, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 5, 1, 0, 0, 0, 55, 53, 1,
		0, 0, 0, 6, 21, 29, 31, 43, 51, 53,
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

// PCompi2NParserInit initializes any static state used to implement PCompi2NParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPCompi2NParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PCompi2NParserInit() {
	staticData := &PCompi2NParserStaticData
	staticData.once.Do(pcompi2nParserInit)
}

// NewPCompi2NParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPCompi2NParser(input antlr.TokenStream) *PCompi2NParser {
	PCompi2NParserInit()
	this := new(PCompi2NParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &PCompi2NParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "PCompi2N.g4"

	return this
}

// PCompi2NParser tokens.
const (
	PCompi2NParserEOF      = antlr.TokenEOF
	PCompi2NParserT__0     = 1
	PCompi2NParserT__1     = 2
	PCompi2NParserT__2     = 3
	PCompi2NParserT__3     = 4
	PCompi2NParserT__4     = 5
	PCompi2NParserT__5     = 6
	PCompi2NParserT__6     = 7
	PCompi2NParserT__7     = 8
	PCompi2NParserT__8     = 9
	PCompi2NParserT__9     = 10
	PCompi2NParserT__10    = 11
	PCompi2NParserT__11    = 12
	PCompi2NParserT__12    = 13
	PCompi2NParserPARDER   = 14
	PCompi2NParserPARIZQ   = 15
	PCompi2NParserENTERO   = 16
	PCompi2NParserDECIMAL  = 17
	PCompi2NParserENBLANCO = 18
	PCompi2NParserID       = 19
)

// PCompi2NParser rules.
const (
	PCompi2NParserRULE_s    = 0
	PCompi2NParserRULE_cond = 1
	PCompi2NParserRULE_expr = 2
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
	p.RuleIndex = PCompi2NParserRULE_s
	return p
}

func InitEmptySContext(p *SContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PCompi2NParserRULE_s
}

func (*SContext) IsSContext() {}

func NewSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SContext {
	var p = new(SContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PCompi2NParserRULE_s

	return p
}

func (s *SContext) GetParser() antlr.Parser { return s.parser }

func (s *SContext) GetC1() ICondContext { return s.c1 }

func (s *SContext) SetC1(v ICondContext) { s.c1 = v }

func (s *SContext) EOF() antlr.TerminalNode {
	return s.GetToken(PCompi2NParserEOF, 0)
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
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.EnterS(s)
	}
}

func (s *SContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.ExitS(s)
	}
}

func (s *SContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PCompi2NVisitor:
		return t.VisitS(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PCompi2NParser) S() (localctx ISContext) {
	localctx = NewSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PCompi2NParserRULE_s)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(6)

		var _x = p.cond(0)

		localctx.(*SContext).c1 = _x
	}
	{
		p.SetState(7)
		p.Match(PCompi2NParserEOF)
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
	// IsCondContext differentiates from other interfaces.
	IsCondContext()
}

type CondContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCondContext() *CondContext {
	var p = new(CondContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PCompi2NParserRULE_cond
	return p
}

func InitEmptyCondContext(p *CondContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PCompi2NParserRULE_cond
}

func (*CondContext) IsCondContext() {}

func NewCondContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CondContext {
	var p = new(CondContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PCompi2NParserRULE_cond

	return p
}

func (s *CondContext) GetParser() antlr.Parser { return s.parser }

func (s *CondContext) CopyAll(ctx *CondContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *CondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CondOrContext struct {
	CondContext
	c1 ICondContext
	op antlr.Token
	c2 ICondContext
}

func NewCondOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CondOrContext {
	var p = new(CondOrContext)

	InitEmptyCondContext(&p.CondContext)
	p.parser = parser
	p.CopyAll(ctx.(*CondContext))

	return p
}

func (s *CondOrContext) GetOp() antlr.Token { return s.op }

func (s *CondOrContext) SetOp(v antlr.Token) { s.op = v }

func (s *CondOrContext) GetC1() ICondContext { return s.c1 }

func (s *CondOrContext) GetC2() ICondContext { return s.c2 }

func (s *CondOrContext) SetC1(v ICondContext) { s.c1 = v }

func (s *CondOrContext) SetC2(v ICondContext) { s.c2 = v }

func (s *CondOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondOrContext) AllCond() []ICondContext {
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

func (s *CondOrContext) Cond(i int) ICondContext {
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

func (s *CondOrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.EnterCondOr(s)
	}
}

func (s *CondOrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.ExitCondOr(s)
	}
}

func (s *CondOrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PCompi2NVisitor:
		return t.VisitCondOr(s)

	default:
		return t.VisitChildren(s)
	}
}

type CondAndContext struct {
	CondContext
	c1 ICondContext
	op antlr.Token
	c2 ICondContext
}

func NewCondAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CondAndContext {
	var p = new(CondAndContext)

	InitEmptyCondContext(&p.CondContext)
	p.parser = parser
	p.CopyAll(ctx.(*CondContext))

	return p
}

func (s *CondAndContext) GetOp() antlr.Token { return s.op }

func (s *CondAndContext) SetOp(v antlr.Token) { s.op = v }

func (s *CondAndContext) GetC1() ICondContext { return s.c1 }

func (s *CondAndContext) GetC2() ICondContext { return s.c2 }

func (s *CondAndContext) SetC1(v ICondContext) { s.c1 = v }

func (s *CondAndContext) SetC2(v ICondContext) { s.c2 = v }

func (s *CondAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondAndContext) AllCond() []ICondContext {
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

func (s *CondAndContext) Cond(i int) ICondContext {
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

func (s *CondAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.EnterCondAnd(s)
	}
}

func (s *CondAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.ExitCondAnd(s)
	}
}

func (s *CondAndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PCompi2NVisitor:
		return t.VisitCondAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

type CondIContext struct {
	CondContext
	op antlr.Token
	c  ICondContext
}

func NewCondIContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CondIContext {
	var p = new(CondIContext)

	InitEmptyCondContext(&p.CondContext)
	p.parser = parser
	p.CopyAll(ctx.(*CondContext))

	return p
}

func (s *CondIContext) GetOp() antlr.Token { return s.op }

func (s *CondIContext) SetOp(v antlr.Token) { s.op = v }

func (s *CondIContext) GetC() ICondContext { return s.c }

func (s *CondIContext) SetC(v ICondContext) { s.c = v }

func (s *CondIContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondIContext) Cond() ICondContext {
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

func (s *CondIContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.EnterCondI(s)
	}
}

func (s *CondIContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.ExitCondI(s)
	}
}

func (s *CondIContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PCompi2NVisitor:
		return t.VisitCondI(s)

	default:
		return t.VisitChildren(s)
	}
}

type CondBoolContext struct {
	CondContext
	bool_ antlr.Token
}

func NewCondBoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CondBoolContext {
	var p = new(CondBoolContext)

	InitEmptyCondContext(&p.CondContext)
	p.parser = parser
	p.CopyAll(ctx.(*CondContext))

	return p
}

func (s *CondBoolContext) GetBool_() antlr.Token { return s.bool_ }

func (s *CondBoolContext) SetBool_(v antlr.Token) { s.bool_ = v }

func (s *CondBoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondBoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.EnterCondBool(s)
	}
}

func (s *CondBoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.ExitCondBool(s)
	}
}

func (s *CondBoolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PCompi2NVisitor:
		return t.VisitCondBool(s)

	default:
		return t.VisitChildren(s)
	}
}

type CondParContext struct {
	CondContext
	c ICondContext
}

func NewCondParContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CondParContext {
	var p = new(CondParContext)

	InitEmptyCondContext(&p.CondContext)
	p.parser = parser
	p.CopyAll(ctx.(*CondContext))

	return p
}

func (s *CondParContext) GetC() ICondContext { return s.c }

func (s *CondParContext) SetC(v ICondContext) { s.c = v }

func (s *CondParContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondParContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(PCompi2NParserPARIZQ, 0)
}

func (s *CondParContext) PARDER() antlr.TerminalNode {
	return s.GetToken(PCompi2NParserPARDER, 0)
}

func (s *CondParContext) Cond() ICondContext {
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

func (s *CondParContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.EnterCondPar(s)
	}
}

func (s *CondParContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.ExitCondPar(s)
	}
}

func (s *CondParContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PCompi2NVisitor:
		return t.VisitCondPar(s)

	default:
		return t.VisitChildren(s)
	}
}

type CondOperContext struct {
	CondContext
	e1  IExprContext
	opr antlr.Token
	e2  IExprContext
}

func NewCondOperContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CondOperContext {
	var p = new(CondOperContext)

	InitEmptyCondContext(&p.CondContext)
	p.parser = parser
	p.CopyAll(ctx.(*CondContext))

	return p
}

func (s *CondOperContext) GetOpr() antlr.Token { return s.opr }

func (s *CondOperContext) SetOpr(v antlr.Token) { s.opr = v }

func (s *CondOperContext) GetE1() IExprContext { return s.e1 }

func (s *CondOperContext) GetE2() IExprContext { return s.e2 }

func (s *CondOperContext) SetE1(v IExprContext) { s.e1 = v }

func (s *CondOperContext) SetE2(v IExprContext) { s.e2 = v }

func (s *CondOperContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondOperContext) AllExpr() []IExprContext {
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

func (s *CondOperContext) Expr(i int) IExprContext {
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

func (s *CondOperContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.EnterCondOper(s)
	}
}

func (s *CondOperContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.ExitCondOper(s)
	}
}

func (s *CondOperContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PCompi2NVisitor:
		return t.VisitCondOper(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PCompi2NParser) Cond() (localctx ICondContext) {
	return p.cond(0)
}

func (p *PCompi2NParser) cond(_p int) (localctx ICondContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewCondContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ICondContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, PCompi2NParserRULE_cond, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewCondIContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(10)

			var _m = p.Match(PCompi2NParserT__0)

			localctx.(*CondIContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(11)

			var _x = p.cond(6)

			localctx.(*CondIContext).c = _x
		}

	case 2:
		localctx = NewCondBoolContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(12)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*CondBoolContext).bool_ = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == PCompi2NParserT__1 || _la == PCompi2NParserT__2) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*CondBoolContext).bool_ = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 3:
		localctx = NewCondOperContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(13)

			var _x = p.expr(0)

			localctx.(*CondOperContext).e1 = _x
		}
		{
			p.SetState(14)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*CondOperContext).opr = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&240) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*CondOperContext).opr = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(15)

			var _x = p.expr(0)

			localctx.(*CondOperContext).e2 = _x
		}

	case 4:
		localctx = NewCondParContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(17)
			p.Match(PCompi2NParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(18)

			var _x = p.cond(0)

			localctx.(*CondParContext).c = _x
		}
		{
			p.SetState(19)
			p.Match(PCompi2NParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(31)
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
			p.SetState(29)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewCondAndContext(p, NewCondContext(p, _parentctx, _parentState))
				localctx.(*CondAndContext).c1 = _prevctx

				p.PushNewRecursionContext(localctx, _startState, PCompi2NParserRULE_cond)
				p.SetState(23)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(24)

					var _m = p.Match(PCompi2NParserT__7)

					localctx.(*CondAndContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(25)

					var _x = p.cond(4)

					localctx.(*CondAndContext).c2 = _x
				}

			case 2:
				localctx = NewCondOrContext(p, NewCondContext(p, _parentctx, _parentState))
				localctx.(*CondOrContext).c1 = _prevctx

				p.PushNewRecursionContext(localctx, _startState, PCompi2NParserRULE_cond)
				p.SetState(26)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(27)

					var _m = p.Match(PCompi2NParserT__8)

					localctx.(*CondOrContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(28)

					var _x = p.cond(3)

					localctx.(*CondOrContext).c2 = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(33)
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

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PCompi2NParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PCompi2NParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PCompi2NParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type EnumContext struct {
	ExprContext
	n antlr.Token
}

func NewEnumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EnumContext {
	var p = new(EnumContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EnumContext) GetN() antlr.Token { return s.n }

func (s *EnumContext) SetN(v antlr.Token) { s.n = v }

func (s *EnumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(PCompi2NParserENTERO, 0)
}

func (s *EnumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.EnterEnum(s)
	}
}

func (s *EnumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.ExitEnum(s)
	}
}

func (s *EnumContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PCompi2NVisitor:
		return t.VisitEnum(s)

	default:
		return t.VisitChildren(s)
	}
}

type EidContext struct {
	ExprContext
	id antlr.Token
}

func NewEidContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EidContext {
	var p = new(EidContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EidContext) GetId() antlr.Token { return s.id }

func (s *EidContext) SetId(v antlr.Token) { s.id = v }

func (s *EidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EidContext) ID() antlr.TerminalNode {
	return s.GetToken(PCompi2NParserID, 0)
}

func (s *EidContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.EnterEid(s)
	}
}

func (s *EidContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.ExitEid(s)
	}
}

func (s *EidContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PCompi2NVisitor:
		return t.VisitEid(s)

	default:
		return t.VisitChildren(s)
	}
}

type EsrContext struct {
	ExprContext
	e1 IExprContext
	op antlr.Token
	e2 IExprContext
}

func NewEsrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EsrContext {
	var p = new(EsrContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EsrContext) GetOp() antlr.Token { return s.op }

func (s *EsrContext) SetOp(v antlr.Token) { s.op = v }

func (s *EsrContext) GetE1() IExprContext { return s.e1 }

func (s *EsrContext) GetE2() IExprContext { return s.e2 }

func (s *EsrContext) SetE1(v IExprContext) { s.e1 = v }

func (s *EsrContext) SetE2(v IExprContext) { s.e2 = v }

func (s *EsrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EsrContext) AllExpr() []IExprContext {
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

func (s *EsrContext) Expr(i int) IExprContext {
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

func (s *EsrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.EnterEsr(s)
	}
}

func (s *EsrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.ExitEsr(s)
	}
}

func (s *EsrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PCompi2NVisitor:
		return t.VisitEsr(s)

	default:
		return t.VisitChildren(s)
	}
}

type EpaContext struct {
	ExprContext
	e1 IExprContext
}

func NewEpaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EpaContext {
	var p = new(EpaContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EpaContext) GetE1() IExprContext { return s.e1 }

func (s *EpaContext) SetE1(v IExprContext) { s.e1 = v }

func (s *EpaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EpaContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(PCompi2NParserPARIZQ, 0)
}

func (s *EpaContext) PARDER() antlr.TerminalNode {
	return s.GetToken(PCompi2NParserPARDER, 0)
}

func (s *EpaContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EpaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.EnterEpa(s)
	}
}

func (s *EpaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.ExitEpa(s)
	}
}

func (s *EpaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PCompi2NVisitor:
		return t.VisitEpa(s)

	default:
		return t.VisitChildren(s)
	}
}

type ENegContext struct {
	ExprContext
	op antlr.Token
	e1 IExprContext
}

func NewENegContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ENegContext {
	var p = new(ENegContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ENegContext) GetOp() antlr.Token { return s.op }

func (s *ENegContext) SetOp(v antlr.Token) { s.op = v }

func (s *ENegContext) GetE1() IExprContext { return s.e1 }

func (s *ENegContext) SetE1(v IExprContext) { s.e1 = v }

func (s *ENegContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ENegContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ENegContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.EnterENeg(s)
	}
}

func (s *ENegContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.ExitENeg(s)
	}
}

func (s *ENegContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PCompi2NVisitor:
		return t.VisitENeg(s)

	default:
		return t.VisitChildren(s)
	}
}

type EmdContext struct {
	ExprContext
	e1 IExprContext
	op antlr.Token
	e2 IExprContext
}

func NewEmdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EmdContext {
	var p = new(EmdContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EmdContext) GetOp() antlr.Token { return s.op }

func (s *EmdContext) SetOp(v antlr.Token) { s.op = v }

func (s *EmdContext) GetE1() IExprContext { return s.e1 }

func (s *EmdContext) GetE2() IExprContext { return s.e2 }

func (s *EmdContext) SetE1(v IExprContext) { s.e1 = v }

func (s *EmdContext) SetE2(v IExprContext) { s.e2 = v }

func (s *EmdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmdContext) AllExpr() []IExprContext {
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

func (s *EmdContext) Expr(i int) IExprContext {
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

func (s *EmdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.EnterEmd(s)
	}
}

func (s *EmdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.ExitEmd(s)
	}
}

func (s *EmdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PCompi2NVisitor:
		return t.VisitEmd(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PCompi2NParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *PCompi2NParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, PCompi2NParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PCompi2NParserT__9:
		localctx = NewENegContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(35)

			var _m = p.Match(PCompi2NParserT__9)

			localctx.(*ENegContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(36)

			var _x = p.expr(6)

			localctx.(*ENegContext).e1 = _x
		}

	case PCompi2NParserPARIZQ:
		localctx = NewEpaContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(37)
			p.Match(PCompi2NParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(38)

			var _x = p.expr(0)

			localctx.(*EpaContext).e1 = _x
		}
		{
			p.SetState(39)
			p.Match(PCompi2NParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PCompi2NParserENTERO:
		localctx = NewEnumContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(41)

			var _m = p.Match(PCompi2NParserENTERO)

			localctx.(*EnumContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PCompi2NParserID:
		localctx = NewEidContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(42)

			var _m = p.Match(PCompi2NParserID)

			localctx.(*EidContext).id = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(53)
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
			p.SetState(51)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewEmdContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*EmdContext).e1 = _prevctx

				p.PushNewRecursionContext(localctx, _startState, PCompi2NParserRULE_expr)
				p.SetState(45)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(46)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EmdContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == PCompi2NParserT__10 || _la == PCompi2NParserT__11) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EmdContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(47)

					var _x = p.expr(6)

					localctx.(*EmdContext).e2 = _x
				}

			case 2:
				localctx = NewEsrContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*EsrContext).e1 = _prevctx

				p.PushNewRecursionContext(localctx, _startState, PCompi2NParserRULE_expr)
				p.SetState(48)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(49)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EsrContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == PCompi2NParserT__9 || _la == PCompi2NParserT__12) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EsrContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(50)

					var _x = p.expr(5)

					localctx.(*EsrContext).e2 = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(55)
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

func (p *PCompi2NParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *CondContext = nil
		if localctx != nil {
			t = localctx.(*CondContext)
		}
		return p.Cond_Sempred(t, predIndex)

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

func (p *PCompi2NParser) Cond_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PCompi2NParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
