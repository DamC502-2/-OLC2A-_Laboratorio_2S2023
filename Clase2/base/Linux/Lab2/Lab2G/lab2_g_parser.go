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
		"", "','", "'*'", "'+'", "'-'", "')'", "'('", "'='", "'/'", "'int'",
		"'Console.out'", "", "", "", "", "'return'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "PARDER", "PARIZQ", "ASIG", "DIV", "INT", "CONSOLE",
		"ENBLANCO", "ENTERO", "DECIMAL", "ID", "RETURN", "UL_C",
	}
	staticData.RuleNames = []string{
		"s", "l_sentencias", "sentencia", "asignacion", "e",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 16, 65, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 1, 0, 1, 0, 1, 1, 5, 1, 14, 8, 1, 10, 1, 12, 1, 17, 9, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 5, 2, 24, 8, 2, 10, 2, 12, 2, 27, 9, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 3, 2, 35, 8, 2, 1, 2, 3, 2, 38, 8, 2, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 52, 8,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 60, 8, 4, 10, 4, 12, 4, 63,
		9, 4, 1, 4, 0, 1, 8, 5, 0, 2, 4, 6, 8, 0, 2, 2, 0, 2, 2, 8, 8, 1, 0, 3,
		4, 69, 0, 10, 1, 0, 0, 0, 2, 15, 1, 0, 0, 0, 4, 37, 1, 0, 0, 0, 6, 39,
		1, 0, 0, 0, 8, 51, 1, 0, 0, 0, 10, 11, 3, 2, 1, 0, 11, 1, 1, 0, 0, 0, 12,
		14, 3, 4, 2, 0, 13, 12, 1, 0, 0, 0, 14, 17, 1, 0, 0, 0, 15, 13, 1, 0, 0,
		0, 15, 16, 1, 0, 0, 0, 16, 3, 1, 0, 0, 0, 17, 15, 1, 0, 0, 0, 18, 19, 5,
		10, 0, 0, 19, 20, 5, 6, 0, 0, 20, 25, 3, 8, 4, 0, 21, 22, 5, 1, 0, 0, 22,
		24, 3, 8, 4, 0, 23, 21, 1, 0, 0, 0, 24, 27, 1, 0, 0, 0, 25, 23, 1, 0, 0,
		0, 25, 26, 1, 0, 0, 0, 26, 28, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 28, 29,
		5, 5, 0, 0, 29, 38, 1, 0, 0, 0, 30, 31, 5, 9, 0, 0, 31, 34, 5, 14, 0, 0,
		32, 33, 5, 7, 0, 0, 33, 35, 3, 8, 4, 0, 34, 32, 1, 0, 0, 0, 34, 35, 1,
		0, 0, 0, 35, 38, 1, 0, 0, 0, 36, 38, 3, 6, 3, 0, 37, 18, 1, 0, 0, 0, 37,
		30, 1, 0, 0, 0, 37, 36, 1, 0, 0, 0, 38, 5, 1, 0, 0, 0, 39, 40, 5, 14, 0,
		0, 40, 41, 5, 7, 0, 0, 41, 42, 3, 8, 4, 0, 42, 7, 1, 0, 0, 0, 43, 44, 6,
		4, -1, 0, 44, 52, 5, 14, 0, 0, 45, 52, 5, 13, 0, 0, 46, 52, 5, 12, 0, 0,
		47, 48, 5, 6, 0, 0, 48, 49, 3, 8, 4, 0, 49, 50, 5, 5, 0, 0, 50, 52, 1,
		0, 0, 0, 51, 43, 1, 0, 0, 0, 51, 45, 1, 0, 0, 0, 51, 46, 1, 0, 0, 0, 51,
		47, 1, 0, 0, 0, 52, 61, 1, 0, 0, 0, 53, 54, 10, 3, 0, 0, 54, 55, 7, 0,
		0, 0, 55, 60, 3, 8, 4, 4, 56, 57, 10, 2, 0, 0, 57, 58, 7, 1, 0, 0, 58,
		60, 3, 8, 4, 3, 59, 53, 1, 0, 0, 0, 59, 56, 1, 0, 0, 0, 60, 63, 1, 0, 0,
		0, 61, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 9, 1, 0, 0, 0, 63, 61, 1,
		0, 0, 0, 7, 15, 25, 34, 37, 51, 59, 61,
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
	Lab2_GParserPARDER   = 5
	Lab2_GParserPARIZQ   = 6
	Lab2_GParserASIG     = 7
	Lab2_GParserDIV      = 8
	Lab2_GParserINT      = 9
	Lab2_GParserCONSOLE  = 10
	Lab2_GParserENBLANCO = 11
	Lab2_GParserENTERO   = 12
	Lab2_GParserDECIMAL  = 13
	Lab2_GParserID       = 14
	Lab2_GParserRETURN   = 15
	Lab2_GParserUL_C     = 16
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

type S_LSentenciasContext struct {
	SContext
}

func NewS_LSentenciasContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_LSentenciasContext {
	var p = new(S_LSentenciasContext)

	InitEmptySContext(&p.SContext)
	p.parser = parser
	p.CopyAll(ctx.(*SContext))

	return p
}

func (s *S_LSentenciasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_LSentenciasContext) L_sentencias() IL_sentenciasContext {
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

func (s *S_LSentenciasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.EnterS_LSentencias(s)
	}
}

func (s *S_LSentenciasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.ExitS_LSentencias(s)
	}
}

func (s *S_LSentenciasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Lab2_GVisitor:
		return t.VisitS_LSentencias(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Lab2_GParser) S() (localctx ISContext) {
	localctx = NewSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Lab2_GParserRULE_s)
	localctx = NewS_LSentenciasContext(p, localctx)
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

type L_SentenciasContext struct {
	L_sentenciasContext
}

func NewL_SentenciasContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *L_SentenciasContext {
	var p = new(L_SentenciasContext)

	InitEmptyL_sentenciasContext(&p.L_sentenciasContext)
	p.parser = parser
	p.CopyAll(ctx.(*L_sentenciasContext))

	return p
}

func (s *L_SentenciasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *L_SentenciasContext) AllSentencia() []ISentenciaContext {
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

func (s *L_SentenciasContext) Sentencia(i int) ISentenciaContext {
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

func (s *L_SentenciasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.EnterL_Sentencias(s)
	}
}

func (s *L_SentenciasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.ExitL_Sentencias(s)
	}
}

func (s *L_SentenciasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Lab2_GVisitor:
		return t.VisitL_Sentencias(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Lab2_GParser) L_sentencias() (localctx IL_sentenciasContext) {
	localctx = NewL_sentenciasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Lab2_GParserRULE_l_sentencias)
	var _la int

	localctx = NewL_SentenciasContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17920) != 0 {
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

type S_ConsolaContext struct {
	SentenciaContext
}

func NewS_ConsolaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_ConsolaContext {
	var p = new(S_ConsolaContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_ConsolaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_ConsolaContext) CONSOLE() antlr.TerminalNode {
	return s.GetToken(Lab2_GParserCONSOLE, 0)
}

func (s *S_ConsolaContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(Lab2_GParserPARIZQ, 0)
}

func (s *S_ConsolaContext) AllE() []IEContext {
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

func (s *S_ConsolaContext) E(i int) IEContext {
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

func (s *S_ConsolaContext) PARDER() antlr.TerminalNode {
	return s.GetToken(Lab2_GParserPARDER, 0)
}

func (s *S_ConsolaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.EnterS_Consola(s)
	}
}

func (s *S_ConsolaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.ExitS_Consola(s)
	}
}

func (s *S_ConsolaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Lab2_GVisitor:
		return t.VisitS_Consola(s)

	default:
		return t.VisitChildren(s)
	}
}

type S_DeclaracionContext struct {
	SentenciaContext
}

func NewS_DeclaracionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_DeclaracionContext {
	var p = new(S_DeclaracionContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_DeclaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_DeclaracionContext) INT() antlr.TerminalNode {
	return s.GetToken(Lab2_GParserINT, 0)
}

func (s *S_DeclaracionContext) ID() antlr.TerminalNode {
	return s.GetToken(Lab2_GParserID, 0)
}

func (s *S_DeclaracionContext) ASIG() antlr.TerminalNode {
	return s.GetToken(Lab2_GParserASIG, 0)
}

func (s *S_DeclaracionContext) E() IEContext {
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

func (s *S_DeclaracionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.EnterS_Declaracion(s)
	}
}

func (s *S_DeclaracionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Lab2_GListener); ok {
		listenerT.ExitS_Declaracion(s)
	}
}

func (s *S_DeclaracionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Lab2_GVisitor:
		return t.VisitS_Declaracion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Lab2_GParser) Sentencia() (localctx ISentenciaContext) {
	localctx = NewSentenciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, Lab2_GParserRULE_sentencia)
	var _la int

	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Lab2_GParserCONSOLE:
		localctx = NewS_ConsolaContext(p, localctx)
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
		localctx = NewS_DeclaracionContext(p, localctx)
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
	localctx = NewAsigContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(39)
		p.Match(Lab2_GParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(40)
		p.Match(Lab2_GParserASIG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(41)
		p.e(0)
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
	operacion antlr.Token
}

func NewE_muldivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *E_muldivContext {
	var p = new(E_muldivContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *E_muldivContext) GetOperacion() antlr.Token { return s.operacion }

func (s *E_muldivContext) SetOperacion(v antlr.Token) { s.operacion = v }

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
	p.SetState(51)
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
			p.SetState(44)
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
			p.SetState(45)
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
			p.SetState(46)
			p.Match(Lab2_GParserENTERO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Lab2_GParserPARIZQ:
		localctx = NewE_parContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(47)
			p.Match(Lab2_GParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(48)
			p.e(0)
		}
		{
			p.SetState(49)
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
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(59)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewE_muldivContext(p, NewEContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, Lab2_GParserRULE_e)
				p.SetState(53)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(54)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*E_muldivContext).operacion = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == Lab2_GParserT__1 || _la == Lab2_GParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*E_muldivContext).operacion = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(55)
					p.e(4)
				}

			case 2:
				localctx = NewE_sumresContext(p, NewEContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, Lab2_GParserRULE_e)
				p.SetState(56)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(57)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*E_sumresContext).operacion = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == Lab2_GParserT__2 || _la == Lab2_GParserT__3) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*E_sumresContext).operacion = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(58)
					p.e(3)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
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
