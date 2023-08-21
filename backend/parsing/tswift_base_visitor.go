// Code generated from .\parsing\tswift.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parsing // tswift
import "github.com/antlr4-go/antlr/v4"

type BasetswiftVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasetswiftVisitor) VisitJson(ctx *JsonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetswiftVisitor) VisitObj(ctx *ObjContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetswiftVisitor) VisitPair(ctx *PairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetswiftVisitor) VisitArr(ctx *ArrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetswiftVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}
