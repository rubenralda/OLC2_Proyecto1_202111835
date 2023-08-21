// Code generated from .\parsing\tswift.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parsing // tswift
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by tswiftParser.
type tswiftVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by tswiftParser#json.
	VisitJson(ctx *JsonContext) interface{}

	// Visit a parse tree produced by tswiftParser#obj.
	VisitObj(ctx *ObjContext) interface{}

	// Visit a parse tree produced by tswiftParser#pair.
	VisitPair(ctx *PairContext) interface{}

	// Visit a parse tree produced by tswiftParser#arr.
	VisitArr(ctx *ArrContext) interface{}

	// Visit a parse tree produced by tswiftParser#value.
	VisitValue(ctx *ValueContext) interface{}
}
