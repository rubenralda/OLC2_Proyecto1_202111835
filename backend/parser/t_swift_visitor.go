// Code generated from .\parser\T_swift.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // T_swift
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by T_swiftParser.
type T_swiftVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by T_swiftParser#inicio.
	VisitInicio(ctx *InicioContext) interface{}

	// Visit a parse tree produced by T_swiftParser#instrucciones_globales.
	VisitInstrucciones_globales(ctx *Instrucciones_globalesContext) interface{}

	// Visit a parse tree produced by T_swiftParser#intruccion_global.
	VisitIntruccion_global(ctx *Intruccion_globalContext) interface{}

	// Visit a parse tree produced by T_swiftParser#function_declaracion.
	VisitFunction_declaracion(ctx *Function_declaracionContext) interface{}

	// Visit a parse tree produced by T_swiftParser#lista_parametros.
	VisitLista_parametros(ctx *Lista_parametrosContext) interface{}

	// Visit a parse tree produced by T_swiftParser#declaracion_parametro.
	VisitDeclaracion_parametro(ctx *Declaracion_parametroContext) interface{}

	// Visit a parse tree produced by T_swiftParser#code_block.
	VisitCode_block(ctx *Code_blockContext) interface{}

	// Visit a parse tree produced by T_swiftParser#struct_declaracion.
	VisitStruct_declaracion(ctx *Struct_declaracionContext) interface{}

	// Visit a parse tree produced by T_swiftParser#declarar_atributo.
	VisitDeclarar_atributo(ctx *Declarar_atributoContext) interface{}

	// Visit a parse tree produced by T_swiftParser#declarar_funcion_sc.
	VisitDeclarar_funcion_sc(ctx *Declarar_funcion_scContext) interface{}

	// Visit a parse tree produced by T_swiftParser#instrucciones.
	VisitInstrucciones(ctx *InstruccionesContext) interface{}

	// Visit a parse tree produced by T_swiftParser#instruccion.
	VisitInstruccion(ctx *InstruccionContext) interface{}

	// Visit a parse tree produced by T_swiftParser#declaracion.
	VisitDeclaracion(ctx *DeclaracionContext) interface{}

	// Visit a parse tree produced by T_swiftParser#loop_statement.
	VisitLoop_statement(ctx *Loop_statementContext) interface{}

	// Visit a parse tree produced by T_swiftParser#declarar_if.
	VisitDeclarar_if(ctx *Declarar_ifContext) interface{}

	// Visit a parse tree produced by T_swiftParser#declarar_guard.
	VisitDeclarar_guard(ctx *Declarar_guardContext) interface{}

	// Visit a parse tree produced by T_swiftParser#declarar_switch.
	VisitDeclarar_switch(ctx *Declarar_switchContext) interface{}

	// Visit a parse tree produced by T_swiftParser#control_transfer_statement.
	VisitControl_transfer_statement(ctx *Control_transfer_statementContext) interface{}

	// Visit a parse tree produced by T_swiftParser#llamadas_funciones.
	VisitLlamadas_funciones(ctx *Llamadas_funcionesContext) interface{}

	// Visit a parse tree produced by T_swiftParser#llamada_normal.
	VisitLlamada_normal(ctx *Llamada_normalContext) interface{}

	// Visit a parse tree produced by T_swiftParser#lista_argumentos.
	VisitLista_argumentos(ctx *Lista_argumentosContext) interface{}

	// Visit a parse tree produced by T_swiftParser#declaracion_argumento.
	VisitDeclaracion_argumento(ctx *Declaracion_argumentoContext) interface{}

	// Visit a parse tree produced by T_swiftParser#llamada_metodos.
	VisitLlamada_metodos(ctx *Llamada_metodosContext) interface{}

	// Visit a parse tree produced by T_swiftParser#atributos_vector_empty.
	VisitAtributos_vector_empty(ctx *Atributos_vector_emptyContext) interface{}

	// Visit a parse tree produced by T_swiftParser#atributos_vector_count.
	VisitAtributos_vector_count(ctx *Atributos_vector_countContext) interface{}

	// Visit a parse tree produced by T_swiftParser#atributos_generales.
	VisitAtributos_generales(ctx *Atributos_generalesContext) interface{}

	// Visit a parse tree produced by T_swiftParser#asignar_atributos_normal.
	VisitAsignar_atributos_normal(ctx *Asignar_atributos_normalContext) interface{}

	// Visit a parse tree produced by T_swiftParser#incre_atributos_normal.
	VisitIncre_atributos_normal(ctx *Incre_atributos_normalContext) interface{}

	// Visit a parse tree produced by T_swiftParser#decre_atributos_normal.
	VisitDecre_atributos_normal(ctx *Decre_atributos_normalContext) interface{}

	// Visit a parse tree produced by T_swiftParser#for_in_statement.
	VisitFor_in_statement(ctx *For_in_statementContext) interface{}

	// Visit a parse tree produced by T_swiftParser#rango.
	VisitRango(ctx *RangoContext) interface{}

	// Visit a parse tree produced by T_swiftParser#while_statement.
	VisitWhile_statement(ctx *While_statementContext) interface{}

	// Visit a parse tree produced by T_swiftParser#if_simple.
	VisitIf_simple(ctx *If_simpleContext) interface{}

	// Visit a parse tree produced by T_swiftParser#else_final.
	VisitElse_final(ctx *Else_finalContext) interface{}

	// Visit a parse tree produced by T_swiftParser#siguiente_if.
	VisitSiguiente_if(ctx *Siguiente_ifContext) interface{}

	// Visit a parse tree produced by T_swiftParser#guard_statement.
	VisitGuard_statement(ctx *Guard_statementContext) interface{}

	// Visit a parse tree produced by T_swiftParser#switch_statement.
	VisitSwitch_statement(ctx *Switch_statementContext) interface{}

	// Visit a parse tree produced by T_swiftParser#switch_case.
	VisitSwitch_case(ctx *Switch_caseContext) interface{}

	// Visit a parse tree produced by T_swiftParser#default_case.
	VisitDefault_case(ctx *Default_caseContext) interface{}

	// Visit a parse tree produced by T_swiftParser#break_statement.
	VisitBreak_statement(ctx *Break_statementContext) interface{}

	// Visit a parse tree produced by T_swiftParser#continue_statement.
	VisitContinue_statement(ctx *Continue_statementContext) interface{}

	// Visit a parse tree produced by T_swiftParser#return_statement.
	VisitReturn_statement(ctx *Return_statementContext) interface{}

	// Visit a parse tree produced by T_swiftParser#constant_declaracion.
	VisitConstant_declaracion(ctx *Constant_declaracionContext) interface{}

	// Visit a parse tree produced by T_swiftParser#variable_declaracion.
	VisitVariable_declaracion(ctx *Variable_declaracionContext) interface{}

	// Visit a parse tree produced by T_swiftParser#anotacion_tipo.
	VisitAnotacion_tipo(ctx *Anotacion_tipoContext) interface{}

	// Visit a parse tree produced by T_swiftParser#tipos.
	VisitTipos(ctx *TiposContext) interface{}

	// Visit a parse tree produced by T_swiftParser#array_comun.
	VisitArray_comun(ctx *Array_comunContext) interface{}

	// Visit a parse tree produced by T_swiftParser#array_vacio.
	VisitArray_vacio(ctx *Array_vacioContext) interface{}

	// Visit a parse tree produced by T_swiftParser#definicion_vector.
	VisitDefinicion_vector(ctx *Definicion_vectorContext) interface{}

	// Visit a parse tree produced by T_swiftParser#lista_expresiones.
	VisitLista_expresiones(ctx *Lista_expresionesContext) interface{}

	// Visit a parse tree produced by T_swiftParser#funcion_print.
	VisitFuncion_print(ctx *Funcion_printContext) interface{}

	// Visit a parse tree produced by T_swiftParser#funcion_append.
	VisitFuncion_append(ctx *Funcion_appendContext) interface{}

	// Visit a parse tree produced by T_swiftParser#funcion_removeLast.
	VisitFuncion_removeLast(ctx *Funcion_removeLastContext) interface{}

	// Visit a parse tree produced by T_swiftParser#funcion_removeat.
	VisitFuncion_removeat(ctx *Funcion_removeatContext) interface{}

	// Visit a parse tree produced by T_swiftParser#funcion_int.
	VisitFuncion_int(ctx *Funcion_intContext) interface{}

	// Visit a parse tree produced by T_swiftParser#funcion_float.
	VisitFuncion_float(ctx *Funcion_floatContext) interface{}

	// Visit a parse tree produced by T_swiftParser#funcion_string.
	VisitFuncion_string(ctx *Funcion_stringContext) interface{}

	// Visit a parse tree produced by T_swiftParser#asignacion_normal.
	VisitAsignacion_normal(ctx *Asignacion_normalContext) interface{}

	// Visit a parse tree produced by T_swiftParser#asignacion_incremento.
	VisitAsignacion_incremento(ctx *Asignacion_incrementoContext) interface{}

	// Visit a parse tree produced by T_swiftParser#asignacion_decremento.
	VisitAsignacion_decremento(ctx *Asignacion_decrementoContext) interface{}

	// Visit a parse tree produced by T_swiftParser#asignacion_vector.
	VisitAsignacion_vector(ctx *Asignacion_vectorContext) interface{}

	// Visit a parse tree produced by T_swiftParser#asignacion_incremento_vector.
	VisitAsignacion_incremento_vector(ctx *Asignacion_incremento_vectorContext) interface{}

	// Visit a parse tree produced by T_swiftParser#asignacion_decremento_vector.
	VisitAsignacion_decremento_vector(ctx *Asignacion_decremento_vectorContext) interface{}

	// Visit a parse tree produced by T_swiftParser#expresion_llamada.
	VisitExpresion_llamada(ctx *Expresion_llamadaContext) interface{}

	// Visit a parse tree produced by T_swiftParser#expresion_id.
	VisitExpresion_id(ctx *Expresion_idContext) interface{}

	// Visit a parse tree produced by T_swiftParser#valor_primitivo.
	VisitValor_primitivo(ctx *Valor_primitivoContext) interface{}

	// Visit a parse tree produced by T_swiftParser#expresion_atributos.
	VisitExpresion_atributos(ctx *Expresion_atributosContext) interface{}

	// Visit a parse tree produced by T_swiftParser#expresion_unario.
	VisitExpresion_unario(ctx *Expresion_unarioContext) interface{}

	// Visit a parse tree produced by T_swiftParser#expresion_paren.
	VisitExpresion_paren(ctx *Expresion_parenContext) interface{}

	// Visit a parse tree produced by T_swiftParser#expresion_rela.
	VisitExpresion_rela(ctx *Expresion_relaContext) interface{}

	// Visit a parse tree produced by T_swiftParser#expresion_arit.
	VisitExpresion_arit(ctx *Expresion_aritContext) interface{}

	// Visit a parse tree produced by T_swiftParser#expresion_compa.
	VisitExpresion_compa(ctx *Expresion_compaContext) interface{}

	// Visit a parse tree produced by T_swiftParser#expresion_struct_dupla.
	VisitExpresion_struct_dupla(ctx *Expresion_struct_duplaContext) interface{}

	// Visit a parse tree produced by T_swiftParser#expresion_vector.
	VisitExpresion_vector(ctx *Expresion_vectorContext) interface{}

	// Visit a parse tree produced by T_swiftParser#expresion_nega.
	VisitExpresion_nega(ctx *Expresion_negaContext) interface{}

	// Visit a parse tree produced by T_swiftParser#l_duble.
	VisitL_duble(ctx *L_dubleContext) interface{}

	// Visit a parse tree produced by T_swiftParser#primitivo_caracter.
	VisitPrimitivo_caracter(ctx *Primitivo_caracterContext) interface{}

	// Visit a parse tree produced by T_swiftParser#primitivo_int.
	VisitPrimitivo_int(ctx *Primitivo_intContext) interface{}

	// Visit a parse tree produced by T_swiftParser#primitivo_float.
	VisitPrimitivo_float(ctx *Primitivo_floatContext) interface{}

	// Visit a parse tree produced by T_swiftParser#primitivo_string.
	VisitPrimitivo_string(ctx *Primitivo_stringContext) interface{}

	// Visit a parse tree produced by T_swiftParser#primitivo_bool.
	VisitPrimitivo_bool(ctx *Primitivo_boolContext) interface{}

	// Visit a parse tree produced by T_swiftParser#primitivo_nulo.
	VisitPrimitivo_nulo(ctx *Primitivo_nuloContext) interface{}
}
