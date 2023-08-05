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
		"", "','", "'*'", "'+'", "'-'", "')'", "'('", "'='", "'/'", "'int'",
		"'Console.out'", "", "", "", "", "'return'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "PARDER", "PARIZQ", "ASIG", "DIV", "INT", "CONSOLE",
		"ENBLANCO", "ENTERO", "DECIMAL", "ID", "RETURN", "UL_C",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "PARDER", "PARIZQ", "ASIG", "DIV", "INT",
		"CONSOLE", "ENBLANCO", "DIG", "ENTERO", "DECIMAL", "ID", "RETURN", "UL_C",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 16, 125, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		10, 4, 10, 69, 8, 10, 11, 10, 12, 10, 70, 1, 10, 1, 10, 1, 11, 1, 11, 1,
		12, 4, 12, 78, 8, 12, 11, 12, 12, 12, 79, 1, 13, 4, 13, 83, 8, 13, 11,
		13, 12, 13, 84, 1, 13, 1, 13, 4, 13, 89, 8, 13, 11, 13, 12, 13, 90, 1,
		14, 1, 14, 5, 14, 95, 8, 14, 10, 14, 12, 14, 98, 9, 14, 1, 14, 1, 14, 4,
		14, 102, 8, 14, 11, 14, 12, 14, 103, 3, 14, 106, 8, 14, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 119,
		8, 16, 10, 16, 12, 16, 122, 9, 16, 1, 16, 1, 16, 0, 0, 17, 1, 1, 3, 2,
		5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 0, 25,
		12, 27, 13, 29, 14, 31, 15, 33, 16, 1, 0, 5, 3, 0, 9, 10, 13, 13, 32, 32,
		1, 0, 48, 57, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95,
		95, 97, 122, 2, 0, 10, 10, 13, 13, 131, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0,
		0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0,
		0, 0, 0, 21, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1,
		0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 1, 35, 1, 0, 0, 0, 3, 37,
		1, 0, 0, 0, 5, 39, 1, 0, 0, 0, 7, 41, 1, 0, 0, 0, 9, 43, 1, 0, 0, 0, 11,
		45, 1, 0, 0, 0, 13, 47, 1, 0, 0, 0, 15, 49, 1, 0, 0, 0, 17, 51, 1, 0, 0,
		0, 19, 55, 1, 0, 0, 0, 21, 68, 1, 0, 0, 0, 23, 74, 1, 0, 0, 0, 25, 77,
		1, 0, 0, 0, 27, 82, 1, 0, 0, 0, 29, 105, 1, 0, 0, 0, 31, 107, 1, 0, 0,
		0, 33, 114, 1, 0, 0, 0, 35, 36, 5, 44, 0, 0, 36, 2, 1, 0, 0, 0, 37, 38,
		5, 42, 0, 0, 38, 4, 1, 0, 0, 0, 39, 40, 5, 43, 0, 0, 40, 6, 1, 0, 0, 0,
		41, 42, 5, 45, 0, 0, 42, 8, 1, 0, 0, 0, 43, 44, 5, 41, 0, 0, 44, 10, 1,
		0, 0, 0, 45, 46, 5, 40, 0, 0, 46, 12, 1, 0, 0, 0, 47, 48, 5, 61, 0, 0,
		48, 14, 1, 0, 0, 0, 49, 50, 5, 47, 0, 0, 50, 16, 1, 0, 0, 0, 51, 52, 5,
		105, 0, 0, 52, 53, 5, 110, 0, 0, 53, 54, 5, 116, 0, 0, 54, 18, 1, 0, 0,
		0, 55, 56, 5, 67, 0, 0, 56, 57, 5, 111, 0, 0, 57, 58, 5, 110, 0, 0, 58,
		59, 5, 115, 0, 0, 59, 60, 5, 111, 0, 0, 60, 61, 5, 108, 0, 0, 61, 62, 5,
		101, 0, 0, 62, 63, 5, 46, 0, 0, 63, 64, 5, 111, 0, 0, 64, 65, 5, 117, 0,
		0, 65, 66, 5, 116, 0, 0, 66, 20, 1, 0, 0, 0, 67, 69, 7, 0, 0, 0, 68, 67,
		1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0,
		71, 72, 1, 0, 0, 0, 72, 73, 6, 10, 0, 0, 73, 22, 1, 0, 0, 0, 74, 75, 7,
		1, 0, 0, 75, 24, 1, 0, 0, 0, 76, 78, 3, 23, 11, 0, 77, 76, 1, 0, 0, 0,
		78, 79, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 26, 1,
		0, 0, 0, 81, 83, 3, 23, 11, 0, 82, 81, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0,
		84, 82, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 88, 5,
		46, 0, 0, 87, 89, 3, 23, 11, 0, 88, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0,
		90, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 28, 1, 0, 0, 0, 92, 96, 7,
		2, 0, 0, 93, 95, 7, 3, 0, 0, 94, 93, 1, 0, 0, 0, 95, 98, 1, 0, 0, 0, 96,
		94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 106, 1, 0, 0, 0, 98, 96, 1, 0,
		0, 0, 99, 101, 5, 95, 0, 0, 100, 102, 7, 3, 0, 0, 101, 100, 1, 0, 0, 0,
		102, 103, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104,
		106, 1, 0, 0, 0, 105, 92, 1, 0, 0, 0, 105, 99, 1, 0, 0, 0, 106, 30, 1,
		0, 0, 0, 107, 108, 5, 114, 0, 0, 108, 109, 5, 101, 0, 0, 109, 110, 5, 116,
		0, 0, 110, 111, 5, 117, 0, 0, 111, 112, 5, 114, 0, 0, 112, 113, 5, 110,
		0, 0, 113, 32, 1, 0, 0, 0, 114, 115, 5, 47, 0, 0, 115, 116, 5, 47, 0, 0,
		116, 120, 1, 0, 0, 0, 117, 119, 8, 4, 0, 0, 118, 117, 1, 0, 0, 0, 119,
		122, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 123,
		1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 123, 124, 6, 16, 1, 0, 124, 34, 1, 0,
		0, 0, 9, 0, 70, 79, 84, 90, 96, 103, 105, 120, 2, 6, 0, 0, 0, 1, 0,
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
	Lab2_GLexerPARDER   = 5
	Lab2_GLexerPARIZQ   = 6
	Lab2_GLexerASIG     = 7
	Lab2_GLexerDIV      = 8
	Lab2_GLexerINT      = 9
	Lab2_GLexerCONSOLE  = 10
	Lab2_GLexerENBLANCO = 11
	Lab2_GLexerENTERO   = 12
	Lab2_GLexerDECIMAL  = 13
	Lab2_GLexerID       = 14
	Lab2_GLexerRETURN   = 15
	Lab2_GLexerUL_C     = 16
)
