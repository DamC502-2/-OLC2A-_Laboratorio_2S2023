// Code generated from Lab1_Lexer.g4 by ANTLR 4.13.0. DO NOT EDIT.

package Lab1P

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

type Lab1_Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var Lab1_LexerLexerStaticData struct {
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

func lab1_lexerLexerInit() {
	staticData := &Lab1_LexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "')'", "'('", "", "", "", "", "'return'",
	}
	staticData.SymbolicNames = []string{
		"", "PARDER", "PARIZQ", "ENBLANCO", "ENTERO", "DECIMAL", "ID", "RETURN",
		"UL_C",
	}
	staticData.RuleNames = []string{
		"PARDER", "PARIZQ", "ENBLANCO", "DIG", "ENTERO", "DECIMAL", "ID", "RETURN",
		"UL_C",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 8, 81, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 2, 4, 2, 25, 8, 2, 11, 2, 12, 2, 26, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 4, 4, 4, 34, 8, 4, 11, 4, 12, 4, 35, 1, 5, 4, 5, 39, 8, 5, 11, 5, 12,
		5, 40, 1, 5, 1, 5, 4, 5, 45, 8, 5, 11, 5, 12, 5, 46, 1, 6, 1, 6, 5, 6,
		51, 8, 6, 10, 6, 12, 6, 54, 9, 6, 1, 6, 1, 6, 4, 6, 58, 8, 6, 11, 6, 12,
		6, 59, 3, 6, 62, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 8, 1, 8, 5, 8, 75, 8, 8, 10, 8, 12, 8, 78, 9, 8, 1, 8, 1, 8, 0,
		0, 9, 1, 1, 3, 2, 5, 3, 7, 0, 9, 4, 11, 5, 13, 6, 15, 7, 17, 8, 1, 0, 5,
		3, 0, 9, 10, 13, 13, 32, 32, 1, 0, 48, 57, 3, 0, 65, 90, 95, 95, 97, 122,
		4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 10, 10, 13, 13, 87, 0, 1,
		1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11,
		1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 1,
		19, 1, 0, 0, 0, 3, 21, 1, 0, 0, 0, 5, 24, 1, 0, 0, 0, 7, 30, 1, 0, 0, 0,
		9, 33, 1, 0, 0, 0, 11, 38, 1, 0, 0, 0, 13, 61, 1, 0, 0, 0, 15, 63, 1, 0,
		0, 0, 17, 70, 1, 0, 0, 0, 19, 20, 5, 41, 0, 0, 20, 2, 1, 0, 0, 0, 21, 22,
		5, 40, 0, 0, 22, 4, 1, 0, 0, 0, 23, 25, 7, 0, 0, 0, 24, 23, 1, 0, 0, 0,
		25, 26, 1, 0, 0, 0, 26, 24, 1, 0, 0, 0, 26, 27, 1, 0, 0, 0, 27, 28, 1,
		0, 0, 0, 28, 29, 6, 2, 0, 0, 29, 6, 1, 0, 0, 0, 30, 31, 7, 1, 0, 0, 31,
		8, 1, 0, 0, 0, 32, 34, 3, 7, 3, 0, 33, 32, 1, 0, 0, 0, 34, 35, 1, 0, 0,
		0, 35, 33, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 10, 1, 0, 0, 0, 37, 39,
		3, 7, 3, 0, 38, 37, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 38, 1, 0, 0, 0,
		40, 41, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 44, 5, 46, 0, 0, 43, 45, 3,
		7, 3, 0, 44, 43, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 46,
		47, 1, 0, 0, 0, 47, 12, 1, 0, 0, 0, 48, 52, 7, 2, 0, 0, 49, 51, 7, 3, 0,
		0, 50, 49, 1, 0, 0, 0, 51, 54, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 52, 53,
		1, 0, 0, 0, 53, 62, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 55, 57, 5, 95, 0, 0,
		56, 58, 7, 3, 0, 0, 57, 56, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 57, 1,
		0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 62, 1, 0, 0, 0, 61, 48, 1, 0, 0, 0, 61,
		55, 1, 0, 0, 0, 62, 14, 1, 0, 0, 0, 63, 64, 5, 114, 0, 0, 64, 65, 5, 101,
		0, 0, 65, 66, 5, 116, 0, 0, 66, 67, 5, 117, 0, 0, 67, 68, 5, 114, 0, 0,
		68, 69, 5, 110, 0, 0, 69, 16, 1, 0, 0, 0, 70, 71, 5, 47, 0, 0, 71, 72,
		5, 47, 0, 0, 72, 76, 1, 0, 0, 0, 73, 75, 8, 4, 0, 0, 74, 73, 1, 0, 0, 0,
		75, 78, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 79, 1,
		0, 0, 0, 78, 76, 1, 0, 0, 0, 79, 80, 6, 8, 1, 0, 80, 18, 1, 0, 0, 0, 9,
		0, 26, 35, 40, 46, 52, 59, 61, 76, 2, 6, 0, 0, 0, 1, 0,
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

// Lab1_LexerInit initializes any static state used to implement Lab1_Lexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewLab1_Lexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func Lab1_LexerInit() {
	staticData := &Lab1_LexerLexerStaticData
	staticData.once.Do(lab1_lexerLexerInit)
}

// NewLab1_Lexer produces a new lexer instance for the optional input antlr.CharStream.
func NewLab1_Lexer(input antlr.CharStream) *Lab1_Lexer {
	Lab1_LexerInit()
	l := new(Lab1_Lexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &Lab1_LexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Lab1_Lexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Lab1_Lexer tokens.
const (
	Lab1_LexerPARDER   = 1
	Lab1_LexerPARIZQ   = 2
	Lab1_LexerENBLANCO = 3
	Lab1_LexerENTERO   = 4
	Lab1_LexerDECIMAL  = 5
	Lab1_LexerID       = 6
	Lab1_LexerRETURN   = 7
	Lab1_LexerUL_C     = 8
)
