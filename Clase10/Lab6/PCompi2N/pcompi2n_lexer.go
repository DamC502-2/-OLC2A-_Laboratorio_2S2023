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
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "PARDER",
		"PARIZQ", "DIG", "ENTERO", "DECIMAL", "ENBLANCO", "ID",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 30, 196, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 1,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19,
		1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1,
		24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 4, 27, 160, 8, 27, 11, 27, 12, 27,
		161, 1, 28, 4, 28, 165, 8, 28, 11, 28, 12, 28, 166, 1, 28, 1, 28, 4, 28,
		171, 8, 28, 11, 28, 12, 28, 172, 1, 29, 4, 29, 176, 8, 29, 11, 29, 12,
		29, 177, 1, 29, 1, 29, 1, 30, 1, 30, 5, 30, 184, 8, 30, 10, 30, 12, 30,
		187, 9, 30, 1, 30, 1, 30, 4, 30, 191, 8, 30, 11, 30, 12, 30, 192, 3, 30,
		195, 8, 30, 0, 0, 31, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8,
		17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
		35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26,
		53, 0, 55, 27, 57, 28, 59, 29, 61, 30, 1, 0, 4, 1, 0, 48, 57, 3, 0, 9,
		10, 13, 13, 32, 32, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90,
		95, 95, 97, 122, 201, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0,
		0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0,
		0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1,
		0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29,
		1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0,
		37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0,
		0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0,
		0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0,
		0, 0, 1, 63, 1, 0, 0, 0, 3, 66, 1, 0, 0, 0, 5, 71, 1, 0, 0, 0, 7, 78, 1,
		0, 0, 0, 9, 80, 1, 0, 0, 0, 11, 82, 1, 0, 0, 0, 13, 87, 1, 0, 0, 0, 15,
		90, 1, 0, 0, 0, 17, 98, 1, 0, 0, 0, 19, 108, 1, 0, 0, 0, 21, 110, 1, 0,
		0, 0, 23, 115, 1, 0, 0, 0, 25, 121, 1, 0, 0, 0, 27, 123, 1, 0, 0, 0, 29,
		125, 1, 0, 0, 0, 31, 128, 1, 0, 0, 0, 33, 131, 1, 0, 0, 0, 35, 134, 1,
		0, 0, 0, 37, 137, 1, 0, 0, 0, 39, 141, 1, 0, 0, 0, 41, 144, 1, 0, 0, 0,
		43, 146, 1, 0, 0, 0, 45, 148, 1, 0, 0, 0, 47, 150, 1, 0, 0, 0, 49, 152,
		1, 0, 0, 0, 51, 154, 1, 0, 0, 0, 53, 156, 1, 0, 0, 0, 55, 159, 1, 0, 0,
		0, 57, 164, 1, 0, 0, 0, 59, 175, 1, 0, 0, 0, 61, 194, 1, 0, 0, 0, 63, 64,
		5, 105, 0, 0, 64, 65, 5, 102, 0, 0, 65, 2, 1, 0, 0, 0, 66, 67, 5, 101,
		0, 0, 67, 68, 5, 108, 0, 0, 68, 69, 5, 115, 0, 0, 69, 70, 5, 101, 0, 0,
		70, 4, 1, 0, 0, 0, 71, 72, 5, 115, 0, 0, 72, 73, 5, 119, 0, 0, 73, 74,
		5, 105, 0, 0, 74, 75, 5, 116, 0, 0, 75, 76, 5, 99, 0, 0, 76, 77, 5, 104,
		0, 0, 77, 6, 1, 0, 0, 0, 78, 79, 5, 123, 0, 0, 79, 8, 1, 0, 0, 0, 80, 81,
		5, 125, 0, 0, 81, 10, 1, 0, 0, 0, 82, 83, 5, 99, 0, 0, 83, 84, 5, 97, 0,
		0, 84, 85, 5, 115, 0, 0, 85, 86, 5, 101, 0, 0, 86, 12, 1, 0, 0, 0, 87,
		88, 5, 45, 0, 0, 88, 89, 5, 62, 0, 0, 89, 14, 1, 0, 0, 0, 90, 91, 5, 100,
		0, 0, 91, 92, 5, 101, 0, 0, 92, 93, 5, 102, 0, 0, 93, 94, 5, 97, 0, 0,
		94, 95, 5, 117, 0, 0, 95, 96, 5, 108, 0, 0, 96, 97, 5, 116, 0, 0, 97, 16,
		1, 0, 0, 0, 98, 99, 5, 115, 0, 0, 99, 100, 5, 101, 0, 0, 100, 101, 5, 110,
		0, 0, 101, 102, 5, 116, 0, 0, 102, 103, 5, 101, 0, 0, 103, 104, 5, 110,
		0, 0, 104, 105, 5, 99, 0, 0, 105, 106, 5, 105, 0, 0, 106, 107, 5, 97, 0,
		0, 107, 18, 1, 0, 0, 0, 108, 109, 5, 33, 0, 0, 109, 20, 1, 0, 0, 0, 110,
		111, 5, 116, 0, 0, 111, 112, 5, 114, 0, 0, 112, 113, 5, 117, 0, 0, 113,
		114, 5, 101, 0, 0, 114, 22, 1, 0, 0, 0, 115, 116, 5, 102, 0, 0, 116, 117,
		5, 97, 0, 0, 117, 118, 5, 108, 0, 0, 118, 119, 5, 115, 0, 0, 119, 120,
		5, 101, 0, 0, 120, 24, 1, 0, 0, 0, 121, 122, 5, 60, 0, 0, 122, 26, 1, 0,
		0, 0, 123, 124, 5, 62, 0, 0, 124, 28, 1, 0, 0, 0, 125, 126, 5, 61, 0, 0,
		126, 127, 5, 61, 0, 0, 127, 30, 1, 0, 0, 0, 128, 129, 5, 33, 0, 0, 129,
		130, 5, 61, 0, 0, 130, 32, 1, 0, 0, 0, 131, 132, 5, 60, 0, 0, 132, 133,
		5, 61, 0, 0, 133, 34, 1, 0, 0, 0, 134, 135, 5, 62, 0, 0, 135, 136, 5, 61,
		0, 0, 136, 36, 1, 0, 0, 0, 137, 138, 5, 65, 0, 0, 138, 139, 5, 78, 0, 0,
		139, 140, 5, 68, 0, 0, 140, 38, 1, 0, 0, 0, 141, 142, 5, 79, 0, 0, 142,
		143, 5, 82, 0, 0, 143, 40, 1, 0, 0, 0, 144, 145, 5, 45, 0, 0, 145, 42,
		1, 0, 0, 0, 146, 147, 5, 42, 0, 0, 147, 44, 1, 0, 0, 0, 148, 149, 5, 47,
		0, 0, 149, 46, 1, 0, 0, 0, 150, 151, 5, 43, 0, 0, 151, 48, 1, 0, 0, 0,
		152, 153, 5, 41, 0, 0, 153, 50, 1, 0, 0, 0, 154, 155, 5, 40, 0, 0, 155,
		52, 1, 0, 0, 0, 156, 157, 7, 0, 0, 0, 157, 54, 1, 0, 0, 0, 158, 160, 3,
		53, 26, 0, 159, 158, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 159, 1, 0,
		0, 0, 161, 162, 1, 0, 0, 0, 162, 56, 1, 0, 0, 0, 163, 165, 3, 53, 26, 0,
		164, 163, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 164, 1, 0, 0, 0, 166,
		167, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 170, 5, 46, 0, 0, 169, 171,
		3, 53, 26, 0, 170, 169, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 170, 1,
		0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 58, 1, 0, 0, 0, 174, 176, 7, 1, 0,
		0, 175, 174, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 177,
		178, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 180, 6, 29, 0, 0, 180, 60,
		1, 0, 0, 0, 181, 185, 7, 2, 0, 0, 182, 184, 7, 3, 0, 0, 183, 182, 1, 0,
		0, 0, 184, 187, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0,
		186, 195, 1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 188, 190, 5, 95, 0, 0, 189,
		191, 7, 3, 0, 0, 190, 189, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 190,
		1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 195, 1, 0, 0, 0, 194, 181, 1, 0,
		0, 0, 194, 188, 1, 0, 0, 0, 195, 62, 1, 0, 0, 0, 8, 0, 161, 166, 172, 177,
		185, 192, 194, 1, 6, 0, 0,
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
	PCompi2NLexerT__13    = 14
	PCompi2NLexerT__14    = 15
	PCompi2NLexerT__15    = 16
	PCompi2NLexerT__16    = 17
	PCompi2NLexerT__17    = 18
	PCompi2NLexerT__18    = 19
	PCompi2NLexerT__19    = 20
	PCompi2NLexerT__20    = 21
	PCompi2NLexerT__21    = 22
	PCompi2NLexerT__22    = 23
	PCompi2NLexerT__23    = 24
	PCompi2NLexerPARDER   = 25
	PCompi2NLexerPARIZQ   = 26
	PCompi2NLexerENTERO   = 27
	PCompi2NLexerDECIMAL  = 28
	PCompi2NLexerENBLANCO = 29
	PCompi2NLexerID       = 30
)
