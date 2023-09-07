// Code generated from .\parser\T_swift.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // T_swift
import "github.com/antlr4-go/antlr/v4"

type BaseT_swiftVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseT_swiftVisitor) VisitInicio(ctx *InicioContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitInstrucciones_globales(ctx *Instrucciones_globalesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitIntruccion_global(ctx *Intruccion_globalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitFunction_declaracion(ctx *Function_declaracionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitLista_parametros(ctx *Lista_parametrosContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitDeclaracion_parametro(ctx *Declaracion_parametroContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitCode_block(ctx *Code_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitStruct_declaracion(ctx *Struct_declaracionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitDeclarar_atributo(ctx *Declarar_atributoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitDeclarar_funcion_sc(ctx *Declarar_funcion_scContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitInstrucciones(ctx *InstruccionesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitInstruccion(ctx *InstruccionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitDeclaracion(ctx *DeclaracionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitLoop_statement(ctx *Loop_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitDeclarar_if(ctx *Declarar_ifContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitDeclarar_guard(ctx *Declarar_guardContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitDeclarar_switch(ctx *Declarar_switchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitControl_transfer_statement(ctx *Control_transfer_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitLlamadas_funciones(ctx *Llamadas_funcionesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitLlamada_normal(ctx *Llamada_normalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitLista_argumentos(ctx *Lista_argumentosContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitDeclaracion_argumento(ctx *Declaracion_argumentoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitLlamada_metodos(ctx *Llamada_metodosContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitAtributos_vector_empty(ctx *Atributos_vector_emptyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitAtributos_vector_count(ctx *Atributos_vector_countContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitAtributos_generales(ctx *Atributos_generalesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitAsignar_atributos_normal(ctx *Asignar_atributos_normalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitIncre_atributos_normal(ctx *Incre_atributos_normalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitDecre_atributos_normal(ctx *Decre_atributos_normalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitFor_in_statement(ctx *For_in_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitRango(ctx *RangoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitWhile_statement(ctx *While_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitIf_simple(ctx *If_simpleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitElse_final(ctx *Else_finalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitSiguiente_if(ctx *Siguiente_ifContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitGuard_statement(ctx *Guard_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitSwitch_statement(ctx *Switch_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitSwitch_case(ctx *Switch_caseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitDefault_case(ctx *Default_caseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitBreak_statement(ctx *Break_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitContinue_statement(ctx *Continue_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitReturn_statement(ctx *Return_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitConstant_declaracion(ctx *Constant_declaracionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitVariable_declaracion(ctx *Variable_declaracionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitAnotacion_tipo(ctx *Anotacion_tipoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitTipos(ctx *TiposContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitArray_declaracion(ctx *Array_declaracionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitDefinicion_vector(ctx *Definicion_vectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitLista_expresiones(ctx *Lista_expresionesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitFuncion_print(ctx *Funcion_printContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitFuncion_append(ctx *Funcion_appendContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitFuncion_removeLast(ctx *Funcion_removeLastContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitFuncion_removeat(ctx *Funcion_removeatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitFuncion_int(ctx *Funcion_intContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitFuncion_float(ctx *Funcion_floatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitFuncion_string(ctx *Funcion_stringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitAsignacion_normal(ctx *Asignacion_normalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitAsignacion_incremento(ctx *Asignacion_incrementoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitAsignacion_decremento(ctx *Asignacion_decrementoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitAsignacion_vector(ctx *Asignacion_vectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitAsignacion_incremento_vector(ctx *Asignacion_incremento_vectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitAsignacion_decremento_vector(ctx *Asignacion_decremento_vectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitExpresion_llamada(ctx *Expresion_llamadaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitExpresion_id(ctx *Expresion_idContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitValor_primitivo(ctx *Valor_primitivoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitExpresion_atributos(ctx *Expresion_atributosContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitExpresion_unario(ctx *Expresion_unarioContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitExpresion_paren(ctx *Expresion_parenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitExpresion_rela(ctx *Expresion_relaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitExpresion_arit(ctx *Expresion_aritContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitExpresion_compa(ctx *Expresion_compaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitExpresion_struct_dupla(ctx *Expresion_struct_duplaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitExpresion_vector(ctx *Expresion_vectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitExpresion_nega(ctx *Expresion_negaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitL_duble(ctx *L_dubleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitPrimitivo_caracter(ctx *Primitivo_caracterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitPrimitivo_int(ctx *Primitivo_intContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitPrimitivo_float(ctx *Primitivo_floatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitPrimitivo_string(ctx *Primitivo_stringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitPrimitivo_bool(ctx *Primitivo_boolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseT_swiftVisitor) VisitPrimitivo_nulo(ctx *Primitivo_nuloContext) interface{} {
	return v.VisitChildren(ctx)
}
