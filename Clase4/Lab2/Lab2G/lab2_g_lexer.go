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
		"", "','", "'+='", "'-'", "'!'", "'*'", "'+'", "')'", "'('", "'='",
		"'/'", "'int'", "'Console.out'", "", "", "", "", "", "'return'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "PARDER", "PARIZQ", "ASIG", "DIV", "INT",
		"CONSOLE", "STRING", "ENBLANCO", "ENTERO", "DECIMAL", "ID", "RETURN",
		"UL_C",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "PARDER", "PARIZQ",
		"ASIG", "DIV", "INT", "CONSOLE", "STRING", "ENBLANCO", "DIG", "ENTERO",
		"DECIMAL", "ID", "RETURN", "UL_C",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 19, 145, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 12, 1, 12, 5, 12, 81, 8, 12, 10, 12, 12, 12, 84, 9, 12, 1,
		12, 1, 12, 1, 13, 4, 13, 89, 8, 13, 11, 13, 12, 13, 90, 1, 13, 1, 13, 1,
		14, 1, 14, 1, 15, 4, 15, 98, 8, 15, 11, 15, 12, 15, 99, 1, 16, 4, 16, 103,
		8, 16, 11, 16, 12, 16, 104, 1, 16, 1, 16, 4, 16, 109, 8, 16, 11, 16, 12,
		16, 110, 1, 17, 1, 17, 5, 17, 115, 8, 17, 10, 17, 12, 17, 118, 9, 17, 1,
		17, 1, 17, 4, 17, 122, 8, 17, 11, 17, 12, 17, 123, 3, 17, 126, 8, 17, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19,
		5, 19, 139, 8, 19, 10, 19, 12, 19, 142, 9, 19, 1, 19, 1, 19, 0, 0, 20,
		1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 0, 31, 15, 33, 16, 35, 17, 37, 18, 39, 19,
		1, 0, 6, 1, 0, 34, 34, 3, 0, 9, 10, 13, 13, 32, 32, 1, 0, 48, 57, 3, 0,
		65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 10,
		10, 13, 13, 152, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0,
		0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0,
		0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0,
		0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 31, 1,
		0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39,
		1, 0, 0, 0, 1, 41, 1, 0, 0, 0, 3, 43, 1, 0, 0, 0, 5, 46, 1, 0, 0, 0, 7,
		48, 1, 0, 0, 0, 9, 50, 1, 0, 0, 0, 11, 52, 1, 0, 0, 0, 13, 54, 1, 0, 0,
		0, 15, 56, 1, 0, 0, 0, 17, 58, 1, 0, 0, 0, 19, 60, 1, 0, 0, 0, 21, 62,
		1, 0, 0, 0, 23, 66, 1, 0, 0, 0, 25, 78, 1, 0, 0, 0, 27, 88, 1, 0, 0, 0,
		29, 94, 1, 0, 0, 0, 31, 97, 1, 0, 0, 0, 33, 102, 1, 0, 0, 0, 35, 125, 1,
		0, 0, 0, 37, 127, 1, 0, 0, 0, 39, 134, 1, 0, 0, 0, 41, 42, 5, 44, 0, 0,
		42, 2, 1, 0, 0, 0, 43, 44, 5, 43, 0, 0, 44, 45, 5, 61, 0, 0, 45, 4, 1,
		0, 0, 0, 46, 47, 5, 45, 0, 0, 47, 6, 1, 0, 0, 0, 48, 49, 5, 33, 0, 0, 49,
		8, 1, 0, 0, 0, 50, 51, 5, 42, 0, 0, 51, 10, 1, 0, 0, 0, 52, 53, 5, 43,
		0, 0, 53, 12, 1, 0, 0, 0, 54, 55, 5, 41, 0, 0, 55, 14, 1, 0, 0, 0, 56,
		57, 5, 40, 0, 0, 57, 16, 1, 0, 0, 0, 58, 59, 5, 61, 0, 0, 59, 18, 1, 0,
		0, 0, 60, 61, 5, 47, 0, 0, 61, 20, 1, 0, 0, 0, 62, 63, 5, 105, 0, 0, 63,
		64, 5, 110, 0, 0, 64, 65, 5, 116, 0, 0, 65, 22, 1, 0, 0, 0, 66, 67, 5,
		67, 0, 0, 67, 68, 5, 111, 0, 0, 68, 69, 5, 110, 0, 0, 69, 70, 5, 115, 0,
		0, 70, 71, 5, 111, 0, 0, 71, 72, 5, 108, 0, 0, 72, 73, 5, 101, 0, 0, 73,
		74, 5, 46, 0, 0, 74, 75, 5, 111, 0, 0, 75, 76, 5, 117, 0, 0, 76, 77, 5,
		116, 0, 0, 77, 24, 1, 0, 0, 0, 78, 82, 5, 34, 0, 0, 79, 81, 8, 0, 0, 0,
		80, 79, 1, 0, 0, 0, 81, 84, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 82, 83, 1,
		0, 0, 0, 83, 85, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 85, 86, 5, 34, 0, 0, 86,
		26, 1, 0, 0, 0, 87, 89, 7, 1, 0, 0, 88, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0,
		0, 90, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 93,
		6, 13, 0, 0, 93, 28, 1, 0, 0, 0, 94, 95, 7, 2, 0, 0, 95, 30, 1, 0, 0, 0,
		96, 98, 3, 29, 14, 0, 97, 96, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 97, 1,
		0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 32, 1, 0, 0, 0, 101, 103, 3, 29, 14,
		0, 102, 101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 102, 1, 0, 0, 0, 104,
		105, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 108, 5, 46, 0, 0, 107, 109,
		3, 29, 14, 0, 108, 107, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 108, 1,
		0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 34, 1, 0, 0, 0, 112, 116, 7, 3, 0,
		0, 113, 115, 7, 4, 0, 0, 114, 113, 1, 0, 0, 0, 115, 118, 1, 0, 0, 0, 116,
		114, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 126, 1, 0, 0, 0, 118, 116,
		1, 0, 0, 0, 119, 121, 5, 95, 0, 0, 120, 122, 7, 4, 0, 0, 121, 120, 1, 0,
		0, 0, 122, 123, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0,
		124, 126, 1, 0, 0, 0, 125, 112, 1, 0, 0, 0, 125, 119, 1, 0, 0, 0, 126,
		36, 1, 0, 0, 0, 127, 128, 5, 114, 0, 0, 128, 129, 5, 101, 0, 0, 129, 130,
		5, 116, 0, 0, 130, 131, 5, 117, 0, 0, 131, 132, 5, 114, 0, 0, 132, 133,
		5, 110, 0, 0, 133, 38, 1, 0, 0, 0, 134, 135, 5, 47, 0, 0, 135, 136, 5,
		47, 0, 0, 136, 140, 1, 0, 0, 0, 137, 139, 8, 5, 0, 0, 138, 137, 1, 0, 0,
		0, 139, 142, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141,
		143, 1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 143, 144, 6, 19, 1, 0, 144, 40,
		1, 0, 0, 0, 10, 0, 82, 90, 99, 104, 110, 116, 123, 125, 140, 2, 6, 0, 0,
		0, 1, 0,
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
	Lab2_GLexerPARDER   = 7
	Lab2_GLexerPARIZQ   = 8
	Lab2_GLexerASIG     = 9
	Lab2_GLexerDIV      = 10
	Lab2_GLexerINT      = 11
	Lab2_GLexerCONSOLE  = 12
	Lab2_GLexerSTRING   = 13
	Lab2_GLexerENBLANCO = 14
	Lab2_GLexerENTERO   = 15
	Lab2_GLexerDECIMAL  = 16
	Lab2_GLexerID       = 17
	Lab2_GLexerRETURN   = 18
	Lab2_GLexerUL_C     = 19
)
