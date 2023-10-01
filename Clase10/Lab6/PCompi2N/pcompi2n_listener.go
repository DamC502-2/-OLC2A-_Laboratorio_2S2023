// Code generated from PCompi2N.g4 by ANTLR 4.13.0. DO NOT EDIT.

package PCompi2N // PCompi2N
import "github.com/antlr4-go/antlr/v4"

// PCompi2NListener is a complete listener for a parse tree produced by PCompi2NParser.
type PCompi2NListener interface {
	antlr.ParseTreeListener

	// EnterS is called when entering the s production.
	EnterS(c *SContext)

	// EnterIFsen is called when entering the IFsen production.
	EnterIFsen(c *IFsenContext)

	// EnterSwitch is called when entering the switch production.
	EnterSwitch(c *SwitchContext)

	// EnterCase is called when entering the case production.
	EnterCase(c *CaseContext)

	// EnterDefaultc is called when entering the defaultc production.
	EnterDefaultc(c *DefaultcContext)

	// EnterBlock is called when entering the Block production.
	EnterBlock(c *BlockContext)

	// EnterSentenciaEj is called when entering the SentenciaEj production.
	EnterSentenciaEj(c *SentenciaEjContext)

	// EnterSentenciaSwitch is called when entering the SentenciaSwitch production.
	EnterSentenciaSwitch(c *SentenciaSwitchContext)

	// EnterCondOr is called when entering the CondOr production.
	EnterCondOr(c *CondOrContext)

	// EnterCondAnd is called when entering the CondAnd production.
	EnterCondAnd(c *CondAndContext)

	// EnterCondI is called when entering the CondI production.
	EnterCondI(c *CondIContext)

	// EnterCondBool is called when entering the CondBool production.
	EnterCondBool(c *CondBoolContext)

	// EnterCondPar is called when entering the CondPar production.
	EnterCondPar(c *CondParContext)

	// EnterCondOper is called when entering the CondOper production.
	EnterCondOper(c *CondOperContext)

	// EnterEnum is called when entering the Enum production.
	EnterEnum(c *EnumContext)

	// EnterEid is called when entering the Eid production.
	EnterEid(c *EidContext)

	// EnterEsr is called when entering the Esr production.
	EnterEsr(c *EsrContext)

	// EnterEpa is called when entering the Epa production.
	EnterEpa(c *EpaContext)

	// EnterENeg is called when entering the ENeg production.
	EnterENeg(c *ENegContext)

	// EnterEmd is called when entering the Emd production.
	EnterEmd(c *EmdContext)

	// ExitS is called when exiting the s production.
	ExitS(c *SContext)

	// ExitIFsen is called when exiting the IFsen production.
	ExitIFsen(c *IFsenContext)

	// ExitSwitch is called when exiting the switch production.
	ExitSwitch(c *SwitchContext)

	// ExitCase is called when exiting the case production.
	ExitCase(c *CaseContext)

	// ExitDefaultc is called when exiting the defaultc production.
	ExitDefaultc(c *DefaultcContext)

	// ExitBlock is called when exiting the Block production.
	ExitBlock(c *BlockContext)

	// ExitSentenciaEj is called when exiting the SentenciaEj production.
	ExitSentenciaEj(c *SentenciaEjContext)

	// ExitSentenciaSwitch is called when exiting the SentenciaSwitch production.
	ExitSentenciaSwitch(c *SentenciaSwitchContext)

	// ExitCondOr is called when exiting the CondOr production.
	ExitCondOr(c *CondOrContext)

	// ExitCondAnd is called when exiting the CondAnd production.
	ExitCondAnd(c *CondAndContext)

	// ExitCondI is called when exiting the CondI production.
	ExitCondI(c *CondIContext)

	// ExitCondBool is called when exiting the CondBool production.
	ExitCondBool(c *CondBoolContext)

	// ExitCondPar is called when exiting the CondPar production.
	ExitCondPar(c *CondParContext)

	// ExitCondOper is called when exiting the CondOper production.
	ExitCondOper(c *CondOperContext)

	// ExitEnum is called when exiting the Enum production.
	ExitEnum(c *EnumContext)

	// ExitEid is called when exiting the Eid production.
	ExitEid(c *EidContext)

	// ExitEsr is called when exiting the Esr production.
	ExitEsr(c *EsrContext)

	// ExitEpa is called when exiting the Epa production.
	ExitEpa(c *EpaContext)

	// ExitENeg is called when exiting the ENeg production.
	ExitENeg(c *ENegContext)

	// ExitEmd is called when exiting the Emd production.
	ExitEmd(c *EmdContext)
}
