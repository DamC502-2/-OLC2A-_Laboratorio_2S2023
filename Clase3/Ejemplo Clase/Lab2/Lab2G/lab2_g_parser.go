// Code generated from Lab2_G.g4 by ANTLR 4.13.0. DO NOT EDIT.

package Lab2G // Lab2_G
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

type Lab2_GParser struct {
	*antlr.BaseParser
}

var Lab2_GParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func lab2_gParserInit() {
	staticData := &Lab2_GParserStaticData
	staticData.LiteralNames = []string{
		"", "','", "'while'", "'{'", "'}'", "'+='", "'-'", "'!'", "'*'", "'+'",
		"')'", "'('", "'='", "'/'", "'int'", "'Console.out'", "", "", "", "",
		"'return'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "PARDER", "PARIZQ", "ASIG",
		"DIV", "INT", "CONSOLE", "ENBLANCO", "ENTERO", "DECIMAL", "ID", "RETURN",
		"UL_C",
	}
	staticData.RuleNames = []string{
		"s", "l_sentencias", "sentencia", "asignacion", "e",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 21, 79, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 1, 0, 1, 0, 1, 1, 5, 1, 14, 8, 1, 10, 1, 12, 1, 17, 9, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 5, 2, 24, 8, 2, 10, 2, 12, 2, 27, 9, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 3, 2, 35, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 3, 2, 46, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3,
		3, 54, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		3, 4, 66, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 74, 8, 4, 10,
		4, 12, 4, 77, 9, 4, 1, 4, 0, 1, 8, 5, 0, 2, 4, 6, 8, 0, 3, 1, 0, 6, 7,
		2, 0, 8, 8, 13, 13, 2, 0, 6, 6, 9, 9, 86, 0, 10, 1, 0, 0, 0, 2, 15, 1,
		0, 0, 0, 4, 45, 1, 0, 0, 0, 6, 53, 1, 0, 0, 0, 8, 65, 1, 0, 0, 0, 10, 11,
		3, 2, 1, 0, 11, 1, 1, 0, 0, 0, 12, 14, 3, 4, 2, 0, 13, 12, 1, 0, 0, 0,
		14, 17, 1, 0, 0, 0, 15, 13, 1, 0, 0, 0, 15, 16, 1, 0, 0, 0, 16, 3, 1, 0,
		0, 0, 17, 15, 1, 0, 0, 0, 18, 19, 5, 15, 0, 0, 19, 20, 5, 11, 0, 0, 20,
		25, 3, 8, 4, 0, 21, 22, 5, 1, 0, 0, 22, 24, 3, 8, 4, 0, 23, 21, 1, 0, 0,
		0, 24, 27, 1, 0, 0, 0, 25, 23, 1, 0, 0, 0, 25, 26, 1, 0, 0, 0, 26, 28,
		1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 28, 29, 5, 10, 0, 0, 29, 46, 1, 0, 0, 0,
		30, 31, 5, 14, 0, 0, 31, 34, 5, 19, 0, 0, 32, 33, 5, 12, 0, 0, 33, 35,
		3, 8, 4, 0, 34, 32, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 46, 1, 0, 0, 0,
		36, 46, 3, 6, 3, 0, 37, 38, 5, 2, 0, 0, 38, 39, 5, 10, 0, 0, 39, 40, 3,
		8, 4, 0, 40, 41, 5, 11, 0, 0, 41, 42, 5, 3, 0, 0, 42, 43, 3, 2, 1, 0, 43,
		44, 5, 4, 0, 0, 44, 46, 1, 0, 0, 0, 45, 18, 1, 0, 0, 0, 45, 30, 1, 0, 0,
		0, 45, 36, 1, 0, 0, 0, 45, 37, 1, 0, 0, 0, 46, 5, 1, 0, 0, 0, 47, 48, 5,
		19, 0, 0, 48, 49, 5, 5, 0, 0, 49, 54, 3, 8, 4, 0, 50, 51, 5, 19, 0, 0,
		51, 52, 5, 12, 0, 0, 52, 54, 3, 8, 4, 0, 53, 47, 1, 0, 0, 0, 53, 50, 1,
		0, 0, 0, 54, 7, 1, 0, 0, 0, 55, 56, 6, 4, -1, 0, 56, 66, 5, 19, 0, 0, 57,
		66, 5, 18, 0, 0, 58, 66, 5, 17, 0, 0, 59, 60, 7, 0, 0, 0, 60, 66, 3, 8,
		4, 4, 61, 62, 5, 11, 0, 0, 62, 63, 3, 8, 4, 0, 63, 64, 5, 10, 0, 0, 64,
		66, 1, 0, 0, 0, 65, 55, 1, 0, 0, 0, 65, 57, 1, 0, 0, 0, 65, 58, 1, 0, 0,
		0, 65, 59, 1, 0, 0, 0, 65, 61, 1, 0, 0, 0, 66, 75, 1, 0, 0, 0, 67, 68,
		10, 3, 0, 0, 68, 69, 7, 1, 0, 0, 69, 74, 3, 8, 4, 4, 70, 71, 10, 2, 0,
		0, 71, 72, 7, 2, 0, 0, 72, 74, 3, 8, 4, 3, 73, 67, 1, 0, 0, 0, 73, 70,
		1, 0, 0, 0, 74, 77, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0,
		76, 9, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 8, 15, 25, 34, 45, 53, 65, 73, 75,
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

// Lab2_GParserInit initializes any static state used to implement Lab2_GParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLab2_GParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Lab2_GParserInit() {
	staticData := &Lab2_GParserStaticData
	staticData.once.Do(lab2_gParserInit)
}

// NewLab2_GParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLab2_GParser(input antlr.TokenStream) *Lab2_GParser {
	Lab2_GParserInit()
	this := new(Lab2_GParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &Lab2_GParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Lab2_G.g4"

	return this
}

// Lab2_GParser tokens.
const (
	Lab2_GParserEOF      = antlr.TokenEOF
	Lab2_GParserT__0     = 1
	Lab2_GParserT__1     = 2
	Lab2_GParserT__2     = 3
	Lab2_GParserT__3     = 4
	Lab2_GParserT__4     = 5
	Lab2_GParserT__5     = 6
	Lab2_GParserT__6     = 7
	Lab2_GParserT__7     = 8
	Lab2_GParserT__8     = 9
	Lab2_GParserPARDER   = 10
	Lab2_GParserPARIZQ   = 11
	Lab2_GParserASIG     = 12
	Lab2_GParserDIV      = 13
	Lab2_GParserINT      = 14
	Lab2_GParserCONSOLE  = 15
	Lab2_GParserENBLANCO = 16
	Lab2_GParserENTERO   = 17
	Lab2_GParserDECIMAL  = 18
	Lab2_GParserID       = 19
	Lab2_GParserRETURN   = 20
	Lab2_GParserUL_C     = 21
)

// Lab2_GParser rules.
const (
	Lab2_GParserRULE_s            = 0
	Lab2_GParserRULE_l_sentencias = 1
	Lab2_GParserRULE_sentencia    = 2
	Lab2_GParserRULE_asignacion   = 3
	Lab2_GParserRULE_e            = 4
)

// ISContext is an interface to support dynamic dispatch.
type ISContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
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
	p.RuleIndex = Lab2_GParserRULE_s
	return p
}

func InitEmptySContext(p *SContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Lab2_GParserRULE_s
}

func (*SContext) IsSContext() {}

func NewSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SContext {
	var p = new(SContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Lab2_GParserRULE_s

	return p
}

func (s *SContext) GetParser() antlr.Parser { return s.parser }

func (s *SContext) CopyAll(ctx *SContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SLSentenciasContext struct {
	SContext
}

func NewSLSentenciasContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SLSentenciasContext {
	var p = new(SLSentenciasContext)

	InitEmptySContext(&p.SContext)
	p.parser = parser
	p.CopyAll(ctx.(*SContext))

	return p
}

func (s *SLSentenciasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SLSentenciasContext) L_sentencias() IL_sentenciasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IL_sentenciasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IL_sentenciasContext)
}

func (s *SLSentenciasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.EnterSLSentencias(s)
	}
}

func (s *SLSentenciasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.ExitSLSentencias(s)
	}
}

func (s *SLSentenciasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Lab2_GVisitor:
		return t.VisitSLSentencias(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Lab2_GParser) S() (localctx ISContext) {
	localctx = NewSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Lab2_GParserRULE_s)
	localctx = NewSLSentenciasContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(10)
		p.L_sentencias()
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

// IL_sentenciasContext is an interface to support dynamic dispatch.
type IL_sentenciasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsL_sentenciasContext differentiates from other interfaces.
	IsL_sentenciasContext()
}

type L_sentenciasContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyL_sentenciasContext() *L_sentenciasContext {
	var p = new(L_sentenciasContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Lab2_GParserRULE_l_sentencias
	return p
}

func InitEmptyL_sentenciasContext(p *L_sentenciasContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Lab2_GParserRULE_l_sentencias
}

func (*L_sentenciasContext) IsL_sentenciasContext() {}

func NewL_sentenciasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *L_sentenciasContext {
	var p = new(L_sentenciasContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Lab2_GParserRULE_l_sentencias

	return p
}

func (s *L_sentenciasContext) GetParser() antlr.Parser { return s.parser }

func (s *L_sentenciasContext) CopyAll(ctx *L_sentenciasContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *L_sentenciasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *L_sentenciasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type L_SentenciaContext struct {
	L_sentenciasContext
}

func NewL_SentenciaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *L_SentenciaContext {
	var p = new(L_SentenciaContext)

	InitEmptyL_sentenciasContext(&p.L_sentenciasContext)
	p.parser = parser
	p.CopyAll(ctx.(*L_sentenciasContext))

	return p
}

func (s *L_SentenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *L_SentenciaContext) AllSentencia() []ISentenciaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISentenciaContext); ok {
			len++
		}
	}

	tst := make([]ISentenciaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISentenciaContext); ok {
			tst[i] = t.(ISentenciaContext)
			i++
		}
	}

	return tst
}

func (s *L_SentenciaContext) Sentencia(i int) ISentenciaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentenciaContext); ok {
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

	return t.(ISentenciaContext)
}

func (s *L_SentenciaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.EnterL_Sentencia(s)
	}
}

func (s *L_SentenciaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.ExitL_Sentencia(s)
	}
}

func (s *L_SentenciaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Lab2_GVisitor:
		return t.VisitL_Sentencia(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Lab2_GParser) L_sentencias() (localctx IL_sentenciasContext) {
	localctx = NewL_sentenciasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Lab2_GParserRULE_l_sentencias)
	var _la int

	localctx = NewL_SentenciaContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&573444) != 0 {
		{
			p.SetState(12)
			p.Sentencia()
		}

		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// ISentenciaContext is an interface to support dynamic dispatch.
type ISentenciaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSentenciaContext differentiates from other interfaces.
	IsSentenciaContext()
}

type SentenciaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySentenciaContext() *SentenciaContext {
	var p = new(SentenciaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Lab2_GParserRULE_sentencia
	return p
}

func InitEmptySentenciaContext(p *SentenciaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Lab2_GParserRULE_sentencia
}

func (*SentenciaContext) IsSentenciaContext() {}

func NewSentenciaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SentenciaContext {
	var p = new(SentenciaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Lab2_GParserRULE_sentencia

	return p
}

func (s *SentenciaContext) GetParser() antlr.Parser { return s.parser }

func (s *SentenciaContext) CopyAll(ctx *SentenciaContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SentenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentenciaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type S_AsigContext struct {
	SentenciaContext
}

func NewS_AsigContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_AsigContext {
	var p = new(S_AsigContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_AsigContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_AsigContext) Asignacion() IAsignacionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignacionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignacionContext)
}

func (s *S_AsigContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.EnterS_Asig(s)
	}
}

func (s *S_AsigContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.ExitS_Asig(s)
	}
}

func (s *S_AsigContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Lab2_GVisitor:
		return t.VisitS_Asig(s)

	default:
		return t.VisitChildren(s)
	}
}

type SConsolaContext struct {
	SentenciaContext
}

func NewSConsolaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SConsolaContext {
	var p = new(SConsolaContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *SConsolaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SConsolaContext) CONSOLE() antlr.TerminalNode {
	return s.GetToken(Lab2_GParserCONSOLE, 0)
}

func (s *SConsolaContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(Lab2_GParserPARIZQ, 0)
}

func (s *SConsolaContext) AllE() []IEContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEContext); ok {
			len++
		}
	}

	tst := make([]IEContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEContext); ok {
			tst[i] = t.(IEContext)
			i++
		}
	}

	return tst
}

func (s *SConsolaContext) E(i int) IEContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
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

	return t.(IEContext)
}

func (s *SConsolaContext) PARDER() antlr.TerminalNode {
	return s.GetToken(Lab2_GParserPARDER, 0)
}

func (s *SConsolaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.EnterSConsola(s)
	}
}

func (s *SConsolaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.ExitSConsola(s)
	}
}

func (s *SConsolaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Lab2_GVisitor:
		return t.VisitSConsola(s)

	default:
		return t.VisitChildren(s)
	}
}

type SWhileContext struct {
	SentenciaContext
}

func NewSWhileContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SWhileContext {
	var p = new(SWhileContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *SWhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SWhileContext) PARDER() antlr.TerminalNode {
	return s.GetToken(Lab2_GParserPARDER, 0)
}

func (s *SWhileContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *SWhileContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(Lab2_GParserPARIZQ, 0)
}

func (s *SWhileContext) L_sentencias() IL_sentenciasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IL_sentenciasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IL_sentenciasContext)
}

func (s *SWhileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.EnterSWhile(s)
	}
}

func (s *SWhileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.ExitSWhile(s)
	}
}

func (s *SWhileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Lab2_GVisitor:
		return t.VisitSWhile(s)

	default:
		return t.VisitChildren(s)
	}
}

type SDeclaracionContext struct {
	SentenciaContext
}

func NewSDeclaracionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SDeclaracionContext {
	var p = new(SDeclaracionContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *SDeclaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SDeclaracionContext) INT() antlr.TerminalNode {
	return s.GetToken(Lab2_GParserINT, 0)
}

func (s *SDeclaracionContext) ID() antlr.TerminalNode {
	return s.GetToken(Lab2_GParserID, 0)
}

func (s *SDeclaracionContext) ASIG() antlr.TerminalNode {
	return s.GetToken(Lab2_GParserASIG, 0)
}

func (s *SDeclaracionContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *SDeclaracionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.EnterSDeclaracion(s)
	}
}

func (s *SDeclaracionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.ExitSDeclaracion(s)
	}
}

func (s *SDeclaracionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Lab2_GVisitor:
		return t.VisitSDeclaracion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Lab2_GParser) Sentencia() (localctx ISentenciaContext) {
	localctx = NewSentenciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, Lab2_GParserRULE_sentencia)
	var _la int

	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Lab2_GParserCONSOLE:
		localctx = NewSConsolaContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(18)
			p.Match(Lab2_GParserCONSOLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(19)
			p.Match(Lab2_GParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(20)
			p.e(0)
		}
		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == Lab2_GParserT__0 {
			{
				p.SetState(21)
				p.Match(Lab2_GParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(22)
				p.e(0)
			}

			p.SetState(27)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(28)
			p.Match(Lab2_GParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Lab2_GParserINT:
		localctx = NewSDeclaracionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(30)
			p.Match(Lab2_GParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(31)
			p.Match(Lab2_GParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Lab2_GParserASIG {
			{
				p.SetState(32)
				p.Match(Lab2_GParserASIG)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(33)
				p.e(0)
			}

		}

	case Lab2_GParserID:
		localctx = NewS_AsigContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(36)
			p.Asignacion()
		}

	case Lab2_GParserT__1:
		localctx = NewSWhileContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(37)
			p.Match(Lab2_GParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(38)
			p.Match(Lab2_GParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(39)
			p.e(0)
		}
		{
			p.SetState(40)
			p.Match(Lab2_GParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(41)
			p.Match(Lab2_GParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(42)
			p.L_sentencias()
		}
		{
			p.SetState(43)
			p.Match(Lab2_GParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

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

// IAsignacionContext is an interface to support dynamic dispatch.
type IAsignacionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAsignacionContext differentiates from other interfaces.
	IsAsignacionContext()
}

type AsignacionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsignacionContext() *AsignacionContext {
	var p = new(AsignacionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Lab2_GParserRULE_asignacion
	return p
}

func InitEmptyAsignacionContext(p *AsignacionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Lab2_GParserRULE_asignacion
}

func (*AsignacionContext) IsAsignacionContext() {}

func NewAsignacionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignacionContext {
	var p = new(AsignacionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Lab2_GParserRULE_asignacion

	return p
}

func (s *AsignacionContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignacionContext) CopyAll(ctx *AsignacionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AsignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SumAsgContext struct {
	AsignacionContext
}

func NewSumAsgContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SumAsgContext {
	var p = new(SumAsgContext)

	InitEmptyAsignacionContext(&p.AsignacionContext)
	p.parser = parser
	p.CopyAll(ctx.(*AsignacionContext))

	return p
}

func (s *SumAsgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SumAsgContext) ID() antlr.TerminalNode {
	return s.GetToken(Lab2_GParserID, 0)
}

func (s *SumAsgContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *SumAsgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.EnterSumAsg(s)
	}
}

func (s *SumAsgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.ExitSumAsg(s)
	}
}

func (s *SumAsgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Lab2_GVisitor:
		return t.VisitSumAsg(s)

	default:
		return t.VisitChildren(s)
	}
}

type AsigContext struct {
	AsignacionContext
}

func NewAsigContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsigContext {
	var p = new(AsigContext)

	InitEmptyAsignacionContext(&p.AsignacionContext)
	p.parser = parser
	p.CopyAll(ctx.(*AsignacionContext))

	return p
}

func (s *AsigContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsigContext) ID() antlr.TerminalNode {
	return s.GetToken(Lab2_GParserID, 0)
}

func (s *AsigContext) ASIG() antlr.TerminalNode {
	return s.GetToken(Lab2_GParserASIG, 0)
}

func (s *AsigContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *AsigContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.EnterAsig(s)
	}
}

func (s *AsigContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.ExitAsig(s)
	}
}

func (s *AsigContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Lab2_GVisitor:
		return t.VisitAsig(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Lab2_GParser) Asignacion() (localctx IAsignacionContext) {
	localctx = NewAsignacionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, Lab2_GParserRULE_asignacion)
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSumAsgContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(47)
			p.Match(Lab2_GParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(48)
			p.Match(Lab2_GParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(49)
			p.e(0)
		}

	case 2:
		localctx = NewAsigContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(50)
			p.Match(Lab2_GParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(51)
			p.Match(Lab2_GParserASIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(52)
			p.e(0)
		}

	case antlr.ATNInvalidAltNumber:
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

// IEContext is an interface to support dynamic dispatch.
type IEContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsEContext differentiates from other interfaces.
	IsEContext()
}

type EContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEContext() *EContext {
	var p = new(EContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Lab2_GParserRULE_e
	return p
}

func InitEmptyEContext(p *EContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Lab2_GParserRULE_e
}

func (*EContext) IsEContext() {}

func NewEContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EContext {
	var p = new(EContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Lab2_GParserRULE_e

	return p
}

func (s *EContext) GetParser() antlr.Parser { return s.parser }

func (s *EContext) CopyAll(ctx *EContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *EContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type E_DecimalContext struct {
	EContext
}

func NewE_DecimalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *E_DecimalContext {
	var p = new(E_DecimalContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *E_DecimalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *E_DecimalContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(Lab2_GParserDECIMAL, 0)
}

func (s *E_DecimalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.EnterE_Decimal(s)
	}
}

func (s *E_DecimalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.ExitE_Decimal(s)
	}
}

func (s *E_DecimalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Lab2_GVisitor:
		return t.VisitE_Decimal(s)

	default:
		return t.VisitChildren(s)
	}
}

type E_EnteroContext struct {
	EContext
}

func NewE_EnteroContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *E_EnteroContext {
	var p = new(E_EnteroContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *E_EnteroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *E_EnteroContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(Lab2_GParserENTERO, 0)
}

func (s *E_EnteroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.EnterE_Entero(s)
	}
}

func (s *E_EnteroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.ExitE_Entero(s)
	}
}

func (s *E_EnteroContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Lab2_GVisitor:
		return t.VisitE_Entero(s)

	default:
		return t.VisitChildren(s)
	}
}

type E_muldivContext struct {
	EContext
	op antlr.Token
}

func NewE_muldivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *E_muldivContext {
	var p = new(E_muldivContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *E_muldivContext) GetOp() antlr.Token { return s.op }

func (s *E_muldivContext) SetOp(v antlr.Token) { s.op = v }

func (s *E_muldivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *E_muldivContext) AllE() []IEContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEContext); ok {
			len++
		}
	}

	tst := make([]IEContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEContext); ok {
			tst[i] = t.(IEContext)
			i++
		}
	}

	return tst
}

func (s *E_muldivContext) E(i int) IEContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
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

	return t.(IEContext)
}

func (s *E_muldivContext) DIV() antlr.TerminalNode {
	return s.GetToken(Lab2_GParserDIV, 0)
}

func (s *E_muldivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.EnterE_muldiv(s)
	}
}

func (s *E_muldivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.ExitE_muldiv(s)
	}
}

func (s *E_muldivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Lab2_GVisitor:
		return t.VisitE_muldiv(s)

	default:
		return t.VisitChildren(s)
	}
}

type E_IdContext struct {
	EContext
}

func NewE_IdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *E_IdContext {
	var p = new(E_IdContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *E_IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *E_IdContext) ID() antlr.TerminalNode {
	return s.GetToken(Lab2_GParserID, 0)
}

func (s *E_IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.EnterE_Id(s)
	}
}

func (s *E_IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.ExitE_Id(s)
	}
}

func (s *E_IdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Lab2_GVisitor:
		return t.VisitE_Id(s)

	default:
		return t.VisitChildren(s)
	}
}

type E_sumresContext struct {
	EContext
	operacion antlr.Token
}

func NewE_sumresContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *E_sumresContext {
	var p = new(E_sumresContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *E_sumresContext) GetOperacion() antlr.Token { return s.operacion }

func (s *E_sumresContext) SetOperacion(v antlr.Token) { s.operacion = v }

func (s *E_sumresContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *E_sumresContext) AllE() []IEContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEContext); ok {
			len++
		}
	}

	tst := make([]IEContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEContext); ok {
			tst[i] = t.(IEContext)
			i++
		}
	}

	return tst
}

func (s *E_sumresContext) E(i int) IEContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
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

	return t.(IEContext)
}

func (s *E_sumresContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.EnterE_sumres(s)
	}
}

func (s *E_sumresContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.ExitE_sumres(s)
	}
}

func (s *E_sumresContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Lab2_GVisitor:
		return t.VisitE_sumres(s)

	default:
		return t.VisitChildren(s)
	}
}

type E_NegContext struct {
	EContext
	op antlr.Token
}

func NewE_NegContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *E_NegContext {
	var p = new(E_NegContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *E_NegContext) GetOp() antlr.Token { return s.op }

func (s *E_NegContext) SetOp(v antlr.Token) { s.op = v }

func (s *E_NegContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *E_NegContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *E_NegContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.EnterE_Neg(s)
	}
}

func (s *E_NegContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.ExitE_Neg(s)
	}
}

func (s *E_NegContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Lab2_GVisitor:
		return t.VisitE_Neg(s)

	default:
		return t.VisitChildren(s)
	}
}

type E_parContext struct {
	EContext
}

func NewE_parContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *E_parContext {
	var p = new(E_parContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *E_parContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *E_parContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(Lab2_GParserPARIZQ, 0)
}

func (s *E_parContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *E_parContext) PARDER() antlr.TerminalNode {
	return s.GetToken(Lab2_GParserPARDER, 0)
}

func (s *E_parContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.EnterE_par(s)
	}
}

func (s *E_parContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.ExitE_par(s)
	}
}

func (s *E_parContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Lab2_GVisitor:
		return t.VisitE_par(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Lab2_GParser) E() (localctx IEContext) {
	return p.e(0)
}

func (p *Lab2_GParser) e(_p int) (localctx IEContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewEContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IEContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, Lab2_GParserRULE_e, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Lab2_GParserID:
		localctx = NewE_IdContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(56)
			p.Match(Lab2_GParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Lab2_GParserDECIMAL:
		localctx = NewE_DecimalContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(57)
			p.Match(Lab2_GParserDECIMAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Lab2_GParserENTERO:
		localctx = NewE_EnteroContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(58)
			p.Match(Lab2_GParserENTERO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Lab2_GParserT__5, Lab2_GParserT__6:
		localctx = NewE_NegContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(59)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*E_NegContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == Lab2_GParserT__5 || _la == Lab2_GParserT__6) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*E_NegContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(60)
			p.e(4)
		}

	case Lab2_GParserPARIZQ:
		localctx = NewE_parContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(61)
			p.Match(Lab2_GParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(62)
			p.e(0)
		}
		{
			p.SetState(63)
			p.Match(Lab2_GParserPARDER)
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
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(73)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
			case 1:
				localctx = NewE_muldivContext(p, NewEContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, Lab2_GParserRULE_e)
				p.SetState(67)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(68)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*E_muldivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == Lab2_GParserT__7 || _la == Lab2_GParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*E_muldivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(69)
					p.e(4)
				}

			case 2:
				localctx = NewE_sumresContext(p, NewEContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, Lab2_GParserRULE_e)
				p.SetState(70)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(71)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*E_sumresContext).operacion = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == Lab2_GParserT__5 || _la == Lab2_GParserT__8) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*E_sumresContext).operacion = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(72)
					p.e(3)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
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

func (p *Lab2_GParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
		var t *EContext = nil
		if localctx != nil {
			t = localctx.(*EContext)
		}
		return p.E_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Lab2_GParser) E_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
