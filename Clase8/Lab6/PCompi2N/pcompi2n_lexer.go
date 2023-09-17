// Code generated from PCompi2N.g4 by ANTLR 4.13.0. DO NOT EDIT.

package PCompi2N

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type PCompi2NLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var PCompi2NLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func pcompi2nlexerLexerInit() {
	staticData := &PCompi2NLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'!'", "'true'", "'false'", "'<'", "'>'", "'=='", "'!='", "'AND'",
		"'OR'", "'-'", "'*'", "'/'", "'+'", "')'", "'('",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "PARDER", "PARIZQ",
		"ENTERO", "DECIMAL", "ENBLANCO", "ID",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "PARDER", "PARIZQ", "DIG", "ENTERO",
		"DECIMAL", "ENBLANCO", "ID",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 19, 123, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 4, 16, 87, 8, 16,
		11, 16, 12, 16, 88, 1, 17, 4, 17, 92, 8, 17, 11, 17, 12, 17, 93, 1, 17,
		1, 17, 4, 17, 98, 8, 17, 11, 17, 12, 17, 99, 1, 18, 4, 18, 103, 8, 18,
		11, 18, 12, 18, 104, 1, 18, 1, 18, 1, 19, 1, 19, 5, 19, 111, 8, 19, 10,
		19, 12, 19, 114, 9, 19, 1, 19, 1, 19, 4, 19, 118, 8, 19, 11, 19, 12, 19,
		119, 3, 19, 122, 8, 19, 0, 0, 20, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6,
		13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31,
		0, 33, 16, 35, 17, 37, 18, 39, 19, 1, 0, 4, 1, 0, 48, 57, 3, 0, 9, 10,
		13, 13, 32, 32, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95,
		95, 97, 122, 128, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0,
		0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0,
		0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0,
		0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1,
		0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39,
		1, 0, 0, 0, 1, 41, 1, 0, 0, 0, 3, 43, 1, 0, 0, 0, 5, 48, 1, 0, 0, 0, 7,
		54, 1, 0, 0, 0, 9, 56, 1, 0, 0, 0, 11, 58, 1, 0, 0, 0, 13, 61, 1, 0, 0,
		0, 15, 64, 1, 0, 0, 0, 17, 68, 1, 0, 0, 0, 19, 71, 1, 0, 0, 0, 21, 73,
		1, 0, 0, 0, 23, 75, 1, 0, 0, 0, 25, 77, 1, 0, 0, 0, 27, 79, 1, 0, 0, 0,
		29, 81, 1, 0, 0, 0, 31, 83, 1, 0, 0, 0, 33, 86, 1, 0, 0, 0, 35, 91, 1,
		0, 0, 0, 37, 102, 1, 0, 0, 0, 39, 121, 1, 0, 0, 0, 41, 42, 5, 33, 0, 0,
		42, 2, 1, 0, 0, 0, 43, 44, 5, 116, 0, 0, 44, 45, 5, 114, 0, 0, 45, 46,
		5, 117, 0, 0, 46, 47, 5, 101, 0, 0, 47, 4, 1, 0, 0, 0, 48, 49, 5, 102,
		0, 0, 49, 50, 5, 97, 0, 0, 50, 51, 5, 108, 0, 0, 51, 52, 5, 115, 0, 0,
		52, 53, 5, 101, 0, 0, 53, 6, 1, 0, 0, 0, 54, 55, 5, 60, 0, 0, 55, 8, 1,
		0, 0, 0, 56, 57, 5, 62, 0, 0, 57, 10, 1, 0, 0, 0, 58, 59, 5, 61, 0, 0,
		59, 60, 5, 61, 0, 0, 60, 12, 1, 0, 0, 0, 61, 62, 5, 33, 0, 0, 62, 63, 5,
		61, 0, 0, 63, 14, 1, 0, 0, 0, 64, 65, 5, 65, 0, 0, 65, 66, 5, 78, 0, 0,
		66, 67, 5, 68, 0, 0, 67, 16, 1, 0, 0, 0, 68, 69, 5, 79, 0, 0, 69, 70, 5,
		82, 0, 0, 70, 18, 1, 0, 0, 0, 71, 72, 5, 45, 0, 0, 72, 20, 1, 0, 0, 0,
		73, 74, 5, 42, 0, 0, 74, 22, 1, 0, 0, 0, 75, 76, 5, 47, 0, 0, 76, 24, 1,
		0, 0, 0, 77, 78, 5, 43, 0, 0, 78, 26, 1, 0, 0, 0, 79, 80, 5, 41, 0, 0,
		80, 28, 1, 0, 0, 0, 81, 82, 5, 40, 0, 0, 82, 30, 1, 0, 0, 0, 83, 84, 7,
		0, 0, 0, 84, 32, 1, 0, 0, 0, 85, 87, 3, 31, 15, 0, 86, 85, 1, 0, 0, 0,
		87, 88, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 34, 1,
		0, 0, 0, 90, 92, 3, 31, 15, 0, 91, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0,
		93, 91, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 97, 5,
		46, 0, 0, 96, 98, 3, 31, 15, 0, 97, 96, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0,
		99, 97, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 36, 1, 0, 0, 0, 101, 103,
		7, 1, 0, 0, 102, 101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 102, 1, 0,
		0, 0, 104, 105, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 107, 6, 18, 0, 0,
		107, 38, 1, 0, 0, 0, 108, 112, 7, 2, 0, 0, 109, 111, 7, 3, 0, 0, 110, 109,
		1, 0, 0, 0, 111, 114, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 112, 113, 1, 0,
		0, 0, 113, 122, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 115, 117, 5, 95, 0, 0,
		116, 118, 7, 3, 0, 0, 117, 116, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119,
		117, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 122, 1, 0, 0, 0, 121, 108,
		1, 0, 0, 0, 121, 115, 1, 0, 0, 0, 122, 40, 1, 0, 0, 0, 8, 0, 88, 93, 99,
		104, 112, 119, 121, 1, 6, 0, 0,
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

// PCompi2NLexerInit initializes any static state used to implement PCompi2NLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewPCompi2NLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func PCompi2NLexerInit() {
	staticData := &PCompi2NLexerLexerStaticData
	staticData.once.Do(pcompi2nlexerLexerInit)
}

// NewPCompi2NLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewPCompi2NLexer(input antlr.CharStream) *PCompi2NLexer {
	PCompi2NLexerInit()
	l := new(PCompi2NLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &PCompi2NLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "PCompi2N.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// PCompi2NLexer tokens.
const (
	PCompi2NLexerT__0     = 1
	PCompi2NLexerT__1     = 2
	PCompi2NLexerT__2     = 3
	PCompi2NLexerT__3     = 4
	PCompi2NLexerT__4     = 5
	PCompi2NLexerT__5     = 6
	PCompi2NLexerT__6     = 7
	PCompi2NLexerT__7     = 8
	PCompi2NLexerT__8     = 9
	PCompi2NLexerT__9     = 10
	PCompi2NLexerT__10    = 11
	PCompi2NLexerT__11    = 12
	PCompi2NLexerT__12    = 13
	PCompi2NLexerPARDER   = 14
	PCompi2NLexerPARIZQ   = 15
	PCompi2NLexerENTERO   = 16
	PCompi2NLexerDECIMAL  = 17
	PCompi2NLexerENBLANCO = 18
	PCompi2NLexerID       = 19
)
