// Code generated from .\parser\T_swift.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // T_swift
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type T_swiftParser struct {
	*antlr.BaseParser
}

var T_swiftParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func t_swiftParserInit() {
	staticData := &T_swiftParserStaticData
	staticData.LiteralNames = []string{
		"", "';'", "'func'", "'('", "')'", "'->'", "','", "'_'", "':'", "'inout'",
		"'{'", "'}'", "'.'", "'IsEmpty'", "'count'", "'for'", "'in'", "'...'",
		"'while'", "'if'", "'else'", "'guard'", "'continue'", "'break'", "'return'",
		"'switch'", "'case'", "'default'", "'let'", "'='", "'var'", "'?'", "'String'",
		"'Int'", "'Float'", "'Bool'", "'Character'", "'['", "']'", "'print'",
		"'append'", "'removeLast'", "'remove'", "'at'", "'+='", "'-='", "'!'",
		"'-'", "'/'", "'%'", "'*'", "'+'", "'<'", "'<='", "'>='", "'>'", "'=='",
		"'!='", "'&&'", "'||'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "WS", "Block_comment", "Line_comment",
		"Int", "Float", "Caracter", "String", "Bool", "Identificador",
	}
	staticData.RuleNames = []string{
		"inicio", "instrucciones_globales", "intruccion_global", "function_declaracion",
		"lista_parametros", "code_block", "instrucciones", "instruccion", "declaracion",
		"loop_statement", "branch_statement", "control_transfer_statement",
		"llamadas_funciones", "atributos", "for_in_statement", "rango", "while_statement",
		"if_statement", "guard_statement", "switch_statement", "switch_case",
		"default_case", "break_statement", "continue_statement", "return_statement",
		"constant_declaracion", "variable_declaracion", "anotacion_tipo", "tipos",
		"array_declaracion", "definicion_vector", "lista_expresiones", "funcion_print",
		"funcion_append", "funcion_removeLast", "funcion_removeat", "funcion_int",
		"funcion_float", "funcion_string", "asignacion", "expresion", "primitivos",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 68, 474, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 1,
		0, 1, 0, 1, 0, 1, 1, 5, 1, 89, 8, 1, 10, 1, 12, 1, 92, 9, 1, 1, 2, 1, 2,
		3, 2, 96, 8, 2, 1, 2, 1, 2, 3, 2, 100, 8, 2, 1, 2, 1, 2, 3, 2, 104, 8,
		2, 1, 2, 1, 2, 3, 2, 108, 8, 2, 1, 2, 1, 2, 3, 2, 112, 8, 2, 1, 2, 3, 2,
		115, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 121, 8, 3, 1, 3, 1, 3, 1, 3, 3,
		3, 126, 8, 3, 1, 3, 1, 3, 1, 4, 1, 4, 3, 4, 132, 8, 4, 1, 4, 1, 4, 3, 4,
		136, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 142, 8, 4, 1, 4, 1, 4, 3, 4, 146,
		8, 4, 1, 4, 3, 4, 149, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 5, 6, 156, 8,
		6, 10, 6, 12, 6, 159, 9, 6, 1, 7, 1, 7, 3, 7, 163, 8, 7, 1, 7, 1, 7, 3,
		7, 167, 8, 7, 1, 7, 1, 7, 3, 7, 171, 8, 7, 1, 7, 1, 7, 3, 7, 175, 8, 7,
		1, 7, 1, 7, 3, 7, 179, 8, 7, 1, 7, 1, 7, 3, 7, 183, 8, 7, 3, 7, 185, 8,
		7, 1, 8, 1, 8, 1, 8, 3, 8, 190, 8, 8, 1, 9, 1, 9, 3, 9, 194, 8, 9, 1, 10,
		1, 10, 1, 10, 3, 10, 199, 8, 10, 1, 11, 1, 11, 1, 11, 3, 11, 204, 8, 11,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 213, 8, 12, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 221, 8, 13, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 3, 14, 228, 8, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 3, 17, 256, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 270, 8, 19, 10, 19, 12, 19, 273,
		9, 19, 1, 19, 3, 19, 276, 8, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 3, 20, 285, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 291, 8,
		21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 3, 24, 299, 8, 24, 1, 25,
		1, 25, 1, 25, 3, 25, 304, 8, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 317, 8, 26, 1, 26, 1, 26,
		3, 26, 321, 8, 26, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 346, 8, 30, 1, 31, 1, 31, 1,
		31, 5, 31, 351, 8, 31, 10, 31, 12, 31, 354, 9, 31, 1, 32, 1, 32, 1, 32,
		1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39,
		1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1,
		39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39,
		1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 3, 39, 428, 8, 39, 1, 40, 1,
		40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40,
		1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 3, 40, 448, 8, 40, 1, 40, 1,
		40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40,
		5, 40, 462, 8, 40, 10, 40, 12, 40, 465, 9, 40, 1, 41, 1, 41, 1, 41, 1,
		41, 1, 41, 3, 41, 472, 8, 41, 1, 41, 0, 1, 80, 42, 0, 2, 4, 6, 8, 10, 12,
		14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48,
		50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 0,
		7, 2, 0, 7, 7, 68, 68, 1, 0, 22, 24, 1, 0, 32, 36, 1, 0, 48, 50, 2, 0,
		47, 47, 51, 51, 1, 0, 52, 57, 1, 0, 58, 59, 509, 0, 84, 1, 0, 0, 0, 2,
		90, 1, 0, 0, 0, 4, 114, 1, 0, 0, 0, 6, 116, 1, 0, 0, 0, 8, 148, 1, 0, 0,
		0, 10, 150, 1, 0, 0, 0, 12, 157, 1, 0, 0, 0, 14, 184, 1, 0, 0, 0, 16, 189,
		1, 0, 0, 0, 18, 193, 1, 0, 0, 0, 20, 198, 1, 0, 0, 0, 22, 203, 1, 0, 0,
		0, 24, 212, 1, 0, 0, 0, 26, 220, 1, 0, 0, 0, 28, 222, 1, 0, 0, 0, 30, 231,
		1, 0, 0, 0, 32, 235, 1, 0, 0, 0, 34, 255, 1, 0, 0, 0, 36, 257, 1, 0, 0,
		0, 38, 265, 1, 0, 0, 0, 40, 279, 1, 0, 0, 0, 42, 286, 1, 0, 0, 0, 44, 292,
		1, 0, 0, 0, 46, 294, 1, 0, 0, 0, 48, 296, 1, 0, 0, 0, 50, 300, 1, 0, 0,
		0, 52, 320, 1, 0, 0, 0, 54, 322, 1, 0, 0, 0, 56, 325, 1, 0, 0, 0, 58, 327,
		1, 0, 0, 0, 60, 345, 1, 0, 0, 0, 62, 347, 1, 0, 0, 0, 64, 355, 1, 0, 0,
		0, 66, 360, 1, 0, 0, 0, 68, 367, 1, 0, 0, 0, 70, 373, 1, 0, 0, 0, 72, 382,
		1, 0, 0, 0, 74, 387, 1, 0, 0, 0, 76, 392, 1, 0, 0, 0, 78, 427, 1, 0, 0,
		0, 80, 447, 1, 0, 0, 0, 82, 471, 1, 0, 0, 0, 84, 85, 3, 2, 1, 0, 85, 86,
		5, 0, 0, 1, 86, 1, 1, 0, 0, 0, 87, 89, 3, 4, 2, 0, 88, 87, 1, 0, 0, 0,
		89, 92, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 3, 1, 0,
		0, 0, 92, 90, 1, 0, 0, 0, 93, 95, 3, 16, 8, 0, 94, 96, 5, 1, 0, 0, 95,
		94, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 115, 1, 0, 0, 0, 97, 99, 3, 18,
		9, 0, 98, 100, 5, 1, 0, 0, 99, 98, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100,
		115, 1, 0, 0, 0, 101, 103, 3, 20, 10, 0, 102, 104, 5, 1, 0, 0, 103, 102,
		1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 115, 1, 0, 0, 0, 105, 107, 3, 78,
		39, 0, 106, 108, 5, 1, 0, 0, 107, 106, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0,
		108, 115, 1, 0, 0, 0, 109, 111, 3, 24, 12, 0, 110, 112, 5, 1, 0, 0, 111,
		110, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 115, 1, 0, 0, 0, 113, 115,
		3, 6, 3, 0, 114, 93, 1, 0, 0, 0, 114, 97, 1, 0, 0, 0, 114, 101, 1, 0, 0,
		0, 114, 105, 1, 0, 0, 0, 114, 109, 1, 0, 0, 0, 114, 113, 1, 0, 0, 0, 115,
		5, 1, 0, 0, 0, 116, 117, 5, 2, 0, 0, 117, 118, 5, 68, 0, 0, 118, 120, 5,
		3, 0, 0, 119, 121, 3, 8, 4, 0, 120, 119, 1, 0, 0, 0, 120, 121, 1, 0, 0,
		0, 121, 122, 1, 0, 0, 0, 122, 125, 5, 4, 0, 0, 123, 124, 5, 5, 0, 0, 124,
		126, 3, 56, 28, 0, 125, 123, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 127,
		1, 0, 0, 0, 127, 128, 3, 10, 5, 0, 128, 7, 1, 0, 0, 0, 129, 131, 5, 6,
		0, 0, 130, 132, 7, 0, 0, 0, 131, 130, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0,
		132, 133, 1, 0, 0, 0, 133, 135, 5, 8, 0, 0, 134, 136, 5, 9, 0, 0, 135,
		134, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 138,
		3, 56, 28, 0, 138, 139, 3, 8, 4, 0, 139, 149, 1, 0, 0, 0, 140, 142, 7,
		0, 0, 0, 141, 140, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 143, 1, 0, 0,
		0, 143, 145, 5, 8, 0, 0, 144, 146, 5, 9, 0, 0, 145, 144, 1, 0, 0, 0, 145,
		146, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 149, 3, 56, 28, 0, 148, 129,
		1, 0, 0, 0, 148, 141, 1, 0, 0, 0, 149, 9, 1, 0, 0, 0, 150, 151, 5, 10,
		0, 0, 151, 152, 3, 12, 6, 0, 152, 153, 5, 11, 0, 0, 153, 11, 1, 0, 0, 0,
		154, 156, 3, 14, 7, 0, 155, 154, 1, 0, 0, 0, 156, 159, 1, 0, 0, 0, 157,
		155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 13, 1, 0, 0, 0, 159, 157, 1,
		0, 0, 0, 160, 162, 3, 16, 8, 0, 161, 163, 5, 1, 0, 0, 162, 161, 1, 0, 0,
		0, 162, 163, 1, 0, 0, 0, 163, 185, 1, 0, 0, 0, 164, 166, 3, 18, 9, 0, 165,
		167, 5, 1, 0, 0, 166, 165, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 185,
		1, 0, 0, 0, 168, 170, 3, 20, 10, 0, 169, 171, 5, 1, 0, 0, 170, 169, 1,
		0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 185, 1, 0, 0, 0, 172, 174, 3, 22, 11,
		0, 173, 175, 5, 1, 0, 0, 174, 173, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175,
		185, 1, 0, 0, 0, 176, 178, 3, 78, 39, 0, 177, 179, 5, 1, 0, 0, 178, 177,
		1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 185, 1, 0, 0, 0, 180, 182, 3, 24,
		12, 0, 181, 183, 5, 1, 0, 0, 182, 181, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0,
		183, 185, 1, 0, 0, 0, 184, 160, 1, 0, 0, 0, 184, 164, 1, 0, 0, 0, 184,
		168, 1, 0, 0, 0, 184, 172, 1, 0, 0, 0, 184, 176, 1, 0, 0, 0, 184, 180,
		1, 0, 0, 0, 185, 15, 1, 0, 0, 0, 186, 190, 3, 50, 25, 0, 187, 190, 3, 52,
		26, 0, 188, 190, 3, 58, 29, 0, 189, 186, 1, 0, 0, 0, 189, 187, 1, 0, 0,
		0, 189, 188, 1, 0, 0, 0, 190, 17, 1, 0, 0, 0, 191, 194, 3, 28, 14, 0, 192,
		194, 3, 32, 16, 0, 193, 191, 1, 0, 0, 0, 193, 192, 1, 0, 0, 0, 194, 19,
		1, 0, 0, 0, 195, 199, 3, 34, 17, 0, 196, 199, 3, 36, 18, 0, 197, 199, 3,
		38, 19, 0, 198, 195, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0, 198, 197, 1, 0,
		0, 0, 199, 21, 1, 0, 0, 0, 200, 204, 3, 44, 22, 0, 201, 204, 3, 46, 23,
		0, 202, 204, 3, 48, 24, 0, 203, 200, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0,
		203, 202, 1, 0, 0, 0, 204, 23, 1, 0, 0, 0, 205, 213, 3, 64, 32, 0, 206,
		213, 3, 66, 33, 0, 207, 213, 3, 68, 34, 0, 208, 213, 3, 70, 35, 0, 209,
		213, 3, 72, 36, 0, 210, 213, 3, 74, 37, 0, 211, 213, 3, 76, 38, 0, 212,
		205, 1, 0, 0, 0, 212, 206, 1, 0, 0, 0, 212, 207, 1, 0, 0, 0, 212, 208,
		1, 0, 0, 0, 212, 209, 1, 0, 0, 0, 212, 210, 1, 0, 0, 0, 212, 211, 1, 0,
		0, 0, 213, 25, 1, 0, 0, 0, 214, 215, 5, 68, 0, 0, 215, 216, 5, 12, 0, 0,
		216, 221, 5, 13, 0, 0, 217, 218, 5, 68, 0, 0, 218, 219, 5, 12, 0, 0, 219,
		221, 5, 14, 0, 0, 220, 214, 1, 0, 0, 0, 220, 217, 1, 0, 0, 0, 221, 27,
		1, 0, 0, 0, 222, 223, 5, 15, 0, 0, 223, 224, 5, 68, 0, 0, 224, 227, 5,
		16, 0, 0, 225, 228, 3, 80, 40, 0, 226, 228, 3, 30, 15, 0, 227, 225, 1,
		0, 0, 0, 227, 226, 1, 0, 0, 0, 228, 229, 1, 0, 0, 0, 229, 230, 3, 10, 5,
		0, 230, 29, 1, 0, 0, 0, 231, 232, 3, 80, 40, 0, 232, 233, 5, 17, 0, 0,
		233, 234, 3, 80, 40, 0, 234, 31, 1, 0, 0, 0, 235, 236, 5, 18, 0, 0, 236,
		237, 3, 80, 40, 0, 237, 238, 3, 10, 5, 0, 238, 33, 1, 0, 0, 0, 239, 240,
		5, 19, 0, 0, 240, 241, 3, 80, 40, 0, 241, 242, 3, 10, 5, 0, 242, 256, 1,
		0, 0, 0, 243, 244, 5, 19, 0, 0, 244, 245, 3, 80, 40, 0, 245, 246, 3, 10,
		5, 0, 246, 247, 5, 20, 0, 0, 247, 248, 3, 10, 5, 0, 248, 256, 1, 0, 0,
		0, 249, 250, 5, 19, 0, 0, 250, 251, 3, 80, 40, 0, 251, 252, 3, 10, 5, 0,
		252, 253, 5, 20, 0, 0, 253, 254, 3, 34, 17, 0, 254, 256, 1, 0, 0, 0, 255,
		239, 1, 0, 0, 0, 255, 243, 1, 0, 0, 0, 255, 249, 1, 0, 0, 0, 256, 35, 1,
		0, 0, 0, 257, 258, 5, 21, 0, 0, 258, 259, 3, 80, 40, 0, 259, 260, 5, 20,
		0, 0, 260, 261, 5, 10, 0, 0, 261, 262, 3, 12, 6, 0, 262, 263, 7, 1, 0,
		0, 263, 264, 5, 11, 0, 0, 264, 37, 1, 0, 0, 0, 265, 266, 5, 25, 0, 0, 266,
		267, 3, 80, 40, 0, 267, 271, 5, 10, 0, 0, 268, 270, 3, 40, 20, 0, 269,
		268, 1, 0, 0, 0, 270, 273, 1, 0, 0, 0, 271, 269, 1, 0, 0, 0, 271, 272,
		1, 0, 0, 0, 272, 275, 1, 0, 0, 0, 273, 271, 1, 0, 0, 0, 274, 276, 3, 42,
		21, 0, 275, 274, 1, 0, 0, 0, 275, 276, 1, 0, 0, 0, 276, 277, 1, 0, 0, 0,
		277, 278, 5, 11, 0, 0, 278, 39, 1, 0, 0, 0, 279, 280, 5, 26, 0, 0, 280,
		281, 3, 80, 40, 0, 281, 282, 5, 8, 0, 0, 282, 284, 3, 12, 6, 0, 283, 285,
		5, 23, 0, 0, 284, 283, 1, 0, 0, 0, 284, 285, 1, 0, 0, 0, 285, 41, 1, 0,
		0, 0, 286, 287, 5, 27, 0, 0, 287, 288, 5, 8, 0, 0, 288, 290, 3, 12, 6,
		0, 289, 291, 5, 23, 0, 0, 290, 289, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0, 291,
		43, 1, 0, 0, 0, 292, 293, 5, 23, 0, 0, 293, 45, 1, 0, 0, 0, 294, 295, 5,
		22, 0, 0, 295, 47, 1, 0, 0, 0, 296, 298, 5, 24, 0, 0, 297, 299, 3, 80,
		40, 0, 298, 297, 1, 0, 0, 0, 298, 299, 1, 0, 0, 0, 299, 49, 1, 0, 0, 0,
		300, 301, 5, 28, 0, 0, 301, 303, 5, 68, 0, 0, 302, 304, 3, 54, 27, 0, 303,
		302, 1, 0, 0, 0, 303, 304, 1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305, 306,
		5, 29, 0, 0, 306, 307, 3, 80, 40, 0, 307, 51, 1, 0, 0, 0, 308, 309, 5,
		30, 0, 0, 309, 310, 5, 68, 0, 0, 310, 311, 3, 54, 27, 0, 311, 312, 5, 31,
		0, 0, 312, 321, 1, 0, 0, 0, 313, 314, 5, 30, 0, 0, 314, 316, 5, 68, 0,
		0, 315, 317, 3, 54, 27, 0, 316, 315, 1, 0, 0, 0, 316, 317, 1, 0, 0, 0,
		317, 318, 1, 0, 0, 0, 318, 319, 5, 29, 0, 0, 319, 321, 3, 80, 40, 0, 320,
		308, 1, 0, 0, 0, 320, 313, 1, 0, 0, 0, 321, 53, 1, 0, 0, 0, 322, 323, 5,
		8, 0, 0, 323, 324, 3, 56, 28, 0, 324, 55, 1, 0, 0, 0, 325, 326, 7, 2, 0,
		0, 326, 57, 1, 0, 0, 0, 327, 328, 5, 30, 0, 0, 328, 329, 5, 68, 0, 0, 329,
		330, 5, 8, 0, 0, 330, 331, 5, 37, 0, 0, 331, 332, 3, 56, 28, 0, 332, 333,
		5, 38, 0, 0, 333, 334, 3, 60, 30, 0, 334, 59, 1, 0, 0, 0, 335, 336, 5,
		29, 0, 0, 336, 337, 5, 37, 0, 0, 337, 338, 3, 62, 31, 0, 338, 339, 5, 38,
		0, 0, 339, 346, 1, 0, 0, 0, 340, 341, 5, 29, 0, 0, 341, 342, 5, 37, 0,
		0, 342, 346, 5, 38, 0, 0, 343, 344, 5, 29, 0, 0, 344, 346, 5, 68, 0, 0,
		345, 335, 1, 0, 0, 0, 345, 340, 1, 0, 0, 0, 345, 343, 1, 0, 0, 0, 346,
		61, 1, 0, 0, 0, 347, 352, 3, 80, 40, 0, 348, 349, 5, 6, 0, 0, 349, 351,
		3, 80, 40, 0, 350, 348, 1, 0, 0, 0, 351, 354, 1, 0, 0, 0, 352, 350, 1,
		0, 0, 0, 352, 353, 1, 0, 0, 0, 353, 63, 1, 0, 0, 0, 354, 352, 1, 0, 0,
		0, 355, 356, 5, 39, 0, 0, 356, 357, 5, 3, 0, 0, 357, 358, 3, 62, 31, 0,
		358, 359, 5, 4, 0, 0, 359, 65, 1, 0, 0, 0, 360, 361, 5, 68, 0, 0, 361,
		362, 5, 12, 0, 0, 362, 363, 5, 40, 0, 0, 363, 364, 5, 3, 0, 0, 364, 365,
		3, 80, 40, 0, 365, 366, 5, 4, 0, 0, 366, 67, 1, 0, 0, 0, 367, 368, 5, 68,
		0, 0, 368, 369, 5, 12, 0, 0, 369, 370, 5, 41, 0, 0, 370, 371, 5, 3, 0,
		0, 371, 372, 5, 4, 0, 0, 372, 69, 1, 0, 0, 0, 373, 374, 5, 68, 0, 0, 374,
		375, 5, 12, 0, 0, 375, 376, 5, 42, 0, 0, 376, 377, 5, 3, 0, 0, 377, 378,
		5, 43, 0, 0, 378, 379, 5, 8, 0, 0, 379, 380, 3, 80, 40, 0, 380, 381, 5,
		4, 0, 0, 381, 71, 1, 0, 0, 0, 382, 383, 5, 33, 0, 0, 383, 384, 5, 3, 0,
		0, 384, 385, 3, 80, 40, 0, 385, 386, 5, 4, 0, 0, 386, 73, 1, 0, 0, 0, 387,
		388, 5, 34, 0, 0, 388, 389, 5, 3, 0, 0, 389, 390, 3, 80, 40, 0, 390, 391,
		5, 4, 0, 0, 391, 75, 1, 0, 0, 0, 392, 393, 5, 32, 0, 0, 393, 394, 5, 3,
		0, 0, 394, 395, 3, 80, 40, 0, 395, 396, 5, 4, 0, 0, 396, 77, 1, 0, 0, 0,
		397, 398, 5, 68, 0, 0, 398, 399, 5, 29, 0, 0, 399, 428, 3, 80, 40, 0, 400,
		401, 5, 68, 0, 0, 401, 402, 5, 44, 0, 0, 402, 428, 3, 80, 40, 0, 403, 404,
		5, 68, 0, 0, 404, 405, 5, 45, 0, 0, 405, 428, 3, 80, 40, 0, 406, 407, 5,
		68, 0, 0, 407, 408, 5, 37, 0, 0, 408, 409, 3, 80, 40, 0, 409, 410, 5, 38,
		0, 0, 410, 411, 5, 29, 0, 0, 411, 412, 3, 80, 40, 0, 412, 428, 1, 0, 0,
		0, 413, 414, 5, 68, 0, 0, 414, 415, 5, 37, 0, 0, 415, 416, 3, 80, 40, 0,
		416, 417, 5, 38, 0, 0, 417, 418, 5, 44, 0, 0, 418, 419, 3, 80, 40, 0, 419,
		428, 1, 0, 0, 0, 420, 421, 5, 68, 0, 0, 421, 422, 5, 37, 0, 0, 422, 423,
		3, 80, 40, 0, 423, 424, 5, 38, 0, 0, 424, 425, 5, 45, 0, 0, 425, 426, 3,
		80, 40, 0, 426, 428, 1, 0, 0, 0, 427, 397, 1, 0, 0, 0, 427, 400, 1, 0,
		0, 0, 427, 403, 1, 0, 0, 0, 427, 406, 1, 0, 0, 0, 427, 413, 1, 0, 0, 0,
		427, 420, 1, 0, 0, 0, 428, 79, 1, 0, 0, 0, 429, 430, 6, 40, -1, 0, 430,
		448, 3, 82, 41, 0, 431, 448, 3, 26, 13, 0, 432, 448, 3, 24, 12, 0, 433,
		434, 5, 68, 0, 0, 434, 435, 5, 37, 0, 0, 435, 436, 3, 80, 40, 0, 436, 437,
		5, 38, 0, 0, 437, 448, 1, 0, 0, 0, 438, 448, 5, 68, 0, 0, 439, 440, 5,
		3, 0, 0, 440, 441, 3, 80, 40, 0, 441, 442, 5, 4, 0, 0, 442, 448, 1, 0,
		0, 0, 443, 444, 5, 46, 0, 0, 444, 448, 3, 80, 40, 6, 445, 446, 5, 47, 0,
		0, 446, 448, 3, 80, 40, 5, 447, 429, 1, 0, 0, 0, 447, 431, 1, 0, 0, 0,
		447, 432, 1, 0, 0, 0, 447, 433, 1, 0, 0, 0, 447, 438, 1, 0, 0, 0, 447,
		439, 1, 0, 0, 0, 447, 443, 1, 0, 0, 0, 447, 445, 1, 0, 0, 0, 448, 463,
		1, 0, 0, 0, 449, 450, 10, 4, 0, 0, 450, 451, 7, 3, 0, 0, 451, 462, 3, 80,
		40, 5, 452, 453, 10, 3, 0, 0, 453, 454, 7, 4, 0, 0, 454, 462, 3, 80, 40,
		4, 455, 456, 10, 2, 0, 0, 456, 457, 7, 5, 0, 0, 457, 462, 3, 80, 40, 3,
		458, 459, 10, 1, 0, 0, 459, 460, 7, 6, 0, 0, 460, 462, 3, 80, 40, 2, 461,
		449, 1, 0, 0, 0, 461, 452, 1, 0, 0, 0, 461, 455, 1, 0, 0, 0, 461, 458,
		1, 0, 0, 0, 462, 465, 1, 0, 0, 0, 463, 461, 1, 0, 0, 0, 463, 464, 1, 0,
		0, 0, 464, 81, 1, 0, 0, 0, 465, 463, 1, 0, 0, 0, 466, 472, 5, 65, 0, 0,
		467, 472, 5, 63, 0, 0, 468, 472, 5, 64, 0, 0, 469, 472, 5, 66, 0, 0, 470,
		472, 5, 67, 0, 0, 471, 466, 1, 0, 0, 0, 471, 467, 1, 0, 0, 0, 471, 468,
		1, 0, 0, 0, 471, 469, 1, 0, 0, 0, 471, 470, 1, 0, 0, 0, 472, 83, 1, 0,
		0, 0, 45, 90, 95, 99, 103, 107, 111, 114, 120, 125, 131, 135, 141, 145,
		148, 157, 162, 166, 170, 174, 178, 182, 184, 189, 193, 198, 203, 212, 220,
		227, 255, 271, 275, 284, 290, 298, 303, 316, 320, 345, 352, 427, 447, 461,
		463, 471,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// T_swiftParserInit initializes any static state used to implement T_swiftParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewT_swiftParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func T_swiftParserInit() {
	staticData := &T_swiftParserStaticData
	staticData.once.Do(t_swiftParserInit)
}

// NewT_swiftParser produces a new parser instance for the optional input antlr.TokenStream.
func NewT_swiftParser(input antlr.TokenStream) *T_swiftParser {
	T_swiftParserInit()
	this := new(T_swiftParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &T_swiftParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "T_swift.g4"

	return this
}

// T_swiftParser tokens.
const (
	T_swiftParserEOF           = antlr.TokenEOF
	T_swiftParserT__0          = 1
	T_swiftParserT__1          = 2
	T_swiftParserT__2          = 3
	T_swiftParserT__3          = 4
	T_swiftParserT__4          = 5
	T_swiftParserT__5          = 6
	T_swiftParserT__6          = 7
	T_swiftParserT__7          = 8
	T_swiftParserT__8          = 9
	T_swiftParserT__9          = 10
	T_swiftParserT__10         = 11
	T_swiftParserT__11         = 12
	T_swiftParserT__12         = 13
	T_swiftParserT__13         = 14
	T_swiftParserT__14         = 15
	T_swiftParserT__15         = 16
	T_swiftParserT__16         = 17
	T_swiftParserT__17         = 18
	T_swiftParserT__18         = 19
	T_swiftParserT__19         = 20
	T_swiftParserT__20         = 21
	T_swiftParserT__21         = 22
	T_swiftParserT__22         = 23
	T_swiftParserT__23         = 24
	T_swiftParserT__24         = 25
	T_swiftParserT__25         = 26
	T_swiftParserT__26         = 27
	T_swiftParserT__27         = 28
	T_swiftParserT__28         = 29
	T_swiftParserT__29         = 30
	T_swiftParserT__30         = 31
	T_swiftParserT__31         = 32
	T_swiftParserT__32         = 33
	T_swiftParserT__33         = 34
	T_swiftParserT__34         = 35
	T_swiftParserT__35         = 36
	T_swiftParserT__36         = 37
	T_swiftParserT__37         = 38
	T_swiftParserT__38         = 39
	T_swiftParserT__39         = 40
	T_swiftParserT__40         = 41
	T_swiftParserT__41         = 42
	T_swiftParserT__42         = 43
	T_swiftParserT__43         = 44
	T_swiftParserT__44         = 45
	T_swiftParserT__45         = 46
	T_swiftParserT__46         = 47
	T_swiftParserT__47         = 48
	T_swiftParserT__48         = 49
	T_swiftParserT__49         = 50
	T_swiftParserT__50         = 51
	T_swiftParserT__51         = 52
	T_swiftParserT__52         = 53
	T_swiftParserT__53         = 54
	T_swiftParserT__54         = 55
	T_swiftParserT__55         = 56
	T_swiftParserT__56         = 57
	T_swiftParserT__57         = 58
	T_swiftParserT__58         = 59
	T_swiftParserWS            = 60
	T_swiftParserBlock_comment = 61
	T_swiftParserLine_comment  = 62
	T_swiftParserInt           = 63
	T_swiftParserFloat         = 64
	T_swiftParserCaracter      = 65
	T_swiftParserString_       = 66
	T_swiftParserBool          = 67
	T_swiftParserIdentificador = 68
)

// T_swiftParser rules.
const (
	T_swiftParserRULE_inicio                     = 0
	T_swiftParserRULE_instrucciones_globales     = 1
	T_swiftParserRULE_intruccion_global          = 2
	T_swiftParserRULE_function_declaracion       = 3
	T_swiftParserRULE_lista_parametros           = 4
	T_swiftParserRULE_code_block                 = 5
	T_swiftParserRULE_instrucciones              = 6
	T_swiftParserRULE_instruccion                = 7
	T_swiftParserRULE_declaracion                = 8
	T_swiftParserRULE_loop_statement             = 9
	T_swiftParserRULE_branch_statement           = 10
	T_swiftParserRULE_control_transfer_statement = 11
	T_swiftParserRULE_llamadas_funciones         = 12
	T_swiftParserRULE_atributos                  = 13
	T_swiftParserRULE_for_in_statement           = 14
	T_swiftParserRULE_rango                      = 15
	T_swiftParserRULE_while_statement            = 16
	T_swiftParserRULE_if_statement               = 17
	T_swiftParserRULE_guard_statement            = 18
	T_swiftParserRULE_switch_statement           = 19
	T_swiftParserRULE_switch_case                = 20
	T_swiftParserRULE_default_case               = 21
	T_swiftParserRULE_break_statement            = 22
	T_swiftParserRULE_continue_statement         = 23
	T_swiftParserRULE_return_statement           = 24
	T_swiftParserRULE_constant_declaracion       = 25
	T_swiftParserRULE_variable_declaracion       = 26
	T_swiftParserRULE_anotacion_tipo             = 27
	T_swiftParserRULE_tipos                      = 28
	T_swiftParserRULE_array_declaracion          = 29
	T_swiftParserRULE_definicion_vector          = 30
	T_swiftParserRULE_lista_expresiones          = 31
	T_swiftParserRULE_funcion_print              = 32
	T_swiftParserRULE_funcion_append             = 33
	T_swiftParserRULE_funcion_removeLast         = 34
	T_swiftParserRULE_funcion_removeat           = 35
	T_swiftParserRULE_funcion_int                = 36
	T_swiftParserRULE_funcion_float              = 37
	T_swiftParserRULE_funcion_string             = 38
	T_swiftParserRULE_asignacion                 = 39
	T_swiftParserRULE_expresion                  = 40
	T_swiftParserRULE_primitivos                 = 41
)

// IInicioContext is an interface to support dynamic dispatch.
type IInicioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Instrucciones_globales() IInstrucciones_globalesContext
	EOF() antlr.TerminalNode

	// IsInicioContext differentiates from other interfaces.
	IsInicioContext()
}

type InicioContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInicioContext() *InicioContext {
	var p = new(InicioContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_inicio
	return p
}

func InitEmptyInicioContext(p *InicioContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_inicio
}

func (*InicioContext) IsInicioContext() {}

func NewInicioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InicioContext {
	var p = new(InicioContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_inicio

	return p
}

func (s *InicioContext) GetParser() antlr.Parser { return s.parser }

func (s *InicioContext) Instrucciones_globales() IInstrucciones_globalesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstrucciones_globalesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstrucciones_globalesContext)
}

func (s *InicioContext) EOF() antlr.TerminalNode {
	return s.GetToken(T_swiftParserEOF, 0)
}

func (s *InicioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InicioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InicioContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitInicio(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Inicio() (localctx IInicioContext) {
	localctx = NewInicioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, T_swiftParserRULE_inicio)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.Instrucciones_globales()
	}
	{
		p.SetState(85)
		p.Match(T_swiftParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstrucciones_globalesContext is an interface to support dynamic dispatch.
type IInstrucciones_globalesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIntruccion_global() []IIntruccion_globalContext
	Intruccion_global(i int) IIntruccion_globalContext

	// IsInstrucciones_globalesContext differentiates from other interfaces.
	IsInstrucciones_globalesContext()
}

type Instrucciones_globalesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstrucciones_globalesContext() *Instrucciones_globalesContext {
	var p = new(Instrucciones_globalesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_instrucciones_globales
	return p
}

func InitEmptyInstrucciones_globalesContext(p *Instrucciones_globalesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_instrucciones_globales
}

func (*Instrucciones_globalesContext) IsInstrucciones_globalesContext() {}

func NewInstrucciones_globalesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instrucciones_globalesContext {
	var p = new(Instrucciones_globalesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_instrucciones_globales

	return p
}

func (s *Instrucciones_globalesContext) GetParser() antlr.Parser { return s.parser }

func (s *Instrucciones_globalesContext) AllIntruccion_global() []IIntruccion_globalContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIntruccion_globalContext); ok {
			len++
		}
	}

	tst := make([]IIntruccion_globalContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIntruccion_globalContext); ok {
			tst[i] = t.(IIntruccion_globalContext)
			i++
		}
	}

	return tst
}

func (s *Instrucciones_globalesContext) Intruccion_global(i int) IIntruccion_globalContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntruccion_globalContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntruccion_globalContext)
}

func (s *Instrucciones_globalesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instrucciones_globalesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instrucciones_globalesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitInstrucciones_globales(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Instrucciones_globales() (localctx IInstrucciones_globalesContext) {
	localctx = NewInstrucciones_globalesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, T_swiftParserRULE_instrucciones_globales)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&581199233028) != 0) || _la == T_swiftParserIdentificador {
		{
			p.SetState(87)
			p.Intruccion_global()
		}

		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIntruccion_globalContext is an interface to support dynamic dispatch.
type IIntruccion_globalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Declaracion() IDeclaracionContext
	Loop_statement() ILoop_statementContext
	Branch_statement() IBranch_statementContext
	Asignacion() IAsignacionContext
	Llamadas_funciones() ILlamadas_funcionesContext
	Function_declaracion() IFunction_declaracionContext

	// IsIntruccion_globalContext differentiates from other interfaces.
	IsIntruccion_globalContext()
}

type Intruccion_globalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntruccion_globalContext() *Intruccion_globalContext {
	var p = new(Intruccion_globalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_intruccion_global
	return p
}

func InitEmptyIntruccion_globalContext(p *Intruccion_globalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_intruccion_global
}

func (*Intruccion_globalContext) IsIntruccion_globalContext() {}

func NewIntruccion_globalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Intruccion_globalContext {
	var p = new(Intruccion_globalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_intruccion_global

	return p
}

func (s *Intruccion_globalContext) GetParser() antlr.Parser { return s.parser }

func (s *Intruccion_globalContext) Declaracion() IDeclaracionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracionContext)
}

func (s *Intruccion_globalContext) Loop_statement() ILoop_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoop_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoop_statementContext)
}

func (s *Intruccion_globalContext) Branch_statement() IBranch_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBranch_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBranch_statementContext)
}

func (s *Intruccion_globalContext) Asignacion() IAsignacionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignacionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignacionContext)
}

func (s *Intruccion_globalContext) Llamadas_funciones() ILlamadas_funcionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamadas_funcionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamadas_funcionesContext)
}

func (s *Intruccion_globalContext) Function_declaracion() IFunction_declaracionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_declaracionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_declaracionContext)
}

func (s *Intruccion_globalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Intruccion_globalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Intruccion_globalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitIntruccion_global(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Intruccion_global() (localctx IIntruccion_globalContext) {
	localctx = NewIntruccion_globalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, T_swiftParserRULE_intruccion_global)
	var _la int

	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(93)
			p.Declaracion()
		}
		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(94)
				p.Match(T_swiftParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(97)
			p.Loop_statement()
		}
		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(98)
				p.Match(T_swiftParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(101)
			p.Branch_statement()
		}
		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(102)
				p.Match(T_swiftParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(105)
			p.Asignacion()
		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(106)
				p.Match(T_swiftParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(109)
			p.Llamadas_funciones()
		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(110)
				p.Match(T_swiftParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(113)
			p.Function_declaracion()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunction_declaracionContext is an interface to support dynamic dispatch.
type IFunction_declaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identificador() antlr.TerminalNode
	Code_block() ICode_blockContext
	Lista_parametros() ILista_parametrosContext
	Tipos() ITiposContext

	// IsFunction_declaracionContext differentiates from other interfaces.
	IsFunction_declaracionContext()
}

type Function_declaracionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_declaracionContext() *Function_declaracionContext {
	var p = new(Function_declaracionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_function_declaracion
	return p
}

func InitEmptyFunction_declaracionContext(p *Function_declaracionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_function_declaracion
}

func (*Function_declaracionContext) IsFunction_declaracionContext() {}

func NewFunction_declaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_declaracionContext {
	var p = new(Function_declaracionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_function_declaracion

	return p
}

func (s *Function_declaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_declaracionContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Function_declaracionContext) Code_block() ICode_blockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICode_blockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICode_blockContext)
}

func (s *Function_declaracionContext) Lista_parametros() ILista_parametrosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_parametrosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_parametrosContext)
}

func (s *Function_declaracionContext) Tipos() ITiposContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITiposContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITiposContext)
}

func (s *Function_declaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_declaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_declaracionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitFunction_declaracion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Function_declaracion() (localctx IFunction_declaracionContext) {
	localctx = NewFunction_declaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, T_swiftParserRULE_function_declaracion)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.Match(T_swiftParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(117)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(118)
		p.Match(T_swiftParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-6)) & ^0x3f) == 0 && ((int64(1)<<(_la-6))&4611686018427387911) != 0 {
		{
			p.SetState(119)
			p.Lista_parametros()
		}

	}
	{
		p.SetState(122)
		p.Match(T_swiftParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T_swiftParserT__4 {
		{
			p.SetState(123)
			p.Match(T_swiftParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(124)
			p.Tipos()
		}

	}
	{
		p.SetState(127)
		p.Code_block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILista_parametrosContext is an interface to support dynamic dispatch.
type ILista_parametrosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Tipos() ITiposContext
	Lista_parametros() ILista_parametrosContext
	Identificador() antlr.TerminalNode

	// IsLista_parametrosContext differentiates from other interfaces.
	IsLista_parametrosContext()
}

type Lista_parametrosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLista_parametrosContext() *Lista_parametrosContext {
	var p = new(Lista_parametrosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_lista_parametros
	return p
}

func InitEmptyLista_parametrosContext(p *Lista_parametrosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_lista_parametros
}

func (*Lista_parametrosContext) IsLista_parametrosContext() {}

func NewLista_parametrosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lista_parametrosContext {
	var p = new(Lista_parametrosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_lista_parametros

	return p
}

func (s *Lista_parametrosContext) GetParser() antlr.Parser { return s.parser }

func (s *Lista_parametrosContext) Tipos() ITiposContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITiposContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITiposContext)
}

func (s *Lista_parametrosContext) Lista_parametros() ILista_parametrosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_parametrosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_parametrosContext)
}

func (s *Lista_parametrosContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Lista_parametrosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lista_parametrosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Lista_parametrosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitLista_parametros(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Lista_parametros() (localctx ILista_parametrosContext) {
	localctx = NewLista_parametrosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, T_swiftParserRULE_lista_parametros)
	var _la int

	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T_swiftParserT__5:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(129)
			p.Match(T_swiftParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__6 || _la == T_swiftParserIdentificador {
			{
				p.SetState(130)
				_la = p.GetTokenStream().LA(1)

				if !(_la == T_swiftParserT__6 || _la == T_swiftParserIdentificador) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(133)
			p.Match(T_swiftParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__8 {
			{
				p.SetState(134)
				p.Match(T_swiftParserT__8)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(137)
			p.Tipos()
		}
		{
			p.SetState(138)
			p.Lista_parametros()
		}

	case T_swiftParserT__6, T_swiftParserT__7, T_swiftParserIdentificador:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__6 || _la == T_swiftParserIdentificador {
			{
				p.SetState(140)
				_la = p.GetTokenStream().LA(1)

				if !(_la == T_swiftParserT__6 || _la == T_swiftParserIdentificador) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(143)
			p.Match(T_swiftParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__8 {
			{
				p.SetState(144)
				p.Match(T_swiftParserT__8)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(147)
			p.Tipos()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICode_blockContext is an interface to support dynamic dispatch.
type ICode_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Instrucciones() IInstruccionesContext

	// IsCode_blockContext differentiates from other interfaces.
	IsCode_blockContext()
}

type Code_blockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCode_blockContext() *Code_blockContext {
	var p = new(Code_blockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_code_block
	return p
}

func InitEmptyCode_blockContext(p *Code_blockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_code_block
}

func (*Code_blockContext) IsCode_blockContext() {}

func NewCode_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Code_blockContext {
	var p = new(Code_blockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_code_block

	return p
}

func (s *Code_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Code_blockContext) Instrucciones() IInstruccionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Code_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Code_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Code_blockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitCode_block(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Code_block() (localctx ICode_blockContext) {
	localctx = NewCode_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, T_swiftParserRULE_code_block)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.Match(T_swiftParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(151)
		p.Instrucciones()
	}
	{
		p.SetState(152)
		p.Match(T_swiftParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstruccionesContext is an interface to support dynamic dispatch.
type IInstruccionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllInstruccion() []IInstruccionContext
	Instruccion(i int) IInstruccionContext

	// IsInstruccionesContext differentiates from other interfaces.
	IsInstruccionesContext()
}

type InstruccionesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstruccionesContext() *InstruccionesContext {
	var p = new(InstruccionesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_instrucciones
	return p
}

func InitEmptyInstruccionesContext(p *InstruccionesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_instrucciones
}

func (*InstruccionesContext) IsInstruccionesContext() {}

func NewInstruccionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionesContext {
	var p = new(InstruccionesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_instrucciones

	return p
}

func (s *InstruccionesContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionesContext) AllInstruccion() []IInstruccionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstruccionContext); ok {
			len++
		}
	}

	tst := make([]IInstruccionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstruccionContext); ok {
			tst[i] = t.(IInstruccionContext)
			i++
		}
	}

	return tst
}

func (s *InstruccionesContext) Instruccion(i int) IInstruccionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionContext)
}

func (s *InstruccionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitInstrucciones(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Instrucciones() (localctx IInstruccionesContext) {
	localctx = NewInstruccionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, T_swiftParserRULE_instrucciones)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(154)
				p.Instruccion()
			}

		}
		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstruccionContext is an interface to support dynamic dispatch.
type IInstruccionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Declaracion() IDeclaracionContext
	Loop_statement() ILoop_statementContext
	Branch_statement() IBranch_statementContext
	Control_transfer_statement() IControl_transfer_statementContext
	Asignacion() IAsignacionContext
	Llamadas_funciones() ILlamadas_funcionesContext

	// IsInstruccionContext differentiates from other interfaces.
	IsInstruccionContext()
}

type InstruccionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstruccionContext() *InstruccionContext {
	var p = new(InstruccionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_instruccion
	return p
}

func InitEmptyInstruccionContext(p *InstruccionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_instruccion
}

func (*InstruccionContext) IsInstruccionContext() {}

func NewInstruccionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionContext {
	var p = new(InstruccionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_instruccion

	return p
}

func (s *InstruccionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionContext) Declaracion() IDeclaracionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracionContext)
}

func (s *InstruccionContext) Loop_statement() ILoop_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoop_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoop_statementContext)
}

func (s *InstruccionContext) Branch_statement() IBranch_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBranch_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBranch_statementContext)
}

func (s *InstruccionContext) Control_transfer_statement() IControl_transfer_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IControl_transfer_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IControl_transfer_statementContext)
}

func (s *InstruccionContext) Asignacion() IAsignacionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignacionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignacionContext)
}

func (s *InstruccionContext) Llamadas_funciones() ILlamadas_funcionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamadas_funcionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamadas_funcionesContext)
}

func (s *InstruccionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitInstruccion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Instruccion() (localctx IInstruccionContext) {
	localctx = NewInstruccionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, T_swiftParserRULE_instruccion)
	var _la int

	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(160)
			p.Declaracion()
		}
		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(161)
				p.Match(T_swiftParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(164)
			p.Loop_statement()
		}
		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(165)
				p.Match(T_swiftParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(168)
			p.Branch_statement()
		}
		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(169)
				p.Match(T_swiftParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(172)
			p.Control_transfer_statement()
		}
		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(173)
				p.Match(T_swiftParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(176)
			p.Asignacion()
		}
		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(177)
				p.Match(T_swiftParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(180)
			p.Llamadas_funciones()
		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(181)
				p.Match(T_swiftParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclaracionContext is an interface to support dynamic dispatch.
type IDeclaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Constant_declaracion() IConstant_declaracionContext
	Variable_declaracion() IVariable_declaracionContext
	Array_declaracion() IArray_declaracionContext

	// IsDeclaracionContext differentiates from other interfaces.
	IsDeclaracionContext()
}

type DeclaracionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaracionContext() *DeclaracionContext {
	var p = new(DeclaracionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_declaracion
	return p
}

func InitEmptyDeclaracionContext(p *DeclaracionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_declaracion
}

func (*DeclaracionContext) IsDeclaracionContext() {}

func NewDeclaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaracionContext {
	var p = new(DeclaracionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_declaracion

	return p
}

func (s *DeclaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaracionContext) Constant_declaracion() IConstant_declaracionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstant_declaracionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstant_declaracionContext)
}

func (s *DeclaracionContext) Variable_declaracion() IVariable_declaracionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariable_declaracionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariable_declaracionContext)
}

func (s *DeclaracionContext) Array_declaracion() IArray_declaracionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_declaracionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_declaracionContext)
}

func (s *DeclaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaracionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitDeclaracion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Declaracion() (localctx IDeclaracionContext) {
	localctx = NewDeclaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, T_swiftParserRULE_declaracion)
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(186)
			p.Constant_declaracion()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(187)
			p.Variable_declaracion()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(188)
			p.Array_declaracion()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILoop_statementContext is an interface to support dynamic dispatch.
type ILoop_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	For_in_statement() IFor_in_statementContext
	While_statement() IWhile_statementContext

	// IsLoop_statementContext differentiates from other interfaces.
	IsLoop_statementContext()
}

type Loop_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoop_statementContext() *Loop_statementContext {
	var p = new(Loop_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_loop_statement
	return p
}

func InitEmptyLoop_statementContext(p *Loop_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_loop_statement
}

func (*Loop_statementContext) IsLoop_statementContext() {}

func NewLoop_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Loop_statementContext {
	var p = new(Loop_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_loop_statement

	return p
}

func (s *Loop_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Loop_statementContext) For_in_statement() IFor_in_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFor_in_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFor_in_statementContext)
}

func (s *Loop_statementContext) While_statement() IWhile_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhile_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhile_statementContext)
}

func (s *Loop_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Loop_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Loop_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitLoop_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Loop_statement() (localctx ILoop_statementContext) {
	localctx = NewLoop_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, T_swiftParserRULE_loop_statement)
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T_swiftParserT__14:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(191)
			p.For_in_statement()
		}

	case T_swiftParserT__17:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(192)
			p.While_statement()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBranch_statementContext is an interface to support dynamic dispatch.
type IBranch_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBranch_statementContext differentiates from other interfaces.
	IsBranch_statementContext()
}

type Branch_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBranch_statementContext() *Branch_statementContext {
	var p = new(Branch_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_branch_statement
	return p
}

func InitEmptyBranch_statementContext(p *Branch_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_branch_statement
}

func (*Branch_statementContext) IsBranch_statementContext() {}

func NewBranch_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Branch_statementContext {
	var p = new(Branch_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_branch_statement

	return p
}

func (s *Branch_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Branch_statementContext) CopyAll(ctx *Branch_statementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Branch_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Branch_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Declarar_guardContext struct {
	Branch_statementContext
}

func NewDeclarar_guardContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Declarar_guardContext {
	var p = new(Declarar_guardContext)

	InitEmptyBranch_statementContext(&p.Branch_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*Branch_statementContext))

	return p
}

func (s *Declarar_guardContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declarar_guardContext) Guard_statement() IGuard_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGuard_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGuard_statementContext)
}

func (s *Declarar_guardContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitDeclarar_guard(s)

	default:
		return t.VisitChildren(s)
	}
}

type Declarar_ifContext struct {
	Branch_statementContext
}

func NewDeclarar_ifContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Declarar_ifContext {
	var p = new(Declarar_ifContext)

	InitEmptyBranch_statementContext(&p.Branch_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*Branch_statementContext))

	return p
}

func (s *Declarar_ifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declarar_ifContext) If_statement() IIf_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_statementContext)
}

func (s *Declarar_ifContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitDeclarar_if(s)

	default:
		return t.VisitChildren(s)
	}
}

type Declarar_switchContext struct {
	Branch_statementContext
}

func NewDeclarar_switchContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Declarar_switchContext {
	var p = new(Declarar_switchContext)

	InitEmptyBranch_statementContext(&p.Branch_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*Branch_statementContext))

	return p
}

func (s *Declarar_switchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declarar_switchContext) Switch_statement() ISwitch_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitch_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitch_statementContext)
}

func (s *Declarar_switchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitDeclarar_switch(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Branch_statement() (localctx IBranch_statementContext) {
	localctx = NewBranch_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, T_swiftParserRULE_branch_statement)
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T_swiftParserT__18:
		localctx = NewDeclarar_ifContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(195)
			p.If_statement()
		}

	case T_swiftParserT__20:
		localctx = NewDeclarar_guardContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(196)
			p.Guard_statement()
		}

	case T_swiftParserT__24:
		localctx = NewDeclarar_switchContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(197)
			p.Switch_statement()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IControl_transfer_statementContext is an interface to support dynamic dispatch.
type IControl_transfer_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Break_statement() IBreak_statementContext
	Continue_statement() IContinue_statementContext
	Return_statement() IReturn_statementContext

	// IsControl_transfer_statementContext differentiates from other interfaces.
	IsControl_transfer_statementContext()
}

type Control_transfer_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyControl_transfer_statementContext() *Control_transfer_statementContext {
	var p = new(Control_transfer_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_control_transfer_statement
	return p
}

func InitEmptyControl_transfer_statementContext(p *Control_transfer_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_control_transfer_statement
}

func (*Control_transfer_statementContext) IsControl_transfer_statementContext() {}

func NewControl_transfer_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Control_transfer_statementContext {
	var p = new(Control_transfer_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_control_transfer_statement

	return p
}

func (s *Control_transfer_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Control_transfer_statementContext) Break_statement() IBreak_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreak_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreak_statementContext)
}

func (s *Control_transfer_statementContext) Continue_statement() IContinue_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinue_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinue_statementContext)
}

func (s *Control_transfer_statementContext) Return_statement() IReturn_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturn_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturn_statementContext)
}

func (s *Control_transfer_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Control_transfer_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Control_transfer_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitControl_transfer_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Control_transfer_statement() (localctx IControl_transfer_statementContext) {
	localctx = NewControl_transfer_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, T_swiftParserRULE_control_transfer_statement)
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T_swiftParserT__22:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(200)
			p.Break_statement()
		}

	case T_swiftParserT__21:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(201)
			p.Continue_statement()
		}

	case T_swiftParserT__23:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(202)
			p.Return_statement()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILlamadas_funcionesContext is an interface to support dynamic dispatch.
type ILlamadas_funcionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Funcion_print() IFuncion_printContext
	Funcion_append() IFuncion_appendContext
	Funcion_removeLast() IFuncion_removeLastContext
	Funcion_removeat() IFuncion_removeatContext
	Funcion_int() IFuncion_intContext
	Funcion_float() IFuncion_floatContext
	Funcion_string() IFuncion_stringContext

	// IsLlamadas_funcionesContext differentiates from other interfaces.
	IsLlamadas_funcionesContext()
}

type Llamadas_funcionesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLlamadas_funcionesContext() *Llamadas_funcionesContext {
	var p = new(Llamadas_funcionesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_llamadas_funciones
	return p
}

func InitEmptyLlamadas_funcionesContext(p *Llamadas_funcionesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_llamadas_funciones
}

func (*Llamadas_funcionesContext) IsLlamadas_funcionesContext() {}

func NewLlamadas_funcionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Llamadas_funcionesContext {
	var p = new(Llamadas_funcionesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_llamadas_funciones

	return p
}

func (s *Llamadas_funcionesContext) GetParser() antlr.Parser { return s.parser }

func (s *Llamadas_funcionesContext) Funcion_print() IFuncion_printContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncion_printContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncion_printContext)
}

func (s *Llamadas_funcionesContext) Funcion_append() IFuncion_appendContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncion_appendContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncion_appendContext)
}

func (s *Llamadas_funcionesContext) Funcion_removeLast() IFuncion_removeLastContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncion_removeLastContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncion_removeLastContext)
}

func (s *Llamadas_funcionesContext) Funcion_removeat() IFuncion_removeatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncion_removeatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncion_removeatContext)
}

func (s *Llamadas_funcionesContext) Funcion_int() IFuncion_intContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncion_intContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncion_intContext)
}

func (s *Llamadas_funcionesContext) Funcion_float() IFuncion_floatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncion_floatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncion_floatContext)
}

func (s *Llamadas_funcionesContext) Funcion_string() IFuncion_stringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncion_stringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncion_stringContext)
}

func (s *Llamadas_funcionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Llamadas_funcionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Llamadas_funcionesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitLlamadas_funciones(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Llamadas_funciones() (localctx ILlamadas_funcionesContext) {
	localctx = NewLlamadas_funcionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, T_swiftParserRULE_llamadas_funciones)
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(205)
			p.Funcion_print()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(206)
			p.Funcion_append()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(207)
			p.Funcion_removeLast()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(208)
			p.Funcion_removeat()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(209)
			p.Funcion_int()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(210)
			p.Funcion_float()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(211)
			p.Funcion_string()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAtributosContext is an interface to support dynamic dispatch.
type IAtributosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAtributosContext differentiates from other interfaces.
	IsAtributosContext()
}

type AtributosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtributosContext() *AtributosContext {
	var p = new(AtributosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_atributos
	return p
}

func InitEmptyAtributosContext(p *AtributosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_atributos
}

func (*AtributosContext) IsAtributosContext() {}

func NewAtributosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtributosContext {
	var p = new(AtributosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_atributos

	return p
}

func (s *AtributosContext) GetParser() antlr.Parser { return s.parser }

func (s *AtributosContext) CopyAll(ctx *AtributosContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AtributosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtributosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Atributos_vector_emptyContext struct {
	AtributosContext
}

func NewAtributos_vector_emptyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Atributos_vector_emptyContext {
	var p = new(Atributos_vector_emptyContext)

	InitEmptyAtributosContext(&p.AtributosContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtributosContext))

	return p
}

func (s *Atributos_vector_emptyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Atributos_vector_emptyContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Atributos_vector_emptyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitAtributos_vector_empty(s)

	default:
		return t.VisitChildren(s)
	}
}

type Atributos_vector_countContext struct {
	AtributosContext
}

func NewAtributos_vector_countContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Atributos_vector_countContext {
	var p = new(Atributos_vector_countContext)

	InitEmptyAtributosContext(&p.AtributosContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtributosContext))

	return p
}

func (s *Atributos_vector_countContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Atributos_vector_countContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Atributos_vector_countContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitAtributos_vector_count(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Atributos() (localctx IAtributosContext) {
	localctx = NewAtributosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, T_swiftParserRULE_atributos)
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAtributos_vector_emptyContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(214)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(215)
			p.Match(T_swiftParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(216)
			p.Match(T_swiftParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewAtributos_vector_countContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(217)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(218)
			p.Match(T_swiftParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(219)
			p.Match(T_swiftParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFor_in_statementContext is an interface to support dynamic dispatch.
type IFor_in_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identificador() antlr.TerminalNode
	Code_block() ICode_blockContext
	Expresion() IExpresionContext
	Rango() IRangoContext

	// IsFor_in_statementContext differentiates from other interfaces.
	IsFor_in_statementContext()
}

type For_in_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_in_statementContext() *For_in_statementContext {
	var p = new(For_in_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_for_in_statement
	return p
}

func InitEmptyFor_in_statementContext(p *For_in_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_for_in_statement
}

func (*For_in_statementContext) IsFor_in_statementContext() {}

func NewFor_in_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_in_statementContext {
	var p = new(For_in_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_for_in_statement

	return p
}

func (s *For_in_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *For_in_statementContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *For_in_statementContext) Code_block() ICode_blockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICode_blockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICode_blockContext)
}

func (s *For_in_statementContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *For_in_statementContext) Rango() IRangoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRangoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRangoContext)
}

func (s *For_in_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_in_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *For_in_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitFor_in_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) For_in_statement() (localctx IFor_in_statementContext) {
	localctx = NewFor_in_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, T_swiftParserRULE_for_in_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(222)
		p.Match(T_swiftParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(223)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(224)
		p.Match(T_swiftParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(225)
			p.expresion(0)
		}

	case 2:
		{
			p.SetState(226)
			p.Rango()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(229)
		p.Code_block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRangoContext is an interface to support dynamic dispatch.
type IRangoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpresion() []IExpresionContext
	Expresion(i int) IExpresionContext

	// IsRangoContext differentiates from other interfaces.
	IsRangoContext()
}

type RangoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangoContext() *RangoContext {
	var p = new(RangoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_rango
	return p
}

func InitEmptyRangoContext(p *RangoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_rango
}

func (*RangoContext) IsRangoContext() {}

func NewRangoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangoContext {
	var p = new(RangoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_rango

	return p
}

func (s *RangoContext) GetParser() antlr.Parser { return s.parser }

func (s *RangoContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *RangoContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *RangoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitRango(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Rango() (localctx IRangoContext) {
	localctx = NewRangoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, T_swiftParserRULE_rango)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(231)
		p.expresion(0)
	}
	{
		p.SetState(232)
		p.Match(T_swiftParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(233)
		p.expresion(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhile_statementContext is an interface to support dynamic dispatch.
type IWhile_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expresion() IExpresionContext
	Code_block() ICode_blockContext

	// IsWhile_statementContext differentiates from other interfaces.
	IsWhile_statementContext()
}

type While_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhile_statementContext() *While_statementContext {
	var p = new(While_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_while_statement
	return p
}

func InitEmptyWhile_statementContext(p *While_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_while_statement
}

func (*While_statementContext) IsWhile_statementContext() {}

func NewWhile_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *While_statementContext {
	var p = new(While_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_while_statement

	return p
}

func (s *While_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *While_statementContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *While_statementContext) Code_block() ICode_blockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICode_blockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICode_blockContext)
}

func (s *While_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *While_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitWhile_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) While_statement() (localctx IWhile_statementContext) {
	localctx = NewWhile_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, T_swiftParserRULE_while_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		p.Match(T_swiftParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(236)
		p.expresion(0)
	}
	{
		p.SetState(237)
		p.Code_block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIf_statementContext is an interface to support dynamic dispatch.
type IIf_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIf_statementContext differentiates from other interfaces.
	IsIf_statementContext()
}

type If_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_statementContext() *If_statementContext {
	var p = new(If_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_if_statement
	return p
}

func InitEmptyIf_statementContext(p *If_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_if_statement
}

func (*If_statementContext) IsIf_statementContext() {}

func NewIf_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_statementContext {
	var p = new(If_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_if_statement

	return p
}

func (s *If_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *If_statementContext) CopyAll(ctx *If_statementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *If_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Else_finalContext struct {
	If_statementContext
}

func NewElse_finalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Else_finalContext {
	var p = new(Else_finalContext)

	InitEmptyIf_statementContext(&p.If_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*If_statementContext))

	return p
}

func (s *Else_finalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Else_finalContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Else_finalContext) AllCode_block() []ICode_blockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICode_blockContext); ok {
			len++
		}
	}

	tst := make([]ICode_blockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICode_blockContext); ok {
			tst[i] = t.(ICode_blockContext)
			i++
		}
	}

	return tst
}

func (s *Else_finalContext) Code_block(i int) ICode_blockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICode_blockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICode_blockContext)
}

func (s *Else_finalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitElse_final(s)

	default:
		return t.VisitChildren(s)
	}
}

type If_simpleContext struct {
	If_statementContext
}

func NewIf_simpleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *If_simpleContext {
	var p = new(If_simpleContext)

	InitEmptyIf_statementContext(&p.If_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*If_statementContext))

	return p
}

func (s *If_simpleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_simpleContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *If_simpleContext) Code_block() ICode_blockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICode_blockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICode_blockContext)
}

func (s *If_simpleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitIf_simple(s)

	default:
		return t.VisitChildren(s)
	}
}

type Siguiente_ifContext struct {
	If_statementContext
}

func NewSiguiente_ifContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Siguiente_ifContext {
	var p = new(Siguiente_ifContext)

	InitEmptyIf_statementContext(&p.If_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*If_statementContext))

	return p
}

func (s *Siguiente_ifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Siguiente_ifContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Siguiente_ifContext) Code_block() ICode_blockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICode_blockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICode_blockContext)
}

func (s *Siguiente_ifContext) If_statement() IIf_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_statementContext)
}

func (s *Siguiente_ifContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitSiguiente_if(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) If_statement() (localctx IIf_statementContext) {
	localctx = NewIf_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, T_swiftParserRULE_if_statement)
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIf_simpleContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(239)
			p.Match(T_swiftParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(240)
			p.expresion(0)
		}
		{
			p.SetState(241)
			p.Code_block()
		}

	case 2:
		localctx = NewElse_finalContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(243)
			p.Match(T_swiftParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(244)
			p.expresion(0)
		}
		{
			p.SetState(245)
			p.Code_block()
		}
		{
			p.SetState(246)
			p.Match(T_swiftParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(247)
			p.Code_block()
		}

	case 3:
		localctx = NewSiguiente_ifContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(249)
			p.Match(T_swiftParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(250)
			p.expresion(0)
		}
		{
			p.SetState(251)
			p.Code_block()
		}
		{
			p.SetState(252)
			p.Match(T_swiftParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(253)
			p.If_statement()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGuard_statementContext is an interface to support dynamic dispatch.
type IGuard_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expresion() IExpresionContext
	Instrucciones() IInstruccionesContext

	// IsGuard_statementContext differentiates from other interfaces.
	IsGuard_statementContext()
}

type Guard_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGuard_statementContext() *Guard_statementContext {
	var p = new(Guard_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_guard_statement
	return p
}

func InitEmptyGuard_statementContext(p *Guard_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_guard_statement
}

func (*Guard_statementContext) IsGuard_statementContext() {}

func NewGuard_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Guard_statementContext {
	var p = new(Guard_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_guard_statement

	return p
}

func (s *Guard_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Guard_statementContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Guard_statementContext) Instrucciones() IInstruccionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Guard_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Guard_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Guard_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitGuard_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Guard_statement() (localctx IGuard_statementContext) {
	localctx = NewGuard_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, T_swiftParserRULE_guard_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(257)
		p.Match(T_swiftParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(258)
		p.expresion(0)
	}
	{
		p.SetState(259)
		p.Match(T_swiftParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(260)
		p.Match(T_swiftParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(261)
		p.Instrucciones()
	}
	{
		p.SetState(262)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&29360128) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(263)
		p.Match(T_swiftParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitch_statementContext is an interface to support dynamic dispatch.
type ISwitch_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expresion() IExpresionContext
	AllSwitch_case() []ISwitch_caseContext
	Switch_case(i int) ISwitch_caseContext
	Default_case() IDefault_caseContext

	// IsSwitch_statementContext differentiates from other interfaces.
	IsSwitch_statementContext()
}

type Switch_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitch_statementContext() *Switch_statementContext {
	var p = new(Switch_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_switch_statement
	return p
}

func InitEmptySwitch_statementContext(p *Switch_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_switch_statement
}

func (*Switch_statementContext) IsSwitch_statementContext() {}

func NewSwitch_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Switch_statementContext {
	var p = new(Switch_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_switch_statement

	return p
}

func (s *Switch_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Switch_statementContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Switch_statementContext) AllSwitch_case() []ISwitch_caseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISwitch_caseContext); ok {
			len++
		}
	}

	tst := make([]ISwitch_caseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISwitch_caseContext); ok {
			tst[i] = t.(ISwitch_caseContext)
			i++
		}
	}

	return tst
}

func (s *Switch_statementContext) Switch_case(i int) ISwitch_caseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitch_caseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitch_caseContext)
}

func (s *Switch_statementContext) Default_case() IDefault_caseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefault_caseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefault_caseContext)
}

func (s *Switch_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Switch_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitSwitch_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Switch_statement() (localctx ISwitch_statementContext) {
	localctx = NewSwitch_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, T_swiftParserRULE_switch_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(265)
		p.Match(T_swiftParserT__24)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(266)
		p.expresion(0)
	}
	{
		p.SetState(267)
		p.Match(T_swiftParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T_swiftParserT__25 {
		{
			p.SetState(268)
			p.Switch_case()
		}

		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T_swiftParserT__26 {
		{
			p.SetState(274)
			p.Default_case()
		}

	}
	{
		p.SetState(277)
		p.Match(T_swiftParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitch_caseContext is an interface to support dynamic dispatch.
type ISwitch_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expresion() IExpresionContext
	Instrucciones() IInstruccionesContext

	// IsSwitch_caseContext differentiates from other interfaces.
	IsSwitch_caseContext()
}

type Switch_caseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitch_caseContext() *Switch_caseContext {
	var p = new(Switch_caseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_switch_case
	return p
}

func InitEmptySwitch_caseContext(p *Switch_caseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_switch_case
}

func (*Switch_caseContext) IsSwitch_caseContext() {}

func NewSwitch_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Switch_caseContext {
	var p = new(Switch_caseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_switch_case

	return p
}

func (s *Switch_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *Switch_caseContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Switch_caseContext) Instrucciones() IInstruccionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Switch_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Switch_caseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitSwitch_case(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Switch_case() (localctx ISwitch_caseContext) {
	localctx = NewSwitch_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, T_swiftParserRULE_switch_case)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(279)
		p.Match(T_swiftParserT__25)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(280)
		p.expresion(0)
	}
	{
		p.SetState(281)
		p.Match(T_swiftParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(282)
		p.Instrucciones()
	}
	p.SetState(284)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T_swiftParserT__22 {
		{
			p.SetState(283)
			p.Match(T_swiftParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDefault_caseContext is an interface to support dynamic dispatch.
type IDefault_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Instrucciones() IInstruccionesContext

	// IsDefault_caseContext differentiates from other interfaces.
	IsDefault_caseContext()
}

type Default_caseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefault_caseContext() *Default_caseContext {
	var p = new(Default_caseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_default_case
	return p
}

func InitEmptyDefault_caseContext(p *Default_caseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_default_case
}

func (*Default_caseContext) IsDefault_caseContext() {}

func NewDefault_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Default_caseContext {
	var p = new(Default_caseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_default_case

	return p
}

func (s *Default_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *Default_caseContext) Instrucciones() IInstruccionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Default_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Default_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Default_caseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitDefault_case(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Default_case() (localctx IDefault_caseContext) {
	localctx = NewDefault_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, T_swiftParserRULE_default_case)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(286)
		p.Match(T_swiftParserT__26)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(287)
		p.Match(T_swiftParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(288)
		p.Instrucciones()
	}
	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T_swiftParserT__22 {
		{
			p.SetState(289)
			p.Match(T_swiftParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBreak_statementContext is an interface to support dynamic dispatch.
type IBreak_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBreak_statementContext differentiates from other interfaces.
	IsBreak_statementContext()
}

type Break_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreak_statementContext() *Break_statementContext {
	var p = new(Break_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_break_statement
	return p
}

func InitEmptyBreak_statementContext(p *Break_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_break_statement
}

func (*Break_statementContext) IsBreak_statementContext() {}

func NewBreak_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Break_statementContext {
	var p = new(Break_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_break_statement

	return p
}

func (s *Break_statementContext) GetParser() antlr.Parser { return s.parser }
func (s *Break_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Break_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Break_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitBreak_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Break_statement() (localctx IBreak_statementContext) {
	localctx = NewBreak_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, T_swiftParserRULE_break_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(292)
		p.Match(T_swiftParserT__22)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContinue_statementContext is an interface to support dynamic dispatch.
type IContinue_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsContinue_statementContext differentiates from other interfaces.
	IsContinue_statementContext()
}

type Continue_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinue_statementContext() *Continue_statementContext {
	var p = new(Continue_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_continue_statement
	return p
}

func InitEmptyContinue_statementContext(p *Continue_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_continue_statement
}

func (*Continue_statementContext) IsContinue_statementContext() {}

func NewContinue_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Continue_statementContext {
	var p = new(Continue_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_continue_statement

	return p
}

func (s *Continue_statementContext) GetParser() antlr.Parser { return s.parser }
func (s *Continue_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Continue_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Continue_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitContinue_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Continue_statement() (localctx IContinue_statementContext) {
	localctx = NewContinue_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, T_swiftParserRULE_continue_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(294)
		p.Match(T_swiftParserT__21)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturn_statementContext is an interface to support dynamic dispatch.
type IReturn_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expresion() IExpresionContext

	// IsReturn_statementContext differentiates from other interfaces.
	IsReturn_statementContext()
}

type Return_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturn_statementContext() *Return_statementContext {
	var p = new(Return_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_return_statement
	return p
}

func InitEmptyReturn_statementContext(p *Return_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_return_statement
}

func (*Return_statementContext) IsReturn_statementContext() {}

func NewReturn_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Return_statementContext {
	var p = new(Return_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_return_statement

	return p
}

func (s *Return_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Return_statementContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Return_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Return_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Return_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitReturn_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Return_statement() (localctx IReturn_statementContext) {
	localctx = NewReturn_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, T_swiftParserRULE_return_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(296)
		p.Match(T_swiftParserT__23)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(298)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(297)
			p.expresion(0)
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstant_declaracionContext is an interface to support dynamic dispatch.
type IConstant_declaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identificador() antlr.TerminalNode
	Expresion() IExpresionContext
	Anotacion_tipo() IAnotacion_tipoContext

	// IsConstant_declaracionContext differentiates from other interfaces.
	IsConstant_declaracionContext()
}

type Constant_declaracionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstant_declaracionContext() *Constant_declaracionContext {
	var p = new(Constant_declaracionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_constant_declaracion
	return p
}

func InitEmptyConstant_declaracionContext(p *Constant_declaracionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_constant_declaracion
}

func (*Constant_declaracionContext) IsConstant_declaracionContext() {}

func NewConstant_declaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Constant_declaracionContext {
	var p = new(Constant_declaracionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_constant_declaracion

	return p
}

func (s *Constant_declaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *Constant_declaracionContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Constant_declaracionContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Constant_declaracionContext) Anotacion_tipo() IAnotacion_tipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnotacion_tipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnotacion_tipoContext)
}

func (s *Constant_declaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Constant_declaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Constant_declaracionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitConstant_declaracion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Constant_declaracion() (localctx IConstant_declaracionContext) {
	localctx = NewConstant_declaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, T_swiftParserRULE_constant_declaracion)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(300)
		p.Match(T_swiftParserT__27)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(301)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(303)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T_swiftParserT__7 {
		{
			p.SetState(302)
			p.Anotacion_tipo()
		}

	}
	{
		p.SetState(305)
		p.Match(T_swiftParserT__28)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(306)
		p.expresion(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariable_declaracionContext is an interface to support dynamic dispatch.
type IVariable_declaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identificador() antlr.TerminalNode
	Anotacion_tipo() IAnotacion_tipoContext
	Expresion() IExpresionContext

	// IsVariable_declaracionContext differentiates from other interfaces.
	IsVariable_declaracionContext()
}

type Variable_declaracionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariable_declaracionContext() *Variable_declaracionContext {
	var p = new(Variable_declaracionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_variable_declaracion
	return p
}

func InitEmptyVariable_declaracionContext(p *Variable_declaracionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_variable_declaracion
}

func (*Variable_declaracionContext) IsVariable_declaracionContext() {}

func NewVariable_declaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Variable_declaracionContext {
	var p = new(Variable_declaracionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_variable_declaracion

	return p
}

func (s *Variable_declaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *Variable_declaracionContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Variable_declaracionContext) Anotacion_tipo() IAnotacion_tipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnotacion_tipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnotacion_tipoContext)
}

func (s *Variable_declaracionContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Variable_declaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Variable_declaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Variable_declaracionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitVariable_declaracion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Variable_declaracion() (localctx IVariable_declaracionContext) {
	localctx = NewVariable_declaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, T_swiftParserRULE_variable_declaracion)
	var _la int

	p.SetState(320)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(308)
			p.Match(T_swiftParserT__29)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(309)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(310)
			p.Anotacion_tipo()
		}
		{
			p.SetState(311)
			p.Match(T_swiftParserT__30)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(313)
			p.Match(T_swiftParserT__29)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(314)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(316)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__7 {
			{
				p.SetState(315)
				p.Anotacion_tipo()
			}

		}
		{
			p.SetState(318)
			p.Match(T_swiftParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(319)
			p.expresion(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAnotacion_tipoContext is an interface to support dynamic dispatch.
type IAnotacion_tipoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Tipos() ITiposContext

	// IsAnotacion_tipoContext differentiates from other interfaces.
	IsAnotacion_tipoContext()
}

type Anotacion_tipoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnotacion_tipoContext() *Anotacion_tipoContext {
	var p = new(Anotacion_tipoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_anotacion_tipo
	return p
}

func InitEmptyAnotacion_tipoContext(p *Anotacion_tipoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_anotacion_tipo
}

func (*Anotacion_tipoContext) IsAnotacion_tipoContext() {}

func NewAnotacion_tipoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Anotacion_tipoContext {
	var p = new(Anotacion_tipoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_anotacion_tipo

	return p
}

func (s *Anotacion_tipoContext) GetParser() antlr.Parser { return s.parser }

func (s *Anotacion_tipoContext) Tipos() ITiposContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITiposContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITiposContext)
}

func (s *Anotacion_tipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Anotacion_tipoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Anotacion_tipoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitAnotacion_tipo(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Anotacion_tipo() (localctx IAnotacion_tipoContext) {
	localctx = NewAnotacion_tipoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, T_swiftParserRULE_anotacion_tipo)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(322)
		p.Match(T_swiftParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(323)
		p.Tipos()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITiposContext is an interface to support dynamic dispatch.
type ITiposContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTiposContext differentiates from other interfaces.
	IsTiposContext()
}

type TiposContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTiposContext() *TiposContext {
	var p = new(TiposContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_tipos
	return p
}

func InitEmptyTiposContext(p *TiposContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_tipos
}

func (*TiposContext) IsTiposContext() {}

func NewTiposContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TiposContext {
	var p = new(TiposContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_tipos

	return p
}

func (s *TiposContext) GetParser() antlr.Parser { return s.parser }
func (s *TiposContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TiposContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TiposContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitTipos(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Tipos() (localctx ITiposContext) {
	localctx = NewTiposContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, T_swiftParserRULE_tipos)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(325)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&133143986176) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArray_declaracionContext is an interface to support dynamic dispatch.
type IArray_declaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identificador() antlr.TerminalNode
	Tipos() ITiposContext
	Definicion_vector() IDefinicion_vectorContext

	// IsArray_declaracionContext differentiates from other interfaces.
	IsArray_declaracionContext()
}

type Array_declaracionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_declaracionContext() *Array_declaracionContext {
	var p = new(Array_declaracionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_array_declaracion
	return p
}

func InitEmptyArray_declaracionContext(p *Array_declaracionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_array_declaracion
}

func (*Array_declaracionContext) IsArray_declaracionContext() {}

func NewArray_declaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_declaracionContext {
	var p = new(Array_declaracionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_array_declaracion

	return p
}

func (s *Array_declaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_declaracionContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Array_declaracionContext) Tipos() ITiposContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITiposContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITiposContext)
}

func (s *Array_declaracionContext) Definicion_vector() IDefinicion_vectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefinicion_vectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefinicion_vectorContext)
}

func (s *Array_declaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_declaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_declaracionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitArray_declaracion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Array_declaracion() (localctx IArray_declaracionContext) {
	localctx = NewArray_declaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, T_swiftParserRULE_array_declaracion)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(327)
		p.Match(T_swiftParserT__29)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(328)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(329)
		p.Match(T_swiftParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(330)
		p.Match(T_swiftParserT__36)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(331)
		p.Tipos()
	}
	{
		p.SetState(332)
		p.Match(T_swiftParserT__37)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(333)
		p.Definicion_vector()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDefinicion_vectorContext is an interface to support dynamic dispatch.
type IDefinicion_vectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Lista_expresiones() ILista_expresionesContext
	Identificador() antlr.TerminalNode

	// IsDefinicion_vectorContext differentiates from other interfaces.
	IsDefinicion_vectorContext()
}

type Definicion_vectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinicion_vectorContext() *Definicion_vectorContext {
	var p = new(Definicion_vectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_definicion_vector
	return p
}

func InitEmptyDefinicion_vectorContext(p *Definicion_vectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_definicion_vector
}

func (*Definicion_vectorContext) IsDefinicion_vectorContext() {}

func NewDefinicion_vectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Definicion_vectorContext {
	var p = new(Definicion_vectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_definicion_vector

	return p
}

func (s *Definicion_vectorContext) GetParser() antlr.Parser { return s.parser }

func (s *Definicion_vectorContext) Lista_expresiones() ILista_expresionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_expresionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_expresionesContext)
}

func (s *Definicion_vectorContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Definicion_vectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Definicion_vectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Definicion_vectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitDefinicion_vector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Definicion_vector() (localctx IDefinicion_vectorContext) {
	localctx = NewDefinicion_vectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, T_swiftParserRULE_definicion_vector)
	p.SetState(345)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(335)
			p.Match(T_swiftParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(336)
			p.Match(T_swiftParserT__36)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(337)
			p.Lista_expresiones()
		}
		{
			p.SetState(338)
			p.Match(T_swiftParserT__37)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(340)
			p.Match(T_swiftParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(341)
			p.Match(T_swiftParserT__36)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(342)
			p.Match(T_swiftParserT__37)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(343)
			p.Match(T_swiftParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(344)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILista_expresionesContext is an interface to support dynamic dispatch.
type ILista_expresionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpresion() []IExpresionContext
	Expresion(i int) IExpresionContext

	// IsLista_expresionesContext differentiates from other interfaces.
	IsLista_expresionesContext()
}

type Lista_expresionesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLista_expresionesContext() *Lista_expresionesContext {
	var p = new(Lista_expresionesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_lista_expresiones
	return p
}

func InitEmptyLista_expresionesContext(p *Lista_expresionesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_lista_expresiones
}

func (*Lista_expresionesContext) IsLista_expresionesContext() {}

func NewLista_expresionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lista_expresionesContext {
	var p = new(Lista_expresionesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_lista_expresiones

	return p
}

func (s *Lista_expresionesContext) GetParser() antlr.Parser { return s.parser }

func (s *Lista_expresionesContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *Lista_expresionesContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Lista_expresionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lista_expresionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Lista_expresionesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitLista_expresiones(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Lista_expresiones() (localctx ILista_expresionesContext) {
	localctx = NewLista_expresionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, T_swiftParserRULE_lista_expresiones)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(347)
		p.expresion(0)
	}
	p.SetState(352)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T_swiftParserT__5 {
		{
			p.SetState(348)
			p.Match(T_swiftParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(349)
			p.expresion(0)
		}

		p.SetState(354)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncion_printContext is an interface to support dynamic dispatch.
type IFuncion_printContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Lista_expresiones() ILista_expresionesContext

	// IsFuncion_printContext differentiates from other interfaces.
	IsFuncion_printContext()
}

type Funcion_printContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncion_printContext() *Funcion_printContext {
	var p = new(Funcion_printContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_funcion_print
	return p
}

func InitEmptyFuncion_printContext(p *Funcion_printContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_funcion_print
}

func (*Funcion_printContext) IsFuncion_printContext() {}

func NewFuncion_printContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Funcion_printContext {
	var p = new(Funcion_printContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_funcion_print

	return p
}

func (s *Funcion_printContext) GetParser() antlr.Parser { return s.parser }

func (s *Funcion_printContext) Lista_expresiones() ILista_expresionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_expresionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_expresionesContext)
}

func (s *Funcion_printContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Funcion_printContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Funcion_printContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitFuncion_print(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Funcion_print() (localctx IFuncion_printContext) {
	localctx = NewFuncion_printContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, T_swiftParserRULE_funcion_print)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(355)
		p.Match(T_swiftParserT__38)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(356)
		p.Match(T_swiftParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(357)
		p.Lista_expresiones()
	}
	{
		p.SetState(358)
		p.Match(T_swiftParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncion_appendContext is an interface to support dynamic dispatch.
type IFuncion_appendContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identificador() antlr.TerminalNode
	Expresion() IExpresionContext

	// IsFuncion_appendContext differentiates from other interfaces.
	IsFuncion_appendContext()
}

type Funcion_appendContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncion_appendContext() *Funcion_appendContext {
	var p = new(Funcion_appendContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_funcion_append
	return p
}

func InitEmptyFuncion_appendContext(p *Funcion_appendContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_funcion_append
}

func (*Funcion_appendContext) IsFuncion_appendContext() {}

func NewFuncion_appendContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Funcion_appendContext {
	var p = new(Funcion_appendContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_funcion_append

	return p
}

func (s *Funcion_appendContext) GetParser() antlr.Parser { return s.parser }

func (s *Funcion_appendContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Funcion_appendContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Funcion_appendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Funcion_appendContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Funcion_appendContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitFuncion_append(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Funcion_append() (localctx IFuncion_appendContext) {
	localctx = NewFuncion_appendContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, T_swiftParserRULE_funcion_append)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(360)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(361)
		p.Match(T_swiftParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(362)
		p.Match(T_swiftParserT__39)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(363)
		p.Match(T_swiftParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(364)
		p.expresion(0)
	}
	{
		p.SetState(365)
		p.Match(T_swiftParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncion_removeLastContext is an interface to support dynamic dispatch.
type IFuncion_removeLastContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identificador() antlr.TerminalNode

	// IsFuncion_removeLastContext differentiates from other interfaces.
	IsFuncion_removeLastContext()
}

type Funcion_removeLastContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncion_removeLastContext() *Funcion_removeLastContext {
	var p = new(Funcion_removeLastContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_funcion_removeLast
	return p
}

func InitEmptyFuncion_removeLastContext(p *Funcion_removeLastContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_funcion_removeLast
}

func (*Funcion_removeLastContext) IsFuncion_removeLastContext() {}

func NewFuncion_removeLastContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Funcion_removeLastContext {
	var p = new(Funcion_removeLastContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_funcion_removeLast

	return p
}

func (s *Funcion_removeLastContext) GetParser() antlr.Parser { return s.parser }

func (s *Funcion_removeLastContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Funcion_removeLastContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Funcion_removeLastContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Funcion_removeLastContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitFuncion_removeLast(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Funcion_removeLast() (localctx IFuncion_removeLastContext) {
	localctx = NewFuncion_removeLastContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, T_swiftParserRULE_funcion_removeLast)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(367)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(368)
		p.Match(T_swiftParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(369)
		p.Match(T_swiftParserT__40)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(370)
		p.Match(T_swiftParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(371)
		p.Match(T_swiftParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncion_removeatContext is an interface to support dynamic dispatch.
type IFuncion_removeatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identificador() antlr.TerminalNode
	Expresion() IExpresionContext

	// IsFuncion_removeatContext differentiates from other interfaces.
	IsFuncion_removeatContext()
}

type Funcion_removeatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncion_removeatContext() *Funcion_removeatContext {
	var p = new(Funcion_removeatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_funcion_removeat
	return p
}

func InitEmptyFuncion_removeatContext(p *Funcion_removeatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_funcion_removeat
}

func (*Funcion_removeatContext) IsFuncion_removeatContext() {}

func NewFuncion_removeatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Funcion_removeatContext {
	var p = new(Funcion_removeatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_funcion_removeat

	return p
}

func (s *Funcion_removeatContext) GetParser() antlr.Parser { return s.parser }

func (s *Funcion_removeatContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Funcion_removeatContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Funcion_removeatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Funcion_removeatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Funcion_removeatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitFuncion_removeat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Funcion_removeat() (localctx IFuncion_removeatContext) {
	localctx = NewFuncion_removeatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, T_swiftParserRULE_funcion_removeat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(373)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(374)
		p.Match(T_swiftParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(375)
		p.Match(T_swiftParserT__41)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(376)
		p.Match(T_swiftParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(377)
		p.Match(T_swiftParserT__42)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(378)
		p.Match(T_swiftParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(379)
		p.expresion(0)
	}
	{
		p.SetState(380)
		p.Match(T_swiftParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncion_intContext is an interface to support dynamic dispatch.
type IFuncion_intContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expresion() IExpresionContext

	// IsFuncion_intContext differentiates from other interfaces.
	IsFuncion_intContext()
}

type Funcion_intContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncion_intContext() *Funcion_intContext {
	var p = new(Funcion_intContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_funcion_int
	return p
}

func InitEmptyFuncion_intContext(p *Funcion_intContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_funcion_int
}

func (*Funcion_intContext) IsFuncion_intContext() {}

func NewFuncion_intContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Funcion_intContext {
	var p = new(Funcion_intContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_funcion_int

	return p
}

func (s *Funcion_intContext) GetParser() antlr.Parser { return s.parser }

func (s *Funcion_intContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Funcion_intContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Funcion_intContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Funcion_intContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitFuncion_int(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Funcion_int() (localctx IFuncion_intContext) {
	localctx = NewFuncion_intContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, T_swiftParserRULE_funcion_int)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(382)
		p.Match(T_swiftParserT__32)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(383)
		p.Match(T_swiftParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(384)
		p.expresion(0)
	}
	{
		p.SetState(385)
		p.Match(T_swiftParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncion_floatContext is an interface to support dynamic dispatch.
type IFuncion_floatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expresion() IExpresionContext

	// IsFuncion_floatContext differentiates from other interfaces.
	IsFuncion_floatContext()
}

type Funcion_floatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncion_floatContext() *Funcion_floatContext {
	var p = new(Funcion_floatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_funcion_float
	return p
}

func InitEmptyFuncion_floatContext(p *Funcion_floatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_funcion_float
}

func (*Funcion_floatContext) IsFuncion_floatContext() {}

func NewFuncion_floatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Funcion_floatContext {
	var p = new(Funcion_floatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_funcion_float

	return p
}

func (s *Funcion_floatContext) GetParser() antlr.Parser { return s.parser }

func (s *Funcion_floatContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Funcion_floatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Funcion_floatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Funcion_floatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitFuncion_float(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Funcion_float() (localctx IFuncion_floatContext) {
	localctx = NewFuncion_floatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, T_swiftParserRULE_funcion_float)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(387)
		p.Match(T_swiftParserT__33)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(388)
		p.Match(T_swiftParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(389)
		p.expresion(0)
	}
	{
		p.SetState(390)
		p.Match(T_swiftParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncion_stringContext is an interface to support dynamic dispatch.
type IFuncion_stringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expresion() IExpresionContext

	// IsFuncion_stringContext differentiates from other interfaces.
	IsFuncion_stringContext()
}

type Funcion_stringContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncion_stringContext() *Funcion_stringContext {
	var p = new(Funcion_stringContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_funcion_string
	return p
}

func InitEmptyFuncion_stringContext(p *Funcion_stringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_funcion_string
}

func (*Funcion_stringContext) IsFuncion_stringContext() {}

func NewFuncion_stringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Funcion_stringContext {
	var p = new(Funcion_stringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_funcion_string

	return p
}

func (s *Funcion_stringContext) GetParser() antlr.Parser { return s.parser }

func (s *Funcion_stringContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Funcion_stringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Funcion_stringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Funcion_stringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitFuncion_string(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Funcion_string() (localctx IFuncion_stringContext) {
	localctx = NewFuncion_stringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, T_swiftParserRULE_funcion_string)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(392)
		p.Match(T_swiftParserT__31)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(393)
		p.Match(T_swiftParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(394)
		p.expresion(0)
	}
	{
		p.SetState(395)
		p.Match(T_swiftParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAsignacionContext is an interface to support dynamic dispatch.
type IAsignacionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAsignacionContext differentiates from other interfaces.
	IsAsignacionContext()
}

type AsignacionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsignacionContext() *AsignacionContext {
	var p = new(AsignacionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_asignacion
	return p
}

func InitEmptyAsignacionContext(p *AsignacionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_asignacion
}

func (*AsignacionContext) IsAsignacionContext() {}

func NewAsignacionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignacionContext {
	var p = new(AsignacionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_asignacion

	return p
}

func (s *AsignacionContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignacionContext) CopyAll(ctx *AsignacionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AsignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Asignacion_normalContext struct {
	AsignacionContext
}

func NewAsignacion_normalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Asignacion_normalContext {
	var p = new(Asignacion_normalContext)

	InitEmptyAsignacionContext(&p.AsignacionContext)
	p.parser = parser
	p.CopyAll(ctx.(*AsignacionContext))

	return p
}

func (s *Asignacion_normalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Asignacion_normalContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Asignacion_normalContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Asignacion_normalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitAsignacion_normal(s)

	default:
		return t.VisitChildren(s)
	}
}

type Asignacion_decrementoContext struct {
	AsignacionContext
}

func NewAsignacion_decrementoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Asignacion_decrementoContext {
	var p = new(Asignacion_decrementoContext)

	InitEmptyAsignacionContext(&p.AsignacionContext)
	p.parser = parser
	p.CopyAll(ctx.(*AsignacionContext))

	return p
}

func (s *Asignacion_decrementoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Asignacion_decrementoContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Asignacion_decrementoContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Asignacion_decrementoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitAsignacion_decremento(s)

	default:
		return t.VisitChildren(s)
	}
}

type Asignacion_vectorContext struct {
	AsignacionContext
}

func NewAsignacion_vectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Asignacion_vectorContext {
	var p = new(Asignacion_vectorContext)

	InitEmptyAsignacionContext(&p.AsignacionContext)
	p.parser = parser
	p.CopyAll(ctx.(*AsignacionContext))

	return p
}

func (s *Asignacion_vectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Asignacion_vectorContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Asignacion_vectorContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *Asignacion_vectorContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Asignacion_vectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitAsignacion_vector(s)

	default:
		return t.VisitChildren(s)
	}
}

type Asignacion_incremento_vectorContext struct {
	AsignacionContext
}

func NewAsignacion_incremento_vectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Asignacion_incremento_vectorContext {
	var p = new(Asignacion_incremento_vectorContext)

	InitEmptyAsignacionContext(&p.AsignacionContext)
	p.parser = parser
	p.CopyAll(ctx.(*AsignacionContext))

	return p
}

func (s *Asignacion_incremento_vectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Asignacion_incremento_vectorContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Asignacion_incremento_vectorContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *Asignacion_incremento_vectorContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Asignacion_incremento_vectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitAsignacion_incremento_vector(s)

	default:
		return t.VisitChildren(s)
	}
}

type Asignacion_incrementoContext struct {
	AsignacionContext
}

func NewAsignacion_incrementoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Asignacion_incrementoContext {
	var p = new(Asignacion_incrementoContext)

	InitEmptyAsignacionContext(&p.AsignacionContext)
	p.parser = parser
	p.CopyAll(ctx.(*AsignacionContext))

	return p
}

func (s *Asignacion_incrementoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Asignacion_incrementoContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Asignacion_incrementoContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Asignacion_incrementoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitAsignacion_incremento(s)

	default:
		return t.VisitChildren(s)
	}
}

type Asignacion_decremento_vectorContext struct {
	AsignacionContext
}

func NewAsignacion_decremento_vectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Asignacion_decremento_vectorContext {
	var p = new(Asignacion_decremento_vectorContext)

	InitEmptyAsignacionContext(&p.AsignacionContext)
	p.parser = parser
	p.CopyAll(ctx.(*AsignacionContext))

	return p
}

func (s *Asignacion_decremento_vectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Asignacion_decremento_vectorContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Asignacion_decremento_vectorContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *Asignacion_decremento_vectorContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Asignacion_decremento_vectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitAsignacion_decremento_vector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Asignacion() (localctx IAsignacionContext) {
	localctx = NewAsignacionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, T_swiftParserRULE_asignacion)
	p.SetState(427)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAsignacion_normalContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(397)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(398)
			p.Match(T_swiftParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(399)
			p.expresion(0)
		}

	case 2:
		localctx = NewAsignacion_incrementoContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(400)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(401)
			p.Match(T_swiftParserT__43)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(402)
			p.expresion(0)
		}

	case 3:
		localctx = NewAsignacion_decrementoContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(403)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(404)
			p.Match(T_swiftParserT__44)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(405)
			p.expresion(0)
		}

	case 4:
		localctx = NewAsignacion_vectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(406)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(407)
			p.Match(T_swiftParserT__36)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(408)
			p.expresion(0)
		}
		{
			p.SetState(409)
			p.Match(T_swiftParserT__37)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(410)
			p.Match(T_swiftParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(411)
			p.expresion(0)
		}

	case 5:
		localctx = NewAsignacion_incremento_vectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(413)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(414)
			p.Match(T_swiftParserT__36)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(415)
			p.expresion(0)
		}
		{
			p.SetState(416)
			p.Match(T_swiftParserT__37)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(417)
			p.Match(T_swiftParserT__43)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(418)
			p.expresion(0)
		}

	case 6:
		localctx = NewAsignacion_decremento_vectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(420)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(421)
			p.Match(T_swiftParserT__36)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(422)
			p.expresion(0)
		}
		{
			p.SetState(423)
			p.Match(T_swiftParserT__37)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(424)
			p.Match(T_swiftParserT__44)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(425)
			p.expresion(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpresionContext is an interface to support dynamic dispatch.
type IExpresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpresionContext differentiates from other interfaces.
	IsExpresionContext()
}

type ExpresionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpresionContext() *ExpresionContext {
	var p = new(ExpresionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_expresion
	return p
}

func InitEmptyExpresionContext(p *ExpresionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_expresion
}

func (*ExpresionContext) IsExpresionContext() {}

func NewExpresionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpresionContext {
	var p = new(ExpresionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_expresion

	return p
}

func (s *ExpresionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpresionContext) CopyAll(ctx *ExpresionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Expresion_llamadaContext struct {
	ExpresionContext
}

func NewExpresion_llamadaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expresion_llamadaContext {
	var p = new(Expresion_llamadaContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *Expresion_llamadaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expresion_llamadaContext) Llamadas_funciones() ILlamadas_funcionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamadas_funcionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamadas_funcionesContext)
}

func (s *Expresion_llamadaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitExpresion_llamada(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expresion_idContext struct {
	ExpresionContext
}

func NewExpresion_idContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expresion_idContext {
	var p = new(Expresion_idContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *Expresion_idContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expresion_idContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Expresion_idContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitExpresion_id(s)

	default:
		return t.VisitChildren(s)
	}
}

type Valor_primitivoContext struct {
	ExpresionContext
}

func NewValor_primitivoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Valor_primitivoContext {
	var p = new(Valor_primitivoContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *Valor_primitivoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Valor_primitivoContext) Primitivos() IPrimitivosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitivosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitivosContext)
}

func (s *Valor_primitivoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitValor_primitivo(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expresion_atributosContext struct {
	ExpresionContext
}

func NewExpresion_atributosContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expresion_atributosContext {
	var p = new(Expresion_atributosContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *Expresion_atributosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expresion_atributosContext) Atributos() IAtributosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtributosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtributosContext)
}

func (s *Expresion_atributosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitExpresion_atributos(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expresion_unarioContext struct {
	ExpresionContext
}

func NewExpresion_unarioContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expresion_unarioContext {
	var p = new(Expresion_unarioContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *Expresion_unarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expresion_unarioContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Expresion_unarioContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitExpresion_unario(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expresion_parenContext struct {
	ExpresionContext
}

func NewExpresion_parenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expresion_parenContext {
	var p = new(Expresion_parenContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *Expresion_parenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expresion_parenContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Expresion_parenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitExpresion_paren(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expresion_relaContext struct {
	ExpresionContext
	op antlr.Token
}

func NewExpresion_relaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expresion_relaContext {
	var p = new(Expresion_relaContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *Expresion_relaContext) GetOp() antlr.Token { return s.op }

func (s *Expresion_relaContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expresion_relaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expresion_relaContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *Expresion_relaContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Expresion_relaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitExpresion_rela(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expresion_aritContext struct {
	ExpresionContext
	op antlr.Token
}

func NewExpresion_aritContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expresion_aritContext {
	var p = new(Expresion_aritContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *Expresion_aritContext) GetOp() antlr.Token { return s.op }

func (s *Expresion_aritContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expresion_aritContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expresion_aritContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *Expresion_aritContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Expresion_aritContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitExpresion_arit(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expresion_compaContext struct {
	ExpresionContext
	op antlr.Token
}

func NewExpresion_compaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expresion_compaContext {
	var p = new(Expresion_compaContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *Expresion_compaContext) GetOp() antlr.Token { return s.op }

func (s *Expresion_compaContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expresion_compaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expresion_compaContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *Expresion_compaContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Expresion_compaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitExpresion_compa(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expresion_vectorContext struct {
	ExpresionContext
}

func NewExpresion_vectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expresion_vectorContext {
	var p = new(Expresion_vectorContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *Expresion_vectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expresion_vectorContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Expresion_vectorContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Expresion_vectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitExpresion_vector(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expresion_negaContext struct {
	ExpresionContext
	op antlr.Token
}

func NewExpresion_negaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expresion_negaContext {
	var p = new(Expresion_negaContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *Expresion_negaContext) GetOp() antlr.Token { return s.op }

func (s *Expresion_negaContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expresion_negaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expresion_negaContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Expresion_negaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitExpresion_nega(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Expresion() (localctx IExpresionContext) {
	return p.expresion(0)
}

func (p *T_swiftParser) expresion(_p int) (localctx IExpresionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpresionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpresionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 80
	p.EnterRecursionRule(localctx, 80, T_swiftParserRULE_expresion, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(447)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) {
	case 1:
		localctx = NewValor_primitivoContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(430)
			p.Primitivos()
		}

	case 2:
		localctx = NewExpresion_atributosContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(431)
			p.Atributos()
		}

	case 3:
		localctx = NewExpresion_llamadaContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(432)
			p.Llamadas_funciones()
		}

	case 4:
		localctx = NewExpresion_vectorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(433)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(434)
			p.Match(T_swiftParserT__36)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(435)
			p.expresion(0)
		}
		{
			p.SetState(436)
			p.Match(T_swiftParserT__37)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewExpresion_idContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(438)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewExpresion_parenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(439)
			p.Match(T_swiftParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(440)
			p.expresion(0)
		}
		{
			p.SetState(441)
			p.Match(T_swiftParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewExpresion_negaContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(443)

			var _m = p.Match(T_swiftParserT__45)

			localctx.(*Expresion_negaContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(444)
			p.expresion(6)
		}

	case 8:
		localctx = NewExpresion_unarioContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(445)
			p.Match(T_swiftParserT__46)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(446)
			p.expresion(5)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(463)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(461)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresion_aritContext(p, NewExpresionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, T_swiftParserRULE_expresion)
				p.SetState(449)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(450)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expresion_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1970324836974592) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expresion_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(451)
					p.expresion(5)
				}

			case 2:
				localctx = NewExpresion_aritContext(p, NewExpresionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, T_swiftParserRULE_expresion)
				p.SetState(452)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(453)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expresion_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == T_swiftParserT__46 || _la == T_swiftParserT__50) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expresion_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(454)
					p.expresion(4)
				}

			case 3:
				localctx = NewExpresion_compaContext(p, NewExpresionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, T_swiftParserRULE_expresion)
				p.SetState(455)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(456)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expresion_compaContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&283726776524341248) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expresion_compaContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(457)
					p.expresion(3)
				}

			case 4:
				localctx = NewExpresion_relaContext(p, NewExpresionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, T_swiftParserRULE_expresion)
				p.SetState(458)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(459)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expresion_relaContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == T_swiftParserT__57 || _la == T_swiftParserT__58) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expresion_relaContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(460)
					p.expresion(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(465)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimitivosContext is an interface to support dynamic dispatch.
type IPrimitivosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrimitivosContext differentiates from other interfaces.
	IsPrimitivosContext()
}

type PrimitivosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitivosContext() *PrimitivosContext {
	var p = new(PrimitivosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_primitivos
	return p
}

func InitEmptyPrimitivosContext(p *PrimitivosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_primitivos
}

func (*PrimitivosContext) IsPrimitivosContext() {}

func NewPrimitivosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitivosContext {
	var p = new(PrimitivosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_primitivos

	return p
}

func (s *PrimitivosContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitivosContext) CopyAll(ctx *PrimitivosContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PrimitivosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitivosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Primitivo_boolContext struct {
	PrimitivosContext
}

func NewPrimitivo_boolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Primitivo_boolContext {
	var p = new(Primitivo_boolContext)

	InitEmptyPrimitivosContext(&p.PrimitivosContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimitivosContext))

	return p
}

func (s *Primitivo_boolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primitivo_boolContext) Bool() antlr.TerminalNode {
	return s.GetToken(T_swiftParserBool, 0)
}

func (s *Primitivo_boolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitPrimitivo_bool(s)

	default:
		return t.VisitChildren(s)
	}
}

type Primitivo_stringContext struct {
	PrimitivosContext
}

func NewPrimitivo_stringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Primitivo_stringContext {
	var p = new(Primitivo_stringContext)

	InitEmptyPrimitivosContext(&p.PrimitivosContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimitivosContext))

	return p
}

func (s *Primitivo_stringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primitivo_stringContext) String_() antlr.TerminalNode {
	return s.GetToken(T_swiftParserString_, 0)
}

func (s *Primitivo_stringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitPrimitivo_string(s)

	default:
		return t.VisitChildren(s)
	}
}

type Primitivo_floatContext struct {
	PrimitivosContext
}

func NewPrimitivo_floatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Primitivo_floatContext {
	var p = new(Primitivo_floatContext)

	InitEmptyPrimitivosContext(&p.PrimitivosContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimitivosContext))

	return p
}

func (s *Primitivo_floatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primitivo_floatContext) Float() antlr.TerminalNode {
	return s.GetToken(T_swiftParserFloat, 0)
}

func (s *Primitivo_floatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitPrimitivo_float(s)

	default:
		return t.VisitChildren(s)
	}
}

type Primitivo_caracterContext struct {
	PrimitivosContext
}

func NewPrimitivo_caracterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Primitivo_caracterContext {
	var p = new(Primitivo_caracterContext)

	InitEmptyPrimitivosContext(&p.PrimitivosContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimitivosContext))

	return p
}

func (s *Primitivo_caracterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primitivo_caracterContext) Caracter() antlr.TerminalNode {
	return s.GetToken(T_swiftParserCaracter, 0)
}

func (s *Primitivo_caracterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitPrimitivo_caracter(s)

	default:
		return t.VisitChildren(s)
	}
}

type Primitivo_intContext struct {
	PrimitivosContext
}

func NewPrimitivo_intContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Primitivo_intContext {
	var p = new(Primitivo_intContext)

	InitEmptyPrimitivosContext(&p.PrimitivosContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimitivosContext))

	return p
}

func (s *Primitivo_intContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primitivo_intContext) Int() antlr.TerminalNode {
	return s.GetToken(T_swiftParserInt, 0)
}

func (s *Primitivo_intContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitPrimitivo_int(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Primitivos() (localctx IPrimitivosContext) {
	localctx = NewPrimitivosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, T_swiftParserRULE_primitivos)
	p.SetState(471)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T_swiftParserCaracter:
		localctx = NewPrimitivo_caracterContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(466)
			p.Match(T_swiftParserCaracter)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case T_swiftParserInt:
		localctx = NewPrimitivo_intContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(467)
			p.Match(T_swiftParserInt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case T_swiftParserFloat:
		localctx = NewPrimitivo_floatContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(468)
			p.Match(T_swiftParserFloat)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case T_swiftParserString_:
		localctx = NewPrimitivo_stringContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(469)
			p.Match(T_swiftParserString_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case T_swiftParserBool:
		localctx = NewPrimitivo_boolContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(470)
			p.Match(T_swiftParserBool)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *T_swiftParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 40:
		var t *ExpresionContext = nil
		if localctx != nil {
			t = localctx.(*ExpresionContext)
		}
		return p.Expresion_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *T_swiftParser) Expresion_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
