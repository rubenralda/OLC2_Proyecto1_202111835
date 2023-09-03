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
		"funcion_append", "funcion_removeLast", "funcion_removeat", "asignacion",
		"expresion", "primitivos",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 68, 449, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 1, 0, 1, 0, 1, 0, 1, 1, 5, 1, 83, 8, 1, 10,
		1, 12, 1, 86, 9, 1, 1, 2, 1, 2, 3, 2, 90, 8, 2, 1, 2, 1, 2, 3, 2, 94, 8,
		2, 1, 2, 1, 2, 3, 2, 98, 8, 2, 1, 2, 1, 2, 3, 2, 102, 8, 2, 1, 2, 1, 2,
		3, 2, 106, 8, 2, 1, 2, 3, 2, 109, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 115,
		8, 3, 1, 3, 1, 3, 1, 3, 3, 3, 120, 8, 3, 1, 3, 1, 3, 1, 4, 1, 4, 3, 4,
		126, 8, 4, 1, 4, 1, 4, 3, 4, 130, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 136,
		8, 4, 1, 4, 1, 4, 3, 4, 140, 8, 4, 1, 4, 3, 4, 143, 8, 4, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 6, 5, 6, 150, 8, 6, 10, 6, 12, 6, 153, 9, 6, 1, 7, 1, 7, 3,
		7, 157, 8, 7, 1, 7, 1, 7, 3, 7, 161, 8, 7, 1, 7, 1, 7, 3, 7, 165, 8, 7,
		1, 7, 1, 7, 3, 7, 169, 8, 7, 1, 7, 1, 7, 3, 7, 173, 8, 7, 1, 7, 1, 7, 3,
		7, 177, 8, 7, 3, 7, 179, 8, 7, 1, 8, 1, 8, 1, 8, 3, 8, 184, 8, 8, 1, 9,
		1, 9, 3, 9, 188, 8, 9, 1, 10, 1, 10, 1, 10, 3, 10, 193, 8, 10, 1, 11, 1,
		11, 1, 11, 3, 11, 198, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 204, 8,
		12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 212, 8, 13, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 219, 8, 14, 1, 14, 1, 14, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 3, 17, 247, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 261, 8, 19, 10, 19, 12,
		19, 264, 9, 19, 1, 19, 3, 19, 267, 8, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 3, 20, 276, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21,
		282, 8, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 3, 24, 290, 8, 24,
		1, 25, 1, 25, 1, 25, 3, 25, 295, 8, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 308, 8, 26, 1, 26,
		1, 26, 3, 26, 312, 8, 26, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 337, 8, 30, 1, 31, 1,
		31, 1, 31, 5, 31, 342, 8, 31, 10, 31, 12, 31, 345, 9, 31, 1, 32, 1, 32,
		1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 1, 36, 3, 36, 404, 8, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37,
		1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 37, 3, 37, 423, 8, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37,
		1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 5, 37, 437, 8, 37, 10, 37, 12,
		37, 440, 9, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 3, 38, 447, 8, 38, 1,
		38, 0, 1, 74, 39, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
		30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64,
		66, 68, 70, 72, 74, 76, 0, 7, 2, 0, 7, 7, 68, 68, 1, 0, 22, 24, 1, 0, 32,
		36, 1, 0, 48, 50, 2, 0, 47, 47, 51, 51, 1, 0, 52, 57, 1, 0, 58, 59, 483,
		0, 78, 1, 0, 0, 0, 2, 84, 1, 0, 0, 0, 4, 108, 1, 0, 0, 0, 6, 110, 1, 0,
		0, 0, 8, 142, 1, 0, 0, 0, 10, 144, 1, 0, 0, 0, 12, 151, 1, 0, 0, 0, 14,
		178, 1, 0, 0, 0, 16, 183, 1, 0, 0, 0, 18, 187, 1, 0, 0, 0, 20, 192, 1,
		0, 0, 0, 22, 197, 1, 0, 0, 0, 24, 203, 1, 0, 0, 0, 26, 211, 1, 0, 0, 0,
		28, 213, 1, 0, 0, 0, 30, 222, 1, 0, 0, 0, 32, 226, 1, 0, 0, 0, 34, 246,
		1, 0, 0, 0, 36, 248, 1, 0, 0, 0, 38, 256, 1, 0, 0, 0, 40, 270, 1, 0, 0,
		0, 42, 277, 1, 0, 0, 0, 44, 283, 1, 0, 0, 0, 46, 285, 1, 0, 0, 0, 48, 287,
		1, 0, 0, 0, 50, 291, 1, 0, 0, 0, 52, 311, 1, 0, 0, 0, 54, 313, 1, 0, 0,
		0, 56, 316, 1, 0, 0, 0, 58, 318, 1, 0, 0, 0, 60, 336, 1, 0, 0, 0, 62, 338,
		1, 0, 0, 0, 64, 346, 1, 0, 0, 0, 66, 351, 1, 0, 0, 0, 68, 358, 1, 0, 0,
		0, 70, 364, 1, 0, 0, 0, 72, 403, 1, 0, 0, 0, 74, 422, 1, 0, 0, 0, 76, 446,
		1, 0, 0, 0, 78, 79, 3, 2, 1, 0, 79, 80, 5, 0, 0, 1, 80, 1, 1, 0, 0, 0,
		81, 83, 3, 4, 2, 0, 82, 81, 1, 0, 0, 0, 83, 86, 1, 0, 0, 0, 84, 82, 1,
		0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 3, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 87,
		89, 3, 16, 8, 0, 88, 90, 5, 1, 0, 0, 89, 88, 1, 0, 0, 0, 89, 90, 1, 0,
		0, 0, 90, 109, 1, 0, 0, 0, 91, 93, 3, 18, 9, 0, 92, 94, 5, 1, 0, 0, 93,
		92, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 109, 1, 0, 0, 0, 95, 97, 3, 20,
		10, 0, 96, 98, 5, 1, 0, 0, 97, 96, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98,
		109, 1, 0, 0, 0, 99, 101, 3, 72, 36, 0, 100, 102, 5, 1, 0, 0, 101, 100,
		1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 109, 1, 0, 0, 0, 103, 105, 3, 24,
		12, 0, 104, 106, 5, 1, 0, 0, 105, 104, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0,
		106, 109, 1, 0, 0, 0, 107, 109, 3, 6, 3, 0, 108, 87, 1, 0, 0, 0, 108, 91,
		1, 0, 0, 0, 108, 95, 1, 0, 0, 0, 108, 99, 1, 0, 0, 0, 108, 103, 1, 0, 0,
		0, 108, 107, 1, 0, 0, 0, 109, 5, 1, 0, 0, 0, 110, 111, 5, 2, 0, 0, 111,
		112, 5, 68, 0, 0, 112, 114, 5, 3, 0, 0, 113, 115, 3, 8, 4, 0, 114, 113,
		1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 119, 5, 4,
		0, 0, 117, 118, 5, 5, 0, 0, 118, 120, 3, 56, 28, 0, 119, 117, 1, 0, 0,
		0, 119, 120, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 122, 3, 10, 5, 0, 122,
		7, 1, 0, 0, 0, 123, 125, 5, 6, 0, 0, 124, 126, 7, 0, 0, 0, 125, 124, 1,
		0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 129, 5, 8, 0,
		0, 128, 130, 5, 9, 0, 0, 129, 128, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130,
		131, 1, 0, 0, 0, 131, 132, 3, 56, 28, 0, 132, 133, 3, 8, 4, 0, 133, 143,
		1, 0, 0, 0, 134, 136, 7, 0, 0, 0, 135, 134, 1, 0, 0, 0, 135, 136, 1, 0,
		0, 0, 136, 137, 1, 0, 0, 0, 137, 139, 5, 8, 0, 0, 138, 140, 5, 9, 0, 0,
		139, 138, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141,
		143, 3, 56, 28, 0, 142, 123, 1, 0, 0, 0, 142, 135, 1, 0, 0, 0, 143, 9,
		1, 0, 0, 0, 144, 145, 5, 10, 0, 0, 145, 146, 3, 12, 6, 0, 146, 147, 5,
		11, 0, 0, 147, 11, 1, 0, 0, 0, 148, 150, 3, 14, 7, 0, 149, 148, 1, 0, 0,
		0, 150, 153, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152,
		13, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 154, 156, 3, 16, 8, 0, 155, 157,
		5, 1, 0, 0, 156, 155, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 179, 1, 0,
		0, 0, 158, 160, 3, 18, 9, 0, 159, 161, 5, 1, 0, 0, 160, 159, 1, 0, 0, 0,
		160, 161, 1, 0, 0, 0, 161, 179, 1, 0, 0, 0, 162, 164, 3, 20, 10, 0, 163,
		165, 5, 1, 0, 0, 164, 163, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 179,
		1, 0, 0, 0, 166, 168, 3, 22, 11, 0, 167, 169, 5, 1, 0, 0, 168, 167, 1,
		0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 179, 1, 0, 0, 0, 170, 172, 3, 72, 36,
		0, 171, 173, 5, 1, 0, 0, 172, 171, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173,
		179, 1, 0, 0, 0, 174, 176, 3, 24, 12, 0, 175, 177, 5, 1, 0, 0, 176, 175,
		1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 179, 1, 0, 0, 0, 178, 154, 1, 0,
		0, 0, 178, 158, 1, 0, 0, 0, 178, 162, 1, 0, 0, 0, 178, 166, 1, 0, 0, 0,
		178, 170, 1, 0, 0, 0, 178, 174, 1, 0, 0, 0, 179, 15, 1, 0, 0, 0, 180, 184,
		3, 50, 25, 0, 181, 184, 3, 52, 26, 0, 182, 184, 3, 58, 29, 0, 183, 180,
		1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 183, 182, 1, 0, 0, 0, 184, 17, 1, 0,
		0, 0, 185, 188, 3, 28, 14, 0, 186, 188, 3, 32, 16, 0, 187, 185, 1, 0, 0,
		0, 187, 186, 1, 0, 0, 0, 188, 19, 1, 0, 0, 0, 189, 193, 3, 34, 17, 0, 190,
		193, 3, 36, 18, 0, 191, 193, 3, 38, 19, 0, 192, 189, 1, 0, 0, 0, 192, 190,
		1, 0, 0, 0, 192, 191, 1, 0, 0, 0, 193, 21, 1, 0, 0, 0, 194, 198, 3, 44,
		22, 0, 195, 198, 3, 46, 23, 0, 196, 198, 3, 48, 24, 0, 197, 194, 1, 0,
		0, 0, 197, 195, 1, 0, 0, 0, 197, 196, 1, 0, 0, 0, 198, 23, 1, 0, 0, 0,
		199, 204, 3, 64, 32, 0, 200, 204, 3, 66, 33, 0, 201, 204, 3, 68, 34, 0,
		202, 204, 3, 70, 35, 0, 203, 199, 1, 0, 0, 0, 203, 200, 1, 0, 0, 0, 203,
		201, 1, 0, 0, 0, 203, 202, 1, 0, 0, 0, 204, 25, 1, 0, 0, 0, 205, 206, 5,
		68, 0, 0, 206, 207, 5, 12, 0, 0, 207, 212, 5, 13, 0, 0, 208, 209, 5, 68,
		0, 0, 209, 210, 5, 12, 0, 0, 210, 212, 5, 14, 0, 0, 211, 205, 1, 0, 0,
		0, 211, 208, 1, 0, 0, 0, 212, 27, 1, 0, 0, 0, 213, 214, 5, 15, 0, 0, 214,
		215, 5, 68, 0, 0, 215, 218, 5, 16, 0, 0, 216, 219, 3, 74, 37, 0, 217, 219,
		3, 30, 15, 0, 218, 216, 1, 0, 0, 0, 218, 217, 1, 0, 0, 0, 219, 220, 1,
		0, 0, 0, 220, 221, 3, 10, 5, 0, 221, 29, 1, 0, 0, 0, 222, 223, 5, 63, 0,
		0, 223, 224, 5, 17, 0, 0, 224, 225, 5, 63, 0, 0, 225, 31, 1, 0, 0, 0, 226,
		227, 5, 18, 0, 0, 227, 228, 3, 74, 37, 0, 228, 229, 3, 10, 5, 0, 229, 33,
		1, 0, 0, 0, 230, 231, 5, 19, 0, 0, 231, 232, 3, 74, 37, 0, 232, 233, 3,
		10, 5, 0, 233, 247, 1, 0, 0, 0, 234, 235, 5, 19, 0, 0, 235, 236, 3, 74,
		37, 0, 236, 237, 3, 10, 5, 0, 237, 238, 5, 20, 0, 0, 238, 239, 3, 10, 5,
		0, 239, 247, 1, 0, 0, 0, 240, 241, 5, 19, 0, 0, 241, 242, 3, 74, 37, 0,
		242, 243, 3, 10, 5, 0, 243, 244, 5, 20, 0, 0, 244, 245, 3, 34, 17, 0, 245,
		247, 1, 0, 0, 0, 246, 230, 1, 0, 0, 0, 246, 234, 1, 0, 0, 0, 246, 240,
		1, 0, 0, 0, 247, 35, 1, 0, 0, 0, 248, 249, 5, 21, 0, 0, 249, 250, 3, 74,
		37, 0, 250, 251, 5, 20, 0, 0, 251, 252, 5, 10, 0, 0, 252, 253, 3, 12, 6,
		0, 253, 254, 7, 1, 0, 0, 254, 255, 5, 11, 0, 0, 255, 37, 1, 0, 0, 0, 256,
		257, 5, 25, 0, 0, 257, 258, 3, 74, 37, 0, 258, 262, 5, 10, 0, 0, 259, 261,
		3, 40, 20, 0, 260, 259, 1, 0, 0, 0, 261, 264, 1, 0, 0, 0, 262, 260, 1,
		0, 0, 0, 262, 263, 1, 0, 0, 0, 263, 266, 1, 0, 0, 0, 264, 262, 1, 0, 0,
		0, 265, 267, 3, 42, 21, 0, 266, 265, 1, 0, 0, 0, 266, 267, 1, 0, 0, 0,
		267, 268, 1, 0, 0, 0, 268, 269, 5, 11, 0, 0, 269, 39, 1, 0, 0, 0, 270,
		271, 5, 26, 0, 0, 271, 272, 3, 74, 37, 0, 272, 273, 5, 8, 0, 0, 273, 275,
		3, 12, 6, 0, 274, 276, 5, 23, 0, 0, 275, 274, 1, 0, 0, 0, 275, 276, 1,
		0, 0, 0, 276, 41, 1, 0, 0, 0, 277, 278, 5, 27, 0, 0, 278, 279, 5, 8, 0,
		0, 279, 281, 3, 12, 6, 0, 280, 282, 5, 23, 0, 0, 281, 280, 1, 0, 0, 0,
		281, 282, 1, 0, 0, 0, 282, 43, 1, 0, 0, 0, 283, 284, 5, 23, 0, 0, 284,
		45, 1, 0, 0, 0, 285, 286, 5, 22, 0, 0, 286, 47, 1, 0, 0, 0, 287, 289, 5,
		24, 0, 0, 288, 290, 3, 74, 37, 0, 289, 288, 1, 0, 0, 0, 289, 290, 1, 0,
		0, 0, 290, 49, 1, 0, 0, 0, 291, 292, 5, 28, 0, 0, 292, 294, 5, 68, 0, 0,
		293, 295, 3, 54, 27, 0, 294, 293, 1, 0, 0, 0, 294, 295, 1, 0, 0, 0, 295,
		296, 1, 0, 0, 0, 296, 297, 5, 29, 0, 0, 297, 298, 3, 74, 37, 0, 298, 51,
		1, 0, 0, 0, 299, 300, 5, 30, 0, 0, 300, 301, 5, 68, 0, 0, 301, 302, 3,
		54, 27, 0, 302, 303, 5, 31, 0, 0, 303, 312, 1, 0, 0, 0, 304, 305, 5, 30,
		0, 0, 305, 307, 5, 68, 0, 0, 306, 308, 3, 54, 27, 0, 307, 306, 1, 0, 0,
		0, 307, 308, 1, 0, 0, 0, 308, 309, 1, 0, 0, 0, 309, 310, 5, 29, 0, 0, 310,
		312, 3, 74, 37, 0, 311, 299, 1, 0, 0, 0, 311, 304, 1, 0, 0, 0, 312, 53,
		1, 0, 0, 0, 313, 314, 5, 8, 0, 0, 314, 315, 3, 56, 28, 0, 315, 55, 1, 0,
		0, 0, 316, 317, 7, 2, 0, 0, 317, 57, 1, 0, 0, 0, 318, 319, 5, 30, 0, 0,
		319, 320, 5, 68, 0, 0, 320, 321, 5, 8, 0, 0, 321, 322, 5, 37, 0, 0, 322,
		323, 3, 56, 28, 0, 323, 324, 5, 38, 0, 0, 324, 325, 3, 60, 30, 0, 325,
		59, 1, 0, 0, 0, 326, 327, 5, 29, 0, 0, 327, 328, 5, 37, 0, 0, 328, 329,
		3, 62, 31, 0, 329, 330, 5, 38, 0, 0, 330, 337, 1, 0, 0, 0, 331, 332, 5,
		29, 0, 0, 332, 333, 5, 37, 0, 0, 333, 337, 5, 38, 0, 0, 334, 335, 5, 29,
		0, 0, 335, 337, 5, 68, 0, 0, 336, 326, 1, 0, 0, 0, 336, 331, 1, 0, 0, 0,
		336, 334, 1, 0, 0, 0, 337, 61, 1, 0, 0, 0, 338, 343, 3, 74, 37, 0, 339,
		340, 5, 6, 0, 0, 340, 342, 3, 74, 37, 0, 341, 339, 1, 0, 0, 0, 342, 345,
		1, 0, 0, 0, 343, 341, 1, 0, 0, 0, 343, 344, 1, 0, 0, 0, 344, 63, 1, 0,
		0, 0, 345, 343, 1, 0, 0, 0, 346, 347, 5, 39, 0, 0, 347, 348, 5, 3, 0, 0,
		348, 349, 3, 74, 37, 0, 349, 350, 5, 4, 0, 0, 350, 65, 1, 0, 0, 0, 351,
		352, 5, 68, 0, 0, 352, 353, 5, 12, 0, 0, 353, 354, 5, 40, 0, 0, 354, 355,
		5, 3, 0, 0, 355, 356, 3, 74, 37, 0, 356, 357, 5, 4, 0, 0, 357, 67, 1, 0,
		0, 0, 358, 359, 5, 68, 0, 0, 359, 360, 5, 12, 0, 0, 360, 361, 5, 41, 0,
		0, 361, 362, 5, 3, 0, 0, 362, 363, 5, 4, 0, 0, 363, 69, 1, 0, 0, 0, 364,
		365, 5, 68, 0, 0, 365, 366, 5, 12, 0, 0, 366, 367, 5, 42, 0, 0, 367, 368,
		5, 3, 0, 0, 368, 369, 5, 43, 0, 0, 369, 370, 5, 8, 0, 0, 370, 371, 3, 74,
		37, 0, 371, 372, 5, 4, 0, 0, 372, 71, 1, 0, 0, 0, 373, 374, 5, 68, 0, 0,
		374, 375, 5, 29, 0, 0, 375, 404, 3, 74, 37, 0, 376, 377, 5, 68, 0, 0, 377,
		378, 5, 44, 0, 0, 378, 404, 3, 74, 37, 0, 379, 380, 5, 68, 0, 0, 380, 381,
		5, 45, 0, 0, 381, 404, 3, 74, 37, 0, 382, 383, 5, 68, 0, 0, 383, 384, 5,
		37, 0, 0, 384, 385, 3, 74, 37, 0, 385, 386, 5, 38, 0, 0, 386, 387, 5, 29,
		0, 0, 387, 388, 3, 74, 37, 0, 388, 404, 1, 0, 0, 0, 389, 390, 5, 68, 0,
		0, 390, 391, 5, 37, 0, 0, 391, 392, 3, 74, 37, 0, 392, 393, 5, 38, 0, 0,
		393, 394, 5, 44, 0, 0, 394, 395, 3, 74, 37, 0, 395, 404, 1, 0, 0, 0, 396,
		397, 5, 68, 0, 0, 397, 398, 5, 37, 0, 0, 398, 399, 3, 74, 37, 0, 399, 400,
		5, 38, 0, 0, 400, 401, 5, 45, 0, 0, 401, 402, 3, 74, 37, 0, 402, 404, 1,
		0, 0, 0, 403, 373, 1, 0, 0, 0, 403, 376, 1, 0, 0, 0, 403, 379, 1, 0, 0,
		0, 403, 382, 1, 0, 0, 0, 403, 389, 1, 0, 0, 0, 403, 396, 1, 0, 0, 0, 404,
		73, 1, 0, 0, 0, 405, 406, 6, 37, -1, 0, 406, 423, 3, 76, 38, 0, 407, 423,
		3, 26, 13, 0, 408, 409, 5, 68, 0, 0, 409, 410, 5, 37, 0, 0, 410, 411, 3,
		74, 37, 0, 411, 412, 5, 38, 0, 0, 412, 423, 1, 0, 0, 0, 413, 423, 5, 68,
		0, 0, 414, 415, 5, 3, 0, 0, 415, 416, 3, 74, 37, 0, 416, 417, 5, 4, 0,
		0, 417, 423, 1, 0, 0, 0, 418, 419, 5, 46, 0, 0, 419, 423, 3, 74, 37, 6,
		420, 421, 5, 47, 0, 0, 421, 423, 3, 74, 37, 5, 422, 405, 1, 0, 0, 0, 422,
		407, 1, 0, 0, 0, 422, 408, 1, 0, 0, 0, 422, 413, 1, 0, 0, 0, 422, 414,
		1, 0, 0, 0, 422, 418, 1, 0, 0, 0, 422, 420, 1, 0, 0, 0, 423, 438, 1, 0,
		0, 0, 424, 425, 10, 4, 0, 0, 425, 426, 7, 3, 0, 0, 426, 437, 3, 74, 37,
		5, 427, 428, 10, 3, 0, 0, 428, 429, 7, 4, 0, 0, 429, 437, 3, 74, 37, 4,
		430, 431, 10, 2, 0, 0, 431, 432, 7, 5, 0, 0, 432, 437, 3, 74, 37, 3, 433,
		434, 10, 1, 0, 0, 434, 435, 7, 6, 0, 0, 435, 437, 3, 74, 37, 2, 436, 424,
		1, 0, 0, 0, 436, 427, 1, 0, 0, 0, 436, 430, 1, 0, 0, 0, 436, 433, 1, 0,
		0, 0, 437, 440, 1, 0, 0, 0, 438, 436, 1, 0, 0, 0, 438, 439, 1, 0, 0, 0,
		439, 75, 1, 0, 0, 0, 440, 438, 1, 0, 0, 0, 441, 447, 5, 65, 0, 0, 442,
		447, 5, 63, 0, 0, 443, 447, 5, 64, 0, 0, 444, 447, 5, 66, 0, 0, 445, 447,
		5, 67, 0, 0, 446, 441, 1, 0, 0, 0, 446, 442, 1, 0, 0, 0, 446, 443, 1, 0,
		0, 0, 446, 444, 1, 0, 0, 0, 446, 445, 1, 0, 0, 0, 447, 77, 1, 0, 0, 0,
		45, 84, 89, 93, 97, 101, 105, 108, 114, 119, 125, 129, 135, 139, 142, 151,
		156, 160, 164, 168, 172, 176, 178, 183, 187, 192, 197, 203, 211, 218, 246,
		262, 266, 275, 281, 289, 294, 307, 311, 336, 343, 403, 422, 436, 438, 446,
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
	T_swiftParserRULE_asignacion                 = 36
	T_swiftParserRULE_expresion                  = 37
	T_swiftParserRULE_primitivos                 = 38
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
		p.SetState(78)
		p.Instrucciones_globales()
	}
	{
		p.SetState(79)
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
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&551134461956) != 0) || _la == T_swiftParserIdentificador {
		{
			p.SetState(81)
			p.Intruccion_global()
		}

		p.SetState(86)
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

	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(87)
			p.Declaracion()
		}
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(88)
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
			p.SetState(91)
			p.Loop_statement()
		}
		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(92)
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
			p.SetState(95)
			p.Branch_statement()
		}
		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(96)
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
			p.SetState(99)
			p.Asignacion()
		}
		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(100)
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
			p.SetState(103)
			p.Llamadas_funciones()
		}
		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(104)
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
			p.SetState(107)
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
		p.SetState(110)
		p.Match(T_swiftParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(111)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(112)
		p.Match(T_swiftParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-6)) & ^0x3f) == 0 && ((int64(1)<<(_la-6))&4611686018427387911) != 0 {
		{
			p.SetState(113)
			p.Lista_parametros()
		}

	}
	{
		p.SetState(116)
		p.Match(T_swiftParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T_swiftParserT__4 {
		{
			p.SetState(117)
			p.Match(T_swiftParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(118)
			p.Tipos()
		}

	}
	{
		p.SetState(121)
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

	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T_swiftParserT__5:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(123)
			p.Match(T_swiftParserT__5)
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

		if _la == T_swiftParserT__6 || _la == T_swiftParserIdentificador {
			{
				p.SetState(124)
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
			p.SetState(127)
			p.Match(T_swiftParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__8 {
			{
				p.SetState(128)
				p.Match(T_swiftParserT__8)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(131)
			p.Tipos()
		}
		{
			p.SetState(132)
			p.Lista_parametros()
		}

	case T_swiftParserT__6, T_swiftParserT__7, T_swiftParserIdentificador:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__6 || _la == T_swiftParserIdentificador {
			{
				p.SetState(134)
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
			p.SetState(137)
			p.Match(T_swiftParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__8 {
			{
				p.SetState(138)
				p.Match(T_swiftParserT__8)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(141)
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
		p.SetState(144)
		p.Match(T_swiftParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.Instrucciones()
	}
	{
		p.SetState(146)
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
	p.SetState(151)
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
				p.SetState(148)
				p.Instruccion()
			}

		}
		p.SetState(153)
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

	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(154)
			p.Declaracion()
		}
		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(155)
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
			p.SetState(158)
			p.Loop_statement()
		}
		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(159)
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
			p.SetState(162)
			p.Branch_statement()
		}
		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(163)
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
			p.SetState(166)
			p.Control_transfer_statement()
		}
		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(167)
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
			p.SetState(170)
			p.Asignacion()
		}
		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(171)
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
			p.SetState(174)
			p.Llamadas_funciones()
		}
		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(175)
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
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(180)
			p.Constant_declaracion()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(181)
			p.Variable_declaracion()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(182)
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
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T_swiftParserT__14:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(185)
			p.For_in_statement()
		}

	case T_swiftParserT__17:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(186)
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
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T_swiftParserT__18:
		localctx = NewDeclarar_ifContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(189)
			p.If_statement()
		}

	case T_swiftParserT__20:
		localctx = NewDeclarar_guardContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(190)
			p.Guard_statement()
		}

	case T_swiftParserT__24:
		localctx = NewDeclarar_switchContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(191)
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
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T_swiftParserT__22:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(194)
			p.Break_statement()
		}

	case T_swiftParserT__21:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(195)
			p.Continue_statement()
		}

	case T_swiftParserT__23:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(196)
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
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(199)
			p.Funcion_print()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(200)
			p.Funcion_append()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(201)
			p.Funcion_removeLast()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(202)
			p.Funcion_removeat()
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
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAtributos_vector_emptyContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(205)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(206)
			p.Match(T_swiftParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(207)
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
			p.SetState(208)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(209)
			p.Match(T_swiftParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(210)
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
		p.SetState(213)
		p.Match(T_swiftParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
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
		p.Match(T_swiftParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(216)
			p.expresion(0)
		}

	case 2:
		{
			p.SetState(217)
			p.Rango()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(220)
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
	AllInt() []antlr.TerminalNode
	Int(i int) antlr.TerminalNode

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

func (s *RangoContext) AllInt() []antlr.TerminalNode {
	return s.GetTokens(T_swiftParserInt)
}

func (s *RangoContext) Int(i int) antlr.TerminalNode {
	return s.GetToken(T_swiftParserInt, i)
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
		p.SetState(222)
		p.Match(T_swiftParserInt)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(223)
		p.Match(T_swiftParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(224)
		p.Match(T_swiftParserInt)
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
		p.SetState(226)
		p.Match(T_swiftParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(227)
		p.expresion(0)
	}
	{
		p.SetState(228)
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
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIf_simpleContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(230)
			p.Match(T_swiftParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(231)
			p.expresion(0)
		}
		{
			p.SetState(232)
			p.Code_block()
		}

	case 2:
		localctx = NewElse_finalContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(234)
			p.Match(T_swiftParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(235)
			p.expresion(0)
		}
		{
			p.SetState(236)
			p.Code_block()
		}
		{
			p.SetState(237)
			p.Match(T_swiftParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(238)
			p.Code_block()
		}

	case 3:
		localctx = NewSiguiente_ifContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(240)
			p.Match(T_swiftParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(241)
			p.expresion(0)
		}
		{
			p.SetState(242)
			p.Code_block()
		}
		{
			p.SetState(243)
			p.Match(T_swiftParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(244)
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
		p.SetState(248)
		p.Match(T_swiftParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(249)
		p.expresion(0)
	}
	{
		p.SetState(250)
		p.Match(T_swiftParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(251)
		p.Match(T_swiftParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(252)
		p.Instrucciones()
	}
	{
		p.SetState(253)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&29360128) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(254)
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
		p.SetState(256)
		p.Match(T_swiftParserT__24)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(257)
		p.expresion(0)
	}
	{
		p.SetState(258)
		p.Match(T_swiftParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T_swiftParserT__25 {
		{
			p.SetState(259)
			p.Switch_case()
		}

		p.SetState(264)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T_swiftParserT__26 {
		{
			p.SetState(265)
			p.Default_case()
		}

	}
	{
		p.SetState(268)
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
		p.SetState(270)
		p.Match(T_swiftParserT__25)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(271)
		p.expresion(0)
	}
	{
		p.SetState(272)
		p.Match(T_swiftParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(273)
		p.Instrucciones()
	}
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T_swiftParserT__22 {
		{
			p.SetState(274)
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
		p.SetState(277)
		p.Match(T_swiftParserT__26)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(278)
		p.Match(T_swiftParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(279)
		p.Instrucciones()
	}
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T_swiftParserT__22 {
		{
			p.SetState(280)
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
		p.SetState(283)
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
		p.SetState(285)
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
		p.SetState(287)
		p.Match(T_swiftParserT__23)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(289)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(288)
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
		p.SetState(291)
		p.Match(T_swiftParserT__27)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(292)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T_swiftParserT__7 {
		{
			p.SetState(293)
			p.Anotacion_tipo()
		}

	}
	{
		p.SetState(296)
		p.Match(T_swiftParserT__28)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(297)
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

	p.SetState(311)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(299)
			p.Match(T_swiftParserT__29)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(300)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(301)
			p.Anotacion_tipo()
		}
		{
			p.SetState(302)
			p.Match(T_swiftParserT__30)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(304)
			p.Match(T_swiftParserT__29)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(305)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(307)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__7 {
			{
				p.SetState(306)
				p.Anotacion_tipo()
			}

		}
		{
			p.SetState(309)
			p.Match(T_swiftParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(310)
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
		p.SetState(313)
		p.Match(T_swiftParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(314)
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
		p.SetState(316)
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
		p.SetState(318)
		p.Match(T_swiftParserT__29)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(319)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(320)
		p.Match(T_swiftParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(321)
		p.Match(T_swiftParserT__36)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(322)
		p.Tipos()
	}
	{
		p.SetState(323)
		p.Match(T_swiftParserT__37)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(324)
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
	p.SetState(336)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(326)
			p.Match(T_swiftParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(327)
			p.Match(T_swiftParserT__36)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(328)
			p.Lista_expresiones()
		}
		{
			p.SetState(329)
			p.Match(T_swiftParserT__37)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(331)
			p.Match(T_swiftParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(332)
			p.Match(T_swiftParserT__36)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(333)
			p.Match(T_swiftParserT__37)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(334)
			p.Match(T_swiftParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(335)
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
		p.SetState(338)
		p.expresion(0)
	}
	p.SetState(343)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T_swiftParserT__5 {
		{
			p.SetState(339)
			p.Match(T_swiftParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(340)
			p.expresion(0)
		}

		p.SetState(345)
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
	Expresion() IExpresionContext

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

func (s *Funcion_printContext) Expresion() IExpresionContext {
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
		p.SetState(346)
		p.Match(T_swiftParserT__38)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(347)
		p.Match(T_swiftParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(348)
		p.expresion(0)
	}
	{
		p.SetState(349)
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
		p.SetState(351)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(352)
		p.Match(T_swiftParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(353)
		p.Match(T_swiftParserT__39)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(354)
		p.Match(T_swiftParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(355)
		p.expresion(0)
	}
	{
		p.SetState(356)
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
		p.SetState(358)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(359)
		p.Match(T_swiftParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(360)
		p.Match(T_swiftParserT__40)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(361)
		p.Match(T_swiftParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(362)
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
		p.SetState(364)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(365)
		p.Match(T_swiftParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(366)
		p.Match(T_swiftParserT__41)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(367)
		p.Match(T_swiftParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(368)
		p.Match(T_swiftParserT__42)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(369)
		p.Match(T_swiftParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(370)
		p.expresion(0)
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
	p.EnterRule(localctx, 72, T_swiftParserRULE_asignacion)
	p.SetState(403)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAsignacion_normalContext(p, localctx)
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
			p.Match(T_swiftParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(375)
			p.expresion(0)
		}

	case 2:
		localctx = NewAsignacion_incrementoContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(376)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(377)
			p.Match(T_swiftParserT__43)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(378)
			p.expresion(0)
		}

	case 3:
		localctx = NewAsignacion_decrementoContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(379)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(380)
			p.Match(T_swiftParserT__44)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(381)
			p.expresion(0)
		}

	case 4:
		localctx = NewAsignacion_vectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(382)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(383)
			p.Match(T_swiftParserT__36)
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
			p.Match(T_swiftParserT__37)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(386)
			p.Match(T_swiftParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(387)
			p.expresion(0)
		}

	case 5:
		localctx = NewAsignacion_incremento_vectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(389)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(390)
			p.Match(T_swiftParserT__36)
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
			p.Match(T_swiftParserT__37)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(393)
			p.Match(T_swiftParserT__43)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(394)
			p.expresion(0)
		}

	case 6:
		localctx = NewAsignacion_decremento_vectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(396)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(397)
			p.Match(T_swiftParserT__36)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(398)
			p.expresion(0)
		}
		{
			p.SetState(399)
			p.Match(T_swiftParserT__37)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(400)
			p.Match(T_swiftParserT__44)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(401)
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
	_startState := 74
	p.EnterRecursionRule(localctx, 74, T_swiftParserRULE_expresion, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(422)
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
			p.SetState(406)
			p.Primitivos()
		}

	case 2:
		localctx = NewExpresion_atributosContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(407)
			p.Atributos()
		}

	case 3:
		localctx = NewExpresion_vectorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(408)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(409)
			p.Match(T_swiftParserT__36)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(410)
			p.expresion(0)
		}
		{
			p.SetState(411)
			p.Match(T_swiftParserT__37)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewExpresion_idContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(413)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewExpresion_parenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(414)
			p.Match(T_swiftParserT__2)
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
			p.Match(T_swiftParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewExpresion_negaContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(418)

			var _m = p.Match(T_swiftParserT__45)

			localctx.(*Expresion_negaContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(419)
			p.expresion(6)
		}

	case 7:
		localctx = NewExpresion_unarioContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(420)
			p.Match(T_swiftParserT__46)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(421)
			p.expresion(5)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(438)
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
			p.SetState(436)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresion_aritContext(p, NewExpresionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, T_swiftParserRULE_expresion)
				p.SetState(424)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(425)

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
					p.SetState(426)
					p.expresion(5)
				}

			case 2:
				localctx = NewExpresion_aritContext(p, NewExpresionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, T_swiftParserRULE_expresion)
				p.SetState(427)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(428)

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
					p.SetState(429)
					p.expresion(4)
				}

			case 3:
				localctx = NewExpresion_compaContext(p, NewExpresionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, T_swiftParserRULE_expresion)
				p.SetState(430)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(431)

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
					p.SetState(432)
					p.expresion(3)
				}

			case 4:
				localctx = NewExpresion_relaContext(p, NewExpresionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, T_swiftParserRULE_expresion)
				p.SetState(433)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(434)

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
					p.SetState(435)
					p.expresion(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(440)
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
	p.EnterRule(localctx, 76, T_swiftParserRULE_primitivos)
	p.SetState(446)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T_swiftParserCaracter:
		localctx = NewPrimitivo_caracterContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(441)
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
			p.SetState(442)
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
			p.SetState(443)
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
			p.SetState(444)
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
			p.SetState(445)
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
	case 37:
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
