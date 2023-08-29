// Code generated from .\parser\T_swift.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type T_swiftLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var T_swiftLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func t_swiftlexerLexerInit() {
	staticData := &T_swiftLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "';'", "'{'", "'}'", "'for'", "'in'", "'while'", "'if'", "'else'",
		"'guard'", "'switch'", "'case'", "':'", "'default'", "'break'", "'continue'",
		"'return'", "'let '", "'='", "'var'", "'?'", "'String'", "'Int'", "'Float'",
		"'Bool'", "'['", "']'", "','", "'print'", "'('", "')'", "'-'", "'*'",
		"'/'", "'%'", "'+'", "'=='", "'!='", "'>'", "'<'", "'<='", "'>='", "'!'",
		"'&&'", "'||'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "WS", "Block_comment", "Line_comment",
		"Int", "Float", "Caracter", "String", "Bool", "Identificador",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
		"T__33", "T__34", "T__35", "T__36", "T__37", "T__38", "T__39", "T__40",
		"T__41", "T__42", "T__43", "WS", "Block_comment", "Line_comment", "Int",
		"Float", "Caracter", "String", "Quoted_text", "Quoted_text_item", "Bool",
		"Identificador", "Identifier_head", "Identifier_character", "Identifier_characters",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 53, 380, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11,
		1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30,
		1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1,
		35, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39,
		1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1,
		43, 1, 44, 4, 44, 281, 8, 44, 11, 44, 12, 44, 282, 1, 44, 1, 44, 1, 45,
		1, 45, 1, 45, 1, 45, 1, 45, 5, 45, 292, 8, 45, 10, 45, 12, 45, 295, 9,
		45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 46, 5, 46,
		306, 8, 46, 10, 46, 12, 46, 309, 9, 46, 1, 46, 3, 46, 312, 8, 46, 1, 46,
		1, 46, 1, 47, 4, 47, 317, 8, 47, 11, 47, 12, 47, 318, 1, 48, 4, 48, 322,
		8, 48, 11, 48, 12, 48, 323, 1, 48, 1, 48, 4, 48, 328, 8, 48, 11, 48, 12,
		48, 329, 3, 48, 332, 8, 48, 1, 49, 1, 49, 1, 49, 1, 49, 1, 50, 1, 50, 3,
		50, 340, 8, 50, 1, 50, 1, 50, 1, 51, 4, 51, 345, 8, 51, 11, 51, 12, 51,
		346, 1, 52, 1, 52, 1, 52, 3, 52, 352, 8, 52, 1, 53, 1, 53, 1, 53, 1, 53,
		1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 3, 53, 363, 8, 53, 1, 54, 1, 54, 3,
		54, 367, 8, 54, 1, 55, 3, 55, 370, 8, 55, 1, 56, 1, 56, 3, 56, 374, 8,
		56, 1, 57, 4, 57, 377, 8, 57, 11, 57, 12, 57, 378, 2, 293, 307, 0, 58,
		1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20,
		41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29,
		59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38,
		77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 93, 47,
		95, 48, 97, 49, 99, 50, 101, 51, 103, 0, 105, 0, 107, 52, 109, 53, 111,
		0, 113, 0, 115, 0, 1, 0, 6, 3, 0, 0, 0, 9, 13, 32, 32, 1, 1, 10, 10, 1,
		0, 48, 57, 7, 0, 34, 34, 39, 39, 48, 48, 92, 92, 110, 110, 114, 114, 116,
		116, 4, 0, 10, 10, 13, 13, 34, 34, 92, 92, 3, 0, 65, 90, 95, 95, 97, 122,
		389, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0,
		0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1,
		0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23,
		1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0,
		31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0,
		0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0,
		0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0,
		0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1,
		0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69,
		1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0,
		77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0,
		0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0,
		0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0,
		0, 0, 0, 101, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 1, 117,
		1, 0, 0, 0, 3, 119, 1, 0, 0, 0, 5, 121, 1, 0, 0, 0, 7, 123, 1, 0, 0, 0,
		9, 127, 1, 0, 0, 0, 11, 130, 1, 0, 0, 0, 13, 136, 1, 0, 0, 0, 15, 139,
		1, 0, 0, 0, 17, 144, 1, 0, 0, 0, 19, 150, 1, 0, 0, 0, 21, 157, 1, 0, 0,
		0, 23, 162, 1, 0, 0, 0, 25, 164, 1, 0, 0, 0, 27, 172, 1, 0, 0, 0, 29, 178,
		1, 0, 0, 0, 31, 187, 1, 0, 0, 0, 33, 194, 1, 0, 0, 0, 35, 199, 1, 0, 0,
		0, 37, 201, 1, 0, 0, 0, 39, 205, 1, 0, 0, 0, 41, 207, 1, 0, 0, 0, 43, 214,
		1, 0, 0, 0, 45, 218, 1, 0, 0, 0, 47, 224, 1, 0, 0, 0, 49, 229, 1, 0, 0,
		0, 51, 231, 1, 0, 0, 0, 53, 233, 1, 0, 0, 0, 55, 235, 1, 0, 0, 0, 57, 241,
		1, 0, 0, 0, 59, 243, 1, 0, 0, 0, 61, 245, 1, 0, 0, 0, 63, 247, 1, 0, 0,
		0, 65, 249, 1, 0, 0, 0, 67, 251, 1, 0, 0, 0, 69, 253, 1, 0, 0, 0, 71, 255,
		1, 0, 0, 0, 73, 258, 1, 0, 0, 0, 75, 261, 1, 0, 0, 0, 77, 263, 1, 0, 0,
		0, 79, 265, 1, 0, 0, 0, 81, 268, 1, 0, 0, 0, 83, 271, 1, 0, 0, 0, 85, 273,
		1, 0, 0, 0, 87, 276, 1, 0, 0, 0, 89, 280, 1, 0, 0, 0, 91, 286, 1, 0, 0,
		0, 93, 301, 1, 0, 0, 0, 95, 316, 1, 0, 0, 0, 97, 321, 1, 0, 0, 0, 99, 333,
		1, 0, 0, 0, 101, 337, 1, 0, 0, 0, 103, 344, 1, 0, 0, 0, 105, 351, 1, 0,
		0, 0, 107, 362, 1, 0, 0, 0, 109, 364, 1, 0, 0, 0, 111, 369, 1, 0, 0, 0,
		113, 373, 1, 0, 0, 0, 115, 376, 1, 0, 0, 0, 117, 118, 5, 59, 0, 0, 118,
		2, 1, 0, 0, 0, 119, 120, 5, 123, 0, 0, 120, 4, 1, 0, 0, 0, 121, 122, 5,
		125, 0, 0, 122, 6, 1, 0, 0, 0, 123, 124, 5, 102, 0, 0, 124, 125, 5, 111,
		0, 0, 125, 126, 5, 114, 0, 0, 126, 8, 1, 0, 0, 0, 127, 128, 5, 105, 0,
		0, 128, 129, 5, 110, 0, 0, 129, 10, 1, 0, 0, 0, 130, 131, 5, 119, 0, 0,
		131, 132, 5, 104, 0, 0, 132, 133, 5, 105, 0, 0, 133, 134, 5, 108, 0, 0,
		134, 135, 5, 101, 0, 0, 135, 12, 1, 0, 0, 0, 136, 137, 5, 105, 0, 0, 137,
		138, 5, 102, 0, 0, 138, 14, 1, 0, 0, 0, 139, 140, 5, 101, 0, 0, 140, 141,
		5, 108, 0, 0, 141, 142, 5, 115, 0, 0, 142, 143, 5, 101, 0, 0, 143, 16,
		1, 0, 0, 0, 144, 145, 5, 103, 0, 0, 145, 146, 5, 117, 0, 0, 146, 147, 5,
		97, 0, 0, 147, 148, 5, 114, 0, 0, 148, 149, 5, 100, 0, 0, 149, 18, 1, 0,
		0, 0, 150, 151, 5, 115, 0, 0, 151, 152, 5, 119, 0, 0, 152, 153, 5, 105,
		0, 0, 153, 154, 5, 116, 0, 0, 154, 155, 5, 99, 0, 0, 155, 156, 5, 104,
		0, 0, 156, 20, 1, 0, 0, 0, 157, 158, 5, 99, 0, 0, 158, 159, 5, 97, 0, 0,
		159, 160, 5, 115, 0, 0, 160, 161, 5, 101, 0, 0, 161, 22, 1, 0, 0, 0, 162,
		163, 5, 58, 0, 0, 163, 24, 1, 0, 0, 0, 164, 165, 5, 100, 0, 0, 165, 166,
		5, 101, 0, 0, 166, 167, 5, 102, 0, 0, 167, 168, 5, 97, 0, 0, 168, 169,
		5, 117, 0, 0, 169, 170, 5, 108, 0, 0, 170, 171, 5, 116, 0, 0, 171, 26,
		1, 0, 0, 0, 172, 173, 5, 98, 0, 0, 173, 174, 5, 114, 0, 0, 174, 175, 5,
		101, 0, 0, 175, 176, 5, 97, 0, 0, 176, 177, 5, 107, 0, 0, 177, 28, 1, 0,
		0, 0, 178, 179, 5, 99, 0, 0, 179, 180, 5, 111, 0, 0, 180, 181, 5, 110,
		0, 0, 181, 182, 5, 116, 0, 0, 182, 183, 5, 105, 0, 0, 183, 184, 5, 110,
		0, 0, 184, 185, 5, 117, 0, 0, 185, 186, 5, 101, 0, 0, 186, 30, 1, 0, 0,
		0, 187, 188, 5, 114, 0, 0, 188, 189, 5, 101, 0, 0, 189, 190, 5, 116, 0,
		0, 190, 191, 5, 117, 0, 0, 191, 192, 5, 114, 0, 0, 192, 193, 5, 110, 0,
		0, 193, 32, 1, 0, 0, 0, 194, 195, 5, 108, 0, 0, 195, 196, 5, 101, 0, 0,
		196, 197, 5, 116, 0, 0, 197, 198, 5, 32, 0, 0, 198, 34, 1, 0, 0, 0, 199,
		200, 5, 61, 0, 0, 200, 36, 1, 0, 0, 0, 201, 202, 5, 118, 0, 0, 202, 203,
		5, 97, 0, 0, 203, 204, 5, 114, 0, 0, 204, 38, 1, 0, 0, 0, 205, 206, 5,
		63, 0, 0, 206, 40, 1, 0, 0, 0, 207, 208, 5, 83, 0, 0, 208, 209, 5, 116,
		0, 0, 209, 210, 5, 114, 0, 0, 210, 211, 5, 105, 0, 0, 211, 212, 5, 110,
		0, 0, 212, 213, 5, 103, 0, 0, 213, 42, 1, 0, 0, 0, 214, 215, 5, 73, 0,
		0, 215, 216, 5, 110, 0, 0, 216, 217, 5, 116, 0, 0, 217, 44, 1, 0, 0, 0,
		218, 219, 5, 70, 0, 0, 219, 220, 5, 108, 0, 0, 220, 221, 5, 111, 0, 0,
		221, 222, 5, 97, 0, 0, 222, 223, 5, 116, 0, 0, 223, 46, 1, 0, 0, 0, 224,
		225, 5, 66, 0, 0, 225, 226, 5, 111, 0, 0, 226, 227, 5, 111, 0, 0, 227,
		228, 5, 108, 0, 0, 228, 48, 1, 0, 0, 0, 229, 230, 5, 91, 0, 0, 230, 50,
		1, 0, 0, 0, 231, 232, 5, 93, 0, 0, 232, 52, 1, 0, 0, 0, 233, 234, 5, 44,
		0, 0, 234, 54, 1, 0, 0, 0, 235, 236, 5, 112, 0, 0, 236, 237, 5, 114, 0,
		0, 237, 238, 5, 105, 0, 0, 238, 239, 5, 110, 0, 0, 239, 240, 5, 116, 0,
		0, 240, 56, 1, 0, 0, 0, 241, 242, 5, 40, 0, 0, 242, 58, 1, 0, 0, 0, 243,
		244, 5, 41, 0, 0, 244, 60, 1, 0, 0, 0, 245, 246, 5, 45, 0, 0, 246, 62,
		1, 0, 0, 0, 247, 248, 5, 42, 0, 0, 248, 64, 1, 0, 0, 0, 249, 250, 5, 47,
		0, 0, 250, 66, 1, 0, 0, 0, 251, 252, 5, 37, 0, 0, 252, 68, 1, 0, 0, 0,
		253, 254, 5, 43, 0, 0, 254, 70, 1, 0, 0, 0, 255, 256, 5, 61, 0, 0, 256,
		257, 5, 61, 0, 0, 257, 72, 1, 0, 0, 0, 258, 259, 5, 33, 0, 0, 259, 260,
		5, 61, 0, 0, 260, 74, 1, 0, 0, 0, 261, 262, 5, 62, 0, 0, 262, 76, 1, 0,
		0, 0, 263, 264, 5, 60, 0, 0, 264, 78, 1, 0, 0, 0, 265, 266, 5, 60, 0, 0,
		266, 267, 5, 61, 0, 0, 267, 80, 1, 0, 0, 0, 268, 269, 5, 62, 0, 0, 269,
		270, 5, 61, 0, 0, 270, 82, 1, 0, 0, 0, 271, 272, 5, 33, 0, 0, 272, 84,
		1, 0, 0, 0, 273, 274, 5, 38, 0, 0, 274, 275, 5, 38, 0, 0, 275, 86, 1, 0,
		0, 0, 276, 277, 5, 124, 0, 0, 277, 278, 5, 124, 0, 0, 278, 88, 1, 0, 0,
		0, 279, 281, 7, 0, 0, 0, 280, 279, 1, 0, 0, 0, 281, 282, 1, 0, 0, 0, 282,
		280, 1, 0, 0, 0, 282, 283, 1, 0, 0, 0, 283, 284, 1, 0, 0, 0, 284, 285,
		6, 44, 0, 0, 285, 90, 1, 0, 0, 0, 286, 287, 5, 47, 0, 0, 287, 288, 5, 42,
		0, 0, 288, 293, 1, 0, 0, 0, 289, 292, 3, 91, 45, 0, 290, 292, 9, 0, 0,
		0, 291, 289, 1, 0, 0, 0, 291, 290, 1, 0, 0, 0, 292, 295, 1, 0, 0, 0, 293,
		294, 1, 0, 0, 0, 293, 291, 1, 0, 0, 0, 294, 296, 1, 0, 0, 0, 295, 293,
		1, 0, 0, 0, 296, 297, 5, 42, 0, 0, 297, 298, 5, 47, 0, 0, 298, 299, 1,
		0, 0, 0, 299, 300, 6, 45, 0, 0, 300, 92, 1, 0, 0, 0, 301, 302, 5, 47, 0,
		0, 302, 303, 5, 47, 0, 0, 303, 307, 1, 0, 0, 0, 304, 306, 9, 0, 0, 0, 305,
		304, 1, 0, 0, 0, 306, 309, 1, 0, 0, 0, 307, 308, 1, 0, 0, 0, 307, 305,
		1, 0, 0, 0, 308, 311, 1, 0, 0, 0, 309, 307, 1, 0, 0, 0, 310, 312, 7, 1,
		0, 0, 311, 310, 1, 0, 0, 0, 312, 313, 1, 0, 0, 0, 313, 314, 6, 46, 0, 0,
		314, 94, 1, 0, 0, 0, 315, 317, 7, 2, 0, 0, 316, 315, 1, 0, 0, 0, 317, 318,
		1, 0, 0, 0, 318, 316, 1, 0, 0, 0, 318, 319, 1, 0, 0, 0, 319, 96, 1, 0,
		0, 0, 320, 322, 7, 2, 0, 0, 321, 320, 1, 0, 0, 0, 322, 323, 1, 0, 0, 0,
		323, 321, 1, 0, 0, 0, 323, 324, 1, 0, 0, 0, 324, 331, 1, 0, 0, 0, 325,
		327, 5, 46, 0, 0, 326, 328, 7, 2, 0, 0, 327, 326, 1, 0, 0, 0, 328, 329,
		1, 0, 0, 0, 329, 327, 1, 0, 0, 0, 329, 330, 1, 0, 0, 0, 330, 332, 1, 0,
		0, 0, 331, 325, 1, 0, 0, 0, 331, 332, 1, 0, 0, 0, 332, 98, 1, 0, 0, 0,
		333, 334, 5, 34, 0, 0, 334, 335, 3, 105, 52, 0, 335, 336, 5, 34, 0, 0,
		336, 100, 1, 0, 0, 0, 337, 339, 5, 34, 0, 0, 338, 340, 3, 103, 51, 0, 339,
		338, 1, 0, 0, 0, 339, 340, 1, 0, 0, 0, 340, 341, 1, 0, 0, 0, 341, 342,
		5, 34, 0, 0, 342, 102, 1, 0, 0, 0, 343, 345, 3, 105, 52, 0, 344, 343, 1,
		0, 0, 0, 345, 346, 1, 0, 0, 0, 346, 344, 1, 0, 0, 0, 346, 347, 1, 0, 0,
		0, 347, 104, 1, 0, 0, 0, 348, 349, 5, 92, 0, 0, 349, 352, 7, 3, 0, 0, 350,
		352, 8, 4, 0, 0, 351, 348, 1, 0, 0, 0, 351, 350, 1, 0, 0, 0, 352, 106,
		1, 0, 0, 0, 353, 354, 5, 116, 0, 0, 354, 355, 5, 114, 0, 0, 355, 356, 5,
		117, 0, 0, 356, 363, 5, 101, 0, 0, 357, 358, 5, 102, 0, 0, 358, 359, 5,
		97, 0, 0, 359, 360, 5, 108, 0, 0, 360, 361, 5, 115, 0, 0, 361, 363, 5,
		101, 0, 0, 362, 353, 1, 0, 0, 0, 362, 357, 1, 0, 0, 0, 363, 108, 1, 0,
		0, 0, 364, 366, 3, 111, 55, 0, 365, 367, 3, 115, 57, 0, 366, 365, 1, 0,
		0, 0, 366, 367, 1, 0, 0, 0, 367, 110, 1, 0, 0, 0, 368, 370, 7, 5, 0, 0,
		369, 368, 1, 0, 0, 0, 370, 112, 1, 0, 0, 0, 371, 374, 7, 2, 0, 0, 372,
		374, 3, 111, 55, 0, 373, 371, 1, 0, 0, 0, 373, 372, 1, 0, 0, 0, 374, 114,
		1, 0, 0, 0, 375, 377, 3, 113, 56, 0, 376, 375, 1, 0, 0, 0, 377, 378, 1,
		0, 0, 0, 378, 376, 1, 0, 0, 0, 378, 379, 1, 0, 0, 0, 379, 116, 1, 0, 0,
		0, 18, 0, 282, 291, 293, 307, 311, 318, 323, 329, 331, 339, 346, 351, 362,
		366, 369, 373, 378, 1, 0, 1, 0,
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

// T_swiftLexerInit initializes any static state used to implement T_swiftLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewT_swiftLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func T_swiftLexerInit() {
	staticData := &T_swiftLexerLexerStaticData
	staticData.once.Do(t_swiftlexerLexerInit)
}

// NewT_swiftLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewT_swiftLexer(input antlr.CharStream) *T_swiftLexer {
	T_swiftLexerInit()
	l := new(T_swiftLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &T_swiftLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "T_swift.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// T_swiftLexer tokens.
const (
	T_swiftLexerT__0          = 1
	T_swiftLexerT__1          = 2
	T_swiftLexerT__2          = 3
	T_swiftLexerT__3          = 4
	T_swiftLexerT__4          = 5
	T_swiftLexerT__5          = 6
	T_swiftLexerT__6          = 7
	T_swiftLexerT__7          = 8
	T_swiftLexerT__8          = 9
	T_swiftLexerT__9          = 10
	T_swiftLexerT__10         = 11
	T_swiftLexerT__11         = 12
	T_swiftLexerT__12         = 13
	T_swiftLexerT__13         = 14
	T_swiftLexerT__14         = 15
	T_swiftLexerT__15         = 16
	T_swiftLexerT__16         = 17
	T_swiftLexerT__17         = 18
	T_swiftLexerT__18         = 19
	T_swiftLexerT__19         = 20
	T_swiftLexerT__20         = 21
	T_swiftLexerT__21         = 22
	T_swiftLexerT__22         = 23
	T_swiftLexerT__23         = 24
	T_swiftLexerT__24         = 25
	T_swiftLexerT__25         = 26
	T_swiftLexerT__26         = 27
	T_swiftLexerT__27         = 28
	T_swiftLexerT__28         = 29
	T_swiftLexerT__29         = 30
	T_swiftLexerT__30         = 31
	T_swiftLexerT__31         = 32
	T_swiftLexerT__32         = 33
	T_swiftLexerT__33         = 34
	T_swiftLexerT__34         = 35
	T_swiftLexerT__35         = 36
	T_swiftLexerT__36         = 37
	T_swiftLexerT__37         = 38
	T_swiftLexerT__38         = 39
	T_swiftLexerT__39         = 40
	T_swiftLexerT__40         = 41
	T_swiftLexerT__41         = 42
	T_swiftLexerT__42         = 43
	T_swiftLexerT__43         = 44
	T_swiftLexerWS            = 45
	T_swiftLexerBlock_comment = 46
	T_swiftLexerLine_comment  = 47
	T_swiftLexerInt           = 48
	T_swiftLexerFloat         = 49
	T_swiftLexerCaracter      = 50
	T_swiftLexerString_       = 51
	T_swiftLexerBool          = 52
	T_swiftLexerIdentificador = 53
)