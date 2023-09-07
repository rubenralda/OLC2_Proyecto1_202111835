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
		"", "';'", "'func'", "'('", "')'", "'->'", "','", "':'", "'inout'",
		"'{'", "'}'", "'struct'", "'let'", "'var'", "'='", "'mutating'", "'&'",
		"'.'", "'IsEmpty'", "'count'", "'+='", "'-='", "'for'", "'in'", "'...'",
		"'while'", "'if'", "'else'", "'guard'", "'continue'", "'break'", "'return'",
		"'switch'", "'case'", "'default'", "'?'", "'String'", "'Int'", "'Float'",
		"'Bool'", "'Character'", "'['", "']'", "'print'", "'append'", "'removeLast'",
		"'remove'", "'at'", "'!'", "'-'", "'/'", "'%'", "'*'", "'+'", "'<'",
		"'<='", "'>='", "'>'", "'=='", "'!='", "'&&'", "'||'", "'nil'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "WS", "Block_comment",
		"Line_comment", "Int", "Float", "Caracter", "String", "Bool", "Identificador",
	}
	staticData.RuleNames = []string{
		"inicio", "instrucciones_globales", "intruccion_global", "function_declaracion",
		"lista_parametros", "declaracion_parametro", "code_block", "struct_declaracion",
		"lista_atributos", "instrucciones", "instruccion", "declaracion", "loop_statement",
		"branch_statement", "control_transfer_statement", "llamadas_funciones",
		"llamada_normal", "lista_argumentos", "declaracion_argumento", "llamada_metodos",
		"atributos", "asignar_atributos", "for_in_statement", "rango", "while_statement",
		"if_statement", "guard_statement", "switch_statement", "switch_case",
		"default_case", "break_statement", "continue_statement", "return_statement",
		"constant_declaracion", "variable_declaracion", "anotacion_tipo", "tipos",
		"array_declaracion", "definicion_vector", "lista_expresiones", "funcion_print",
		"funcion_append", "funcion_removeLast", "funcion_removeat", "funcion_int",
		"funcion_float", "funcion_string", "asignacion", "expresion", "l_duble",
		"primitivos",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 71, 617, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 1, 0, 1, 0, 1, 0, 1, 1,
		5, 1, 107, 8, 1, 10, 1, 12, 1, 110, 9, 1, 1, 2, 1, 2, 3, 2, 114, 8, 2,
		1, 2, 1, 2, 3, 2, 118, 8, 2, 1, 2, 1, 2, 3, 2, 122, 8, 2, 1, 2, 1, 2, 3,
		2, 126, 8, 2, 1, 2, 1, 2, 3, 2, 130, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2,
		136, 8, 2, 3, 2, 138, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 144, 8, 3, 1,
		3, 1, 3, 1, 3, 3, 3, 149, 8, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 5, 4, 156,
		8, 4, 10, 4, 12, 4, 159, 9, 4, 1, 5, 3, 5, 162, 8, 5, 1, 5, 1, 5, 1, 5,
		3, 5, 167, 8, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7,
		1, 7, 5, 7, 179, 8, 7, 10, 7, 12, 7, 182, 9, 7, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 8, 1, 8, 3, 8, 190, 8, 8, 1, 8, 1, 8, 3, 8, 194, 8, 8, 1, 8, 3, 8, 197,
		8, 8, 1, 8, 3, 8, 200, 8, 8, 1, 8, 3, 8, 203, 8, 8, 1, 9, 5, 9, 206, 8,
		9, 10, 9, 12, 9, 209, 9, 9, 1, 10, 1, 10, 3, 10, 213, 8, 10, 1, 10, 1,
		10, 3, 10, 217, 8, 10, 1, 10, 1, 10, 3, 10, 221, 8, 10, 1, 10, 1, 10, 3,
		10, 225, 8, 10, 1, 10, 1, 10, 3, 10, 229, 8, 10, 1, 10, 1, 10, 3, 10, 233,
		8, 10, 1, 10, 1, 10, 3, 10, 237, 8, 10, 3, 10, 239, 8, 10, 1, 11, 1, 11,
		1, 11, 3, 11, 244, 8, 11, 1, 12, 1, 12, 3, 12, 248, 8, 12, 1, 13, 1, 13,
		1, 13, 3, 13, 253, 8, 13, 1, 14, 1, 14, 1, 14, 3, 14, 258, 8, 14, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 269, 8,
		15, 1, 16, 1, 16, 1, 16, 3, 16, 274, 8, 16, 1, 16, 1, 16, 1, 17, 1, 17,
		1, 17, 5, 17, 281, 8, 17, 10, 17, 12, 17, 284, 9, 17, 1, 18, 1, 18, 3,
		18, 288, 8, 18, 1, 18, 3, 18, 291, 8, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 3, 19, 300, 8, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 4, 20, 313, 8, 20, 11, 20, 12,
		20, 314, 3, 20, 317, 8, 20, 1, 21, 1, 21, 1, 21, 4, 21, 322, 8, 21, 11,
		21, 12, 21, 323, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 4, 21, 331, 8, 21,
		11, 21, 12, 21, 332, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 4, 21, 340, 8,
		21, 11, 21, 12, 21, 341, 1, 21, 1, 21, 3, 21, 346, 8, 21, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 3, 22, 353, 8, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 3, 25, 381, 8, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 5, 27, 395, 8, 27, 10, 27, 12, 27, 398,
		9, 27, 1, 27, 3, 27, 401, 8, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 3, 28, 410, 8, 28, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29, 416, 8,
		29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 3, 32, 424, 8, 32, 1, 33,
		1, 33, 1, 33, 3, 33, 429, 8, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 442, 8, 34, 1, 34, 1, 34,
		3, 34, 446, 8, 34, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38,
		1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 3, 38, 471, 8, 38, 1, 39, 1, 39, 1,
		39, 5, 39, 476, 8, 39, 10, 39, 12, 39, 479, 9, 39, 1, 40, 1, 40, 1, 40,
		1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 42, 1,
		42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43,
		1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1,
		45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 47, 1, 47, 1, 47,
		1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1,
		47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47,
		1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 3, 47, 553, 8, 47, 1, 48, 1,
		48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48,
		1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1,
		48, 3, 48, 578, 8, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48,
		1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 5, 48, 592, 8, 48, 10, 48, 12, 48, 595,
		9, 48, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 5, 49, 604, 8,
		49, 10, 49, 12, 49, 607, 9, 49, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50,
		3, 50, 615, 8, 50, 1, 50, 0, 1, 96, 51, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52,
		54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88,
		90, 92, 94, 96, 98, 100, 0, 7, 1, 0, 12, 13, 1, 0, 29, 31, 2, 0, 36, 40,
		71, 71, 1, 0, 50, 52, 2, 0, 49, 49, 53, 53, 1, 0, 54, 59, 1, 0, 60, 61,
		669, 0, 102, 1, 0, 0, 0, 2, 108, 1, 0, 0, 0, 4, 137, 1, 0, 0, 0, 6, 139,
		1, 0, 0, 0, 8, 152, 1, 0, 0, 0, 10, 161, 1, 0, 0, 0, 12, 170, 1, 0, 0,
		0, 14, 174, 1, 0, 0, 0, 16, 202, 1, 0, 0, 0, 18, 207, 1, 0, 0, 0, 20, 238,
		1, 0, 0, 0, 22, 243, 1, 0, 0, 0, 24, 247, 1, 0, 0, 0, 26, 252, 1, 0, 0,
		0, 28, 257, 1, 0, 0, 0, 30, 268, 1, 0, 0, 0, 32, 270, 1, 0, 0, 0, 34, 277,
		1, 0, 0, 0, 36, 287, 1, 0, 0, 0, 38, 294, 1, 0, 0, 0, 40, 316, 1, 0, 0,
		0, 42, 345, 1, 0, 0, 0, 44, 347, 1, 0, 0, 0, 46, 356, 1, 0, 0, 0, 48, 360,
		1, 0, 0, 0, 50, 380, 1, 0, 0, 0, 52, 382, 1, 0, 0, 0, 54, 390, 1, 0, 0,
		0, 56, 404, 1, 0, 0, 0, 58, 411, 1, 0, 0, 0, 60, 417, 1, 0, 0, 0, 62, 419,
		1, 0, 0, 0, 64, 421, 1, 0, 0, 0, 66, 425, 1, 0, 0, 0, 68, 445, 1, 0, 0,
		0, 70, 447, 1, 0, 0, 0, 72, 450, 1, 0, 0, 0, 74, 452, 1, 0, 0, 0, 76, 470,
		1, 0, 0, 0, 78, 472, 1, 0, 0, 0, 80, 480, 1, 0, 0, 0, 82, 485, 1, 0, 0,
		0, 84, 492, 1, 0, 0, 0, 86, 498, 1, 0, 0, 0, 88, 507, 1, 0, 0, 0, 90, 512,
		1, 0, 0, 0, 92, 517, 1, 0, 0, 0, 94, 552, 1, 0, 0, 0, 96, 577, 1, 0, 0,
		0, 98, 596, 1, 0, 0, 0, 100, 614, 1, 0, 0, 0, 102, 103, 3, 2, 1, 0, 103,
		104, 5, 0, 0, 1, 104, 1, 1, 0, 0, 0, 105, 107, 3, 4, 2, 0, 106, 105, 1,
		0, 0, 0, 107, 110, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 108, 109, 1, 0, 0,
		0, 109, 3, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 111, 113, 3, 22, 11, 0, 112,
		114, 5, 1, 0, 0, 113, 112, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 138,
		1, 0, 0, 0, 115, 117, 3, 24, 12, 0, 116, 118, 5, 1, 0, 0, 117, 116, 1,
		0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 138, 1, 0, 0, 0, 119, 121, 3, 26, 13,
		0, 120, 122, 5, 1, 0, 0, 121, 120, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122,
		138, 1, 0, 0, 0, 123, 125, 3, 94, 47, 0, 124, 126, 5, 1, 0, 0, 125, 124,
		1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 138, 1, 0, 0, 0, 127, 129, 3, 30,
		15, 0, 128, 130, 5, 1, 0, 0, 129, 128, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0,
		130, 138, 1, 0, 0, 0, 131, 138, 3, 6, 3, 0, 132, 138, 3, 14, 7, 0, 133,
		135, 3, 42, 21, 0, 134, 136, 5, 1, 0, 0, 135, 134, 1, 0, 0, 0, 135, 136,
		1, 0, 0, 0, 136, 138, 1, 0, 0, 0, 137, 111, 1, 0, 0, 0, 137, 115, 1, 0,
		0, 0, 137, 119, 1, 0, 0, 0, 137, 123, 1, 0, 0, 0, 137, 127, 1, 0, 0, 0,
		137, 131, 1, 0, 0, 0, 137, 132, 1, 0, 0, 0, 137, 133, 1, 0, 0, 0, 138,
		5, 1, 0, 0, 0, 139, 140, 5, 2, 0, 0, 140, 141, 5, 71, 0, 0, 141, 143, 5,
		3, 0, 0, 142, 144, 3, 8, 4, 0, 143, 142, 1, 0, 0, 0, 143, 144, 1, 0, 0,
		0, 144, 145, 1, 0, 0, 0, 145, 148, 5, 4, 0, 0, 146, 147, 5, 5, 0, 0, 147,
		149, 3, 72, 36, 0, 148, 146, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 150,
		1, 0, 0, 0, 150, 151, 3, 12, 6, 0, 151, 7, 1, 0, 0, 0, 152, 157, 3, 10,
		5, 0, 153, 154, 5, 6, 0, 0, 154, 156, 3, 10, 5, 0, 155, 153, 1, 0, 0, 0,
		156, 159, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158,
		9, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 160, 162, 5, 71, 0, 0, 161, 160, 1,
		0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 164, 5, 71, 0,
		0, 164, 166, 5, 7, 0, 0, 165, 167, 5, 8, 0, 0, 166, 165, 1, 0, 0, 0, 166,
		167, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 169, 3, 72, 36, 0, 169, 11,
		1, 0, 0, 0, 170, 171, 5, 9, 0, 0, 171, 172, 3, 18, 9, 0, 172, 173, 5, 10,
		0, 0, 173, 13, 1, 0, 0, 0, 174, 175, 5, 11, 0, 0, 175, 176, 5, 71, 0, 0,
		176, 180, 5, 9, 0, 0, 177, 179, 3, 16, 8, 0, 178, 177, 1, 0, 0, 0, 179,
		182, 1, 0, 0, 0, 180, 178, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 183,
		1, 0, 0, 0, 182, 180, 1, 0, 0, 0, 183, 184, 5, 10, 0, 0, 184, 15, 1, 0,
		0, 0, 185, 186, 7, 0, 0, 0, 186, 189, 5, 71, 0, 0, 187, 188, 5, 7, 0, 0,
		188, 190, 3, 72, 36, 0, 189, 187, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190,
		193, 1, 0, 0, 0, 191, 192, 5, 14, 0, 0, 192, 194, 3, 96, 48, 0, 193, 191,
		1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 196, 1, 0, 0, 0, 195, 197, 5, 1,
		0, 0, 196, 195, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 203, 1, 0, 0, 0,
		198, 200, 5, 15, 0, 0, 199, 198, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200,
		201, 1, 0, 0, 0, 201, 203, 3, 6, 3, 0, 202, 185, 1, 0, 0, 0, 202, 199,
		1, 0, 0, 0, 203, 17, 1, 0, 0, 0, 204, 206, 3, 20, 10, 0, 205, 204, 1, 0,
		0, 0, 206, 209, 1, 0, 0, 0, 207, 205, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0,
		208, 19, 1, 0, 0, 0, 209, 207, 1, 0, 0, 0, 210, 212, 3, 22, 11, 0, 211,
		213, 5, 1, 0, 0, 212, 211, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 239,
		1, 0, 0, 0, 214, 216, 3, 24, 12, 0, 215, 217, 5, 1, 0, 0, 216, 215, 1,
		0, 0, 0, 216, 217, 1, 0, 0, 0, 217, 239, 1, 0, 0, 0, 218, 220, 3, 26, 13,
		0, 219, 221, 5, 1, 0, 0, 220, 219, 1, 0, 0, 0, 220, 221, 1, 0, 0, 0, 221,
		239, 1, 0, 0, 0, 222, 224, 3, 28, 14, 0, 223, 225, 5, 1, 0, 0, 224, 223,
		1, 0, 0, 0, 224, 225, 1, 0, 0, 0, 225, 239, 1, 0, 0, 0, 226, 228, 3, 94,
		47, 0, 227, 229, 5, 1, 0, 0, 228, 227, 1, 0, 0, 0, 228, 229, 1, 0, 0, 0,
		229, 239, 1, 0, 0, 0, 230, 232, 3, 30, 15, 0, 231, 233, 5, 1, 0, 0, 232,
		231, 1, 0, 0, 0, 232, 233, 1, 0, 0, 0, 233, 239, 1, 0, 0, 0, 234, 236,
		3, 42, 21, 0, 235, 237, 5, 1, 0, 0, 236, 235, 1, 0, 0, 0, 236, 237, 1,
		0, 0, 0, 237, 239, 1, 0, 0, 0, 238, 210, 1, 0, 0, 0, 238, 214, 1, 0, 0,
		0, 238, 218, 1, 0, 0, 0, 238, 222, 1, 0, 0, 0, 238, 226, 1, 0, 0, 0, 238,
		230, 1, 0, 0, 0, 238, 234, 1, 0, 0, 0, 239, 21, 1, 0, 0, 0, 240, 244, 3,
		66, 33, 0, 241, 244, 3, 68, 34, 0, 242, 244, 3, 74, 37, 0, 243, 240, 1,
		0, 0, 0, 243, 241, 1, 0, 0, 0, 243, 242, 1, 0, 0, 0, 244, 23, 1, 0, 0,
		0, 245, 248, 3, 44, 22, 0, 246, 248, 3, 48, 24, 0, 247, 245, 1, 0, 0, 0,
		247, 246, 1, 0, 0, 0, 248, 25, 1, 0, 0, 0, 249, 253, 3, 50, 25, 0, 250,
		253, 3, 52, 26, 0, 251, 253, 3, 54, 27, 0, 252, 249, 1, 0, 0, 0, 252, 250,
		1, 0, 0, 0, 252, 251, 1, 0, 0, 0, 253, 27, 1, 0, 0, 0, 254, 258, 3, 60,
		30, 0, 255, 258, 3, 62, 31, 0, 256, 258, 3, 64, 32, 0, 257, 254, 1, 0,
		0, 0, 257, 255, 1, 0, 0, 0, 257, 256, 1, 0, 0, 0, 258, 29, 1, 0, 0, 0,
		259, 269, 3, 80, 40, 0, 260, 269, 3, 82, 41, 0, 261, 269, 3, 84, 42, 0,
		262, 269, 3, 86, 43, 0, 263, 269, 3, 88, 44, 0, 264, 269, 3, 90, 45, 0,
		265, 269, 3, 92, 46, 0, 266, 269, 3, 32, 16, 0, 267, 269, 3, 38, 19, 0,
		268, 259, 1, 0, 0, 0, 268, 260, 1, 0, 0, 0, 268, 261, 1, 0, 0, 0, 268,
		262, 1, 0, 0, 0, 268, 263, 1, 0, 0, 0, 268, 264, 1, 0, 0, 0, 268, 265,
		1, 0, 0, 0, 268, 266, 1, 0, 0, 0, 268, 267, 1, 0, 0, 0, 269, 31, 1, 0,
		0, 0, 270, 271, 5, 71, 0, 0, 271, 273, 5, 3, 0, 0, 272, 274, 3, 34, 17,
		0, 273, 272, 1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 275, 1, 0, 0, 0, 275,
		276, 5, 4, 0, 0, 276, 33, 1, 0, 0, 0, 277, 282, 3, 36, 18, 0, 278, 279,
		5, 6, 0, 0, 279, 281, 3, 36, 18, 0, 280, 278, 1, 0, 0, 0, 281, 284, 1,
		0, 0, 0, 282, 280, 1, 0, 0, 0, 282, 283, 1, 0, 0, 0, 283, 35, 1, 0, 0,
		0, 284, 282, 1, 0, 0, 0, 285, 286, 5, 71, 0, 0, 286, 288, 5, 7, 0, 0, 287,
		285, 1, 0, 0, 0, 287, 288, 1, 0, 0, 0, 288, 290, 1, 0, 0, 0, 289, 291,
		5, 16, 0, 0, 290, 289, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0, 291, 292, 1, 0,
		0, 0, 292, 293, 3, 96, 48, 0, 293, 37, 1, 0, 0, 0, 294, 295, 5, 71, 0,
		0, 295, 296, 5, 17, 0, 0, 296, 297, 5, 71, 0, 0, 297, 299, 5, 3, 0, 0,
		298, 300, 3, 34, 17, 0, 299, 298, 1, 0, 0, 0, 299, 300, 1, 0, 0, 0, 300,
		301, 1, 0, 0, 0, 301, 302, 5, 4, 0, 0, 302, 39, 1, 0, 0, 0, 303, 304, 5,
		71, 0, 0, 304, 305, 5, 17, 0, 0, 305, 317, 5, 18, 0, 0, 306, 307, 5, 71,
		0, 0, 307, 308, 5, 17, 0, 0, 308, 317, 5, 19, 0, 0, 309, 312, 5, 71, 0,
		0, 310, 311, 5, 17, 0, 0, 311, 313, 5, 71, 0, 0, 312, 310, 1, 0, 0, 0,
		313, 314, 1, 0, 0, 0, 314, 312, 1, 0, 0, 0, 314, 315, 1, 0, 0, 0, 315,
		317, 1, 0, 0, 0, 316, 303, 1, 0, 0, 0, 316, 306, 1, 0, 0, 0, 316, 309,
		1, 0, 0, 0, 317, 41, 1, 0, 0, 0, 318, 321, 5, 71, 0, 0, 319, 320, 5, 17,
		0, 0, 320, 322, 5, 71, 0, 0, 321, 319, 1, 0, 0, 0, 322, 323, 1, 0, 0, 0,
		323, 321, 1, 0, 0, 0, 323, 324, 1, 0, 0, 0, 324, 325, 1, 0, 0, 0, 325,
		326, 5, 14, 0, 0, 326, 346, 3, 96, 48, 0, 327, 330, 5, 71, 0, 0, 328, 329,
		5, 17, 0, 0, 329, 331, 5, 71, 0, 0, 330, 328, 1, 0, 0, 0, 331, 332, 1,
		0, 0, 0, 332, 330, 1, 0, 0, 0, 332, 333, 1, 0, 0, 0, 333, 334, 1, 0, 0,
		0, 334, 335, 5, 20, 0, 0, 335, 346, 3, 96, 48, 0, 336, 339, 5, 71, 0, 0,
		337, 338, 5, 17, 0, 0, 338, 340, 5, 71, 0, 0, 339, 337, 1, 0, 0, 0, 340,
		341, 1, 0, 0, 0, 341, 339, 1, 0, 0, 0, 341, 342, 1, 0, 0, 0, 342, 343,
		1, 0, 0, 0, 343, 344, 5, 21, 0, 0, 344, 346, 3, 96, 48, 0, 345, 318, 1,
		0, 0, 0, 345, 327, 1, 0, 0, 0, 345, 336, 1, 0, 0, 0, 346, 43, 1, 0, 0,
		0, 347, 348, 5, 22, 0, 0, 348, 349, 5, 71, 0, 0, 349, 352, 5, 23, 0, 0,
		350, 353, 3, 96, 48, 0, 351, 353, 3, 46, 23, 0, 352, 350, 1, 0, 0, 0, 352,
		351, 1, 0, 0, 0, 353, 354, 1, 0, 0, 0, 354, 355, 3, 12, 6, 0, 355, 45,
		1, 0, 0, 0, 356, 357, 3, 96, 48, 0, 357, 358, 5, 24, 0, 0, 358, 359, 3,
		96, 48, 0, 359, 47, 1, 0, 0, 0, 360, 361, 5, 25, 0, 0, 361, 362, 3, 96,
		48, 0, 362, 363, 3, 12, 6, 0, 363, 49, 1, 0, 0, 0, 364, 365, 5, 26, 0,
		0, 365, 366, 3, 96, 48, 0, 366, 367, 3, 12, 6, 0, 367, 381, 1, 0, 0, 0,
		368, 369, 5, 26, 0, 0, 369, 370, 3, 96, 48, 0, 370, 371, 3, 12, 6, 0, 371,
		372, 5, 27, 0, 0, 372, 373, 3, 12, 6, 0, 373, 381, 1, 0, 0, 0, 374, 375,
		5, 26, 0, 0, 375, 376, 3, 96, 48, 0, 376, 377, 3, 12, 6, 0, 377, 378, 5,
		27, 0, 0, 378, 379, 3, 50, 25, 0, 379, 381, 1, 0, 0, 0, 380, 364, 1, 0,
		0, 0, 380, 368, 1, 0, 0, 0, 380, 374, 1, 0, 0, 0, 381, 51, 1, 0, 0, 0,
		382, 383, 5, 28, 0, 0, 383, 384, 3, 96, 48, 0, 384, 385, 5, 27, 0, 0, 385,
		386, 5, 9, 0, 0, 386, 387, 3, 18, 9, 0, 387, 388, 7, 1, 0, 0, 388, 389,
		5, 10, 0, 0, 389, 53, 1, 0, 0, 0, 390, 391, 5, 32, 0, 0, 391, 392, 3, 96,
		48, 0, 392, 396, 5, 9, 0, 0, 393, 395, 3, 56, 28, 0, 394, 393, 1, 0, 0,
		0, 395, 398, 1, 0, 0, 0, 396, 394, 1, 0, 0, 0, 396, 397, 1, 0, 0, 0, 397,
		400, 1, 0, 0, 0, 398, 396, 1, 0, 0, 0, 399, 401, 3, 58, 29, 0, 400, 399,
		1, 0, 0, 0, 400, 401, 1, 0, 0, 0, 401, 402, 1, 0, 0, 0, 402, 403, 5, 10,
		0, 0, 403, 55, 1, 0, 0, 0, 404, 405, 5, 33, 0, 0, 405, 406, 3, 96, 48,
		0, 406, 407, 5, 7, 0, 0, 407, 409, 3, 18, 9, 0, 408, 410, 5, 30, 0, 0,
		409, 408, 1, 0, 0, 0, 409, 410, 1, 0, 0, 0, 410, 57, 1, 0, 0, 0, 411, 412,
		5, 34, 0, 0, 412, 413, 5, 7, 0, 0, 413, 415, 3, 18, 9, 0, 414, 416, 5,
		30, 0, 0, 415, 414, 1, 0, 0, 0, 415, 416, 1, 0, 0, 0, 416, 59, 1, 0, 0,
		0, 417, 418, 5, 30, 0, 0, 418, 61, 1, 0, 0, 0, 419, 420, 5, 29, 0, 0, 420,
		63, 1, 0, 0, 0, 421, 423, 5, 31, 0, 0, 422, 424, 3, 96, 48, 0, 423, 422,
		1, 0, 0, 0, 423, 424, 1, 0, 0, 0, 424, 65, 1, 0, 0, 0, 425, 426, 5, 12,
		0, 0, 426, 428, 5, 71, 0, 0, 427, 429, 3, 70, 35, 0, 428, 427, 1, 0, 0,
		0, 428, 429, 1, 0, 0, 0, 429, 430, 1, 0, 0, 0, 430, 431, 5, 14, 0, 0, 431,
		432, 3, 96, 48, 0, 432, 67, 1, 0, 0, 0, 433, 434, 5, 13, 0, 0, 434, 435,
		5, 71, 0, 0, 435, 436, 3, 70, 35, 0, 436, 437, 5, 35, 0, 0, 437, 446, 1,
		0, 0, 0, 438, 439, 5, 13, 0, 0, 439, 441, 5, 71, 0, 0, 440, 442, 3, 70,
		35, 0, 441, 440, 1, 0, 0, 0, 441, 442, 1, 0, 0, 0, 442, 443, 1, 0, 0, 0,
		443, 444, 5, 14, 0, 0, 444, 446, 3, 96, 48, 0, 445, 433, 1, 0, 0, 0, 445,
		438, 1, 0, 0, 0, 446, 69, 1, 0, 0, 0, 447, 448, 5, 7, 0, 0, 448, 449, 3,
		72, 36, 0, 449, 71, 1, 0, 0, 0, 450, 451, 7, 2, 0, 0, 451, 73, 1, 0, 0,
		0, 452, 453, 5, 13, 0, 0, 453, 454, 5, 71, 0, 0, 454, 455, 5, 7, 0, 0,
		455, 456, 5, 41, 0, 0, 456, 457, 3, 72, 36, 0, 457, 458, 5, 42, 0, 0, 458,
		459, 3, 76, 38, 0, 459, 75, 1, 0, 0, 0, 460, 461, 5, 14, 0, 0, 461, 462,
		5, 41, 0, 0, 462, 463, 3, 78, 39, 0, 463, 464, 5, 42, 0, 0, 464, 471, 1,
		0, 0, 0, 465, 466, 5, 14, 0, 0, 466, 467, 5, 41, 0, 0, 467, 471, 5, 42,
		0, 0, 468, 469, 5, 14, 0, 0, 469, 471, 5, 71, 0, 0, 470, 460, 1, 0, 0,
		0, 470, 465, 1, 0, 0, 0, 470, 468, 1, 0, 0, 0, 471, 77, 1, 0, 0, 0, 472,
		477, 3, 96, 48, 0, 473, 474, 5, 6, 0, 0, 474, 476, 3, 96, 48, 0, 475, 473,
		1, 0, 0, 0, 476, 479, 1, 0, 0, 0, 477, 475, 1, 0, 0, 0, 477, 478, 1, 0,
		0, 0, 478, 79, 1, 0, 0, 0, 479, 477, 1, 0, 0, 0, 480, 481, 5, 43, 0, 0,
		481, 482, 5, 3, 0, 0, 482, 483, 3, 78, 39, 0, 483, 484, 5, 4, 0, 0, 484,
		81, 1, 0, 0, 0, 485, 486, 5, 71, 0, 0, 486, 487, 5, 17, 0, 0, 487, 488,
		5, 44, 0, 0, 488, 489, 5, 3, 0, 0, 489, 490, 3, 96, 48, 0, 490, 491, 5,
		4, 0, 0, 491, 83, 1, 0, 0, 0, 492, 493, 5, 71, 0, 0, 493, 494, 5, 17, 0,
		0, 494, 495, 5, 45, 0, 0, 495, 496, 5, 3, 0, 0, 496, 497, 5, 4, 0, 0, 497,
		85, 1, 0, 0, 0, 498, 499, 5, 71, 0, 0, 499, 500, 5, 17, 0, 0, 500, 501,
		5, 46, 0, 0, 501, 502, 5, 3, 0, 0, 502, 503, 5, 47, 0, 0, 503, 504, 5,
		7, 0, 0, 504, 505, 3, 96, 48, 0, 505, 506, 5, 4, 0, 0, 506, 87, 1, 0, 0,
		0, 507, 508, 5, 37, 0, 0, 508, 509, 5, 3, 0, 0, 509, 510, 3, 96, 48, 0,
		510, 511, 5, 4, 0, 0, 511, 89, 1, 0, 0, 0, 512, 513, 5, 38, 0, 0, 513,
		514, 5, 3, 0, 0, 514, 515, 3, 96, 48, 0, 515, 516, 5, 4, 0, 0, 516, 91,
		1, 0, 0, 0, 517, 518, 5, 36, 0, 0, 518, 519, 5, 3, 0, 0, 519, 520, 3, 96,
		48, 0, 520, 521, 5, 4, 0, 0, 521, 93, 1, 0, 0, 0, 522, 523, 5, 71, 0, 0,
		523, 524, 5, 14, 0, 0, 524, 553, 3, 96, 48, 0, 525, 526, 5, 71, 0, 0, 526,
		527, 5, 20, 0, 0, 527, 553, 3, 96, 48, 0, 528, 529, 5, 71, 0, 0, 529, 530,
		5, 21, 0, 0, 530, 553, 3, 96, 48, 0, 531, 532, 5, 71, 0, 0, 532, 533, 5,
		41, 0, 0, 533, 534, 3, 96, 48, 0, 534, 535, 5, 42, 0, 0, 535, 536, 5, 14,
		0, 0, 536, 537, 3, 96, 48, 0, 537, 553, 1, 0, 0, 0, 538, 539, 5, 71, 0,
		0, 539, 540, 5, 41, 0, 0, 540, 541, 3, 96, 48, 0, 541, 542, 5, 42, 0, 0,
		542, 543, 5, 20, 0, 0, 543, 544, 3, 96, 48, 0, 544, 553, 1, 0, 0, 0, 545,
		546, 5, 71, 0, 0, 546, 547, 5, 41, 0, 0, 547, 548, 3, 96, 48, 0, 548, 549,
		5, 42, 0, 0, 549, 550, 5, 21, 0, 0, 550, 551, 3, 96, 48, 0, 551, 553, 1,
		0, 0, 0, 552, 522, 1, 0, 0, 0, 552, 525, 1, 0, 0, 0, 552, 528, 1, 0, 0,
		0, 552, 531, 1, 0, 0, 0, 552, 538, 1, 0, 0, 0, 552, 545, 1, 0, 0, 0, 553,
		95, 1, 0, 0, 0, 554, 555, 6, 48, -1, 0, 555, 578, 3, 100, 50, 0, 556, 578,
		3, 40, 20, 0, 557, 558, 5, 71, 0, 0, 558, 559, 5, 41, 0, 0, 559, 560, 3,
		96, 48, 0, 560, 561, 5, 42, 0, 0, 561, 578, 1, 0, 0, 0, 562, 578, 3, 30,
		15, 0, 563, 564, 5, 71, 0, 0, 564, 565, 5, 3, 0, 0, 565, 566, 3, 98, 49,
		0, 566, 567, 5, 4, 0, 0, 567, 578, 1, 0, 0, 0, 568, 578, 5, 71, 0, 0, 569,
		570, 5, 3, 0, 0, 570, 571, 3, 96, 48, 0, 571, 572, 5, 4, 0, 0, 572, 578,
		1, 0, 0, 0, 573, 574, 5, 48, 0, 0, 574, 578, 3, 96, 48, 6, 575, 576, 5,
		49, 0, 0, 576, 578, 3, 96, 48, 5, 577, 554, 1, 0, 0, 0, 577, 556, 1, 0,
		0, 0, 577, 557, 1, 0, 0, 0, 577, 562, 1, 0, 0, 0, 577, 563, 1, 0, 0, 0,
		577, 568, 1, 0, 0, 0, 577, 569, 1, 0, 0, 0, 577, 573, 1, 0, 0, 0, 577,
		575, 1, 0, 0, 0, 578, 593, 1, 0, 0, 0, 579, 580, 10, 4, 0, 0, 580, 581,
		7, 3, 0, 0, 581, 592, 3, 96, 48, 5, 582, 583, 10, 3, 0, 0, 583, 584, 7,
		4, 0, 0, 584, 592, 3, 96, 48, 4, 585, 586, 10, 2, 0, 0, 586, 587, 7, 5,
		0, 0, 587, 592, 3, 96, 48, 3, 588, 589, 10, 1, 0, 0, 589, 590, 7, 6, 0,
		0, 590, 592, 3, 96, 48, 2, 591, 579, 1, 0, 0, 0, 591, 582, 1, 0, 0, 0,
		591, 585, 1, 0, 0, 0, 591, 588, 1, 0, 0, 0, 592, 595, 1, 0, 0, 0, 593,
		591, 1, 0, 0, 0, 593, 594, 1, 0, 0, 0, 594, 97, 1, 0, 0, 0, 595, 593, 1,
		0, 0, 0, 596, 597, 5, 71, 0, 0, 597, 598, 5, 7, 0, 0, 598, 605, 3, 96,
		48, 0, 599, 600, 5, 6, 0, 0, 600, 601, 5, 71, 0, 0, 601, 602, 5, 7, 0,
		0, 602, 604, 3, 96, 48, 0, 603, 599, 1, 0, 0, 0, 604, 607, 1, 0, 0, 0,
		605, 603, 1, 0, 0, 0, 605, 606, 1, 0, 0, 0, 606, 99, 1, 0, 0, 0, 607, 605,
		1, 0, 0, 0, 608, 615, 5, 68, 0, 0, 609, 615, 5, 66, 0, 0, 610, 615, 5,
		67, 0, 0, 611, 615, 5, 69, 0, 0, 612, 615, 5, 70, 0, 0, 613, 615, 5, 62,
		0, 0, 614, 608, 1, 0, 0, 0, 614, 609, 1, 0, 0, 0, 614, 610, 1, 0, 0, 0,
		614, 611, 1, 0, 0, 0, 614, 612, 1, 0, 0, 0, 614, 613, 1, 0, 0, 0, 615,
		101, 1, 0, 0, 0, 62, 108, 113, 117, 121, 125, 129, 135, 137, 143, 148,
		157, 161, 166, 180, 189, 193, 196, 199, 202, 207, 212, 216, 220, 224, 228,
		232, 236, 238, 243, 247, 252, 257, 268, 273, 282, 287, 290, 299, 314, 316,
		323, 332, 341, 345, 352, 380, 396, 400, 409, 415, 423, 428, 441, 445, 470,
		477, 552, 577, 591, 593, 605, 614,
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
	T_swiftParserT__59         = 60
	T_swiftParserT__60         = 61
	T_swiftParserT__61         = 62
	T_swiftParserWS            = 63
	T_swiftParserBlock_comment = 64
	T_swiftParserLine_comment  = 65
	T_swiftParserInt           = 66
	T_swiftParserFloat         = 67
	T_swiftParserCaracter      = 68
	T_swiftParserString_       = 69
	T_swiftParserBool          = 70
	T_swiftParserIdentificador = 71
)

// T_swiftParser rules.
const (
	T_swiftParserRULE_inicio                     = 0
	T_swiftParserRULE_instrucciones_globales     = 1
	T_swiftParserRULE_intruccion_global          = 2
	T_swiftParserRULE_function_declaracion       = 3
	T_swiftParserRULE_lista_parametros           = 4
	T_swiftParserRULE_declaracion_parametro      = 5
	T_swiftParserRULE_code_block                 = 6
	T_swiftParserRULE_struct_declaracion         = 7
	T_swiftParserRULE_lista_atributos            = 8
	T_swiftParserRULE_instrucciones              = 9
	T_swiftParserRULE_instruccion                = 10
	T_swiftParserRULE_declaracion                = 11
	T_swiftParserRULE_loop_statement             = 12
	T_swiftParserRULE_branch_statement           = 13
	T_swiftParserRULE_control_transfer_statement = 14
	T_swiftParserRULE_llamadas_funciones         = 15
	T_swiftParserRULE_llamada_normal             = 16
	T_swiftParserRULE_lista_argumentos           = 17
	T_swiftParserRULE_declaracion_argumento      = 18
	T_swiftParserRULE_llamada_metodos            = 19
	T_swiftParserRULE_atributos                  = 20
	T_swiftParserRULE_asignar_atributos          = 21
	T_swiftParserRULE_for_in_statement           = 22
	T_swiftParserRULE_rango                      = 23
	T_swiftParserRULE_while_statement            = 24
	T_swiftParserRULE_if_statement               = 25
	T_swiftParserRULE_guard_statement            = 26
	T_swiftParserRULE_switch_statement           = 27
	T_swiftParserRULE_switch_case                = 28
	T_swiftParserRULE_default_case               = 29
	T_swiftParserRULE_break_statement            = 30
	T_swiftParserRULE_continue_statement         = 31
	T_swiftParserRULE_return_statement           = 32
	T_swiftParserRULE_constant_declaracion       = 33
	T_swiftParserRULE_variable_declaracion       = 34
	T_swiftParserRULE_anotacion_tipo             = 35
	T_swiftParserRULE_tipos                      = 36
	T_swiftParserRULE_array_declaracion          = 37
	T_swiftParserRULE_definicion_vector          = 38
	T_swiftParserRULE_lista_expresiones          = 39
	T_swiftParserRULE_funcion_print              = 40
	T_swiftParserRULE_funcion_append             = 41
	T_swiftParserRULE_funcion_removeLast         = 42
	T_swiftParserRULE_funcion_removeat           = 43
	T_swiftParserRULE_funcion_int                = 44
	T_swiftParserRULE_funcion_float              = 45
	T_swiftParserRULE_funcion_string             = 46
	T_swiftParserRULE_asignacion                 = 47
	T_swiftParserRULE_expresion                  = 48
	T_swiftParserRULE_l_duble                    = 49
	T_swiftParserRULE_primitivos                 = 50
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
		p.SetState(102)
		p.Instrucciones_globales()
	}
	{
		p.SetState(103)
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
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9281797634052) != 0) || _la == T_swiftParserIdentificador {
		{
			p.SetState(105)
			p.Intruccion_global()
		}

		p.SetState(110)
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
	Struct_declaracion() IStruct_declaracionContext
	Asignar_atributos() IAsignar_atributosContext

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

func (s *Intruccion_globalContext) Struct_declaracion() IStruct_declaracionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_declaracionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStruct_declaracionContext)
}

func (s *Intruccion_globalContext) Asignar_atributos() IAsignar_atributosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignar_atributosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignar_atributosContext)
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

	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(111)
			p.Declaracion()
		}
		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(112)
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
			p.SetState(115)
			p.Loop_statement()
		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(116)
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
			p.SetState(119)
			p.Branch_statement()
		}
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(120)
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
			p.SetState(123)
			p.Asignacion()
		}
		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(124)
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
			p.SetState(127)
			p.Llamadas_funciones()
		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(128)
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
			p.SetState(131)
			p.Function_declaracion()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(132)
			p.Struct_declaracion()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(133)
			p.Asignar_atributos()
		}
		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(134)
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
		p.SetState(139)
		p.Match(T_swiftParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(140)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(141)
		p.Match(T_swiftParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T_swiftParserIdentificador {
		{
			p.SetState(142)
			p.Lista_parametros()
		}

	}
	{
		p.SetState(145)
		p.Match(T_swiftParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T_swiftParserT__4 {
		{
			p.SetState(146)
			p.Match(T_swiftParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(147)
			p.Tipos()
		}

	}
	{
		p.SetState(150)
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
	AllDeclaracion_parametro() []IDeclaracion_parametroContext
	Declaracion_parametro(i int) IDeclaracion_parametroContext

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

func (s *Lista_parametrosContext) AllDeclaracion_parametro() []IDeclaracion_parametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclaracion_parametroContext); ok {
			len++
		}
	}

	tst := make([]IDeclaracion_parametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclaracion_parametroContext); ok {
			tst[i] = t.(IDeclaracion_parametroContext)
			i++
		}
	}

	return tst
}

func (s *Lista_parametrosContext) Declaracion_parametro(i int) IDeclaracion_parametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracion_parametroContext); ok {
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

	return t.(IDeclaracion_parametroContext)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		p.Declaracion_parametro()
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T_swiftParserT__5 {
		{
			p.SetState(153)
			p.Match(T_swiftParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
			p.Declaracion_parametro()
		}

		p.SetState(159)
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

// IDeclaracion_parametroContext is an interface to support dynamic dispatch.
type IDeclaracion_parametroContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRefencia returns the refencia token.
	GetRefencia() antlr.Token

	// SetRefencia sets the refencia token.
	SetRefencia(antlr.Token)

	// Getter signatures
	AllIdentificador() []antlr.TerminalNode
	Identificador(i int) antlr.TerminalNode
	Tipos() ITiposContext

	// IsDeclaracion_parametroContext differentiates from other interfaces.
	IsDeclaracion_parametroContext()
}

type Declaracion_parametroContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	refencia antlr.Token
}

func NewEmptyDeclaracion_parametroContext() *Declaracion_parametroContext {
	var p = new(Declaracion_parametroContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_declaracion_parametro
	return p
}

func InitEmptyDeclaracion_parametroContext(p *Declaracion_parametroContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_declaracion_parametro
}

func (*Declaracion_parametroContext) IsDeclaracion_parametroContext() {}

func NewDeclaracion_parametroContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declaracion_parametroContext {
	var p = new(Declaracion_parametroContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_declaracion_parametro

	return p
}

func (s *Declaracion_parametroContext) GetParser() antlr.Parser { return s.parser }

func (s *Declaracion_parametroContext) GetRefencia() antlr.Token { return s.refencia }

func (s *Declaracion_parametroContext) SetRefencia(v antlr.Token) { s.refencia = v }

func (s *Declaracion_parametroContext) AllIdentificador() []antlr.TerminalNode {
	return s.GetTokens(T_swiftParserIdentificador)
}

func (s *Declaracion_parametroContext) Identificador(i int) antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, i)
}

func (s *Declaracion_parametroContext) Tipos() ITiposContext {
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

func (s *Declaracion_parametroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracion_parametroContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Declaracion_parametroContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitDeclaracion_parametro(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Declaracion_parametro() (localctx IDeclaracion_parametroContext) {
	localctx = NewDeclaracion_parametroContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, T_swiftParserRULE_declaracion_parametro)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(161)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(160)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(163)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(164)
		p.Match(T_swiftParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T_swiftParserT__7 {
		{
			p.SetState(165)

			var _m = p.Match(T_swiftParserT__7)

			localctx.(*Declaracion_parametroContext).refencia = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(168)
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
	p.EnterRule(localctx, 12, T_swiftParserRULE_code_block)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.Match(T_swiftParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(171)
		p.Instrucciones()
	}
	{
		p.SetState(172)
		p.Match(T_swiftParserT__9)
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

// IStruct_declaracionContext is an interface to support dynamic dispatch.
type IStruct_declaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identificador() antlr.TerminalNode
	AllLista_atributos() []ILista_atributosContext
	Lista_atributos(i int) ILista_atributosContext

	// IsStruct_declaracionContext differentiates from other interfaces.
	IsStruct_declaracionContext()
}

type Struct_declaracionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_declaracionContext() *Struct_declaracionContext {
	var p = new(Struct_declaracionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_struct_declaracion
	return p
}

func InitEmptyStruct_declaracionContext(p *Struct_declaracionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_struct_declaracion
}

func (*Struct_declaracionContext) IsStruct_declaracionContext() {}

func NewStruct_declaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_declaracionContext {
	var p = new(Struct_declaracionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_struct_declaracion

	return p
}

func (s *Struct_declaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_declaracionContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Struct_declaracionContext) AllLista_atributos() []ILista_atributosContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILista_atributosContext); ok {
			len++
		}
	}

	tst := make([]ILista_atributosContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILista_atributosContext); ok {
			tst[i] = t.(ILista_atributosContext)
			i++
		}
	}

	return tst
}

func (s *Struct_declaracionContext) Lista_atributos(i int) ILista_atributosContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_atributosContext); ok {
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

	return t.(ILista_atributosContext)
}

func (s *Struct_declaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_declaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Struct_declaracionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitStruct_declaracion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Struct_declaracion() (localctx IStruct_declaracionContext) {
	localctx = NewStruct_declaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, T_swiftParserRULE_struct_declaracion)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		p.Match(T_swiftParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(175)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(176)
		p.Match(T_swiftParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&45060) != 0 {
		{
			p.SetState(177)
			p.Lista_atributos()
		}

		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(183)
		p.Match(T_swiftParserT__9)
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

// ILista_atributosContext is an interface to support dynamic dispatch.
type ILista_atributosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLista_atributosContext differentiates from other interfaces.
	IsLista_atributosContext()
}

type Lista_atributosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLista_atributosContext() *Lista_atributosContext {
	var p = new(Lista_atributosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_lista_atributos
	return p
}

func InitEmptyLista_atributosContext(p *Lista_atributosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_lista_atributos
}

func (*Lista_atributosContext) IsLista_atributosContext() {}

func NewLista_atributosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lista_atributosContext {
	var p = new(Lista_atributosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_lista_atributos

	return p
}

func (s *Lista_atributosContext) GetParser() antlr.Parser { return s.parser }

func (s *Lista_atributosContext) CopyAll(ctx *Lista_atributosContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Lista_atributosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lista_atributosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Declarar_atributoContext struct {
	Lista_atributosContext
	tipo antlr.Token
}

func NewDeclarar_atributoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Declarar_atributoContext {
	var p = new(Declarar_atributoContext)

	InitEmptyLista_atributosContext(&p.Lista_atributosContext)
	p.parser = parser
	p.CopyAll(ctx.(*Lista_atributosContext))

	return p
}

func (s *Declarar_atributoContext) GetTipo() antlr.Token { return s.tipo }

func (s *Declarar_atributoContext) SetTipo(v antlr.Token) { s.tipo = v }

func (s *Declarar_atributoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declarar_atributoContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Declarar_atributoContext) Tipos() ITiposContext {
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

func (s *Declarar_atributoContext) Expresion() IExpresionContext {
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

func (s *Declarar_atributoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitDeclarar_atributo(s)

	default:
		return t.VisitChildren(s)
	}
}

type Declarar_funcion_scContext struct {
	Lista_atributosContext
	m antlr.Token
}

func NewDeclarar_funcion_scContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Declarar_funcion_scContext {
	var p = new(Declarar_funcion_scContext)

	InitEmptyLista_atributosContext(&p.Lista_atributosContext)
	p.parser = parser
	p.CopyAll(ctx.(*Lista_atributosContext))

	return p
}

func (s *Declarar_funcion_scContext) GetM() antlr.Token { return s.m }

func (s *Declarar_funcion_scContext) SetM(v antlr.Token) { s.m = v }

func (s *Declarar_funcion_scContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declarar_funcion_scContext) Function_declaracion() IFunction_declaracionContext {
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

func (s *Declarar_funcion_scContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitDeclarar_funcion_sc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Lista_atributos() (localctx ILista_atributosContext) {
	localctx = NewLista_atributosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, T_swiftParserRULE_lista_atributos)
	var _la int

	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T_swiftParserT__11, T_swiftParserT__12:
		localctx = NewDeclarar_atributoContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(185)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Declarar_atributoContext).tipo = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == T_swiftParserT__11 || _la == T_swiftParserT__12) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Declarar_atributoContext).tipo = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(186)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(189)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__6 {
			{
				p.SetState(187)
				p.Match(T_swiftParserT__6)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(188)
				p.Tipos()
			}

		}
		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__13 {
			{
				p.SetState(191)
				p.Match(T_swiftParserT__13)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(192)
				p.expresion(0)
			}

		}
		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(195)
				p.Match(T_swiftParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case T_swiftParserT__1, T_swiftParserT__14:
		localctx = NewDeclarar_funcion_scContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__14 {
			{
				p.SetState(198)

				var _m = p.Match(T_swiftParserT__14)

				localctx.(*Declarar_funcion_scContext).m = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(201)
			p.Function_declaracion()
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
	p.EnterRule(localctx, 18, T_swiftParserRULE_instrucciones)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(204)
				p.Instruccion()
			}

		}
		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
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
	Asignar_atributos() IAsignar_atributosContext

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

func (s *InstruccionContext) Asignar_atributos() IAsignar_atributosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignar_atributosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignar_atributosContext)
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
	p.EnterRule(localctx, 20, T_swiftParserRULE_instruccion)
	var _la int

	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(210)
			p.Declaracion()
		}
		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(211)
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
			p.SetState(214)
			p.Loop_statement()
		}
		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(215)
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
			p.SetState(218)
			p.Branch_statement()
		}
		p.SetState(220)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(219)
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
			p.SetState(222)
			p.Control_transfer_statement()
		}
		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(223)
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
			p.SetState(226)
			p.Asignacion()
		}
		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(227)
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
			p.SetState(230)
			p.Llamadas_funciones()
		}
		p.SetState(232)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(231)
				p.Match(T_swiftParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(234)
			p.Asignar_atributos()
		}
		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(235)
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
	p.EnterRule(localctx, 22, T_swiftParserRULE_declaracion)
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(240)
			p.Constant_declaracion()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(241)
			p.Variable_declaracion()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(242)
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
	p.EnterRule(localctx, 24, T_swiftParserRULE_loop_statement)
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T_swiftParserT__21:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(245)
			p.For_in_statement()
		}

	case T_swiftParserT__24:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(246)
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
	p.EnterRule(localctx, 26, T_swiftParserRULE_branch_statement)
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T_swiftParserT__25:
		localctx = NewDeclarar_ifContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(249)
			p.If_statement()
		}

	case T_swiftParserT__27:
		localctx = NewDeclarar_guardContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(250)
			p.Guard_statement()
		}

	case T_swiftParserT__31:
		localctx = NewDeclarar_switchContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(251)
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
	p.EnterRule(localctx, 28, T_swiftParserRULE_control_transfer_statement)
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T_swiftParserT__29:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(254)
			p.Break_statement()
		}

	case T_swiftParserT__28:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(255)
			p.Continue_statement()
		}

	case T_swiftParserT__30:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(256)
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
	Llamada_normal() ILlamada_normalContext
	Llamada_metodos() ILlamada_metodosContext

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

func (s *Llamadas_funcionesContext) Llamada_normal() ILlamada_normalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamada_normalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamada_normalContext)
}

func (s *Llamadas_funcionesContext) Llamada_metodos() ILlamada_metodosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamada_metodosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamada_metodosContext)
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
	p.EnterRule(localctx, 30, T_swiftParserRULE_llamadas_funciones)
	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(259)
			p.Funcion_print()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(260)
			p.Funcion_append()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(261)
			p.Funcion_removeLast()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(262)
			p.Funcion_removeat()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(263)
			p.Funcion_int()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(264)
			p.Funcion_float()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(265)
			p.Funcion_string()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(266)
			p.Llamada_normal()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(267)
			p.Llamada_metodos()
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

// ILlamada_normalContext is an interface to support dynamic dispatch.
type ILlamada_normalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identificador() antlr.TerminalNode
	Lista_argumentos() ILista_argumentosContext

	// IsLlamada_normalContext differentiates from other interfaces.
	IsLlamada_normalContext()
}

type Llamada_normalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLlamada_normalContext() *Llamada_normalContext {
	var p = new(Llamada_normalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_llamada_normal
	return p
}

func InitEmptyLlamada_normalContext(p *Llamada_normalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_llamada_normal
}

func (*Llamada_normalContext) IsLlamada_normalContext() {}

func NewLlamada_normalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Llamada_normalContext {
	var p = new(Llamada_normalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_llamada_normal

	return p
}

func (s *Llamada_normalContext) GetParser() antlr.Parser { return s.parser }

func (s *Llamada_normalContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Llamada_normalContext) Lista_argumentos() ILista_argumentosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_argumentosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_argumentosContext)
}

func (s *Llamada_normalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Llamada_normalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Llamada_normalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitLlamada_normal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Llamada_normal() (localctx ILlamada_normalContext) {
	localctx = NewLlamada_normalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, T_swiftParserRULE_llamada_normal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(270)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(271)
		p.Match(T_swiftParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4612539720486944776) != 0) || ((int64((_la-66)) & ^0x3f) == 0 && ((int64(1)<<(_la-66))&63) != 0) {
		{
			p.SetState(272)
			p.Lista_argumentos()
		}

	}
	{
		p.SetState(275)
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

// ILista_argumentosContext is an interface to support dynamic dispatch.
type ILista_argumentosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDeclaracion_argumento() []IDeclaracion_argumentoContext
	Declaracion_argumento(i int) IDeclaracion_argumentoContext

	// IsLista_argumentosContext differentiates from other interfaces.
	IsLista_argumentosContext()
}

type Lista_argumentosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLista_argumentosContext() *Lista_argumentosContext {
	var p = new(Lista_argumentosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_lista_argumentos
	return p
}

func InitEmptyLista_argumentosContext(p *Lista_argumentosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_lista_argumentos
}

func (*Lista_argumentosContext) IsLista_argumentosContext() {}

func NewLista_argumentosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lista_argumentosContext {
	var p = new(Lista_argumentosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_lista_argumentos

	return p
}

func (s *Lista_argumentosContext) GetParser() antlr.Parser { return s.parser }

func (s *Lista_argumentosContext) AllDeclaracion_argumento() []IDeclaracion_argumentoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclaracion_argumentoContext); ok {
			len++
		}
	}

	tst := make([]IDeclaracion_argumentoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclaracion_argumentoContext); ok {
			tst[i] = t.(IDeclaracion_argumentoContext)
			i++
		}
	}

	return tst
}

func (s *Lista_argumentosContext) Declaracion_argumento(i int) IDeclaracion_argumentoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracion_argumentoContext); ok {
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

	return t.(IDeclaracion_argumentoContext)
}

func (s *Lista_argumentosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lista_argumentosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Lista_argumentosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitLista_argumentos(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Lista_argumentos() (localctx ILista_argumentosContext) {
	localctx = NewLista_argumentosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, T_swiftParserRULE_lista_argumentos)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(277)
		p.Declaracion_argumento()
	}
	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T_swiftParserT__5 {
		{
			p.SetState(278)
			p.Match(T_swiftParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(279)
			p.Declaracion_argumento()
		}

		p.SetState(284)
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

// IDeclaracion_argumentoContext is an interface to support dynamic dispatch.
type IDeclaracion_argumentoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetR returns the r token.
	GetR() antlr.Token

	// SetR sets the r token.
	SetR(antlr.Token)

	// Getter signatures
	Expresion() IExpresionContext
	Identificador() antlr.TerminalNode

	// IsDeclaracion_argumentoContext differentiates from other interfaces.
	IsDeclaracion_argumentoContext()
}

type Declaracion_argumentoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	r      antlr.Token
}

func NewEmptyDeclaracion_argumentoContext() *Declaracion_argumentoContext {
	var p = new(Declaracion_argumentoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_declaracion_argumento
	return p
}

func InitEmptyDeclaracion_argumentoContext(p *Declaracion_argumentoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_declaracion_argumento
}

func (*Declaracion_argumentoContext) IsDeclaracion_argumentoContext() {}

func NewDeclaracion_argumentoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declaracion_argumentoContext {
	var p = new(Declaracion_argumentoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_declaracion_argumento

	return p
}

func (s *Declaracion_argumentoContext) GetParser() antlr.Parser { return s.parser }

func (s *Declaracion_argumentoContext) GetR() antlr.Token { return s.r }

func (s *Declaracion_argumentoContext) SetR(v antlr.Token) { s.r = v }

func (s *Declaracion_argumentoContext) Expresion() IExpresionContext {
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

func (s *Declaracion_argumentoContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Declaracion_argumentoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracion_argumentoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Declaracion_argumentoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitDeclaracion_argumento(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Declaracion_argumento() (localctx IDeclaracion_argumentoContext) {
	localctx = NewDeclaracion_argumentoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, T_swiftParserRULE_declaracion_argumento)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(287)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(285)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(286)
			p.Match(T_swiftParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T_swiftParserT__15 {
		{
			p.SetState(289)

			var _m = p.Match(T_swiftParserT__15)

			localctx.(*Declaracion_argumentoContext).r = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(292)
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

// ILlamada_metodosContext is an interface to support dynamic dispatch.
type ILlamada_metodosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentificador() []antlr.TerminalNode
	Identificador(i int) antlr.TerminalNode
	Lista_argumentos() ILista_argumentosContext

	// IsLlamada_metodosContext differentiates from other interfaces.
	IsLlamada_metodosContext()
}

type Llamada_metodosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLlamada_metodosContext() *Llamada_metodosContext {
	var p = new(Llamada_metodosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_llamada_metodos
	return p
}

func InitEmptyLlamada_metodosContext(p *Llamada_metodosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_llamada_metodos
}

func (*Llamada_metodosContext) IsLlamada_metodosContext() {}

func NewLlamada_metodosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Llamada_metodosContext {
	var p = new(Llamada_metodosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_llamada_metodos

	return p
}

func (s *Llamada_metodosContext) GetParser() antlr.Parser { return s.parser }

func (s *Llamada_metodosContext) AllIdentificador() []antlr.TerminalNode {
	return s.GetTokens(T_swiftParserIdentificador)
}

func (s *Llamada_metodosContext) Identificador(i int) antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, i)
}

func (s *Llamada_metodosContext) Lista_argumentos() ILista_argumentosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_argumentosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_argumentosContext)
}

func (s *Llamada_metodosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Llamada_metodosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Llamada_metodosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitLlamada_metodos(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Llamada_metodos() (localctx ILlamada_metodosContext) {
	localctx = NewLlamada_metodosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, T_swiftParserRULE_llamada_metodos)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(294)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(295)
		p.Match(T_swiftParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(296)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(297)
		p.Match(T_swiftParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4612539720486944776) != 0) || ((int64((_la-66)) & ^0x3f) == 0 && ((int64(1)<<(_la-66))&63) != 0) {
		{
			p.SetState(298)
			p.Lista_argumentos()
		}

	}
	{
		p.SetState(301)
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

type Atributos_generalesContext struct {
	AtributosContext
}

func NewAtributos_generalesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Atributos_generalesContext {
	var p = new(Atributos_generalesContext)

	InitEmptyAtributosContext(&p.AtributosContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtributosContext))

	return p
}

func (s *Atributos_generalesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Atributos_generalesContext) AllIdentificador() []antlr.TerminalNode {
	return s.GetTokens(T_swiftParserIdentificador)
}

func (s *Atributos_generalesContext) Identificador(i int) antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, i)
}

func (s *Atributos_generalesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitAtributos_generales(s)

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
	p.EnterRule(localctx, 40, T_swiftParserRULE_atributos)
	var _alt int

	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAtributos_vector_emptyContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(303)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(304)
			p.Match(T_swiftParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(305)
			p.Match(T_swiftParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewAtributos_vector_countContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(306)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(307)
			p.Match(T_swiftParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(308)
			p.Match(T_swiftParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewAtributos_generalesContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(309)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(312)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(310)
					p.Match(T_swiftParserT__16)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(311)
					p.Match(T_swiftParserIdentificador)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(314)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext())
			if p.HasError() {
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

// IAsignar_atributosContext is an interface to support dynamic dispatch.
type IAsignar_atributosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAsignar_atributosContext differentiates from other interfaces.
	IsAsignar_atributosContext()
}

type Asignar_atributosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsignar_atributosContext() *Asignar_atributosContext {
	var p = new(Asignar_atributosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_asignar_atributos
	return p
}

func InitEmptyAsignar_atributosContext(p *Asignar_atributosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_asignar_atributos
}

func (*Asignar_atributosContext) IsAsignar_atributosContext() {}

func NewAsignar_atributosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Asignar_atributosContext {
	var p = new(Asignar_atributosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_asignar_atributos

	return p
}

func (s *Asignar_atributosContext) GetParser() antlr.Parser { return s.parser }

func (s *Asignar_atributosContext) CopyAll(ctx *Asignar_atributosContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Asignar_atributosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Asignar_atributosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Incre_atributos_normalContext struct {
	Asignar_atributosContext
}

func NewIncre_atributos_normalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Incre_atributos_normalContext {
	var p = new(Incre_atributos_normalContext)

	InitEmptyAsignar_atributosContext(&p.Asignar_atributosContext)
	p.parser = parser
	p.CopyAll(ctx.(*Asignar_atributosContext))

	return p
}

func (s *Incre_atributos_normalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Incre_atributos_normalContext) AllIdentificador() []antlr.TerminalNode {
	return s.GetTokens(T_swiftParserIdentificador)
}

func (s *Incre_atributos_normalContext) Identificador(i int) antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, i)
}

func (s *Incre_atributos_normalContext) Expresion() IExpresionContext {
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

func (s *Incre_atributos_normalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitIncre_atributos_normal(s)

	default:
		return t.VisitChildren(s)
	}
}

type Decre_atributos_normalContext struct {
	Asignar_atributosContext
}

func NewDecre_atributos_normalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Decre_atributos_normalContext {
	var p = new(Decre_atributos_normalContext)

	InitEmptyAsignar_atributosContext(&p.Asignar_atributosContext)
	p.parser = parser
	p.CopyAll(ctx.(*Asignar_atributosContext))

	return p
}

func (s *Decre_atributos_normalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Decre_atributos_normalContext) AllIdentificador() []antlr.TerminalNode {
	return s.GetTokens(T_swiftParserIdentificador)
}

func (s *Decre_atributos_normalContext) Identificador(i int) antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, i)
}

func (s *Decre_atributos_normalContext) Expresion() IExpresionContext {
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

func (s *Decre_atributos_normalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitDecre_atributos_normal(s)

	default:
		return t.VisitChildren(s)
	}
}

type Asignar_atributos_normalContext struct {
	Asignar_atributosContext
}

func NewAsignar_atributos_normalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Asignar_atributos_normalContext {
	var p = new(Asignar_atributos_normalContext)

	InitEmptyAsignar_atributosContext(&p.Asignar_atributosContext)
	p.parser = parser
	p.CopyAll(ctx.(*Asignar_atributosContext))

	return p
}

func (s *Asignar_atributos_normalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Asignar_atributos_normalContext) AllIdentificador() []antlr.TerminalNode {
	return s.GetTokens(T_swiftParserIdentificador)
}

func (s *Asignar_atributos_normalContext) Identificador(i int) antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, i)
}

func (s *Asignar_atributos_normalContext) Expresion() IExpresionContext {
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

func (s *Asignar_atributos_normalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitAsignar_atributos_normal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Asignar_atributos() (localctx IAsignar_atributosContext) {
	localctx = NewAsignar_atributosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, T_swiftParserRULE_asignar_atributos)
	var _la int

	p.SetState(345)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAsignar_atributos_normalContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(318)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(321)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == T_swiftParserT__16 {
			{
				p.SetState(319)
				p.Match(T_swiftParserT__16)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(320)
				p.Match(T_swiftParserIdentificador)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(323)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(325)
			p.Match(T_swiftParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(326)
			p.expresion(0)
		}

	case 2:
		localctx = NewIncre_atributos_normalContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(327)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(330)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == T_swiftParserT__16 {
			{
				p.SetState(328)
				p.Match(T_swiftParserT__16)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(329)
				p.Match(T_swiftParserIdentificador)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(332)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(334)
			p.Match(T_swiftParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(335)
			p.expresion(0)
		}

	case 3:
		localctx = NewDecre_atributos_normalContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(336)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(339)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == T_swiftParserT__16 {
			{
				p.SetState(337)
				p.Match(T_swiftParserT__16)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(338)
				p.Match(T_swiftParserIdentificador)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(341)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(343)
			p.Match(T_swiftParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(344)
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
	p.EnterRule(localctx, 44, T_swiftParserRULE_for_in_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(347)
		p.Match(T_swiftParserT__21)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(348)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(349)
		p.Match(T_swiftParserT__22)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(352)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(350)
			p.expresion(0)
		}

	case 2:
		{
			p.SetState(351)
			p.Rango()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(354)
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
	p.EnterRule(localctx, 46, T_swiftParserRULE_rango)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(356)
		p.expresion(0)
	}
	{
		p.SetState(357)
		p.Match(T_swiftParserT__23)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(358)
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
	p.EnterRule(localctx, 48, T_swiftParserRULE_while_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(360)
		p.Match(T_swiftParserT__24)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(361)
		p.expresion(0)
	}
	{
		p.SetState(362)
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
	p.EnterRule(localctx, 50, T_swiftParserRULE_if_statement)
	p.SetState(380)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIf_simpleContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(364)
			p.Match(T_swiftParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(365)
			p.expresion(0)
		}
		{
			p.SetState(366)
			p.Code_block()
		}

	case 2:
		localctx = NewElse_finalContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(368)
			p.Match(T_swiftParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(369)
			p.expresion(0)
		}
		{
			p.SetState(370)
			p.Code_block()
		}
		{
			p.SetState(371)
			p.Match(T_swiftParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(372)
			p.Code_block()
		}

	case 3:
		localctx = NewSiguiente_ifContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(374)
			p.Match(T_swiftParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(375)
			p.expresion(0)
		}
		{
			p.SetState(376)
			p.Code_block()
		}
		{
			p.SetState(377)
			p.Match(T_swiftParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(378)
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
	p.EnterRule(localctx, 52, T_swiftParserRULE_guard_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(382)
		p.Match(T_swiftParserT__27)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(383)
		p.expresion(0)
	}
	{
		p.SetState(384)
		p.Match(T_swiftParserT__26)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(385)
		p.Match(T_swiftParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(386)
		p.Instrucciones()
	}
	{
		p.SetState(387)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3758096384) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(388)
		p.Match(T_swiftParserT__9)
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
	p.EnterRule(localctx, 54, T_swiftParserRULE_switch_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(390)
		p.Match(T_swiftParserT__31)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(391)
		p.expresion(0)
	}
	{
		p.SetState(392)
		p.Match(T_swiftParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(396)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T_swiftParserT__32 {
		{
			p.SetState(393)
			p.Switch_case()
		}

		p.SetState(398)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(400)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T_swiftParserT__33 {
		{
			p.SetState(399)
			p.Default_case()
		}

	}
	{
		p.SetState(402)
		p.Match(T_swiftParserT__9)
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
	p.EnterRule(localctx, 56, T_swiftParserRULE_switch_case)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(404)
		p.Match(T_swiftParserT__32)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(405)
		p.expresion(0)
	}
	{
		p.SetState(406)
		p.Match(T_swiftParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(407)
		p.Instrucciones()
	}
	p.SetState(409)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T_swiftParserT__29 {
		{
			p.SetState(408)
			p.Match(T_swiftParserT__29)
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
	p.EnterRule(localctx, 58, T_swiftParserRULE_default_case)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(411)
		p.Match(T_swiftParserT__33)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(412)
		p.Match(T_swiftParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(413)
		p.Instrucciones()
	}
	p.SetState(415)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T_swiftParserT__29 {
		{
			p.SetState(414)
			p.Match(T_swiftParserT__29)
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
	p.EnterRule(localctx, 60, T_swiftParserRULE_break_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(417)
		p.Match(T_swiftParserT__29)
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
	p.EnterRule(localctx, 62, T_swiftParserRULE_continue_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(419)
		p.Match(T_swiftParserT__28)
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
	p.EnterRule(localctx, 64, T_swiftParserRULE_return_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(421)
		p.Match(T_swiftParserT__30)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(423)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 50, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(422)
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
	p.EnterRule(localctx, 66, T_swiftParserRULE_constant_declaracion)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(425)
		p.Match(T_swiftParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(426)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(428)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T_swiftParserT__6 {
		{
			p.SetState(427)
			p.Anotacion_tipo()
		}

	}
	{
		p.SetState(430)
		p.Match(T_swiftParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(431)
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
	p.EnterRule(localctx, 68, T_swiftParserRULE_variable_declaracion)
	var _la int

	p.SetState(445)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 53, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(433)
			p.Match(T_swiftParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(434)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(435)
			p.Anotacion_tipo()
		}
		{
			p.SetState(436)
			p.Match(T_swiftParserT__34)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(438)
			p.Match(T_swiftParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(439)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(441)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__6 {
			{
				p.SetState(440)
				p.Anotacion_tipo()
			}

		}
		{
			p.SetState(443)
			p.Match(T_swiftParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(444)
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
	p.EnterRule(localctx, 70, T_swiftParserRULE_anotacion_tipo)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(447)
		p.Match(T_swiftParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(448)
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

	// Getter signatures
	Identificador() antlr.TerminalNode

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

func (s *TiposContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

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
	p.EnterRule(localctx, 72, T_swiftParserRULE_tipos)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(450)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-36)) & ^0x3f) == 0 && ((int64(1)<<(_la-36))&34359738399) != 0) {
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
	p.EnterRule(localctx, 74, T_swiftParserRULE_array_declaracion)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(452)
		p.Match(T_swiftParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(453)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(454)
		p.Match(T_swiftParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(455)
		p.Match(T_swiftParserT__40)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(456)
		p.Tipos()
	}
	{
		p.SetState(457)
		p.Match(T_swiftParserT__41)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(458)
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
	p.EnterRule(localctx, 76, T_swiftParserRULE_definicion_vector)
	p.SetState(470)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 54, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(460)
			p.Match(T_swiftParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(461)
			p.Match(T_swiftParserT__40)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(462)
			p.Lista_expresiones()
		}
		{
			p.SetState(463)
			p.Match(T_swiftParserT__41)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(465)
			p.Match(T_swiftParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(466)
			p.Match(T_swiftParserT__40)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(467)
			p.Match(T_swiftParserT__41)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(468)
			p.Match(T_swiftParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(469)
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
	p.EnterRule(localctx, 78, T_swiftParserRULE_lista_expresiones)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(472)
		p.expresion(0)
	}
	p.SetState(477)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T_swiftParserT__5 {
		{
			p.SetState(473)
			p.Match(T_swiftParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(474)
			p.expresion(0)
		}

		p.SetState(479)
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
	p.EnterRule(localctx, 80, T_swiftParserRULE_funcion_print)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(480)
		p.Match(T_swiftParserT__42)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(481)
		p.Match(T_swiftParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(482)
		p.Lista_expresiones()
	}
	{
		p.SetState(483)
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
	p.EnterRule(localctx, 82, T_swiftParserRULE_funcion_append)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(485)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(486)
		p.Match(T_swiftParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(487)
		p.Match(T_swiftParserT__43)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(488)
		p.Match(T_swiftParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(489)
		p.expresion(0)
	}
	{
		p.SetState(490)
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
	p.EnterRule(localctx, 84, T_swiftParserRULE_funcion_removeLast)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(492)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(493)
		p.Match(T_swiftParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(494)
		p.Match(T_swiftParserT__44)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(495)
		p.Match(T_swiftParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(496)
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
	p.EnterRule(localctx, 86, T_swiftParserRULE_funcion_removeat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(498)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(499)
		p.Match(T_swiftParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(500)
		p.Match(T_swiftParserT__45)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(501)
		p.Match(T_swiftParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(502)
		p.Match(T_swiftParserT__46)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(503)
		p.Match(T_swiftParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(504)
		p.expresion(0)
	}
	{
		p.SetState(505)
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
	p.EnterRule(localctx, 88, T_swiftParserRULE_funcion_int)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(507)
		p.Match(T_swiftParserT__36)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(508)
		p.Match(T_swiftParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(509)
		p.expresion(0)
	}
	{
		p.SetState(510)
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
	p.EnterRule(localctx, 90, T_swiftParserRULE_funcion_float)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(512)
		p.Match(T_swiftParserT__37)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(513)
		p.Match(T_swiftParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(514)
		p.expresion(0)
	}
	{
		p.SetState(515)
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
	p.EnterRule(localctx, 92, T_swiftParserRULE_funcion_string)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(517)
		p.Match(T_swiftParserT__35)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(518)
		p.Match(T_swiftParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(519)
		p.expresion(0)
	}
	{
		p.SetState(520)
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
	p.EnterRule(localctx, 94, T_swiftParserRULE_asignacion)
	p.SetState(552)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 56, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAsignacion_normalContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(522)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(523)
			p.Match(T_swiftParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(524)
			p.expresion(0)
		}

	case 2:
		localctx = NewAsignacion_incrementoContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(525)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(526)
			p.Match(T_swiftParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(527)
			p.expresion(0)
		}

	case 3:
		localctx = NewAsignacion_decrementoContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(528)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(529)
			p.Match(T_swiftParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(530)
			p.expresion(0)
		}

	case 4:
		localctx = NewAsignacion_vectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(531)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(532)
			p.Match(T_swiftParserT__40)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(533)
			p.expresion(0)
		}
		{
			p.SetState(534)
			p.Match(T_swiftParserT__41)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(535)
			p.Match(T_swiftParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(536)
			p.expresion(0)
		}

	case 5:
		localctx = NewAsignacion_incremento_vectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(538)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(539)
			p.Match(T_swiftParserT__40)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(540)
			p.expresion(0)
		}
		{
			p.SetState(541)
			p.Match(T_swiftParserT__41)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(542)
			p.Match(T_swiftParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(543)
			p.expresion(0)
		}

	case 6:
		localctx = NewAsignacion_decremento_vectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(545)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(546)
			p.Match(T_swiftParserT__40)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(547)
			p.expresion(0)
		}
		{
			p.SetState(548)
			p.Match(T_swiftParserT__41)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(549)
			p.Match(T_swiftParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(550)
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

type Expresion_struct_duplaContext struct {
	ExpresionContext
}

func NewExpresion_struct_duplaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expresion_struct_duplaContext {
	var p = new(Expresion_struct_duplaContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *Expresion_struct_duplaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expresion_struct_duplaContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Expresion_struct_duplaContext) L_duble() IL_dubleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IL_dubleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IL_dubleContext)
}

func (s *Expresion_struct_duplaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitExpresion_struct_dupla(s)

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
	_startState := 96
	p.EnterRecursionRule(localctx, 96, T_swiftParserRULE_expresion, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(577)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 57, p.GetParserRuleContext()) {
	case 1:
		localctx = NewValor_primitivoContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(555)
			p.Primitivos()
		}

	case 2:
		localctx = NewExpresion_atributosContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(556)
			p.Atributos()
		}

	case 3:
		localctx = NewExpresion_vectorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(557)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(558)
			p.Match(T_swiftParserT__40)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(559)
			p.expresion(0)
		}
		{
			p.SetState(560)
			p.Match(T_swiftParserT__41)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewExpresion_llamadaContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(562)
			p.Llamadas_funciones()
		}

	case 5:
		localctx = NewExpresion_struct_duplaContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(563)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(564)
			p.Match(T_swiftParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(565)
			p.L_duble()
		}
		{
			p.SetState(566)
			p.Match(T_swiftParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewExpresion_idContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(568)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewExpresion_parenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(569)
			p.Match(T_swiftParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(570)
			p.expresion(0)
		}
		{
			p.SetState(571)
			p.Match(T_swiftParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewExpresion_negaContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(573)

			var _m = p.Match(T_swiftParserT__47)

			localctx.(*Expresion_negaContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(574)
			p.expresion(6)
		}

	case 9:
		localctx = NewExpresion_unarioContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(575)
			p.Match(T_swiftParserT__48)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(576)
			p.expresion(5)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(593)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 59, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(591)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 58, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresion_aritContext(p, NewExpresionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, T_swiftParserRULE_expresion)
				p.SetState(579)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(580)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expresion_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7881299347898368) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expresion_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(581)
					p.expresion(5)
				}

			case 2:
				localctx = NewExpresion_aritContext(p, NewExpresionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, T_swiftParserRULE_expresion)
				p.SetState(582)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(583)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expresion_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == T_swiftParserT__48 || _la == T_swiftParserT__52) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expresion_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(584)
					p.expresion(4)
				}

			case 3:
				localctx = NewExpresion_compaContext(p, NewExpresionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, T_swiftParserRULE_expresion)
				p.SetState(585)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(586)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expresion_compaContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1134907106097364992) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expresion_compaContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(587)
					p.expresion(3)
				}

			case 4:
				localctx = NewExpresion_relaContext(p, NewExpresionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, T_swiftParserRULE_expresion)
				p.SetState(588)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(589)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expresion_relaContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == T_swiftParserT__59 || _la == T_swiftParserT__60) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expresion_relaContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(590)
					p.expresion(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(595)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 59, p.GetParserRuleContext())
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

// IL_dubleContext is an interface to support dynamic dispatch.
type IL_dubleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentificador() []antlr.TerminalNode
	Identificador(i int) antlr.TerminalNode
	AllExpresion() []IExpresionContext
	Expresion(i int) IExpresionContext

	// IsL_dubleContext differentiates from other interfaces.
	IsL_dubleContext()
}

type L_dubleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyL_dubleContext() *L_dubleContext {
	var p = new(L_dubleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_l_duble
	return p
}

func InitEmptyL_dubleContext(p *L_dubleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_l_duble
}

func (*L_dubleContext) IsL_dubleContext() {}

func NewL_dubleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *L_dubleContext {
	var p = new(L_dubleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_l_duble

	return p
}

func (s *L_dubleContext) GetParser() antlr.Parser { return s.parser }

func (s *L_dubleContext) AllIdentificador() []antlr.TerminalNode {
	return s.GetTokens(T_swiftParserIdentificador)
}

func (s *L_dubleContext) Identificador(i int) antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, i)
}

func (s *L_dubleContext) AllExpresion() []IExpresionContext {
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

func (s *L_dubleContext) Expresion(i int) IExpresionContext {
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

func (s *L_dubleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *L_dubleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *L_dubleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitL_duble(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) L_duble() (localctx IL_dubleContext) {
	localctx = NewL_dubleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, T_swiftParserRULE_l_duble)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(596)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(597)
		p.Match(T_swiftParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(598)
		p.expresion(0)
	}
	p.SetState(605)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T_swiftParserT__5 {
		{
			p.SetState(599)
			p.Match(T_swiftParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(600)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(601)
			p.Match(T_swiftParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(602)
			p.expresion(0)
		}

		p.SetState(607)
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

type Primitivo_nuloContext struct {
	PrimitivosContext
}

func NewPrimitivo_nuloContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Primitivo_nuloContext {
	var p = new(Primitivo_nuloContext)

	InitEmptyPrimitivosContext(&p.PrimitivosContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimitivosContext))

	return p
}

func (s *Primitivo_nuloContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primitivo_nuloContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitPrimitivo_nulo(s)

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
	p.EnterRule(localctx, 100, T_swiftParserRULE_primitivos)
	p.SetState(614)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T_swiftParserCaracter:
		localctx = NewPrimitivo_caracterContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(608)
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
			p.SetState(609)
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
			p.SetState(610)
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
			p.SetState(611)
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
			p.SetState(612)
			p.Match(T_swiftParserBool)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case T_swiftParserT__61:
		localctx = NewPrimitivo_nuloContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(613)
			p.Match(T_swiftParserT__61)
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
	case 48:
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
