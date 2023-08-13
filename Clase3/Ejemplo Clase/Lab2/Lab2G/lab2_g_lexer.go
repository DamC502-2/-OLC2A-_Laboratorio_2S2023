// Code generated from Lab2_G.g4 by ANTLR 4.13.0. DO NOT EDIT.

package Lab2G

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

type Lab2_GLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var Lab2_GLexerLexerStaticData struct {
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

func lab2_glexerLexerInit() {
	staticData := &Lab2_GLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"PARDER", "PARIZQ", "ASIG", "DIV", "INT", "CONSOLE", "ENBLANCO", "DIG",
		"ENTERO", "DECIMAL", "ID", "RETURN", "UL_C",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 21, 150, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 4, 15, 94, 8, 15, 11, 15, 12, 15,
		95, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 4, 17, 103, 8, 17, 11, 17, 12, 17,
		104, 1, 18, 4, 18, 108, 8, 18, 11, 18, 12, 18, 109, 1, 18, 1, 18, 4, 18,
		114, 8, 18, 11, 18, 12, 18, 115, 1, 19, 1, 19, 5, 19, 120, 8, 19, 10, 19,
		12, 19, 123, 9, 19, 1, 19, 1, 19, 4, 19, 127, 8, 19, 11, 19, 12, 19, 128,
		3, 19, 131, 8, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		21, 1, 21, 1, 21, 1, 21, 5, 21, 144, 8, 21, 10, 21, 12, 21, 147, 9, 21,
		1, 21, 1, 21, 0, 0, 22, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 0,
		35, 17, 37, 18, 39, 19, 41, 20, 43, 21, 1, 0, 5, 3, 0, 9, 10, 13, 13, 32,
		32, 1, 0, 48, 57, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90,
		95, 95, 97, 122, 2, 0, 10, 10, 13, 13, 156, 0, 1, 1, 0, 0, 0, 0, 3, 1,
		0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1,
		0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19,
		1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0,
		27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0,
		0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0,
		0, 1, 45, 1, 0, 0, 0, 3, 47, 1, 0, 0, 0, 5, 53, 1, 0, 0, 0, 7, 55, 1, 0,
		0, 0, 9, 57, 1, 0, 0, 0, 11, 60, 1, 0, 0, 0, 13, 62, 1, 0, 0, 0, 15, 64,
		1, 0, 0, 0, 17, 66, 1, 0, 0, 0, 19, 68, 1, 0, 0, 0, 21, 70, 1, 0, 0, 0,
		23, 72, 1, 0, 0, 0, 25, 74, 1, 0, 0, 0, 27, 76, 1, 0, 0, 0, 29, 80, 1,
		0, 0, 0, 31, 93, 1, 0, 0, 0, 33, 99, 1, 0, 0, 0, 35, 102, 1, 0, 0, 0, 37,
		107, 1, 0, 0, 0, 39, 130, 1, 0, 0, 0, 41, 132, 1, 0, 0, 0, 43, 139, 1,
		0, 0, 0, 45, 46, 5, 44, 0, 0, 46, 2, 1, 0, 0, 0, 47, 48, 5, 119, 0, 0,
		48, 49, 5, 104, 0, 0, 49, 50, 5, 105, 0, 0, 50, 51, 5, 108, 0, 0, 51, 52,
		5, 101, 0, 0, 52, 4, 1, 0, 0, 0, 53, 54, 5, 123, 0, 0, 54, 6, 1, 0, 0,
		0, 55, 56, 5, 125, 0, 0, 56, 8, 1, 0, 0, 0, 57, 58, 5, 43, 0, 0, 58, 59,
		5, 61, 0, 0, 59, 10, 1, 0, 0, 0, 60, 61, 5, 45, 0, 0, 61, 12, 1, 0, 0,
		0, 62, 63, 5, 33, 0, 0, 63, 14, 1, 0, 0, 0, 64, 65, 5, 42, 0, 0, 65, 16,
		1, 0, 0, 0, 66, 67, 5, 43, 0, 0, 67, 18, 1, 0, 0, 0, 68, 69, 5, 41, 0,
		0, 69, 20, 1, 0, 0, 0, 70, 71, 5, 40, 0, 0, 71, 22, 1, 0, 0, 0, 72, 73,
		5, 61, 0, 0, 73, 24, 1, 0, 0, 0, 74, 75, 5, 47, 0, 0, 75, 26, 1, 0, 0,
		0, 76, 77, 5, 105, 0, 0, 77, 78, 5, 110, 0, 0, 78, 79, 5, 116, 0, 0, 79,
		28, 1, 0, 0, 0, 80, 81, 5, 67, 0, 0, 81, 82, 5, 111, 0, 0, 82, 83, 5, 110,
		0, 0, 83, 84, 5, 115, 0, 0, 84, 85, 5, 111, 0, 0, 85, 86, 5, 108, 0, 0,
		86, 87, 5, 101, 0, 0, 87, 88, 5, 46, 0, 0, 88, 89, 5, 111, 0, 0, 89, 90,
		5, 117, 0, 0, 90, 91, 5, 116, 0, 0, 91, 30, 1, 0, 0, 0, 92, 94, 7, 0, 0,
		0, 93, 92, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 95, 96,
		1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 98, 6, 15, 0, 0, 98, 32, 1, 0, 0, 0,
		99, 100, 7, 1, 0, 0, 100, 34, 1, 0, 0, 0, 101, 103, 3, 33, 16, 0, 102,
		101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 102, 1, 0, 0, 0, 104, 105,
		1, 0, 0, 0, 105, 36, 1, 0, 0, 0, 106, 108, 3, 33, 16, 0, 107, 106, 1, 0,
		0, 0, 108, 109, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0,
		110, 111, 1, 0, 0, 0, 111, 113, 5, 46, 0, 0, 112, 114, 3, 33, 16, 0, 113,
		112, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 115, 116,
		1, 0, 0, 0, 116, 38, 1, 0, 0, 0, 117, 121, 7, 2, 0, 0, 118, 120, 7, 3,
		0, 0, 119, 118, 1, 0, 0, 0, 120, 123, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0,
		121, 122, 1, 0, 0, 0, 122, 131, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 124,
		126, 5, 95, 0, 0, 125, 127, 7, 3, 0, 0, 126, 125, 1, 0, 0, 0, 127, 128,
		1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 131, 1, 0,
		0, 0, 130, 117, 1, 0, 0, 0, 130, 124, 1, 0, 0, 0, 131, 40, 1, 0, 0, 0,
		132, 133, 5, 114, 0, 0, 133, 134, 5, 101, 0, 0, 134, 135, 5, 116, 0, 0,
		135, 136, 5, 117, 0, 0, 136, 137, 5, 114, 0, 0, 137, 138, 5, 110, 0, 0,
		138, 42, 1, 0, 0, 0, 139, 140, 5, 47, 0, 0, 140, 141, 5, 47, 0, 0, 141,
		145, 1, 0, 0, 0, 142, 144, 8, 4, 0, 0, 143, 142, 1, 0, 0, 0, 144, 147,
		1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 148, 1, 0,
		0, 0, 147, 145, 1, 0, 0, 0, 148, 149, 6, 21, 1, 0, 149, 44, 1, 0, 0, 0,
		9, 0, 95, 104, 109, 115, 121, 128, 130, 145, 2, 6, 0, 0, 0, 1, 0,
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

// Lab2_GLexerInit initializes any static state used to implement Lab2_GLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewLab2_GLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func Lab2_GLexerInit() {
	staticData := &Lab2_GLexerLexerStaticData
	staticData.once.Do(lab2_glexerLexerInit)
}

// NewLab2_GLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewLab2_GLexer(input antlr.CharStream) *Lab2_GLexer {
	Lab2_GLexerInit()
	l := new(Lab2_GLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &Lab2_GLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Lab2_G.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Lab2_GLexer tokens.
const (
	Lab2_GLexerT__0     = 1
	Lab2_GLexerT__1     = 2
	Lab2_GLexerT__2     = 3
	Lab2_GLexerT__3     = 4
	Lab2_GLexerT__4     = 5
	Lab2_GLexerT__5     = 6
	Lab2_GLexerT__6     = 7
	Lab2_GLexerT__7     = 8
	Lab2_GLexerT__8     = 9
	Lab2_GLexerPARDER   = 10
	Lab2_GLexerPARIZQ   = 11
	Lab2_GLexerASIG     = 12
	Lab2_GLexerDIV      = 13
	Lab2_GLexerINT      = 14
	Lab2_GLexerCONSOLE  = 15
	Lab2_GLexerENBLANCO = 16
	Lab2_GLexerENTERO   = 17
	Lab2_GLexerDECIMAL  = 18
	Lab2_GLexerID       = 19
	Lab2_GLexerRETURN   = 20
	Lab2_GLexerUL_C     = 21
)
