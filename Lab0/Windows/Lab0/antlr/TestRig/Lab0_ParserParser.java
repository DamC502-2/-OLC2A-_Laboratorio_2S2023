// Generated from Lab0_Parser.g4 by ANTLR 4.13.0
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class Lab0_ParserParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.0", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		PARDER=1, PARIZQ=2, ENBLANCO=3, ENTERO=4, DECIMAL=5, RETURN=6, ID=7;
	public static final int
		RULE_s = 0, RULE_a = 1;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "a"
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

	@Override
	public String getGrammarFileName() { return "Lab0_Parser.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public Lab0_ParserParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SContext extends ParserRuleContext {
		public AContext a() {
			return getRuleContext(AContext.class,0);
		}
		public SContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_s; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Lab0_ParserListener ) ((Lab0_ParserListener)listener).enterS(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Lab0_ParserListener ) ((Lab0_ParserListener)listener).exitS(this);
		}
	}

	public final SContext s() throws RecognitionException {
		SContext _localctx = new SContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_s);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(4);
			a();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AContext extends ParserRuleContext {
		public TerminalNode PARIZQ() { return getToken(Lab0_ParserParser.PARIZQ, 0); }
		public AContext a() {
			return getRuleContext(AContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(Lab0_ParserParser.PARDER, 0); }
		public TerminalNode ENTERO() { return getToken(Lab0_ParserParser.ENTERO, 0); }
		public TerminalNode DECIMAL() { return getToken(Lab0_ParserParser.DECIMAL, 0); }
		public TerminalNode ID() { return getToken(Lab0_ParserParser.ID, 0); }
		public AContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_a; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Lab0_ParserListener ) ((Lab0_ParserListener)listener).enterA(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Lab0_ParserListener ) ((Lab0_ParserListener)listener).exitA(this);
		}
	}

	public final AContext a() throws RecognitionException {
		AContext _localctx = new AContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_a);
		try {
			setState(14);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PARIZQ:
				enterOuterAlt(_localctx, 1);
				{
				setState(6);
				match(PARIZQ);
				setState(7);
				a();
				setState(8);
				match(PARDER);
				}
				break;
			case ENTERO:
				enterOuterAlt(_localctx, 2);
				{
				setState(10);
				match(ENTERO);
				}
				break;
			case DECIMAL:
				enterOuterAlt(_localctx, 3);
				{
				setState(11);
				match(DECIMAL);
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 4);
				{
				setState(12);
				match(ID);
				}
				break;
			case EOF:
			case PARDER:
				enterOuterAlt(_localctx, 5);
				{
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\u0004\u0001\u0007\u0011\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001"+
		"\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0003\u0001\u000f\b\u0001"+
		"\u0001\u0001\u0000\u0000\u0002\u0000\u0002\u0000\u0000\u0012\u0000\u0004"+
		"\u0001\u0000\u0000\u0000\u0002\u000e\u0001\u0000\u0000\u0000\u0004\u0005"+
		"\u0003\u0002\u0001\u0000\u0005\u0001\u0001\u0000\u0000\u0000\u0006\u0007"+
		"\u0005\u0002\u0000\u0000\u0007\b\u0003\u0002\u0001\u0000\b\t\u0005\u0001"+
		"\u0000\u0000\t\u000f\u0001\u0000\u0000\u0000\n\u000f\u0005\u0004\u0000"+
		"\u0000\u000b\u000f\u0005\u0005\u0000\u0000\f\u000f\u0005\u0007\u0000\u0000"+
		"\r\u000f\u0001\u0000\u0000\u0000\u000e\u0006\u0001\u0000\u0000\u0000\u000e"+
		"\n\u0001\u0000\u0000\u0000\u000e\u000b\u0001\u0000\u0000\u0000\u000e\f"+
		"\u0001\u0000\u0000\u0000\u000e\r\u0001\u0000\u0000\u0000\u000f\u0003\u0001"+
		"\u0000\u0000\u0000\u0001\u000e";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}