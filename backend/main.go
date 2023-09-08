package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"main/ambito"
	"main/arbol"
	"main/parser"
	"os"
	"os/exec"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type parser_visitor struct {
	parser.BaseT_swiftVisitor
	//raiz []arbol.BaseNodo //slice con todos los objetos camuflados por basenodo
}

func (v *parser_visitor) Visit(raiz antlr.ParseTree) interface{} {
	return raiz.Accept(v).([]arbol.BaseNodo)
}

func (v *parser_visitor) VisitInicio(ctx *parser.InicioContext) interface{} {
	return ctx.Instrucciones_globales().Accept(v)
}

func (v *parser_visitor) VisitInstrucciones_globales(ctx *parser.Instrucciones_globalesContext) interface{} {
	var instrucciones []arbol.BaseNodo
	for _, instruccion := range ctx.AllIntruccion_global() {
		fdf := instruccion.Accept(v).(arbol.BaseNodo)
		instrucciones = append(instrucciones, fdf)
	}
	return instrucciones
}

func (v *parser_visitor) VisitIntruccion_global(ctx *parser.Intruccion_globalContext) interface{} {
	if ctx.Declaracion() != nil {
		return ctx.Declaracion().Accept(v).(arbol.BaseNodo)
	} else if ctx.Loop_statement() != nil {
		return ctx.Loop_statement().Accept(v).(arbol.BaseNodo)
	} else if ctx.Branch_statement() != nil {
		return ctx.Branch_statement().Accept(v).(arbol.BaseNodo)
	} else if ctx.Llamadas_funciones() != nil {
		return ctx.Llamadas_funciones().Accept(v).(arbol.BaseNodo)
	} else if ctx.Asignacion() != nil {
		return ctx.Asignacion().Accept(v).(arbol.BaseNodo)
	} else if ctx.Function_declaracion() != nil {
		return ctx.Function_declaracion().Accept(v).(arbol.BaseNodo)
	} else if ctx.Struct_declaracion() != nil {
		return ctx.Struct_declaracion().Accept(v).(arbol.BaseNodo)
	} else if ctx.Asignar_atributos() != nil {
		return ctx.Asignar_atributos().Accept(v).(arbol.BaseNodo)
	}
	return nil
}

// METODOS PARA INSTRUCCIONES DE FUNCIONES

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
	} else if ctx.Llamadas_funciones() != nil {
		return ctx.Llamadas_funciones().Accept(v).(arbol.BaseNodo)
	} else if ctx.Asignacion() != nil {
		return ctx.Asignacion().Accept(v).(arbol.BaseNodo)
	} else if ctx.Control_transfer_statement() != nil {
		return ctx.Control_transfer_statement().Accept(v).(arbol.BaseNodo)
	} else if ctx.Asignar_atributos() != nil {
		return ctx.Asignar_atributos().Accept(v).(arbol.BaseNodo)
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

// METODOS PARA EXPRESION

func (v *parser_visitor) VisitExpresion_id(ctx *parser.Expresion_idContext) interface{} {
	ide := arbol.Id_expresion{Id: ctx.Identificador().GetText()}
	return arbol.Expresion{Valor1: ide, Operacion: "identificador"}
}

func (v *parser_visitor) VisitExpresion_vector(ctx *parser.Expresion_vectorContext) interface{} {
	ide := arbol.Id_vector{Id: ctx.Identificador().GetText(), Indice: ctx.Expresion().Accept(v).(arbol.BaseNodo)}
	return arbol.Expresion{Valor1: ide, Operacion: "vector"}
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

func (v *parser_visitor) VisitExpresion_atributos(ctx *parser.Expresion_atributosContext) interface{} {
	return arbol.Expresion{Valor1: ctx.Atributos().Accept(v).(arbol.BaseNodo), Operacion: "atributos"}
}

func (v *parser_visitor) VisitExpresion_llamada(ctx *parser.Expresion_llamadaContext) interface{} {
	return arbol.Expresion{Valor1: ctx.Llamadas_funciones().Accept(v).(arbol.BaseNodo), Operacion: "llamada"}
}

func (v *parser_visitor) VisitExpresion_struct_dupla(ctx *parser.Expresion_struct_duplaContext) interface{} {
	dd := arbol.Declarar_objeto{Id: ctx.Identificador().GetText(), Dupla: ctx.L_duble().Accept(v).([]arbol.Dupla_atributos)}
	return arbol.Expresion{Valor1: dd, Operacion: "objeto"}
}

func (v *parser_visitor) VisitL_duble(ctx *parser.L_dubleContext) interface{} {
	var atributos []arbol.Dupla_atributos
	for indice, instruccion := range ctx.AllExpresion() {
		expresion := instruccion.Accept(v).(arbol.BaseNodo)
		atributo := arbol.Dupla_atributos{Id_externo: ctx.Identificador(indice).GetText(),
			Expresion: expresion}
		atributos = append(atributos, atributo)
	}
	return atributos
}

// METODOS PARA VALOR PRIMITIVO

func (v *parser_visitor) VisitValor_primitivo(ctx *parser.Valor_primitivoContext) interface{} {
	aa := arbol.Expresion{Valor1: ctx.Primitivos().Accept(v).(arbol.BaseNodo), Operacion: "primitivo"}
	//fmt.Println(aa)
	return aa
}

func (v *parser_visitor) VisitPrimitivo_int(ctx *parser.Primitivo_intContext) interface{} {
	aa := arbol.Primitivos{Tipo: "Int", Valor: ctx.Int().GetText()}
	//fmt.Println(aa)
	return &aa
}

func (v *parser_visitor) VisitPrimitivo_float(ctx *parser.Primitivo_floatContext) interface{} {
	aa := arbol.Primitivos{Tipo: "Float", Valor: ctx.Float().GetText()}
	//fmt.Println(aa)
	return &aa
}

func (v *parser_visitor) VisitPrimitivo_string(ctx *parser.Primitivo_stringContext) interface{} {
	aa := arbol.Primitivos{Tipo: "String", Valor: ctx.String_().GetText()}
	//fmt.Println(aa)
	return &aa
}

func (v *parser_visitor) VisitPrimitivo_bool(ctx *parser.Primitivo_boolContext) interface{} {
	aa := arbol.Primitivos{Tipo: "Bool", Valor: ctx.Bool().GetText()}
	//fmt.Println(aa)
	return &aa
}

func (v *parser_visitor) VisitPrimitivo_caracter(ctx *parser.Primitivo_caracterContext) interface{} {
	aa := arbol.Primitivos{Tipo: "Caracter", Valor: ctx.Caracter().GetText()}
	//fmt.Println(aa)
	return &aa
}

func (v *parser_visitor) VisitPrimitivo_nulo(ctx *parser.Primitivo_nuloContext) interface{} {
	return arbol.Primitivos{Tipo: "nulo", Valor: ctx.GetText()}
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

func (v *parser_visitor) VisitAsignacion_vector(ctx *parser.Asignacion_vectorContext) interface{} {
	return arbol.Asignacion_vector{Id: ctx.Identificador().GetText(),
		Expresion: ctx.Expresion(1).Accept(v).(arbol.BaseNodo),
		Indice:    ctx.Expresion(0).Accept(v).(arbol.BaseNodo)}
}

func (v *parser_visitor) VisitAsignacion_incremento_vector(ctx *parser.Asignacion_incremento_vectorContext) interface{} {
	return arbol.Incremento_vector{Id: ctx.Identificador().GetText(),
		Expresion: ctx.Expresion(1).Accept(v).(arbol.BaseNodo),
		Indice:    ctx.Expresion(0).Accept(v).(arbol.BaseNodo)}
}

func (v *parser_visitor) VisitAsignacion_decremento_vector(ctx *parser.Asignacion_decremento_vectorContext) interface{} {
	return arbol.Decremento_vector{Id: ctx.Identificador().GetText(),
		Expresion: ctx.Expresion(1).Accept(v).(arbol.BaseNodo),
		Indice:    ctx.Expresion(0).Accept(v).(arbol.BaseNodo)}
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
	var rango_inicio arbol.BaseNodo
	var rango_final arbol.BaseNodo
	if ctx.Rango() != nil {
		rango_inicio = ctx.Rango().Expresion(0).Accept(v).(arbol.BaseNodo)
		rango_final = ctx.Rango().Expresion(1).Accept(v).(arbol.BaseNodo)
	}
	return arbol.Loop_for_in{Expresion: expresion, Inicio: rango_inicio, Final: rango_final,
		Sentencias: ctx.Code_block().Accept(v).([]arbol.BaseNodo), Id: ctx.Identificador().GetText()}
}

// METODOS PARA CONTROL DE TRANSFERENCIA
func (v *parser_visitor) VisitControl_transfer_statement(ctx *parser.Control_transfer_statementContext) interface{} {
	var sentencia arbol.Control_transfer
	if ctx.Break_statement() != nil {
		sentencia = arbol.Control_transfer{Sentencia_break: true, Sentencia_continue: false, Sentencia_return: false}
	} else if ctx.Continue_statement() != nil {
		sentencia = arbol.Control_transfer{Sentencia_break: false, Sentencia_continue: true, Sentencia_return: false}
	} else {
		sentencia = arbol.Control_transfer{Sentencia_break: false, Sentencia_continue: false, Sentencia_return: true,
			Retorno: ctx.Return_statement().Accept(v).(arbol.BaseNodo)}
	}
	return sentencia
}

func (v *parser_visitor) VisitReturn_statement(ctx *parser.Return_statementContext) interface{} {
	var expresion arbol.BaseNodo
	tipo := false
	if ctx.Expresion() != nil {
		expresion = ctx.Expresion().Accept(v).(arbol.BaseNodo)
		tipo = true
	}
	return arbol.Sentencia_return{Expresion: expresion, Tipo_retorno: tipo}
}

// METODOS DE SENTENCIA GUARD

func (v *parser_visitor) VisitDeclarar_guard(ctx *parser.Declarar_guardContext) interface{} {
	return ctx.Guard_statement().Accept(v).(arbol.BaseNodo)
}

func (v *parser_visitor) VisitGuard_statement(ctx *parser.Guard_statementContext) interface{} {
	return arbol.Sentencia_guard{Expresion: ctx.Expresion().Accept(v).(arbol.BaseNodo),
		Sentencias: ctx.Instrucciones().Accept(v).([]arbol.BaseNodo)}
}

// METODOS DECLARACION DE VECTOR

func (v *parser_visitor) VisitArray_comun(ctx *parser.Array_comunContext) interface{} {
	var lista_valores []arbol.BaseNodo
	otro_id := ""
	if ctx.Definicion_vector().Lista_expresiones() != nil {
		lista_valores = ctx.Definicion_vector().Lista_expresiones().Accept(v).([]arbol.BaseNodo)
	}
	if ctx.Definicion_vector().Identificador() != nil {
		otro_id = ctx.Definicion_vector().Identificador().GetText()
	}
	return arbol.Declarar_vector{Tipo: ctx.Tipos().GetText(), ID: ctx.Identificador().GetText(),
		Lista_expresion: lista_valores, ID_otro_vector: otro_id}
}

func (v *parser_visitor) VisitArray_vacio(ctx *parser.Array_vacioContext) interface{} {
	var lista_valores []arbol.BaseNodo
	otro_id := ""
	return arbol.Declarar_vector{Tipo: ctx.Tipos().GetText(), ID: ctx.Identificador().GetText(),
		Lista_expresion: lista_valores, ID_otro_vector: otro_id}
}

func (v *parser_visitor) VisitLista_expresiones(ctx *parser.Lista_expresionesContext) interface{} {
	var expresiones []arbol.BaseNodo
	for _, item := range ctx.AllExpresion() {
		expresiones = append(expresiones, item.Accept(v).(arbol.BaseNodo))
	}
	return expresiones
}

// METODOS PARA LLAMADAS DE FUNCIONES

func (v *parser_visitor) VisitLlamadas_funciones(ctx *parser.Llamadas_funcionesContext) interface{} {
	if ctx.Funcion_print() != nil {
		return ctx.Funcion_print().Accept(v).(arbol.BaseNodo)
	} else if ctx.Funcion_append() != nil {
		return ctx.Funcion_append().Accept(v).(arbol.BaseNodo)
	} else if ctx.Funcion_removeLast() != nil {
		return ctx.Funcion_removeLast().Accept(v).(arbol.BaseNodo)
	} else if ctx.Funcion_removeat() != nil {
		return ctx.Funcion_removeat().Accept(v).(arbol.BaseNodo)
	} else if ctx.Funcion_float() != nil {
		return ctx.Funcion_float().Accept(v).(arbol.BaseNodo)
	} else if ctx.Funcion_int() != nil {
		return ctx.Funcion_int().Accept(v).(arbol.BaseNodo)
	} else if ctx.Funcion_string() != nil {
		return ctx.Funcion_string().Accept(v).(arbol.BaseNodo)
	} else if ctx.Llamada_normal() != nil {
		return ctx.Llamada_normal().Accept(v).(arbol.BaseNodo)
	} else if ctx.Llamada_metodos() != nil {
		return ctx.Llamada_metodos().Accept(v).(arbol.BaseNodo)
	}
	return nil
}

// METODOS DE LOS METODOS Y ATRIBUTOS DE UN VECTOR

func (v *parser_visitor) VisitFuncion_append(ctx *parser.Funcion_appendContext) interface{} {
	return arbol.Funcion_append{Id: ctx.Identificador().GetText(),
		Expresion: ctx.Expresion().Accept(v).(arbol.BaseNodo)}
}
func (v *parser_visitor) VisitFuncion_removeLast(ctx *parser.Funcion_removeLastContext) interface{} {
	return arbol.Funcion_removeLast{Id: ctx.Identificador().GetText()}
}
func (v *parser_visitor) VisitFuncion_removeat(ctx *parser.Funcion_removeatContext) interface{} {
	return arbol.Funcion_removeat{Id: ctx.Identificador().GetText(),
		Expresion: ctx.Expresion().Accept(v).(arbol.BaseNodo)}
}

func (v *parser_visitor) VisitAtributos_vector_count(ctx *parser.Atributos_vector_countContext) interface{} {
	return arbol.Vector_count{Id: ctx.Identificador().GetText()}
}
func (v *parser_visitor) VisitAtributos_vector_empty(ctx *parser.Atributos_vector_emptyContext) interface{} {
	return arbol.Vector_isempty{Id: ctx.Identificador().GetText()}
}

// METODOS PARA FUNCIONES EMBEBIDAS

func (v *parser_visitor) VisitFuncion_print(ctx *parser.Funcion_printContext) interface{} {
	lista_valores := ctx.Lista_expresiones().Accept(v).([]arbol.BaseNodo)
	return arbol.Funcion_print{Lista_expresion: lista_valores}
}

func (v *parser_visitor) VisitFuncion_float(ctx *parser.Funcion_floatContext) interface{} {
	return arbol.Funcion_float{Expresion: ctx.Expresion().Accept(v).(arbol.BaseNodo)}
}

func (v *parser_visitor) VisitFuncion_int(ctx *parser.Funcion_intContext) interface{} {
	return arbol.Funcion_int{Expresion: ctx.Expresion().Accept(v).(arbol.BaseNodo)}
}

func (v *parser_visitor) VisitFuncion_string(ctx *parser.Funcion_stringContext) interface{} {
	return arbol.Funcion_string{Expresion: ctx.Expresion().Accept(v).(arbol.BaseNodo)}
}

// METODOS PARA DECLARAR STRUCTS

func (v *parser_visitor) VisitStruct_declaracion(ctx *parser.Struct_declaracionContext) interface{} {
	var atributos []arbol.BaseNodo
	for _, item := range ctx.AllLista_atributos() {
		atributos = append(atributos, item.Accept(v).(arbol.BaseNodo))
	}
	return arbol.Declarar_struct{Id: ctx.Identificador().GetText(), Atributos: atributos}
}

func (v *parser_visitor) VisitDeclarar_atributo(ctx *parser.Declarar_atributoContext) interface{} {
	tipo := ""
	if ctx.GetTipo().GetText() == "let" {
		tipo = "constante"
	} else if ctx.GetTipo().GetText() == "var" {
		tipo = "variable"
	}
	primitivo := ""
	if ctx.Tipos() != nil {
		primitivo = ctx.Tipos().GetText()
	}
	var expresion arbol.BaseNodo
	if ctx.Expresion() != nil {
		expresion = ctx.Expresion().Accept(v).(arbol.BaseNodo)
	}
	return arbol.Declaracion_atributo{Id: ctx.Identificador().GetText(), Primitivo: primitivo,
		Expresion: expresion, Tipo: tipo}
}

func (v *parser_visitor) VisitDeclarar_funcion_sc(ctx *parser.Declarar_funcion_scContext) interface{} {
	funcion := ctx.Function_declaracion().Accept(v).(arbol.Declarar_funcion)
	if ctx.GetM() != nil {
		funcion.Mutable = true
	}
	return funcion
}

// METODO PARA ATRIBUTOS GENERALES

func (v *parser_visitor) VisitAtributos_generales(ctx *parser.Atributos_generalesContext) interface{} {
	id_objeto := ctx.Ide_atributo(0).Accept(v).(arbol.Atributo_ide)
	var atributos []arbol.Atributo_ide
	for _, atributo := range ctx.AllIde_atributo()[1:] {
		atributos = append(atributos, atributo.Accept(v).(arbol.Atributo_ide))
	}
	return arbol.Atributo_general{ID_inicial: id_objeto,
		Lista_atributos: atributos}
}

func (v *parser_visitor) VisitIde_atributo_simple(ctx *parser.Ide_atributo_simpleContext) interface{} {
	return arbol.Atributo_ide{ID: ctx.Identificador().GetText(), Vector: false}
}
func (v *parser_visitor) VisitIde_atributo_vector(ctx *parser.Ide_atributo_vectorContext) interface{} {
	return arbol.Atributo_ide{ID: ctx.Identificador().GetText(), Vector: true, Indice: ctx.Expresion().Accept(v).(arbol.BaseNodo)}
}

func (v *parser_visitor) VisitAsignar_atributos_normal(ctx *parser.Asignar_atributos_normalContext) interface{} {
	id_objeto := ctx.Ide_atributo(0).Accept(v).(arbol.Atributo_ide)
	var atributos []arbol.Atributo_ide
	for _, atributo := range ctx.AllIde_atributo()[1:] {
		atributos = append(atributos, atributo.Accept(v).(arbol.Atributo_ide))
	}
	return arbol.Asignar_atributos{ID_inicial: id_objeto, Lista_atributos: atributos,
		Expresion: ctx.Expresion().Accept(v).(arbol.BaseNodo)}
}

func (v *parser_visitor) VisitIncre_atributos_normal(ctx *parser.Incre_atributos_normalContext) interface{} {
	id_objeto := ctx.Ide_atributo(0).Accept(v).(arbol.Atributo_ide)
	var atributos []arbol.Atributo_ide
	for _, atributo := range ctx.AllIde_atributo()[1:] {
		atributos = append(atributos, atributo.Accept(v).(arbol.Atributo_ide))
	}
	return arbol.Incremento_atributo{ID_inicial: id_objeto, Lista_atributos: atributos,
		Expresion: ctx.Expresion().Accept(v).(arbol.BaseNodo)}
}

func (v *parser_visitor) VisitDecre_atributos_normal(ctx *parser.Decre_atributos_normalContext) interface{} {
	id_objeto := ctx.Ide_atributo(0).Accept(v).(arbol.Atributo_ide)
	var atributos []arbol.Atributo_ide
	for _, atributo := range ctx.AllIde_atributo()[1:] {
		atributos = append(atributos, atributo.Accept(v).(arbol.Atributo_ide))
	}
	return arbol.Decremento_atributo{ID_inicial: id_objeto, Lista_atributos: atributos,
		Expresion: ctx.Expresion().Accept(v).(arbol.BaseNodo)}
}

/*func (v *parser_visitor) VisitAsignar_self(ctx *parser.Asignar_selfContext) interface{} {
	id_objeto := "self"
	var atributos []string
	for _, atributo := range ctx.AllIdentificador() {
		atributos = append(atributos, atributo.GetText())
	}
	return arbol.Asignar_atributos{ID_inicial: id_objeto, Lista_atributos: atributos,
		Expresion: ctx.Expresion().Accept(v).(arbol.BaseNodo)}
}*/

// DECLARACION DE FUNCIONES

func (v *parser_visitor) VisitFunction_declaracion(ctx *parser.Function_declaracionContext) interface{} {
	var lista []arbol.Lista_parametros
	if ctx.Lista_parametros() != nil {
		lista = ctx.Lista_parametros().Accept(v).([]arbol.Lista_parametros)
	}
	tipos := ""
	if ctx.Tipos() != nil {
		tipos = ctx.Tipos().GetText()
	}
	return arbol.Declarar_funcion{Id: ctx.Identificador().GetText(), Lista_parametros: lista,
		Tipo_retorno: tipos, Sentencias: ctx.Code_block().Accept(v).([]arbol.BaseNodo), Mutable: false}
}

func (v *parser_visitor) VisitLista_parametros(ctx *parser.Lista_parametrosContext) interface{} {
	var lista []arbol.Lista_parametros
	for _, item := range ctx.AllDeclaracion_parametro() {
		lista = append(lista, item.Accept(v).(arbol.Lista_parametros))
	}
	return lista
}

func (v *parser_visitor) VisitDeclaracion_parametro(ctx *parser.Declaracion_parametroContext) interface{} {
	externo := ""
	interno := ""
	if len(ctx.AllIdentificador()) > 1 {
		externo = ctx.Identificador(0).GetText()
		interno = ctx.Identificador(1).GetText()
	} else {
		interno = ctx.Identificador(0).GetText()
	}
	referencia := false
	if ctx.GetRefencia() != nil {
		referencia = true
	}
	return arbol.Lista_parametros{Id_interno: interno, Id_externo: externo, Referencia: referencia,
		Primitivo: ctx.Tipos().GetText()}
}

// METODOS LLAMADA FUNCION
func (v *parser_visitor) VisitLlamada_normal(ctx *parser.Llamada_normalContext) interface{} {
	var lista []arbol.Lista_argumentos
	if ctx.Lista_argumentos() != nil {
		lista = ctx.Lista_argumentos().Accept(v).([]arbol.Lista_argumentos)
	}
	var dupla []arbol.Dupla_atributos
	for _, declarar := range lista {
		ejecutar := arbol.Dupla_atributos(declarar)
		dupla = append(dupla, ejecutar)
	}
	declara_objeto := arbol.Declarar_objeto{Id: ctx.Identificador().GetText(), Dupla: dupla}
	return arbol.Llamada_funcion{Id: ctx.Identificador().GetText(), Lista_argumentos: lista,
		Declarar_objeto_amb: declara_objeto}
}

func (v *parser_visitor) VisitLista_argumentos(ctx *parser.Lista_argumentosContext) interface{} {
	var lista []arbol.Lista_argumentos
	for _, item := range ctx.AllDeclaracion_argumento() {
		lista = append(lista, item.Accept(v).(arbol.Lista_argumentos))
	}
	return lista
}

func (v *parser_visitor) VisitDeclaracion_argumento(ctx *parser.Declaracion_argumentoContext) interface{} {
	id_externo := ""
	if ctx.Identificador() != nil {
		id_externo = ctx.Identificador().GetText()
	}
	referencia := false
	if ctx.GetR() != nil {
		referencia = true
	}
	return arbol.Lista_argumentos{Id_externo: id_externo, Referencia: referencia,
		Expresion: ctx.Expresion().Accept(v).(arbol.BaseNodo)}
}

func (v *parser_visitor) VisitLlamada_metodos(ctx *parser.Llamada_metodosContext) interface{} {
	var lista []arbol.Lista_argumentos
	if ctx.Lista_argumentos() != nil {
		lista = ctx.Lista_argumentos().Accept(v).([]arbol.Lista_argumentos)
	}
	return arbol.Llamada_metodo{Id_objeto: ctx.Identificador(0).GetText(),
		Id_metodo:        ctx.Identificador(1).GetText(),
		Lista_argumentos: lista}
}

type Resp struct {
	Salida        string
	Err           bool
	Reporte_error string
	Message       string
}

type Message struct {
	Codigo string `json:"Codigo"`
}

func handleVisitor(c *fiber.Ctx) error {

	/*fichero, err := antlr.NewFileStream("Struct2.swift")
	if err != nil {
		fmt.Println("No se pudo abrir el archivo")
	}*/
	var message Message
	if err := c.BodyParser(&message); err != nil {
		return err
	}
	code := message.Codigo
	input := antlr.NewInputStream(code)
	lexer := parser.NewT_swiftLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewT_swiftParser(tokens)
	p.BuildParseTrees = true
	visitor := parser_visitor{}
	resultado := visitor.Visit(p.Inicio()).([]arbol.BaseNodo)
	ambito_global := &ambito.Ambito{NombreAmbito: "global"}
	for _, linea := range resultado {
		ejecutada := linea.Ejecutar(ambito_global)
		switch ejecutada.(type) {
		case arbol.Control_transfer:
			panic("Error controles de tranferencia no pueden ir aqui")
		case arbol.Sentencia_return:
			panic("Error el return no esta dentro de una funcion")
		case nil: // si es nil no hace nada
		default:
			panic("No deberia pasar")
		}
	}
	for _, local := range ambito_global.Locales {
		fmt.Println(local)
	}
	//fmt.Println(ambito_global.Locales)
	response := Resp{
		Salida:  arbol.Salid_programa,
		Err:     false,
		Message: "<3 Ejecución realizada con éxito <3",
	}
	arbol.Salid_programa = ""
	return c.Status(fiber.StatusOK).JSON(response)
}

func reporte(c *fiber.Ctx) error {
	var message Message
	if err := c.BodyParser(&message); err != nil {
		return err
	}
	code := message.Codigo
	if err := ioutil.WriteFile("./parser/reporte/entrada.swift", []byte(code), 0644); err != nil {
		log.Fatal(err)
	}
	// Comando de generación que deseas ejecutar
	comando := "go generate ./..."

	// Crear un objeto Cmd
	cmd := exec.Command("bash", "-c", comando)

	// Configurar la salida estándar para ver los resultados
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Ejecutar el comando
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	response := Resp{
		Salida:  "",
		Err:     false,
		Message: "Ejecución realizada con éxito",
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Post("/ejecutar", handleVisitor)
	app.Post("/reporte", reporte)
	app.Listen(":3000")
}
