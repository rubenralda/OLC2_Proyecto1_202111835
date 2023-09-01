package main

import (
	"fmt"
	"main/ambito"
	"main/arbol"
	"main/parser"

	"github.com/antlr4-go/antlr/v4"
)

type parser_visitor struct {
	parser.BaseT_swiftVisitor
	//raiz []arbol.BaseNodo //slice con todos los objetos camuflados por basenodo
}

func (v *parser_visitor) Visit(raiz antlr.ParseTree) interface{} {
	return raiz.Accept(v).([]arbol.BaseNodo)
}

func (v *parser_visitor) VisitInicio(ctx *parser.InicioContext) interface{} {
	return ctx.Instrucciones().Accept(v)
}

func (v *parser_visitor) VisitInstrucciones(ctx *parser.InstruccionesContext) interface{} {
	var instrucciones []arbol.BaseNodo
	for _, instruccion := range ctx.AllInstruccion() {
		fdf := instruccion.Accept(v).(arbol.BaseNodo)
		instrucciones = append(instrucciones, fdf)
	}
	return instrucciones
}

func (v *parser_visitor) VisitInstruccion(ctx *parser.InstruccionContext) interface{} {
	if ctx.Declaracion() != nil {
		return ctx.Declaracion().Accept(v).(arbol.BaseNodo)
	} else if ctx.Loop_statement() != nil {
		return ctx.Loop_statement().Accept(v).(arbol.BaseNodo)
	} else if ctx.Branch_statement() != nil {
		return ctx.Branch_statement().Accept(v).(arbol.BaseNodo)
	} else if ctx.Funcion_print() != nil {
		return ctx.Funcion_print().Accept(v).(arbol.BaseNodo)
	} else if ctx.Asignacion() != nil {
		return ctx.Asignacion().Accept(v).(arbol.BaseNodo)
	}
	return nil
}

func (v *parser_visitor) VisitDeclaracion(ctx *parser.DeclaracionContext) interface{} {
	if ctx.Variable_declaracion() != nil {
		return ctx.Variable_declaracion().Accept(v).(arbol.BaseNodo)
	} else if ctx.Constant_declaracion() != nil {
		return ctx.Constant_declaracion().Accept(v).(arbol.BaseNodo)
	} else if ctx.Array_declaracion() != nil {
		return ctx.Array_declaracion().Accept(v).(arbol.BaseNodo)
	}
	return nil
}

func (v *parser_visitor) VisitExpresion_id(ctx *parser.Expresion_idContext) interface{} {
	ide := arbol.Id_expresion{Id: ctx.Identificador().GetText()}
	return arbol.Expresion{Valor1: ide, Operacion: "identificador"}
}

func (v *parser_visitor) VisitExpresion_arit(ctx *parser.Expresion_aritContext) interface{} {
	aa := arbol.Expresion{Valor1: ctx.Expresion(0).Accept(v).(arbol.BaseNodo),
		Valor2: ctx.Expresion(1).Accept(v).(arbol.BaseNodo), Operacion: ctx.GetOp().GetText()}
	return aa
}

func (v *parser_visitor) VisitExpresion_unario(ctx *parser.Expresion_unarioContext) interface{} {
	aa := arbol.Expresion{Valor1: ctx.Expresion().Accept(v).(arbol.BaseNodo), Operacion: "negacion"}
	return aa
}

func (v *parser_visitor) VisitExpresion_paren(ctx *parser.Expresion_parenContext) interface{} {
	return ctx.Expresion().Accept(v).(arbol.BaseNodo)
}

func (v *parser_visitor) VisitExpresion_compa(ctx *parser.Expresion_compaContext) interface{} {
	aa := arbol.Expresion{Valor1: ctx.Expresion(0).Accept(v).(arbol.BaseNodo),
		Valor2: ctx.Expresion(1).Accept(v).(arbol.BaseNodo), Operacion: ctx.GetOp().GetText()}
	return aa
}

func (v *parser_visitor) VisitExpresion_nega(ctx *parser.Expresion_negaContext) interface{} {
	aa := arbol.Expresion{Valor1: ctx.Expresion().Accept(v).(arbol.BaseNodo), Operacion: "!"}
	return aa
}

func (v *parser_visitor) VisitExpresion_rela(ctx *parser.Expresion_relaContext) interface{} {
	aa := arbol.Expresion{Valor1: ctx.Expresion(0).Accept(v).(arbol.BaseNodo),
		Valor2: ctx.Expresion(1).Accept(v).(arbol.BaseNodo), Operacion: ctx.GetOp().GetText()}
	return aa
}

func (v *parser_visitor) VisitValor_primitivo(ctx *parser.Valor_primitivoContext) interface{} {
	aa := arbol.Expresion{Valor1: ctx.Primitivos().Accept(v).(arbol.BaseNodo), Operacion: "primitivo"}
	//fmt.Println(aa)
	return aa
}

func (v *parser_visitor) VisitFuncion_print(ctx *parser.Funcion_printContext) interface{} {
	aa := arbol.Imprimir{Valor: ctx.Expresion().Accept(v).(arbol.BaseNodo)}
	//fmt.Println(aa)
	return &aa
}

func (v *parser_visitor) VisitPrimitivo_int(ctx *parser.Primitivo_intContext) interface{} {
	aa := arbol.Primitivos{Tipo: "int", Valor: ctx.Int().GetText()}
	//fmt.Println(aa)
	return &aa
}

func (v *parser_visitor) VisitPrimitivo_float(ctx *parser.Primitivo_floatContext) interface{} {
	aa := arbol.Primitivos{Tipo: "float", Valor: ctx.Float().GetText()}
	//fmt.Println(aa)
	return &aa
}

func (v *parser_visitor) VisitPrimitivo_string(ctx *parser.Primitivo_stringContext) interface{} {
	aa := arbol.Primitivos{Tipo: "string", Valor: ctx.String_().GetText()}
	//fmt.Println(aa)
	return &aa
}

func (v *parser_visitor) VisitPrimitivo_bool(ctx *parser.Primitivo_boolContext) interface{} {
	aa := arbol.Primitivos{Tipo: "bool", Valor: ctx.Bool().GetText()}
	//fmt.Println(aa)
	return &aa
}

func (v *parser_visitor) VisitPrimitivo_caracter(ctx *parser.Primitivo_caracterContext) interface{} {
	aa := arbol.Primitivos{Tipo: "caracter", Valor: ctx.Caracter().GetText()}
	//fmt.Println(aa)
	return &aa
}

// metodos de la gramatica declaracion

func (v *parser_visitor) VisitVariable_declaracion(ctx *parser.Variable_declaracionContext) interface{} {
	tipo := ""
	var expresion arbol.BaseNodo
	if ctx.Anotacion_tipo() != nil {
		tipo = ctx.Anotacion_tipo().Accept(v).(string)
	}
	if ctx.Expresion() != nil {
		expresion = ctx.Expresion().Accept(v).(arbol.BaseNodo)
	}
	return arbol.Declarar_variable{Tipo: tipo, Expresion: expresion, Id: ctx.Identificador().GetText()}
}

// metodo de declaracion de constantes

func (v *parser_visitor) VisitConstant_declaracion(ctx *parser.Constant_declaracionContext) interface{} {
	tipo := ""
	if ctx.Anotacion_tipo() != nil {
		tipo = ctx.Anotacion_tipo().Accept(v).(string)
	}
	return arbol.Declarar_constante{Tipo: tipo, Expresion: ctx.Expresion().Accept(v).(arbol.BaseNodo), Id: ctx.Identificador().GetText()}
}

func (v *parser_visitor) VisitAnotacion_tipo(ctx *parser.Anotacion_tipoContext) interface{} {
	return ctx.Tipos().GetText()
}

// metodos de la gramatica ASIGNACION

func (v *parser_visitor) VisitAsignacion_decremento(ctx *parser.Asignacion_decrementoContext) interface{} {
	return arbol.Decremento_variable{Id: ctx.Identificador().GetText(), Expresion: ctx.Expresion().Accept(v).(arbol.BaseNodo)}
}

func (v *parser_visitor) VisitAsignacion_incremento(ctx *parser.Asignacion_incrementoContext) interface{} {
	return arbol.Incremento_variable{Id: ctx.Identificador().GetText(), Expresion: ctx.Expresion().Accept(v).(arbol.BaseNodo)}
}

func (v *parser_visitor) VisitAsignacion_normal(ctx *parser.Asignacion_normalContext) interface{} {
	return arbol.Asignacion_variable{Id: ctx.Identificador().GetText(), Expresion: ctx.Expresion().Accept(v).(arbol.BaseNodo)}
}

// metodos para la sentencia if else

func (v *parser_visitor) VisitDeclarar_if(ctx *parser.Declarar_ifContext) interface{} {
	return ctx.If_statement().Accept(v).(arbol.BaseNodo)
}

func (v *parser_visitor) VisitCode_block(ctx *parser.Code_blockContext) interface{} {
	return ctx.Instrucciones().Accept(v).([]arbol.BaseNodo)
}

func (v *parser_visitor) VisitIf_simple(ctx *parser.If_simpleContext) interface{} {
	return arbol.Sentencia_if{Expresion: ctx.Expresion().Accept(v).(arbol.BaseNodo),
		Sentencias: ctx.Code_block().Accept(v).([]arbol.BaseNodo)}
}

func (v *parser_visitor) VisitElse_final(ctx *parser.Else_finalContext) interface{} {
	return arbol.Sentencia_if{Expresion: ctx.Expresion().Accept(v).(arbol.BaseNodo),
		Sentencias:      ctx.Code_block(0).Accept(v).([]arbol.BaseNodo),
		Sentencias_else: ctx.Code_block(1).Accept(v).([]arbol.BaseNodo)}
}

func (v *parser_visitor) VisitSiguiente_if(ctx *parser.Siguiente_ifContext) interface{} {
	siguiente := ctx.If_statement().Accept(v).(arbol.BaseNodo)
	return arbol.Sentencia_if{Expresion: ctx.Expresion().Accept(v).(arbol.BaseNodo),
		Sentencias: ctx.Code_block().Accept(v).([]arbol.BaseNodo),
		Siguiente:  siguiente}
}

// METODOS PARA EL SWITCH

func (v *parser_visitor) VisitDeclarar_switch(ctx *parser.Declarar_switchContext) interface{} {
	return ctx.Switch_statement().Accept(v).(arbol.BaseNodo)
}

func (v *parser_visitor) VisitSwitch_statement(ctx *parser.Switch_statementContext) interface{} {
	var lista_case []arbol.Sentencia_case
	for _, item_case := range ctx.AllSwitch_case() {
		fdf := item_case.Accept(v).(arbol.Sentencia_case)
		lista_case = append(lista_case, fdf)
	}
	var default_case arbol.BaseNodo
	if ctx.Default_case() != nil {
		default_case = ctx.Default_case().Accept(v).(arbol.BaseNodo)
	}
	return arbol.Sentencia_switch{Expresion: ctx.Expresion().Accept(v).(arbol.BaseNodo),
		Lista_case: lista_case, Default_case: default_case}
}

func (v *parser_visitor) VisitSwitch_case(ctx *parser.Switch_caseContext) interface{} {
	return arbol.Sentencia_case{Expresion: ctx.Expresion().Accept(v).(arbol.BaseNodo),
		Sentencias: ctx.Instrucciones().Accept(v).([]arbol.BaseNodo)}
}

func (v *parser_visitor) VisitDefault_case(ctx *parser.Default_caseContext) interface{} {
	return arbol.Default_case{Sentencias: ctx.Instrucciones().Accept(v).([]arbol.BaseNodo)}
}

// METODOS PARA EL LOOP

func (v *parser_visitor) VisitLoop_statement(ctx *parser.Loop_statementContext) interface{} {
	if ctx.For_in_statement() != nil {
		return ctx.For_in_statement().Accept(v).(arbol.BaseNodo)
	} else if ctx.While_statement() != nil {
		return ctx.While_statement().Accept(v).(arbol.BaseNodo)
	}
	return nil
}

// METODO PARA EL WHILE

func (v *parser_visitor) VisitWhile_statement(ctx *parser.While_statementContext) interface{} {
	return arbol.Loop_while{Expresion: ctx.Expresion().Accept(v).(arbol.BaseNodo),
		Sentencias: ctx.Code_block().Accept(v).([]arbol.BaseNodo)}
}

// METODOS PARA EL FOR IN

func (v *parser_visitor) VisitFor_in_statement(ctx *parser.For_in_statementContext) interface{} {
	var expresion arbol.BaseNodo
	if ctx.Expresion() != nil {
		expresion = ctx.Expresion().Accept(v).(arbol.BaseNodo)
	}
	rango_inicio := ""
	rango_final := ""
	if ctx.Rango() != nil {
		rango_inicio = ctx.Rango().Int(0).GetText()
		rango_final = ctx.Rango().Int(1).GetText()
	}
	return arbol.Loop_for_in{Expresion: expresion, Inicio: rango_inicio, Final: rango_final,
		Sentencias: ctx.Code_block().Accept(v).([]arbol.BaseNodo), Id: ctx.Identificador().GetText()}
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
	ambito_global := &ambito.Ambito{NombreAmbito: "global"}
	for _, linea := range resultado {
		linea.Ejecutar(ambito_global)
	}
	for _, local := range ambito_global.Locales {
		fmt.Println(local)
	}
	fmt.Println(ambito_global.Locales)

}
