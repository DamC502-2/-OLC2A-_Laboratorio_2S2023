// Code generated from Lab1_G.g4 by ANTLR 4.13.0. DO NOT EDIT.

package Lab1P // Lab1_G
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

type Lab1_GParser struct {
	*antlr.BaseParser
}

var Lab1_GParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func lab1_gParserInit() {
	staticData := &Lab1_GParserStaticData
	staticData.LiteralNames = []string{
		"", "')'", "'('", "", "", "", "", "'return'",
	}
	staticData.SymbolicNames = []string{
		"", "PARDER", "PARIZQ", "ENBLANCO", "ENTERO", "DECIMAL", "ID", "RETURN",
		"UL_C",
	}
	staticData.RuleNames = []string{
		"s", "a",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 8, 17, 2, 0, 7, 0, 2, 1, 7, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 15, 8, 1, 1, 1, 0, 0, 2, 0, 2, 0, 0, 18,
		0, 4, 1, 0, 0, 0, 2, 14, 1, 0, 0, 0, 4, 5, 3, 2, 1, 0, 5, 1, 1, 0, 0, 0,
		6, 7, 5, 2, 0, 0, 7, 8, 3, 2, 1, 0, 8, 9, 5, 1, 0, 0, 9, 15, 1, 0, 0, 0,
		10, 15, 5, 4, 0, 0, 11, 15, 5, 5, 0, 0, 12, 15, 5, 6, 0, 0, 13, 15, 1,
		0, 0, 0, 14, 6, 1, 0, 0, 0, 14, 10, 1, 0, 0, 0, 14, 11, 1, 0, 0, 0, 14,
		12, 1, 0, 0, 0, 14, 13, 1, 0, 0, 0, 15, 3, 1, 0, 0, 0, 1, 14,
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

// Lab1_GParserInit initializes any static state used to implement Lab1_GParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLab1_GParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Lab1_GParserInit() {
	staticData := &Lab1_GParserStaticData
	staticData.once.Do(lab1_gParserInit)
}

// NewLab1_GParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLab1_GParser(input antlr.TokenStream) *Lab1_GParser {
	Lab1_GParserInit()
	this := new(Lab1_GParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &Lab1_GParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Lab1_G.g4"

	return this
}

// Lab1_GParser tokens.
const (
	Lab1_GParserEOF      = antlr.TokenEOF
	Lab1_GParserPARDER   = 1
	Lab1_GParserPARIZQ   = 2
	Lab1_GParserENBLANCO = 3
	Lab1_GParserENTERO   = 4
	Lab1_GParserDECIMAL  = 5
	Lab1_GParserID       = 6
	Lab1_GParserRETURN   = 7
	Lab1_GParserUL_C     = 8
)

// Lab1_GParser rules.
const (
	Lab1_GParserRULE_s = 0
	Lab1_GParserRULE_a = 1
)

// ISContext is an interface to support dynamic dispatch.
type ISContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	A() IAContext

	// IsSContext differentiates from other interfaces.
	IsSContext()
}

type SContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySContext() *SContext {
	var p = new(SContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Lab1_GParserRULE_s
	return p
}

func InitEmptySContext(p *SContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Lab1_GParserRULE_s
}

func (*SContext) IsSContext() {}

func NewSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SContext {
	var p = new(SContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Lab1_GParserRULE_s

	return p
}

func (s *SContext) GetParser() antlr.Parser { return s.parser }

func (s *SContext) A() IAContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAContext)
}

func (s *SContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab1_GListener); ok {
		listenerT.EnterS(s)
	}
}

func (s *SContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab1_GListener); ok {
		listenerT.ExitS(s)
	}
}

func (s *SContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Lab1_GVisitor:
		return t.VisitS(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Lab1_GParser) S() (localctx ISContext) {
	localctx = NewSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Lab1_GParserRULE_s)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(4)
		p.A()
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

// IAContext is an interface to support dynamic dispatch.
type IAContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAContext differentiates from other interfaces.
	IsAContext()
}

type AContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAContext() *AContext {
	var p = new(AContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Lab1_GParserRULE_a
	return p
}

func InitEmptyAContext(p *AContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Lab1_GParserRULE_a
}

func (*AContext) IsAContext() {}

func NewAContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AContext {
	var p = new(AContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Lab1_GParserRULE_a

	return p
}

func (s *AContext) GetParser() antlr.Parser { return s.parser }

func (s *AContext) CopyAll(ctx *AContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AEPSILUMContext struct {
	AContext
}

func NewAEPSILUMContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AEPSILUMContext {
	var p = new(AEPSILUMContext)

	InitEmptyAContext(&p.AContext)
	p.parser = parser
	p.CopyAll(ctx.(*AContext))

	return p
}

func (s *AEPSILUMContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AEPSILUMContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab1_GListener); ok {
		listenerT.EnterAEPSILUM(s)
	}
}

func (s *AEPSILUMContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab1_GListener); ok {
		listenerT.ExitAEPSILUM(s)
	}
}

func (s *AEPSILUMContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Lab1_GVisitor:
		return t.VisitAEPSILUM(s)

	default:
		return t.VisitChildren(s)
	}
}

type ADECIMALContext struct {
	AContext
}

func NewADECIMALContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ADECIMALContext {
	var p = new(ADECIMALContext)

	InitEmptyAContext(&p.AContext)
	p.parser = parser
	p.CopyAll(ctx.(*AContext))

	return p
}

func (s *ADECIMALContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ADECIMALContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(Lab1_GParserDECIMAL, 0)
}

func (s *ADECIMALContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab1_GListener); ok {
		listenerT.EnterADECIMAL(s)
	}
}

func (s *ADECIMALContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab1_GListener); ok {
		listenerT.ExitADECIMAL(s)
	}
}

func (s *ADECIMALContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Lab1_GVisitor:
		return t.VisitADECIMAL(s)

	default:
		return t.VisitChildren(s)
	}
}

type AIDContext struct {
	AContext
}

func NewAIDContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AIDContext {
	var p = new(AIDContext)

	InitEmptyAContext(&p.AContext)
	p.parser = parser
	p.CopyAll(ctx.(*AContext))

	return p
}

func (s *AIDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AIDContext) ID() antlr.TerminalNode {
	return s.GetToken(Lab1_GParserID, 0)
}

func (s *AIDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab1_GListener); ok {
		listenerT.EnterAID(s)
	}
}

func (s *AIDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab1_GListener); ok {
		listenerT.ExitAID(s)
	}
}

func (s *AIDContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Lab1_GVisitor:
		return t.VisitAID(s)

	default:
		return t.VisitChildren(s)
	}
}

type AENTEROContext struct {
	AContext
}

func NewAENTEROContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AENTEROContext {
	var p = new(AENTEROContext)

	InitEmptyAContext(&p.AContext)
	p.parser = parser
	p.CopyAll(ctx.(*AContext))

	return p
}

func (s *AENTEROContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AENTEROContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(Lab1_GParserENTERO, 0)
}

func (s *AENTEROContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab1_GListener); ok {
		listenerT.EnterAENTERO(s)
	}
}

func (s *AENTEROContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab1_GListener); ok {
		listenerT.ExitAENTERO(s)
	}
}

func (s *AENTEROContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Lab1_GVisitor:
		return t.VisitAENTERO(s)

	default:
		return t.VisitChildren(s)
	}
}

type A0Context struct {
	AContext
}

func NewA0Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *A0Context {
	var p = new(A0Context)

	InitEmptyAContext(&p.AContext)
	p.parser = parser
	p.CopyAll(ctx.(*AContext))

	return p
}

func (s *A0Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *A0Context) PARIZQ() antlr.TerminalNode {
	return s.GetToken(Lab1_GParserPARIZQ, 0)
}

func (s *A0Context) A() IAContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAContext)
}

func (s *A0Context) PARDER() antlr.TerminalNode {
	return s.GetToken(Lab1_GParserPARDER, 0)
}

func (s *A0Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab1_GListener); ok {
		listenerT.EnterA0(s)
	}
}

func (s *A0Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab1_GListener); ok {
		listenerT.ExitA0(s)
	}
}

func (s *A0Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Lab1_GVisitor:
		return t.VisitA0(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Lab1_GParser) A() (localctx IAContext) {
	localctx = NewAContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Lab1_GParserRULE_a)
	p.SetState(14)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Lab1_GParserPARIZQ:
		localctx = NewA0Context(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(6)
			p.Match(Lab1_GParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(7)
			p.A()
		}
		{
			p.SetState(8)
			p.Match(Lab1_GParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Lab1_GParserENTERO:
		localctx = NewAENTEROContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(10)
			p.Match(Lab1_GParserENTERO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Lab1_GParserDECIMAL:
		localctx = NewADECIMALContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(11)
			p.Match(Lab1_GParserDECIMAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Lab1_GParserID:
		localctx = NewAIDContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(12)
			p.Match(Lab1_GParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Lab1_GParserEOF, Lab1_GParserPARDER:
		localctx = NewAEPSILUMContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)

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
