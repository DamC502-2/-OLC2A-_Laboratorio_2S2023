// Code generated from PCompi2.g4 by ANTLR 4.13.0. DO NOT EDIT.

package PCompi2

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

type PCompi2Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var PCompi2LexerLexerStaticData struct {
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

func pcompi2lexerLexerInit() {
	staticData := &PCompi2LexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'*'", "'/'", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='", "')'",
		"'('",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "PARDER", "PARIZQ", "ENTERO", "DECIMAL",
		"ENBLANCO", "ID",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "PARDER",
		"PARIZQ", "DIG", "ENTERO", "DECIMAL", "ENBLANCO", "ID",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 14, 95, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 11, 4, 11, 59, 8, 11, 11, 11, 12, 11, 60, 1, 12, 4, 12, 64, 8,
		12, 11, 12, 12, 12, 65, 1, 12, 1, 12, 4, 12, 70, 8, 12, 11, 12, 12, 12,
		71, 1, 13, 4, 13, 75, 8, 13, 11, 13, 12, 13, 76, 1, 13, 1, 13, 1, 14, 1,
		14, 5, 14, 83, 8, 14, 10, 14, 12, 14, 86, 9, 14, 1, 14, 1, 14, 4, 14, 90,
		8, 14, 11, 14, 12, 14, 91, 3, 14, 94, 8, 14, 0, 0, 15, 1, 1, 3, 2, 5, 3,
		7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 0, 23, 11, 25, 12,
		27, 13, 29, 14, 1, 0, 4, 1, 0, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32, 3,
		0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 100,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0,
		0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 1, 31, 1, 0, 0, 0, 3, 33, 1,
		0, 0, 0, 5, 35, 1, 0, 0, 0, 7, 38, 1, 0, 0, 0, 9, 41, 1, 0, 0, 0, 11, 43,
		1, 0, 0, 0, 13, 45, 1, 0, 0, 0, 15, 48, 1, 0, 0, 0, 17, 51, 1, 0, 0, 0,
		19, 53, 1, 0, 0, 0, 21, 55, 1, 0, 0, 0, 23, 58, 1, 0, 0, 0, 25, 63, 1,
		0, 0, 0, 27, 74, 1, 0, 0, 0, 29, 93, 1, 0, 0, 0, 31, 32, 5, 42, 0, 0, 32,
		2, 1, 0, 0, 0, 33, 34, 5, 47, 0, 0, 34, 4, 1, 0, 0, 0, 35, 36, 5, 61, 0,
		0, 36, 37, 5, 61, 0, 0, 37, 6, 1, 0, 0, 0, 38, 39, 5, 33, 0, 0, 39, 40,
		5, 61, 0, 0, 40, 8, 1, 0, 0, 0, 41, 42, 5, 62, 0, 0, 42, 10, 1, 0, 0, 0,
		43, 44, 5, 60, 0, 0, 44, 12, 1, 0, 0, 0, 45, 46, 5, 62, 0, 0, 46, 47, 5,
		61, 0, 0, 47, 14, 1, 0, 0, 0, 48, 49, 5, 60, 0, 0, 49, 50, 5, 61, 0, 0,
		50, 16, 1, 0, 0, 0, 51, 52, 5, 41, 0, 0, 52, 18, 1, 0, 0, 0, 53, 54, 5,
		40, 0, 0, 54, 20, 1, 0, 0, 0, 55, 56, 7, 0, 0, 0, 56, 22, 1, 0, 0, 0, 57,
		59, 3, 21, 10, 0, 58, 57, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 58, 1, 0,
		0, 0, 60, 61, 1, 0, 0, 0, 61, 24, 1, 0, 0, 0, 62, 64, 3, 21, 10, 0, 63,
		62, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0,
		0, 66, 67, 1, 0, 0, 0, 67, 69, 5, 46, 0, 0, 68, 70, 3, 21, 10, 0, 69, 68,
		1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0,
		72, 26, 1, 0, 0, 0, 73, 75, 7, 1, 0, 0, 74, 73, 1, 0, 0, 0, 75, 76, 1,
		0, 0, 0, 76, 74, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78,
		79, 6, 13, 0, 0, 79, 28, 1, 0, 0, 0, 80, 84, 7, 2, 0, 0, 81, 83, 7, 3,
		0, 0, 82, 81, 1, 0, 0, 0, 83, 86, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 84, 85,
		1, 0, 0, 0, 85, 94, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 87, 89, 5, 95, 0, 0,
		88, 90, 7, 3, 0, 0, 89, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 89, 1,
		0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 94, 1, 0, 0, 0, 93, 80, 1, 0, 0, 0, 93,
		87, 1, 0, 0, 0, 94, 30, 1, 0, 0, 0, 8, 0, 60, 65, 71, 76, 84, 91, 93, 1,
		6, 0, 0,
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

// PCompi2LexerInit initializes any static state used to implement PCompi2Lexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewPCompi2Lexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func PCompi2LexerInit() {
	staticData := &PCompi2LexerLexerStaticData
	staticData.once.Do(pcompi2lexerLexerInit)
}

// NewPCompi2Lexer produces a new lexer instance for the optional input antlr.CharStream.
func NewPCompi2Lexer(input antlr.CharStream) *PCompi2Lexer {
	PCompi2LexerInit()
	l := new(PCompi2Lexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &PCompi2LexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "PCompi2.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// PCompi2Lexer tokens.
const (
	PCompi2LexerT__0     = 1
	PCompi2LexerT__1     = 2
	PCompi2LexerT__2     = 3
	PCompi2LexerT__3     = 4
	PCompi2LexerT__4     = 5
	PCompi2LexerT__5     = 6
	PCompi2LexerT__6     = 7
	PCompi2LexerT__7     = 8
	PCompi2LexerPARDER   = 9
	PCompi2LexerPARIZQ   = 10
	PCompi2LexerENTERO   = 11
	PCompi2LexerDECIMAL  = 12
	PCompi2LexerENBLANCO = 13
	PCompi2LexerID       = 14
)
