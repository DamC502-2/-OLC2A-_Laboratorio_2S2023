// Generated from Lab0_Parser.g4 by ANTLR 4.13.0
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class Lab0_ParserLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.13.0", RuntimeMetaData.VERSION); }

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
		"\u0004\u0000\u0007D\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002\u0001"+
		"\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004"+
		"\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007"+
		"\u0007\u0007\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0002"+
		"\u0004\u0002\u0017\b\u0002\u000b\u0002\f\u0002\u0018\u0001\u0002\u0001"+
		"\u0002\u0001\u0003\u0001\u0003\u0001\u0004\u0004\u0004 \b\u0004\u000b"+
		"\u0004\f\u0004!\u0001\u0005\u0004\u0005%\b\u0005\u000b\u0005\f\u0005&"+
		"\u0001\u0005\u0001\u0005\u0004\u0005+\b\u0005\u000b\u0005\f\u0005,\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0007\u0001\u0007\u0005\u00078\b\u0007\n\u0007\f\u0007;\t"+
		"\u0007\u0001\u0007\u0001\u0007\u0004\u0007?\b\u0007\u000b\u0007\f\u0007"+
		"@\u0003\u0007C\b\u0007\u0000\u0000\b\u0001\u0001\u0003\u0002\u0005\u0003"+
		"\u0007\u0000\t\u0004\u000b\u0005\r\u0006\u000f\u0007\u0001\u0000\u0004"+
		"\u0003\u0000\t\n\r\r  \u0001\u000009\u0003\u0000AZ__az\u0004\u000009A"+
		"Z__azI\u0000\u0001\u0001\u0000\u0000\u0000\u0000\u0003\u0001\u0000\u0000"+
		"\u0000\u0000\u0005\u0001\u0000\u0000\u0000\u0000\t\u0001\u0000\u0000\u0000"+
		"\u0000\u000b\u0001\u0000\u0000\u0000\u0000\r\u0001\u0000\u0000\u0000\u0000"+
		"\u000f\u0001\u0000\u0000\u0000\u0001\u0011\u0001\u0000\u0000\u0000\u0003"+
		"\u0013\u0001\u0000\u0000\u0000\u0005\u0016\u0001\u0000\u0000\u0000\u0007"+
		"\u001c\u0001\u0000\u0000\u0000\t\u001f\u0001\u0000\u0000\u0000\u000b$"+
		"\u0001\u0000\u0000\u0000\r.\u0001\u0000\u0000\u0000\u000fB\u0001\u0000"+
		"\u0000\u0000\u0011\u0012\u0005)\u0000\u0000\u0012\u0002\u0001\u0000\u0000"+
		"\u0000\u0013\u0014\u0005(\u0000\u0000\u0014\u0004\u0001\u0000\u0000\u0000"+
		"\u0015\u0017\u0007\u0000\u0000\u0000\u0016\u0015\u0001\u0000\u0000\u0000"+
		"\u0017\u0018\u0001\u0000\u0000\u0000\u0018\u0016\u0001\u0000\u0000\u0000"+
		"\u0018\u0019\u0001\u0000\u0000\u0000\u0019\u001a\u0001\u0000\u0000\u0000"+
		"\u001a\u001b\u0006\u0002\u0000\u0000\u001b\u0006\u0001\u0000\u0000\u0000"+
		"\u001c\u001d\u0007\u0001\u0000\u0000\u001d\b\u0001\u0000\u0000\u0000\u001e"+
		" \u0003\u0007\u0003\u0000\u001f\u001e\u0001\u0000\u0000\u0000 !\u0001"+
		"\u0000\u0000\u0000!\u001f\u0001\u0000\u0000\u0000!\"\u0001\u0000\u0000"+
		"\u0000\"\n\u0001\u0000\u0000\u0000#%\u0003\u0007\u0003\u0000$#\u0001\u0000"+
		"\u0000\u0000%&\u0001\u0000\u0000\u0000&$\u0001\u0000\u0000\u0000&\'\u0001"+
		"\u0000\u0000\u0000\'(\u0001\u0000\u0000\u0000(*\u0005.\u0000\u0000)+\u0003"+
		"\u0007\u0003\u0000*)\u0001\u0000\u0000\u0000+,\u0001\u0000\u0000\u0000"+
		",*\u0001\u0000\u0000\u0000,-\u0001\u0000\u0000\u0000-\f\u0001\u0000\u0000"+
		"\u0000./\u0005r\u0000\u0000/0\u0005e\u0000\u000001\u0005t\u0000\u0000"+
		"12\u0005u\u0000\u000023\u0005r\u0000\u000034\u0005n\u0000\u00004\u000e"+
		"\u0001\u0000\u0000\u000059\u0007\u0002\u0000\u000068\u0007\u0003\u0000"+
		"\u000076\u0001\u0000\u0000\u00008;\u0001\u0000\u0000\u000097\u0001\u0000"+
		"\u0000\u00009:\u0001\u0000\u0000\u0000:C\u0001\u0000\u0000\u0000;9\u0001"+
		"\u0000\u0000\u0000<>\u0005_\u0000\u0000=?\u0007\u0003\u0000\u0000>=\u0001"+
		"\u0000\u0000\u0000?@\u0001\u0000\u0000\u0000@>\u0001\u0000\u0000\u0000"+
		"@A\u0001\u0000\u0000\u0000AC\u0001\u0000\u0000\u0000B5\u0001\u0000\u0000"+
		"\u0000B<\u0001\u0000\u0000\u0000C\u0010\u0001\u0000\u0000\u0000\b\u0000"+
		"\u0018!&,9@B\u0001\u0006\u0000\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}