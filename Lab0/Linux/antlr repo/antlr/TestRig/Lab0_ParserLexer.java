// Generated from Lab0_Parser.g4 by ANTLR 4.7.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class Lab0_ParserLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.7.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		PARDER=1, PARIZQ=2, ENBLANCO=3, ENTERO=4, DECIMAL=5, RETURN=6, ID=7;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"PARDER", "PARIZQ", "ENBLANCO", "DIG", "ENTERO", "DECIMAL", "RETURN", 
			"ID"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "')'", "'('", null, null, null, "'return'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "PARDER", "PARIZQ", "ENBLANCO", "ENTERO", "DECIMAL", "RETURN", 
			"ID"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public Lab0_ParserLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Lab0_Parser.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\tF\b\1\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\3\2\3\2\3\3\3\3"+
		"\3\4\6\4\31\n\4\r\4\16\4\32\3\4\3\4\3\5\3\5\3\6\6\6\"\n\6\r\6\16\6#\3"+
		"\7\6\7\'\n\7\r\7\16\7(\3\7\3\7\6\7-\n\7\r\7\16\7.\3\b\3\b\3\b\3\b\3\b"+
		"\3\b\3\b\3\t\3\t\7\t:\n\t\f\t\16\t=\13\t\3\t\3\t\6\tA\n\t\r\t\16\tB\5"+
		"\tE\n\t\2\2\n\3\3\5\4\7\5\t\2\13\6\r\7\17\b\21\t\3\2\6\5\2\13\f\17\17"+
		"\"\"\3\2\62;\5\2C\\aac|\6\2\62;C\\aac|\2K\2\3\3\2\2\2\2\5\3\2\2\2\2\7"+
		"\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\3\23\3\2\2"+
		"\2\5\25\3\2\2\2\7\30\3\2\2\2\t\36\3\2\2\2\13!\3\2\2\2\r&\3\2\2\2\17\60"+
		"\3\2\2\2\21D\3\2\2\2\23\24\7+\2\2\24\4\3\2\2\2\25\26\7*\2\2\26\6\3\2\2"+
		"\2\27\31\t\2\2\2\30\27\3\2\2\2\31\32\3\2\2\2\32\30\3\2\2\2\32\33\3\2\2"+
		"\2\33\34\3\2\2\2\34\35\b\4\2\2\35\b\3\2\2\2\36\37\t\3\2\2\37\n\3\2\2\2"+
		" \"\5\t\5\2! \3\2\2\2\"#\3\2\2\2#!\3\2\2\2#$\3\2\2\2$\f\3\2\2\2%\'\5\t"+
		"\5\2&%\3\2\2\2\'(\3\2\2\2(&\3\2\2\2()\3\2\2\2)*\3\2\2\2*,\7\60\2\2+-\5"+
		"\t\5\2,+\3\2\2\2-.\3\2\2\2.,\3\2\2\2./\3\2\2\2/\16\3\2\2\2\60\61\7t\2"+
		"\2\61\62\7g\2\2\62\63\7v\2\2\63\64\7w\2\2\64\65\7t\2\2\65\66\7p\2\2\66"+
		"\20\3\2\2\2\67;\t\4\2\28:\t\5\2\298\3\2\2\2:=\3\2\2\2;9\3\2\2\2;<\3\2"+
		"\2\2<E\3\2\2\2=;\3\2\2\2>@\7a\2\2?A\t\5\2\2@?\3\2\2\2AB\3\2\2\2B@\3\2"+
		"\2\2BC\3\2\2\2CE\3\2\2\2D\67\3\2\2\2D>\3\2\2\2E\22\3\2\2\2\n\2\32#(.;"+
		"BD\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}