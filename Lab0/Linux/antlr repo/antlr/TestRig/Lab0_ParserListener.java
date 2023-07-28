// Generated from Lab0_Parser.g4 by ANTLR 4.7.2
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link Lab0_ParserParser}.
 */
public interface Lab0_ParserListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link Lab0_ParserParser#s}.
	 * @param ctx the parse tree
	 */
	void enterS(Lab0_ParserParser.SContext ctx);
	/**
	 * Exit a parse tree produced by {@link Lab0_ParserParser#s}.
	 * @param ctx the parse tree
	 */
	void exitS(Lab0_ParserParser.SContext ctx);
	/**
	 * Enter a parse tree produced by {@link Lab0_ParserParser#a}.
	 * @param ctx the parse tree
	 */
	void enterA(Lab0_ParserParser.AContext ctx);
	/**
	 * Exit a parse tree produced by {@link Lab0_ParserParser#a}.
	 * @param ctx the parse tree
	 */
	void exitA(Lab0_ParserParser.AContext ctx);
}