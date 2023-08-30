// Code generated from .\parser\T_swift.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // T_swift
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by T_swiftParser.
type T_swiftVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by T_swiftParser#inicio.
	VisitInicio(ctx *InicioContext) interface{}

	// Visit a parse tree produced by T_swiftParser#instrucciones.
	VisitInstrucciones(ctx *InstruccionesContext) interface{}

	// Visit a parse tree produced by T_swiftParser#instruccion.
	VisitInstruccion(ctx *InstruccionContext) interface{}

	// Visit a parse tree produced by T_swiftParser#loop_statement.
	VisitLoop_statement(ctx *Loop_statementContext) interface{}

	// Visit a parse tree produced by T_swiftParser#code_block.
	VisitCode_block(ctx *Code_blockContext) interface{}

	// Visit a parse tree produced by T_swiftParser#for_in_statement.
	VisitFor_in_statement(ctx *For_in_statementContext) interface{}

	// Visit a parse tree produced by T_swiftParser#while_statement.
	VisitWhile_statement(ctx *While_statementContext) interface{}

	// Visit a parse tree produced by T_swiftParser#branch_statement.
	VisitBranch_statement(ctx *Branch_statementContext) interface{}

	// Visit a parse tree produced by T_swiftParser#if_statement.
	VisitIf_statement(ctx *If_statementContext) interface{}

	// Visit a parse tree produced by T_swiftParser#else_clause.
	VisitElse_clause(ctx *Else_clauseContext) interface{}

	// Visit a parse tree produced by T_swiftParser#guard_statement.
	VisitGuard_statement(ctx *Guard_statementContext) interface{}

	// Visit a parse tree produced by T_swiftParser#switch_statement.
	VisitSwitch_statement(ctx *Switch_statementContext) interface{}

	// Visit a parse tree produced by T_swiftParser#switch_cases.
	VisitSwitch_cases(ctx *Switch_casesContext) interface{}

	// Visit a parse tree produced by T_swiftParser#switch_case.
	VisitSwitch_case(ctx *Switch_caseContext) interface{}

	// Visit a parse tree produced by T_swiftParser#case_label.
	VisitCase_label(ctx *Case_labelContext) interface{}

	// Visit a parse tree produced by T_swiftParser#default_label.
	VisitDefault_label(ctx *Default_labelContext) interface{}

	// Visit a parse tree produced by T_swiftParser#control_transfer_statement.
	VisitControl_transfer_statement(ctx *Control_transfer_statementContext) interface{}

	// Visit a parse tree produced by T_swiftParser#break_statement.
	VisitBreak_statement(ctx *Break_statementContext) interface{}

	// Visit a parse tree produced by T_swiftParser#continue_statement.
	VisitContinue_statement(ctx *Continue_statementContext) interface{}

	// Visit a parse tree produced by T_swiftParser#return_statement.
	VisitReturn_statement(ctx *Return_statementContext) interface{}

	// Visit a parse tree produced by T_swiftParser#declaracion.
	VisitDeclaracion(ctx *DeclaracionContext) interface{}

	// Visit a parse tree produced by T_swiftParser#constant_declaracion.
	VisitConstant_declaracion(ctx *Constant_declaracionContext) interface{}

	// Visit a parse tree produced by T_swiftParser#variable_declaracion.
	VisitVariable_declaracion(ctx *Variable_declaracionContext) interface{}

	// Visit a parse tree produced by T_swiftParser#anotacion_tipo.
	VisitAnotacion_tipo(ctx *Anotacion_tipoContext) interface{}

	// Visit a parse tree produced by T_swiftParser#tipos.
	VisitTipos(ctx *TiposContext) interface{}

	// Visit a parse tree produced by T_swiftParser#array_declaracion.
	VisitArray_declaracion(ctx *Array_declaracionContext) interface{}

	// Visit a parse tree produced by T_swiftParser#definicion_vector.
	VisitDefinicion_vector(ctx *Definicion_vectorContext) interface{}

	// Visit a parse tree produced by T_swiftParser#lista_expresiones.
	VisitLista_expresiones(ctx *Lista_expresionesContext) interface{}

	// Visit a parse tree produced by T_swiftParser#funcion_print.
	VisitFuncion_print(ctx *Funcion_printContext) interface{}

	// Visit a parse tree produced by T_swiftParser#asignacion_normal.
	VisitAsignacion_normal(ctx *Asignacion_normalContext) interface{}

	// Visit a parse tree produced by T_swiftParser#asignacion_incremento.
	VisitAsignacion_incremento(ctx *Asignacion_incrementoContext) interface{}

	// Visit a parse tree produced by T_swiftParser#asignacion_decremento.
	VisitAsignacion_decremento(ctx *Asignacion_decrementoContext) interface{}

	// Visit a parse tree produced by T_swiftParser#expresion_id.
	VisitExpresion_id(ctx *Expresion_idContext) interface{}

	// Visit a parse tree produced by T_swiftParser#valor_primitivo.
	VisitValor_primitivo(ctx *Valor_primitivoContext) interface{}

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

	// Visit a parse tree produced by T_swiftParser#expresion_nega.
	VisitExpresion_nega(ctx *Expresion_negaContext) interface{}

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
}
