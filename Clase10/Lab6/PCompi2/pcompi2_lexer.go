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
		"", "'!'", "'AND'", "'OR'", "'true'", "'false'", "'-'", "'*'", "'/'",
		"'+'", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='", "')'", "'('",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "PARDER",
		"PARIZQ", "ENTERO", "DECIMAL", "ENBLANCO", "ID",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "PARDER", "PARIZQ",
		"DIG", "ENTERO", "DECIMAL", "ENBLANCO", "ID",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 21, 133, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 4, 18, 97, 8, 18,
		11, 18, 12, 18, 98, 1, 19, 4, 19, 102, 8, 19, 11, 19, 12, 19, 103, 1, 19,
		1, 19, 4, 19, 108, 8, 19, 11, 19, 12, 19, 109, 1, 20, 4, 20, 113, 8, 20,
		11, 20, 12, 20, 114, 1, 20, 1, 20, 1, 21, 1, 21, 5, 21, 121, 8, 21, 10,
		21, 12, 21, 124, 9, 21, 1, 21, 1, 21, 4, 21, 128, 8, 21, 11, 21, 12, 21,
		129, 3, 21, 132, 8, 21, 0, 0, 22, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6,
		13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31,
		16, 33, 17, 35, 0, 37, 18, 39, 19, 41, 20, 43, 21, 1, 0, 4, 1, 0, 48, 57,
		3, 0, 9, 10, 13, 13, 32, 32, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57,
		65, 90, 95, 95, 97, 122, 138, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5,
		1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0,
		0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 37, 1, 0, 0,
		0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 1, 45, 1, 0,
		0, 0, 3, 47, 1, 0, 0, 0, 5, 51, 1, 0, 0, 0, 7, 54, 1, 0, 0, 0, 9, 59, 1,
		0, 0, 0, 11, 65, 1, 0, 0, 0, 13, 67, 1, 0, 0, 0, 15, 69, 1, 0, 0, 0, 17,
		71, 1, 0, 0, 0, 19, 73, 1, 0, 0, 0, 21, 76, 1, 0, 0, 0, 23, 79, 1, 0, 0,
		0, 25, 81, 1, 0, 0, 0, 27, 83, 1, 0, 0, 0, 29, 86, 1, 0, 0, 0, 31, 89,
		1, 0, 0, 0, 33, 91, 1, 0, 0, 0, 35, 93, 1, 0, 0, 0, 37, 96, 1, 0, 0, 0,
		39, 101, 1, 0, 0, 0, 41, 112, 1, 0, 0, 0, 43, 131, 1, 0, 0, 0, 45, 46,
		5, 33, 0, 0, 46, 2, 1, 0, 0, 0, 47, 48, 5, 65, 0, 0, 48, 49, 5, 78, 0,
		0, 49, 50, 5, 68, 0, 0, 50, 4, 1, 0, 0, 0, 51, 52, 5, 79, 0, 0, 52, 53,
		5, 82, 0, 0, 53, 6, 1, 0, 0, 0, 54, 55, 5, 116, 0, 0, 55, 56, 5, 114, 0,
		0, 56, 57, 5, 117, 0, 0, 57, 58, 5, 101, 0, 0, 58, 8, 1, 0, 0, 0, 59, 60,
		5, 102, 0, 0, 60, 61, 5, 97, 0, 0, 61, 62, 5, 108, 0, 0, 62, 63, 5, 115,
		0, 0, 63, 64, 5, 101, 0, 0, 64, 10, 1, 0, 0, 0, 65, 66, 5, 45, 0, 0, 66,
		12, 1, 0, 0, 0, 67, 68, 5, 42, 0, 0, 68, 14, 1, 0, 0, 0, 69, 70, 5, 47,
		0, 0, 70, 16, 1, 0, 0, 0, 71, 72, 5, 43, 0, 0, 72, 18, 1, 0, 0, 0, 73,
		74, 5, 61, 0, 0, 74, 75, 5, 61, 0, 0, 75, 20, 1, 0, 0, 0, 76, 77, 5, 33,
		0, 0, 77, 78, 5, 61, 0, 0, 78, 22, 1, 0, 0, 0, 79, 80, 5, 62, 0, 0, 80,
		24, 1, 0, 0, 0, 81, 82, 5, 60, 0, 0, 82, 26, 1, 0, 0, 0, 83, 84, 5, 62,
		0, 0, 84, 85, 5, 61, 0, 0, 85, 28, 1, 0, 0, 0, 86, 87, 5, 60, 0, 0, 87,
		88, 5, 61, 0, 0, 88, 30, 1, 0, 0, 0, 89, 90, 5, 41, 0, 0, 90, 32, 1, 0,
		0, 0, 91, 92, 5, 40, 0, 0, 92, 34, 1, 0, 0, 0, 93, 94, 7, 0, 0, 0, 94,
		36, 1, 0, 0, 0, 95, 97, 3, 35, 17, 0, 96, 95, 1, 0, 0, 0, 97, 98, 1, 0,
		0, 0, 98, 96, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 38, 1, 0, 0, 0, 100,
		102, 3, 35, 17, 0, 101, 100, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 101,
		1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105, 107, 5, 46,
		0, 0, 106, 108, 3, 35, 17, 0, 107, 106, 1, 0, 0, 0, 108, 109, 1, 0, 0,
		0, 109, 107, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 40, 1, 0, 0, 0, 111,
		113, 7, 1, 0, 0, 112, 111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 112,
		1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 117, 6, 20,
		0, 0, 117, 42, 1, 0, 0, 0, 118, 122, 7, 2, 0, 0, 119, 121, 7, 3, 0, 0,
		120, 119, 1, 0, 0, 0, 121, 124, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 122,
		123, 1, 0, 0, 0, 123, 132, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 125, 127,
		5, 95, 0, 0, 126, 128, 7, 3, 0, 0, 127, 126, 1, 0, 0, 0, 128, 129, 1, 0,
		0, 0, 129, 127, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 132, 1, 0, 0, 0,
		131, 118, 1, 0, 0, 0, 131, 125, 1, 0, 0, 0, 132, 44, 1, 0, 0, 0, 8, 0,
		98, 103, 109, 114, 122, 129, 131, 1, 6, 0, 0,
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
	PCompi2LexerT__8     = 9
	PCompi2LexerT__9     = 10
	PCompi2LexerT__10    = 11
	PCompi2LexerT__11    = 12
	PCompi2LexerT__12    = 13
	PCompi2LexerT__13    = 14
	PCompi2LexerT__14    = 15
	PCompi2LexerPARDER   = 16
	PCompi2LexerPARIZQ   = 17
	PCompi2LexerENTERO   = 18
	PCompi2LexerDECIMAL  = 19
	PCompi2LexerENBLANCO = 20
	PCompi2LexerID       = 21
)
