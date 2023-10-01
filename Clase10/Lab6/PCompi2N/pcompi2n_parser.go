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
		"", "'if'", "'else'", "'switch'", "'{'", "'}'", "'case'", "'->'", "'default'",
		"'sentencia'", "'!'", "'true'", "'false'", "'<'", "'>'", "'=='", "'!='",
		"'<='", "'>='", "'AND'", "'OR'", "'-'", "'*'", "'/'", "'+'", "')'",
		"'('",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "PARDER", "PARIZQ", "ENTERO", "DECIMAL",
		"ENBLANCO", "ID",
	}
	staticData.RuleNames = []string{
		"s", "if", "switch", "case", "defaultc", "blockS", "sentencia", "cond",
		"expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 30, 124, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 28, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 3, 1, 36, 8, 1, 3, 1, 38, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 5,
		2, 44, 8, 2, 10, 2, 12, 2, 47, 9, 2, 1, 2, 3, 2, 50, 8, 2, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 3, 5,
		65, 8, 5, 1, 5, 1, 5, 1, 6, 4, 6, 70, 8, 6, 11, 6, 12, 6, 71, 1, 6, 3,
		6, 75, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 3, 7, 89, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 97,
		8, 7, 10, 7, 12, 7, 100, 9, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 3, 8, 111, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8,
		119, 8, 8, 10, 8, 12, 8, 122, 9, 8, 1, 8, 0, 2, 14, 16, 9, 0, 2, 4, 6,
		8, 10, 12, 14, 16, 0, 4, 1, 0, 11, 12, 1, 0, 13, 18, 1, 0, 22, 23, 2, 0,
		21, 21, 24, 24, 133, 0, 27, 1, 0, 0, 0, 2, 29, 1, 0, 0, 0, 4, 39, 1, 0,
		0, 0, 6, 53, 1, 0, 0, 0, 8, 58, 1, 0, 0, 0, 10, 62, 1, 0, 0, 0, 12, 74,
		1, 0, 0, 0, 14, 88, 1, 0, 0, 0, 16, 110, 1, 0, 0, 0, 18, 19, 3, 2, 1, 0,
		19, 20, 5, 0, 0, 1, 20, 28, 1, 0, 0, 0, 21, 22, 3, 14, 7, 0, 22, 23, 5,
		0, 0, 1, 23, 28, 1, 0, 0, 0, 24, 25, 3, 4, 2, 0, 25, 26, 5, 0, 0, 1, 26,
		28, 1, 0, 0, 0, 27, 18, 1, 0, 0, 0, 27, 21, 1, 0, 0, 0, 27, 24, 1, 0, 0,
		0, 28, 1, 1, 0, 0, 0, 29, 30, 5, 1, 0, 0, 30, 31, 3, 14, 7, 0, 31, 37,
		3, 10, 5, 0, 32, 35, 5, 2, 0, 0, 33, 36, 3, 2, 1, 0, 34, 36, 3, 10, 5,
		0, 35, 33, 1, 0, 0, 0, 35, 34, 1, 0, 0, 0, 36, 38, 1, 0, 0, 0, 37, 32,
		1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 3, 1, 0, 0, 0, 39, 40, 5, 3, 0, 0,
		40, 41, 3, 16, 8, 0, 41, 45, 5, 4, 0, 0, 42, 44, 3, 6, 3, 0, 43, 42, 1,
		0, 0, 0, 44, 47, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46,
		49, 1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 48, 50, 3, 8, 4, 0, 49, 48, 1, 0, 0,
		0, 49, 50, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 52, 5, 5, 0, 0, 52, 5, 1,
		0, 0, 0, 53, 54, 5, 6, 0, 0, 54, 55, 3, 16, 8, 0, 55, 56, 5, 7, 0, 0, 56,
		57, 3, 10, 5, 0, 57, 7, 1, 0, 0, 0, 58, 59, 5, 8, 0, 0, 59, 60, 5, 7, 0,
		0, 60, 61, 3, 10, 5, 0, 61, 9, 1, 0, 0, 0, 62, 64, 5, 4, 0, 0, 63, 65,
		3, 12, 6, 0, 64, 63, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0,
		66, 67, 5, 5, 0, 0, 67, 11, 1, 0, 0, 0, 68, 70, 5, 9, 0, 0, 69, 68, 1,
		0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72,
		75, 1, 0, 0, 0, 73, 75, 3, 4, 2, 0, 74, 69, 1, 0, 0, 0, 74, 73, 1, 0, 0,
		0, 75, 13, 1, 0, 0, 0, 76, 77, 6, 7, -1, 0, 77, 78, 5, 10, 0, 0, 78, 89,
		3, 14, 7, 6, 79, 89, 7, 0, 0, 0, 80, 81, 3, 16, 8, 0, 81, 82, 7, 1, 0,
		0, 82, 83, 3, 16, 8, 0, 83, 89, 1, 0, 0, 0, 84, 85, 5, 26, 0, 0, 85, 86,
		3, 14, 7, 0, 86, 87, 5, 25, 0, 0, 87, 89, 1, 0, 0, 0, 88, 76, 1, 0, 0,
		0, 88, 79, 1, 0, 0, 0, 88, 80, 1, 0, 0, 0, 88, 84, 1, 0, 0, 0, 89, 98,
		1, 0, 0, 0, 90, 91, 10, 3, 0, 0, 91, 92, 5, 19, 0, 0, 92, 97, 3, 14, 7,
		4, 93, 94, 10, 2, 0, 0, 94, 95, 5, 20, 0, 0, 95, 97, 3, 14, 7, 3, 96, 90,
		1, 0, 0, 0, 96, 93, 1, 0, 0, 0, 97, 100, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0,
		98, 99, 1, 0, 0, 0, 99, 15, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 101, 102,
		6, 8, -1, 0, 102, 103, 5, 21, 0, 0, 103, 111, 3, 16, 8, 6, 104, 105, 5,
		26, 0, 0, 105, 106, 3, 16, 8, 0, 106, 107, 5, 25, 0, 0, 107, 111, 1, 0,
		0, 0, 108, 111, 5, 27, 0, 0, 109, 111, 5, 30, 0, 0, 110, 101, 1, 0, 0,
		0, 110, 104, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 110, 109, 1, 0, 0, 0, 111,
		120, 1, 0, 0, 0, 112, 113, 10, 5, 0, 0, 113, 114, 7, 2, 0, 0, 114, 119,
		3, 16, 8, 6, 115, 116, 10, 4, 0, 0, 116, 117, 7, 3, 0, 0, 117, 119, 3,
		16, 8, 5, 118, 112, 1, 0, 0, 0, 118, 115, 1, 0, 0, 0, 119, 122, 1, 0, 0,
		0, 120, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 17, 1, 0, 0, 0, 122,
		120, 1, 0, 0, 0, 14, 27, 35, 37, 45, 49, 64, 71, 74, 88, 96, 98, 110, 118,
		120,
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
	PCompi2NParserT__13    = 14
	PCompi2NParserT__14    = 15
	PCompi2NParserT__15    = 16
	PCompi2NParserT__16    = 17
	PCompi2NParserT__17    = 18
	PCompi2NParserT__18    = 19
	PCompi2NParserT__19    = 20
	PCompi2NParserT__20    = 21
	PCompi2NParserT__21    = 22
	PCompi2NParserT__22    = 23
	PCompi2NParserT__23    = 24
	PCompi2NParserPARDER   = 25
	PCompi2NParserPARIZQ   = 26
	PCompi2NParserENTERO   = 27
	PCompi2NParserDECIMAL  = 28
	PCompi2NParserENBLANCO = 29
	PCompi2NParserID       = 30
)

// PCompi2NParser rules.
const (
	PCompi2NParserRULE_s         = 0
	PCompi2NParserRULE_if        = 1
	PCompi2NParserRULE_switch    = 2
	PCompi2NParserRULE_case      = 3
	PCompi2NParserRULE_defaultc  = 4
	PCompi2NParserRULE_blockS    = 5
	PCompi2NParserRULE_sentencia = 6
	PCompi2NParserRULE_cond      = 7
	PCompi2NParserRULE_expr      = 8
)

// ISContext is an interface to support dynamic dispatch.
type ISContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	If_() IIfContext
	EOF() antlr.TerminalNode
	Cond() ICondContext
	Switch_() ISwitchContext

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

func (s *SContext) If_() IIfContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfContext)
}

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

func (s *SContext) Switch_() ISwitchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchContext)
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
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PCompi2NParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(18)
			p.If_()
		}
		{
			p.SetState(19)
			p.Match(PCompi2NParserEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PCompi2NParserT__9, PCompi2NParserT__10, PCompi2NParserT__11, PCompi2NParserT__20, PCompi2NParserPARIZQ, PCompi2NParserENTERO, PCompi2NParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(21)
			p.cond(0)
		}
		{
			p.SetState(22)
			p.Match(PCompi2NParserEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PCompi2NParserT__2:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(24)
			p.Switch_()
		}
		{
			p.SetState(25)
			p.Match(PCompi2NParserEOF)
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

// IIfContext is an interface to support dynamic dispatch.
type IIfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIfContext differentiates from other interfaces.
	IsIfContext()
}

type IfContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfContext() *IfContext {
	var p = new(IfContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PCompi2NParserRULE_if
	return p
}

func InitEmptyIfContext(p *IfContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PCompi2NParserRULE_if
}

func (*IfContext) IsIfContext() {}

func NewIfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfContext {
	var p = new(IfContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PCompi2NParserRULE_if

	return p
}

func (s *IfContext) GetParser() antlr.Parser { return s.parser }

func (s *IfContext) CopyAll(ctx *IfContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IFsenContext struct {
	IfContext
}

func NewIFsenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IFsenContext {
	var p = new(IFsenContext)

	InitEmptyIfContext(&p.IfContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfContext))

	return p
}

func (s *IFsenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IFsenContext) Cond() ICondContext {
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

func (s *IFsenContext) AllBlockS() []IBlockSContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockSContext); ok {
			len++
		}
	}

	tst := make([]IBlockSContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockSContext); ok {
			tst[i] = t.(IBlockSContext)
			i++
		}
	}

	return tst
}

func (s *IFsenContext) BlockS(i int) IBlockSContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockSContext); ok {
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

	return t.(IBlockSContext)
}

func (s *IFsenContext) If_() IIfContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfContext)
}

func (s *IFsenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.EnterIFsen(s)
	}
}

func (s *IFsenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.ExitIFsen(s)
	}
}

func (s *IFsenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PCompi2NVisitor:
		return t.VisitIFsen(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PCompi2NParser) If_() (localctx IIfContext) {
	localctx = NewIfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PCompi2NParserRULE_if)
	var _la int

	localctx = NewIFsenContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(29)
		p.Match(PCompi2NParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(30)
		p.cond(0)
	}
	{
		p.SetState(31)
		p.BlockS()
	}
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PCompi2NParserT__1 {
		{
			p.SetState(32)
			p.Match(PCompi2NParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case PCompi2NParserT__0:
			{
				p.SetState(33)
				p.If_()
			}

		case PCompi2NParserT__3:
			{
				p.SetState(34)
				p.BlockS()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

// ISwitchContext is an interface to support dynamic dispatch.
type ISwitchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	AllCase_() []ICaseContext
	Case_(i int) ICaseContext
	Defaultc() IDefaultcContext

	// IsSwitchContext differentiates from other interfaces.
	IsSwitchContext()
}

type SwitchContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchContext() *SwitchContext {
	var p = new(SwitchContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PCompi2NParserRULE_switch
	return p
}

func InitEmptySwitchContext(p *SwitchContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PCompi2NParserRULE_switch
}

func (*SwitchContext) IsSwitchContext() {}

func NewSwitchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchContext {
	var p = new(SwitchContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PCompi2NParserRULE_switch

	return p
}

func (s *SwitchContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchContext) Expr() IExprContext {
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

func (s *SwitchContext) AllCase_() []ICaseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICaseContext); ok {
			len++
		}
	}

	tst := make([]ICaseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICaseContext); ok {
			tst[i] = t.(ICaseContext)
			i++
		}
	}

	return tst
}

func (s *SwitchContext) Case_(i int) ICaseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseContext); ok {
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

	return t.(ICaseContext)
}

func (s *SwitchContext) Defaultc() IDefaultcContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefaultcContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefaultcContext)
}

func (s *SwitchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.EnterSwitch(s)
	}
}

func (s *SwitchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.ExitSwitch(s)
	}
}

func (s *SwitchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PCompi2NVisitor:
		return t.VisitSwitch(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PCompi2NParser) Switch_() (localctx ISwitchContext) {
	localctx = NewSwitchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PCompi2NParserRULE_switch)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(39)
		p.Match(PCompi2NParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(40)
		p.expr(0)
	}
	{
		p.SetState(41)
		p.Match(PCompi2NParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PCompi2NParserT__5 {
		{
			p.SetState(42)
			p.Case_()
		}

		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PCompi2NParserT__7 {
		{
			p.SetState(48)
			p.Defaultc()
		}

	}
	{
		p.SetState(51)
		p.Match(PCompi2NParserT__4)
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

// ICaseContext is an interface to support dynamic dispatch.
type ICaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	BlockS() IBlockSContext

	// IsCaseContext differentiates from other interfaces.
	IsCaseContext()
}

type CaseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseContext() *CaseContext {
	var p = new(CaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PCompi2NParserRULE_case
	return p
}

func InitEmptyCaseContext(p *CaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PCompi2NParserRULE_case
}

func (*CaseContext) IsCaseContext() {}

func NewCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseContext {
	var p = new(CaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PCompi2NParserRULE_case

	return p
}

func (s *CaseContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseContext) Expr() IExprContext {
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

func (s *CaseContext) BlockS() IBlockSContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockSContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockSContext)
}

func (s *CaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.EnterCase(s)
	}
}

func (s *CaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.ExitCase(s)
	}
}

func (s *CaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PCompi2NVisitor:
		return t.VisitCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PCompi2NParser) Case_() (localctx ICaseContext) {
	localctx = NewCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PCompi2NParserRULE_case)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Match(PCompi2NParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(54)
		p.expr(0)
	}
	{
		p.SetState(55)
		p.Match(PCompi2NParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(56)
		p.BlockS()
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

// IDefaultcContext is an interface to support dynamic dispatch.
type IDefaultcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BlockS() IBlockSContext

	// IsDefaultcContext differentiates from other interfaces.
	IsDefaultcContext()
}

type DefaultcContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultcContext() *DefaultcContext {
	var p = new(DefaultcContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PCompi2NParserRULE_defaultc
	return p
}

func InitEmptyDefaultcContext(p *DefaultcContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PCompi2NParserRULE_defaultc
}

func (*DefaultcContext) IsDefaultcContext() {}

func NewDefaultcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultcContext {
	var p = new(DefaultcContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PCompi2NParserRULE_defaultc

	return p
}

func (s *DefaultcContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultcContext) BlockS() IBlockSContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockSContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockSContext)
}

func (s *DefaultcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.EnterDefaultc(s)
	}
}

func (s *DefaultcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.ExitDefaultc(s)
	}
}

func (s *DefaultcContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PCompi2NVisitor:
		return t.VisitDefaultc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PCompi2NParser) Defaultc() (localctx IDefaultcContext) {
	localctx = NewDefaultcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PCompi2NParserRULE_defaultc)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.Match(PCompi2NParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(59)
		p.Match(PCompi2NParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(60)
		p.BlockS()
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

// IBlockSContext is an interface to support dynamic dispatch.
type IBlockSContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBlockSContext differentiates from other interfaces.
	IsBlockSContext()
}

type BlockSContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockSContext() *BlockSContext {
	var p = new(BlockSContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PCompi2NParserRULE_blockS
	return p
}

func InitEmptyBlockSContext(p *BlockSContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PCompi2NParserRULE_blockS
}

func (*BlockSContext) IsBlockSContext() {}

func NewBlockSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockSContext {
	var p = new(BlockSContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PCompi2NParserRULE_blockS

	return p
}

func (s *BlockSContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockSContext) CopyAll(ctx *BlockSContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *BlockSContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockSContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BlockContext struct {
	BlockSContext
}

func NewBlockContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlockContext {
	var p = new(BlockContext)

	InitEmptyBlockSContext(&p.BlockSContext)
	p.parser = parser
	p.CopyAll(ctx.(*BlockSContext))

	return p
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) Sentencia() ISentenciaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentenciaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISentenciaContext)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PCompi2NVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PCompi2NParser) BlockS() (localctx IBlockSContext) {
	localctx = NewBlockSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PCompi2NParserRULE_blockS)
	var _la int

	localctx = NewBlockContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		p.Match(PCompi2NParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PCompi2NParserT__2 || _la == PCompi2NParserT__8 {
		{
			p.SetState(63)
			p.Sentencia()
		}

	}
	{
		p.SetState(66)
		p.Match(PCompi2NParserT__4)
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
	p.RuleIndex = PCompi2NParserRULE_sentencia
	return p
}

func InitEmptySentenciaContext(p *SentenciaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PCompi2NParserRULE_sentencia
}

func (*SentenciaContext) IsSentenciaContext() {}

func NewSentenciaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SentenciaContext {
	var p = new(SentenciaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PCompi2NParserRULE_sentencia

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

type SentenciaEjContext struct {
	SentenciaContext
}

func NewSentenciaEjContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SentenciaEjContext {
	var p = new(SentenciaEjContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *SentenciaEjContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentenciaEjContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.EnterSentenciaEj(s)
	}
}

func (s *SentenciaEjContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.ExitSentenciaEj(s)
	}
}

func (s *SentenciaEjContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PCompi2NVisitor:
		return t.VisitSentenciaEj(s)

	default:
		return t.VisitChildren(s)
	}
}

type SentenciaSwitchContext struct {
	SentenciaContext
}

func NewSentenciaSwitchContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SentenciaSwitchContext {
	var p = new(SentenciaSwitchContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *SentenciaSwitchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentenciaSwitchContext) Switch_() ISwitchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchContext)
}

func (s *SentenciaSwitchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.EnterSentenciaSwitch(s)
	}
}

func (s *SentenciaSwitchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PCompi2NListener); ok {
		listenerT.ExitSentenciaSwitch(s)
	}
}

func (s *SentenciaSwitchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PCompi2NVisitor:
		return t.VisitSentenciaSwitch(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PCompi2NParser) Sentencia() (localctx ISentenciaContext) {
	localctx = NewSentenciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PCompi2NParserRULE_sentencia)
	var _la int

	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PCompi2NParserT__8:
		localctx = NewSentenciaEjContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == PCompi2NParserT__8 {
			{
				p.SetState(68)
				p.Match(PCompi2NParserT__8)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(71)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case PCompi2NParserT__2:
		localctx = NewSentenciaSwitchContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(73)
			p.Switch_()
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
	_startState := 14
	p.EnterRecursionRule(localctx, 14, PCompi2NParserRULE_cond, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		localctx = NewCondIContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(77)

			var _m = p.Match(PCompi2NParserT__9)

			localctx.(*CondIContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(78)

			var _x = p.cond(6)

			localctx.(*CondIContext).c = _x
		}

	case 2:
		localctx = NewCondBoolContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(79)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*CondBoolContext).bool_ = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == PCompi2NParserT__10 || _la == PCompi2NParserT__11) {
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
			p.SetState(80)

			var _x = p.expr(0)

			localctx.(*CondOperContext).e1 = _x
		}
		{
			p.SetState(81)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*CondOperContext).opr = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&516096) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*CondOperContext).opr = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(82)

			var _x = p.expr(0)

			localctx.(*CondOperContext).e2 = _x
		}

	case 4:
		localctx = NewCondParContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(84)
			p.Match(PCompi2NParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)

			var _x = p.cond(0)

			localctx.(*CondParContext).c = _x
		}
		{
			p.SetState(86)
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
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(96)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewCondAndContext(p, NewCondContext(p, _parentctx, _parentState))
				localctx.(*CondAndContext).c1 = _prevctx

				p.PushNewRecursionContext(localctx, _startState, PCompi2NParserRULE_cond)
				p.SetState(90)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(91)

					var _m = p.Match(PCompi2NParserT__18)

					localctx.(*CondAndContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(92)

					var _x = p.cond(4)

					localctx.(*CondAndContext).c2 = _x
				}

			case 2:
				localctx = NewCondOrContext(p, NewCondContext(p, _parentctx, _parentState))
				localctx.(*CondOrContext).c1 = _prevctx

				p.PushNewRecursionContext(localctx, _startState, PCompi2NParserRULE_cond)
				p.SetState(93)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(94)

					var _m = p.Match(PCompi2NParserT__19)

					localctx.(*CondOrContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(95)

					var _x = p.cond(3)

					localctx.(*CondOrContext).c2 = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
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
	_startState := 16
	p.EnterRecursionRule(localctx, 16, PCompi2NParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PCompi2NParserT__20:
		localctx = NewENegContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(102)

			var _m = p.Match(PCompi2NParserT__20)

			localctx.(*ENegContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)

			var _x = p.expr(6)

			localctx.(*ENegContext).e1 = _x
		}

	case PCompi2NParserPARIZQ:
		localctx = NewEpaContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(104)
			p.Match(PCompi2NParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(105)

			var _x = p.expr(0)

			localctx.(*EpaContext).e1 = _x
		}
		{
			p.SetState(106)
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
			p.SetState(108)

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
			p.SetState(109)

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
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(118)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
			case 1:
				localctx = NewEmdContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*EmdContext).e1 = _prevctx

				p.PushNewRecursionContext(localctx, _startState, PCompi2NParserRULE_expr)
				p.SetState(112)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(113)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EmdContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == PCompi2NParserT__21 || _la == PCompi2NParserT__22) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EmdContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(114)

					var _x = p.expr(6)

					localctx.(*EmdContext).e2 = _x
				}

			case 2:
				localctx = NewEsrContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*EsrContext).e1 = _prevctx

				p.PushNewRecursionContext(localctx, _startState, PCompi2NParserRULE_expr)
				p.SetState(115)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(116)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EsrContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == PCompi2NParserT__20 || _la == PCompi2NParserT__23) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EsrContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(117)

					var _x = p.expr(5)

					localctx.(*EsrContext).e2 = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
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
	case 7:
		var t *CondContext = nil
		if localctx != nil {
			t = localctx.(*CondContext)
		}
		return p.Cond_Sempred(t, predIndex)

	case 8:
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
