// Code generated from Lab2_Lexer.g4 by ANTLR 4.13.0. DO NOT EDIT.

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

type Lab2_Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var Lab2_LexerLexerStaticData struct {
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

func lab2_lexerLexerInit() {
	staticData := &Lab2_LexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "')'", "'('", "'='", "'/'", "'int'", "'Console.out'", "", "", "",
		"", "'return'",
	}
	staticData.SymbolicNames = []string{
		"", "PARDER", "PARIZQ", "ASIG", "DIV", "INT", "CONSOLE", "ENBLANCO",
		"ENTERO", "DECIMAL", "ID", "RETURN", "UL_C",
	}
	staticData.RuleNames = []string{
		"PARDER", "PARIZQ", "ASIG", "DIV", "INT", "CONSOLE", "ENBLANCO", "DIG",
		"ENTERO", "DECIMAL", "ID", "RETURN", "UL_C",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 12, 109, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 4, 6, 53, 8, 6, 11, 6, 12,
		6, 54, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 4, 8, 62, 8, 8, 11, 8, 12, 8, 63,
		1, 9, 4, 9, 67, 8, 9, 11, 9, 12, 9, 68, 1, 9, 1, 9, 4, 9, 73, 8, 9, 11,
		9, 12, 9, 74, 1, 10, 1, 10, 5, 10, 79, 8, 10, 10, 10, 12, 10, 82, 9, 10,
		1, 10, 1, 10, 4, 10, 86, 8, 10, 11, 10, 12, 10, 87, 3, 10, 90, 8, 10, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12,
		5, 12, 103, 8, 12, 10, 12, 12, 12, 106, 9, 12, 1, 12, 1, 12, 0, 0, 13,
		1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 0, 17, 8, 19, 9, 21, 10,
		23, 11, 25, 12, 1, 0, 5, 3, 0, 9, 10, 13, 13, 32, 32, 1, 0, 48, 57, 3,
		0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0,
		10, 10, 13, 13, 115, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0,
		0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0,
		0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1,
		0, 0, 0, 0, 25, 1, 0, 0, 0, 1, 27, 1, 0, 0, 0, 3, 29, 1, 0, 0, 0, 5, 31,
		1, 0, 0, 0, 7, 33, 1, 0, 0, 0, 9, 35, 1, 0, 0, 0, 11, 39, 1, 0, 0, 0, 13,
		52, 1, 0, 0, 0, 15, 58, 1, 0, 0, 0, 17, 61, 1, 0, 0, 0, 19, 66, 1, 0, 0,
		0, 21, 89, 1, 0, 0, 0, 23, 91, 1, 0, 0, 0, 25, 98, 1, 0, 0, 0, 27, 28,
		5, 41, 0, 0, 28, 2, 1, 0, 0, 0, 29, 30, 5, 40, 0, 0, 30, 4, 1, 0, 0, 0,
		31, 32, 5, 61, 0, 0, 32, 6, 1, 0, 0, 0, 33, 34, 5, 47, 0, 0, 34, 8, 1,
		0, 0, 0, 35, 36, 5, 105, 0, 0, 36, 37, 5, 110, 0, 0, 37, 38, 5, 116, 0,
		0, 38, 10, 1, 0, 0, 0, 39, 40, 5, 67, 0, 0, 40, 41, 5, 111, 0, 0, 41, 42,
		5, 110, 0, 0, 42, 43, 5, 115, 0, 0, 43, 44, 5, 111, 0, 0, 44, 45, 5, 108,
		0, 0, 45, 46, 5, 101, 0, 0, 46, 47, 5, 46, 0, 0, 47, 48, 5, 111, 0, 0,
		48, 49, 5, 117, 0, 0, 49, 50, 5, 116, 0, 0, 50, 12, 1, 0, 0, 0, 51, 53,
		7, 0, 0, 0, 52, 51, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0,
		54, 55, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 57, 6, 6, 0, 0, 57, 14, 1,
		0, 0, 0, 58, 59, 7, 1, 0, 0, 59, 16, 1, 0, 0, 0, 60, 62, 3, 15, 7, 0, 61,
		60, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 63, 64, 1, 0, 0,
		0, 64, 18, 1, 0, 0, 0, 65, 67, 3, 15, 7, 0, 66, 65, 1, 0, 0, 0, 67, 68,
		1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0,
		70, 72, 5, 46, 0, 0, 71, 73, 3, 15, 7, 0, 72, 71, 1, 0, 0, 0, 73, 74, 1,
		0, 0, 0, 74, 72, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 20, 1, 0, 0, 0, 76,
		80, 7, 2, 0, 0, 77, 79, 7, 3, 0, 0, 78, 77, 1, 0, 0, 0, 79, 82, 1, 0, 0,
		0, 80, 78, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 90, 1, 0, 0, 0, 82, 80,
		1, 0, 0, 0, 83, 85, 5, 95, 0, 0, 84, 86, 7, 3, 0, 0, 85, 84, 1, 0, 0, 0,
		86, 87, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 90, 1,
		0, 0, 0, 89, 76, 1, 0, 0, 0, 89, 83, 1, 0, 0, 0, 90, 22, 1, 0, 0, 0, 91,
		92, 5, 114, 0, 0, 92, 93, 5, 101, 0, 0, 93, 94, 5, 116, 0, 0, 94, 95, 5,
		117, 0, 0, 95, 96, 5, 114, 0, 0, 96, 97, 5, 110, 0, 0, 97, 24, 1, 0, 0,
		0, 98, 99, 5, 47, 0, 0, 99, 100, 5, 47, 0, 0, 100, 104, 1, 0, 0, 0, 101,
		103, 8, 4, 0, 0, 102, 101, 1, 0, 0, 0, 103, 106, 1, 0, 0, 0, 104, 102,
		1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105, 107, 1, 0, 0, 0, 106, 104, 1, 0,
		0, 0, 107, 108, 6, 12, 1, 0, 108, 26, 1, 0, 0, 0, 9, 0, 54, 63, 68, 74,
		80, 87, 89, 104, 2, 6, 0, 0, 0, 1, 0,
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

// Lab2_LexerInit initializes any static state used to implement Lab2_Lexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewLab2_Lexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func Lab2_LexerInit() {
	staticData := &Lab2_LexerLexerStaticData
	staticData.once.Do(lab2_lexerLexerInit)
}

// NewLab2_Lexer produces a new lexer instance for the optional input antlr.CharStream.
func NewLab2_Lexer(input antlr.CharStream) *Lab2_Lexer {
	Lab2_LexerInit()
	l := new(Lab2_Lexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &Lab2_LexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Lab2_Lexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Lab2_Lexer tokens.
const (
	Lab2_LexerPARDER   = 1
	Lab2_LexerPARIZQ   = 2
	Lab2_LexerASIG     = 3
	Lab2_LexerDIV      = 4
	Lab2_LexerINT      = 5
	Lab2_LexerCONSOLE  = 6
	Lab2_LexerENBLANCO = 7
	Lab2_LexerENTERO   = 8
	Lab2_LexerDECIMAL  = 9
	Lab2_LexerID       = 10
	Lab2_LexerRETURN   = 11
	Lab2_LexerUL_C     = 12
)
