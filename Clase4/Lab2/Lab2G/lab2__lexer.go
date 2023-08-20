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
		"", "", "'return'",
	}
	staticData.SymbolicNames = []string{
		"", "PARDER", "PARIZQ", "ASIG", "DIV", "INT", "CONSOLE", "STRING", "ENBLANCO",
		"ENTERO", "DECIMAL", "ID", "RETURN", "UL_C",
	}
	staticData.RuleNames = []string{
		"PARDER", "PARIZQ", "ASIG", "DIV", "INT", "CONSOLE", "STRING", "ENBLANCO",
		"DIG", "ENTERO", "DECIMAL", "ID", "RETURN", "UL_C",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 13, 120, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 1, 0, 1, 1,
		1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 5, 6,
		56, 8, 6, 10, 6, 12, 6, 59, 9, 6, 1, 6, 1, 6, 1, 7, 4, 7, 64, 8, 7, 11,
		7, 12, 7, 65, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 4, 9, 73, 8, 9, 11, 9, 12,
		9, 74, 1, 10, 4, 10, 78, 8, 10, 11, 10, 12, 10, 79, 1, 10, 1, 10, 4, 10,
		84, 8, 10, 11, 10, 12, 10, 85, 1, 11, 1, 11, 5, 11, 90, 8, 11, 10, 11,
		12, 11, 93, 9, 11, 1, 11, 1, 11, 4, 11, 97, 8, 11, 11, 11, 12, 11, 98,
		3, 11, 101, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		13, 1, 13, 1, 13, 1, 13, 5, 13, 114, 8, 13, 10, 13, 12, 13, 117, 9, 13,
		1, 13, 1, 13, 0, 0, 14, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 0, 19, 9, 21, 10, 23, 11, 25, 12, 27, 13, 1, 0, 6, 1, 0, 34, 34,
		3, 0, 9, 10, 13, 13, 32, 32, 1, 0, 48, 57, 3, 0, 65, 90, 95, 95, 97, 122,
		4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 10, 10, 13, 13, 127, 0, 1,
		1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9,
		1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0,
		19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0,
		0, 27, 1, 0, 0, 0, 1, 29, 1, 0, 0, 0, 3, 31, 1, 0, 0, 0, 5, 33, 1, 0, 0,
		0, 7, 35, 1, 0, 0, 0, 9, 37, 1, 0, 0, 0, 11, 41, 1, 0, 0, 0, 13, 53, 1,
		0, 0, 0, 15, 63, 1, 0, 0, 0, 17, 69, 1, 0, 0, 0, 19, 72, 1, 0, 0, 0, 21,
		77, 1, 0, 0, 0, 23, 100, 1, 0, 0, 0, 25, 102, 1, 0, 0, 0, 27, 109, 1, 0,
		0, 0, 29, 30, 5, 41, 0, 0, 30, 2, 1, 0, 0, 0, 31, 32, 5, 40, 0, 0, 32,
		4, 1, 0, 0, 0, 33, 34, 5, 61, 0, 0, 34, 6, 1, 0, 0, 0, 35, 36, 5, 47, 0,
		0, 36, 8, 1, 0, 0, 0, 37, 38, 5, 105, 0, 0, 38, 39, 5, 110, 0, 0, 39, 40,
		5, 116, 0, 0, 40, 10, 1, 0, 0, 0, 41, 42, 5, 67, 0, 0, 42, 43, 5, 111,
		0, 0, 43, 44, 5, 110, 0, 0, 44, 45, 5, 115, 0, 0, 45, 46, 5, 111, 0, 0,
		46, 47, 5, 108, 0, 0, 47, 48, 5, 101, 0, 0, 48, 49, 5, 46, 0, 0, 49, 50,
		5, 111, 0, 0, 50, 51, 5, 117, 0, 0, 51, 52, 5, 116, 0, 0, 52, 12, 1, 0,
		0, 0, 53, 57, 5, 34, 0, 0, 54, 56, 8, 0, 0, 0, 55, 54, 1, 0, 0, 0, 56,
		59, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 60, 1, 0, 0,
		0, 59, 57, 1, 0, 0, 0, 60, 61, 5, 34, 0, 0, 61, 14, 1, 0, 0, 0, 62, 64,
		7, 1, 0, 0, 63, 62, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0,
		65, 66, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 68, 6, 7, 0, 0, 68, 16, 1,
		0, 0, 0, 69, 70, 7, 2, 0, 0, 70, 18, 1, 0, 0, 0, 71, 73, 3, 17, 8, 0, 72,
		71, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 74, 75, 1, 0, 0,
		0, 75, 20, 1, 0, 0, 0, 76, 78, 3, 17, 8, 0, 77, 76, 1, 0, 0, 0, 78, 79,
		1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0,
		81, 83, 5, 46, 0, 0, 82, 84, 3, 17, 8, 0, 83, 82, 1, 0, 0, 0, 84, 85, 1,
		0, 0, 0, 85, 83, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 22, 1, 0, 0, 0, 87,
		91, 7, 3, 0, 0, 88, 90, 7, 4, 0, 0, 89, 88, 1, 0, 0, 0, 90, 93, 1, 0, 0,
		0, 91, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 101, 1, 0, 0, 0, 93, 91,
		1, 0, 0, 0, 94, 96, 5, 95, 0, 0, 95, 97, 7, 4, 0, 0, 96, 95, 1, 0, 0, 0,
		97, 98, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 101, 1,
		0, 0, 0, 100, 87, 1, 0, 0, 0, 100, 94, 1, 0, 0, 0, 101, 24, 1, 0, 0, 0,
		102, 103, 5, 114, 0, 0, 103, 104, 5, 101, 0, 0, 104, 105, 5, 116, 0, 0,
		105, 106, 5, 117, 0, 0, 106, 107, 5, 114, 0, 0, 107, 108, 5, 110, 0, 0,
		108, 26, 1, 0, 0, 0, 109, 110, 5, 47, 0, 0, 110, 111, 5, 47, 0, 0, 111,
		115, 1, 0, 0, 0, 112, 114, 8, 5, 0, 0, 113, 112, 1, 0, 0, 0, 114, 117,
		1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 118, 1, 0,
		0, 0, 117, 115, 1, 0, 0, 0, 118, 119, 6, 13, 1, 0, 119, 28, 1, 0, 0, 0,
		10, 0, 57, 65, 74, 79, 85, 91, 98, 100, 115, 2, 6, 0, 0, 0, 1, 0,
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
	Lab2_LexerSTRING   = 7
	Lab2_LexerENBLANCO = 8
	Lab2_LexerENTERO   = 9
	Lab2_LexerDECIMAL  = 10
	Lab2_LexerID       = 11
	Lab2_LexerRETURN   = 12
	Lab2_LexerUL_C     = 13
)
