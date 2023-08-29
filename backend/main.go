package main

import (
	"fmt"
	"main/arbol"
	"main/parser"

	"github.com/antlr4-go/antlr/v4"
)

type parser_visitor struct {
	parser.BaseT_swiftVisitor
	raiz []arbol.BaseNodo //slice con todos los objetos camuflados por basenodo
}

func (v *parser_visitor) VisitInicio(ctx *parser.InicioContext) interface{} {
	return ctx.Instrucciones().Accept(v)
}

func (v *parser_visitor) VisitInstrucciones(ctx *parser.InstruccionesContext) interface{} {
	for _, instruccion := range ctx.AllInstruccion() {
		fdf := instruccion.Accept(v).(arbol.BaseNodo)
		v.raiz = append(v.raiz, fdf)
	}
	return v.raiz
}

func (v *parser_visitor) VisitInstruccion(ctx *parser.InstruccionContext) interface{} {
	if ctx.Declaracion() != nil {
		return ctx.Declaracion().Accept(v).(arbol.BaseNodo)
	} else if ctx.Loop_statement() != nil {
		return ctx.Loop_statement().Accept(v).(arbol.BaseNodo)
	} else if ctx.Funcion_print() != nil {
		return ctx.Funcion_print().Accept(v).(arbol.BaseNodo)
	}
	return nil
}

func (v *parser_visitor) VisitDeclaracion(ctx *parser.DeclaracionContext) interface{} {
	return ctx.Variable_declaracion().Accept(v).(arbol.BaseNodo)
}

func (v *parser_visitor) VisitVariable_declaracion(ctx *parser.Variable_declaracionContext) interface{} {
	return ctx.Expresion().Accept(v).(arbol.BaseNodo)
}

func (v *parser_visitor) Visit(raiz antlr.ParseTree) interface{} {
	return raiz.Accept(v).([]arbol.BaseNodo)
}

func (v *parser_visitor) VisitExpresion_arit(ctx *parser.Expresion_aritContext) interface{} {
	aa := arbol.Expresion{Valor1: ctx.GetRight().Accept(v).(arbol.BaseNodo),
		Valor2: ctx.GetLeft().Accept(v).(arbol.BaseNodo), Operacion: ctx.GetOp().GetText()}
	return aa
}

func (v *parser_visitor) VisitExpresion_unario(ctx *parser.Expresion_unarioContext) interface{} {
	aa := arbol.Expresion{Valor1: ctx.Expresion().Accept(v).(arbol.BaseNodo), Operacion: "-"}
	return aa
}

func (v *parser_visitor) VisitExpresion_paren(ctx *parser.Expresion_parenContext) interface{} {
	return ctx.Expresion().Accept(v).(arbol.BaseNodo)
}

func (v *parser_visitor) VisitExpresion_compa(ctx *parser.Expresion_compaContext) interface{} {
	aa := arbol.Expresion{Valor1: ctx.GetRight().Accept(v).(arbol.BaseNodo),
		Valor2: ctx.GetLeft().Accept(v).(arbol.BaseNodo), Operacion: ctx.GetOp().GetText()}
	return aa
}

func (v *parser_visitor) VisitExpresion_nega(ctx *parser.Expresion_negaContext) interface{} {
	aa := arbol.Expresion{Valor1: ctx.GetLeft().Accept(v).(arbol.BaseNodo), Operacion: "!"}
	return aa
}

func (v *parser_visitor) VisitExpresion_rela(ctx *parser.Expresion_relaContext) interface{} {
	aa := arbol.Expresion{Valor1: ctx.GetRight().Accept(v).(arbol.BaseNodo),
		Valor2: ctx.GetLeft().Accept(v).(arbol.BaseNodo), Operacion: ctx.GetOp().GetText()}
	return aa
}

func (v *parser_visitor) VisitValor_primitivo(ctx *parser.Valor_primitivoContext) interface{} {
	aa := arbol.Expresion{Valor1: ctx.Primitivos().Accept(v).(arbol.BaseNodo), Operacion: "primitivo"}
	fmt.Println(aa)
	return aa
}

func (v *parser_visitor) VisitFuncion_print(ctx *parser.Funcion_printContext) interface{} {
	aa := arbol.Imprimir{Valor: ctx.Expresion().Accept(v).(arbol.BaseNodo)}
	fmt.Println(aa)
	return &aa
}

func (v *parser_visitor) VisitPrimitivo_int(ctx *parser.Primitivo_intContext) interface{} {
	aa := arbol.Primitivos{Tipo: "int", Valor: ctx.Int().GetText()}
	fmt.Println(aa)
	return &aa
}

func (v *parser_visitor) VisitPrimitivo_float(ctx *parser.Primitivo_floatContext) interface{} {
	aa := arbol.Primitivos{Tipo: "float", Valor: ctx.Float().GetText()}
	fmt.Println(aa)
	return &aa
}

func (v *parser_visitor) VisitPrimitivo_string(ctx *parser.Primitivo_stringContext) interface{} {
	aa := arbol.Primitivos{Tipo: "string", Valor: ctx.String_().GetText()}
	fmt.Println(aa)
	return &aa
}

func (v *parser_visitor) VisitPrimitivo_bool(ctx *parser.Primitivo_boolContext) interface{} {
	aa := arbol.Primitivos{Tipo: "bool", Valor: ctx.Bool().GetText()}
	fmt.Println(aa)
	return &aa
}

func (v *parser_visitor) VisitPrimitivo_caracter(ctx *parser.Primitivo_caracterContext) interface{} {
	aa := arbol.Primitivos{Tipo: "caracter", Valor: ctx.Caracter().GetText()}
	fmt.Println(aa)
	return &aa
}

func main() {
	fichero, err := antlr.NewFileStream("entrada.txt")
	if err != nil {
		fmt.Println("No se pudo abrir el archivo")
	}
	lexer := parser.NewT_swiftLexer(fichero)
	tokens := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewT_swiftParser(tokens)
	p.BuildParseTrees = true
	visitor := parser_visitor{}
	resultado := visitor.Visit(p.Inicio()).([]arbol.BaseNodo)
	for _, linea := range resultado {
		linea.Ejecutar()
	}

}
