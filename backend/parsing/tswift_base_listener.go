// Code generated from .\parsing\tswift.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parsing // tswift
import "github.com/antlr4-go/antlr/v4"

// BasetswiftListener is a complete listener for a parse tree produced by tswiftParser.
type BasetswiftListener struct{}

var _ tswiftListener = &BasetswiftListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasetswiftListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasetswiftListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasetswiftListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasetswiftListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterJson is called when production json is entered.
func (s *BasetswiftListener) EnterJson(ctx *JsonContext) {}

// ExitJson is called when production json is exited.
func (s *BasetswiftListener) ExitJson(ctx *JsonContext) {}

// EnterObj is called when production obj is entered.
func (s *BasetswiftListener) EnterObj(ctx *ObjContext) {}

// ExitObj is called when production obj is exited.
func (s *BasetswiftListener) ExitObj(ctx *ObjContext) {}

// EnterPair is called when production pair is entered.
func (s *BasetswiftListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BasetswiftListener) ExitPair(ctx *PairContext) {}

// EnterArr is called when production arr is entered.
func (s *BasetswiftListener) EnterArr(ctx *ArrContext) {}

// ExitArr is called when production arr is exited.
func (s *BasetswiftListener) ExitArr(ctx *ArrContext) {}

// EnterValue is called when production value is entered.
func (s *BasetswiftListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BasetswiftListener) ExitValue(ctx *ValueContext) {}
